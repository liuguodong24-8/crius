package model

import (
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// TableCardCancel 注销卡表
type TableCardCancel struct {
	ID             uuid.UUID        `gorm:"column:id"`
	CardID         *uuid.UUID       `gorm:"column:card_id"`
	AccountID      *fields.UUIDArr  `gorm:"column:account_id"`
	BankAccount    string           `gorm:"column:bank_account"`
	BankName       string           `gorm:"column:bank_name"`
	MoneyReceiver  string           `gorm:"column:money_receiver"`
	Reason         string           `gorm:"column:reason"`
	ApplyStaffID   *uuid.UUID       `gorm:"column:apply_staff_id"`
	ApplyAt        *time.Time       `gorm:"column:apply_at"`
	Status         CardCancelStatus `gorm:"column:status"`
	ExamineStaffID *uuid.UUID       `gorm:"column:examine_staff_id"`
	ExamineAt      *time.Time       `gorm:"column:examine_at"`
	RefundValue    int32            `gorm:"column:refund_value"`
	RejectReason   string           `gorm:"column:reject_reason"`
	MerchantID     *uuid.UUID       `gorm:"column:merchant_id"`
	CreatedAt      *time.Time       `gorm:"column:created_at"`
	UpdatedAt      *time.Time       `gorm:"column:updated_at"`
}

// TableName 表名
func (TableCardCancel) TableName() string {
	return "member_account.card_cancel"
}

// CardCancelStatus ...
type CardCancelStatus string

const (
	// CardCancelStatusApply 申请
	CardCancelStatusApply = "apply"
	// CardCancelStatusRefund 已退款
	CardCancelStatusRefund = "refund"
	// CardCancelStatusReject 已驳回
	CardCancelStatusReject = "reject"
)

// CardCancel 卡注销
func CardCancel(cancelCard TableCardCancel) error {
	return entity.Conn.Transaction(func(tx *gorm.DB) error {
		//卡状态更新
		if err := tx.Model(&TableCard{}).Scopes(idCondition(*cancelCard.CardID)).Update("status", CardStatusCancelling).Error; err != nil {
			return err
		}
		//注销表add
		return entity.Conn.Create(&cancelCard).Error
	})

}
