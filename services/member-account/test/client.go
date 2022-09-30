package test

import (
	"context"
	"fmt"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/member-account/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type client struct {
	client proto.MemberAccountServerClient
}

func newClient() *client {
	conn, err := grpc.Dial("127.0.0.1:36512", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("new client err:%s", err.Error()))
	}

	return &client{
		client: proto.NewMemberAccountServerClient(conn),
	}
}

func newContext() context.Context {
	ids := []string{"1d6fac48-77df-4395-8a88-e1ec425baffe", "1d6fac48-77df-4395-8a88-e1ec425baff2"}
	arr, _ := fields.StringArrToUUIDArr(ids)
	md := metadata.New(map[string]string{
		"sleuth_code": fmt.Sprintf("%d", time.Now().Unix()),
		"merchant_id": "1d6fac48-77df-4395-8a88-e1ec425baffe",
		"staff_id":    "4b26a6e9-94a6-4876-810c-7c6ef408922f",
		"branch_ids":  arr.ToMetadataString(),
	})
	return metadata.NewOutgoingContext(context.Background(), md)
}
