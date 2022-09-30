package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_UpdatePointRuleStatus 修改积分规则状态
func Test_UpdatePointRuleStatus(t *testing.T) {
	fmt.Println("begin Test_UpdatePointRuleStatus ===========")
	client := newClient()
	res, err := client.UpdatePointRuleStatus(newContext(), &proto.UpdateStatusRequest{
		Id:     "39d3771d-aeb1-4000-abe9-8a5e300a83c9",
		Status: "closed",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
