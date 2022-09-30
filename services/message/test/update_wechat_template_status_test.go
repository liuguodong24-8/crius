package test

// // Test_UpdateWechatTemplateStatus 修改微信模版状态
// func Test_UpdateWechatTemplateStatus(t *testing.T) {
// 	ids := []string{"1d6fac48-77df-4395-8a88-e1ec425baffe", "1d6fac48-77df-4395-8a88-e1ec425baff2"}
// 	arr, _ := fields.StringArrToUUIDArr(ids)

// 	md := metadata.New(map[string]string{
// 		"sleuth_code": fmt.Sprintf("%d", time.Now().Unix()),
// 		"merchant_id": "1d6fac48-77df-4395-8a88-e1ec425baffe",
// 		"staff_id":    uuid.NewV4().String(),
// 		"branch_ids":  arr.ToMetadataString(),
// 	})

// 	ctx := metadata.NewOutgoingContext(context.Background(), md)
// 	client, err := message.NewClient(context.Background(), message.ClientConfig{Address: "127.0.0.1:33456"})
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	res, err := client.UpdateWechatTemplateStatus(ctx, message.UpdateWechatTemplateStatusRequest{
// 		ID:     "963649e2-35ba-41af-b79a-ac619f17a1cc",
// 		Status: "closed",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Println(res)
// }
