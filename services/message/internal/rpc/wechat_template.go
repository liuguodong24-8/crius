package rpc

import (
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/message/internal/cache"
	"gitlab.omytech.com.cn/micro-service/message/internal/model"
	"gitlab.omytech.com.cn/micro-service/message/internal/wechat"
	"gitlab.omytech.com.cn/micro-service/message/proto"
	"gorm.io/gorm"
)

// CreateWechatTemplate 创建微信模版
func (s *Server) CreateWechatTemplate(ctx context.Context, req *proto.CreateWechatTemplateRequest) (*proto.CreateWechatTemplateResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("CreateWechatTemplate")
	metadata := pkgs.GetMetadata(ctx)
	if len(metadata.MerchantID.String()) == 0 {
		util.Logger.WithMetadata(ctx).Error("新增微信模版参数错误")
		return &proto.CreateWechatTemplateResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	template := model.WechatTemplate{
		ID:              uuid.NewV4(),
		MerchantID:      metadata.MerchantID,
		TemplateName:    req.TemplateName,
		TemplateCode:    req.TemplateCode,
		OfficialLink:    req.OfficialLink,
		MiniprogramLink: req.MiniprogramLink,
		Category:        req.Category,
		CategoryKey:     req.CategoryKey,
		Status:          util.StatusOpened,
	}

	template.TemplateContent = toWechatContent(req.Content)
	after := pkgs.MakeParams(template)

	if err := s.database.Conn.Create(&template).Error; nil != err {
		util.Logger.WithFields("template", logger.MakeFields(template)).WithError(err).Error("新增微信模版，数据库保存错位")
		return &proto.CreateWechatTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库保存错误:%s", err.Error()),
		}, nil
	}

	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: template.TableName(),
		TableID:           template.ID,
		Method:            model.CreateMethod,
		After:             &after,
	})

	return &proto.CreateWechatTemplateResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// ListWechatTemplate 微信模版列表
func (s *Server) ListWechatTemplate(ctx context.Context, req *proto.ListWechatTemplateRequest) (*proto.ListWechatTemplateResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("ListWechatTemplate")
	metadata := pkgs.GetMetadata(ctx)
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))
	if len(req.TemplateName) > 0 {
		scopes = append(scopes, util.ColumnLikeScope("template_name", req.TemplateName))
	}
	if len(req.Status) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("status", req.Status))
	}

	var total int64
	var templates []model.WechatTemplate

	if req.WithPage {
		scopes = append(scopes, util.PaginationScope(req.Offset, req.Limit))
	}

	s.database.Conn.Model(&model.WechatTemplate{}).Scopes(scopes...).Count(&total)
	if total == 0 {
		return &proto.ListWechatTemplateResponse{
			ErrorCode:    pkgs.Success,
			ErrorMessage: "",
		}, nil
	}

	orderBy := "updated_at desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}

	if err := s.database.Conn.Model(&model.WechatTemplate{}).Scopes(scopes...).Order(orderBy).Find(&templates).Error; nil != err {
		util.Logger.WithError(err).Error("查询微信模版列表错误")
		return &proto.ListWechatTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据查询错误:%s", err.Error()),
		}, nil
	}

	data := make([]*proto.WechatTemplate, 0)

	for _, v := range templates {
		data = append(data, s.toProtoWechatTemplate(ctx, v))
	}

	return &proto.ListWechatTemplateResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ListWechatTemplateData{
			Templates: data,
			Total:     total,
		},
	}, nil
}

// UpdateWechatTemplateStatus 修改微信模版状态
func (s *Server) UpdateWechatTemplateStatus(ctx context.Context, req *proto.UpdateWechatTemplateStatusRequest) (*proto.UpdateWechatTemplateStatusResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("UpdateWechatTemplateStatus")
	metadata := pkgs.GetMetadata(ctx)
	if len(req.Id) == 0 {
		util.Logger.WithMetadata(ctx).Error("修改微信模版状态，参数错误")
		return &proto.UpdateWechatTemplateStatusResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "修改微信模版状态，参数错误",
		}, nil
	}

	var template model.WechatTemplate
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&template).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.UpdateWechatTemplateStatusResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.UpdateWechatTemplateStatusResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误：%s", err.Error()),
		}, nil
	}

	update := map[string]interface{}{
		`status`: util.StringToStatus(req.Status),
	}

	if err := s.database.Conn.Model(&model.WechatTemplate{}).Where("id = ?", template.ID).Updates(update).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(update)).WithError(err).Error("修改微信模板状态失败")
		return &proto.UpdateWechatTemplateStatusResponse{
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

	return &proto.UpdateWechatTemplateStatusResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// ShowWechatTemplate 查看单个微信模版详情
func (s *Server) ShowWechatTemplate(ctx context.Context, req *proto.ShowWechatTemplateRequest) (*proto.ShowWechatTemplateResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("ShowWechatTemplate")
	if len(req.Id) == 0 {
		util.Logger.WithMetadata(ctx).Error("查看微信模版详情，参数错误")
		return &proto.ShowWechatTemplateResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	var template model.WechatTemplate
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&template).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.ShowWechatTemplateResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.ShowWechatTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误：%s", err.Error()),
		}, nil
	}

	return &proto.ShowWechatTemplateResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         s.toProtoWechatTemplate(ctx, template),
	}, nil
}

// UpdateWechatTemplate 修改微信模版消息
func (s *Server) UpdateWechatTemplate(ctx context.Context, req *proto.UpdateWechatTemplateRequest) (*proto.UpdateWechatTemplateResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("UpdateWechatTemplate")
	metadata := pkgs.GetMetadata(ctx)
	if len(req.Id) == 0 {
		util.Logger.WithMetadata(ctx).Error("修改微信模版,参数错误")
		return &proto.UpdateWechatTemplateResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "修改微信模版参数错误",
		}, nil
	}

	var template model.WechatTemplate
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&template).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.UpdateWechatTemplateResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.UpdateWechatTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误：%s", err.Error()),
		}, nil
	}

	content := &model.WechatTemplateContent{}
	if nil != req.Content {
		if req.Content.First != nil {
			content.First = model.WechatTemplateContentBase{
				Value: req.Content.First.Value,
				Color: req.Content.First.Color,
			}
		}

		if req.Content.Remark != nil {
			content.Remark = model.WechatTemplateContentBase{
				Value: req.Content.Remark.Value,
				Color: req.Content.Remark.Color,
			}
		}

		var detail []model.WechatTemplateContentDetail
		for _, v := range req.Content.Detail {
			detail = append(detail, model.WechatTemplateContentDetail{
				Name:  v.Name,
				Value: v.Value,
				Color: v.Color,
			})
		}

		content.Detail = detail
	}

	update := map[string]interface{}{
		"template_name":    req.TemplateName,
		"template_code":    req.TemplateCode,
		"template_content": toWechatContent(req.Content),
		"official_link":    req.OfficialLink,
		"miniprogram_link": req.MiniprogramLink,
	}

	if err := s.database.Conn.Model(&model.WechatTemplate{}).Where("id = ?", template.ID).Updates(update).Error; nil != err {
		util.Logger.WithFields("update", logger.MakeFields(update)).WithError(err).Error("修改微信模版失败")
		return &proto.UpdateWechatTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("修改数据库失败:%s", err.Error()),
		}, nil
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

	return &proto.UpdateWechatTemplateResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// SendWechatTemplate 发送微信模版消息
func (s *Server) SendWechatTemplate(ctx context.Context, req *proto.SendWechatTemplateRequest) (*proto.SendWechatTemplateResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("SendWechatTemplate")

	var template model.WechatTemplate
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.TemplateId)).First(&template).Error; nil != err {
		return &proto.SendWechatTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误：%s", err.Error()),
		}, nil
	}

	res, err := wechat.SendWechatTemplate(ctx, s.crius, template, req)
	if nil != err {
		go s.saveWechatTemplateRecord(ctx, req, res, model.WechatStatusFail)
		util.Logger.WithMetadata(ctx).WithError(err).Error("发送微信模版消息错误")
		return &proto.SendWechatTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("发送微信模版消息失败:%s", err.Error()),
		}, nil
	}

	util.Logger.WithMetadata(ctx).WithFields("send template response", logger.MakeFields(res)).Info("发送微信模版消息返回")

	go s.saveWechatTemplateRecord(ctx, req, res, model.WechatStatusSuccess)

	return &proto.SendWechatTemplateResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

func (s *Server) saveWechatTemplateRecord(ctx context.Context, req *proto.SendWechatTemplateRequest, res *wechat.SendWechatTemplateResponse, status model.WechatStatus) {
	metadata := pkgs.GetMetadata(ctx)
	branchID := uuid.FromStringOrNil(req.BranchId)
	wechatResponse := pkgs.MakeParams(res)
	request := pkgs.MakeParams(req)
	stat := model.WechatStat{
		ID:             uuid.NewV4(),
		MerchantID:     &metadata.MerchantID,
		BranchID:       &branchID,
		MessageType:    req.MessageType,
		System:         req.System,
		Request:        &request,
		WechatResponse: &wechatResponse,
		Status:         status,
		CreatedAt:      pkgs.NullTime{},
		UpdatedAt:      pkgs.NullTime{},
	}

	if req.WechatUser != nil {
		if len(req.WechatUser.MemberId) > 0 {
			memberID := uuid.FromStringOrNil(req.WechatUser.MemberId)
			stat.MemberID = &memberID
		}

		if len(req.WechatUser.MemberWechatId) > 0 {
			memberWechatID := uuid.FromStringOrNil(req.WechatUser.MemberWechatId)
			stat.MemberWechatID = &memberWechatID
		}
	}

	if err := s.database.Conn.Create(&stat).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("wechat stat", logger.MakeFields(stat)).WithError(err).Error("保存微信模版消息发送记录失败")
	}
}

func (s *Server) toProtoWechatTemplate(ctx context.Context, template model.WechatTemplate) *proto.WechatTemplate {
	res := &proto.WechatTemplate{
		Id:           template.ID.String(),
		TemplateName: template.TemplateName,
		TemplateCode: template.TemplateCode,
		Content:      nil,
		Status:       template.Status.String(),
		CreatedAt:    template.CreatedAt.Time.Unix(),
		Category:     template.Category,
		CategoryKey:  template.CategoryKey,
	}

	linkCache, err := cache.GetWechatLink(ctx, s.cache)
	if len(template.OfficialLink) > 0 {
		res.OfficialLink = template.OfficialLink
		if err == nil {
			res.OfficialLinkName = linkDesc(linkCache.Official, template.OfficialLink)
		}
	}

	if len(template.MiniprogramLink) > 0 {
		res.MiniprogramLink = template.MiniprogramLink
		if err == nil {
			res.MiniprogramLinkName = linkDesc(linkCache.Miniprogram, template.MiniprogramLink)
		}
	}

	content := &proto.WechatTemplateContent{}

	if template.TemplateContent != nil {
		content.First = &proto.WechatTemplateContentBase{
			Value: template.TemplateContent.First.Value,
			Color: template.TemplateContent.First.Color,
		}

		content.Remark = &proto.WechatTemplateContentBase{
			Value: template.TemplateContent.Remark.Value,
			Color: template.TemplateContent.Remark.Color,
		}

		var detail []*proto.WechatTemplateContentDetail
		for _, v := range template.TemplateContent.Detail {
			detail = append(detail, &proto.WechatTemplateContentDetail{
				Name:  v.Name,
				Value: v.Value,
				Color: v.Color,
			})
		}

		content.Detail = detail
	}

	res.Content = content
	return res
}

func linkDesc(items []cache.LinkDetail, link string) string {
	for _, v := range items {
		if v.URL == link {
			return v.Name
		}
	}

	return ""
}

func toWechatContent(content *proto.WechatTemplateContent) *model.WechatTemplateContent {
	res := &model.WechatTemplateContent{}

	if nil != content {
		if content.First != nil {
			res.First = model.WechatTemplateContentBase{
				Value: content.First.Value,
				Color: content.First.Color,
			}
		}

		if content.Remark != nil {
			res.Remark = model.WechatTemplateContentBase{
				Value: content.Remark.Value,
				Color: content.Remark.Color,
			}
		}

		var detail []model.WechatTemplateContentDetail
		for _, v := range content.Detail {
			detail = append(detail, model.WechatTemplateContentDetail{
				Name:  v.Name,
				Value: v.Value,
				Color: v.Color,
			})
		}

		res.Detail = detail
	}

	return res
}
