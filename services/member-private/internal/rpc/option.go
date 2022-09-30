package rpc

import (
	"context"
	"errors"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/member-private/internal/model"
	"gitlab.omytech.com.cn/micro-service/member-private/proto"
	merchantBasic "gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// ListPromotionOption 列表
func (s *Server) ListPromotionOption(ctx context.Context, req *proto.ListPromotionOptionRequest) (*proto.ListPromotionOptionResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ListPromotionOption")

	scopes := s.listPromotionOptionScopes(req)

	var total int64
	var options []model.Option

	s.database.Conn.Model(&model.Option{}).Scopes(scopes...).Count(&total)
	if req.WithPage {
		scopes = append(scopes, util.PaginationScope(req.Offset, req.Limit))
	}
	orderBy := "created_at desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}
	if err := s.database.Conn.Model(&model.Option{}).Scopes(scopes...).Order(orderBy).Find(&options).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("查询优惠方案列表错误")
		return &proto.ListPromotionOptionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据查询错误:%s", err.Error()),
		}, nil
	}

	data := make([]*proto.PromotionOption, 0)
	for i := range options {
		option := toProtoPromotionOption(options[i])
		// 当有门店时，需要获取门店商品最新价格
		if req.BranchId != "" {
			productIDs := make([]string, 0)
			packageIDs := make([]string, 0)
			for _, v := range option.Products {
				productIDs = append(productIDs, v.Id)
			}
			for _, v := range option.Packages {
				packageIDs = append(packageIDs, v.Id)
			}
			resp, err := s.merchantBasic().MultiGetGoodsAndPackages(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &merchantBasic.MultiGetGoodsAndPackagesRequest{
				GoodsIds:   productIDs,
				PackageIds: packageIDs,
				BranchId:   req.BranchId,
			})
			if err != nil || resp.ErrorCode != pkgs.Success {
				util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error(fmt.Sprintf("查询商品套餐错误:%v", err))
				return &proto.ListPromotionOptionResponse{
					ErrorCode:    pkgs.ErrInternal,
					ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
				}, err
			}
			m := make(map[string]merchantBasic.GoodsAndPackageItem, 0)
			for _, v := range resp.Data.Goods {
				m[v.Id] = *v
			}
			for _, v := range resp.Data.Packages {
				m[v.Id] = *v
			}
			for j := range option.Products {
				if p, ok := m[option.Products[j].Id]; ok {
					option.Products[j].Code = p.PosCode
					option.Products[j].Title = p.Name
					option.Products[j].Price = p.Price
				}
			}
			for j := range option.Packages {
				if p, ok := m[option.Packages[j].Id]; ok {
					option.Packages[j].Code = p.PosCode
					option.Packages[j].Title = p.Name
					option.Packages[j].Price = p.Price
				}
			}
		}
		data = append(data, option)
	}

	return &proto.ListPromotionOptionResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ListPromotionOptionResponse_ListPromotionOptionData{
			Data:  data,
			Total: int32(total),
		},
	}, nil
}

// listPromotionOptionScopes 组装请求scope
func (s *Server) listPromotionOptionScopes(req *proto.ListPromotionOptionRequest) []func(db *gorm.DB) *gorm.DB {
	var scopes []func(db *gorm.DB) *gorm.DB
	if len(req.PromotionId) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("promotion_id", req.PromotionId))
	}
	if len(req.Name) > 0 {
		scopes = append(scopes, util.ColumnLikeScope("name", req.Name))
	}

	if len(req.Status) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("status", req.Status))
	}

	if len(req.Ids) > 0 {
		scopes = append(scopes, util.ColumnInScope("id", makeInterfaceSlice(req.Ids)))
	}

	if len(req.PromotionIds) > 0 {
		scopes = append(scopes, util.ColumnInScope("promotion_id", makeInterfaceSlice(req.PromotionIds)))
	}

	if len(req.TagId) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("tag_id", req.TagId))
	}

	if len(req.TagIds) > 0 {
		scopes = append(scopes, util.ColumnInScope("tag_id", makeInterfaceSlice(req.TagIds)))
	}

	//是否过滤优惠方案组条件
	var promotionScopes []func(db *gorm.DB) *gorm.DB
	if req.FilterPromotion != nil {
		if len(req.FilterPromotion.Status) > 0 {
			promotionScopes = append(promotionScopes, util.ColumnEqualScope("status", req.FilterPromotion.Status))
		}
		if req.FilterPromotion.Begin > 0 {
			promotionScopes = append(promotionScopes, util.ColumnSymbolScope("begin_at", "<=", fields.UnixToDateTime(req.FilterPromotion.Begin)))
		}
		if req.FilterPromotion.End > 0 {
			promotionScopes = append(promotionScopes, util.ColumnSymbolScope("end_at", ">=", fields.UnixToDateTime(req.FilterPromotion.End)))
		}
		if len(req.FilterPromotion.BranchId) > 0 {
			promotionScopes = append(promotionScopes, util.ArrayAnyScope("branch_ids", req.FilterPromotion.BranchId))
		}
	}
	if len(promotionScopes) > 0 {
		var promotionIDs []uuid.UUID
		s.database.Conn.Model(&model.Promotion{}).Scopes(promotionScopes...).Pluck("id", &promotionIDs)
		scopes = append(scopes, util.ColumnInScope("promotion_id", makeUUIDInterfaceSlice(promotionIDs)))
	}

	return scopes
}

// CreatePromotionOption 创建优惠方案
func (s *Server) CreatePromotionOption(ctx context.Context, req *proto.CreatePromotionOptionRequest) (*proto.OptionResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreatePromotionOption")

	if len(req.PromotionId) == 0 || len(req.Name) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("新增优惠方案参数错误")
		return &proto.OptionResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	option := model.Option{
		ID:            uuid.NewV4(),
		Name:          req.Name,
		MerchantID:    metadata.MerchantID,
		PromotionID:   uuid.FromStringOrNil(req.PromotionId),
		RechargeValue: req.RechargeValue,
		BaseValue:     req.BaseValue,
		GiftValue:     req.GiftValue,
		Describe:      req.Describe,
		Products:      makeParams(req.Products),
		Packages:      makeParams(req.Packages),
		Tickets:       makeParams(req.Tickets),
	}

	// 名字去重
	if len(req.Status) == 0 {
		option.Status = util.StatusOpened
	} else {
		option.Status = util.StringToStatus(req.Status)
	}

	if len(req.TagId) > 0 {
		tagID := uuid.FromStringOrNil(req.TagId)
		option.TagID = &tagID
	}

	if err := s.database.Conn.Create(&option).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("option", logger.MakeFields(option)).WithError(err).Info("创建优惠方案错误")
		return &proto.OptionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("新增优惠方案数据库操作错误:%s", err.Error()),
		}, nil
	}

	after := pkgs.MakeParams(option)
	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: option.TableName(),
		TableID:           option.ID,
		Method:            model.CreateMethod,
		After:             &after,
	})

	return &proto.OptionResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "创建成功",
	}, nil
}

// ShowPromotionOption 查询优惠方案详情
func (s *Server) ShowPromotionOption(ctx context.Context, req *proto.ShowPromotionOptionRequest) (*proto.ShowPromotionOptionResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowPromotionOption")
	if len(req.OptionId) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("查看优惠方案详情,参数错误")
		return &proto.ShowPromotionOptionResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	var option model.Option
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.OptionId)).First(&option).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.ShowPromotionOptionResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.ShowPromotionOptionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, err
	}

	return &proto.ShowPromotionOptionResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         toProtoPromotionOption(option),
	}, nil
}

func toProtoPromotionOption(option model.Option) *proto.PromotionOption {
	var products, packages, tickets []*proto.ProductPackage
	if option.Products != nil {
		for _, v := range *option.Products {
			products = append(products, &proto.ProductPackage{
				Id:     v.ID.String(),
				Code:   v.Code,
				Number: v.Count,
				Price:  v.Price,
				Title:  v.Title,
				Unit:   v.Unit,
			})
		}
	}
	if option.Packages != nil {
		for _, v := range *option.Packages {
			packages = append(packages, &proto.ProductPackage{
				Id:     v.ID.String(),
				Code:   v.Code,
				Number: v.Count,
				Price:  v.Price,
				Title:  v.Title,
				Unit:   v.Unit,
			})
		}
	}
	if option.Tickets != nil {
		for _, v := range *option.Tickets {
			tickets = append(tickets, &proto.ProductPackage{
				Id:     v.ID.String(),
				Code:   v.Code,
				Number: v.Count,
				Price:  v.Price,
				Title:  v.Title,
				Unit:   v.Unit,
			})
		}
	}
	res := &proto.PromotionOption{
		Id:            option.ID.String(),
		PromotionId:   option.PromotionID.String(),
		Name:          option.Name,
		Status:        option.Status.String(),
		RechargeValue: option.RechargeValue,
		BaseValue:     option.BaseValue,
		GiftValue:     option.GiftValue,
		Describe:      option.Describe,
		TagId:         option.TagID.String(),
		CreatedAt:     option.CreatedAt.Time.Unix(),
		Products:      products,
		Packages:      packages,
		Tickets:       tickets,
	}

	if option.TagID != nil {
		res.TagId = option.TagID.String()
	}

	return res
}

// UpdatePromotionOption 修改优惠方案
func (s *Server) UpdatePromotionOption(ctx context.Context, req *proto.UpdatePromotionOptionRequest) (*proto.OptionResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdatePromotionOption")
	if len(req.Id) == 0 || len(req.Name) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("修改优惠方案参数错误")
		return &proto.OptionResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	var option model.Option
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&option).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.OptionResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.OptionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, err
	}

	update := map[string]interface{}{
		"name":           req.Name,
		"recharge_value": req.RechargeValue,
		"base_value":     req.BaseValue,
		"gift_value":     req.GiftValue,
		"describe":       req.Describe,
		"products":       makeParams(req.Products),
		"packages":       makeParams(req.Packages),
		"tickets":        makeParams(req.Tickets),
	}

	if len(req.Status) > 0 {
		update["status"] = util.StringToStatus(req.Status)
	}

	if len(req.TagId) == 0 {
		update["tag_id"] = nil
	} else {
		tagID := uuid.FromStringOrNil(req.TagId)
		update["tag_id"] = &tagID
	}

	if err := s.database.Conn.Model(&model.Option{}).Where("id = ?", option.ID).Updates(update).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("update", logger.MakeFields(update)).WithError(err).Error("修改优惠方案失败")
		return &proto.OptionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}
	before := pkgs.MakeParams(option)
	// 暂时用update当作信息
	after := pkgs.MakeParams(update)
	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: option.TableName(),
		TableID:           option.ID,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
	})

	return &proto.OptionResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "更新成功",
	}, nil
}

// UpdatePromotionOptionStatus 修改优惠方案
func (s *Server) UpdatePromotionOptionStatus(ctx context.Context, req *proto.UpdateStatusRequest) (*proto.UpdateStatusResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdatePromotionOption")
	if len(req.Id) == 0 || len(req.Status) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("修改优惠方案状态参数错误")
		return &proto.UpdateStatusResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	var option model.Option
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&option).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.UpdateStatusResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.UpdateStatusResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, err
	}

	update := map[string]interface{}{
		"status": util.StringToStatus(req.Status),
	}

	if err := s.database.Conn.Model(&model.Option{}).Where("id = ?", option.ID).Updates(update).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("update", logger.MakeFields(update)).WithError(err).Error("修改优惠方案状态失败")
		return &proto.UpdateStatusResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}
	before := pkgs.MakeParams(option)
	// 暂时用update当作信息
	after := pkgs.MakeParams(update)
	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: option.TableName(),
		TableID:           option.ID,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
	})

	return &proto.UpdateStatusResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "更新成功",
	}, nil
}

func makeInterfaceSlice(items []string) []interface{} {
	var res []interface{}
	for _, item := range items {
		res = append(res, item)
	}

	return res
}

func makeUUIDInterfaceSlice(items []uuid.UUID) []interface{} {
	var res []interface{}
	for _, item := range items {
		res = append(res, item.String())
	}

	return res
}

func makeParams(pp []*proto.ProductPackage) *model.Params {
	if len(pp) == 0 {
		return nil
	}
	params := new(model.Params)
	for _, v := range pp {
		*params = append(*params, model.Param{
			ID:    uuid.FromStringOrNil(v.Id),
			Code:  v.Code,
			Title: v.Title,
			Count: v.Number,
			Unit:  v.Unit,
		})
	}
	return params
}
