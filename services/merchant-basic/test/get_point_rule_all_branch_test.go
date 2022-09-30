package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_GetPointRuleAllBranch 获取积分规则已设置门店合集
func Test_GetPointRuleAllBranch(t *testing.T) {
	fmt.Println("begin Test_GetPointRuleAllBranch ===========")
	client := newClient()

	res, err := client.GetPointRuleAllBranch(newContext(), &proto.Empty{})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res.Data.BranchIds)

}
