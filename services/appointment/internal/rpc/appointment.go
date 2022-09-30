package rpc

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	redigo "github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	mqMessage "gitlab.omytech.com.cn/micro-service/Crius/pkgs/message"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/config"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gitlab.omytech.com.cn/micro-service/appointment/util"
	merchantBasic "gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

const appointmentURL = "%s/api/third/branch/preorder"

// CreateAppointment 创建预约
func (s *Server) CreateAppointment(ctx context.Context, req *proto.CreateAppointmentRequest) (*proto.CreateAppointmentResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreateAppointment")
	resp := &proto.CreateAppointmentResponse{
		ErrorCode: pkgs.Success,
	}

	packageID := uuid.FromStringOrNil(req.ThemePackageId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	operator := pkgs.GetMetadata(ctx).StaffID
	agent := pkgs.GetMetadata(ctx).UserAgent
	branchID := uuid.FromStringOrNil(req.Appointment.BranchId)
	lockID := uuid.FromStringOrNil(req.LockId)
	memberID := uuid.FromStringOrNil(req.Appointment.MemberId)
	roomTypeID := uuid.FromStringOrNil(req.Appointment.RoomTypeId)
	appointmentAt := time.Unix(int64(req.Appointment.AppointmentAt), 0)
	appointmentDate := time.Unix(int64(req.Appointment.AppointmentDate), 0)
	relatedID := uuid.NewV4()
	depositFee := req.Appointment.DepositFee
	var extend model.TableAppointmentExtend

	if agent == pkgs.UserAgentWechat {
		c, err := s.showTemplateConfig(merchantID, branchID, roomTypeID, appointmentDate)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("CreateAppointment 获取模板配置错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "创建预约失败"
			return resp, nil
		}
		depositFee = c.DepositFee
	}

	// 主题套餐
	if packageID != uuid.Nil {
		p, err := model.ShowAppointmentThemePackage(packageID)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("CreateAppointment 获取主题套餐数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "创建预约失败"
			return resp, nil
		}
		roomTypes := make([]*proto.AppointmentThemePackage_RoomType, 0)
		err = json.Unmarshal([]byte(p.RoomTypes.JSON()), &roomTypes)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("CreateAppointment 解析主题套餐数据错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "创建预约失败"
			return resp, nil
		}
		packagePrice := int32(0)
		for _, v := range roomTypes {
			if v.Id == roomTypeID.String() {
				depositFee += v.Price
				packagePrice = v.Price
				break
			}
		}
		extend = model.TableAppointmentExtend{
			PackageID:     &p.ID,
			Packages:      p.Packages,
			Decoration:    p.Decoration,
			Staffing:      p.Staffing,
			CustomConfigs: p.CustomConfigs,
			ThemeID:       p.ThemeID,
			PackageName:   p.Name,
			PackagePrice:  packagePrice,
		}
	}

	appointmentConfig, err := s.getAppointmentConfig(merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 获取预约配置错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建预约失败"
		return resp, nil
	}

	keepAt := appointmentAt.Add(time.Duration(appointmentConfig.KeepTime) * time.Minute)
	if extend.ThemeID != nil {
		// 主题预约保留时间
		keepAt = appointmentAt.Add(time.Duration(appointmentConfig.ThemeKeepTime) * time.Minute)
	}

	//过期时间
	businessTime, err := s.getBranchBusinessTime(ctx, branchID, int32(appointmentDate.Unix()))
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 获取门店营业时间错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建预约失败"
		return resp, nil
	}
	hms, err := time.Parse("15:04:05", businessTime.EndTime)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 获取门店营业时间解析错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建预约失败"
		return resp, nil
	}
	expireAt := appointmentDate.Add(time.Duration(hms.Hour()) * time.Hour).Add(time.Duration(hms.Minute()) * time.Minute).
		Add(time.Duration(hms.Second()) * time.Second).Add(time.Duration(hms.Nanosecond()) * time.Nanosecond)
	if businessTime.IsNextDay {
		expireAt = expireAt.Add(time.Hour * 24)
	}

	appointment := model.TableAppointment{
		ID:               uuid.NewV4(),
		MerchantID:       &merchantID,
		BranchID:         &branchID,
		CalledCode:       req.Appointment.CalledCode,
		CalledPhone:      req.Appointment.CalledPhone,
		AppointmentCode:  req.Appointment.AppointmentCode,
		AppointmentPhone: req.Appointment.AppointmentPhone,
		Name:             req.Appointment.Name,
		Gender:           int8(req.Appointment.Gender),
		Way:              int8(req.Appointment.Way),
		CustomerNum:      int16(req.Appointment.CustomerNum),
		AppointmentAt:    &appointmentAt,
		AppointmentDate:  &appointmentDate,
		ExpireAt:         &expireAt,
		RoomTypeID:       &roomTypeID,
		DepositFee:       depositFee,
		FlowerCake:       req.Appointment.FlowerCake,
		FlowerCakeRemark: req.Appointment.FlowerCakeRemark,
		Remark:           req.Appointment.Remark,
		Status:           model.AppointmentStatusAppointed,
		RelatedID:        &relatedID,
		ChargingWay:      int8(req.Appointment.ChargingWay),
		KeepAt:           &keepAt,
		MemberID:         &memberID,
	}
	if req.Appointment.CalledPhone != "" {
		appointment.CalledCode = req.Appointment.CalledCode
		appointment.CalledPhone = req.Appointment.CalledPhone
	} else {
		appointment.CalledCode = req.Appointment.AppointmentCode
		appointment.CalledPhone = req.Appointment.AppointmentPhone
	}
	if appointment.CalledCode == "" {
		appointment.CalledCode = "86"
	}
	extend.AppointmentID = appointment.ID
	code, err := model.ShowAppointmentCodeSeq()
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateAppointment 生成预约号错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约失败"
		return resp, nil
	}
	extend.Code = code
	if operator != uuid.Nil {
		appointment.Operator = &operator
	}

	if memberID != uuid.Nil {
		appointment.MemberID = &memberID
	}

	if err := s.createAppointmentLogic(lockID, &appointment, &extend, appointmentConfig.PaymentTime); err != nil {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建预约失败"
		return resp, nil
	}

	ap := model.Appointment{TableAppointment: appointment, TableAppointmentExtend: extend}
	// 不需要支付，预约直接成功，通知咨客系统
	if appointment.DepositFee == 0 {
		go s.AppointmentHTTPRequest(&ap)
	}
	after := pkgs.MakeParams(ap)
	go s.SaveSnapshot(ctx, model.TableSnapshot{
		SnapShotTableName: model.TableAppointment{}.TableName(),
		TableID:           nil,
		Method:            model.CreateMethod,
		After:             &after,
		RelatedID:         appointment.RelatedID,
		StaffID:           operator,
	})
	resp.Data = appointment.ID.String()
	taskMessage := mqMessage.TaskMessage{
		Category: mqMessage.Appointment,
		MemberID: memberID,
		Time:     time.Now(),
	}
	bs, _ := json.Marshal(&taskMessage)
	if err := s.mq.Publish(config.Setting.Mqtt.TaskTopic, 2, true, bs); err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 预约任务推送错误:%v", err))
	}

	return resp, nil
}

// UpdateAppointment 更新预约
func (s *Server) UpdateAppointment(ctx context.Context, req *proto.UpdateAppointmentRequest) (*proto.UpdateAppointmentResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointment")
	resp := &proto.UpdateAppointmentResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	agent := pkgs.GetMetadata(ctx).UserAgent
	id := uuid.FromStringOrNil(req.Appointment.Id)
	operator := pkgs.GetMetadata(ctx).StaffID
	branchID := uuid.FromStringOrNil(req.Appointment.BranchId)
	lockID := uuid.FromStringOrNil(req.LockId)
	roomTypeID := uuid.FromStringOrNil(req.Appointment.RoomTypeId)
	appointmentAt := time.Unix(int64(req.Appointment.AppointmentAt), 0)
	appointmentDate := time.Unix(int64(req.Appointment.AppointmentDate), 0)
	memberID := uuid.FromStringOrNil(req.Appointment.MemberId)
	depositFee := req.Appointment.DepositFee
	packageID := uuid.FromStringOrNil(req.ThemePackageId)
	if agent == pkgs.UserAgentWechat {
		c, err := s.showTemplateConfig(merchantID, branchID, roomTypeID, appointmentDate)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("UpdateAppointment 获取模板配置错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "更新预约失败"
			return resp, nil
		}
		depositFee = c.DepositFee
	}

	//过期时间
	businessTime, err := s.getBranchBusinessTime(ctx, branchID, int32(appointmentDate.Unix()))
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateAppointment 获取门店营业时间错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约失败"
		return resp, nil
	}
	hms, err := time.Parse("15:04:05", businessTime.EndTime)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateAppointment 获取门店营业时间解析错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约失败"
		return resp, nil
	}
	expireAt := appointmentDate.Add(time.Duration(hms.Hour()) * time.Hour).Add(time.Duration(hms.Minute()) * time.Minute).
		Add(time.Duration(hms.Second()) * time.Second).Add(time.Duration(hms.Nanosecond()) * time.Nanosecond)
	if businessTime.IsNextDay {
		expireAt = expireAt.Add(time.Hour * 24)
	}
	tabAppointment := new(model.Appointment)
	if err := s.database.Conn.Model(&model.TableAppointment{}).Select("appointment.*", "appointment_extend.*").Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").
		Scopes(crius.ColumnEqualScope("id", id)).Take(tabAppointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约数据未找到"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdateAppointment 更新预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约失败"
		return resp, nil
	}
	oldBefore := pkgs.MakeParams(tabAppointment)

	// 保存AppointmentDateCounter id
	dateCount := new(model.TableAppointmentDateCounter)
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnEqualScope("branch_id", branchID),
		crius.ColumnEqualScope("appointment_time", appointmentAt), model.WayAnd(int8(req.Appointment.Way)),
		crius.ColumnEqualScope("room_group_id", roomTypeID)).Take(&dateCount).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约房间不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdateAppointment 获取预约剩余数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约失败"
		return resp, nil
	}

	ap := model.TableAppointment{
		ID:               uuid.NewV4(),
		MerchantID:       &merchantID,
		BranchID:         &branchID,
		AppointmentCode:  req.Appointment.AppointmentCode,
		AppointmentPhone: req.Appointment.AppointmentPhone,
		Name:             req.Appointment.Name,
		Gender:           int8(req.Appointment.Gender),
		Way:              int8(req.Appointment.Way),
		CustomerNum:      int16(req.Appointment.CustomerNum),
		AppointmentAt:    &appointmentAt,
		AppointmentDate:  &appointmentDate,
		ExpireAt:         &expireAt,
		RoomTypeID:       &roomTypeID,
		DepositFee:       depositFee,
		FlowerCake:       req.Appointment.FlowerCake,
		FlowerCakeRemark: req.Appointment.FlowerCakeRemark,
		Remark:           req.Appointment.Remark,
		Operator:         &operator,
		RelatedID:        tabAppointment.RelatedID,
		DateCounterID:    &dateCount.ID,
		ChargingWay:      int8(req.Appointment.ChargingWay),
		CalledPhone:      req.Appointment.CalledPhone,
		CalledCode:       req.Appointment.CalledCode,
		MemberID:         &memberID,
	}
	if operator != uuid.Nil {
		ap.Operator = &operator
	}

	var extend model.TableAppointmentExtend
	// 主题套餐
	if packageID != uuid.Nil {
		p, err := model.ShowAppointmentThemePackage(packageID)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("UpdateAppointment 获取主题套餐数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "更新预约失败"
			return resp, nil
		}
		roomTypes := make([]*proto.AppointmentThemePackage_RoomType, 0)
		err = json.Unmarshal([]byte(p.RoomTypes.JSON()), &roomTypes)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("UpdateAppointment 解析主题套餐数据错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "更新预约失败"
			return resp, nil
		}
		packagePrice := int32(0)
		for _, v := range roomTypes {
			if v.Id == roomTypeID.String() {
				ap.DepositFee += v.Price
				packagePrice = v.Price
				break
			}
		}
		extend = model.TableAppointmentExtend{
			PackageID:     &p.ID,
			Packages:      p.Packages,
			Decoration:    p.Decoration,
			Staffing:      p.Staffing,
			CustomConfigs: p.CustomConfigs,
			ThemeID:       p.ThemeID,
			PackageName:   p.Name,
			PackagePrice:  packagePrice,
		}
	}
	extend.AppointmentID = ap.ID
	code, err := model.ShowAppointmentCodeSeq()
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateAppointment 生成预约号错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约失败"
		return resp, nil
	}
	extend.Code = code

	appointment := model.Appointment{TableAppointment: ap, TableAppointmentExtend: extend}
	oldMethod, method, err := s.updateAppointment(*tabAppointment, appointment, lockID)
	if err != nil {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约失败"
		return resp, nil
	}

	if oldMethod != "" {
		oldAfter := pkgs.MakeParams(tabAppointment)
		go s.SaveSnapshot(ctx, model.TableSnapshot{
			SnapShotTableName: model.TableAppointment{}.TableName(),
			TableID:           nil,
			Method:            oldMethod,
			Before:            &oldBefore,
			After:             &oldAfter,
			RelatedID:         tabAppointment.RelatedID,
			StaffID:           operator,
		})
	}
	resp.Data = &proto.UpdateAppointmentData{
		AppointmentId: tabAppointment.ID.String(),
	}
	if method.String() != "" {
		after := pkgs.MakeParams(appointment)
		go s.SaveSnapshot(ctx, model.TableSnapshot{
			SnapShotTableName: model.TableAppointment{}.TableName(),
			TableID:           nil,
			Method:            method,
			After:             &after,
			RelatedID:         appointment.RelatedID,
			StaffID:           operator,
		})
		resp.Data.NewAppointmentId = appointment.ID.String()
	}

	return resp, nil
}

// PayAppointment 支付预约款
func (s *Server) PayAppointment(ctx context.Context, req *proto.PayAppointmentRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("PayAppointment")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}
	var staffID *uuid.UUID
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if pkgs.GetMetadata(ctx).StaffID != uuid.Nil {
		staffID = new(uuid.UUID)
		*staffID = pkgs.GetMetadata(ctx).StaffID
	}
	id := uuid.FromStringOrNil(req.Id)

	appointment := new(model.Appointment) // 要修改的订单
	if err := s.database.Conn.Model(&model.TableAppointment{}).Select("appointment.*", "appointment_extend.*").Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").
		Scopes(crius.ColumnEqualScope("id", id)).Take(appointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约数据不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("PayAppointment 查询预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}

	if appointment.Status != model.AppointmentStatusArrearage {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "预约状态错误"
		return resp, nil
	}
	before := pkgs.MakeParams(appointment)

	config, err := s.getAppointmentConfig(merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("PayAppointment 获取预约配置错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}

	keepTime := config.KeepTime
	if appointment.PackageID != nil {
		keepTime = config.ThemeKeepTime
	}

	keepAt := appointment.AppointmentAt.Add(time.Duration(keepTime) * time.Minute)
	appointment.KeepAt = &keepAt
	appointment.Status = model.AppointmentStatusAppointed
	appointment.Operator = staffID

	oldAppointment := new(model.Appointment) // 以前更新订单可能产生的老订单
	oldAppointmentSend := true               // 是否通知咨客旧订单取消
	// 当新的预约订单状态变成已预约，旧的预约取消
	if err := s.database.Conn.Model(&model.TableAppointment{}).Select("appointment.*", "appointment_extend.*").Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").
		Scopes(crius.ColumnEqualScope("appointment.related_id", appointment.RelatedID), crius.ColumnNotEqualScope("appointment.id", appointment.ID),
			crius.ColumnInScope("appointment.status", []interface{}{model.AppointmentStatusAppointed, model.AppointmentStatusArrearage})).Take(oldAppointment).Error; err != nil && err != gorm.ErrRecordNotFound {
		crius.Logger.Error(fmt.Sprintf("PayAppointment 获取旧预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	} else if err != gorm.ErrRecordNotFound { // 旧预约存在，取消旧预约
		if oldAppointment.Status == model.AppointmentStatusArrearage {
			oldAppointment.Status = model.AppointmentStatusCancelled
			oldAppointmentSend = false // 旧订单状态为待支付，取消不需要通知咨客
		} else { // 旧订单支付完成
			if oldAppointment.DepositFee == 0 {
				oldAppointment.Status = model.AppointmentStatusCancelled
			} else {
				oldAppointment.Status = model.AppointmentStatusRefunding
				now := time.Now()
				oldAppointment.RefundingAt = &now
			}
		}
		oldAppointment.CancelledReason = model.AppointmentChangeCancelReason
		oldAppointment.Operator = staffID
		oldAppointment.Sended = false
	}
	after := pkgs.MakeParams(appointment)

	tx := s.database.Conn.Begin()
	// 更新新预约
	if err := tx.Save(&(appointment.TableAppointment)).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("PayAppointment 更新预约状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	if err := tx.Scopes(crius.ColumnEqualScope("appointment_id", id)).Save(&(appointment.TableAppointmentExtend)).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("PayAppointment 更新预约状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}

	// 更新旧预约
	if oldAppointment.ID != uuid.Nil {
		if err := tx.Save(&oldAppointment.TableAppointment).Error; err != nil {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("PayAppointment 保存旧预约数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "更新预约状态失败"
			return resp, nil
		}
		if err := tx.Scopes(crius.ColumnEqualScope("appointment_id", oldAppointment.ID)).Save(&oldAppointment.TableAppointmentExtend).Error; err != nil {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("PayAppointment 保存旧预约数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "更新预约状态失败"
			return resp, nil
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("PayAppointment 提交事务数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	// 通知咨客
	go s.AppointmentHTTPRequest(appointment)
	if oldAppointment.ID != uuid.Nil && oldAppointmentSend {
		go s.AppointmentHTTPRequest(oldAppointment)
	}

	snapshot := model.TableSnapshot{
		SnapShotTableName: model.TableAppointment{}.TableName(),
		TableID:           nil,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
		RelatedID:         appointment.RelatedID,
	}
	if staffID != nil {
		snapshot.StaffID = *staffID
	}
	go s.SaveSnapshot(ctx, snapshot)

	return resp, nil
}

// CancelAppointment 取消预约
func (s *Server) CancelAppointment(ctx context.Context, req *proto.CancelAppointmentRequest) (*proto.CancelAppointmentResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CancelAppointment")
	resp := &proto.CancelAppointmentResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	var staffID *uuid.UUID
	if pkgs.GetMetadata(ctx).StaffID != uuid.Nil {
		staffID = new(uuid.UUID)
		*staffID = pkgs.GetMetadata(ctx).StaffID
	}
	id := uuid.FromStringOrNil(req.Id)
	appointment := new(model.Appointment) // 要修改的订单快照的变量
	if err := s.database.Conn.Model(&model.TableAppointment{}).Select("appointment.*", "appointment_extend.*").Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").
		Scopes(crius.ColumnEqualScope("id", id)).Take(appointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约数据不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("CancelAppointment 查询预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	if appointment.Status != model.AppointmentStatusArrearage && appointment.Status != model.AppointmentStatusAppointed {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "预约状态错误"
		return resp, nil
	}
	before := pkgs.MakeParams(appointment)

	config, err := s.getAppointmentConfig(merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CancelAppointment 获取预约配置错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}

	cancelTime := config.CancelTime
	refundAfter := config.RefundPercentAfter
	refundBefore := config.RefundPercentBefore
	if appointment.PackageID != nil {
		cancelTime = config.ThemeCancelTime
		refundBefore = config.ThemeRefundPercentBefore
		refundAfter = config.ThemeRefundPercentAfter
	}

	appointmentSend := true
	if appointment.Status == model.AppointmentStatusArrearage { // 待支付，不通知咨客
		appointmentSend = false
	}

	if appointment.Status == model.AppointmentStatusAppointed && appointment.DepositFee != 0 { // 已预约有订金，状态为退款
		appointment.Status = model.AppointmentStatusRefunding
		now := time.Now()
		appointment.RefundingAt = &now
	} else { // 没有订金直接取消
		appointment.Status = model.AppointmentStatusCancelled
	}
	appointment.CancelledReason = req.OperateMessage
	appointment.RefundAmount = appointment.DepositFee * int32(refundBefore) / 100      // 提前退款金额
	if util.BeyondAppointCancelTime(float64(cancelTime), *appointment.AppointmentAt) { // 超过时间退款金额及记录违约
		appointment.RefundAmount = appointment.DepositFee * int32(refundAfter) / 100
		appointment.Breach = true
		appointment.BreachReason = model.AppointmentBreachReasonCancelLate
	}
	appointment.Operator = staffID

	after := pkgs.MakeParams(appointment)

	tx := s.database.Conn.Begin()
	if err := tx.Save(&(appointment.TableAppointment)).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CancelAppointment 更新预约状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	if err := tx.Scopes(crius.ColumnEqualScope("appointment_id", id)).Save(&(appointment.TableAppointmentExtend)).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CancelAppointment 更新预约状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	if err := tx.Model(&model.TableAppointmentDateCounter{}).Scopes(crius.ColumnEqualScope("id", appointment.DateCounterID)).
		Update("appoint_num", gorm.Expr("appoint_num-1")).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CancelAppointment 预约数量减少数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CancelAppointment 提交事务数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	if appointmentSend { // 通知咨客
		go s.AppointmentHTTPRequest(appointment)
	}

	snapshot := model.TableSnapshot{
		SnapShotTableName: model.TableAppointment{}.TableName(),
		TableID:           nil,
		Method:            model.CancelMethod,
		Before:            &before,
		After:             &after,
		RelatedID:         appointment.RelatedID,
	}
	if staffID != nil {
		snapshot.StaffID = *staffID
	}
	go s.SaveSnapshot(ctx, snapshot)
	resp.Data = appointment.RefundAmount
	return resp, nil
}

// ArriveAppointment 预约到店
func (s *Server) ArriveAppointment(ctx context.Context, req *proto.ArriveAppointmentRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ArriveAppointment")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	var staffID *uuid.UUID
	if pkgs.GetMetadata(ctx).StaffID != uuid.Nil {
		staffID = new(uuid.UUID)
		*staffID = pkgs.GetMetadata(ctx).StaffID
	}
	id := uuid.FromStringOrNil(req.Id)

	appointment := new(model.Appointment) // 要修改的订单快照的变量
	if err := s.database.Conn.Model(&model.TableAppointment{}).Select("appointment.*", "appointment_extend.*").Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").
		Scopes(crius.ColumnEqualScope("id", id)).Take(appointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约数据不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ArriveAppointment 查询预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	if appointment.Status != model.AppointmentStatusAppointed {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "预约状态错误"
		return resp, nil
	}
	before := pkgs.MakeParams(appointment)

	if time.Now().After(*appointment.KeepAt) {
		appointment.Breach = true
		appointment.BreachReason = model.AppointmentBreachReasonLate
	}
	arrive := time.Unix(int64(req.ArrivedAt), 0)
	appointment.ArrivedAt = &arrive
	appointment.Operator = staffID

	after := pkgs.MakeParams(appointment)

	tx := s.database.Conn.Begin()
	if err := tx.Save(&(appointment.TableAppointment)).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("ArriveAppointment 更新预约状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	if err := tx.Scopes(crius.ColumnEqualScope("appointment_id", id)).Save(&(appointment.TableAppointmentExtend)).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("ArriveAppointment 更新预约状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("ArriveAppointment 提交事务数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}

	snapshot := model.TableSnapshot{
		SnapShotTableName: model.TableAppointment{}.TableName(),
		TableID:           nil,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
		RelatedID:         appointment.RelatedID,
	}
	if staffID != nil {
		snapshot.StaffID = *staffID
	}
	go s.SaveSnapshot(ctx, snapshot)

	return resp, nil
}

// RefundAppointment 预约退款成功
func (s *Server) RefundAppointment(ctx context.Context, req *proto.RefundAppointmentRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("RefundAppointment")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	var staffID *uuid.UUID
	if pkgs.GetMetadata(ctx).StaffID != uuid.Nil {
		staffID = new(uuid.UUID)
		*staffID = pkgs.GetMetadata(ctx).StaffID
	}
	id := uuid.FromStringOrNil(req.Id)

	appointment := new(model.Appointment) // 要修改的订单快照的变量
	if err := s.database.Conn.Model(&model.TableAppointment{}).Select("appointment.*", "appointment_extend.*").Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").
		Scopes(crius.ColumnEqualScope("id", id)).Take(appointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约数据不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("RefundAppointment 查询预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	if appointment.Status != model.AppointmentStatusRefunding {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "预约状态错误"
		return resp, nil
	}
	before := pkgs.MakeParams(appointment)

	now := time.Now()
	appointment.RefundedAt = &now
	appointment.Status = model.AppointmentStatusRefunded
	appointment.Operator = staffID
	after := pkgs.MakeParams(appointment)

	tx := s.database.Conn.Begin()
	if err := tx.Save(&(appointment.TableAppointment)).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("RefundAppointment 更新预约状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}
	if err := tx.Scopes(crius.ColumnEqualScope("appointment_id", id)).Save(&(appointment.TableAppointmentExtend)).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("RefundAppointment 更新预约状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("RefundAppointment 提交事务数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约状态失败"
		return resp, nil
	}

	snapshot := model.TableSnapshot{
		SnapShotTableName: model.TableAppointment{}.TableName(),
		TableID:           nil,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
		RelatedID:         appointment.RelatedID,
	}
	if staffID != nil {
		snapshot.StaffID = *staffID
	}
	go s.SaveSnapshot(ctx, snapshot)

	return resp, nil
}

// GetAppointments 获取预约数据列表
func (s *Server) GetAppointments(ctx context.Context, req *proto.GetAppointmentsRequest) (*proto.GetAppointmentsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointments")
	resp := &proto.GetAppointmentsResponse{
		ErrorCode: pkgs.Success,
	}

	var dateStart, dateEnd time.Time
	var count int64
	var appointments []model.Appointment
	userAgent := pkgs.GetMetadata(ctx).UserAgent
	branchID := uuid.FromStringOrNil(req.BranchId)
	memberID := uuid.FromStringOrNil(req.MemberId)
	wechatID := uuid.FromStringOrNil(req.WechatId)
	themeID := uuid.FromStringOrNil(req.ThemeId)
	packageID := uuid.FromStringOrNil(req.PackageId)
	roomTypeIds := make([]uuid.UUID, 0)
	if req.DateStart != 0 {
		dateStart = time.Unix(int64(req.DateStart), 0)
	}
	if req.DateEnd != 0 {
		dateEnd = time.Unix(int64(req.DateEnd), 0)
	}
	for _, v := range req.RoomTypeId {
		if id := uuid.FromStringOrNil(v); id != uuid.Nil {
			roomTypeIds = append(roomTypeIds, id)
		}
	}

	var scopes = []func(db *gorm.DB) *gorm.DB{
		crius.ColumnInScopeDefault("appointment.room_type_id", roomTypeIds),
		crius.ColumnEqualScopeDefault("appointment.branch_id", branchID),
		model.AppointmentPhoneTailLike(req.PhoneTail),
		crius.ColumnInScopeDefault("appointment.status", req.Status),
		crius.ColumnEqualScopeDefault("appointment.appointment_phone", req.Phone),
		crius.ColumnEqualScopeDefault("appointment.way", int8(req.Way)),
		model.IsDepositPayed(int8(req.IsDepositPayed)),
		model.DateRangeCondition(dateStart, dateEnd),
		model.FlowerCakeCondition(int8(req.FlowerCake)),
		crius.ColumnEqualScopeDefault("appointment.called_phone", req.CalledPhone),
		crius.ColumnEqualScopeDefault("appointment.called_code", req.CalledCode),
		model.AnyPhoneMemberIDWechatIDCondition(req.AnyPhone, memberID, wechatID),
	}
	if themeID != uuid.Nil && packageID == uuid.Nil {
		scopes = append(scopes, crius.ColumnEqualScopeDefault("appointment_extend.theme_id", themeID))
	}
	if packageID != uuid.Nil {
		scopes = append(scopes, crius.ColumnEqualScopeDefault("appointment_extend.package_id", packageID))
	}
	db := s.database.Conn.Model(&model.TableAppointment{}).
		Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").
		Scopes(scopes...)
	//当请求方式预约系统后台，不过滤删除预约
	if userAgent == pkgs.UserAgentWeb {
		db = db.Unscoped()
	}
	if err := db.Count(&count).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointments 获取列表数量数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取预约列表失败"
		return resp, nil
	}
	if count == 0 {
		resp.Data = &proto.GetAppointmentsData{Total: 0}
		return resp, nil
	}

	if err := db.Select("*").Scopes(model.PagingCondition(req.Offset, req.Limit)).Order("appointment.created_at desc").Find(&appointments).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointments 获取列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取预约列表失败"
		return resp, nil
	}

	resp.Data = &proto.GetAppointmentsData{
		Total: int32(count),
	}

	for _, v := range appointments {
		theme := new(proto.ThemeRoomType)
		// 如果套餐不为空，查询主题和套餐
		if v.PackageID != nil {
			p := toProtoAppointmentThemePackage(&model.TableAppointmentThemePackage{
				Name:          v.PackageName,
				Packages:      v.Packages,
				Decoration:    v.Decoration,
				Staffing:      v.Staffing,
				CustomConfigs: v.CustomConfigs,
				ID:            *v.PackageID,
				RoomTypes:     &pkgs.ParamsArr{{"id": v.RoomTypeID, "price": v.PackagePrice}},
			})
			themeCategory, err := model.ShowAppointmentThemeWithCategory(*v.ThemeID)
			if err != nil && err != gorm.ErrRecordNotFound {
				crius.Logger.Error(fmt.Sprintf("GetWechatAppointments 查询预约主题数据库错误:%v", err))
				resp.ErrorCode = pkgs.ErrInternal
				resp.ErrorMessage = "查询预约数据失败"
				return resp, nil
			}
			theme = &proto.ThemeRoomType{AppointmentTheme: toProtoAppointmentTheme(&themeCategory.TableAppointmentTheme), ThemePackage: p, CategoryName: themeCategory.CategoryName}
		}
		resp.Data.Appointments = append(resp.Data.Appointments, &proto.GetAppointmentsData_Data{
			Appointment: toProtoFullAppointment(v),
			Theme:       theme,
		})
	}

	return resp, nil
}

// GetAppointmentRemaining 获取剩余预约数据列表
func (s *Server) GetAppointmentRemaining(ctx context.Context, req *proto.GetAppointmentRemainingRequest) (*proto.GetAppointmentRemainingResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentRemaining")
	resp := &proto.GetAppointmentRemainingResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	branchID := uuid.FromStringOrNil(req.BranchId)
	businessDate := time.Unix(int64(req.BusinessDate), 0)

	var calendars []model.TableTemplateCalendar
	var dateCounter []model.TableAppointmentDateCounter
	var configs []model.TableAppointmentTemplateConfig
	var templateID uuid.UUID

	if err := s.deleteExpiredAppointmentLock(merchantID); err != nil {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取剩余预约列表失败"
		return resp, nil
	}

	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnEqualScope("branch_id", branchID),
		crius.ColumnEqualScope("business_date", businessDate)).Find(&calendars).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentRemaining 获取模板日历数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取剩余预约列表失败"
		return resp, nil
	}
	if len(calendars) > 1 {
		for _, v := range calendars {
			if v.Category == model.CalendarHoliday {
				templateID = v.TemplateID
			}
		}
	} else if len(calendars) == 1 {
		templateID = calendars[0].TemplateID
	} else {
		// 没有数据
		return resp, nil
	}

	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("template_id", templateID)).Find(&configs).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentRemaining 获取模板配置数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取剩余预约列表失败"
		return resp, nil
	}

	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnEqualScope("branch_id", branchID),
		crius.ColumnEqualScope("appointment_date", businessDate)).Find(&dateCounter).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentRemaining 获取临时日历数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取剩余预约列表失败"
		return resp, nil
	}

	remaining := getAppointmentRemainingRespData(configs, businessDate, dateCounter, req.Way)

	for _, v := range remaining {
		resp.Data = append(resp.Data, v)
	}

	return resp, nil
}

// GetAppointmentRemainingFast 快速预约获取剩余预约数据列表
func (s *Server) GetAppointmentRemainingFast(ctx context.Context, req *proto.GetAppointmentRemainingFastRequest) (*proto.GetAppointmentRemainingFastResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentRemainingFast")
	resp := &proto.GetAppointmentRemainingFastResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	businessDate := time.Unix(int64(req.BusinessDate), 0)

	var calendars []model.TableTemplateCalendar
	var configs []model.AppointmentTemplateConfig
	businessTime := fields.StringToLocalTime(req.Time)

	if err := s.deleteExpiredAppointmentLock(merchantID); err != nil {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取剩余预约列表失败"
		return resp, nil
	}

	//门店距离关联map
	branchMap := make(map[uuid.UUID]int32)
	branchIDs := make([]uuid.UUID, 0)
	for _, v := range req.Branches {
		branchID := uuid.FromStringOrNil(v.BranchId)
		branchIDs = append(branchIDs, branchID)
		branchMap[branchID] = v.Distance
	}

	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnInScopeDefault("branch_id", branchIDs),
		crius.ColumnInScope("business_date", []interface{}{businessDate.Add(-time.Hour * 24), businessDate})).Find(&calendars).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentRemainingFast 获取模板日历数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取剩余预约列表失败"
		return resp, nil
	}

	//日历过滤map
	calendarMap := make(map[string]uuid.UUID)
	businessTimeMap := make(map[string]merchantBasic.ShowBranchAppointmentBusinessTimeData)
	templateIDs := make([]interface{}, 0)
	for _, calendar := range calendars {
		if _, ok := calendarMap[fmt.Sprintf("%v_%v", calendar.BranchID, calendar.BusinessDate.Time.Format("2006-01-02"))]; !ok || calendar.Category == model.CalendarHoliday {
			businessTimeResp, err := s.merchantBasic().ShowBranchAppointmentBusinessTime(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &merchantBasic.ShowBranchAppointmentBusinessTimeRequest{
				BranchId: calendar.BranchID.String(),
				Date:     int32(calendar.BusinessDate.ToUnix()),
			})
			if err != nil {
				crius.Logger.Error(fmt.Sprintf("GetAppointmentRemainingFast 获取模板门店营业时间错误err:%v, resp:%v", err, businessTimeResp))
				resp.ErrorCode = pkgs.ErrInternal
				resp.ErrorMessage = "获取剩余预约列表失败"
				return resp, nil
			}
			if businessTimeResp.ErrorCode != pkgs.Success {
				// 查询营业时间不存在，当天不营业
				if businessTimeResp.ErrorCode == pkgs.ErrNotFound {
					continue
				}
				crius.Logger.Error(fmt.Sprintf("GetAppointmentRemainingFast 获取模板门店营业时间错误 resp:%v", businessTimeResp))
				resp.ErrorCode = pkgs.ErrInternal
				resp.ErrorMessage = "获取剩余预约列表失败"
				return resp, nil
			}
			// 没有开放预约
			if !businessTimeResp.Data.OpenAppointment {
				continue
			}

			businessTimeMap[fmt.Sprintf("%v_%v", calendar.BranchID, calendar.BusinessDate.Time.Format("2006-01-02"))] = *businessTimeResp.Data
			calendarMap[fmt.Sprintf("%v_%v", calendar.BranchID, calendar.BusinessDate.Time.Format("2006-01-02"))] = calendar.TemplateID
			templateIDs = append(templateIDs, calendar.TemplateID)
		}
	}

	templates := make([]model.TableAppointmentTemplate, 0)
	if err := s.database.Conn.Scopes(crius.ColumnInScope("id", templateIDs)).Find(&templates).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentRemainingFast 获取模板数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取剩余预约列表失败"
		return resp, nil
	}

	branchTemplateMap := make(map[uuid.UUID]templateBusinessDate)
	yesterdayTemplateIDs := make([]interface{}, 0)
	todayTemplateIDs := make([]interface{}, 0)
	for _, template := range templates {
		for k, v := range calendarMap {
			if v == template.ID {
				strs := strings.Split(k, "_")
				branchID := uuid.FromStringOrNil(strs[0])
				date, _ := time.Parse("2006-01-02", strs[1])
				endDate := date
				var beginTime, endTime string
				if template.BeginTime.String() > businessTimeMap[k].BeginTime {
					beginTime = template.BeginTime.String()
				} else {
					beginTime = businessTimeMap[k].BeginTime
				}
				start, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", date.Format("2006-01-02"), beginTime))

				isNextDay := false
				if template.IsNextDay == 1 {
					isNextDay = true
				}
				if isNextDay == businessTimeMap[k].IsNextDay {
					if template.EndTime.String() > businessTimeMap[k].EndTime {
						endTime = businessTimeMap[k].BeginTime
					} else {
						endTime = template.EndTime.String()
					}
					if isNextDay {
						endDate = date.Add(time.Hour * 24)
					}
				} else {
					if !isNextDay {
						endTime = template.EndTime.String()
					} else {
						endTime = businessTimeMap[k].BeginTime
					}
				}
				end, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", endDate.Format("2006-01-02"), endTime))
				appointmentAt, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", businessDate.Format("2006-01-02"), req.Time))

				log.Println(start, end, appointmentAt)
				if !appointmentAt.After(end) && !appointmentAt.Before(start) {
					if date.Before(businessDate) {
						yesterdayTemplateIDs = append(yesterdayTemplateIDs, template.ID)
					} else {
						todayTemplateIDs = append(todayTemplateIDs, template.ID)
					}
					branchTemplateMap[branchID] = templateBusinessDate{templateID: template.ID, businessDate: date}
				}
			}
		}
	}

	// 获取前一天的预约时间数据
	var todayConfigs, yesterdayConfigs []model.AppointmentTemplateConfig
	if len(yesterdayTemplateIDs) != 0 {
		if err := s.database.Conn.Table("appointment.appointment_template_configure t1").Select("t1.id, t1.room_type_id, t1.template_id, t1.advance_day, t1.deposit_fee, value").
			Joins("left join jsonb_array_elements(t1.configure) c on 1=1").Scopes(crius.ColumnInScope("template_id", yesterdayTemplateIDs),
			model.ConfigureTimeCondition(businessTime, true), model.ConfigureWayAnd(int8(req.Way))).Order("(c::json->>'time')::time asc").Find(&yesterdayConfigs).Error; err != nil {
			crius.Logger.Error(fmt.Sprintf("GetAppointmentRemainingFast 获取模板配置数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "获取剩余预约列表失败"
			return resp, nil
		}
	}

	// 获取当天的预约时间数据
	if len(todayTemplateIDs) != 0 {
		if err := s.database.Conn.Table("appointment.appointment_template_configure t1").Select("t1.id, t1.room_type_id, t1.template_id, t1.advance_day, t1.deposit_fee, value").
			Joins("left join jsonb_array_elements(t1.configure) c on 1=1").Scopes(crius.ColumnInScope("template_id", todayTemplateIDs),
			model.ConfigureTimeCondition(businessTime, false), model.ConfigureWayAnd(int8(req.Way))).Order("(c::json->>'is_next_day')::bool asc, (c::json->>'time')::time asc").Find(&todayConfigs).Error; err != nil {
			crius.Logger.Error(fmt.Sprintf("GetAppointmentRemainingFast 获取模板配置数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "获取剩余预约列表失败"
			return resp, nil
		}
	}

	// 整合两天的预约时段数据
	if len(yesterdayConfigs) == 0 {
		configs = todayConfigs
	} else if len(todayConfigs) == 0 {
		configs = yesterdayConfigs
	} else {
		for _, yesterday := range yesterdayConfigs {
			for i := 0; i < len(todayConfigs); {
				if yesterday.Value.Get("time").(string) > todayConfigs[i].Value.Get("time").(string) && !(todayConfigs[i].Value.Get("is_next_day").(bool)) {
					configs = append(configs, todayConfigs[i])
					todayConfigs = append(todayConfigs[:i], todayConfigs[i+1:]...)
				} else {
					configs = append(configs, yesterday)
					break
				}
			}
		}
		configs = append(configs, todayConfigs...)
	}

	resp.Data = &proto.GetAppointmentRemainingFastData{Recommends: s.getAppointmentRemainingFastRespData(ctx, branchIDs, configs, branchTemplateMap, req.Num, req.Limit, merchantID, branchMap)}

	return resp, nil
}

// GetAppointmentRecord 获取预约操作记录
func (s *Server) GetAppointmentRecord(ctx context.Context, req *proto.GetAppointmentRecordRequest) (*proto.GetAppointmentRecordResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentRecord")
	resp := &proto.GetAppointmentRecordResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	tabAppointment := new(model.TableAppointment)
	if err := s.database.Conn.Unscoped().Scopes(crius.ColumnEqualScope("id", id)).Take(tabAppointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约数据未找到"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("GetAppointmentRecord 获取预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取预约操作记录失败"
		return resp, nil
	}

	snapshots := make([]model.TableSnapshot, 0)
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("related_id", tabAppointment.RelatedID)).Find(&snapshots).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentRecord 获取快照数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取预约操作记录失败"
		return resp, nil
	}

	for _, v := range snapshots {
		resp.Data = append(resp.Data, toProtoAppointmentSnapshot(v))
	}
	return resp, nil
}

// SaveAppointmentTempNumber 修改临时总数
func (s *Server) SaveAppointmentTempNumber(ctx context.Context, req *proto.SaveAppointmentTempNumberRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("SaveAppointmentTempNumber")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	branchID := uuid.FromStringOrNil(req.BranchId)
	roomTypeID := uuid.FromStringOrNil(req.RoomTypeId)
	appointmentAt := time.Unix(int64(req.AppointmentAt), 0)
	appointmentDate := time.Unix(int64(req.AppointmentDate), 0)

	tx := s.database.Conn.Begin()
	dateCounter := new(model.TableAppointmentDateCounter)
	err := tx.Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnEqualScope("branch_id", branchID), crius.ColumnEqualScope("appointment_time", appointmentAt),
		model.WayAnd(int8(req.Way)), crius.ColumnEqualScope("room_group_id", roomTypeID)).Take(dateCounter).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentTempNumber 查询临时总数数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "修改总数失败"
		return resp, nil
	}

	if err == gorm.ErrRecordNotFound {
		dateCounter.ID = uuid.NewV4()
		dateCounter.MerchantID = &merchantID
		dateCounter.BranchID = &branchID
		dateCounter.AppointmentTime = &appointmentAt
		dateCounter.Way = int8(req.Way)
		dateCounter.RoomGroupID = &roomTypeID
		dateCounter.Number = req.Num
		dateCounter.AppointmentDate = &appointmentDate
		if err := tx.Create(dateCounter).Error; err != nil {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("SaveAppointmentTempNumber 创建临时总数数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "修改总数失败"
			return resp, nil
		}
	} else if err := tx.Model(&model.TableAppointmentDateCounter{}).Scopes(crius.ColumnEqualScope("id", dateCounter.ID)).Update("number", req.Num).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentTempNumber 修改临时总数数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "修改总数失败"
		return resp, nil
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentTempNumber 提交数据库事务错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "修改总数失败"
		return resp, nil
	}

	return resp, nil
}

// ShowAppointment 查询单条预约数据
func (s *Server) ShowAppointment(ctx context.Context, req *proto.ShowAppointmentRequest) (*proto.ShowAppointmentResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowAppointment")
	resp := &proto.ShowAppointmentResponse{
		ErrorCode: pkgs.Success,
	}
	id := uuid.FromStringOrNil(req.Id)

	appointment := new(model.Appointment)
	if err := s.database.Conn.Unscoped().Select("appointment.*", "appointment_extend.*").Model(&model.TableAppointment{}).Scopes(crius.ColumnEqualScope("appointment.id", id)).
		Joins("left join appointment.appointment_extend on appointment.id=appointment_extend.appointment_id").Take(appointment).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("ShowAppointment 查询数据库错误:%v", err))
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "未找到预约数据"
			return resp, nil
		}
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "查询预约数据失败"
		return resp, nil
	}

	resp.Data = &proto.ShowAppointmentResponse_Data{}

	if appointment.ThemeID != nil {
		themeCategory, err := model.ShowAppointmentThemeWithCategory(*appointment.ThemeID)
		if err != nil && err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("ShowAppointment 查询预约主题数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "查询预约数据失败"
			return resp, nil
		}
		if err == nil {
			p := toProtoAppointmentThemePackage(&model.TableAppointmentThemePackage{
				Name:          appointment.PackageName,
				Packages:      appointment.Packages,
				Decoration:    appointment.Decoration,
				Staffing:      appointment.Staffing,
				CustomConfigs: appointment.CustomConfigs,
				ID:            *appointment.PackageID,
				RoomTypes:     &pkgs.ParamsArr{{"id": appointment.RoomTypeID, "price": appointment.PackagePrice}},
			})
			resp.Data.Theme = &proto.ThemeRoomType{AppointmentTheme: toProtoAppointmentTheme(&themeCategory.TableAppointmentTheme), ThemePackage: p, CategoryName: themeCategory.CategoryName}
		}
	}

	resp.Data.Appointment = toProtoFullAppointment(*appointment)
	return resp, nil
}

// UpdateAppointmentTradeInfo 更新预约交易信息
func (s *Server) UpdateAppointmentTradeInfo(ctx context.Context, req *proto.UpdateAppointmentTradeInfoRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentTradeInfo")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	id := uuid.FromStringOrNil(req.Id)
	tradeID := uuid.FromStringOrNil(req.TradeId)
	refundID := uuid.FromStringOrNil(req.RefundId)

	tabAppointment := new(model.Appointment)
	if err := s.database.Conn.Unscoped().Model(&model.TableAppointment{}).Select("appointment.*", "appointment_extend.*").Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").
		Scopes(crius.ColumnEqualScope("id", id)).Take(tabAppointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约数据不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdateAppointmentTradeInfo 查询预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约交易失败"
		return resp, nil
	}

	if tabAppointment.Status != model.AppointmentStatusArrearage && req.Status != "" {
		crius.Logger.Error(fmt.Sprintf("UpdateAppointmentTradeInfo 预约状态不能完成更改,%v", tabAppointment.Status))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "更新预约交易失败"
		return resp, nil
	}

	if model.AppointmentStatus(req.Status) == model.AppointmentStatusAppointed {
		config, err := s.getAppointmentConfig(merchantID)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("UpdateAppointmentTradeInfo 获取预约配置错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "更新预约交易失败"
			return resp, nil
		}
		keepAt := tabAppointment.AppointmentAt.Add(time.Duration(config.KeepTime) * time.Minute)
		if tabAppointment.AppointmentID != uuid.Nil {
			keepAt = tabAppointment.AppointmentAt.Add(time.Duration(config.ThemeKeepTime) * time.Minute)
		}
		tabAppointment.KeepAt = &keepAt
	}

	if tradeID != uuid.Nil {
		tabAppointment.TradeID = &tradeID
	}

	tabAppointment.Status = model.AppointmentStatus(req.Status)
	tabAppointment.TradeType = req.TradeType

	tx := s.database.Conn.Begin()
	if err := tx.Updates(tabAppointment.TableAppointment).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateAppointmentTradeInfo 更新预约支付信息数据库错误:%v", err))
		tx.Rollback()
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约交易失败"
		return resp, nil
	}
	if refundID != uuid.Nil {
		if err := tx.Scopes(crius.ColumnEqualScope("appointment_id", tabAppointment.ID)).Update("refund_id", refundID).Error; err != nil {
			crius.Logger.Error(fmt.Sprintf("UpdateAppointmentTradeInfo 更新预约退款信息数据库错误:%v", err))
			tx.Rollback()
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "更新预约交易失败"
			return resp, nil
		}
	}
	if err := tx.Commit().Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateAppointmentTradeInfo 事务提交数据库错误:%v", err))
		tx.Rollback()
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约交易失败"
		return resp, nil
	}

	return resp, nil
}

// GetAppointmentLimit 获取手机号对应预约是否限制
func (s *Server) GetAppointmentLimit(ctx context.Context, req *proto.GetAppointmentLimitRequest) (*proto.GetAppointmentLimitResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentLimit")
	resp := &proto.GetAppointmentLimitResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	memberID := uuid.FromStringOrNil(req.MemberId)
	var start, end time.Time
	if req.DateStart != 0 && req.DateEnd != 0 {
		start, end = time.Unix(int64(req.DateStart), 0), time.Unix(int64(req.DateEnd), 0)
	}

	var appointments []model.TableAppointment
	if err := s.database.Conn.Scopes(crius.ColumnEqualScopeDefault("appointment_code", req.PhoneCode), crius.ColumnEqualScopeDefault("appointment_phone", req.Phone),
		crius.ColumnEqualScopeDefault("status", req.Status), model.DateRangeCondition(start, end), model.AnyPhoneMemberIDCondition(req.AnyPhone, memberID)).Find(&appointments).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentLimit 获取预约配置错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取预约限制失败"
		return resp, nil
	}
	branchNumMap := make(map[uuid.UUID]int32)
	for _, v := range appointments {
		branchNumMap[*v.BranchID]++
	}
	wayNumMap := make(map[int8]int32)
	for _, v := range appointments {
		wayNumMap[v.Way]++
	}
	resp.Data = &proto.GetAppointmentLimitData{}
	for k, v := range branchNumMap {
		resp.Data.BranchNum = append(resp.Data.BranchNum, &proto.BranchAppointmentNum{BranchId: k.String(), Num: v})
	}
	for k, v := range wayNumMap {
		resp.Data.WayNum = append(resp.Data.WayNum, &proto.WayAppointmentNum{Num: v, Way: int32(k)})
	}
	if model.AppointmentStatus(req.Status) == model.AppointmentStatusAppointed {
		config, err := s.getAppointmentConfig(merchantID)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("GetAppointmentLimit 获取预约配置错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "获取预约限制失败"
			return resp, nil
		}
		resp.Data.Limit = int32(config.OrderLimit)
	}

	return resp, nil
}

// DeleteAppointment 删除预约
func (s *Server) DeleteAppointment(ctx context.Context, req *proto.DeleteAppointmentRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("DeleteAppointment")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	appointment := new(model.TableAppointment)
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("id", id)).Take(appointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, nil
		}
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除预约失败"
		return resp, nil
	}

	if appointment.Status != model.AppointmentStatusExpired && appointment.Status != model.AppointmentStatusCancelled {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "预约状态不允许删除"
		return resp, nil
	}

	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("id", id)).Delete(&model.TableAppointment{}).Error; err != nil {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除预约失败"
		return resp, nil
	}
	return resp, nil
}

// GetAppointmentBreachLimit 获取违约订单数量
func (s *Server) GetAppointmentBreachLimit(ctx context.Context, req *proto.GetAppointmentBreachLimitRequest) (*proto.GetAppointmentBreachLimitResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentBreachLimit")
	resp := &proto.GetAppointmentBreachLimitResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	config, err := s.getAppointmentConfig(merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentBreachLimit 获取预约配置错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取违约订单数量失败"
		return resp, nil
	}
	now := time.Now()
	start := now.Add(-time.Duration(config.BreachMonths*30*24) * time.Hour)

	var num int64
	if err := s.database.Conn.Model(&model.TableAppointment{}).Scopes(crius.ColumnEqualScope("appointment_phone", req.Phone), crius.ColumnEqualScopeDefault("appointment_code", req.PhoneCode),
		crius.ColumnEqualScope("breach", true), model.DateRangeCondition(start, now)).Count(&num).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentBreachLimit 获取违约数量数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取违约订单数量失败"
		return resp, nil
	}

	resp.Data = &proto.GetAppointmentBreachLimitData{Num: int32(num), Limit: int32(config.BreachTotal)}
	return resp, nil
}

// AppointmentOpenRoom 开房
func (s *Server) AppointmentOpenRoom(ctx context.Context, req *proto.AppointmentOpenRoomRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("AppointmentOpenRoom")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	openAt := time.Unix(int64(req.OpenAt), 0)
	roomID := uuid.FromStringOrNil(req.RoomId)
	appointment, err := model.ShowFullAppointment(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		crius.Logger.Error(fmt.Sprintf("AppointmentOpenRoom 获取预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "开房失败"
		return resp, nil
	}
	if err == gorm.ErrRecordNotFound {
		crius.Logger.Error(fmt.Sprintf("AppointmentOpenRoom 获取预约不存在:%v", id))
		resp.ErrorCode = pkgs.ErrNotFound
		resp.ErrorMessage = "预约不存在"
		return resp, nil
	}
	before := pkgs.MakeParams(appointment)

	appointment.OpenRoomID = &roomID
	appointment.OpenAt = &openAt
	appointment.Status = model.AppointmentStatusArrived
	if err := model.OpenRoom(id, roomID, openAt); err != nil {
		crius.Logger.Error(fmt.Sprintf("AppointmentOpenRoom 预约开房数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "开房失败"
		return resp, nil
	}

	after := pkgs.MakeParams(appointment)

	go s.SaveSnapshot(ctx, model.TableSnapshot{
		SnapShotTableName: model.TableAppointment{}.TableName(),
		TableID:           nil,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
		RelatedID:         appointment.RelatedID,
		StaffID:           pkgs.GetMetadata(ctx).StaffID,
	})
	return resp, nil
}

// AppointmentCancelOpenRoom 取消开房
func (s *Server) AppointmentCancelOpenRoom(ctx context.Context, req *proto.AppointmentCancelOpenRoomRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("AppointmentCancelOpenRoom")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	appointment, err := model.ShowFullAppointment(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		crius.Logger.Error(fmt.Sprintf("AppointmentCancelOpenRoom 获取预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "取消开房失败"
		return resp, nil
	}
	if err == gorm.ErrRecordNotFound {
		crius.Logger.Error(fmt.Sprintf("AppointmentCancelOpenRoom 获取预约不存在:%v", id))
		resp.ErrorCode = pkgs.ErrNotFound
		resp.ErrorMessage = "预约不存在"
		return resp, nil
	}
	before := pkgs.MakeParams(appointment)

	breach := false
	var breachReason model.AppointmentBreachReason
	status := model.AppointmentStatusAppointed
	if appointment.ExpireAt != nil && time.Now().After(*appointment.ExpireAt) {
		breach = true
		status = model.AppointmentStatusExpired
		breachReason = model.AppointmentBreachReasonNotArrive
	}
	appointment.ArrivedAt = nil
	appointment.Status = status
	appointment.Breach = breach
	appointment.BreachReason = breachReason
	appointment.OpenRoomID = nil
	appointment.OpenAt = nil
	if err := model.CancelOpenRoom(appointment); err != nil {
		crius.Logger.Error(fmt.Sprintf("AppointmentOpenRoom 预约开房数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "取消开房失败"
		return resp, nil
	}

	after := pkgs.MakeParams(appointment)

	go s.SaveSnapshot(ctx, model.TableSnapshot{
		SnapShotTableName: model.TableAppointment{}.TableName(),
		TableID:           nil,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
		RelatedID:         appointment.RelatedID,
		StaffID:           pkgs.GetMetadata(ctx).StaffID,
	})
	return resp, nil
}

func (s *Server) getBranchBusinessTime(ctx context.Context, branchID uuid.UUID, date int32) (*merchantBasic.ShowBranchAppointmentBusinessTimeData, error) {
	data := new(merchantBasic.ShowBranchAppointmentBusinessTimeData)
	bs, err := s.cache.Get(fmt.Sprintf(model.BranchAppointmentBusinessTimeKey, branchID.String(), date))
	if err == nil {
		err = json.Unmarshal(bs, data)
		if err == nil {
			return data, nil
		}
	}

	businessTime, err := s.merchantBasic().ShowBranchAppointmentBusinessTime(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &merchantBasic.ShowBranchAppointmentBusinessTimeRequest{BranchId: branchID.String(), Date: date})
	if err != nil {
		return nil, err
	}
	if businessTime.ErrorCode != pkgs.Success {
		return nil, fmt.Errorf("获取门店营业时间失败")
	}
	bs, err = json.Marshal(businessTime.Data)
	if err != nil {
		return nil, fmt.Errorf("json解析门店营业时间失败")
	}
	s.cache.Set(fmt.Sprintf(model.BranchAppointmentBusinessTimeKey, branchID.String(), date), bs)
	return businessTime.Data, nil
}

func (s *Server) deleteExpiredAppointmentLock(merchantID uuid.UUID) error {
	// 删除过期锁
	zset, err := s.redis.ZRANGEWITHSCORES(model.AppointmentZsetLockKey, 0, -1)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("删除过期锁redis错误:%v", err))
		return err
	}
	zsetMembers := make([]interface{}, 0)
	hashKeys := make([]interface{}, 0)
	tx := s.database.Conn.Begin()
	for k, v := range zset {
		expireTime, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("解析锁时间戳错误:%v", err))
			return err
		}
		if util.AppointLockExpired(expireTime) {
			zsetMembers = append(zsetMembers, k)
			hashKeys = append(hashKeys, fmt.Sprintf(model.AppointmentHashLockKey, k))
			appointmentLock := new(model.AppointmentLock)
			if err := s.redis.HGETALLSTRUCT(fmt.Sprintf(model.AppointmentHashLockKey, k), appointmentLock); err != nil {
				crius.Logger.Error(fmt.Sprintf("获取锁redis错误:%v", err))
				continue
			}
			if err := tx.Model(&model.TableAppointmentDateCounter{}).Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnEqualScope("branch_id", appointmentLock.BranchID),
				crius.ColumnEqualScope("appointment_time", time.Unix(int64(appointmentLock.AppointmentAt), 0)), model.WayAnd(appointmentLock.Way),
				crius.ColumnEqualScope("room_group_id", appointmentLock.RoomGroupID)).Update("appoint_num", gorm.Expr("appoint_num-1")).Error; err != nil {
				tx.Rollback()
				crius.Logger.Error(fmt.Sprintf("锁定数量自减数据库错误:%v", err))
				return err
			}
		}
	}

	if len(zsetMembers) > 0 {
		if n, err := s.redis.ZREM(model.AppointmentZsetLockKey, zsetMembers...); err != nil || n == 0 {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("删除锁redis错误:%v", err))
			return fmt.Errorf("删除redis锁数量:%v, 错误:%v", n, err)
		}
	}
	s.redis.DEL(hashKeys...)
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("事务提交数据库错误:%v", err))
		return err
	}
	return nil
}

func toProtoFullAppointment(a model.Appointment) *proto.Appointment {
	wechatIDs := make([]*proto.Appointment_WechatID, 0)
	if a.WechatIDs != nil {
		for _, v := range *a.WechatIDs {
			wechatIDs = append(wechatIDs, &proto.Appointment_WechatID{
				Id:   v.ID.String(),
				Time: int32(v.Time.Unix()),
			})
		}
	}
	return &proto.Appointment{
		Id:               a.ID.String(),
		CalledCode:       a.CalledCode,
		CalledPhone:      a.CalledPhone,
		AppointmentCode:  a.AppointmentCode,
		AppointmentPhone: a.AppointmentPhone,
		Name:             a.Name,
		Gender:           int32(a.Gender),
		Way:              int32(a.Way),
		CustomerNum:      int32(a.CustomerNum),
		BranchId:         crius.UUIDToString(a.BranchID),
		AppointmentAt:    crius.TimeUnix32(a.AppointmentAt),
		AppointmentDate:  crius.TimeUnix32(a.AppointmentDate),
		ExpireAt:         crius.TimeUnix32(a.ExpireAt),
		RoomTypeId:       crius.UUIDToString(a.RoomTypeID),
		DepositFee:       a.DepositFee,
		FlowerCake:       a.FlowerCake,
		FlowerCakeRemark: a.FlowerCakeRemark,
		Remark:           a.Remark,
		Operator:         crius.UUIDToString(a.Operator),
		Status:           string(a.Status),
		CancelledReason:  a.CancelledReason,
		ChargingWay:      int32(a.ChargingWay),
		TradeId:          crius.UUIDToString(a.TradeID),
		RefundId:         crius.UUIDToString(a.RefundID),
		KeepAt:           crius.TimeUnix32(a.KeepAt),
		CancelAt:         crius.TimeUnix32(a.CancelAt),
		TradeType:        a.TradeType,
		ArrivedAt:        crius.TimeUnix32(a.ArrivedAt),
		RefundingAt:      crius.TimeUnix32(a.RefundingAt),
		RefundedAt:       crius.TimeUnix32(a.RefundedAt),
		MemberId:         crius.UUIDToString(a.MemberID),
		RefundAmount:     a.RefundAmount,
		CreatedAt:        crius.TimeUnix32(a.CreatedAt),
		ShareMessage:     a.ShareMessage,
		WechatIds:        wechatIDs,
		Code:             a.Code,
		Breach:           a.Breach,
		BreachReason:     string(a.BreachReason),
	}
}

func toProtoAppointment(a model.TableAppointment) *proto.Appointment {
	wechatIDs := make([]*proto.Appointment_WechatID, 0)
	if a.WechatIDs != nil {
		for _, v := range *a.WechatIDs {
			wechatIDs = append(wechatIDs, &proto.Appointment_WechatID{
				Id:   v.ID.String(),
				Time: int32(v.Time.Unix()),
			})
		}
	}
	return &proto.Appointment{
		Id:               a.ID.String(),
		CalledCode:       a.CalledCode,
		CalledPhone:      a.CalledPhone,
		AppointmentCode:  a.AppointmentCode,
		AppointmentPhone: a.AppointmentPhone,
		Name:             a.Name,
		Gender:           int32(a.Gender),
		Way:              int32(a.Way),
		CustomerNum:      int32(a.CustomerNum),
		BranchId:         crius.UUIDToString(a.BranchID),
		AppointmentAt:    crius.TimeUnix32(a.AppointmentAt),
		AppointmentDate:  crius.TimeUnix32(a.AppointmentDate),
		ExpireAt:         crius.TimeUnix32(a.ExpireAt),
		RoomTypeId:       crius.UUIDToString(a.RoomTypeID),
		DepositFee:       a.DepositFee,
		FlowerCake:       a.FlowerCake,
		FlowerCakeRemark: a.FlowerCakeRemark,
		Remark:           a.Remark,
		Operator:         crius.UUIDToString(a.Operator),
		Status:           string(a.Status),
		CancelledReason:  a.CancelledReason,
		ChargingWay:      int32(a.ChargingWay),
		TradeId:          crius.UUIDToString(a.TradeID),
		KeepAt:           crius.TimeUnix32(a.KeepAt),
		CancelAt:         crius.TimeUnix32(a.CancelAt),
		TradeType:        a.TradeType,
		ArrivedAt:        crius.TimeUnix32(a.ArrivedAt),
		MemberId:         crius.UUIDToString(a.MemberID),
		CreatedAt:        crius.TimeUnix32(a.CreatedAt),
		WechatIds:        wechatIDs,
		Breach:           a.Breach,
		BreachReason:     string(a.BreachReason),
	}

}

func toProtoAppointmentSnapshot(s model.TableSnapshot) *proto.AppointmentRecord {
	var before, after, staffID string
	if s.Before != nil {
		before = s.Before.JSON()
	}
	if s.After != nil {
		after = s.After.JSON()
	}
	if s.StaffID != uuid.Nil {
		staffID = s.StaffID.String()
	}
	return &proto.AppointmentRecord{
		Operate:    s.Method.String(),
		OperatorId: staffID,
		CreatedAt:  int32(s.CreatedAt.Unix()),
		Before:     before,
		After:      after,
	}
}

func getAppointmentRemainingRespData(configs []model.TableAppointmentTemplateConfig, businessDate time.Time, dateCounter []model.TableAppointmentDateCounter, way int32) map[string]*proto.AppointmentRemaining {
	roomNum := make(map[string]struct {
		Num       int32
		Total     int32
		IsNextDay bool
	})
	roomFee := make(map[string]int32)
	// 模板设置房型数量
	now := time.Now()
	now = now.Add(-(time.Duration(now.Hour()) * time.Hour) - (time.Duration(now.Minute()) * time.Minute) - (time.Duration(now.Second()) * time.Second) - (time.Duration(now.Nanosecond()) * time.Nanosecond))
	for _, v := range configs {
		configure := []map[string]interface{}(*v.Configure)
		for _, cfg := range configure {
			if businessDate.Sub(now) > time.Hour*24*time.Duration(v.AdvanceDay) || (way != 0 && way&int32(cfg["way"].(float64)) != way) {
				continue
			}
			value := roomNum[fmt.Sprintf("%s,%s,%d", v.RoomTypeID.String(), cfg["time"].(string), int32(cfg["way"].(float64)))]
			value.Total = int32(cfg["num"].(float64))
			value.Num = value.Total
			value.IsNextDay = cfg["is_next_day"].(bool)
			roomNum[fmt.Sprintf("%s,%s,%d", v.RoomTypeID.String(), cfg["time"].(string), int32(cfg["way"].(float64)))] = value
		}
		roomFee[v.RoomTypeID.String()] = v.DepositFee
	}
	// 临时设置房型数量
	for _, v := range dateCounter {
		if way != 0 && way&int32(v.Way) != way {
			continue
		}
		value, ok := roomNum[fmt.Sprintf("%s,%s,%d", v.RoomGroupID.String(), v.AppointmentTime.Format("15:04:05"), v.Way)]
		if !ok {
			continue
		}
		if v.Number != 0 {
			value.Total = v.Number
		}
		value.Num = value.Total - v.AppointNum
		roomNum[fmt.Sprintf("%s,%s,%d", v.RoomGroupID.String(), v.AppointmentTime.Format("15:04:05"), v.Way)] = value
	}

	remaining := make(map[string]*proto.AppointmentRemaining)
	for k, v := range roomNum {
		strs := strings.Split(k, ",")
		if len(strs) != 3 {
			crius.Logger.Error(fmt.Sprintf("GetAppointmentRemaining 预约剩余数据错误:roomID,time,way  %v", k))
			continue
		}
		id, t := strs[0], strs[1]
		way, _ := strconv.ParseInt(strs[2], 10, 64)
		value, ok := remaining[id]
		if !ok {
			value = &proto.AppointmentRemaining{
				RoomTypeId:  id,
				DepositFee:  roomFee[id],
				RoomTypeNum: []*proto.RoomTypeNum{{Way: int32(way), Num: v.Num, Time: t, Total: v.Total, IsNextDay: v.IsNextDay}},
			}
			remaining[id] = value
		} else {
			value.RoomTypeNum = append(value.RoomTypeNum, &proto.RoomTypeNum{Way: int32(way), Num: v.Num, Time: t, Total: v.Total, IsNextDay: v.IsNextDay})
		}
	}

	return remaining
}

func (s *Server) getAppointmentRemainingFastRespData(ctx context.Context, branchIDs []uuid.UUID, configs []model.AppointmentTemplateConfig, calendarMap map[uuid.UUID]templateBusinessDate,
	num, limit int32, merchantID uuid.UUID, branchMap map[uuid.UUID]int32) []*proto.Recommend {
	now := time.Now()
	now = now.Add(-(time.Duration(now.Hour()) * time.Hour) - (time.Duration(now.Minute()) * time.Minute) - (time.Duration(now.Second()) * time.Second) - (time.Duration(now.Nanosecond()) * time.Nanosecond))

	branchTimeMap := make(map[string]bool)
	recommends := make([]*proto.Recommend, 0)
	for {
		for i := 0; i < len(branchIDs); {
			before := len(branchTimeMap)

			t, isNextDay, rooms := s.appointmentTemplateConfig(branchIDs[i], merchantID, configs, calendarMap, branchTimeMap)

			roomTypeIDs := make([]string, 0)
			roomMap := make(map[uuid.UUID]int32)
			for _, v := range rooms {
				roomTypeIDs = append(roomTypeIDs, v.roomTypeID.String())
				roomMap[v.roomTypeID] = v.depositFee
			}

			//前后map数量没有变化，此门店循环完毕
			if before == len(branchTimeMap) {
				//数组删除循环完毕门店
				branchIDs = append(branchIDs[:i], branchIDs[i+1:]...)
				continue
			}

			roomTypeResp, err := s.merchantBasic().GetRoomTypesByIDs(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &merchantBasic.GetRoomTypesByIDsRequest{Ids: roomTypeIDs})
			if err != nil || roomTypeResp.ErrorCode != pkgs.Success {
				crius.Logger.Error(fmt.Sprintf("GetAppointmentRemainingFast 获取房型列表rpc错误:%v", err))
				i++
				continue
			}
			for _, v := range roomTypeResp.Data {
				if v.CustomerMax < num || v.CustomerMin > num {
					//房型人数不匹配
					continue
				}
				recommends = append(recommends, &proto.Recommend{
					BranchId:    branchIDs[i].String(),
					Time:        t,
					IsNextDay:   isNextDay,
					RoomTypeId:  v.Id,
					RoomName:    v.Name,
					Distance:    branchMap[branchIDs[i]],
					CustomerMax: v.CustomerMax,
					CustomerMin: v.CustomerMin,
					DepositFee:  roomMap[uuid.FromStringOrNil(v.Id)],
				})

				if len(recommends) == int(limit) {
					return recommends
				}
			}
			i++
		}
		// 门店循环完毕跳出循环
		if len(branchIDs) == 0 {
			break
		}
	}
	return recommends
}

func (s *Server) appointmentTemplateConfig(branchID, merchantID uuid.UUID, configs []model.AppointmentTemplateConfig,
	calendarMap map[uuid.UUID]templateBusinessDate, branchTimeMap map[string]bool) (string, bool, []roomTypeDepositFee) {
	var t string
	var isNextDay bool
	rooms := make([]roomTypeDepositFee, 0)
	for _, config := range configs {
		//允许提前预约时间
		if calendarMap[branchID].businessDate.Sub(time.Now()) >= time.Hour*24*time.Duration(config.AdvanceDay) {
			continue
		}
		appointmentAt, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", calendarMap[branchID].businessDate.Format("2006-01-02"), config.Value.Get("time").(string)))
		if config.Value.Get("is_next_day").(bool) {
			appointmentAt.Add(time.Hour * 24)
		}
		if *config.TemplateID == calendarMap[branchID].templateID && t == "" && !branchTimeMap[fmt.Sprintf("%v%v", branchID, appointmentAt)] {
			isNextDay = config.Value.Get("is_next_day").(bool)
			t = config.Value.Get("time").(string)
			branchTimeMap[fmt.Sprintf("%v%v", branchID, appointmentAt)] = true
		}
		if t != config.Value.Get("time").(string) || *config.TemplateID != calendarMap[branchID].templateID {
			continue
		}

		dateCounter := new(model.TableAppointmentDateCounter)
		if err := s.database.Conn.Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnEqualScope("branch_id", branchID),
			crius.ColumnEqualScope("appointment_time", appointmentAt), crius.ColumnEqualScope("room_group_id", config.RoomTypeID),
			crius.ColumnEqualScope("way", config.Value.Get("way").(float64))).Take(dateCounter).Error; err != nil && err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("GetAppointmentRemainingFast 获取已预约数量数据库错误:%v", err))
			continue
		}
		//预约数已满
		if (dateCounter.Number > 0 && dateCounter.Number <= dateCounter.AppointNum) || (dateCounter.Number <= 0 && dateCounter.AppointNum >= int32(config.Value.Get("num").(float64))) {
			continue
		}
		rooms = append(rooms, roomTypeDepositFee{roomTypeID: *config.RoomTypeID, depositFee: config.DepositFee})
	}
	return t, isNextDay, rooms
}

func (s *Server) createAppointmentLogic(lockID uuid.UUID, appointment *model.TableAppointment, extend *model.TableAppointmentExtend, paymentTime int16) error {
	score, err := s.redis.ZSCORE(model.AppointmentZsetLockKey, lockID.String())
	if err != nil && err != redigo.ErrNil {
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 获取锁过期时间redis错误:%v", err))
		return err
	}
	if err == redigo.ErrNil || util.AppointLockExpired(int64(score)) {
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 预约锁过期"))
		return errors.New("预约锁过期")
	}

	// 保存AppointmentDateCounter id
	dateCount := new(model.TableAppointmentDateCounter)
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("merchant_id", appointment.MerchantID), crius.ColumnEqualScope("branch_id", appointment.BranchID),
		crius.ColumnEqualScope("appointment_time", appointment.AppointmentAt), model.WayAnd(appointment.Way),
		crius.ColumnEqualScope("room_group_id", appointment.RoomTypeID)).Take(&dateCount).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("CreateAppointment 获取预约剩余不存在:%v", err))
			return err
		}
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 获取预约剩余数据库错误:%v", err))
		return err
	}
	appointment.DateCounterID = &dateCount.ID

	if appointment.DepositFee > 0 {
		appointment.Status = model.AppointmentStatusArrearage
		keepAt := time.Now().Add(time.Duration(paymentTime) * time.Minute)
		appointment.KeepAt = &keepAt
	}

	tx := s.database.Conn.Begin()
	if err := tx.Create(appointment).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 创建预约数据库错误:%v", err))
		return err
	}

	appointmentCode, err := model.ShowAppointmentCodeSeq()
	if err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 创建预约号错误:%v", err))
		return err
	}
	extend.Code = appointmentCode
	if err := tx.Create(extend).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 创建预约扩展数据库错误:%v", err))
		return err
	}

	if n, err := s.redis.ZREM(model.AppointmentZsetLockKey, lockID.String()); err != nil || n == 0 {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 删除锁redis错误:%v", err))
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CreateAppointment 提交事务数据库错误:%v", err))
		return err
	}

	// 删除锁数据，不处理错误
	s.redis.DEL(fmt.Sprintf(model.AppointmentHashLockKey, lockID.String()))
	return nil
}

func (s *Server) updateAppointment(tabAppointment, appointment model.Appointment, lockID uuid.UUID) (model.SnapshotMethod, model.SnapshotMethod, error) {
	var oldMethod, method model.SnapshotMethod
	now := time.Now()
	config, err := s.getAppointmentConfig(*appointment.MerchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateAppointment 获取预约配置错误:%v", err))
		return oldMethod, method, err
	}
	oldDateCounterID := tabAppointment.DateCounterID
	sended := false

	tx := s.database.Conn.Begin()
	// 如果门店改变或者订金改变
	if tabAppointment.DepositFee != appointment.DepositFee || *tabAppointment.BranchID != *appointment.BranchID ||
		*tabAppointment.RoomTypeID != *appointment.RoomTypeID || *tabAppointment.AppointmentAt != *appointment.AppointmentAt ||
		!((tabAppointment.PackageID == nil && appointment.PackageID == nil) || (tabAppointment.PackageID != nil && appointment.PackageID != nil && *tabAppointment.PackageID == *appointment.PackageID)) {

		score, err := s.redis.ZSCORE(model.AppointmentZsetLockKey, lockID.String())
		if err != nil && err != redigo.ErrNil {
			crius.Logger.Error(fmt.Sprintf("UpdateAppointment 获取锁过期时间redis错误:%v", err))
			return oldMethod, method, err
		}
		if err == redigo.ErrNil || util.AppointLockExpired(int64(score)) {
			crius.Logger.Error(fmt.Sprintf("UpdateAppointment 预约锁过期"))
			return oldMethod, method, errors.New("预约锁过期")
		}

		appointment.Status = model.AppointmentStatusAppointed
		keepAt := appointment.AppointmentAt.Add(time.Duration(config.KeepTime) * time.Minute)
		if appointment.AppointmentID != uuid.Nil {
			keepAt = appointment.AppointmentAt.Add(time.Duration(config.ThemeKeepTime) * time.Minute)
		}

		// 如果订金不为0，则订单状态为未支付
		if appointment.DepositFee > 0 {
			appointment.Status = model.AppointmentStatusArrearage
			keepAt = time.Now().Add(time.Duration(config.PaymentTime) * time.Minute)
		} else { // 新订单订金为0，则旧订单直接取消
			if tabAppointment.DepositFee == 0 {
				tabAppointment.Status = model.AppointmentStatusCancelled
			} else {
				tabAppointment.Status = model.AppointmentStatusRefunding
				tabAppointment.RefundingAt = &now
			}
			tabAppointment.CancelledReason = model.AppointmentChangeCancelReason
			tabAppointment.Operator = appointment.Operator
			tabAppointment.Sended = false

			// 更新旧预约状态
			if err := tx.Scopes(crius.ColumnEqualScope("id", tabAppointment.ID)).Select("status", "cancelled_reason", "operator").Updates(tabAppointment.TableAppointment).Error; err != nil {
				tx.Rollback()
				crius.Logger.Error(fmt.Sprintf("UpdateAppointment 更新预约数据库错误:%v", err))
				return oldMethod, method, err
			}
			// 更新旧预约扩展状态
			if err := tx.Scopes(crius.ColumnEqualScope("appointment_id", tabAppointment.ID)).Select("refunding_at", "sended").Updates(tabAppointment.TableAppointmentExtend).Error; err != nil {
				tx.Rollback()
				crius.Logger.Error(fmt.Sprintf("UpdateAppointment 更新预约数据库错误:%v", err))
				return oldMethod, method, err
			}
			// 取消旧订单，生成新订单 通知咨客
			sended = true
			oldMethod = model.CancelMethod
			// 旧预约房型数量释放
			if err := tx.Model(&model.TableAppointmentDateCounter{}).Scopes(crius.ColumnEqualScope("id", oldDateCounterID)).
				Update("appoint_num", gorm.Expr("appoint_num-1")).Error; err != nil {
				tx.Rollback()
				crius.Logger.Error(fmt.Sprintf("UpdateAppointment 预约数量减少数据库错误:%v", err))
				return oldMethod, method, err
			}
		}
		appointment.KeepAt = &keepAt

		// 创建新预约
		if err := tx.Create(&appointment.TableAppointment).Error; err != nil {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("UpdateAppointment 创建预约数据库错误:%v", err))
			return oldMethod, method, err
		}
		if err := tx.Create(&appointment.TableAppointmentExtend).Error; err != nil {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("UpdateAppointment 创建预约数据库错误:%v", err))
			return oldMethod, method, err
		}
		method = model.CreateMethod

		if n, err := s.redis.ZREM(model.AppointmentZsetLockKey, lockID.String()); err != nil || n == 0 {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("UpdateAppointment 删除锁redis错误:%v", err))
			return oldMethod, method, err
		}
		// 删除锁数据，不处理错误
		s.redis.DEL(fmt.Sprintf(model.AppointmentHashLockKey, lockID.String()))
	} else {
		tabAppointment.BranchID = appointment.BranchID
		tabAppointment.AppointmentCode = appointment.AppointmentCode
		tabAppointment.AppointmentPhone = appointment.AppointmentPhone
		tabAppointment.Name = appointment.Name
		tabAppointment.Gender = appointment.Gender
		tabAppointment.Way = appointment.Way
		tabAppointment.CustomerNum = appointment.CustomerNum
		tabAppointment.AppointmentAt = appointment.AppointmentAt
		tabAppointment.AppointmentDate = appointment.AppointmentDate
		tabAppointment.ExpireAt = appointment.ExpireAt
		tabAppointment.RoomTypeID = appointment.RoomTypeID
		tabAppointment.DepositFee = appointment.DepositFee
		tabAppointment.FlowerCake = appointment.FlowerCake
		tabAppointment.FlowerCakeRemark = appointment.FlowerCakeRemark
		tabAppointment.Remark = appointment.Remark
		tabAppointment.Operator = appointment.Operator
		tabAppointment.DateCounterID = appointment.DateCounterID
		tabAppointment.ChargingWay = appointment.ChargingWay
		tabAppointment.CalledCode = appointment.CalledCode
		tabAppointment.CalledPhone = appointment.CalledPhone
		tabAppointment.MemberID = appointment.MemberID

		// 更新预约信息
		if err := tx.Scopes(crius.ColumnEqualScope("id", tabAppointment.ID)).Save(&(tabAppointment.TableAppointment)).Error; err != nil {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("UpdateAppointment 更新预约数据库错误:%v", err))
			return oldMethod, method, err
		}
		oldMethod = model.UpdateMethod
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("UpdateAppointment 提交事务数据库错误:%v", err))
		return oldMethod, method, err
	}
	// 通知咨客
	if sended {
		go s.AppointmentHTTPRequest(&tabAppointment)
		go s.AppointmentHTTPRequest(&appointment)
	}
	return oldMethod, method, nil
}

func (s *Server) showTemplateConfig(merchantID, branchID, roomTypeID uuid.UUID, appointmentDate time.Time) (*model.TableAppointmentTemplateConfig, error) {
	config := new(model.TableAppointmentTemplateConfig)
	if err := s.database.Conn.Model(&model.TableAppointmentTemplateConfig{}).Joins("inner join appointment.appointment_template template on template.id=appointment_template_configure.template_id").
		Joins("inner join appointment.appointment_template_calendar calendar on template.id=calendar.template_id").
		Scopes(crius.ColumnEqualScope("calendar.merchant_id", merchantID), crius.ColumnEqualScope("calendar.branch_id", branchID), crius.ColumnEqualScope("calendar.business_date", appointmentDate), crius.ColumnEqualScope("appointment_template_configure.room_type_id", roomTypeID)).
		Take(config).Error; err != nil {
		return nil, err
	}
	return config, nil
}

// AppointmentHTTPRequest 预约创建更改通知咨客
func (s *Server) AppointmentHTTPRequest(appointment *model.Appointment) {
	defer crius.CatchException()
	client := http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{}).SignedString([]byte(config.Setting.App.AppointmentSalt))
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest 生成token错误:%v", err))
		return
	}

	branchResp, err := s.merchantBasic().ShowBranch(pkgs.MetadataContent(pkgs.Metadata{MerchantID: *appointment.MerchantID}), &merchantBasic.ShowBranchRequest{Id: appointment.BranchID.String()})
	if err != nil || branchResp == nil || branchResp.ErrorCode != pkgs.Success {
		crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest 获取门店错误:%v, %v", err, branchResp))
		return
	}
	if branchResp.Data.Domain == "" {
		crius.Logger.Error("AppointmentHttpRequest 门店域名为空")
		return
	}

	var method string
	switch appointment.Status {
	case model.AppointmentStatusAppointed, model.AppointmentStatusArrived:
		method = http.MethodPut
	case model.AppointmentStatusCancelled, model.AppointmentStatusRefunding, model.AppointmentStatusRefunded:
		method = http.MethodDelete
	}
	var body *bytes.Reader
	if method == http.MethodPut {
		phone, err := strconv.ParseInt(appointment.AppointmentPhone, 10, 64)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest 解析预约电话错误:%v", err))
		}
		req := appointmentRequest{
			PreorderID:    appointment.ID.String(),
			BranchID:      crius.UUIDToString(appointment.BranchID),
			BeginTime:     appointment.AppointmentAt.Format(ymdhms),
			Nickname:      appointment.Name,
			Phone:         phone,
			DepositFee:    appointment.DepositFee,
			DecorateFee:   appointment.PackagePrice,
			CustomerCount: int32(appointment.CustomerNum),
			Gender:        int32(appointment.Gender),
			Remarks:       appointment.Remark,
			UserID:        crius.UUIDToString(appointment.MemberID),
			StaffID:       crius.UUIDToString(appointment.Operator),
			Code:          appointment.Code,
			BusinessDate:  appointment.AppointmentDate.Format(ymd),
			Way:           fmt.Sprintf("0%d", appointment.Way),
			FromCrius:     1,
		}
		// 查询微信用户id
		wechatMemberResp, err := s.merchantBasic().ShowWechatUserByMember(context.Background(), &merchantBasic.ShowWechatUserByMemberRequest{MemberId: crius.UUIDToString(appointment.MemberID)})
		if err != nil || wechatMemberResp == nil || wechatMemberResp.ErrorCode != pkgs.Success {
			crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest 获取微信用户错误:%v", err))
		} else {
			req.CustomerID = wechatMemberResp.Data.Id
		}
		// 查询房型组id
		roomGroupResp, err := s.merchantBasic().ShowRoomType(context.Background(), &merchantBasic.ShowRoomTypeRequest{Id: crius.UUIDToString(appointment.RoomTypeID)})
		if err != nil || roomGroupResp == nil || roomGroupResp.ErrorCode != pkgs.Success {
			crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest 获取房型组错误:%v", err))
			return
		}
		for _, v := range roomGroupResp.Data.RoomTypeGroupIds {
			req.RoomTypeGroupIDs = append(req.RoomTypeGroupIDs, v)
		}
		// 查询房型名称
		roomTypeResp, err := s.merchantBasic().ShowRoomType(context.Background(), &merchantBasic.ShowRoomTypeRequest{Id: crius.UUIDToString(appointment.RoomTypeID)})
		if err != nil || roomTypeResp == nil || roomTypeResp.ErrorCode != pkgs.Success {
			crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest 获取房型名称错误:%v", err))
		} else {
			req.PlaceTypeName = roomTypeResp.Data.Name
		}
		// 主题
		if appointment.ThemeID != nil && appointment.PackageID != nil {
			theme, err := model.ShowAppointmentThemeWithCategory(*appointment.ThemeID)
			if err != nil {
				crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest 获取主题错误:%v", err))
				return
			}
			p, err := model.ShowAppointmentThemePackage(*appointment.PackageID)
			if err != nil {
				crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest 获取套餐错误:%v", err))
				return
			}
			req.Theme = &appointmentRequestTheme{
				Category:     theme.CategoryName,
				Name:         theme.Name,
				Package:      p.Name,
				ExtraPackage: p.Packages,
				Decoration:   p.Decoration,
				Staffing:     p.Staffing,
			}
			if p.CustomConfigs != nil {
				for _, customConfig := range *p.CustomConfigs {
					k, ok := customConfig["name"].(string)
					if !ok {
						continue
					}
					v, ok := customConfig["config"].(string)
					if !ok {
						continue
					}
					req.Theme.Customize = append(req.Theme.Customize, keyValue{Key: k, Value: v})
				}
			}

		}
		if appointment.ExpireAt == nil {
			req.ExpireTime = appointment.AppointmentAt.Add(15 * time.Minute).Format(ymdhms)
		} else {
			req.ExpireTime = appointment.ExpireAt.Format(ymdhms)
		}
		bs, _ := json.Marshal(req)
		body = bytes.NewReader(bs)
	} else if method == http.MethodDelete {
		req := appointmentRequest{PreorderID: appointment.ID.String(), FromCrius: 1}
		bs, _ := json.Marshal(req)
		body = bytes.NewReader(bs)
	} else {
		// method 不为 delete 和 put直接返回
		return
	}
	request, err := http.NewRequest(method, fmt.Sprintf(appointmentURL, branchResp.Data.Domain), body)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest http request错误:%v, %v", err, fmt.Sprintf(appointmentURL, branchResp.Data.Domain)))
		return
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", token)
	resp, err := client.Do(request)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest http请求错误:%v", err))
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		if err := model.UpdateAppointmentSended(appointment.ID, true); err != nil {
			crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest 更新预约发送状态错误:%v", err))
		}
	} else {
		crius.Logger.Error(fmt.Sprintf("AppointmentHttpRequest http请求失败:%v", resp))
	}
}

type appointmentRequest struct {
	PreorderID       string                   `json:"preorder_id"`
	BranchID         string                   `json:"branch_id"`
	BeginTime        string                   `json:"begin_time"`
	ExpireTime       string                   `json:"expire_time"`
	Nickname         string                   `json:"nickname"`
	Phone            int64                    `json:"phone"`
	RoomTypeGroupIDs []string                 `json:"room_type_group_ids"`
	DepositFee       int32                    `json:"deposit_fee"`
	DecorateFee      int32                    `json:"decorate_fee"`
	CustomerCount    int32                    `json:"customer_count"`
	CustomerID       string                   `json:"customer_id"`
	Gender           int32                    `json:"gender"`
	Remarks          string                   `json:"remarks"`
	UserID           string                   `json:"user_id"`
	StaffID          string                   `json:"staff_id"`
	Code             string                   `json:"code"`
	BusinessDate     string                   `json:"business_date"`
	Theme            *appointmentRequestTheme `json:"theme,omitempty"`
	Way              string                   `json:"way"`
	IsRegulars       int32                    `json:"is_regulars"`
	RegularsStaffID  string                   `json:"regulars_staff_id"`
	Other            appointmentRequestOther  `json:"other"`
	PlaceTypeName    string                   `json:"place_type_name"`
	FromCrius        int32                    `json:"from_crius"` // 来自微服务，固定传1
}

type appointmentRequestTheme struct {
	Category     string      `json:"category"`
	Name         string      `json:"name"`
	Package      string      `json:"package"`
	ExtraPackage interface{} `json:"extra_package"`
	Decoration   string      `json:"decoration"`
	Staffing     string      `json:"staffing"`
	Customize    []keyValue  `json:"customize"`
}

type keyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type appointmentRequestOther struct {
	Birthday   int32 `json:"birthday"`
	Party      int32 `json:"party"`
	NotDisturb int32 `json:"not_disturb"`
}

type roomTypeDepositFee struct {
	roomTypeID uuid.UUID
	depositFee int32
}

type templateBusinessDate struct {
	templateID   uuid.UUID
	businessDate time.Time
}

const (
	ymdhms = "2006-01-02 15:04:05"
	ymd    = "2006-01-02"
)
