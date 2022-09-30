package test

import (
	"context"
	"encoding/json"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_CreateBranch(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.CreateBranch(ctx, &proto.CreateBranchRequest{Branch: &proto.BranchInfo{
		Name:       "fdsfsd",  //门店名称
		ProvinceId: "123456",  //省ID
		CityId:     "4564",    //城市ID
		DistrictId: "456456",  //区ID
		Address:    "bababab", //门店地址
		Phone:      "bababa",  //门店联系电话
		Longitude:  123.123,
		Latitude:   123.456,
	}})
	t.Log(resp, err)
}

func Test_UpdateBranch(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateBranch(ctx, &proto.UpdateBranchRequest{Branch: &proto.BranchInfo{
		Id:         "08ac2585-2b64-48ff-8c42-7016b8740784",
		Name:       "123456",        //门店名称
		ProvinceId: "fdsfsdfdsfds",  //省ID
		CityId:     "fsdfsdfsd",     //城市ID
		DistrictId: "fsdfsdfsd",     //区ID
		Address:    "string5446456", //门店地址
		Phone:      "4546546789",    //门店联系电话
	}})
	t.Log(resp, err)
}

func Test_GetBranches(t *testing.T) {
	c := newClient()
	resp, err := c.GetBranches(newContext(), &proto.GetBranchesRequest{
		Status: "opened",
		Limit:  20,
	})
	bs, _ := json.Marshal(resp)
	t.Logf("%s ================ %v", string(bs), err)
}

func Test_GetBranchesByTagIDs(t *testing.T) {
	c := newClient()
	resp, err := c.GetBranchesByTagIDs(newContext(), &proto.GetBranchesByTagIDsRequest{
		TagIds: []string{"c349547c-5af7-4d59-9442-8c1e56fccd5d", "9d5a12b7-8f9d-4831-8162-bdf66777286d"},
		// Status: "opened",
		// BusinessStatus: []string{"opening"},
	})
	bs, _ := json.Marshal(resp)
	t.Logf("%s ================ %v", string(bs), err)
}

func Test_DeleteBranches(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.DeleteBranch(ctx, &proto.DeleteBranchRequest{
		Id: "6085e75c-f698-4df7-a5f6-872b142ea7bc",
	})
	t.Log(resp, err)
}

func Test_ShowBranch(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.ShowBranch(ctx, &proto.ShowBranchRequest{
		Id: "0028521d-17a7-4295-8dd4-95e1ff2690af",
	})
	t.Log(resp, err)
}

func Test_UpdateBranchStatus(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateBranchStatus(ctx, &proto.UpdateBranchStatusRequest{
		Id:     "0028521d-17a7-4295-8dd4-95e1ff2690af",
		Status: "closed",
	})
	t.Log(resp, err)
}
