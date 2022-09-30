package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/member-account/proto"
)

// Test_ReportPayment 员工交班信息列表
func Test_ReportPayment(t *testing.T) {
	fmt.Println("begin Test_ReportPayment ===========")
	client := newClient()

	//unix, _ := time.ParseInLocation(`2006-01-02`, "2021-08-01", time.Local)

	res, err := client.client.ReportPayment(newContext(), &proto.ReportPaymentRequest{
		BranchId:  "b6e1eb04-321d-4123-b093-456aca53847e",
		BeginTime: 1629684720,
		EndTime:   1629688080,
	})

	fmt.Println(err)

	fmt.Println(res)

	// [{way:"wechat"  way_desc:"微信"  open_fee:360000  recharge_fee:260000}]
}
