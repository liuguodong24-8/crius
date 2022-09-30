package rpc

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// CreateBranchBusinessSpecial 创建特殊营业时间
func (s *Server) CreateBranchBusinessSpecial(ctx context.Context, req *proto.CreateBranchBusinessSpecialRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreateBranchBusinessSpecial")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	branchID := uuid.FromStringOrNil(req.Business.BranchId)
	beginDate, err1 := time.Parse(dateLayout, req.Business.BeginDate)
	endDate, err2 := time.Parse(dateLayout, req.Business.EndDate)
	if err1 != nil || err2 != nil {
		crius.Logger.Error(fmt.Sprintf("CreateBranchBusinessSpecial begindate:%v, enddate:%v,参数错误", req.Business.BeginDate, req.Business.EndDate))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	beginTime := fields.StringToLocalTime(req.Business.BeginTime)
	endTime := fields.StringToLocalTime(req.Business.EndTime)

	business := model.TableBranchBusiness{
		ID:         uuid.NewV4(),
		BranchID:   &branchID,
		BeginDate:  &beginDate,
		EndDate:    &endDate,
		BeginTime:  &beginTime,
		EndTime:    &endTime,
		IsNextDay:  req.Business.IsNextDay,
		MerchantID: &merchantID,
		Status:     req.Business.Status,
		Category:   model.BranchBusinessCategorySpecial,
	}
	err := model.CreateBranchBusiness(business)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateBranchBusinessSpecial 创建特殊营业时间数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建特殊营业时间失败"
		return resp, nil
	}

	return resp, nil
}

// UpdateBranchBusinessSpecial 更新特殊营业时间
func (s *Server) UpdateBranchBusinessSpecial(ctx context.Context, req *proto.UpdateBranchBusinessSpecialRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateBranchBusinessSpecial")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Business.Id)

	beginDate, err1 := time.Parse(dateLayout, req.Business.BeginDate)
	endDate, err2 := time.Parse(dateLayout, req.Business.EndDate)
	if err1 != nil || err2 != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBranchBusinessSpecial begindate:%v, enddate:%v,参数错误", req.Business.BeginDate, req.Business.EndDate))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	beginTime := fields.StringToLocalTime(req.Business.BeginTime)
	endTime := fields.StringToLocalTime(req.Business.EndTime)

	business := model.TableBranchBusiness{
		ID:        id,
		BeginDate: &beginDate,
		EndDate:   &endDate,
		BeginTime: &beginTime,
		EndTime:   &endTime,
		IsNextDay: req.Business.IsNextDay,
		Status:    req.Business.Status,
	}
	err := model.UpdateBranchBusinessSpecial(business)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBranchBusinessSpecial 更新特殊营业时间数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新特殊营业时间失败"
		return resp, nil
	}

	return resp, nil
}

// UpdateBranchBusinessNormal 更新普通营业时间
func (s *Server) UpdateBranchBusinessNormal(ctx context.Context, req *proto.UpdateBranchBusinessNormalRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateBranchBusinessNormal")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	branchID := uuid.FromStringOrNil(req.BranchId)

	businesses := make([]model.TableBranchBusiness, 0)
	for _, v := range req.Businesses {
		weeks := fields.IntArrToInt8Arr(v.Weeks)
		beginTime := fields.StringToLocalTime(v.BeginTime)
		endTime := fields.StringToLocalTime(v.EndTime)
		business := model.TableBranchBusiness{
			ID:         uuid.NewV4(),
			BranchID:   &branchID,
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
	err := model.UpdateBranchBusinessNormal(businesses, branchID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBranchBusinessNormal 更新普通营业时间数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新普通营业时间失败"
		return resp, nil
	}

	return resp, nil
}

// UpdateBranchBusinessStatus 更新营业时间状态
func (s *Server) UpdateBranchBusinessStatus(ctx context.Context, req *proto.UpdateBranchBusinessStatusRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateBranchBusinessStatus")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)

	err := model.UpdateBranchBusinessStatus(id, req.Status)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBranchBusinessStatus 更新营业时间状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新特殊营业时间状态失败"
		return resp, nil
	}

	return resp, nil
}

// GetBranchBusinesses 获取
func (s *Server) GetBranchBusinesses(ctx context.Context, req *proto.GetBranchBusinessesRequest) (*proto.GetBranchBusinessesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetBranchBusinesses")
	resp := &proto.GetBranchBusinessesResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	branchID := uuid.FromStringOrNil(req.BranchId)

	business := model.TableBranchBusiness{
		BranchID:   &branchID,
		MerchantID: &merchantID,
		Status:     req.Status,
		Category:   model.BranchBusinessCategory(req.Category),
	}

	businesses, count, err := model.GetBranchBusiness(business, req.Offset, req.Limit)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranchBusinesses 获取营业时间列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取营业时间列表失败"
		return resp, nil
	}

	resp.Data = &proto.GetBranchBusinessesData{
		Total: int32(count),
	}

	for _, v := range businesses {
		resp.Data.Businesses = append(resp.Data.Businesses, toProtoBusiness(v))
	}
	return resp, nil
}

func toProtoBusiness(v model.TableBranchBusiness) *proto.BranchBusiness {
	var branchID, beginDate, endDate, beginTime, endTime string
	if v.BranchID != nil {
		branchID = v.BranchID.String()
	}
	if v.BeginDate != nil {
		beginDate = v.BeginDate.Format(dateLayout)
	}
	if v.EndDate != nil {
		endDate = v.EndDate.Format(dateLayout)
	}
	if v.BeginTime != nil {
		beginTime = v.BeginTime.String()
	}
	if v.EndTime != nil {
		endTime = v.EndTime.String()
	}
	return &proto.BranchBusiness{
		Id:        v.ID.String(),
		BranchId:  branchID,
		BeginDate: beginDate,
		EndDate:   endDate,
		Weeks:     v.Weeks.SliceInt32(),
		BeginTime: beginTime,
		EndTime:   endTime,
		IsNextDay: v.IsNextDay,
		Status:    v.Status,
		Category:  string(v.Category),
	}
}

// GetBranchLatelyBusiness 获取门店最近一次营业日
func (s *Server) GetBranchLatelyBusiness(ctx context.Context, req *proto.GetBranchLatelyBusinessRequest) (*proto.GetBranchLatelyBusinessResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("GetBranchLatelyBusiness")

	if len(req.BranchId) == 0 {
		return &proto.GetBranchLatelyBusinessResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "门店参数错误",
		}, nil
	}

	branchID := uuid.FromStringOrNil(req.BranchId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	date := time.Now()
	if req.DateTime != 0 {
		date = time.Unix(req.DateTime, 0)
	}
	businesses, err := model.ShowBranchBusinessByBranchIDDate(branchID, merchantID, date)

	if err != nil {
		return &proto.GetBranchLatelyBusinessResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("获取门店营业日错误:%s", err.Error()),
		}, nil
	}

	return &proto.GetBranchLatelyBusinessResponse{
		ErrorCode:    0,
		ErrorMessage: "",
		Data: &proto.GetBranchLatelyBusinessResponse_Business{
			BusinessDate: date.Format(`2006-01-02`),
			BeginTime:    businesses.BeginTime.String(),
			EndTime:      businesses.EndTime.String(),
			IsNextDay:    businesses.IsNextDay,
		},
	}, nil
}
