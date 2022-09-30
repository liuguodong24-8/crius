package model

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// TableAppointmentTemplateConfig 模板配置表结构
type TableAppointmentTemplateConfig struct {
	ID         uuid.UUID       `gorm:"column:id"`
	RoomTypeID *uuid.UUID      `gorm:"column:room_type_id"`
	TemplateID *uuid.UUID      `gorm:"column:template_id"`
	AdvanceDay int16           `gorm:"column:advance_day"`
	DepositFee int32           `gorm:"column:deposit_fee"`
	Configure  *pkgs.ParamsArr `gorm:"column:configure"`
}

// AppointmentTemplateConfig 快速预约模板配置json展开表结构
type AppointmentTemplateConfig struct {
	ID         uuid.UUID    `gorm:"column:id"`
	RoomTypeID *uuid.UUID   `gorm:"column:room_type_id"`
	TemplateID *uuid.UUID   `gorm:"column:template_id"`
	AdvanceDay int16        `gorm:"column:advance_day"`
	DepositFee int32        `gorm:"column:deposit_fee"`
	Value      *pkgs.Params `gorm:"column:value"`
}

// TableName 模板配置表名
func (TableAppointmentTemplateConfig) TableName() string {
	return "appointment.appointment_template_configure"
}
