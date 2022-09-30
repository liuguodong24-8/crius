package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/member-account/proto"
)

// Test_ReportPromotion 优惠方案汇总
func Test_ReportPromotion(t *testing.T) {
	fmt.Println("begin Test_ReportPromotion ===========")
	client := newClient()

	//unix, _ := time.ParseInLocation(`2006-01-02`, "2021-08-01", time.Local)

	res, err := client.client.ReportPromotion(newContext(), &proto.ReportBillDetailRequest{
		BranchId:  "494e19ac-b3ca-4d2e-80d8-aca168152555",
		BeginTime: 1629446596,
		EndTime:   1629446701,
		WithPage:  true,
	})

	fmt.Println(err)

	fmt.Println(res)

	// [{promotion_option_id:"db9a873c-9d58-4bdf-8847-5244362ba198"  promotion_option_name:"option_name"  total:31  open_total:17  recharge_total:14  total_value:9300000}]
}
