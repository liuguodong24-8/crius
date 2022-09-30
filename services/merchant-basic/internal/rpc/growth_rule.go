package rpc

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

const growthGainConsumeType = "consume_type"
const growthGainCost = "cost"

// CreateGrowthRule 新建规则
func (s *Server) CreateGrowthRule(ctx context.Context, request *proto.CreateGrowthRuleRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("CreateGrowthRule params", logger.MakeFields(request)).Info("CreateGrowthRule")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	rule := model.TableGrowthRule{
		ID:         uuid.NewV4(),
		Name:       request.Rule.Name,
		ExpireDay:  request.Rule.ExpireDay,
		Status:     request.Rule.Status,
		MerchantID: merchantID,
	}

	if len(request.Rule.Branches) > 0 {
		var branches []uuid.UUID
		for _, branch := range request.Rule.Branches {
			if id := uuid.FromStringOrNil(branch); id != uuid.Nil {
				branches = append(branches, id)
			}
		}
		//校验门店是否已有规则
		if err := validateGrowthRuleBranches(branches, uuid.Nil, merchantID); err != nil {
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = err.Error()
			return resp, nil
		}
		rule.Branches = (*fields.UUIDArr)(&branches)
	}

	if len(request.Rule.GrowthGain) > 0 {
		var gains model.GrowthRules
		for _, gg := range request.Rule.GrowthGain {
			consumeType := uuid.FromStringOrNil(gg.ConsumeType)
			if consumeType != uuid.Nil {
				gain := model.GrowthRule{
					ConsumeTypeID: consumeType,
					Cost:          int64(gg.Cost),
				}
				gains = append(gains, gain)
			}
		}
		rule.GrowthGain = &gains
	}

	if err := model.CreateGrowthRule(rule); err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateGrowthRule 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
	}
	return resp, nil
}

// GetGrowthRules 列表
func (s *Server) GetGrowthRules(ctx context.Context, request *proto.GetGrowthRulesRequest) (*proto.GetGrowthRulesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("GetGrowthRules params", logger.MakeFields(request)).Info("GetGrowthRules")
	resp := &proto.GetGrowthRulesResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	count, err := model.CountGrowthRules(request.Name, request.Status, merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CountGrowthRules 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}
	var rulesProto []*proto.GrowthRule
	if count > 0 {
		rules, err := model.GetGrowthRules(request.Name, request.Status, merchantID, request.Offset, request.Limit)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("GetGrowthRules 数据库错误, %v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = err.Error()
			return resp, nil
		}

		if len(rules) > 0 {
			for _, rule := range rules {
				ruleProto := toGrowthRuleProto(rule)
				rulesProto = append(rulesProto, ruleProto)
			}
		}
	}

	resp.Data = &proto.GrowthRulesData{
		Rules: rulesProto,
		Total: int32(count),
	}
	return resp, nil
}

// ShowGrowthRule 详情
func (s *Server) ShowGrowthRule(ctx context.Context, request *proto.ShowGrowthRuleRequest) (*proto.ShowGrowthRuleResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("ShowGrowthRule params", logger.MakeFields(request)).Info("ShowGrowthRule")
	resp := &proto.ShowGrowthRuleResponse{
		ErrorCode: pkgs.Success,
	}
	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}
	rule, err := model.ShowGrowthRule(id)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("ShowGrowthRule 数据库错误, %v", err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	resp.Data = toGrowthRuleProto(rule)
	return resp, nil
}

// UpdateGrowthRule 更新
func (s *Server) UpdateGrowthRule(ctx context.Context, request *proto.UpdateGrowthRuleRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("UpdateGrowthRule params", logger.MakeFields(request)).Info("UpdateGrowthRule")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Rule.Id)
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	rule := model.TableGrowthRule{
		ID:        id,
		Name:      request.Rule.Name,
		ExpireDay: request.Rule.ExpireDay,
		Status:    request.Rule.Status,
	}

	if len(request.Rule.Branches) > 0 {
		var branches []uuid.UUID
		for _, branch := range request.Rule.Branches {
			if id := uuid.FromStringOrNil(branch); id != uuid.Nil {
				branches = append(branches, id)
			}
		}
		//校验门店是否已有规则
		if err := validateGrowthRuleBranches(branches, id, merchantID); err != nil {
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = err.Error()
			return resp, nil
		}
		rule.Branches = (*fields.UUIDArr)(&branches)
	}

	if len(request.Rule.GrowthGain) > 0 {
		var gains model.GrowthRules
		for _, gg := range request.Rule.GrowthGain {
			consumeType := uuid.FromStringOrNil(gg.ConsumeType)
			if consumeType != uuid.Nil {
				gain := model.GrowthRule{
					ConsumeTypeID: consumeType,
					Cost:          int64(gg.Cost),
				}
				gains = append(gains, gain)
			}
		}
		rule.GrowthGain = &gains
	}

	if err := model.UpdateGrowthRule(rule); err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateGrowthRule 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	return resp, nil
}

// GetBranchesHasGrowthRule 获取已有规则的门店列表
func (s *Server) GetBranchesHasGrowthRule(ctx context.Context, request *proto.Empty) (*proto.GetBranchesHasGrowthRuleResponse, error) {
	defer crius.CatchException()
	resp := &proto.GetBranchesHasGrowthRuleResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	rules, err := model.GetHasGrowthRuleBranches(uuid.Nil, merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetHasGrowthRuleBranches 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	if len(rules) > 0 {
		var branches []string
		for _, rule := range rules {
			branches = append(branches, rule.Branches.ToStringArr()...)
		}
		resp.Data = branches
	}

	return resp, nil
}

// GetBranchGrowthRule 门店成长值规则
func (s *Server) GetBranchGrowthRule(ctx context.Context, request *proto.GetBranchGrowthRuleRequest) (*proto.GetBranchGrowthRuleResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("ShowGrowthRule params", logger.MakeFields(request)).Info("ShowGrowthRule")
	resp := &proto.GetBranchGrowthRuleResponse{
		ErrorCode: pkgs.Success,
	}
	branchID := uuid.FromStringOrNil(request.BranchId)

	if branchID == uuid.Nil {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}
	rule, err := model.GetBranchGrowthRule(branchID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("GetBranchGrowthRule 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	resp.Data = toGrowthRuleProto(*rule)
	return resp, nil
}

func toGrowthRuleProto(rule model.TableGrowthRule) *proto.GrowthRule {
	var gains []*proto.GrowthGain
	if rule.GrowthGain != nil {
		for _, gg := range *rule.GrowthGain {
			gain := &proto.GrowthGain{
				ConsumeType: gg.ConsumeTypeID.String(),
				Cost:        uint32(gg.Cost),
			}
			gains = append(gains, gain)
		}
	}
	return &proto.GrowthRule{
		Id:         rule.ID.String(),
		Name:       rule.Name,
		GrowthGain: gains,
		ExpireDay:  rule.ExpireDay,
		Branches:   rule.Branches.ToStringArr(),
		Status:     rule.Status,
	}
}

//门店是否在已有规则里
func validateGrowthRuleBranches(branches []uuid.UUID, id, merchantID uuid.UUID) error {
	rules, err := model.GetHasGrowthRuleBranches(id, merchantID)
	if err != nil {
		return err
	}

	if len(rules) == 0 {
		return nil
	}

	usedBranches := make(map[uuid.UUID]int)
	for _, rule := range rules {
		for _, branch := range rule.Branches.Slice() {
			usedBranches[branch] = 1
		}
	}

	for _, branch := range branches {
		if _, ok := usedBranches[branch]; ok {
			return errors.New("门店已被使用")
		}
	}
	return nil
}
