package test

import (
	"context"
	"encoding/json"
	"testing"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_CreateStaff(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.CreateStaff(ctx, &proto.CreateStaffRequest{Staff: &proto.StaffInfo{
		Name:         "bababab",              //门店名称
		Phone:        "bababa",               //门店联系电话
		EmployeeCode: "98hb7aa98r2321yg79qr", //员工工号
		PhoneCode:    "023",                  //区号
		Gender:       1,                      //性别
		EntryAt:      4564,                   //入职时间
	}, Branches: []string{"0028521d-17a7-4295-8dd4-95e1ff2690af"},
		Roles: []string{uuid.NewV4().String()}})
	t.Log(resp, err)
}

func Test_UpdateStaff(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateStaff(ctx, &proto.UpdateStaffRequest{Staff: &proto.StaffInfo{
		Id:           "a7ffd5e1-9838-4638-a1a4-cfd4073b7859",
		Name:         "bababab",       //门店名称
		Phone:        "bababa",        //门店联系电话
		EmployeeCode: "4fs56d4f65sd4", //员工工号
		PhoneCode:    "023",           //区号
		Gender:       1,               //性别
		EntryAt:      4564,            //入职时间
	}, Branches: []string{"0028521d-17a7-4295-8dd4-95e1ff2690af"},
		Roles: []string{uuid.NewV4().String()}})
	t.Log(resp, err)
}

func Test_GetStaffs(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetStaffs(ctx, &proto.GetStaffsRequest{
		Status: "closed",
	})
	bs, _ := json.Marshal(resp)
	t.Logf("%v                     %v", string(bs), err)
}

func Test_GetStaffsByRoleID(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetStaffsByRoleID(ctx, &proto.GetStaffsByRoleIDRequest{
		RoleId: "a0217b7c-72ae-49ce-88e5-6471d088640e",
	})
	bs, _ := json.Marshal(resp)
	t.Logf("%v                     %v", string(bs), err)
}

func Test_DeleteStaff(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.DeleteStaff(ctx, &proto.DeleteStaffRequest{
		Id: "662f6f3f-692b-4f23-bf13-d3ac87bb2a70",
	})
	t.Log(resp, err)
}

func Test_SignIn(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.SignIn(ctx, &proto.SignInRequest{
		Username: "omytech",
		Password: "omytech2021",
	})
	t.Logf("%+v   ================ %v", resp, err)
}

func Test_UpdatePassword(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdatePassword(ctx, &proto.UpdatePasswordRequest{
		Id:          "a7ffd5e1-9838-4638-a1a4-cfd4073b7859",
		NewPassword: "654321",
	})
	t.Logf("%+v   ================ %v", resp, err)
}

func Test_ResetPassword(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.ResetPassword(ctx, &proto.ResetPasswordRequest{
		Id: "a7ffd5e1-9838-4638-a1a4-cfd4073b7859",
	})
	t.Logf("%+v   ================ %v", resp, err)
}

func Test_UpdateStaffStatus(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateStaffStatus(ctx, &proto.UpdateStaffStatusRequest{
		Id:     "a7ffd5e1-9838-4638-a1a4-cfd4073b7859",
		Status: "opened",
	})
	t.Logf("%+v   ================ %v", resp, err)
}

func Test_ShowStaff(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.ShowStaff(ctx, &proto.ShowStaffRequest{
		Id: "cda8145e-5c43-4590-a622-36e963b4423b",
	})
	t.Logf("%+v   ================ %v", resp, err)
}
