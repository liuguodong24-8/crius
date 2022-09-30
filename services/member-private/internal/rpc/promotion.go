package rpc

import (
	"context"
	"fmt"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/member-private/internal/model"
	"gitlab.omytech.com.cn/micro-service/member-private/proto"
	"gorm.io/gorm"
)

// CreatePromotion 创建优惠方案组
func (s *Server) CreatePromotion(ctx context.Context, req *proto.CreatePromotionRequest) (*proto.CreatePromotionResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreatePromotion")
	metadata := pkgs.GetMetadata(ctx)
	today, _ := time.ParseInLocation(fields.DateFormat, time.Now().Format(fields.DateFormat), time.Local)
	if len(metadata.MerchantID.String()) == 0 ||
		len(req.Name) == 0 || req.EndAt <= req.BeginAt ||
		time.Unix(req.BeginAt, 0).Before(today) {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("新增优惠方案组参数错误")
		return &proto.CreatePromotionResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	promotion := model.Promotion{
		ID:         uuid.NewV4(),
		Name:       req.Name,
		MerchantID: metadata.MerchantID,
	}

	var count int64
	if err := s.database.Conn.Model(&model.Promotion{}).Scopes(util.ColumnEqualScope("merchant_id", promotion.MerchantID), util.ColumnEqualScope("name", promotion.Name)).Count(&count).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("新增优惠方案组,数据库过滤名字唯一查询错误")
		return &proto.CreatePromotionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}

	if count > 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("新增优惠方案组,数据库过滤名字唯一错误")
		return &proto.CreatePromotionResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: fmt.Sprintf("新增优惠方案组，名字重复:%s", req.Name),
		}, nil
	}

	if req.BeginAt > 0 {
		begin := fields.UnixToDateTime(req.BeginAt)
		promotion.BeginAt = &begin
	}
	if req.EndAt > 0 {
		end := fields.UnixToDateTime(req.EndAt)
		promotion.EndAt = &end
	}
	if len(req.Status) == 0 {
		promotion.Status = util.StatusOpened
	} else {
		promotion.Status = util.StringToStatus(req.Status)
	}

	if len(req.BranchIds) > 0 {
		if i, e := fields.StringArrToUUIDArr(req.BranchIds); nil == e {
			promotion.BranchIds = i
		}
	}

	if err := s.database.Conn.Create(&promotion).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("promotion", logger.MakeFields(promotion)).WithError(err).Error("新增优惠方案组,保存数据库记录错误")
		return &proto.CreatePromotionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, nil
	}

	after := pkgs.MakeParams(promotion)
	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: promotion.TableName(),
		TableID:           promotion.ID,
		Method:            model.CreateMethod,
		After:             &after,
	})

	return &proto.CreatePromotionResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         toProtoPromotion(promotion),
	}, nil
}

// UpdatePromotion 更新优惠方案
func (s *Server) UpdatePromotion(ctx context.Context, req *proto.UpdatePromotionRequest) (*proto.PromotionResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdatePromotion")

	if len(req.Id) == 0 || len(req.Name) == 0 ||
		req.EndAt <= req.BeginAt {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改优惠方案组,参数错误")
		return &proto.PromotionResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "修改参数错误",
		}, nil
	}

	// 查询记录是否存在
	var promotion model.Promotion
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&promotion).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.PromotionResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.PromotionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, err
	}

	// 名字判断
	var count int64
	if err := s.database.Conn.Model(&model.Promotion{}).Scopes(
		util.ColumnEqualScope("merchant_id", promotion.MerchantID),
		util.ColumnEqualScope("name", req.Name),
		util.ColumnNotEqualScope("id", promotion.ID)).Count(&count).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改优惠方案组,数据库过滤名字唯一错误")
		return &proto.PromotionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, nil
	}

	update := map[string]interface{}{
		"name": req.Name,
	}

	if len(req.Status) > 0 {
		update["status"] = util.StringToStatus(req.Status)
	}

	if req.BeginAt > 0 {
		begin := fields.UnixToDateTime(req.BeginAt)
		update["begin_at"] = &begin
	} else {
		update["begin_at"] = nil
	}
	if req.EndAt > 0 {
		end := fields.UnixToDateTime(req.EndAt)
		update["end_at"] = &end
	} else {
		update["end_at"] = nil
	}

	if len(req.BranchIds) > 0 {
		update["branch_ids"], _ = fields.StringArrToUUIDArr(req.BranchIds)
	} else {
		update["branch_ids"], _ = fields.StringArrToUUIDArr([]string{})
	}

	if err := s.database.Conn.Model(&model.Promotion{}).Where("id = ?", promotion.ID).Updates(update).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(update)).WithError(err).Error("修改优惠方案组失败")
		return &proto.PromotionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}

	before := pkgs.MakeParams(promotion)
	// 暂时用update信息
	after := pkgs.MakeParams(update)
	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: promotion.TableName(),
		TableID:           promotion.ID,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
	})

	return &proto.PromotionResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "更新成功",
	}, nil
}

// ListPromotion 查询优惠方案组列表
func (s *Server) ListPromotion(ctx context.Context, req *proto.ListPromotionRequest) (*proto.ListPromotionResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ListPromotion")

	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))
	if len(req.Name) > 0 {
		scopes = append(scopes, util.ColumnLikeScope("name", req.Name))
	}

	// 非管理员 查询账户权限门店数据
	if metadata.IsAdmin == 0 {
		var ids []interface{}
		for _, v := range metadata.BranchIDs {
			ids = append(ids, v.String())
		}
		scopes = append(scopes, util.ArrayOverlapScope("branch_ids", "uuid", ids))
	}

	if len(req.BranchId) > 0 {
		scopes = append(scopes, util.ArrayAnyScope("branch_ids", req.BranchId))
	}

	if len(req.Status) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("status", req.Status))
	}

	var total int64
	var promotions []model.Promotion

	s.database.Conn.Model(&model.Promotion{}).Scopes(scopes...).Count(&total)
	// 是否需要分页
	if req.WithPage {
		scopes = append(scopes, util.PaginationScope(req.Offset, req.Limit))
	}
	orderBy := "created_at desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}
	if err := s.database.Conn.Model(&model.Promotion{}).Scopes(scopes...).Order(orderBy).Find(&promotions).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("查询优惠方案组列表错误")
		return &proto.ListPromotionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据查询错误:%s", err.Error()),
		}, nil
	}

	data := make([]*proto.Promotion, 0)
	for _, v := range promotions {
		data = append(data, toProtoPromotion(v))
	}

	return &proto.ListPromotionResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ListPromotionResponse_ListPromotionData{
			Data:  data,
			Total: int32(total),
		},
	}, nil
}

// ShowPromotion 查询优惠方案组详情
func (s *Server) ShowPromotion(ctx context.Context, req *proto.ShowPromotionRequest) (*proto.ShowPromotionResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowPromotion")
	if len(req.PromotionId) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("查看优惠方案组详情,参数错误")
		return &proto.ShowPromotionResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	// 查询记录是否存在
	var promotion model.Promotion
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.PromotionId)).First(&promotion).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.ShowPromotionResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.ShowPromotionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, err
	}

	return &proto.ShowPromotionResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         toProtoPromotion(promotion),
	}, nil
}

// UpdatePromotionStatus 修改状态
func (s *Server) UpdatePromotionStatus(ctx context.Context, req *proto.UpdateStatusRequest) (*proto.UpdateStatusResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdatePromotionStatus")

	if len(req.Id) == 0 || len(req.Status) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改优惠方案组状态,参数错误")
		return &proto.UpdateStatusResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "修改参数错误",
		}, nil
	}

	// 查询记录是否存在
	var promotion model.Promotion
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&promotion).Error; nil != err {
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

	if err := s.database.Conn.Model(&model.Promotion{}).Where("id = ?", promotion.ID).Updates(update).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(update)).WithError(err).Error("修改优惠方案组状态失败")
		return &proto.UpdateStatusResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}

	before := pkgs.MakeParams(promotion)
	// 暂时用update信息
	after := pkgs.MakeParams(update)
	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: promotion.TableName(),
		TableID:           promotion.ID,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
	})

	return &proto.UpdateStatusResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "更新成功",
	}, nil
}

func toProtoPromotion(promotion model.Promotion) *proto.Promotion {
	return &proto.Promotion{
		Id:        promotion.ID.String(),
		Name:      promotion.Name,
		BeginAt:   promotion.BeginAt.ToUnix(),
		EndAt:     promotion.EndAt.ToUnix(),
		Status:    promotion.Status.String(),
		CreatedAt: promotion.CreatedAt.Time.Unix(),
		BranchIds: promotion.BranchIds.ToStringArr(),
	}
}
