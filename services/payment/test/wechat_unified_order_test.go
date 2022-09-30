package test

import (
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/payment/proto"
)

func Test_WechatUnifiedOrder(t *testing.T) {
	fmt.Println("test wechat order unified order")
	request := &proto.WechatUnifiedOrderRequest{
		BranchId: "0de40e41-d480-4f89-a6ae-a9f1e1837bff",
		StringMap: map[string]string{
			"body":        "测试微信预约",
			"attach":      uuid.NewV4().String(),
			"fee_type":    "CNY",
			"time_start":  time.Now().Format(`20060102150405`),
			"time_expire": time.Now().Add(time.Hour * 6).Format(`20060102150405`),
			"notify_url":  "http://test.wx.haochang.tv/api/v1/branch/notify",
			"openid":      "oyJUpv0VOy-l1sny2fw-xU6nhgCY",
			"trade_type":  "JSAPI",
		},
		Int64Map: map[string]int64{
			"total_fee": 1,
		},
	}

	client := newClient()

	res, err := client.client.WechatUnifiedOrder(newContext(), request)
	fmt.Println(res)

	fmt.Println(err)

	fmt.Println(res.Data)
}
