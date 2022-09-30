package rpc

import (
	"context"
	"fmt"
	mqMessage "gitlab.omytech.com.cn/micro-service/Crius/pkgs/message"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gorm.io/gorm"
)

// CreateAppointmentThemeFeature 创建主题预约特色
func (s *Server) CreateAppointmentThemeFeature(ctx context.Context, req *proto.CreateAppointmentThemeFeatureRequest) (*proto.Response, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreateAppointmentThemeFeature")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	feature, err := model.ShowAppointmentThemeFeatureByName(req.Feature.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		util.Logger.Error(fmt.Sprintf("CreateAppointmentThemeFeature 查询主题预约特色数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建主题预约特色失败"
		return resp, nil
	}
	if err == nil {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "主题预约特色名称已存在"
		return resp, nil
	}
	feature = &model.TableAppointmentThemeFeature{
		ID:     uuid.NewV4(),
		Name:   req.Feature.Name,
		Weight: req.Feature.Weight,
		Status: util.Status(req.Feature.Status),
		Icon:   req.Feature.Icon,
	}
	if err := model.CreateAppointmentThemeFeature(feature); err != nil {
		util.Logger.Error(fmt.Sprintf("CreateAppointmentThemeFeature 创建主题预约特色数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建主题预约特色失败"
		return resp, nil
	}
	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_ThemeFeature,
	}
	go s.PublishDataChangeEvent(message)
	return resp, nil
}

// UpdateAppointmentThemeFeature 更新主题预约特色
func (s *Server) UpdateAppointmentThemeFeature(ctx context.Context, req *proto.UpdateAppointmentThemeFeatureRequest) (*proto.Response, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentThemeFeature")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Feature.Id)

	feature, err := model.ShowAppointmentThemeFeatureByName(req.Feature.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentThemeFeature 查询主题预约特色数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新主题预约特色失败"
		return resp, nil
	}
	if err == nil && feature.ID != id {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "主题预约特色名称已存在"
		return resp, nil
	}
	feature = &model.TableAppointmentThemeFeature{
		ID:     id,
		Name:   req.Feature.Name,
		Weight: req.Feature.Weight,
		Status: util.Status(req.Feature.Status),
		Icon:   req.Feature.Icon,
	}
	if err := model.UpdateAppointmentThemeFeature(feature); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentThemeFeature 更新主题预约特色数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新主题预约特色失败"
		return resp, nil
	}
	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_ThemeFeature,
	}
	go s.PublishDataChangeEvent(message)
	return resp, nil
}

// UpdateAppointmentThemeFeatureStatus 更新主题预约特色状态
func (s *Server) UpdateAppointmentThemeFeatureStatus(ctx context.Context, req *proto.UpdateAppointmentThemeFeatureStatusRequest) (*proto.Response, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentThemeFeatureStatus")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	if err := model.UpdateAppointmentThemeFeatureStatus(id, util.Status(req.Status)); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentThemeFeatureStatus 更新主题预约特色状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新主题预约特色状态失败"
		return resp, nil
	}
	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_ThemeFeature,
	}
	go s.PublishDataChangeEvent(message)
	return resp, nil
}

// GetAppointmentThemeFeatures 获取主题预约特色列表
func (s *Server) GetAppointmentThemeFeatures(ctx context.Context, req *proto.GetAppointmentThemeFeaturesRequest) (*proto.GetAppointmentThemeFeaturesResponse, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentThemeFeatures")
	resp := &proto.GetAppointmentThemeFeaturesResponse{
		ErrorCode: pkgs.Success,
	}
	features, total, err := model.GetAppointmentThemeFeatures(req.Name, util.Status(req.Status), req.Offset, req.Limit)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetAppointmentThemeFeatures 获取主题预约特色列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取主题预约特色列表失败"
		return resp, nil
	}

	resp.Data = &proto.GetAppointmentThemeFeaturesResponse_Data{Total: int32(total)}
	for i := range features {
		resp.Data.Features = append(resp.Data.Features, toProtoAppointmentThemeFeature(&features[i]))
	}
	return resp, nil
}

// ShowAppointmentThemeFeature 获取主题预约特色
func (s *Server) ShowAppointmentThemeFeature(ctx context.Context, req *proto.ShowAppointmentThemeFeatureRequest) (*proto.ShowAppointmentThemeFeatureResponse, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowAppointmentThemeFeature")
	resp := &proto.ShowAppointmentThemeFeatureResponse{
		ErrorCode: pkgs.Success,
	}
	feature, err := model.ShowAppointmentThemeFeature(uuid.FromStringOrNil(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "主题预约特色不存在"
			return resp, nil
		}
		util.Logger.Error(fmt.Sprintf("ShowAppointmentThemeFeature 获取主题预约特色数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取主题预约特色失败"
		return resp, nil
	}

	resp.Data = toProtoAppointmentThemeFeature(feature)
	return resp, nil
}

func toProtoAppointmentThemeFeature(c *model.TableAppointmentThemeFeature) *proto.AppointmentThemeFeature {
	return &proto.AppointmentThemeFeature{
		Id:     c.ID.String(),
		Name:   c.Name,
		Weight: c.Weight,
		Icon:   c.Icon,
		Status: c.Status.String(),
	}
}
