package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gorm.io/gorm"
)

// TableCardReplace 补卡表
type TableCardReplace struct {
	ID         uuid.UUID      `gorm:"column:id"`
	CurrCardID *uuid.UUID     `gorm:"column:curr_card_id"`
	NewCardID  *uuid.UUID     `gorm:"column:new_card_id"`
	StaffID    *uuid.UUID     `gorm:"column:staff_id"`
	Payments   *pkgs.Params   `gorm:"column:payments"`
	Extra      *pkgs.Params   `gorm:"column:extra"`
	MerchantID *uuid.UUID     `gorm:"column:merchant_id"`
	CreatedAt  *time.Time     `gorm:"column:created_at"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}

// TableName 表名
func (TableCardReplace) TableName() string {
	return "member_account.card_replace"
}

// ReplaceCard 补卡
func ReplaceCard(tx *gorm.DB, currCard, newCard *TableCard, cardReplace *TableCardReplace) error {
	//当前卡注销
	if err := tx.Model(currCard).Updates(*currCard).Error; err != nil {
		return err
	}
	//更新新卡状态和密码
	if err := tx.Model(newCard).Updates(*newCard).Error; err != nil {
		return err
	}
	//补卡表insert
	return tx.Create(cardReplace).Error
}
