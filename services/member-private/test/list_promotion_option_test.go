package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/member-private/proto"
)

// Test_ListPromotionOption 优惠方案列表
func Test_ListPromotionOption(t *testing.T) {
	fmt.Println("begin Test_ListPromotionOption ===========")
	client := newClient()

	res, err := client.client.ListPromotionOption(newContext(), &proto.ListPromotionOptionRequest{
		Status:   "opened",
		OrderBy:  "recharge_value",
		WithPage: false,
		TagIds:   []string{"ef9e133a-2885-4abf-812b-953476568280"},
		FilterPromotion: &proto.ListPromotionOptionRequest_FilterPromotion{
			Status:   "opened",
			BranchId: "ab2bae4a-8952-41b6-9e22-24f46512ecc6",
		},
	})

	fmt.Println(err)

	fmt.Println(res)
	fmt.Println(res.Data)

}
