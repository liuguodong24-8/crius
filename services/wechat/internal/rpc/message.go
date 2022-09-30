package rpc

import (
	"context"
	"fmt"

	"github.com/fideism/golang-wechat/officialaccount/message"

	"gitlab.omytech.com.cn/micro-service/wechat/internal/service"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/wechat/proto"
)

// SendTemplateMessage 发送模版消息
func (s *Server) SendTemplateMessage(ctx context.Context, req *proto.SendTemplateMessageRequest) (*proto.SendTemplateMessageResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("SendTemplateMessage")

	official, _, officialErr := service.NewOfficial(req.Channel)
	if officialErr != nil {
		util.Logger.WithMetadata(ctx).WithError(officialErr).Error("获取公众号配置错误")
		return &proto.SendTemplateMessageResponse{
			ErrorCode:    pkgs.ErrNotFound,
			ErrorMessage: fmt.Sprintf("获取公众号配置错误:%s", officialErr.Error()),
		}, nil
	}

	manager := official.GetTemplate()
	msg := &message.TemplateMessage{
		ToUser:     req.Touser,
		TemplateID: req.TemplateId,
		URL:        req.Url,
		Color:      req.Color,
		Data:       nil,
	}

	if req.MiniProgram != nil {
		msg.MiniProgram = struct {
			AppID    string `json:"appid"`
			PagePath string `json:"pagepath"`
		}{
			AppID:    req.MiniProgram.Appid,
			PagePath: req.MiniProgram.Pagepath,
		}
	}

	data := make(map[string]*message.TemplateDataItem)
	for _, v := range req.Data {
		data[v.Name] = &message.TemplateDataItem{
			Value: v.Value,
			Color: v.Color,
		}
	}
	msg.Data = data

	res, err := manager.Send(msg)
	util.Logger.WithMetadata(ctx).WithFields("wechat response", logger.MakeFields(res)).Info("发送微信模版消息返回")
	if err != nil {
		return &proto.SendTemplateMessageResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("请求微信失败:%s", err.Error()),
		}, nil
	}

	return &proto.SendTemplateMessageResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.SendTemplateMessageResponse_Data{
			ErrCode: res.ErrCode,
			ErrMsg:  res.ErrMsg,
			MsgId:   res.MsgID,
		},
	}, nil
}
