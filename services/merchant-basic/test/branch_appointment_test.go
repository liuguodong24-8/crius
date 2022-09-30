package test

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"merchant_id": "1d6fac48-77df-4395-8a88-e1ec425baffe",
		"staff_id":    uuid.NewV4().String(),
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func Test_SaveBranchAppointment(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.SaveBranchAppointment(ctx, &proto.SaveBranchAppointmentRequest{
		BranchId:               "a44dea1a-6b24-4383-9a79-5df3e2f015e5",
		OpenAppointment:        true,
		AppointmentGranularity: 45,
	})
	t.Log(resp, err)
}

func Test_ShowBranchAppointment(t *testing.T) {
	ctx := newContext()
	c := newClient()
	resp, err := c.ShowBranchAppointment(ctx, &proto.ShowBranchAppointmentRequest{
		BranchId: "25284f5a-0d88-454a-adea-7ea115717e87",
	})
	t.Logf("%+v, %v", resp, err)
}

func Test_UpdateBranchAppointmentRoomType(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.UpdateBranchAppointmentRoomType(ctx, &proto.UpdateBranchAppointmentRoomTypeRequest{
		BranchId: "a44dea1a-6b24-4383-9a79-5df3e2f015e5",
	})
	t.Log(resp, err)
}

func Test_ShowBranchAppointmentBusinessTime(t *testing.T) {
	c := newClient()
	resp, err := c.ShowBranchAppointmentBusinessTime(context.Background(), &proto.ShowBranchAppointmentBusinessTimeRequest{
		BranchId: "1d517206-b7a8-459b-b02c-eca4e972cb83",
		Date:     1628438400,
	})
	t.Logf("%+v, %v", resp, err)
}
