package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TableAccountFreeze 账户冻结表
type TableAccountFreeze struct {
	ID         uuid.UUID  `gorm:"column:id"`
	AccountID  *uuid.UUID `gorm:"column:account_id"`
	Action     string     `gorm:"column:action"`
	Reason     string     `gorm:"column:reason"`
	StaffID    *uuid.UUID `gorm:"column:staff_id"`
	MerchantID *uuid.UUID `gorm:"column:merchant_id"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
}

// TableName 表名
func (TableAccountFreeze) TableName() string {
	return "member_account.account_freeze"
}
