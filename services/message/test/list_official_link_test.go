package test

// // Test_ListOfficialLink 公众号链接
// func Test_ListOfficialLink(t *testing.T) {
// 	fmt.Println("begin Test_ListOfficialLink")
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

// 	res, err := client.ListOfficialLink(ctx)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Println(res)

// }
