package test

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Test_UpdatePointRule 修改积分规则
func Test_UpdatePointRule(t *testing.T) {
	fmt.Println("begin Test_CreatePointRule ===========")
	client := newClient()

	var gainRules []*proto.PointRuleDetail
	gainRules = append(gainRules, &proto.PointRuleDetail{
		CategoryId: "e086a836-2d7a-4762-a887-07910c29ebac",
		Point:      400,
		Fee:        200,
	})

	var useRules []*proto.PointRuleDetail
	useRules = append(useRules, &proto.PointRuleDetail{
		CategoryId: "e086a836-2d7a-4762-a887-07910c29ebac",
		Point:      1,
		Fee:        100,
	})

	res, err := client.UpdatePointRule(newContext(), &proto.UpdatePointRuleRequest{
		Id:          "39d3771d-aeb1-4000-abe9-8a5e300a83c9",
		RuleName:    "直营店积分规则",
		GainRules:   gainRules,
		UseRules:    useRules,
		ValidityDay: 50, // 有效期(天)  永久有效传0
		BranchIds:   []string{"1d6fac48-77df-4395-8a88-e1ec425baffe", "1d6fac48-77df-4395-8a88-e1ec425baff2"},
		Status:      "opened",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)

}
