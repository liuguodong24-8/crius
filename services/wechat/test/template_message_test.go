package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/wechat/proto"
	"google.golang.org/grpc/metadata"
)

// Test_SendTemplateMessage 模版消息
func Test_SendTemplateMessage(t *testing.T) {
	fmt.Println("begin Test_OfficialGetAuthURL")
	ids := []string{"1d6fac48-77df-4395-8a88-e1ec425baffe"}
	arr, _ := fields.StringArrToUUIDArr(ids)

	md := metadata.New(map[string]string{
		"sleuth_code": fmt.Sprintf("%d", time.Now().Unix()),
		"merchant_id": "1d6fac48-77df-4395-8a88-e1ec425baffe",
		"staff_id":    uuid.NewV4().String(),
		"branch_ids":  arr.ToMetadataString(),
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	client := newClient()

	var data []*proto.SendTemplateMessageRequest_Data
	data = append(data, &proto.SendTemplateMessageRequest_Data{
		Name:  "first",
		Value: "预订成功通知",
		Color: "#173177",
	})
	data = append(data, &proto.SendTemplateMessageRequest_Data{
		Name:  "keyword1",
		Value: "天河店",
		Color: "#173177",
	})
	data = append(data, &proto.SendTemplateMessageRequest_Data{
		Name:  "keyword2",
		Value: "08号",
		Color: "#173177",
	})
	data = append(data, &proto.SendTemplateMessageRequest_Data{
		Name:  "keyword3",
		Value: "10分钟",
		Color: "#173177",
	})
	data = append(data, &proto.SendTemplateMessageRequest_Data{
		Name:  "remark",
		Value: "点击详情可以跳转到",
		Color: "#173177",
	})
	res, err := client.client.SendTemplateMessage(ctx, &proto.SendTemplateMessageRequest{
		Channel:     "appointment",
		Touser:      "oyJUpv0VOy-l1sny2fw-xU6nhgCY",
		TemplateId:  "-iPoL-8QzzFvEOZAhRFPFgbrNlE-LZ_QMOrt0_aaFfM",
		Url:         "https://www.baidu.com",
		Color:       "",
		MiniProgram: nil,
		Data:        data,
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
	fmt.Println(res.Data.MsgId)
}
