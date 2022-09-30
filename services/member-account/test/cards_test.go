package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/member-account/proto"
)

func Test_ActivePrimaryCard(t *testing.T) {
	fmt.Println("begin Test_ActivePrimaryCard ===========")
	client := newClient()

	res, err := client.client.ActivePrimaryCard(newContext(), &proto.ActivePrimaryCardRequest{
		Phone:         "13025412563",
		PhoneCode:     "86",
		Name:          "刘",
		CardId:        "ce30be5e-d331-481c-a292-5986b688b1c1",
		Gender:        1,
		BranchId:      "494e19ac-b3ca-4d2e-80d8-aca168152555",
		Payments:      &proto.Payments{Cash: 10000},
		Promotions:    []*proto.PromotionCount{{Count: 1, Id: "4197b2c5-55c8-4a0a-bf3d-3a749490b909"}},
		RechargeValue: 10000,
	})

	fmt.Println(err)

	fmt.Println(res)

	// [{way:"wechat"  way_desc:"微信"  open_fee:360000  recharge_fee:260000}]
}

func Test_RechargeCard(t *testing.T) {
	fmt.Println("begin Test_RechargeCard ===========")
	client := newClient()

	res, err := client.client.RechargeCard(newContext(), &proto.RechargeCardRequest{
		CardId:        "61e71f0e-5f94-4ca9-aa94-7ee8d6c34e62",
		BranchId:      "ab2bae4a-8952-41b6-9e22-24f46512ecc6",
		Payments:      &proto.Payments{Cash: 1000000},
		Promotions:    []*proto.PromotionCount{{Count: 1, Id: "5d831026-28e2-4ce0-b0bc-b4aa74559984"}},
		RechargeValue: 1000000,
	})

	fmt.Println(err)
	fmt.Println(res)
}
