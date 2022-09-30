package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func newClient() proto.AppointmentServerClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)), grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:34512", opts...)
	if err != nil {
		return nil
	}

	return proto.NewAppointmentServerClient(conn)
}

func newContext() context.Context {
	ids := []string{"1d6fac48-77df-4395-8a88-e1ec425baffe", "1d6fac48-77df-4395-8a88-e1ec425baff2"}
	arr, _ := fields.StringArrToUUIDArr(ids)
	md := metadata.New(map[string]string{
		"sleuth_code": fmt.Sprintf("%d", time.Now().Unix()),
		"merchant_id": "1d6fac48-77df-4395-8a88-e1ec425baffe",
		"staff_id":    "4b26a6e9-94a6-4876-810c-7c6ef408922f",
		"branch_ids":  arr.ToMetadataString(),
	})
	return metadata.NewOutgoingContext(context.Background(), md)
}

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

func Test_GetAppointmentRemaining(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetAppointmentRemaining(ctx, &proto.GetAppointmentRemainingRequest{
		BranchId:     "4d59442d-f7b6-472c-8df4-9beda4d48e0b",
		BusinessDate: 1639929600,
		Way:          1,
	})
	t.Logf("%+v, %v", resp, err)
}

// // SaveAppointmentLock
// func Test_SaveAppointmentLock(t *testing.T) {
// 	ctx := context.Background()
// 	c, err := appointment.NewClient(ctx, appointment.ClientConfig{"127.0.0.1:34512"})
// 	t.Log(err)
// 	resp, err := c.SaveAppointmentLock(ctx, appointment.SaveAppointmentLockRequest{
// 		BranchID:        uuid.FromStringOrNil("1d517206-b7a8-459b-b02c-eca4e972cb83"),
// 		RoomGroupID:     uuid.FromStringOrNil("e3490899-11b6-4f76-839b-fe433b8799ec"),
// 		Way:             2,
// 		AppointmentAt:   time.Unix(1624593600, 0),
// 		AppointmentDate: time.Unix(1624550400, 0),
// 	})
// 	t.Log(resp, err)
// }

// func Test_DeleteAppointmentLock(t *testing.T) {
// 	ctx := context.Background()
// 	c, err := appointment.NewClient(ctx, appointment.ClientConfig{"127.0.0.1:34512"})
// 	t.Log(err)
// 	resp, err := c.DeleteAppointmentLock(ctx, appointment.DeleteAppointmentLockRequest{
// 		ID: uuid.FromStringOrNil("dffb3f89-5b50-4de8-9289-2e55aaf0fc3a"),
// 	})
// 	t.Log(resp, err)
// }

// func Test_GetAppointmentRecord(t *testing.T) {
// 	ctx := context.Background()
// 	c, err := appointment.NewClient(ctx, appointment.ClientConfig{"127.0.0.1:34512"})
// 	t.Log(err)
// 	resp, err := c.GetAppointmentRecord(ctx, appointment.GetAppointmentRecordRequest{
// 		ID: uuid.FromStringOrNil("eaf25d92-a901-4ed8-9e2d-d5b9607818b5"),
// 	})
// 	t.Log(resp, err)
// }
func Test_ShowAppointment(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.ShowAppointment(ctx, &proto.ShowAppointmentRequest{
		Id: "ce03ccdd-74c3-4521-8f59-09f204259d38",
	})
	t.Log(resp, err)
}

func Test_CreateAppointment(t *testing.T) {
	c := newClient()

	resp, err := c.CreateAppointment(newContext(), &proto.CreateAppointmentRequest{
		Appointment: &proto.Appointment{CalledCode: "86",
			CalledPhone:      "15669964451",
			AppointmentCode:  "86",
			AppointmentPhone: "13312345678",
			Name:             "fzt",
			Gender:           1,
			Way:              2,
			CustomerNum:      5,
			BranchId:         "cf396088-00b9-4705-8c93-bd8e20d88ad5",
			AppointmentAt:    1642000500,
			AppointmentDate:  1641916800,
			RoomTypeId:       "a1e7369a-5a10-45f8-9038-065e638f9717",
			Operator:         uuid.NewV4().String(),
			ChargingWay:      1,
			FlowerCake:       true,
			DepositFee:       0,
			FlowerCakeRemark: "有生日蛋糕",
			Remark:           "提前开空调",
			MemberId:         "910c695d-ea8d-437c-afce-24655e0d4c88",
		},
		LockId:         "743b8a2d-ba56-4b04-a801-9f4f338cf642",
		ThemePackageId: "2073973a-b247-4059-b48e-ba1751ee7ed5",
	})
	t.Log(resp, err)
}

func Test_GetAppointments(t *testing.T) {
	ctx := context.Background()
	c := newClient()
	resp, err := c.GetAppointments(ctx, &proto.GetAppointmentsRequest{
		MemberId: "910c695d-ea8d-437c-afce-24655e0d4c88",
		WechatId: "38f9f373-352c-41fd-89e5-477558a2ad00",
		AnyPhone: "18723508414",
	})
	t.Log(resp, err)
}

func Test_GetAppointmentRemainingFast(t *testing.T) {
	c := newClient()
	resp, err := c.GetAppointmentRemainingFast(context.Background(), &proto.GetAppointmentRemainingFastRequest{
		Time: "11:17:00",
		Num:  1,
		Branches: []*proto.GetAppointmentRemainingFastBranches{
			{BranchId: "ec42113f-77ff-4faa-a0be-921b704f2406", Distance: 6854},
			{BranchId: "8fef5e49-8a97-4eca-a428-e0beeff34d00", Distance: 2001},
			{BranchId: "25284f5a-0d88-454a-adea-7ea115717e87", Distance: 9700},
			{BranchId: "1d517206-b7a8-459b-b02c-eca4e972cb83", Distance: 34160},
			{BranchId: "22489bb6-25b2-4875-9d82-f3106da7e26e", Distance: 2590},
			{BranchId: "e0dcccf5-8c4a-4fbd-be14-ec702b0ef04d", Distance: 40794},
			{BranchId: "51dc013d-a48f-498f-9efa-c4173427cebf", Distance: 4371},
			{BranchId: "f30f030a-6062-4599-97e9-2367c93cc3b1", Distance: 1457769},
			{BranchId: "f48f1be5-68be-4e5c-8397-8a516c0edb8c", Distance: 9629},
			{BranchId: "ebc50dac-fc4f-4fd8-ab18-434e45960be1", Distance: 69004},
		},
		Way:          1,
		BusinessDate: 1628179200,
		Limit:        3,
	})
	t.Logf("%+v, %v", resp, err)
}

func Test_UpdateAppointment(t *testing.T) {
	c := newClient()
	resp, err := c.UpdateAppointment(context.Background(), &proto.UpdateAppointmentRequest{
		Appointment:    &proto.Appointment{Id: "6bde5550-f31a-4cce-b249-e160e5e3a6e1", CalledPhone: "18723508414", AppointmentPhone: "18723508414", AppointmentCode: "86", Way: 1, CustomerNum: 1, AppointmentAt: 1642762800, AppointmentDate: 1642694400, Remark: "阿巴阿巴阿巴阿巴", RoomTypeId: "64812382-a4f4-498c-a071-4d8ad3d0d841", ChargingWay: 1, BranchId: "ce3435a0-62c9-4355-accd-a224013ed280", MemberId: "910c695d-ea8d-437c-afce-24655e0d4c88", Name: "走到累趴"},
		ThemePackageId: "bf5a482e-a521-4909-9a22-51ecd9b0cd7f",
	})
	t.Logf("%+v, %v", resp, err)
}

func Test_GetAppointmentLimit(t *testing.T) {
	c := newClient()
	resp, err := c.GetAppointmentLimit(context.Background(), &proto.GetAppointmentLimitRequest{
		Phone:  "13388888888",
		Status: "appointed",
	})
	t.Logf("%+v, %v", resp, err)
}
