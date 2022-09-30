package category

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_ListConsumeCategory 消费类型列表
func Test_ListConsumeCategory(t *testing.T) {
	fmt.Println("begin Test_ListConsumeCategory ===========")
	client := newClient()

	res, err := client.client.ListConsumeCategory(newContext(), &proto.ListConsumeCategoryRequest{
		Category: "费",
		Status:   "opened",
		Limit:    10,
		Offset:   0,
		OrderBy:  "created_at asc",
		WithPage: false,
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
