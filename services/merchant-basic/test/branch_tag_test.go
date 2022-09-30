package test

import (
	"context"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_CreateBranchTag(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.CreateBranchTag(ctx, &proto.CreateBranchTagRequest{
		Name: "abababab",
	})
	t.Log(resp, err)
}

func Test_UpdateBranchTag(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateBranchTag(ctx, &proto.UpdateBranchTagRequest{
		Id:   "c349547c-5af7-4d59-9442-8c1e56fccd5d",
		Name: "ababab发斯蒂芬斯蒂芬斯蒂芬ab",
	})
	t.Log(resp, err)
}

func Test_UpdateBranchTagStatus(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateBranchTagStatus(ctx, &proto.UpdateBranchTagStatusRequest{
		Id:     "c349547c-5af7-4d59-9442-8c1e56fccd5d",
		Status: "closed",
	})
	t.Log(resp, err)
}

func Test_GetBranchTags(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetBranchTags(ctx, &proto.GetBranchTagsRequest{
		Status: "closed",
	})
	t.Log(resp, err)
}

func Test_GetBranchTagsByIDs(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetBranchTagsByIDs(ctx, &proto.GetBranchTagsByIDsRequest{
		Ids: []string{"c349547c-5af7-4d59-9442-8c1e56fccd5d", "8a08f210-2a47-430d-83b1-5e0812b8e65a"},
	})
	t.Log(resp, err)
}

func Test_ShowBranchTag(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.ShowBranchTag(ctx, &proto.ShowBranchTagRequest{
		Id: "c349547c-5af7-4d59-9442-8c1e56fccd5d",
	})
	t.Log(resp, err)
}
