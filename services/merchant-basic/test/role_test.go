package test

import (
	"context"
	"encoding/json"
	"testing"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_CreateRole(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.CreateRole(ctx, &proto.CreateRoleRequest{
		Name:        "店长",
		Property:    1,
		Permissions: []int32{1, 2, 3, 4, 5, 6},
	})
	t.Logf("%+v   ================ %v", resp, err)
}

func Test_UpdateRole(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateRole(ctx, &proto.UpdateRoleRequest{
		Id:          "c9aab6f5-a8d8-4518-89d7-fd4f9d1e2942",
		Name:        "店长1",
		Property:    2,
		Permissions: []int32{1, 2, 3, 4, 5, 7},
	})
	t.Logf("%+v   ================ %v", resp, err)
}

func Test_GetRoles(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetRoles(ctx, &proto.GetRolesRequest{
		Name: "店长",
	})
	bs, _ := json.Marshal(resp)
	t.Logf("%+v   ================ %v", string(bs), err)
}

func Test_DeleteRole(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.DeleteRole(ctx, &proto.DeleteRoleRequest{
		Id: uuid.NewV4().String(),
	})
	t.Logf("%+v   ================ %v", resp, err)
}

func Test_UpdateRoleStatus(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateRoleStatus(ctx, &proto.UpdateRoleStatusRequest{
		Id:     "c9aab6f5-a8d8-4518-89d7-fd4f9d1e2942",
		Status: "opened",
	})
	bs, _ := json.Marshal(resp)
	t.Log(string(bs))
	t.Log(err)
}

func Test_GetRoleHistories(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetRoleHistories(ctx, &proto.GetRoleHistoriesRequest{
		Id: "9e4523a3-0b76-42ed-996a-afca8de73276",
	})
	bs, _ := json.Marshal(resp)
	t.Log(string(bs))
	t.Log(err)
}

func Test_ShowRole(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.ShowRole(ctx, &proto.ShowRoleRequest{
		Id: "c9aab6f5-a8d8-4518-89d7-fd4f9d1e2942",
	})
	t.Logf("%+v   ================ %v", resp, err)
}
