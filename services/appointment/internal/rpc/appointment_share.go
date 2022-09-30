package rpc

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gorm.io/gorm"
)

// JoinSharedAppointment 加入分享预约订单
func (s *Server) JoinSharedAppointment(ctx context.Context, req *proto.JoinSharedAppointmentRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("JoinSharedAppointment")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	wechatID := uuid.FromStringOrNil(req.WechatId)
	appointment, err := model.ShowAppointment(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约不存在"
			return resp, nil
		}
		util.Logger.Error(fmt.Sprintf("JoinSharedAppointment 查询预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "加入分享预约失败"
		return resp, nil
	}
	if appointment.Status != model.AppointmentStatusAppointed {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "预约订单状态错误"
		return resp, nil
	}

	if appointment.WechatIDs == nil {
		appointment.WechatIDs = new(model.WechatIDs)
	} else {
		for _, v := range *appointment.WechatIDs {
			if v.ID == wechatID {
				resp.ErrorCode = pkgs.ErrUnprocessableEntity
				resp.ErrorMessage = "用户已加入"
				return resp, nil
			}
		}
	}
	*appointment.WechatIDs = append(*appointment.WechatIDs, model.WechatID{ID: wechatID, Time: time.Now()})
	if err := model.UpdateAppointmentWechatIDs(appointment); err != nil {
		util.Logger.Error(fmt.Sprintf("JoinSharedAppointment 查询预约数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "加入分享预约失败"
		return resp, nil
	}
	return resp, nil
}

// UpdateAppointmentShare 更新预约分享设置
func (s *Server) UpdateAppointmentShare(ctx context.Context, req *proto.UpdateAppointmentShareRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentShare")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	if err := model.UpdateAppointmentShare(id, req.Message); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentShare 更新预约分享设置数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约分享设置失败"
		return resp, nil
	}
	return resp, nil
}
