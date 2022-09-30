package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_ListPointRule 积分规则列表
func Test_ListPointRule(t *testing.T) {
	fmt.Println("begin Test_ListPointRule ===========")
	client := newClient()

	res, err := client.ListPointRule(newContext(), &proto.ListPointRuleRequest{
		RuleName: "规则",
		Status:   "opened",
		Limit:    10,
		Offset:   0,
		OrderBy:  "created_at asc",
		WithPage: true,
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
