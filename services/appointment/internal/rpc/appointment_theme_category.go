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

// CreateAppointmentThemeCategory 创建主题预约分类
func (s *Server) CreateAppointmentThemeCategory(ctx context.Context, req *proto.CreateAppointmentThemeCategoryRequest) (*proto.Response, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreateAppointmentThemeCategory")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	category, err := model.ShowAppointmentThemeCategoryByName(req.Category.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		util.Logger.Error(fmt.Sprintf("CreateAppointmentThemeCategory 查询主题预约分类数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建主题预约分类失败"
		return resp, nil
	}
	if err == nil {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "主题预约分类名称已存在"
		return resp, nil
	}
	category = &model.TableAppointmentThemeCategory{
		ID:     uuid.NewV4(),
		Name:   req.Category.Name,
		Weight: req.Category.Weight,
		Status: util.Status(req.Category.Status),
	}
	if err := model.CreateAppointmentThemeCategory(category); err != nil {
		util.Logger.Error(fmt.Sprintf("CreateAppointmentThemeCategory 创建主题预约分类数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建主题预约分类失败"
		return resp, nil
	}
	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_ThemeCategory,
	}
	go s.PublishDataChangeEvent(message)
	return resp, nil
}

// UpdateAppointmentThemeCategory 更新主题预约分类
func (s *Server) UpdateAppointmentThemeCategory(ctx context.Context, req *proto.UpdateAppointmentThemeCategoryRequest) (*proto.Response, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentThemeCategory")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Category.Id)

	category, err := model.ShowAppointmentThemeCategoryByName(req.Category.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentThemeCategory 查询主题预约分类数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新主题预约分类失败"
		return resp, nil
	}
	if err == nil && category.ID != id {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "主题预约分类名称已存在"
		return resp, nil
	}
	category = &model.TableAppointmentThemeCategory{
		ID:     id,
		Name:   req.Category.Name,
		Weight: req.Category.Weight,
		Status: util.Status(req.Category.Status),
	}
	if err := model.UpdateAppointmentThemeCategory(category); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentThemeCategory 更新主题预约分类数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新主题预约分类失败"
		return resp, nil
	}
	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_ThemeCategory,
	}
	go s.PublishDataChangeEvent(message)
	return resp, nil
}

// UpdateAppointmentThemeCategoryStatus 更新主题预约分类状态
func (s *Server) UpdateAppointmentThemeCategoryStatus(ctx context.Context, req *proto.UpdateAppointmentThemeCategoryStatusRequest) (*proto.Response, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentThemeCategoryStatus")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	if err := model.UpdateAppointmentThemeCategoryStatus(id, util.Status(req.Status)); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentThemeCategoryStatus 更新主题预约分类状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新主题预约分类状态失败"
		return resp, nil
	}
	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_ThemeCategory,
	}
	go s.PublishDataChangeEvent(message)
	return resp, nil
}

// GetAppointmentThemeCategories 获取主题预约分类列表
func (s *Server) GetAppointmentThemeCategories(ctx context.Context, req *proto.GetAppointmentThemeCategoriesRequest) (*proto.GetAppointmentThemeCategoriesResponse, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentThemeCategories")
	resp := &proto.GetAppointmentThemeCategoriesResponse{
		ErrorCode: pkgs.Success,
	}
	categories, total, err := model.GetAppointmentThemeCategories(req.Name, util.Status(req.Status), req.Offset, req.Limit)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetAppointmentThemeCategories 获取主题预约分类列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取主题预约分类列表失败"
		return resp, nil
	}

	resp.Data = &proto.GetAppointmentThemeCategoriesResponse_Data{Total: int32(total)}
	for i := range categories {
		resp.Data.Categories = append(resp.Data.Categories, toProtoAppointmentThemeCategory(&categories[i]))
	}
	return resp, nil
}

// ShowAppointmentThemeCategory 获取主题预约分类
func (s *Server) ShowAppointmentThemeCategory(ctx context.Context, req *proto.ShowAppointmentThemeCategoryRequest) (*proto.ShowAppointmentThemeCategoryResponse, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowAppointmentThemeCategory")
	resp := &proto.ShowAppointmentThemeCategoryResponse{
		ErrorCode: pkgs.Success,
	}
	category, err := model.ShowAppointmentThemeCategory(uuid.FromStringOrNil(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "主题预约分类不存在"
			return resp, nil
		}
		util.Logger.Error(fmt.Sprintf("ShowAppointmentThemeCategory 获取主题预约分类数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取主题预约分类失败"
		return resp, nil
	}

	resp.Data = toProtoAppointmentThemeCategory(category)
	return resp, nil
}

func toProtoAppointmentThemeCategory(c *model.TableAppointmentThemeCategory) *proto.AppointmentThemeCategory {
	return &proto.AppointmentThemeCategory{
		Id:     c.ID.String(),
		Name:   c.Name,
		Weight: c.Weight,
		Status: c.Status.String(),
	}
}
