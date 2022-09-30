package model

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

const appointmentCodeSeqKey = "appointment_code_seq_%s"

// TableAppointmentExtend 预约扩展表
type TableAppointmentExtend struct {
	AppointmentID  uuid.UUID       `gorm:"column:appointment_id"`
	PackageID      *uuid.UUID      `gorm:"column:package_id"`
	Packages       *pkgs.ParamsArr `gorm:"column:packages"`
	Decoration     string          `gorm:"column:decoration"`
	Staffing       string          `gorm:"column:staffing"`
	CustomConfigs  *pkgs.ParamsArr `gorm:"column:custom_configs"`
	ThemeID        *uuid.UUID      `gorm:"column:theme_id"`
	PackageName    string          `gorm:"column:package_name"`
	ShareMessage   string          `gorm:"column:share_message"`
	RefundingAt    *time.Time      `gorm:"column:refunding_at"`
	RefundedAt     *time.Time      `gorm:"column:refunded_at"`
	RefundAmount   int32           `gorm:"column:refund_amount"`
	RefundID       *uuid.UUID      `gorm:"column:refund_id"`
	RefundResponse *pkgs.Params    `gorm:"column:refund_response"`
	PackagePrice   int32           `gorm:"column:package_price"`
	Code           string          `gorm:"column:code"`
	OpenAt         *time.Time      `gorm:"column:open_at"`
	OpenRoomID     *uuid.UUID      `gorm:"column:open_room_id"`
	Sended         bool            `gorm:"column:sended"`
	LoadExtra      string          `gorm:"-"`
}

// TableName 表名
func (TableAppointmentExtend) TableName() string {
	return "appointment.appointment_extend"
}

// ShowAppointmentCodeSeq 获取预约号序列
func ShowAppointmentCodeSeq() (string, error) {
	now := time.Now().Format("20060102")
	key := fmt.Sprintf(appointmentCodeSeqKey, now)
	code, err := redisEntity.INCR(key)
	if err != nil {
		return "", err
	}
	redisEntity.EXPIRE(key, 24*3600)
	return fmt.Sprintf("%s%d", now, code), nil
}

// ShowAppointmentExtend 查询预约扩展
func ShowAppointmentExtend(id uuid.UUID) (*TableAppointmentExtend, error) {
	extend := new(TableAppointmentExtend)
	err := entity.Conn.Scopes(util.ColumnEqualScope("appointment_id", id)).Take(extend).Error
	return extend, err
}

// UpdateAppointmentShare 更新预约分享设置
func UpdateAppointmentShare(id uuid.UUID, message string) error {
	return entity.Conn.Model(&TableAppointmentExtend{}).Scopes(util.ColumnEqualScope("appointment_id", id)).UpdateColumn("share_message", message).Error
}

// UpdateAppointmentSended 更新是否发送到咨客
func UpdateAppointmentSended(id uuid.UUID, s bool) error {
	return entity.Conn.Model(&TableAppointmentExtend{}).Scopes(util.ColumnEqualScope("appointment_id", id)).UpdateColumn("sended", s).Error
}
