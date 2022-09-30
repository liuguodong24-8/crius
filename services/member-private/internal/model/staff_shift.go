package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// StaffShift 员工交班信息
type StaffShift struct {
	ID         uuid.UUID `json:"id" gorm:"column:id"`
	MerchantID uuid.UUID `json:"merchant_id" gorm:"column:merchant_id"`
	BranchID   uuid.UUID `json:"branch_id" gorm:"column:branch_id"`
	StaffID    uuid.UUID `json:"staff_id" gorm:"column:staff_id"`
	BeginTime  time.Time `json:"begin_time" gorm:"column:begin_time"`
	EndTime    time.Time `json:"end_time" gorm:"column:end_time"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
}

// TableName table name
func (s StaffShift) TableName() string {
	return `member_private.staff_shifts`
}
