package test

// // Test_SendWechatTemplate 发送微信模版消息
// func Test_SendWechatTemplate(t *testing.T) {
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

// 	res, err := client.SendWechatTemplate(ctx, message.SendWechatTemplateRequest{
// 		TemplateID:  "a2af69ba-120f-4b9d-8bd2-0058770a0a09",
// 		Content:     templateContent(),
// 		BranchID:    "39f9ee0c-e172-465c-a2c4-95b6e41ba16e",
// 		MessageType: "appointment.deposit_payed",
// 		WechatUser: &message.SendWechatTemplateRequestWechatUser{
// 			MemberID:       "cff20fee-64ea-47ad-83b0-7b0fc1416f40",
// 			MemberWechatID: "",
// 			MemberOpenID:   "oGPx80rLChgd-2vcSg__CLTdgq80",
// 		},
// 		OfficialLink: "https://www.baidu.com",
// 		Miniprogram: &message.SendWechatTemplateRequestMiniProgram{
// 			Appid:    "",
// 			Pagepath: "",
// 		},
// 		Channel: "minik",
// 		System:  "appointment",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Println(res)

// }

// func templateContent() *message.WechatTemplateContent {
// 	var detail []*message.WechatTemplateContentDetail
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword1",
// 		Value: "江北店",
// 		Color: "#d3416c",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword2",
// 		Value: "豪中",
// 		Color: "#eeffdd",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword3",
// 		Value: "2021-07-07 周三",
// 		Color: "#eeffdd",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword4",
// 		Value: "人数：10人",
// 		Color: "#eeffdd",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword5",
// 		Value: "2021-07-08 12:00 13800138000",
// 		Color: "#eeffdd",
// 	})
// 	return &message.WechatTemplateContent{
// 		First: &message.WechatTemplateContentBase{
// 			Value: "亲，支付成功，欢迎你的光临",
// 			Color: "#000000",
// 		},
// 		Detail: detail,
// 		Remark: &message.WechatTemplateContentBase{
// 			Value: "点击详情可以跳转到“我的订单”页面。",
// 			Color: "#0de3f4",
// 		},
// 	}
// }
