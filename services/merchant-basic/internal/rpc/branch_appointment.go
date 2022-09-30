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
	"gorm.io/gorm"
)

const (
	dateLayout = "2006-1-2"
	timeLayout = "15:4"
)

// SaveBranchAppointment 更新
func (s *Server) SaveBranchAppointment(ctx context.Context, request *proto.SaveBranchAppointmentRequest) (*proto.SaveBranchAppointmentResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("SaveBranchAppointment")
	resp := &proto.SaveBranchAppointmentResponse{
		ErrorCode: pkgs.Success,
	}
	openAppointment := int8(1)
	if !request.OpenAppointment {
		openAppointment = 0
	}
	branchID := uuid.FromStringOrNil(request.BranchId)
	paramArr := make(pkgs.ParamsArr, 0)
	for _, v := range request.Vr {
		m := make(map[string]interface{})
		m["url"] = v.Url
		m["name"] = v.Name
		m["description"] = v.Description
		paramArr = append(paramArr, m)
	}
	appointment := model.TableBranchAppointment{
		OpenAppointment:        openAppointment,
		AppointmentGranularity: int8(request.AppointmentGranularity),
		VR:                     &paramArr,
		Video:                  (*fields.StringArr)(&request.Video),
		Environment:            (*fields.StringArr)(&request.Environment),
		Meal:                   (*fields.StringArr)(&request.Meal),
		Price:                  (*fields.StringArr)(&request.Price),
		Hot:                    request.Hot,
		BranchID:               &branchID,
	}

	if err := model.SaveBranchAppointment(appointment); err != nil {
		crius.Logger.Error(fmt.Sprintf("SaveBranchAppointment 创建门店预约数据数据库错误:%v", err))
		resp.ErrorMessage = "创建门店预约数据失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	return resp, nil
}

// ShowBranchAppointment 查询单条
func (s *Server) ShowBranchAppointment(ctx context.Context, request *proto.ShowBranchAppointmentRequest) (*proto.ShowBranchAppointmentResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowBranchAppointment")
	resp := &proto.ShowBranchAppointmentResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(request.BranchId)

	appointment, err := model.ShowBranchAppointment(branchID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "门店预约数据不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowBranchAppointment 查询门店预约数据数据库错误:%v", err))
		resp.ErrorMessage = "查询门店预约数据失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	// businesses, _, err := model.GetBranchBusiness(model.TableBranchBusiness{
	// 	BranchID:   &branchID,
	// 	MerchantID: &merchantID,
	// }, 0, 0)
	// if err != nil {
	// 	crius.Logger.Error(fmt.Sprintf("ShowBranchAppointment 查询门店营业时间数据库错误:%v", err))
	// 	resp.ErrorMessage = "查询门店预约数据失败"
	// 	resp.ErrorCode = pkgs.ErrInternal
	// 	return resp, nil
	// }

	vr := make([]*proto.BranchAppointmentVR, 0)
	if appointment.VR != nil {
		for _, v := range *appointment.VR {
			vr = append(vr, &proto.BranchAppointmentVR{Url: v["url"].(string), Name: v["name"].(string), Description: v["description"].(string)})
		}
	}
	resp.Data = &proto.BranchAppointmentData{
		AppointmentGranularity: int32(appointment.AppointmentGranularity),
		Vr:                     vr,
		Video:                  appointment.Video.Slice(),
		Environment:            appointment.Environment.Slice(),
		Meal:                   appointment.Meal.Slice(),
		Price:                  appointment.Price.Slice(),
		Hot:                    appointment.Hot,
	}
	if appointment.OpenAppointment == 1 {
		resp.Data.OpenAppointment = true
	}
	if appointment.BranchID != nil {
		resp.Data.BranchId = appointment.BranchID.String()
	}
	if appointment.RoomTypes != nil {
		resp.Data.RoomTypes = appointment.RoomTypes.JSON()
	}

	// for _, v := range businesses {
	// 	resp.Data.BusinessHours = append(resp.Data.BusinessHours, toProtoBusiness(v))
	// }

	return resp, nil
}

// UpdateBranchAppointmentRoomType 更新门店房型关联数据
func (s *Server) UpdateBranchAppointmentRoomType(ctx context.Context, request *proto.UpdateBranchAppointmentRoomTypeRequest) (*proto.UpdateBranchAppointmentRoomTypeResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateBranchAppointmentRoomType")
	resp := &proto.UpdateBranchAppointmentRoomTypeResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(request.BranchId)

	roomTypes := make(pkgs.ParamsArr, 0)
	for _, v := range request.RoomType {
		m := make(map[string]interface{})
		m["room_type_id"] = v.RoomTypeId
		m["num"] = v.RoomTypeNum
		roomTypes = append(roomTypes, m)
	}
	appointment := model.TableBranchAppointment{
		BranchID:  &branchID,
		RoomTypes: &roomTypes,
	}

	err := model.UpdateBranchAppointmentRoomType(appointment)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBranchAppointmentRoomType 更新门店预约配置数据库错误:%v", err))
		resp.ErrorMessage = "更新门店预约数据失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	return resp, nil
}

// ShowBranchAppointmentBusinessTime 查询单条
func (s *Server) ShowBranchAppointmentBusinessTime(ctx context.Context, request *proto.ShowBranchAppointmentBusinessTimeRequest) (*proto.ShowBranchAppointmentBusinessTimeResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowBranchAppointmentBusinessTime")
	resp := &proto.ShowBranchAppointmentBusinessTimeResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(request.BranchId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	date := time.Now()

	branchAppointment, err := model.ShowBranchAppointment(branchID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "门店预约数据不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowBranchAppointmentBusinessTime 查询门店预约数据数据库错误:%v", err))
		resp.ErrorMessage = "查询门店预约数据失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	if request.Date != 0 {
		date = time.Unix(int64(request.Date), 0)
	}
	businesses, err := model.ShowBranchBusinessByBranchIDDate(branchID, merchantID, date)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "数据未找到"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowBranchAppointmentBusinessTime 查询门店营业时间数据库错误:%v", err))
		resp.ErrorMessage = "查询门店预约数据失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	openAppointment := false
	if branchAppointment.OpenAppointment == 1 {
		openAppointment = true
	}

	resp.Data = &proto.ShowBranchAppointmentBusinessTimeData{
		BeginTime:       businesses.BeginTime.String(),
		EndTime:         businesses.EndTime.String(),
		IsNextDay:       businesses.IsNextDay,
		Granularity:     int32(branchAppointment.AppointmentGranularity),
		OpenAppointment: openAppointment,
		Hot:             branchAppointment.Hot,
	}

	return resp, nil
}
