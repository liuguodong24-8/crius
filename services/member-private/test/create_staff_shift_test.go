package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/member-private/proto"
)

// Test_CreateStaffShift 创建员工交班信息
func Test_CreateStaffShift(t *testing.T) {
	fmt.Println("begin Test_CreateStaffShift ===========")
	client := newClient()

	res, err := client.client.CreateStaffShift(newContext(), &proto.CreateStaffShiftRequest{
		BranchId: "f969b836-c38b-4680-ad5f-19a28272dc49",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
