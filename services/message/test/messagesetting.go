package test

import (
	"context"
	"fmt"
	//uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/message/proto"
)

func TestMessageSetting() {
	ctx := context.Background()
	c := newClient()

	var specials []*proto.SpecialSetting

	s := proto.SpecialSetting{
		Begin :   "1111",
		End    :  "2222",
		SmsTemplateId   :  "40b91de9-47dd-4ed1-865e-a82eea40e9dd",
		WechatTemplateId :  "b92ddd0c-2c39-4e49-ad3f-b99745e8a020",
	}
	specials = append(specials, &s)

	/*resp, err := c.client.CreateMessageSetting(ctx, &proto.CreateMessageSettingRequest{
		MessageType:              "merchant.staff_edit",
		TriggerType:        "event",
		AdvanceHour: 0,
		SmsTemplateId: "40b91de9-47dd-4ed1-865e-a82eea40e9dd",
		WechatTemplateId: "b92ddd0c-2c39-4e49-ad3f-b99745e8a020",
		SpecialSetting: specials,
		Status: "opened",
	})*/
	resp, err := c.client.UpdateMessageSetting(ctx, &proto.UpdateMessageSettingRequest{
		Id: "76dd0a08-bde1-4631-9f02-0ad4c894fa50",
		MessageType:              "merchant.staff_edit",
		TriggerType:        "event",
		AdvanceHour: 0,
		SmsTemplateId: "40b91de9-47dd-4ed1-865e-a82eea40e9dd",
		WechatTemplateId: "b92ddd0c-2c39-4e49-ad3f-b99745e8a020",
		SpecialSetting: specials,
		Status: "opened",
	})
	fmt.Println(err)
	fmt.Println(resp)
}
