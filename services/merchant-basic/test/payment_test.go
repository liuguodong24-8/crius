package test

import (
	"context"
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

func Test_GetBranchWechatPaymentSetting(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetBranchWechatPaymentSetting(ctx, &proto.GetBranchWechatPaymentSettingRequest{
		BranchId: "0de40e41-d480-4f89-a6ae-a9f1e1837bff",
	})
	t.Log(err)
	fmt.Printf("%+v", resp)
}
