package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TableAppointmentDateCounter 表结构
type TableAppointmentDateCounter struct {
	ID              uuid.UUID  `gorm:"column:id"`
	BranchID        *uuid.UUID `gorm:"column:branch_id"`
	MerchantID      *uuid.UUID `gorm:"column:merchant_id"`
	RoomGroupID     *uuid.UUID `gorm:"column:room_group_id"`
	Way             int8       `gorm:"column:way"`
	AppointmentDate *time.Time `gorm:"column:appointment_date"`
	AppointmentTime *time.Time `gorm:"column:appointment_time"`
	Number          int32      `gorm:"column:number"`
	AppointNum      int32      `gorm:"appoint_num"`
}

// TableName 表名
func (TableAppointmentDateCounter) TableName() string {
	return "appointment.appointment_date_counter"
}
