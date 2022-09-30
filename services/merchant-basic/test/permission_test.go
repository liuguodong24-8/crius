package test

import (
	"context"
	"encoding/json"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_GetPermissions(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetPermissions(ctx, &proto.GetPermissionsRequest{
		Id: "26ff46f5-38d5-4906-a221-08a1349f7fe9",
	})
	bs, _ := json.Marshal(resp)
	t.Logf("%+v   ================ %v", string(bs), err)
}

func Test_CreatePermissions(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.CreatePermissions(ctx, &proto.CreatePermissionsRequest{
		Permissions: []*proto.PermissionInfo{
			{Id: 123, Permission: "fds5fsdfsd4564", Service: "4564554"},
			{Id: 124, Permission: "fdsf", Service: "4564554"},
			{Id: 125, Permission: "fds5fsdfsd4564", Service: "4564554"},
			{Id: 126, Permission: "dsfds", Service: "4564554"},
		},
		Service: "4564554",
	})
	t.Logf("%+v   ================ %v", resp, err)
}
