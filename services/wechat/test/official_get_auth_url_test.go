package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/wechat/proto"
	"google.golang.org/grpc/metadata"
)

// Test_OfficialGetAuthURL 获取公众号授权url
func Test_OfficialGetAuthURL(t *testing.T) {
	fmt.Println("begin Test_OfficialGetAuthURL")
	ids := []string{"1d6fac48-77df-4395-8a88-e1ec425baffe"}
	arr, _ := fields.StringArrToUUIDArr(ids)

	md := metadata.New(map[string]string{
		"sleuth_code": fmt.Sprintf("%d", time.Now().Unix()),
		"merchant_id": "1d6fac48-77df-4395-8a88-e1ec425baffe",
		"staff_id":    uuid.NewV4().String(),
		"branch_ids":  arr.ToMetadataString(),
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	client := newClient()

	res, err := client.client.OfficialGetAuthURL(ctx, &proto.OfficialAuthURLRequest{
		Channel: "appointment",
		Url:     "http://test.wx.haochang.tv/api/v1/branch-map",
		Scope:   "snsapi_base",
		State:   "state",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
	fmt.Println(*res.Data)
}
