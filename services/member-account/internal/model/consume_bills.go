package model

import (
	"time"

	"gorm.io/gorm"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// TableConsumeBills 消费账单明细
type TableConsumeBills struct {
	ID            uuid.UUID    `gorm:"column:id"`
	BaseValue     int32        `gorm:"column:base_value"`
	GiftValue     int32        `gorm:"column:gift_value"`
	Products      *pkgs.Params `gorm:"column:products"`
	Packages      *pkgs.Params `gorm:"column:packages"`
	Tickets       *pkgs.Params `gorm:"column:tickets"`
	BillID        *uuid.UUID   `gorm:"column:bill_id"`
	ConsumeBillID *uuid.UUID   `gorm:"column:consume_bill_id"`
	CreatedAt     *time.Time   `gorm:"column:created_at"`
	UpdatedAt     *time.Time   `gorm:"column:updated_at"`
}

//TableName 表名
func (TableConsumeBills) TableName() string {
	return "member_account.consume_bills"
}

type ConsumeAmount struct {
	ConsumeValue int32
	Products     *pkgs.Params
	Packages     *pkgs.Params
	Tickets      *pkgs.Params
}

//Consume 账户消费
func Consume(tx *gorm.DB, updateBillsMap map[uuid.UUID]map[string]interface{}, createBills4Consume []*TableAccountBill, createConsumeBillsMap map[uuid.UUID]*TableConsumeBills, productNoLeftBillIDs []uuid.UUID, updateProductBills []map[string]interface{}, updateAccounts []map[string]interface{}) error {
	//充值流水更新
	for _, updateBill := range updateBillsMap {
		if err := tx.Model(&TableAccountBill{}).Where("id = ?", updateBill["id"]).Updates(updateBill).Error; err != nil {
			return err
		}
	}
	//流水add消费记录
	for _, createBill := range createBills4Consume {
		if err := tx.Omit("account").Create(createBill).Error; err != nil {
			return err
		}
	}
	//消费流水add
	for _, createConsumeBill := range createConsumeBillsMap {
		if err := tx.Create(createConsumeBill).Error; err != nil {
			return err
		}
	}
	//更新库存
	if len(productNoLeftBillIDs) > 0 {
		if err := tx.Model(&TableProductPackage{}).
			Where("id in (?)", productNoLeftBillIDs).
			UpdateColumn("left", 0).Error; err != nil {
			return err
		}
	}
	//商品流水更新余额
	for _, bill := range updateProductBills {
		//扣了之后还有余额的流水
		if err := tx.Model(&TableProductPackage{}).
			Updates(bill).Error; err != nil {
			return err
		}
	}

	//账户余额更新
	for _, updateAccount := range updateAccounts {
		if err := tx.Model(&TableCardAccount{}).Where("id = ?", updateAccount["id"]).Updates(updateAccount).Error; err != nil {
			return err
		}
	}
	return nil
}
