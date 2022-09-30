package test

import (
	"context"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_ShowMemberByAccuratePhone(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	request := &proto.ShowMemberByAccuratePhoneRequest{
		Phone: "1566996451",
	}
	resp, err := c.ShowMemberByAccuratePhone(ctx, request)
	t.Log(resp, err)
}

func Test_CreateMember(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	request := &proto.CreateMemberRequest{
		Member: &proto.MemberInfo{
			Name:          "李四",
			Phone:         "15669964454",
			PhoneCode:     "86",
			Gender:        1,
			CityCode:      "510923",
			Channel:       "wechat",
			FirstBranchId: "0de40e41-d480-4f89-a6ae-a9f1e1837bff",
		},
	}
	resp, err := c.CreateMember(ctx, request)
	t.Log(resp, err)
}

func Test_CreateOrUpdateWechatUser(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	request := &proto.CreateOrUpdateWechatUserRequest{
		User: &proto.WechatUser{
			Openid:   "test111111",
			Appid:    "222222222222",
			MemberId: "361dd8f6-ba92-4334-9c8e-cba5a9368e57",
		},
	}
	resp, err := c.CreateOrUpdateWechatUser(ctx, request)
	t.Log(resp, err)
}
