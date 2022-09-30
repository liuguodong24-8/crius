package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/member-private/proto"
)

// Test_ListStaffShift 员工交班信息列表
func Test_ListStaffShift(t *testing.T) {
	fmt.Println("begin Test_ListStaffShift ===========")
	client := newClient()

	res, err := client.client.ListStaffShift(newContext(), &proto.ListStaffShiftRequest{
		BranchId: "f969b836-c38b-4680-ad5f-19a28272dc49",
		BeginAt:  0,
		EndAt:    0,
		OrderBy:  "",
	})

	fmt.Println(err)

	fmt.Println(res)
	fmt.Println(res.Data)

}
