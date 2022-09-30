package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/util"
	"gorm.io/gorm"
)

// CreateBranch CreateBranch
func (s *Server) CreateBranch(ctx context.Context, request *proto.CreateBranchRequest) (*proto.CreateBranchResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("CreateBranch params", logger.MakeFields(request)).Info("CreateBranch")
	resp := &proto.CreateBranchResponse{
		ErrorCode: pkgs.Success,
	}

	// 判断参数合法性
	if request.Branch == nil {
		crius.Logger.Error("CreateBranch rpc请求参数nil")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	areaID := uuid.FromStringOrNil(request.Branch.AreaId)
	brandID := uuid.FromStringOrNil(request.Branch.BrandId)

	branch := model.TableBranch{
		ID:             uuid.NewV4(),
		Name:           request.Branch.Name,
		ProvinceID:     request.Branch.ProvinceId,
		CityID:         request.Branch.CityId,
		DistrictID:     request.Branch.DistrictId,
		Address:        request.Branch.Address,
		ContactPhone:   request.Branch.Phone,
		Latitude:       request.Branch.Latitude,
		Longitude:      request.Branch.Longitude,
		Status:         model.StatusOpened,
		MerchantID:     &merchantID,
		Parking:        request.Branch.Parking,
		Weight:         request.Branch.Weight,
		Domain:         request.Branch.Domain,
		BizType:        int8(request.Branch.BizType),
		BusinessStatus: request.Branch.BusinessStatus,
		Alias:          request.Branch.Alias,
		Simplify:       strings.ToUpper(request.Branch.Simplify),
		Location:       request.Branch.Location,
		BrandID:        &brandID,
	}
	if areaID != uuid.Nil {
		branch.AreaID = &areaID
	}

	if request.Branch.OpenedAt != 0 {
		openedAt := time.Unix(int64(request.Branch.OpenedAt), 0)
		branch.OpenedAt = &openedAt
	}

	if len(request.Branch.Photo) != 0 {
		branch.Photo = (*fields.StringArr)(&request.Branch.Photo)
	}

	if _, err := model.ShowBranchByName(branch.Name, merchantID); err != nil {
		if err != gorm.ErrRecordNotFound { //查询数据报错且不是没有找到数据错误
			crius.Logger.Error(fmt.Sprintf("CreateBranch 校验用户输入门店名称数据库错误:%v", err))
			resp.ErrorMessage = "检验门店名称错误"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	} else { //正确查询到数据
		crius.Logger.Error("CreateBranch 用户输入门店名称已经存在")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "门店名称已存在"
		return resp, nil
	}

	for i := 0; i < 10; i++ {
		// 获取数据库门店code序列号
		num := model.GetBranchSequence()
		if num < 0 {
			resp.ErrorMessage = "门店编码生成失败"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
		// 生成门店code
		code := util.GenerateBranchCode(request.Branch.ProvinceId, num)
		// 查询门店code是否已存在
		if err := model.ShowBranchExistsByCode(code, merchantID); err != nil {
			if err != gorm.ErrRecordNotFound { //查询数据报错且不是没有找到数据错误
				crius.Logger.Error(fmt.Sprintf("CreateBranch 校验门店生成code数据库错误:%v", err))
				resp.ErrorMessage = "检验生成门店编码失败"
				resp.ErrorCode = pkgs.ErrInternal
				return resp, nil
			}
			branch.Code = code
			break
		}

		// 此处门店code已存在，重新生成code

	}
	if branch.Code == "" {
		crius.Logger.Error("CreateBranch 门店code生成10次已存在，创建门店失败")
		resp.ErrorMessage = "生成门店编号已存在"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	json.Unmarshal([]byte(request.Branch.Extra), &branch.Extra)

	businesses := make([]model.TableBranchBusiness, 0)
	for _, v := range request.BusinessHours {
		weeks := fields.IntArrToInt8Arr(v.Weeks)
		beginTime := fields.StringToLocalTime(v.BeginTime)
		endTime := fields.StringToLocalTime(v.EndTime)
		business := model.TableBranchBusiness{
			ID:         uuid.NewV4(),
			BranchID:   &branch.ID,
			Weeks:      &weeks,
			BeginTime:  &beginTime,
			EndTime:    &endTime,
			IsNextDay:  v.IsNextDay,
			MerchantID: &merchantID,
			Status:     crius.StatusOpened.String(),
			Category:   model.BranchBusinessCategoryNormal,
		}
		businesses = append(businesses, business)
	}

	err := model.CreateBranch(branch, businesses)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateBranch 新增门店数据库操作错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "新增门店错误"
		return resp, nil
	}

	staffID := pkgs.GetMetadata(ctx).StaffID
	after := pkgs.MakeParams(branch)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: branch.TableName(),
		After:             &after,
		TableID:           &branch.ID,
		Method:            "create",
	}

	resp.Data = branch.ID.String()
	return resp, nil
}

// UpdateBranch UpdateBranch
func (s *Server) UpdateBranch(ctx context.Context, request *proto.UpdateBranchRequest) (*proto.UpdateBranchResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("UpdateBranch params", logger.MakeFields(request)).Info("UpdateBranch")
	resp := &proto.UpdateBranchResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(request.Branch.Id)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	areaID := uuid.FromStringOrNil(request.Branch.AreaId)
	brandID := uuid.FromStringOrNil(request.Branch.BrandId)

	branch := model.TableBranch{
		ID:             branchID,
		Name:           request.Branch.Name,
		ProvinceID:     request.Branch.ProvinceId,
		CityID:         request.Branch.CityId,
		DistrictID:     request.Branch.DistrictId,
		Address:        request.Branch.Address,
		ContactPhone:   request.Branch.Phone,
		Latitude:       request.Branch.Latitude,
		Longitude:      request.Branch.Longitude,
		Parking:        request.Branch.Parking,
		Weight:         request.Branch.Weight,
		Domain:         request.Branch.Domain,
		BizType:        int8(request.Branch.BizType),
		BusinessStatus: request.Branch.BusinessStatus,
		Alias:          request.Branch.Alias,
		Simplify:       strings.ToUpper(request.Branch.Simplify),
		Location:       request.Branch.Location,
		BrandID:        &brandID,
		MerchantID:     &merchantID,
	}

	if areaID != uuid.Nil {
		branch.AreaID = &areaID
	}

	if request.Branch.OpenedAt != 0 {
		openedAt := time.Unix(int64(request.Branch.OpenedAt), 0)
		branch.OpenedAt = &openedAt
	}

	if len(request.Branch.Photo) != 0 {
		branch.Photo = (*fields.StringArr)(&request.Branch.Photo)
	}

	var tabBranch *model.TableBranch
	var err error
	if tabBranch, err = model.ShowBranchByName(branch.Name, merchantID); err != nil {
		if err != gorm.ErrRecordNotFound { //查询数据报错且不是没有找到数据错误
			crius.Logger.Error(fmt.Sprintf("CreateBranch 校验用户输入门店名称数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "校验门店名称错误"
			return resp, nil
		}
	} else { //正确查询到数据
		if tabBranch.ID != branchID {
			crius.Logger.Error("CreateBranch 用户输入门店名称已经存在")
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "门店名称已存在"
			return resp, nil
		}
	}

	businesses := make([]model.TableBranchBusiness, 0)
	for _, v := range request.BusinessHours {
		weeks := fields.IntArrToInt8Arr(v.Weeks)
		beginTime := fields.StringToLocalTime(v.BeginTime)
		endTime := fields.StringToLocalTime(v.EndTime)
		business := model.TableBranchBusiness{
			ID:         uuid.NewV4(),
			BranchID:   &branch.ID,
			Weeks:      &weeks,
			BeginTime:  &beginTime,
			EndTime:    &endTime,
			IsNextDay:  v.IsNextDay,
			MerchantID: &merchantID,
			Status:     crius.StatusOpened.String(),
			Category:   model.BranchBusinessCategoryNormal,
		}
		businesses = append(businesses, business)
	}

	err = model.UpdateBranch(branch, businesses)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBranch 更新门店信息数据库操作错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新门店错误"
		return resp, nil
	}

	branch.Extra = tabBranch.Extra
	branch.Code = tabBranch.Code
	staffID := pkgs.GetMetadata(ctx).StaffID
	before := pkgs.MakeParams(*tabBranch)
	after := pkgs.MakeParams(branch)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: branch.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &branch.ID,
		Method:            "update",
	}

	return resp, nil
}

// UpdateBranchStatus 更新门店状态
func (s *Server) UpdateBranchStatus(ctx context.Context, request *proto.UpdateBranchStatusRequest) (*proto.UpdateBranchStatusResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateBranchStatus")
	resp := &proto.UpdateBranchStatusResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil || (request.Status != model.StatusOpened && request.Status != model.StatusClosed) {
		crius.Logger.Error(fmt.Sprintf("UpdateBranchStatus 请求参数错误"))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	tabBranch, err := model.ShowBranchByID(id)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "查询门店失败"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdateBranchStatus 查询门店错误:%v", err))
		resp.ErrorCode = pkgs.ErrNotFound
		resp.ErrorMessage = "门店不存在"
		return resp, nil
	}

	err = model.UpdateBranchStatus(id, request.Status)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBranchStatus 更新门店状态:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新门店状态失败"
		return resp, nil
	}

	branch := *tabBranch
	branch.Status = crius.Status(request.Status)
	staffID := pkgs.GetMetadata(ctx).StaffID
	before := pkgs.MakeParams(*tabBranch)
	after := pkgs.MakeParams(branch)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: branch.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &branch.ID,
		Method:            "update",
	}

	return resp, nil
}

// GetBranches GetBranches
func (s *Server) GetBranches(ctx context.Context, request *proto.GetBranchesRequest) (*proto.GetBranchesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBranches")
	resp := &proto.GetBranchesResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	areaID := uuid.FromStringOrNil(request.AreaId)
	brandID := uuid.FromStringOrNil(request.BrandId)
	if merchantID == uuid.Nil || (request.Status != model.StatusClosed && request.Status != model.StatusOpened && request.Status != "") {
		crius.Logger.Error(fmt.Sprintf("GetBranches 请求参数错误:%v", merchantID))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	// 查询当前员工的门店
	staffID := uuid.FromStringOrNil(request.StaffId)
	if staffID != uuid.Nil {
		staff, err := model.ShowBasicStaffByID(staffID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				crius.Logger.Error(fmt.Sprintf("GetBranches 查询员工不存在,id:%v", staffID))
				resp.ErrorCode = pkgs.ErrNotFound
				resp.ErrorMessage = "当前登录员工不存在"
				return resp, nil
			}
			crius.Logger.Error(fmt.Sprintf("GetBranches 查询员工数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "查询员工信息错误"
			return resp, nil
		}
		if staff.Admin {
			staffID = uuid.Nil
		}
	}

	branch := model.TableBranch{
		Name:       request.Name,
		ProvinceID: request.ProvinceId,
		CityID:     request.CityId,
		DistrictID: request.DistrictId,
		Status:     crius.Status(request.Status),
	}
	if areaID != uuid.Nil {
		branch.AreaID = &areaID
	}
	branches, count, err := model.GetBranches(branch, request.BusinessStatus, brandID, staffID, merchantID, request.Offset, request.Limit)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranches 获取门店信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取门店信息错误"
		return resp, nil
	}

	resp.Data = &proto.BranchesData{Total: int32(count)}

	for i := range branches {
		protoBranch := toProtoBranch(&branches[i])
		branchBusiness, _, err := model.GetBranchBusiness(model.TableBranchBusiness{
			BranchID:   &branches[i].ID,
			MerchantID: branches[i].MerchantID,
			Category:   model.BranchBusinessCategoryNormal,
		}, 0, 0)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("GetBranches 获取门店营业时间数据库错误:%v", err))
		} else {
			for _, v := range branchBusiness {
				protoBranch.BusinessHours = append(protoBranch.BusinessHours, toProtoBusiness(v))
			}
		}
		resp.Data.Branches = append(resp.Data.Branches, protoBranch)
	}
	return resp, nil
}

// GetBranchesByTagIDs GetBranchesByTagIDs
func (s *Server) GetBranchesByTagIDs(ctx context.Context, request *proto.GetBranchesByTagIDsRequest) (*proto.GetBranchesByTagIDsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBranchesByTagIDs")
	resp := &proto.GetBranchesByTagIDsResponse{
		ErrorCode: pkgs.Success,
	}

	tagIDs := make([]uuid.UUID, 0)
	for _, id := range request.TagIds {
		if tagID := uuid.FromStringOrNil(id); tagID != uuid.Nil {
			tagIDs = append(tagIDs, tagID)
		}
	}

	tags, err := model.GetBranchTagsByIDs(tagIDs)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranchesByTagIDs 根据标签id获取标签数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取门店信息失败"
		return resp, nil
	}

	for _, v := range tags {
		data := &proto.GetBranchesByTagIDsData{
			BranchTag: &proto.BranchTagData{
				Id:        v.ID.String(),
				Name:      v.Name,
				Branches:  v.Branches.ToStringArr(),
				StaffName: v.StaffName,
				Status:    v.Status.String(),
				CreatedAt: int32(v.CreatedAt.Unix()),
				UpdatedAt: int32(v.UpdatedAt.Unix()),
			},
		}
		if v.CreateStaffID != nil {
			data.BranchTag.CreateStaffId = v.CreateStaffID.String()
		}
		branches, err := model.GetBranchesByIDs(v.Branches.Slice(), request.Status, request.BusinessStatus)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("GetBranchesByTagIDs 根据标签id获取门店数据库错误:%v", err))
			resp.ErrorMessage = "获取门店信息失败"
			resp.ErrorCode = pkgs.ErrInternal
			resp.Data = nil
			return resp, nil
		}
		for i := range branches {
			protoBranch := toProtoBranch(&branches[i])
			branchBusiness, _, err := model.GetBranchBusiness(model.TableBranchBusiness{
				BranchID:   &branches[i].ID,
				MerchantID: branches[i].MerchantID,
				Category:   model.BranchBusinessCategoryNormal,
			}, 0, 0)
			if err != nil {
				crius.Logger.Error(fmt.Sprintf("GetBranchesByTagIDs 获取门店营业时间数据库错误:%v", err))
			} else {
				for _, v := range branchBusiness {
					protoBranch.BusinessHours = append(protoBranch.BusinessHours, toProtoBusiness(v))
				}
			}
			data.Branches = append(data.Branches, protoBranch)
		}
		resp.Data = append(resp.Data, data)
	}

	return resp, nil
}

// ShowBranch ShowBranch
func (s *Server) ShowBranch(ctx context.Context, request *proto.ShowBranchRequest) (*proto.ShowBranchResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowBranch")
	resp := &proto.ShowBranchResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(request.Id)
	if branchID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("ShowBranch rpc请求参数错误，门店id:%v", request.Id))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	branch, err := model.ShowBranchByID(branchID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "门店不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowBranch 查询门店数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "查询门店失败"
		return resp, nil
	}
	protoBranch := toProtoBranch(branch)
	branchBusiness, _, err := model.GetBranchBusiness(model.TableBranchBusiness{
		BranchID:   &branch.ID,
		MerchantID: branch.MerchantID,
		Category:   model.BranchBusinessCategoryNormal,
	}, 0, 0)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("ShowBranch 获取门店营业时间数据库错误:%v", err))
	} else {
		for _, v := range branchBusiness {
			protoBranch.BusinessHours = append(protoBranch.BusinessHours, toProtoBusiness(v))
		}
	}
	resp.Data = protoBranch
	return resp, nil
}

// DeleteBranch 删除门店
func (s *Server) DeleteBranch(ctx context.Context, request *proto.DeleteBranchRequest) (*proto.DeleteBranchResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("DeleteBranch")
	resp := &proto.DeleteBranchResponse{
		ErrorCode: pkgs.Success,
	}
	branchID := uuid.FromStringOrNil(request.Id)
	if branchID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("DeleteBranch rpc请求参数错误，门店id:%v", request.Id))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	branch, err := model.ShowBranchByID(branchID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "门店不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("DeleteBranch 获取门店信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除门店错误"
		return resp, nil
	}

	err = model.DeleteBranch(branchID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("DeleteBranch 删除门店信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除门店错误"
		return resp, nil
	}

	staffID := pkgs.GetMetadata(ctx).StaffID
	before := pkgs.MakeParams(branch)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: model.TableBranch{}.TableName(),
		Before:            &before,
		TableID:           &branchID,
		Method:            "delete",
	}

	return resp, nil
}

// GetBranchConsumeRewardRule 消费奖励规则
func (s *Server) GetBranchConsumeRewardRule(ctx context.Context, request *proto.GetBranchConsumeRewardRuleRequest) (*proto.GetBranchConsumeRewardRuleResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBranchConsumeRewardRule")
	resp := &proto.GetBranchConsumeRewardRuleResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	//消费类型 code转id
	categories, err := model.GetConsumeCategoriesByCodes(request.ConsumeTypeCodes, merchantID)
	if err != nil || len(categories) == 0 {
		crius.Logger.Error(fmt.Sprintf("GetConsumeCategoryIDsByCodes 消费类型查询错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "消费类型查询错误"
		return resp, nil
	}
	branchID := uuid.FromStringOrNil(request.BranchId)
	if branchID == uuid.Nil {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "branchID格式错误"
		return resp, nil
	}

	GrowthRule, growthErr := model.GetBranchGrowthRule(branchID)
	if growthErr != nil && !errors.Is(growthErr, gorm.ErrRecordNotFound) {
		crius.Logger.Error(fmt.Sprintf("GetBranchGrowthRule 成长值规则查询错误:%v", growthErr))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "成长值规则查询错误"
		return resp, nil
	}

	PointRule, pointErr := model.GetBranchPointGainRule(branchID)
	if pointErr != nil && !errors.Is(pointErr, gorm.ErrRecordNotFound) {
		crius.Logger.Error(fmt.Sprintf("GetBranchPointGainRule 积分规则查询错误:%v", pointErr))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "积分规则查询错误"
		return resp, nil
	}
	categoriesMap := make(map[uuid.UUID]string)
	for _, cat := range categories {
		categoriesMap[cat.ID] = cat.Code
	}

	var growthGainRules []*proto.ConsumeRewardRule_GrowthRule
	if growthErr == nil {
		for _, rule := range *GrowthRule.GrowthGain {
			if code, ok := categoriesMap[rule.ConsumeTypeID]; ok {
				growthGainRules = append(growthGainRules, &proto.ConsumeRewardRule_GrowthRule{
					ConsumeTypeCode: code,
					Cost:            int32(rule.Cost),
				})
			}
		}
	}

	var pointGainRules []*proto.ConsumeRewardRule_PointRule

	if pointErr == nil {
		for _, rule := range *PointRule.GainRules {
			if code, ok := categoriesMap[rule.CategoryID]; ok {
				pointGainRules = append(pointGainRules, &proto.ConsumeRewardRule_PointRule{
					ConsumeTypeCode: code,
					Fee:             int32(rule.Fee),
					Point:           int32(rule.Point),
				})
			}
		}
	}

	resp.Data = &proto.ConsumeRewardRule{
		GrowthRules: &proto.ConsumeRewardRule_GrowthRules{
			Rules:       growthGainRules,
			ValidityDay: GrowthRule.ExpireDay,
		},
		PointRules: &proto.ConsumeRewardRule_PointRules{
			Rules:       pointGainRules,
			ValidityDay: PointRule.ValidityDay,
		},
	}
	return resp, nil
}

// // UpdateBranchAccount 更新门店账户
// func (s *Server) UpdateBranchAccount(ctx context.Context, request *proto.UpdateBranchAccountRequest) (*proto.UpdateBranchAccountResponse, error) {
// 	defer crius.CatchException()
// 	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateBranchAccount")
// 	resp := &proto.UpdateBranchAccountResponse{
// 		ErrorCode: pkgs.Success,
// 	}
// 	id := uuid.FromStringOrNil(request.Id)
// 	if id == uuid.Nil {
// 		crius.Logger.Error(fmt.Sprintf("UpdateBranchAccount rpc请求参数错误，门店id:%v", request.Id))
// 		resp.ErrorCode = pkgs.ErrUnprocessableEntity
// 		resp.ErrorMessage = "请求参数错误"
// 		return resp, nil
// 	}

// 	tabBranch, err := model.ShowBranchByID(id)
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			resp.ErrorCode = pkgs.ErrInternal
// 			resp.ErrorMessage = "查询门店失败"
// 			return resp, nil
// 		}
// 		crius.Logger.Error(fmt.Sprintf("UpdateBranchAccount 查询门店错误:%v", err))
// 		resp.ErrorCode = pkgs.ErrNotFound
// 		resp.ErrorMessage = "门店不存在"
// 		return resp, nil
// 	}

// 	branch := *tabBranch

// 	err = model.UpdateBranch(branch)
// 	if err != nil {
// 		crius.Logger.Error(fmt.Sprintf("UpdateBranchAccount 更新门店账户信息数据库错误:%v", err))
// 		resp.ErrorCode = pkgs.ErrInternal
// 		resp.ErrorMessage = "更新门店错误"
// 		return resp, nil
// 	}

// 	staffID := pkgs.GetMetadata(ctx).StaffID
// 	before := pkgs.MakeParams(tabBranch)
// 	after := pkgs.MakeParams(branch)
// 	model.SnapShotChan <- &model.TableSnapshot{
// 		ID:                uuid.NewV4(),
// 		StaffID:           &staffID,
// 		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
// 		SnapShotTableName: model.TableBranch{}.TableName(),
// 		Before:            &before,
// 		After:             &after,
// 		TableID:           &id,
// 		Method:            "update",
// 	}

// 	return resp, nil
// }

func toProtoBranch(branch *model.Branch) *proto.BranchInfo {
	var createdAt, openedAt int32 // 处理时间为0时，unix时间为负数
	var extra, areaID, brandID string
	var photo []string
	if branch.CreatedAt != nil {
		createdAt = int32(branch.CreatedAt.Unix())
	}
	if branch.Extra != nil {
		extra = branch.Extra.JSON()
	}
	if branch.OpenedAt != nil {
		openedAt = int32(branch.OpenedAt.Unix())
	}
	if branch.Photo != nil {
		photo = branch.Photo.Slice()
	}
	if branch.AreaID != nil {
		areaID = branch.AreaID.String()
	}
	if branch.BrandID != nil {
		brandID = branch.BrandID.String()
	}
	return &proto.BranchInfo{
		Name:           branch.Name,
		ProvinceId:     branch.ProvinceID,
		CityId:         branch.CityID,
		DistrictId:     branch.DistrictID,
		Address:        branch.Address,
		Phone:          branch.ContactPhone,
		CreatedAt:      createdAt,
		Code:           branch.Code,
		Extra:          extra,
		Id:             branch.ID.String(),
		Longitude:      branch.Longitude,
		Latitude:       branch.Latitude,
		Status:         branch.Status.String(),
		OpenedAt:       openedAt,
		Photo:          photo,
		Parking:        branch.Parking,
		AreaId:         areaID,
		Weight:         int32(branch.Weight),
		Domain:         branch.Domain,
		BizType:        int32(branch.BizType),
		BusinessStatus: branch.BusinessStatus,
		Alias:          branch.Alias,
		Simplify:       branch.Simplify,
		Location:       branch.Location,
		BrandId:        brandID,
		BrandName:      branch.BrandName,
	}
}
