package test

import (
	"context"
	"testing"

	"gitlab.omytech.com.cn/micro-service/appointment/proto"
)

func Test_GetAppointmentThemesByRoomType(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetAppointmentThemesByRoomType(ctx, &proto.GetAppointmentThemesByRoomTypeRequest{
		RoomTypeId: "ed4429de-2c46-43e5-8c8e-da9f1bab53bb",
		BranchId:   "ce3435a0-62c9-4355-accd-a224013ed280",
		Date:       1637856000,
		ThemeId:    "3491113a-a794-439b-831d-d14f67c5bed1",
	})
	t.Logf("%+v, %v", resp, err)
}
