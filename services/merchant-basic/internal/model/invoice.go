package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TableInvoice 小票表
type TableInvoice struct {
	ID         uuid.UUID  `gorm:"column:id"`
	Action     string     `gorm:"column:action"`
	Data       string     `gorm:"column:data"`
	MerchantID *uuid.UUID `gorm:"column:merchant_id"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
}

// TableName 表名
func (TableInvoice) TableName() string {
	return "merchant_basic.invoices"
}

const (
	// InvoiceActionOpenCard 开卡
	InvoiceActionOpenCard = "open_card"

	// InvoiceActionRecharge 充值
	InvoiceActionRecharge = "recharge"
)

// CreateInvoice 保存小票
func CreateInvoice(t *TableInvoice) error {
	return entity.Conn.Create(t).Error
}
