package test

import (
	"context"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_CreateRoomTypeCategory(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.CreateRoomTypeCategory(ctx, &proto.CreateRoomTypeCategoryRequest{
		Name:     "afafa",
		Category: 1,
		Status:   "opened",
	})
	t.Log(resp, err)
}

func Test_GetRoomTypeCategories(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetRoomTypeCategories(ctx, &proto.GetRoomTypeCategoriesRequest{
		Status: "opened",
	})
	t.Log(resp, err)
}

func Test_UpdateRoomTypeCategory(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateRoomTypeCategory(ctx, &proto.UpdateRoomTypeCategoryRequest{
		Id:       "63854e3f-fa3e-4052-84f4-156a7b2e3d1a",
		Name:     "fsdfdsfsdfsd",
		Category: 1,
		Status:   "opened",
	})
	t.Log(resp, err)
}
