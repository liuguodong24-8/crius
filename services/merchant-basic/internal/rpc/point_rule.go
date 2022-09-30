package rpc

import (
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// SetPointRuleDescribe 设置积分规则说明
func (s *Server) SetPointRuleDescribe(ctx context.Context, req *proto.SetPointRuleDescribeRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("SetPointRuleDescribe")

	metadata := pkgs.GetMetadata(ctx)
	var describe model.PointRuleDescribe
	snapshot := model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &metadata.StaffID,
		SleuthCode:        metadata.SleuthCode,
		SnapShotTableName: describe.TableName(),
		TableID:           nil,
		Method:            "update",
		Before:            nil,
		After: &pkgs.Params{
			"graphic_detail": req.Images,
		},
	}

	err := model.DatabaseConn().Model(&model.PointRuleDescribe{}).Scopes(util.ColumnEqualScope("merchant_id", metadata.MerchantID)).First(&describe).Error
	// 数据库查询错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询积分规则说明数据库错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询数据库错误:%s", err.Error()),
		}, nil
	}

	// 不存在 创建
	if errors.Is(err, gorm.ErrRecordNotFound) {
		describe = model.PointRuleDescribe{
			ID:            uuid.NewV4(),
			MerchantID:    &metadata.MerchantID,
			GraphicDetail: (*fields.StringArr)(&req.Images),
		}

		if e := model.DatabaseConn().Create(&describe).Error; nil != e {
			util.Logger.WithMetadata(ctx).WithError(e).Error("新增积分规则说明错误")
			return &proto.Response{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: fmt.Sprintf("保存数据库错误:%s", e.Error()),
			}, nil
		}
	}

	// 修改
	if err == nil {
		if e := model.DatabaseConn().Model(&describe).Updates(model.PointRuleDescribe{GraphicDetail: (*fields.StringArr)(&req.Images)}).Error; nil != e {
			util.Logger.WithMetadata(ctx).WithError(e).Error("修改积分规则说明错误")
			return &proto.Response{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: fmt.Sprintf("保存数据库错误:%s", e.Error()),
			}, nil
		}

		before := pkgs.MakeParams(describe)
		snapshot.Before = &before

	}

	snapshot.TableID = &describe.ID
	model.SnapShotChan <- &snapshot

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// GetPointRuleDescribe 获取积分规则说明
func (s *Server) GetPointRuleDescribe(ctx context.Context, req *proto.Empty) (*proto.GetPointRuleDescribeResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("GetPointRuleDescribe")

	metadata := pkgs.GetMetadata(ctx)
	var describe model.PointRuleDescribe

	if err := model.DatabaseConn().Model(&model.PointRuleDescribe{}).Scopes(util.ColumnEqualScope("merchant_id", metadata.MerchantID)).First(&describe).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.GetPointRuleDescribeResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: fmt.Sprintf("%s not found", metadata.MerchantID.String()),
			}, nil
		}

		return &proto.GetPointRuleDescribeResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询数据库错误:%s", err.Error()),
		}, nil
	}

	return &proto.GetPointRuleDescribeResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.GetPointRuleDescribeResponse_Data{
			Images: describe.GraphicDetail.Slice(),
		},
	}, nil
}

// CreatePointRule 创建积分规则
func (s *Server) CreatePointRule(ctx context.Context, req *proto.CreatePointRuleRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("CreatePointRule")

	exists, err := s.checkPointRule(ctx, req.RuleName, "")
	if !exists && err != nil {
		util.Logger.WithError(err).Error("新增积分规则，校验错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("新增积分规则，校验错误:%s", err.Error()),
		}, nil
	}

	if exists && err != nil {
		return &proto.Response{
			ErrorCode:    pkgs.ErrCheck,
			ErrorMessage: err.Error(),
		}, nil
	}

	metadata := pkgs.GetMetadata(ctx)
	branchIDs, _ := fields.StringArrToUUIDArr(req.BranchIds)
	rule := model.PointRule{
		ID:          uuid.NewV4(),
		MerchantID:  &metadata.MerchantID,
		RuleName:    req.RuleName,
		GainRules:   protoRuleDetail(req.GainRules),
		UseRules:    protoRuleDetail(req.UseRules),
		ValidityDay: req.ValidityDay,
		BranchIDs:   &branchIDs,
		Status:      util.StringToStatus(req.Status),
	}

	if err := model.DatabaseConn().Create(&rule).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("保存积分规则错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("保存数据错误:%s", err.Error()),
		}, nil
	}

	after := pkgs.MakeParams(rule)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &metadata.StaffID,
		SleuthCode:        metadata.SleuthCode,
		SnapShotTableName: rule.TableName(),
		After:             &after,
		TableID:           &rule.ID,
		Method:            "create",
	}

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// UpdatePointRule 修改积分规则
func (s *Server) UpdatePointRule(ctx context.Context, req *proto.UpdatePointRuleRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("UpdatePointRule")

	var rule model.PointRule
	if err := model.DatabaseConn().Model(&model.PointRule{}).Scopes(util.ColumnEqualScope("id", req.Id)).First(&rule).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.Response{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: fmt.Sprintf("%s not found", req.Id),
			}, nil
		}

		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询数据错误:%s", err.Error()),
		}, nil
	}

	exists, err := s.checkPointRule(ctx, req.RuleName, rule.ID.String())
	if !exists && err != nil {
		util.Logger.WithError(err).Error("修改积分规则，校验错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("修改积分规则，校验错误:%s", err.Error()),
		}, nil
	}

	if exists && err != nil {
		return &proto.Response{
			ErrorCode:    pkgs.ErrCheck,
			ErrorMessage: err.Error(),
		}, nil
	}

	before := pkgs.MakeParams(rule)
	rule.RuleName = req.RuleName
	rule.GainRules = protoRuleDetail(req.GainRules)
	rule.UseRules = protoRuleDetail(req.UseRules)
	rule.ValidityDay = req.ValidityDay
	rule.Status = util.StringToStatus(req.Status)
	branchIDs, _ := fields.StringArrToUUIDArr(req.BranchIds)
	rule.BranchIDs = &branchIDs

	if err := model.DatabaseConn().Save(&rule).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("修改保存积分规则错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("更新保存数据库错误"),
		}, nil
	}
	after := pkgs.MakeParams(req)
	metadata := pkgs.GetMetadata(ctx)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &metadata.StaffID,
		SleuthCode:        metadata.SleuthCode,
		SnapShotTableName: rule.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &rule.ID,
		Method:            "update",
	}

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// UpdatePointRuleStatus 修改积分规则状态
func (s *Server) UpdatePointRuleStatus(ctx context.Context, req *proto.UpdateStatusRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("UpdatePointRuleStatus")

	var rule model.PointRule
	if err := model.DatabaseConn().Model(&model.PointRule{}).Scopes(util.ColumnEqualScope("id", req.Id)).First(&rule).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.Response{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: fmt.Sprintf("%s not found", req.Id),
			}, nil
		}

		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询数据错误:%s", err.Error()),
		}, nil
	}

	before := pkgs.MakeParams(rule)
	rule.Status = util.StringToStatus(req.Status)
	after := pkgs.MakeParams(req)

	if err := model.DatabaseConn().Save(&rule).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("修改积分规则状态错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("更新保存数据库错误"),
		}, nil
	}

	metadata := pkgs.GetMetadata(ctx)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &metadata.StaffID,
		SleuthCode:        metadata.SleuthCode,
		SnapShotTableName: rule.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &rule.ID,
		Method:            "update",
	}

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// ListPointRule 积分规则列表
func (s *Server) ListPointRule(ctx context.Context, req *proto.ListPointRuleRequest) (*proto.ListPointRuleResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("ListPointRule")

	metadata := pkgs.GetMetadata(ctx)
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))
	if len(req.RuleName) > 0 {
		scopes = append(scopes, util.ColumnLikeScope("rule_name", req.RuleName))
	}
	if len(req.Status) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("status", req.Status))
	}

	var total int64
	var items []model.PointRule

	model.DatabaseConn().Model(&model.PointRule{}).Scopes(scopes...).Count(&total)
	if req.WithPage {
		scopes = append(scopes, util.PaginationScope(req.Offset, req.Limit))
	}

	orderBy := "created_at desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}

	if err := model.DatabaseConn().Model(&model.PointRule{}).Scopes(scopes...).Order(orderBy).Find(&items).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询积分规则列表错误")
		return &proto.ListPointRuleResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询数据错误:%s", err.Error()),
		}, nil
	}

	var data []*proto.PointRule
	for _, item := range items {
		data = append(data, toProtoPointRule(item))
	}

	return &proto.ListPointRuleResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ListPointRuleResponse_Data{
			Data:  data,
			Total: total,
		},
	}, nil
}

// ShowPointRule 积分规则详情
func (s *Server) ShowPointRule(ctx context.Context, req *proto.ShowPointRuleRequest) (*proto.ShowPointRuleResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("ShowPointRule")

	var rule model.PointRule
	var scopes []func(db *gorm.DB) *gorm.DB
	if len(req.Id) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("id", req.Id))
	}
	if len(req.BranchId) > 0 {
		scopes = append(scopes, util.ArrayAnyScope("branch_ids", req.BranchId))
	}
	if len(scopes) == 0 {
		return &proto.ShowPointRuleResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "查询详情，参数错误",
		}, nil
	}

	if err := model.DatabaseConn().Model(&model.PointRule{}).Scopes(scopes...).Order("updated_at desc").First(&rule).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.ShowPointRuleResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: fmt.Sprintf("%s not found", req.Id),
			}, nil
		}

		return &proto.ShowPointRuleResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询数据错误:%s", err.Error()),
		}, nil
	}

	return &proto.ShowPointRuleResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         toProtoPointRule(rule),
	}, nil
}

// GetPointRuleAllBranch 获取所有已设置积分规则门店
func (s *Server) GetPointRuleAllBranch(ctx context.Context, req *proto.Empty) (*proto.GetPointRuleAllBranchResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("GetPointRuleAllBranch")

	metadata := pkgs.GetMetadata(ctx)

	var items []model.PointRule

	if err := model.DatabaseConn().Model(&model.PointRule{}).Scopes(util.ColumnEqualScope("merchant_id", metadata.MerchantID)).Find(&items).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询积分规则门店错误")
		return &proto.GetPointRuleAllBranchResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询数据错误:%s", err.Error()),
		}, nil
	}

	idsMap := make(map[string]string)
	var ids []string
	for _, item := range items {
		for _, v := range item.BranchIDs.ToStringArr() {
			if _, exists := idsMap[v]; exists {
				continue
			}
			ids = append(ids, v)
			idsMap[v] = v
		}
	}

	return &proto.GetPointRuleAllBranchResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         &proto.GetPointRuleAllBranchResponse_Data{BranchIds: ids},
	}, nil
}

func protoRuleDetail(rules []*proto.PointRuleDetail) *model.Rules {
	var res model.Rules

	for _, rule := range rules {
		res = append(res, model.Rule{
			CategoryID: uuid.FromStringOrNil(rule.CategoryId),
			Point:      rule.Point,
			Fee:        rule.Fee,
		})
	}

	return &res
}

func toProtoPointRule(rule model.PointRule) *proto.PointRule {
	return &proto.PointRule{
		Id:          rule.ID.String(),
		RuleName:    rule.RuleName,
		GainRules:   toProtoPointRuleDetail(rule.GainRules),
		UseRules:    toProtoPointRuleDetail(rule.UseRules),
		ValidityDay: rule.ValidityDay,
		BranchIds:   rule.BranchIDs.ToStringArr(),
		Status:      rule.Status.String(),
		CreatedAt:   rule.CreatedAt.Unix(),
	}
}

func toProtoPointRuleDetail(rules *model.Rules) []*proto.PointRuleDetail {
	if rules == nil {
		return nil
	}

	var res []*proto.PointRuleDetail

	for _, rule := range *rules {
		res = append(res, &proto.PointRuleDetail{
			CategoryId: rule.CategoryID.String(),
			Point:      rule.Point,
			Fee:        rule.Fee,
		})
	}

	return res
}

// GetBranchPointRule 获取门店积分规则
func (s *Server) GetBranchPointRule(ctx context.Context, req *proto.GetBranchPointRuleRequest) (*proto.GetBranchPointRuleResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("GetBranchPointRule")

	gain, use, err := model.GetBranchPointCategoryRule(model.GetBranchPointCategoryRuleRequest{
		MerchantID:   pkgs.GetMetadata(ctx).MerchantID.String(),
		BranchID:     req.BranchId,
		CategoryCode: req.CategoryCode,
	})

	// 没有门店积分规则
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &proto.GetBranchPointRuleResponse{
			ErrorCode:    pkgs.ErrNotFound,
			ErrorMessage: "",
		}, nil
	}

	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("获取门店积分规则错误")
		return &proto.GetBranchPointRuleResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("获取门店积分规则失败:%s", err.Error()),
		}, nil
	}

	gainRule := &proto.GetBranchPointRuleResponse_Rule{}
	if gain != nil {
		gainRule.Point = gain.Point
		gainRule.Fee = gain.Fee
		gainRule.ValidityDay = gain.ValidityDay
	}
	useRule := &proto.GetBranchPointRuleResponse_Rule{}
	if use != nil {
		useRule.Point = use.Point
		useRule.Fee = use.Fee
		useRule.ValidityDay = use.ValidityDay
	}

	return &proto.GetBranchPointRuleResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.GetBranchPointRuleResponse_Data{
			GainRule: gainRule,
			UseRule:  useRule,
		},
	}, nil
}

// GetBranchPointRules 门店积分规则（全量）
func (s *Server) GetBranchPointRules(ctx context.Context, req *proto.GetBranchPointRulesRequest) (*proto.GetBranchPointRulesResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("GetBranchPointRules")
	resp := &proto.GetBranchPointRulesResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	branchID := uuid.FromStringOrNil(req.BranchId)

	rules, err := model.GetBranchPointRules(branchID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetBranchPointRules 获取门店积分规则数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取门店积分规则错误"
		return resp, nil
	}

	var protoData proto.GetBranchPointRulesResponse_Data
	if len(rules) > 0 {
		//所有消费类型的映射
		categories, err := model.GetConsumeCategoryTypes(merchantID)
		if err != nil {
			util.Logger.WithMetadata(ctx).WithError(err).Error("GetConsumeCategoryTypes 获取消费类型映射失败")
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "获取门店积分规则错误"
			return resp, nil
		}
		type ConsumeTypeMapping struct {
			ActiveTypes   *fields.UUIDArr
			OperateTypes  *fields.UUIDArr
			ConsumeType   string
			ConsumeTypeID uuid.UUID
			IsRoomFee     bool
		}
		consumeTypeMap := make(map[uuid.UUID]ConsumeTypeMapping)
		for _, cate := range categories {
			consumeTypeMap[cate.ID] = ConsumeTypeMapping{
				ActiveTypes:   cate.ActiveTypes,
				OperateTypes:  cate.OperatorTypes,
				ConsumeType:   cate.Category,
				ConsumeTypeID: cate.ID,
				IsRoomFee:     cate.IsRoomFee,
			}
		}
		var protoGainRules, protoUseRules []*proto.PointRuleWithConsumeType
		for _, rule := range rules {
			for _, r := range *rule.GainRules {
				if tmpMapping, ok := consumeTypeMap[r.CategoryID]; ok {
					protoGainRules = append(protoGainRules, &proto.PointRuleWithConsumeType{
						Point:         r.Point,
						Fee:           r.Fee,
						OperateType:   tmpMapping.OperateTypes.ToStringArr(),
						ActiveType:    tmpMapping.ActiveTypes.ToStringArr(),
						ConsumeType:   tmpMapping.ConsumeType,
						ConsumeTypeId: tmpMapping.ConsumeTypeID.String(),
						IsRoomFee:     tmpMapping.IsRoomFee,
					})
				}
			}

			for _, r := range *rule.UseRules {
				if tmpMapping, ok := consumeTypeMap[r.CategoryID]; ok {
					protoUseRules = append(protoUseRules, &proto.PointRuleWithConsumeType{
						Point:         r.Point,
						Fee:           r.Fee,
						OperateType:   tmpMapping.OperateTypes.ToStringArr(),
						ActiveType:    tmpMapping.ActiveTypes.ToStringArr(),
						ConsumeType:   tmpMapping.ConsumeType,
						ConsumeTypeId: tmpMapping.ConsumeTypeID.String(),
						IsRoomFee:     tmpMapping.IsRoomFee,
					})
				}
			}
		}
		protoData.GainRule = protoGainRules
		protoData.UseRule = protoUseRules
	}

	resp.Data = &protoData

	return resp, nil
}

func (s *Server) checkPointRule(ctx context.Context, ruleName, ignoreID string) (bool, error) {
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", pkgs.GetMetadata(ctx).MerchantID))
	if len(ignoreID) > 0 {
		scopes = append(scopes, util.ColumnSymbolScope("id", "!=", ignoreID))
	}

	var total int64
	// 判断规则名称
	if err := model.DatabaseConn().Model(&model.PointRule{}).Scopes(scopes...).Where("rule_name", ruleName).Count(&total).Error; nil != err {
		return false, err
	}

	if total > 0 {
		return true, errors.New("规则名称已存在")
	}

	return false, nil
}
