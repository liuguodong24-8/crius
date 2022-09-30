package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/message/proto"
)

// Test_WechatStat 微信发送统计
func Test_WechatStat(t *testing.T) {
	client := newClient()
	ctx := newContext()

	b := fields.StringToDateTime("2010-01-01")
	e := fields.StringToDateTime("2022-01-01")

	res, err := client.client.WechatStat(ctx, &proto.WechatStatRequest{
		BranchId:    "39f9ee0c-e172-465c-a2c4-95b6e41ba16e",
		MessageType: "appointment.deposit_payed",
		BeginDate:   b.ToUnix(),
		EndDate:     e.ToUnix(),
		WithPage:    true,
		Limit:       1,
		Offset:      0,
	})

	fmt.Println(err)

	fmt.Println(res)
}
