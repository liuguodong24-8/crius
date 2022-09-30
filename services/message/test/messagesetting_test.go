package test

import (
	"context"
	"fmt"

	//uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/message/proto"
)

func main() {
	ctx := context.Background()
	c := newClient()
	resp, err := c.client.CreateMessageSetting(ctx, &proto.CreateMessageSettingRequest{
		MessageType:              "appointment.deposit_pay_link",
		TriggerType:        "event",
		AdvanceHour: 0,
		SmsTemplateId: "626c62d8-2b8d-48ea-b1c8-2b3ab649ad10",
		WechatTemplateId: "1511e195-0d2f-40a5-b708-c1e97c58dc7f",
		Status: "opened",
	})
	fmt.Println(err)
	fmt.Println(resp)
}
