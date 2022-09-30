package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_GetPointRuleDescribe 获取积分规则说明
func Test_GetPointRuleDescribe(t *testing.T) {
	fmt.Println("begin Test_GetPointRuleDescribe ===========")
	client := newClient()

	res, err := client.GetPointRuleDescribe(newContext(), &proto.Empty{})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
