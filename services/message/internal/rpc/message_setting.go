package rpc

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cyrnicolase/nulls"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/message/internal/model"
	"gitlab.omytech.com.cn/micro-service/message/proto"
	"gorm.io/gorm"
)

// CreateMessageSetting 创建发送消息设置
func (s *Server) CreateMessageSetting(ctx context.Context, req *proto.CreateMessageSettingRequest) (*proto.CreateMessageSettingResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreateMessageSetting")
	metadata := pkgs.GetMetadata(ctx)
	if len(metadata.MerchantID.String()) == 0 || len(req.MessageType) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("新增消息设置，参数错误")
		return &proto.CreateMessageSettingResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "新增消息设置参数错误",
		}, nil
	}

	if checkSpecialBranchRepeated(s.database.Conn, metadata.MerchantID, req.SpecialBranches, req.MessageType, ``) {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("新建消息设置特殊门店重复")
		return &proto.CreateMessageSettingResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "该消息类型配置已存在，不可重复配置",
		}, nil
	}

	var smsTempalte model.SmsTemplate
	var wechatTempalte model.WechatTemplate
	if err := s.database.Conn.Model(&model.SmsTemplate{}).First(&smsTempalte,"id = ?", req.SmsTemplateId).Error; nil != err {
		if err == gorm.ErrRecordNotFound {
			return &proto.CreateMessageSettingResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "短信模板不存在",
			}, nil
		}
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("查询短信模板错误")
		return &proto.CreateMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: err.Error(),
		}, nil
	}
	if err := s.database.Conn.Model(&model.WechatTemplate{}).First(&wechatTempalte,"id = ?", req.WechatTemplateId).Error; nil != err {
		if err == gorm.ErrRecordNotFound {
			return &proto.CreateMessageSettingResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "微信模板不存在",
			}, nil
		}
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("查询微信模板错误")
		return &proto.CreateMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: err.Error(),
		}, nil
	}

	//获取配置文件
	data, err := getVariables()
	if err != nil {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error(err.Error())
		return &proto.CreateMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: err.Error(),
		}, nil
	}
	messageTypeName := getSmsTempalte(data, req.MessageType)

	extra := pkgs.MakeParams(map[string]interface{}{
		`message_type_name`:    messageTypeName,
		`sms_template_name`:    smsTempalte.Name,
		`wechat_template_name`: wechatTempalte.TemplateName,
	})

	messageSetting := model.MessageSetting{
		ID:          uuid.NewV4(),
		MerchantID:  metadata.MerchantID,
		MessageType: req.MessageType,
		TriggerType: req.TriggerType,
		AdvanceHour: req.AdvanceHour,
		Status:      util.StringToStatus(req.Status),
		Extra:       &extra,
	}

	if len(req.SmsTemplateId) > 0 {
		messageSetting.SmsTemplateID = nulls.NewUUID(uuid.FromStringOrNil(req.SmsTemplateId))
	}

	if len(req.WechatTemplateId) > 0 {
		messageSetting.WechatTemplateID = nulls.NewUUID(uuid.FromStringOrNil(req.WechatTemplateId))
	}

	if len(req.SpecialSetting) > 0 {
		req.SpecialSetting, err = getSpecialSetting(s.database.Conn, req.SpecialSetting)
		if nil != err {
			util.Logger.WithMetadata(ctx).WithFields("messageSetting", logger.MakeFields(messageSetting)).WithError(err).Error("组装微信模板短信模板名错误")
			return &proto.CreateMessageSettingResponse{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: fmt.Sprintf("组装微信模板短信模板名错误:%s", err.Error()),
			}, nil
		}
		messageSetting.SpecialSetting = normalizeSpecialSetting(req.SpecialSetting)
	}

	if len(req.CcList) > 0 {
		messageSetting.CcList = normalizeCcList(req.CcList)
	}

	if len(req.SpecialBranches) > 0 {
		messageSetting.SpecialBranches = normalizeSpecialBranches(req.SpecialBranches)
	}

	if err := s.database.Conn.Create(&messageSetting).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("messageSetting", logger.MakeFields(messageSetting)).WithError(err).Error("新增发送设置，数据库保存错误")
		return &proto.CreateMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库保存错误:%s", err.Error()),
		}, nil
	}

	after := pkgs.MakeParams(messageSetting)
	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SleuthCode:        "",
		SnapShotTableName: messageSetting.TableName(),
		TableID:           messageSetting.ID,
		Method:            model.CreateMethod,
		After:             &after,
	})

	return &proto.CreateMessageSettingResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// UpdateMessageSetting 创建发送消息设置
func (s *Server) UpdateMessageSetting(ctx context.Context, req *proto.UpdateMessageSettingRequest) (*proto.UpdateMessageSettingResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateMessageSetting")
	if len(req.Id) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改消息设置,参数错误")
		return &proto.UpdateMessageSettingResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "修改消息设置参数错误",
		}, nil
	}

	if checkSpecialBranchRepeated(s.database.Conn, metadata.MerchantID, req.SpecialBranches, req.MessageType, req.Id) {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改消息类型配置特殊门店重复")
		return &proto.UpdateMessageSettingResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "该消息类型配置已存在，不可重复配置",
		}, nil
	}

	// 查询
	var messageSetting model.MessageSetting
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&messageSetting).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.UpdateMessageSettingResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}
		return &proto.UpdateMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, nil
	}

	//获取配置文件
	data, err := getVariables()
	if err != nil {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error(err.Error())
		return &proto.UpdateMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: err.Error(),
		}, nil
	}
	messageTypeName := getSmsTempalte(data, req.MessageType)

	extra := pkgs.MakeParams(map[string]interface{}{
		`message_type_name`: messageTypeName,
	})


	req.SpecialSetting, err = getSpecialSetting(s.database.Conn, req.SpecialSetting)
	if nil != err {
		util.Logger.WithMetadata(ctx).WithFields("messageSetting", logger.MakeFields(messageSetting)).WithError(err).Error("组装微信模板短信模板名错误")
		return &proto.UpdateMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("组装微信模板短信模板名错误:%s", err.Error()),
		}, nil
	}

	update := map[string]interface{}{
		`message_type`:     req.MessageType,
		`trigger_type`:     req.TriggerType,
		`advance_hour`:     req.AdvanceHour,
		`special_branches`: normalizeSpecialBranches(req.SpecialBranches),
		`special_setting`:  normalizeSpecialSetting(req.SpecialSetting),
		`cc_list`:          normalizeCcList(req.CcList),
	}
	if len(req.SmsTemplateId) > 0 {
		update[`sms_template_id`] = req.SmsTemplateId
		var smsTempalte model.SmsTemplate
		if err := s.database.Conn.Model(&model.SmsTemplate{}).First(&smsTempalte,"id = ?", req.SmsTemplateId).Error; nil != err {
			if err == gorm.ErrRecordNotFound {
				return &proto.UpdateMessageSettingResponse{
					ErrorCode:    pkgs.ErrNotFound,
					ErrorMessage: "短信模板不存在",
				}, nil
			}
			util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("查询短信模板错误")
			return &proto.UpdateMessageSettingResponse{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: err.Error(),
			}, nil
		}
		extra.Set(`sms_template_name`, smsTempalte.Name)
	}
	if len(req.WechatTemplateId) > 0 {
		update[`wechat_template_id`] = req.WechatTemplateId
		var wechatTempalte model.WechatTemplate
		if err := s.database.Conn.Model(&model.WechatTemplate{}).First(&wechatTempalte,"id = ?", req.WechatTemplateId).Error; nil != err {
			if err == gorm.ErrRecordNotFound {
				return &proto.UpdateMessageSettingResponse{
					ErrorCode:    pkgs.ErrNotFound,
					ErrorMessage: "微信模板不存在",
				}, nil
			}
			util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("查询微信模板错误")
			return &proto.UpdateMessageSettingResponse{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: err.Error(),
			}, nil
		}
		extra.Set(`wechat_template_name`, wechatTempalte.TemplateName)
	}
	if len(req.Status) > 0 {
		update["status"] = util.StringToStatus(req.Status)
	}
	update[`extra`] = extra

	if err := s.database.Conn.Model(&model.MessageSetting{}).Where("id = ?", messageSetting.ID).Updates(update).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(update)).WithError(err).Error("修改消息设置失败")
		return &proto.UpdateMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}

	before := pkgs.MakeParams(messageSetting)
	after := pkgs.MakeParams(update)

	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: messageSetting.TableName(),
		TableID:           messageSetting.ID,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
	})

	return &proto.UpdateMessageSettingResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// ListMessageSetting 消息设置列表
func (s *Server) ListMessageSetting(ctx context.Context, req *proto.ListMessageSettingRequest) (*proto.ListMessageSettingResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ListMessageSetting")
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))
	if len(req.MessageType) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("message_type", req.MessageType))
	}
	if len(req.TriggerType) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("trigger_type", req.TriggerType))
	}
	if len(req.Status) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("status", req.Status))
	}

	var total int64
	var result []model.MessageSetting

	s.database.Conn.Model(&model.MessageSetting{}).Scopes(scopes...).Count(&total)
	if total == 0 {
		return &proto.ListMessageSettingResponse{
			ErrorCode:    pkgs.Success,
			ErrorMessage: "",
		}, nil
	}
	if req.WithPage {
		scopes = append(scopes, util.PaginationScope(req.Offset, req.Limit))
	}
	orderBy := "created_at desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}
	if err := s.database.Conn.Model(&model.MessageSetting{}).Scopes(scopes...).Order(orderBy).Find(&result).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("查询消息设置列表错误")
		return &proto.ListMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据查询错误:%s", err.Error()),
		}, nil
	}

	data := make([]*proto.MessageSetting, 0)
	for _, v := range result {
		data = append(data, toProtoMessageSetting(v))
	}

	return &proto.ListMessageSettingResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ListMessageSettingData{
			Settings: data,
			Total:    total,
		},
	}, nil
}

// ShowMessageSetting 消息设置详情
func (s *Server) ShowMessageSetting(ctx context.Context, req *proto.ShowMessageSettingRequest) (*proto.ShowMessageSettingResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowMessageSetting")
	if len(req.Id) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("查看消息设置详情，参数错误")
		return &proto.ShowMessageSettingResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "消息设置详情，参数错误",
		}, nil
	}

	var messageSetting model.MessageSetting
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&messageSetting).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.ShowMessageSettingResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}

		return &proto.ShowMessageSettingResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, nil
	}

	return &proto.ShowMessageSettingResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         toProtoMessageSetting(messageSetting),
	}, nil
}

// ChangeMessageSettingStatus 更改发送消息设置状态
func (s *Server) ChangeMessageSettingStatus(ctx context.Context, req *proto.ChangeMessageSettingStatusRequest) (*proto.ChangeMessageSettingStatusResponse, error) {
	defer util.CatchException()
	metadata := pkgs.GetMetadata(ctx)
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateMessageSetting")
	if len(req.Id) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改消息设置状态,参数错误")
		return &proto.ChangeMessageSettingStatusResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "修改消息设置状态参数错误",
		}, nil
	}

	// 查询
	var messageSetting model.MessageSetting
	if err := s.database.Conn.Scopes(util.ColumnEqualScope("id", req.Id)).First(&messageSetting).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.ChangeMessageSettingStatusResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "对应记录不存在",
			}, nil
		}
		return &proto.ChangeMessageSettingStatusResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, nil
	}

	update := map[string]interface{}{
		`status`: util.StringToStatus(req.Status),
	}
	if err := s.database.Conn.Model(&model.MessageSetting{}).Where("id = ?", messageSetting.ID).Updates(update).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(update)).WithError(err).Error("修改消息设置状态失败")
		return &proto.ChangeMessageSettingStatusResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}

	before := pkgs.MakeParams(messageSetting)
	after := pkgs.MakeParams(update)

	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: messageSetting.TableName(),
		TableID:           messageSetting.ID,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
	})

	return &proto.ChangeMessageSettingStatusResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

func toProtoMessageSetting(item model.MessageSetting) *proto.MessageSetting {
	ms := proto.MessageSetting{
		Id:              item.ID.String(),
		MessageType:     item.MessageType,
		MessageTypeName: item.Extra.Get(`message_type_name`).(string),
		TriggerType:     item.TriggerType,
		AdvanceHour:     item.AdvanceHour,
		SpecialSetting:  normalizeProtoSpecialSetting(item.SpecialSetting),
		CcList:          normalizeProtoCcList(item.CcList),
		SpecialBranches: item.SpecialBranches.ToStringArr(),
		Status:          item.Status.String(),
		CreatedAt:       item.CreatedAt.Time.Unix(),
	}

	if !uuid.Equal(item.SmsTemplateID.UUID, uuid.Nil) {
		ms.SmsTemplateId = item.SmsTemplateID.UUID.String()
		ms.SmsTemplateName = ``
		if item.Extra.Exists(`sms_template_name`) {
			ms.SmsTemplateName = item.Extra.Get(`sms_template_name`).(string)
		}
	}

	if !uuid.Equal(item.WechatTemplateID.UUID, uuid.Nil) {
		ms.WechatTemplateId = item.WechatTemplateID.UUID.String()
		ms.WechatTemplateName = ``
		if item.Extra.Exists(`wechat_template_name`) {
			ms.WechatTemplateName = item.Extra.Get(`wechat_template_name`).(string)
		}
	}

	return &ms
}

func normalizeProtoSpecialSetting(data *model.SpecialSetting) (result []*proto.SpecialSetting) {
	if data == nil {
		return
	}

	for _, item := range *data {
		s := &proto.SpecialSetting{
			Begin: item.Begin.String(),
			End:   item.End.String(),
		}

		if !uuid.Equal(item.SmsTemplateID.UUID, uuid.Nil) {
			s.SmsTemplateId = item.SmsTemplateID.UUID.String()
			s.SmsTemplateName = item.SmsTemplateName
		}
		if !uuid.Equal(item.WechatTemplateID.UUID, uuid.Nil) {
			s.WechatTemplateId = item.WechatTemplateID.UUID.String()
			s.WechatTemplateName = item.WechatTemplateName
		}

		result = append(result, s)
	}
	return
}

func normalizeProtoCcList(data *model.CcList) (result []*proto.Cc) {
	if data == nil {
		return
	}

	for _, item := range *data {
		result = append(result, &proto.Cc{
			Code:  item.Code,
			Phone: item.Phone,
		})
	}
	return
}

func normalizeSpecialSetting(data []*proto.SpecialSetting) *model.SpecialSetting {
	result := make(model.SpecialSetting, 0)
	for _, item := range data {
		setting := model.RangeSetting{
			Begin: fields.StringToDateTime(item.Begin),
			End:   fields.StringToDateTime(item.End),
		}

		stID := uuid.FromStringOrNil(item.SmsTemplateId)
		wtID := uuid.FromStringOrNil(item.WechatTemplateId)
		if uuid.Equal(stID, uuid.Nil) && uuid.Equal(wtID, uuid.Nil) {
			continue
		}

		if !uuid.Equal(stID, uuid.Nil) {
			setting.SmsTemplateID = nulls.NewUUID(stID)
			setting.SmsTemplateName = item.SmsTemplateName
		}
		if !uuid.Equal(wtID, uuid.Nil) {
			setting.WechatTemplateID = nulls.NewUUID(wtID)
			setting.WechatTemplateName = item.WechatTemplateName
		}

		result = append(result, setting)
	}

	if len(result) > 0 {
		return &result
	}

	return nil
}

func normalizeSpecialBranches(data []string) *fields.UUIDArr {
	if len(data) > 0 {
		result, _ := fields.StringArrToUUIDArr(data)
		return &result
	}

	return nil
}

func normalizeCcList(data []*proto.Cc) *model.CcList {
	result := make(model.CcList, 0)
	for _, item := range data {
		result = append(result, model.Cc{
			Code:  item.Code,
			Phone: item.Phone,
		})
	}

	if len(result) > 0 {
		return &result
	}

	return nil
}

func makeInterfaceSlice(items []string) []interface{} {
	var res []interface{}
	for _, item := range items {
		res = append(res, item)
	}

	return res
}

func checkSpecialBranchRepeated(db *gorm.DB, merchantID uuid.UUID, branches []string, messageType, id string) bool {
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(
		scopes,
		util.ColumnEqualScope("merchant_id", merchantID),
		util.ColumnEqualScope("message_type", messageType),
	)
	if len(branches) > 0 {
		scopes = append(scopes, util.ArrayOverlapScope(`special_branches`, `uuid`, makeInterfaceSlice(branches)))
	} else {
		scopes = append(scopes, util.ColumnNullScope(`special_branches`, true))
	}
	if id != `` {
		scopes = append(scopes, util.ColumnNotEqualScope(`id`, id))
	}

	var total int64
	db.Model(&model.MessageSetting{}).Scopes(scopes...).Count(&total)
	if total > 0 {
		return true
	}

	return false
}

// getVariables 获取配置文件
func getVariables() ([]*proto.MessageVariableResponse_Variable, error) {
	path, err := os.Getwd()
	if err != nil {
		util.Logger.WithError(err).Error("获取配置文件错误")
		return nil, err
	}
	file := fmt.Sprintf("%s/config/variables.json", path)
	exists, err := util.FileExists(file)
	if err != nil {
		util.Logger.WithError(err).Error("获取配置文件错误，校验文件是否存在")
		return nil, err
	}
	if !exists {
		util.Logger.WithError(err).Error("获取配置文件错误，不存在")
		return nil, err
	}

	data, err := dealJSONFile(file)
	if err != nil {
		util.Logger.WithError(err).Error("获取配置文件错误，解析错误")
		return nil, err
	}

	return data, nil
}
// getSmsTempalte 获取messageTypeName
func getSmsTempalte(data []*proto.MessageVariableResponse_Variable, MessageType string) string{
	var messageTypeName string
	Cycle:
	for _, v := range data  {
		for _, val := range v.Message {
			if val.CategoryKey == MessageType {

				messageTypeName = val.Category
				break Cycle
			}
		}
	}

	return messageTypeName
}

// getSpecialSetting 获取配置的微信模板名，短信模板名
func getSpecialSetting(db *gorm.DB, specialSetting []*proto.SpecialSetting) ([]*proto.SpecialSetting, error)  {
	var smsTemplateIds []string
	var wechatTemplateIds []string
	for _, item := range specialSetting {
		smsTemplateIds = append(smsTemplateIds, item.SmsTemplateId)
		wechatTemplateIds = append(wechatTemplateIds, item.WechatTemplateId)
	}

	var smsTempaltes []model.SmsTemplate
	var wechatTempaltes []model.WechatTemplate
	if err := db.Model(&model.SmsTemplate{}).Where("id in (?)", smsTemplateIds).Scan(&smsTempaltes).Error; nil != err {
		return nil, err
	}
	if err := db.Model(&model.WechatTemplate{}).Where("id in (?)", wechatTemplateIds).Scan(&wechatTempaltes).Error; nil != err {
		return nil, err
	}

	for k, val := range specialSetting {
		for _, v := range smsTempaltes{
			if v.ID.String() == val.SmsTemplateId {
				specialSetting[k].SmsTemplateName = v.Name
				break
			}
		}

		for _, v := range wechatTempaltes{
			if v.ID.String() == val.WechatTemplateId {
				specialSetting[k].WechatTemplateName = v.TemplateName
				break
			}
		}
	}
	return specialSetting, nil
}