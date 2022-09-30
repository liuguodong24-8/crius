package test

import (
	"context"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_CreateDistrict(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.CreateDistrict(ctx, &proto.CreateDistrictRequest{
		Name:   "12346",
		Status: "closed",
	})
	t.Log(resp, err)
}

func Test_UpdateDistrict(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateDistrict(ctx, &proto.UpdateDistrictRequest{
		Id:     "1e5a0746-1093-4ae2-abf6-7454030febaa",
		Status: "closed",
	})
	t.Log(resp, err)
}

func Test_GetDistricts(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetDistricts(ctx, &proto.GetDistrictsRequest{
		Status: "closed",
	})
	t.Log(resp, err)
}
