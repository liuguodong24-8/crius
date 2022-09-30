package rpc

import (
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/message/internal/config"
	"gitlab.omytech.com.cn/micro-service/message/internal/model"
	"gitlab.omytech.com.cn/micro-service/message/internal/sms"
	"gitlab.omytech.com.cn/micro-service/message/proto"
	"gorm.io/gorm"
)

// SendSms 发送短信
func (s *Server) SendSms(ctx context.Context, req *proto.SendSmsRequest) (*proto.SendSmsResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("发送短信[begin]")

	// 是否本地调试
	if config.Setting.Sms.Debug {
		s.saveSmsRecord(ctx, req, &sms.SendResponse{
			Code: 0,
			Msg:  fmt.Sprintf("调试发送成功:%s", req.Message),
		})

		return &proto.SendSmsResponse{
			ErrorCode:    pkgs.Success,
			ErrorMessage: "调试发送成功",
		}, nil
	}

	entity, err := sms.NewSms()
	if err != nil {
		return &proto.SendSmsResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("发送短信失败:%s", err.Error()),
		}, nil
	}

	res, smsErr := entity.Send(ctx, sms.SendRequest{
		Sign:     req.Sign,
		AreaCode: req.AreaCode,
		Mobile:   req.Phone,
		Message:  req.Message,
	})
	if smsErr != nil {
		return &proto.SendSmsResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("发送短信失败:%s", smsErr.Error()),
		}, nil
	}

	go s.saveSmsRecord(ctx, req, res)

	if res.Code == 0 {
		return &proto.SendSmsResponse{
			ErrorCode:    pkgs.Success,
			ErrorMessage: "成功",
		}, nil
	}

	return &proto.SendSmsResponse{
		ErrorCode:    pkgs.ErrInternal,
		ErrorMessage: fmt.Sprintf("发送短信失败：%s", res.Msg),
	}, nil
}

func (s *Server) saveSmsRecord(ctx context.Context, req *proto.SendSmsRequest, res *sms.SendResponse) {
	p := pkgs.MakeParams(res)
	metadata := pkgs.GetMetadata(ctx)
	branchID := uuid.FromStringOrNil(req.BranchId)
	// 保存记录
	record := model.SmsStat{
		ID:          uuid.NewV4(),
		BranchID:    &branchID,
		MerchantID:  &metadata.MerchantID,
		MessageType: req.MessageType,
		AreaCode:    req.AreaCode,
		Phone:       req.Phone,
		Sign:        req.Sign,
		System:      req.System,
		Status:      model.SmsStatusSuccess,
		Content:     req.Message,
		Extra:       &p,
	}
	if res.Code != 0 {
		record.Status = model.SmsStatusFail
	}

	if err := s.database.Conn.Create(&record).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).WithFields("sms stat", logger.MakeFields(record)).WithError(err).Error("保存短信发送记录失败")
	}
}

// CreateSmsTemplate 创建短信模版
func (s *Server) CreateSmsTemplate(ctx context.Context, req *proto.CreateSmsTemplateRequest) (*proto.CreateSmsTemplateResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreateSmsTemplate")
	metadata := pkgs.GetMetadata(ctx)
	if len(metadata.MerchantID.String()) == 0 || len(req.Name) == 0 {
		util.Logger.WithMetadata(ctx).Error("新增短信模版参数错误")
		return &proto.CreateSmsTemplateResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	template := model.SmsTemplate{
		ID:          uuid.NewV4(),
		MerchantID:  metadata.MerchantID,
		Name:        req.Name,
		Sign:        req.Sign,
		Category:    req.Category,
		CategoryKey: req.CategoryKey,
		Content:     req.Content,
		Status:      util.StringToStatus(req.Status),
	}

	var count int64
	if err := s.database.Conn.Model(&model.SmsTemplate{}).Scopes(util.ColumnEqualScope("merchant_id", template.MerchantID), util.ColumnEqualScope("name", template.Name)).Count(&count).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("新增短信模版，数据库过滤名字唯一查询错误")
		return &proto.CreateSmsTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, nil
	}

	if count > 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("新增短信模版，名字重复")
		return &proto.CreateSmsTemplateResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: fmt.Sprintf("新增短信模版，名字重复:%s", req.Name),
		}, nil
	}

	if err := s.database.Conn.Create(&template).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("template", logger.MakeFields(template)).WithError(err).Error("新增短信模版，数据库保存错误")
		return &proto.CreateSmsTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库保存错误:%s", err.Error()),
		}, nil
	}

	after := pkgs.MakeParams(template)
	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: template.TableName(),
		TableID:           template.ID,
		Method:            model.CreateMethod,
		After:             &after,
	})

	return &proto.CreateSmsTemplateResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// UpdateSmsTemplate 修改短信模版
func (s *Server) UpdateSmsTemplate(ctx context.Context, req *proto.UpdateSmsTemplateRequest) (*proto.UpdateSmsTemplateResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateSmsTemplate")
	if len(req.Id) == 0 {
		util.Logger.WithMetadata(ctx).Error("修改短信模版,参数错误")
		return &proto.UpdateSmsTemplateResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "修改模版参数错误",
		}, nil
	}

	// 查询
	var template model.SmsTemplate
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&template).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.UpdateSmsTemplateResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}
		return &proto.UpdateSmsTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, nil
	}

	// 名字判断
	var count int64
	if err := s.database.Conn.Model(&model.SmsTemplate{}).Scopes(
		util.ColumnEqualScope("merchant_id", template.MerchantID),
		util.ColumnEqualScope("name", req.Name),
		util.ColumnNotEqualScope("id", template.ID)).Count(&count).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改短信模版,数据库过滤名字唯一错误")
		return &proto.UpdateSmsTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, nil
	}

	if count == 0 {
		return &proto.UpdateSmsTemplateResponse{
			ErrorCode:    pkgs.Success,
			ErrorMessage: "",
		}, nil
	}

	update := map[string]interface{}{
		"sign":    req.Sign,
		"name":    req.Name,
		"content": req.Content,
	}
	if len(req.Status) > 0 {
		update["status"] = util.StringToStatus(req.Status)
	}

	if err := s.database.Conn.Model(&model.SmsTemplate{}).Where("id = ?", template.ID).Updates(update).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("update", logger.MakeFields(update)).WithError(err).Error("修改短信模版失败")
		return &proto.UpdateSmsTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}

	before := pkgs.MakeParams(template)
	after := pkgs.MakeParams(update)

	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: template.TableName(),
		TableID:           template.ID,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
	})

	return &proto.UpdateSmsTemplateResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// ChangeSmsTemplateStatus 更改短信模板状态
func (s *Server) ChangeSmsTemplateStatus(ctx context.Context, req *proto.ChangeSmsTemplateStatusRequest) (*proto.ChangeSmsTemplateStatusResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateSmsTemplate")
	if len(req.Id) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改短信模板状态,参数错误")
		return &proto.ChangeSmsTemplateStatusResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "修改短信模板状态参数错误",
		}, nil
	}

	// 查询
	var SmsTemplate model.SmsTemplate
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&SmsTemplate).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.ChangeSmsTemplateStatusResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}
		return &proto.ChangeSmsTemplateStatusResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, nil
	}

	update := map[string]interface{}{
		`status`: util.StringToStatus(req.Status),
	}
	if err := s.database.Conn.Model(&model.SmsTemplate{}).Where("id = ?", SmsTemplate.ID).Updates(update).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(update)).WithError(err).Error("修改短信模板状态失败")
		return &proto.ChangeSmsTemplateStatusResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}

	before := pkgs.MakeParams(SmsTemplate)
	after := pkgs.MakeParams(update)

	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: SmsTemplate.TableName(),
		TableID:           SmsTemplate.ID,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
	})

	return &proto.ChangeSmsTemplateStatusResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// ListSmsTemplate 短信模版列表
func (s *Server) ListSmsTemplate(ctx context.Context, req *proto.ListSmsTemplateRequest) (*proto.ListSmsTemplateResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ListSmsTemplate")
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))
	if len(req.Name) > 0 {
		scopes = append(scopes, util.ColumnLikeScope("name", req.Name))
	}
	if len(req.Status) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("status", req.Status))
	}
	if len(req.Category) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("category", req.Category))
	}
	if len(req.CategoryKey) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("category_key", req.CategoryKey))
	}
	var total int64
	s.database.Conn.Model(&model.SmsTemplate{}).Scopes(scopes...).Count(&total)
	if total == 0 {
		return &proto.ListSmsTemplateResponse{
			ErrorCode:    pkgs.Success,
			ErrorMessage: "",
		}, nil
	}

	var templates []model.SmsTemplate
	if req.WithPage {
		scopes = append(scopes, util.PaginationScope(req.Offset, req.Limit))
	}
	orderBy := "updated_at desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}
	if err := s.database.Conn.Model(&model.SmsTemplate{}).Scopes(scopes...).Order(orderBy).Find(&templates).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询短信模版列表错误")
		return &proto.ListSmsTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据查询错误:%s", err.Error()),
		}, nil
	}

	data := make([]*proto.SmsTemplate, 0)
	for _, v := range templates {
		data = append(data, toProtoSmsTemplate(v))
	}

	return &proto.ListSmsTemplateResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ListSmsTemplateData{
			Templates: data,
			Total:     total,
		},
	}, nil
}

// ShowSmsTemplate 短信模版详情
func (s *Server) ShowSmsTemplate(ctx context.Context, req *proto.ShowSmsTemplateRequest) (*proto.ShowSmsTemplateResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowSmsTemplate")
	if len(req.Id) == 0 {
		util.Logger.WithMetadata(ctx).Error("查看短信模版详情，参数错误")
		return &proto.ShowSmsTemplateResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	var template model.SmsTemplate
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&template).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.ShowSmsTemplateResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.ShowSmsTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, nil
	}

	return &proto.ShowSmsTemplateResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         toProtoSmsTemplate(template),
	}, nil
}

func toProtoSmsTemplate(item model.SmsTemplate) *proto.SmsTemplate {
	return &proto.SmsTemplate{
		Id:          item.ID.String(),
		Name:        item.Name,
		Sign:        item.Sign,
		Category:    item.Category,
		CategoryKey: item.CategoryKey,
		Content:     item.Content,
		Status:      item.Status.String(),
		CreatedAt:   item.CreatedAt.Time.Unix(),
	}
}
