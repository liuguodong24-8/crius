package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
)

// TableAppointmentTemplate 预约模板表结构
type TableAppointmentTemplate struct {
	ID          uuid.UUID         `gorm:"column:id"`
	MerchantID  *uuid.UUID        `gorm:"column:merchant_id"`
	BranchID    *uuid.UUID        `gorm:"column:branch_id"`
	Name        string            `gorm:"column:name"`
	Color       string            `gorm:"column:color"`
	Status      string            `gorm:"column:status"`
	RoomTypeIDs *fields.UUIDArr   `gorm:"column:room_type_ids"`
	CreatedAt   *time.Time        `gorm:"column:created_at"`
	UpdatedAt   *time.Time        `gorm:"column:updated_at"`
	BeginTime   *fields.LocalTime `gorm:"column:begin_time"`
	EndTime     *fields.LocalTime `gorm:"column:end_time"`
	IsNextDay   int8              `gorm:"column:is_next_day"`
}

// TableName 预约模板表名
func (TableAppointmentTemplate) TableName() string {
	return "appointment.appointment_template"
}
