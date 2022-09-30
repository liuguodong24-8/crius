package wechat

import (
	"context"
	"errors"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/message/internal/micro"
	"gitlab.omytech.com.cn/micro-service/message/internal/model"
	"gitlab.omytech.com.cn/micro-service/message/proto"
	wechatProto "gitlab.omytech.com.cn/micro-service/wechat/proto"
)

// SendWechatTemplateResponse 发送微信模版消息返回
type SendWechatTemplateResponse struct {
	CommonErr
	MsgID int64
}

// SendWechatTemplate 发送微信模版消息
func SendWechatTemplate(ctx context.Context, client *crius.Client, template model.WechatTemplate, req *proto.SendWechatTemplateRequest) (*SendWechatTemplateResponse, error) {
	metadata := pkgs.GetMetadata(ctx)
	newCtx := pkgs.MetadataContent(metadata)

	wechatClient, wechatErr := micro.GetWechatServer(newCtx, client)
	if wechatErr != nil {
		return nil, wechatErr
	}
	// 组装调用request
	request := toWechatSendWechatTemplateRequest(req, template)

	util.Logger.WithMetadata(ctx).WithFields("template request", logger.MakeFields(request)).Info("组装微信模版发送格式完成")

	res, err := wechatClient.Entity.SendTemplateMessage(newCtx, request)
	if err != nil {
		return nil, fmt.Errorf("发送wechat服务，发送模版消息失败:%s", err.Error())
	}

	if res.ErrorCode != pkgs.Success {
		return nil, fmt.Errorf("发送wechat服务，发送模版消息失败:code:%d, %s", res.ErrorCode, res.ErrorMessage)
	}

	if res.Data == nil {
		return nil, errors.New("发送wechat服务，发送模版消息失败, data nil")
	}

	return &SendWechatTemplateResponse{
		CommonErr: CommonErr{
			ErrCode: res.Data.ErrCode,
			ErrMsg:  res.Data.ErrMsg,
		},
		MsgID: res.Data.MsgId,
	}, nil
}

func toWechatSendWechatTemplateRequest(req *proto.SendWechatTemplateRequest, template model.WechatTemplate) *wechatProto.SendTemplateMessageRequest {
	request := &wechatProto.SendTemplateMessageRequest{
		Channel:     req.Channel,
		TemplateId:  template.TemplateCode,
		Url:         req.OfficialLink,
		Color:       "", // 暂无
		MiniProgram: nil,
	}

	// 详细内容
	if req.Content != nil {
		var data []*wechatProto.SendTemplateMessageRequest_Data

		if req.Content.First != nil {
			data = append(data, &wechatProto.SendTemplateMessageRequest_Data{
				Name:  "first",
				Value: req.Content.First.Value,
				Color: req.Content.First.Color,
			})
		}

		for _, v := range req.Content.Detail {
			data = append(data, &wechatProto.SendTemplateMessageRequest_Data{
				Name:  v.Name,
				Value: v.Value,
				Color: v.Color,
			})
		}

		if req.Content.Remark != nil {
			data = append(data, &wechatProto.SendTemplateMessageRequest_Data{
				Name:  "remark",
				Value: req.Content.Remark.Value,
				Color: req.Content.Remark.Color,
			})
		}

		request.Data = data
	}

	// 接收者
	if req.WechatUser != nil {
		request.Touser = req.WechatUser.MemberOpenId
	}

	// 小程序信息
	if req.Miniprogram != nil {
		request.MiniProgram = &wechatProto.SendTemplateMessageRequest_MiniProgram{
			Appid:    req.Miniprogram.Appid,
			Pagepath: req.Miniprogram.Pagepath,
		}
	}

	return request
}
