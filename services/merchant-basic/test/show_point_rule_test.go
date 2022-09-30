package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_ShowPointRule 积分规则详情
func Test_ShowPointRule(t *testing.T) {
	fmt.Println("begin Test_ShowPointRule ===========")
	client := newClient()

	res, err := client.ShowPointRule(newContext(), &proto.ShowPointRuleRequest{
		//Id: "39d3771d-aeb1-4000-abe9-8a5e300a83c9", // 通过ID查询
		BranchId: "a3d0ffa8-d6a8-483e-8b11-445f529d3102", // 查询门店对应积分规则
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
