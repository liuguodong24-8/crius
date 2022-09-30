package test

import (
	"fmt"
	"testing"
	"time"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_GetBranchLatelyBusiness 获取门店指定时间最近一次营业日
func Test_GetBranchLatelyBusiness(t *testing.T) {
	fmt.Println("begin Test_GetBranchLatelyBusiness ===========")

	request := &proto.GetBranchLatelyBusinessRequest{
		BranchId: "f969b836-c38b-4680-ad5f-19a28272dc49",
		DateTime: time.Now().Unix(),
	}

	client := newClient()

	res, err := client.GetBranchLatelyBusiness(newContext(), request)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

	fmt.Println(res.Data)

}
