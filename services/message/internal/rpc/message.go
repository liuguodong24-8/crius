package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/message/internal/model"
	"gitlab.omytech.com.cn/micro-service/message/proto"
)

// GetMessageVariable 获取变量
func (s *Server) GetMessageVariable(ctx context.Context, req *proto.Empty) (*proto.MessageVariableResponse, error) {
	util.Logger.WithMetadata(ctx).Info("GetMessageVariable")
	path, err := os.Getwd()
	if err != nil {
		util.Logger.WithError(err).Error("获取配置文件错误")
		return &proto.MessageVariableResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: "获取配置文件错误",
		}, nil
	}
	file := fmt.Sprintf("%s/config/variables.json", path)
	exists, err := util.FileExists(file)
	if err != nil {
		util.Logger.WithError(err).Error("获取配置文件错误，校验文件是否存在")
		return &proto.MessageVariableResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: "获取配置文件错误, 校验文件是否存在",
		}, nil
	}
	if !exists {
		util.Logger.WithError(err).Error("获取配置文件错误，不存在")
		return &proto.MessageVariableResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: "获取配置文件错误, 不存在",
		}, nil
	}

	data, err := dealJSONFile(file)
	if err != nil {
		util.Logger.WithError(err).Error("获取配置文件错误，解析错误")
		return &proto.MessageVariableResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: "获取配置文件错误, 解析错误",
		}, nil
	}

	return &proto.MessageVariableResponse{
		ErrorCode:    0,
		ErrorMessage: "",
		Data:         data,
	}, nil
}

// Message 消息变量
type Message struct {
	Category       string   `json:"category"`
	CategoryKey    string   `json:"category_key"`
	Variables      []string `json:"variables"`
	Trigger        string   `json:"trigger"`
	SettingDisable bool     `json:"setting_disable"`
}

// Variable 消息变量
type Variable struct {
	System    string    `json:"system"`
	SystemKey string    `json:"system_key"`
	Message   []Message `json:"message"`
}

func dealJSONFile(file string) ([]*proto.MessageVariableResponse_Variable, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var items []Variable
	if err := json.Unmarshal(bytes, &items); nil != err {
		return nil, err
	}

	response := make([]*proto.MessageVariableResponse_Variable, 0)

	for _, item := range items {
		message := make([]*proto.MessageVariableResponse_Variable_Message, 0)

		for _, v := range item.Message {
			message = append(message, &proto.MessageVariableResponse_Variable_Message{
				Category:       v.Category,
				CategoryKey:    v.CategoryKey,
				Variables:      v.Variables,
				Trigger:        v.Trigger,
				SettingDisable: v.SettingDisable,
			})
		}

		response = append(response, &proto.MessageVariableResponse_Variable{
			System:    item.System,
			SystemKey: item.SystemKey,
			Message:   message,
		})
	}

	return response, nil
}

// GetBranchTemplate 获取门店模版信息
func (s *Server) GetBranchTemplate(ctx context.Context, req *proto.GetBranchTemplateRequest) (*proto.GetBranchTemplateResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("GetBranchTemplate")
	if len(req.MessageType) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("获取门店模版信息错误")
		return &proto.GetBranchTemplateResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}
	metadata := pkgs.GetMetadata(ctx)
	var settings []model.MessageSetting
	if err := s.database.Conn.Model(&model.MessageSetting{}).Scopes(
		util.ColumnEqualScope("status", util.StatusOpened),
		util.ColumnEqualScope("message_type", req.MessageType),
		util.ColumnEqualScope("merchant_id", metadata.MerchantID)).
		Order("special_branches,updated_at desc").Find(&settings).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).WithError(err).Error("查询数据库错误")
		return &proto.GetBranchTemplateResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库操作错误:%s", err.Error()),
		}, err
	}

	date := time.Now()
	if req.Time > 0 {
		date = time.Unix(req.Time, 0)
	}
	// 按修改时间排序判断
	for _, setting := range settings {
		// 门店不符合 跳过
		if !judgeSpecialBranches(setting.SpecialBranches, req.BranchId) {
			continue
		}

		smsTemplateID, wechatTemplateID := judgeTemplateID(setting, date)

		cc := make([]*proto.GetBranchTemplateResponse_GetBranchTemplateResponseDate_Cc, 0)
		if nil != setting.CcList {
			for _, c := range *setting.CcList {
				cc = append(cc, &proto.GetBranchTemplateResponse_GetBranchTemplateResponseDate_Cc{
					Code:  c.Code,
					Phone: c.Phone,
				})
			}
		}

		return &proto.GetBranchTemplateResponse{
			ErrorCode:    0,
			ErrorMessage: "",
			Data: &proto.GetBranchTemplateResponse_GetBranchTemplateResponseDate{
				Id:             setting.ID.String(),
				TriggerType:    setting.TriggerType,
				AdvanceHour:    int32(setting.AdvanceHour),
				SmsTemplate:    s.getSmsTemplate(ctx, smsTemplateID),
				WechatTemplate: s.getWechatTemplate(ctx, wechatTemplateID),
				Cc:             cc,
			},
		}, nil
	}

	return &proto.GetBranchTemplateResponse{
		ErrorCode:    pkgs.ErrNotFound,
		ErrorMessage: "没有对应门店模版数据",
	}, nil
}

func judgeTemplateID(setting model.MessageSetting, date time.Time) (smsTemplateID string, wechatTemplateID string) {

	if setting.SmsTemplateID.Valid {
		smsTemplateID = setting.SmsTemplateID.UUID.String()
	}

	if setting.WechatTemplateID.Valid {
		wechatTemplateID = setting.WechatTemplateID.UUID.String()
	}

	// 没有特殊配置
	if nil != setting.SpecialSetting && len(*setting.SpecialSetting) > 0 {
		for _, s := range *setting.SpecialSetting {
			if s.Begin.Time.Before(date) && s.End.Time.Add(time.Hour*24).After(date) {
				if s.SmsTemplateID.Valid {
					smsTemplateID = s.SmsTemplateID.UUID.String()
				}

				if s.WechatTemplateID.Valid {
					wechatTemplateID = s.WechatTemplateID.UUID.String()
				}
				break
			}
		}
	}

	return
}

func judgeSpecialBranches(items *fields.UUIDArr, branchID string) bool {
	if items == nil || len(*items) == 0 {
		return true
	}

	for _, b := range *items {
		if b.String() == branchID {
			return true
		}
	}

	return false
}

func (s *Server) getSmsTemplate(ctx context.Context, id string) *proto.SmsTemplate {
	if len(id) == 0 {
		util.Logger.WithMetadata(ctx).Error(fmt.Sprintf("获取短信模版错误:%s", id))
		return nil
	}
	var smsTemplate model.SmsTemplate
	if err := s.database.Conn.Model(&model.SmsTemplate{}).Where("id = ?", id).First(&smsTemplate).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error(fmt.Sprintf("获取短信模版错误:%s", id))

		return nil
	}

	return toProtoSmsTemplate(smsTemplate)
}

func (s *Server) getWechatTemplate(ctx context.Context, id string) *proto.WechatTemplate {
	if len(id) == 0 {
		util.Logger.WithMetadata(ctx).Error(fmt.Sprintf("获取微信模版错误:%s", id))
		return nil
	}
	var wechatTemplate model.WechatTemplate
	if err := s.database.Conn.Model(&model.WechatTemplate{}).Where("id = ?", id).First(&wechatTemplate).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error(fmt.Sprintf("获取微信模版错误:%s", id))

		return nil
	}

	return s.toProtoWechatTemplate(ctx, wechatTemplate)
}
