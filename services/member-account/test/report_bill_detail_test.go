package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/member-account/proto"
)

// Test_ReportBillDetail 账单详情
func Test_ReportBillDetail(t *testing.T) {
	fmt.Println("begin Test_ReportBillDetail ===========")
	client := newClient()

	//unix, _ := time.ParseInLocation(`2006-01-02`, "2021-08-01", time.Local)

	res, err := client.client.ReportBillDetail(newContext(), &proto.ReportBillDetailRequest{
		//BranchId:   "4b26a6e9-94a6-4876-810c-7c6ef408922f",
		//ReportType: []string{"open", "nobody"}, // recharge
		//BeginTime:  unix.Unix(),
		//EndTime:    unix.Unix() + 3600*24*7,
		WithPage:  true,
		BranchId:  "494e19ac-b3ca-4d2e-80d8-aca168152555",
		BeginTime: 1629446596,
		EndTime:   1629446701,
	})

	fmt.Println(err)

	fmt.Println(res)

	// [{promotion_option_id:"db9a873c-9d58-4bdf-8847-5244362ba198"  promotion_option_name:"option_name"  total:31  open_total:17  recharge_total:14  total_value:9300000}]
}
