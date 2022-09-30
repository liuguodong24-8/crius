package test

// // Test_CreateWechatTemplate 创建微信模版
// func Test_CreateWechatTemplate(t *testing.T) {
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

// 	res, err := client.CreateWechatTemplate(ctx, template())
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Println(res)

// }

// func template() message.CreateWechatTemplateRequest {
// 	var detail []*message.WechatTemplateContentDetail
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword1",
// 		Value: "[门店]",
// 		Color: "#d3416c",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword2",
// 		Value: "[房型]",
// 		Color: "#eeffdd",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword3",
// 		Value: "[到店日期][星期]",
// 		Color: "#eeffdd",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword4",
// 		Value: "人数：[预约人数]",
// 		Color: "#eeffdd",
// 	})
// 	detail = append(detail, &message.WechatTemplateContentDetail{
// 		Name:  "keyword5",
// 		Value: "[到店时间][门店电话]",
// 		Color: "#eeffdd",
// 	})
// 	return message.CreateWechatTemplateRequest{
// 		TemplateName: "普通预约退款成功",
// 		TemplateCode: "CxMz08dvyn9jcdbxkBhxIR6QTrOA47AzhTbeGjutvtY",
// 		Category:     "普通预约成功已付",
// 		CategoryKey:  "appointment.deposit_payed",
// 		Content: &message.WechatTemplateContent{
// 			First: &message.WechatTemplateContentBase{
// 				Value: "亲，支付成功，欢迎你的光临",
// 				Color: "#000000",
// 			},
// 			Detail: detail,
// 			Remark: &message.WechatTemplateContentBase{
// 				Value: "点击详情可以跳转到“我的订单”页面。",
// 				Color: "#0de3f4",
// 			},
// 		},
// 		OfficialLink:    "/index/appointed",
// 		MiniprogramLink: "",
// 	}
// }
