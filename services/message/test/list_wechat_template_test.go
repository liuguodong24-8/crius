package test

// // Test_ListWechatTemplate 微信模版列表
// func Test_ListWechatTemplate(t *testing.T) {
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

// 	res, err := client.ListWechatTemplate(ctx, message.ListWechatTemplateRequest{
// 		TemplateName: "",
// 		Status:       "opened",
// 		Limit:        10,
// 		Offset:       0,
// 		OrderBy:      "created_at desc",
// 		WithPage:     true,
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Println(res)

// 	for _, v := range res.Data {
// 		fmt.Println(*v)
// 		fmt.Println(v.OfficialLink)
// 		fmt.Println(v.OfficialLinkName)
// 	}
// }
