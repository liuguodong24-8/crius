package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_GetBranchRule 获取门店积分规则
func Test_GetBranchRule(t *testing.T) {
	fmt.Println("begin Test_GetBranchRule ===========")

	client := newClient()

	res, err := client.GetBranchPointRule(newContext(), &proto.GetBranchPointRuleRequest{
		BranchId:     "f969b836-c38b-4680-ad5f-19a28272dc49",
		CategoryCode: "003",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

	fmt.Println(res.Data)

}
