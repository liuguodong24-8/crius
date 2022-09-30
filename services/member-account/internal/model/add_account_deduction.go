package model

import (
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// TableAddAccountDeduction 追加扣款表
type TableAddAccountDeduction struct {
	ID         uuid.UUID       `gorm:"column:id"`
	ConsumeIDs *fields.UUIDArr `gorm:"column:consume_ids"`
	BillNumber string          `gorm:"column:bill_number"`
	BranchID   *uuid.UUID      `gorm:"column:branch_id"`
	StaffID    *uuid.UUID      `gorm:"column:staff_id"`
	MerchantID *uuid.UUID      `gorm:"column:merchant_id"`
	Reason     string          `gorm:"column:reason"`
	CreatedAt  *time.Time      `gorm:"column:created_at"`
	UpdatedAt  *time.Time      `gorm:"column:updated_at"`
}

// TableName 表
func (TableAddAccountDeduction) TableName() string {
	return "member_account.add_account_deduction"
}

// AddAccountDeduction 追加扣款
func AddAccountDeduction(tx *gorm.DB, accountDeduction *TableAddAccountDeduction) error {
	//记录表insert
	return tx.Create(accountDeduction).Error
}
