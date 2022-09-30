package category

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_CreateConsumeCategory 创建积分类型
func Test_CreateConsumeCategory(t *testing.T) {
	fmt.Println("begin Test_CreatePointCategory ===========")
	client := newClient()

	res, err := client.client.CreateConsumeCategory(newContext(), &proto.CreateConsumeCategoryRequest{
		Category: "房费",     // 类型名
		Code:     "0002",   // 类型编码
		Status:   "opened", // 状态
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
