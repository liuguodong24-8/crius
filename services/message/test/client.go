package test

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/message/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Client 实例化对象
type client struct {
	client proto.MessageServiceClient
}

// newClient 实例化
func newClient() *client {
	conn, err := grpc.Dial("127.0.0.1:23456", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("new client err:%s", err.Error()))
	}

	return &client{
		client: proto.NewMessageServiceClient(conn),
	}
}

func newContext() context.Context {
	ids := []string{"1d6fac48-77df-4395-8a88-e1ec425baffe", "1d6fac48-77df-4395-8a88-e1ec425baff2"}
	arr, _ := fields.StringArrToUUIDArr(ids)

	md := metadata.New(map[string]string{
		"sleuth_code": fmt.Sprintf("%d", time.Now().Unix()),
		"merchant_id": "1d6fac48-77df-4395-8a88-e1ec425baffe",
		"staff_id":    uuid.NewV4().String(),
		"branch_ids":  arr.ToMetadataString(),
	})

	return metadata.NewOutgoingContext(context.Background(), md)
}
