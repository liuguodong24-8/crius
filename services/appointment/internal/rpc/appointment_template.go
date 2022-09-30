package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gorm.io/gorm"
)

// SaveAppointmentTemplate 保存预约模板
func (s *Server) SaveAppointmentTemplate(ctx context.Context, req *proto.SaveAppointmentTemplateRequest) (*proto.SaveAppointmentTemplateResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("SaveAppointmentTemplate")
	resp := &proto.SaveAppointmentTemplateResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	branchID := uuid.FromStringOrNil(req.Template.BranchId)
	id := uuid.FromStringOrNil(req.Template.Id)
	roomTypeIDs, err := fields.StringArrToUUIDArr(req.Template.RoomTypeIds)
	beginTime := fields.StringToLocalTime(req.Template.BeginTime)
	endTime := fields.StringToLocalTime(req.Template.EndTime)
	if err != nil {
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	tabTemplate := new(model.TableAppointmentTemplate)
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("name", req.Template.Name)).Take(tabTemplate).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("SaveAppointmentTemplate 查询模板名字数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "保存预约模板失败"
			return resp, nil
		}
	} else {
		if tabTemplate.ID != id {
			resp.ErrorMessage = "模板名称已存在"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	template := model.TableAppointmentTemplate{
		MerchantID:  &merchantID,
		BranchID:    &branchID,
		Name:        req.Template.Name,
		Color:       req.Template.Color,
		Status:      req.Template.Status,
		RoomTypeIDs: &roomTypeIDs,
		BeginTime:   &beginTime,
		EndTime:     &endTime,
	}
	if req.Template.IsNextDay {
		template.IsNextDay = 1
	}

	tx := s.database.Conn.Begin()
	if id != uuid.Nil {
		err = tx.Scopes(crius.ColumnEqualScope("id", id)).Updates(&template).Error
	} else {
		id = uuid.NewV4()
		template.ID = id
		err = tx.Create(&template).Error
	}
	if err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentTemplate 创建、更新预约模板数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存预约模板失败"
		return resp, nil
	}

	configs := make([]model.TableAppointmentTemplateConfig, 0)
	for _, v := range req.Config {
		roomTypeID := uuid.FromStringOrNil(v.RoomTypeId)
		templateConfig := model.TableAppointmentTemplateConfig{
			ID:         uuid.NewV4(),
			RoomTypeID: &roomTypeID,
			TemplateID: &id,
			AdvanceDay: int16(v.AdvanceDay),
			DepositFee: v.DepositFee,
		}
		mapArr := make(pkgs.ParamsArr, 0)
		for _, config := range v.Configure {
			m := make(map[string]interface{})
			m["way"] = config.Way
			m["time"] = config.Time
			m["num"] = config.Num
			m["is_next_day"] = config.IsNextDay
			mapArr = append(mapArr, m)
		}
		templateConfig.Configure = &mapArr
		configs = append(configs, templateConfig)
	}

	if err := tx.Scopes(crius.ColumnEqualScope("template_id", id)).
		Delete(&model.TableAppointmentTemplateConfig{}).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentTemplate 删除旧模板配置数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存预约模板配置失败"
		return resp, nil
	}
	if err := tx.Create(&configs).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentTemplate 新增模板配置数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存预约模板配置失败"
		return resp, nil
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentTemplate 事务提交数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存预约模板配置失败"
		return resp, nil
	}

	resp.Data = id.String()
	return resp, nil
}

// GetAppointmentTemplateConfigs 获取预约模板配置
func (s *Server) GetAppointmentTemplateConfigs(ctx context.Context, req *proto.GetAppointmentTemplateConfigsRequest) (*proto.GetAppointmentTemplateConfigsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentTemplateConfigs")
	resp := &proto.GetAppointmentTemplateConfigsResponse{
		ErrorCode: pkgs.Success,
	}

	templateID := uuid.FromStringOrNil(req.TemplateId)

	var configs []model.TableAppointmentTemplateConfig
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("template_id", templateID)).Find(&configs).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentTemplateConfigs 获取预约模板配置数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取预约模板配置失败"
		return resp, nil
	}

	for _, v := range configs {
		resp.Data = append(resp.Data, toProtoAppointmentTemplateConfig(v))
	}

	return resp, nil
}

// ShowAppointmentTemplate 获取预约模板
func (s *Server) ShowAppointmentTemplate(ctx context.Context, req *proto.ShowAppointmentTemplateRequest) (*proto.ShowAppointmentTemplateResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowAppointmentTemplate")
	resp := &proto.ShowAppointmentTemplateResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	template := new(model.TableAppointmentTemplate)
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("id", id)).Take(template).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约模板不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowAppointmentTemplate 获取预约模板数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取预约模板失败"
		return resp, nil
	}

	resp.Data = toProtoAppointmentTemplate(*template)
	return resp, nil
}

// GetAppointmentTemplates 获取模板列表
func (s *Server) GetAppointmentTemplates(ctx context.Context, req *proto.GetAppointmentTemplatesRequest) (*proto.GetAppointmentTemplatesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentTemplates")
	resp := &proto.GetAppointmentTemplatesResponse{
		ErrorCode: pkgs.Success,
	}

	var count int64
	branchID := uuid.FromStringOrNil(req.BranchId)

	templates := make([]model.TableAppointmentTemplate, 0)
	db := s.database.Conn.Scopes(crius.ColumnEqualScopeDefault("branch_id", branchID), crius.ColumnEqualScopeDefault("name", req.Name), crius.ColumnEqualScopeDefault("status", req.Status))
	if err := db.Model(&model.TableAppointmentTemplate{}).Count(&count).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentTemplates 获取预约模板列表数量数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取预约模板列表失败"
		return resp, nil
	}
	if count == 0 {
		resp.Data = &proto.GetAppointmentTemplatesData{Total: 0}
		return resp, nil
	}
	if err := db.Scopes(model.PagingCondition(req.Offset, req.Limit)).Find(&templates).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetAppointmentTemplates 获取预约模板列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取预约模板列表失败"
		return resp, nil
	}

	resp.Data = &proto.GetAppointmentTemplatesData{
		Total: int32(count),
	}
	for _, v := range templates {
		resp.Data.Templates = append(resp.Data.Templates, toProtoAppointmentTemplate(v))
	}

	return resp, nil
}

// UpdateAppointmentTemplateStatus 获取模板列表
func (s *Server) UpdateAppointmentTemplateStatus(ctx context.Context, req *proto.UpdateAppointmentTemplateStatusRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentTemplateStatus")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	if err := s.database.Conn.Model(&model.TableAppointmentTemplate{}).Scopes(crius.ColumnEqualScope("id", id)).Update("status", req.Status).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateAppointmentTemplateStatus 更新预约模板状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新预约模板状态失败"
		return resp, nil
	}
	return resp, nil
}

func toProtoAppointmentTemplateConfig(config model.TableAppointmentTemplateConfig) *proto.AppointmentTemplateConfig {
	var configure []*proto.TemplateRoomConfigColumn
	err := json.Unmarshal([]byte(config.Configure.JSON()), &configure)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("toProtoAppointmentTemplateConfig 类型转换room config错误:%v", err))
	}
	return &proto.AppointmentTemplateConfig{
		Id:         config.ID.String(),
		RoomTypeId: config.RoomTypeID.String(),
		TemplateId: config.TemplateID.String(),
		AdvanceDay: int32(config.AdvanceDay),
		DepositFee: config.DepositFee,
		Configure:  configure,
	}
}

func toProtoAppointmentTemplate(template model.TableAppointmentTemplate) *proto.AppointmentTemplate {
	isNextDay := false
	var beginTime, endTime string
	if template.IsNextDay == 1 {
		isNextDay = true
	}
	if template.BeginTime != nil {
		beginTime = template.BeginTime.String()
	}
	if template.EndTime != nil {
		endTime = template.EndTime.String()
	}
	return &proto.AppointmentTemplate{
		Id:          template.ID.String(),
		BranchId:    template.BranchID.String(),
		Name:        template.Name,
		Color:       template.Color,
		Status:      template.Status,
		RoomTypeIds: template.RoomTypeIDs.ToStringArr(),
		CreatedAt:   int32(template.CreatedAt.Unix()),
		UpdatedAt:   int32(template.UpdatedAt.Unix()),
		BeginTime:   beginTime,
		EndTime:     endTime,
		IsNextDay:   isNextDay,
	}
}
