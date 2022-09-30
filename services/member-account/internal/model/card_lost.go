package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gorm.io/gorm"
)

// TableCardLost 卡挂失表
type TableCardLost struct {
	ID         uuid.UUID      `gorm:"column:id"`
	CardID     *uuid.UUID     `gorm:"column:card_id"`
	StaffID    *uuid.UUID     `gorm:"column:staff_id"`
	Action     string         `gorm:"column:action"`
	Extra      *pkgs.Params   `gorm:"column:extra"`
	MerchantID *uuid.UUID     `gorm:"column:merchant_id"`
	CreatedAt  *time.Time     `gorm:"column:created_at"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}

// TableName 表名
func (TableCardLost) TableName() string {
	return "member_account.card_lost"
}

// UpdateCardStatus 更新卡状态
func UpdateCardStatus(tableCardLost *TableCardLost, cardNextStatus CardStatus) error {
	return entity.Conn.Transaction(func(tx *gorm.DB) error {
		// 更新卡状态
		if err := tx.Model(&TableCard{}).Scopes(idCondition(*tableCardLost.CardID)).Update("status", cardNextStatus).Error; err != nil {
			return err
		}
		//挂失流水表insert
		return tx.Create(tableCardLost).Error
	})
}
