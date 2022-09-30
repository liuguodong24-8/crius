package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_SetPointRuleDescribe 设置积分规则说明
func Test_SetPointRuleDescribe(t *testing.T) {
	fmt.Println("begin Test_SetPointRuleDescribe ===========")
	client := newClient()

	res, err := client.SetPointRuleDescribe(newContext(), &proto.SetPointRuleDescribeRequest{
		Images: []string{
			"http://5b0988e595225.cdn.sohucs.com/images/20180125/325146066393455d8cb9705bdd8900ae.jpeg",
			"http://newssrc.onlinedown.net/d/file/20160811/87c19dc05ef9272d3717d437bd192cf6.jpg",
		},
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
