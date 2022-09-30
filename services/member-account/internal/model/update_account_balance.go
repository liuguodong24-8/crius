package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gorm.io/gorm"
)

// TableUpdateAccountBill 修改账户表
type TableUpdateAccountBill struct {
	ID                    uuid.UUID           `gorm:"column:id"`
	AccountBillID         *uuid.UUID          `gorm:"column:account_bill_id"`
	OldAccountBillBalance *AccountBalanceBill `gorm:"column:old_account_bill_balance"`
	NewAccountBillBalance *AccountBalanceBill `gorm:"column:new_account_bill_balance"`
	BranchID              *uuid.UUID          `gorm:"column:branch_id"`
	StaffID               *uuid.UUID          `gorm:"column:staff_id"`
	MerchantID            *uuid.UUID          `gorm:"column:merchant_id"`
	Reason                string              `gorm:"column:reason"`
	CreatedAt             *time.Time          `gorm:"column:created_at"`
	UpdatedAt             *time.Time          `gorm:"column:updated_at"`
}

// AccountBalanceBill 账户流水余额
type AccountBalanceBill struct {
	BaseValue int32
	GiftValue int32
	Packages  *pkgs.ParamsArr
	Products  *pkgs.ParamsArr
}

func (p AccountBalanceBill) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan ...
func (p *AccountBalanceBill) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("AccountBalanceBill value error")
	}

	if err := json.Unmarshal(bytes, p); nil != err {
		return errors.New("AccountBalanceBill unmarshal error")
	}

	return nil
}

// TableName 表名
func (TableUpdateAccountBill) TableName() string {
	return "member_account.update_account_bill"
}

// UpdateAccountBalance 修改余额
func UpdateAccountBalance(accountID uuid.UUID, afterAccount map[string]interface{}, bills []*TableAccountBill, updateAccountBalance *TableUpdateAccountBill) error {
	return entity.Conn.Transaction(func(tx *gorm.DB) error {
		// 更新账户balance
		if err := tx.Model(&TableCardAccount{}).Where("id = ?", accountID).Updates(afterAccount).Error; err != nil {
			return err
		}
		//流水表insert
		for _, bill := range bills {
			if err := tx.Save(bill).Error; err != nil {
				return err
			}
		}
		//记录表insert
		return tx.Create(updateAccountBalance).Error
	})
}
