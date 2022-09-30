package test

// // Test_UpdateWechatTemplate 修改微信模版
// func Test_UpdateWechatTemplate(t *testing.T) {
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

// 	var detail []*message.WechatTemplateContentDetail
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "thing1",
// 		Value: "燃脂团操",
// 		Color: "#d3416c",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "time2",
// 		Value: "开始时间",
// 		Color: "#eeffdd",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "thing4",
// 		Value: "地点",
// 		Color: "#aabbcc",
// 	})

// 	res, err := client.UpdateWechatTemplate(ctx, message.UpdateWechatTemplateRequest{
// 		ID:           "963649e2-35ba-41af-b79a-ac619f17a1cc",
// 		TemplateName: "预约提醒通知1",
// 		TemplateCode: "0ipsMGVyKuNt4AdU9jpUgYqQ4RjuCm9MP7odJGmHCkE",
// 		Content: &message.WechatTemplateContent{
// 			First: &message.WechatTemplateContentBase{
// 				Value: "预约提醒通知",
// 				Color: "#000000",
// 			},
// 			Detail: nil,
// 			Remark: &message.WechatTemplateContentBase{
// 				Value: "欢迎欢迎",
// 				Color: "#0de3f4",
// 			},
// 		},
// 		OfficialLink:    "/index/cancel",
// 		MiniprogramLink: "",
// 	})

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Println(res)
// }
