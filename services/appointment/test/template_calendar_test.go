package test

import (
	"context"
	"testing"

	"gitlab.omytech.com.cn/micro-service/appointment/proto"
)

func Test_UpdateTemplateCalendar(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateTemplateCalendar(ctx, &proto.UpdateTemplateCalendarRequest{
		BranchId:  "ce3435a0-62c9-4355-accd-a224013ed280",
		BeginDate: 1635091200,
		EndDate:   1635609600,
		Settings: []*proto.UpdateTemplateCalendarRequest_CalendarSetting{{
			TemplateId: "239b8133-674e-4baa-8ab7-d632aae83b25",
			Weeks:      []int32{1, 2, 3, 4, 5, 6, 0},
			Category:   "holiday",
			ThemeIds:   []string{"4ba7e878-de42-4ccb-aa95-518ea5302d63", "2bab9321-06e5-4bb9-a2e5-fb5ff49bb4ea"},
		}},
	})
	t.Logf("%+v, %v", resp, err)
}
