package test

import (
	"context"
	"fmt"
	"time"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func newClient() proto.MerchantBasicServiceClient {
	conn, err := grpc.Dial("127.0.0.1:12346", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("new client err:%s", err.Error()))
	}

	return proto.NewMerchantBasicServiceClient(conn)
}

func newContext() context.Context {
	// ids := []string{"1d6fac48-77df-4395-8a88-e1ec425baffe", "1d6fac48-77df-4395-8a88-e1ec425baff2"}
	// arr, _ := fields.StringArrToUUIDArr(ids)
	md := metadata.New(map[string]string{
		"sleuth_code": fmt.Sprintf("%d", time.Now().Unix()),
		"merchant_id": "1d6fac48-77df-4395-8a88-e1ec425baffe",
		// "staff_id":    uuid.NewV4().String(),
		"staff_id": "f24c596a-da55-43ae-94f8-ed1628a4f2e0",
		// "branch_ids": arr.ToMetadataString(),
	})
	return metadata.NewOutgoingContext(context.Background(), md)
}
