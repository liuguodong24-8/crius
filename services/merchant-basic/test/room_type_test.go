package test

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_CreateRoomType(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.CreateRoomType(ctx, &proto.CreateRoomTypeRequest{
		BranchId:    uuid.NewV4().String(),
		Name:        "afafa",
		CategoryId:  uuid.NewV4().String(),
		Status:      "closed",
		CustomerMin: 1,
		CustomerMax: 2,
		Order:       999,
	})
	t.Log(resp, err)
}

func Test_GetRoomTypes(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetRoomTypes(ctx, &proto.GetRoomTypesRequest{
		Status: "opened",
	})
	t.Log(resp, err)
}

func Test_UpdateRoomType(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateRoomType(ctx, &proto.UpdateRoomTypeRequest{
		Id:     "8f369829-5e69-4175-ac0a-d5c26a07e35f",
		Status: "opened",
	})
	t.Log(resp, err)
}

func Test_ShowRoomType(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.ShowRoomType(ctx, &proto.ShowRoomTypeRequest{
		Id: "d390b6ba-2361-4c1f-afb9-fe81b274e6fb",
	})
	t.Log(resp, err)
}
