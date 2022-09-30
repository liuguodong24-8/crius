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

// Test_OfficialCodeGetUser 公众号code获取user信息
func Test_OfficialCodeGetUser(t *testing.T) {
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

	res, err := client.client.OfficialCodeGetUser(ctx, &proto.OfficialCodeGetUserRequest{
		Channel: "appointment",
		Code:    "0011ZiGa1PP2sB0gK9Ha1gqKga01ZiGT",
	})

	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
	fmt.Println(res.ErrorCode)
	fmt.Println(res.ErrorMessage)
	fmt.Println(res.Data)
}
