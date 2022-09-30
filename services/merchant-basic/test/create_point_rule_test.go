package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_CreatePointRule 创建积分规则
func Test_CreatePointRule(t *testing.T) {
	fmt.Println("begin Test_CreatePointRule ===========")
	client := newClient()

	var gainRules []*proto.PointRuleDetail
	gainRules = append(gainRules, &proto.PointRuleDetail{
		CategoryId: "c8c9ba69-7aff-4828-883e-590d3e88762c", // 类型名
		Point:      1,                                      // 积分
		Fee:        100,                                    // 金额
	})

	var useRules []*proto.PointRuleDetail
	useRules = append(useRules, &proto.PointRuleDetail{
		CategoryId: "c8c9ba69-7aff-4828-883e-590d3e88762c",
		Point:      100,
		Fee:        1,
	})

	res, err := client.CreatePointRule(newContext(), &proto.CreatePointRuleRequest{
		RuleName:    "固定有效期限",                                         // 规则名
		GainRules:   gainRules,                                        // 获取规则
		UseRules:    useRules,                                         // 抵扣规则
		ValidityDay: 90,                                               // 有效期(天)  永久有效传-1
		BranchIds:   []string{"a3d0ffa8-d6a8-483e-8b11-445f529d3102"}, // 门店
		Status:      "opened",                                         // 状态
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
