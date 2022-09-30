package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_CreateMemberAddress 创建用户地址
func Test_CreateMemberAddress(t *testing.T) {
	fmt.Println("begin Test_CreateMemberAddress ===========")
	client := newClient()

	res, err := client.CreateMemberAddress(newContext(), &proto.MemberAddress{
		MemberId:   "cdcf480d-0883-4545-bdb6-faab8181a803",
		Name:       "张三",
		Phone:      "13800138000",
		PhoneCode:  "86",
		ProvinceId: "54",
		CityId:     "540100",
		DistrictId: "540102",
		Address:    "西藏自治区拉萨市城关区",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}

// Test_GetMemberAddress 获取用户地址
func Test_GetMemberAddress(t *testing.T) {
	fmt.Println("begin Test_GetMemberAddress ===========")
	client := newClient()

	res, err := client.GetMemberAddress(newContext(), &proto.GetMemberAddressRequest{
		MemberId: "cdcf480d-0883-4545-bdb6-faab8181a803",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}

// Test_UpdateMemberAddress 修改用户地址
func Test_UpdateMemberAddress(t *testing.T) {
	fmt.Println("begin Test_UpdateMemberAddress ===========")
	client := newClient()

	res, err := client.UpdateMemberAddress(newContext(), &proto.MemberAddress{
		Id:         "64acc05f-e57e-498d-9302-bbd49a074921",
		MemberId:   "cdcf480d-0883-4545-bdb6-faab8181a803",
		Name:       "张三",
		Phone:      "13600136000",
		PhoneCode:  "86",
		ProvinceId: "54",
		CityId:     "540100",
		DistrictId: "540102",
		Address:    "西藏自治区拉萨市城关区布达拉哇哇哇",
		IsDefault:  false,
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}

// Test_GetMemberDefaultAddress 获取用户默认地址
func Test_GetMemberDefaultAddress(t *testing.T) {
	fmt.Println("begin Test_GetMemberAddress ===========")
	client := newClient()

	res, err := client.GetMemberDefaultAddress(newContext(), &proto.GetMemberDefaultAddressRequest{
		MemberId: "cdcf480d-0883-4545-bdb6-faab8181a803",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}

// Test_SetMemberAddressDefault 设置用户默认地址
func Test_SetMemberAddressDefault(t *testing.T) {
	fmt.Println("begin Test_SetMemberAddressDefault ===========")
	client := newClient()

	res, err := client.SetMemberAddressDefault(newContext(), &proto.SetMemberAddressDefaultRequest{
		Id:       "64acc05f-e57e-498d-9302-bbd49a074921",
		MemberId: "cdcf480d-0883-4545-bdb6-faab8181a803",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}

// Test_DeleteMemberAddress 删除用户地址
func Test_DeleteMemberAddress(t *testing.T) {
	fmt.Println("begin Test_DeleteMemberAddress ===========")
	client := newClient()

	res, err := client.DeleteMemberAddress(newContext(), &proto.DeleteMemberAddressRequest{
		Id: "64acc05f-e57e-498d-9302-bbd49a074921",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
