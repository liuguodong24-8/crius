package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// TableCardTransfer 划账表结构
type TableCardTransfer struct {
	ID              uuid.UUID    `gorm:"column:id"`
	SourceAccountID *uuid.UUID   `gorm:"column:source_account_id"`
	DestAccountID   *uuid.UUID   `gorm:"column:dest_account_id"`
	TransferValue   int32        `gorm:"column:transfer_value"`
	StaffID         *uuid.UUID   `gorm:"column:staff_id"`
	Extra           *pkgs.Params `gorm:"column:extra"`
	MerchantID      *uuid.UUID   `gorm:"column:merchant_id"`
	CreatedAt       *time.Time   `gorm:"column:created_at"`
	UpdatedAt       *time.Time   `gorm:"column:updated_at"`
}

// TableName 表名
func (TableCardTransfer) TableName() string {
	return "member_account.card_transfer"
}
