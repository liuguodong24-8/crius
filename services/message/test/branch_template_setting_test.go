package test

// // Test_BranchTemplateSetting 门店模版设置
// func Test_BranchTemplateSetting(t *testing.T) {
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

// 	res, err := client.GetBranchTemplate(ctx, message.GetBranchTemplateRequest{
// 		BranchID:    "8fef5e49-8a97-4eca-a428-e0beeff34d00",
// 		MessageType: "merchant.birthday",
// 		Time:        0,
// 	})

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Println(*res.Data)
// 	fmt.Println("短信模版")
// 	fmt.Println(*res.Data.SmsTemplate)
// 	fmt.Println("微信模版")
// 	fmt.Println(*res.Data.WechatTemplate)
// }
