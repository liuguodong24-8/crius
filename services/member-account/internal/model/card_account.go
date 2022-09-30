package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"

	"gorm.io/gorm"

	"fmt"

	uuid "github.com/satori/go.uuid"
)

// TableCardAccount 卡账户
type TableCardAccount struct {
	ID         uuid.UUID                       `json:"id" gorm:"column:id"`
	MemberID   *uuid.UUID                      `json:"member_id" gorm:"column:member_id"`
	BranchID   *uuid.UUID                      `json:"branch_id" gorm:"column:branch_id"`
	BaseValue  int32                           `json:"base_value" gorm:"column:base_value"`
	GiftValue  int32                           `json:"gift_value" gorm:"column:gift_value"`
	Products   *fields.ProductPackageTicketArr `json:"products" gorm:"column:products"`
	Packages   *fields.ProductPackageTicketArr `json:"packages" gorm:"column:packages"`
	TagID      *uuid.UUID                      `json:"tag_id" gorm:"column:tag_id"`
	Extra      *pkgs.Params                    `json:"extra" gorm:"column:extra"`
	CreatedAt  *time.Time                      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  *time.Time                      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt                  `json:"deleted_at" gorm:"column:deleted_at"`
	Status     AccountStatus                   `json:"status" gorm:"column:status"`
	MerchantID uuid.UUID                       `json:"merchant_id" gorm:"column:merchant_id"`
	Category   CardSubCategory                 `json:"category" gorm:"column:category"`
	FreezeInfo TableAccountFreeze              `gorm:"foreignkey:account_id"`
}

// AccountStatus 账户状态
type AccountStatus string

const (
	// AccountActionFreeze 冻结卡
	AccountActionFreeze = "freeze"
	// AccountActionUnfreeze 解冻卡
	AccountActionUnfreeze = "unfreeze"

	// AccountStatusActivated 账户状态激活
	AccountStatusActivated AccountStatus = "active"
	// AccountStatusFrozen 账户状态冻结
	AccountStatusFrozen AccountStatus = "frozen"
	// AccountStatusCancelled 账户状态注销
	AccountStatusCancelled AccountStatus = "cancelled"
)

// TableName 表名
func (TableCardAccount) TableName() string {
	return "member_account.card_account"
}

// Value 返回数据库可识别类型
func (p TableCardAccount) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan ...
func (p *TableCardAccount) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("util params value")
	}

	if err := json.Unmarshal(bytes, p); nil != err {
		return errors.New("util params value error")
	}

	return nil
}

// CreateCardAccount 创建
func CreateCardAccount(cardAccount TableCardAccount) error {
	return entity.Conn.Create(&cardAccount).Error
}

// ShowCardAccount 查看
func ShowCardAccount(id uuid.UUID) (*TableCardAccount, error) {
	cardAccount := new(TableCardAccount)
	err := entity.Conn.Scopes(idCondition(id)).Take(cardAccount).Error
	return cardAccount, err
}

// ShowCardWithFreeze 带冻结/解冻处理结果的AccountInfo
func ShowCardWithFreeze(accountID uuid.UUID) (*TableCardAccount, error) {
	account := &TableCardAccount{
		ID: accountID,
	}
	err := entity.Conn.Preload("FreezeInfo").First(&account).Error
	return account, err
}

// GetCardAccountsByCardID 根据卡id查询账户信息
func GetCardAccountsByCardID(id uuid.UUID) ([]TableCardAccount, error) {
	accounts := make([]TableCardAccount, 0)
	err := entity.Conn.Model(&TableCardAccount{}).Joins("inner join member_account.cards on card_account.id = ANY(cards.account_ids)").
		Where("cards.id = ?", id).Find(&accounts).Error
	return accounts, err
}

// GetCardAccountsByIDs 批量查询账户
func GetCardAccountsByIDs(ids []uuid.UUID) ([]TableCardAccount, error) {
	accounts := make([]TableCardAccount, 0)
	err := entity.Conn.Where("id in (?)", ids).Find(&accounts).Error
	return accounts, err
}

// CountAccounts 计算条数
func CountAccounts(status string, branchIDs []uuid.UUID, merchantID uuid.UUID) (int64, error) {
	var count int64
	err := entity.Conn.Model(&TableCardAccount{}).Scopes(
		accountBranchesCondition(branchIDs),
		accountStatusCondition(status),
		merchantIDCondition(merchantID)).Count(&count).Error
	return count, err
}

// GetAccounts 获取账户列表
func GetAccounts(status string, branchIDs []uuid.UUID, offset, limit int32, merchantID uuid.UUID) ([]TableCardAccount, error) {
	accounts := make([]TableCardAccount, 0)
	db := entity.Conn.Distinct().Scopes(
		accountBranchesCondition(branchIDs),
		pagingCondition(offset, limit),
		accountStatusCondition(status),
		merchantIDCondition(merchantID))

	err := db.
		Order("created_at DESC").
		Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// UpdateAccountStatusByID 更新账户状态
func UpdateAccountStatusByID(account *TableCardAccount, accountFreeze *TableAccountFreeze) error {
	return entity.Conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(account).UpdateColumn("status", account.Status).Error; err != nil {
			return err
		}

		return tx.Create(accountFreeze).Error
	})

}

// GetAccountByMemberID 根据会员id查询账户列表
func GetAccountByMemberID(memberID uuid.UUID) ([]TableCardAccount, error) {
	var account []TableCardAccount
	err := entity.Conn.Distinct().Scopes(
		memberIDCondition(memberID),
	).
		Order("created_at desc").
		Find(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

// GetAccountByCardCode 根据卡号查询账户
func GetAccountByCardCode(cardCode string, merchantID uuid.UUID) (*TableCardAccount, error) {
	var account TableCardAccount
	err := entity.Conn.Model(&TableCard{}).
		Select("card_account.*").
		Joins("right join member_account.card_account on card_account.id = cards.account_id").
		Where("cards.code = ?", cardCode).
		Where("card_account.merchant_id = ?", merchantID).
		Order("card_account.created_at desc").
		Take(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetAccountByCardID 根据卡id查询账户列表
func GetAccountByCardID(cardID uuid.UUID) ([]TableCardAccount, error) {
	var account []TableCardAccount
	err := entity.Conn.
		Joins("left join member_account.cards on cards.account_id = card_account.id").
		Where("cards.id = ?", cardID).Find(&account).Error
	return account, err
}

func nameLike(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%%s%%", name))
		}
		return db
	}
}

func phoneLike(phone string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if phone != "" {
			db = db.Where("phone like ?", fmt.Sprintf("%%%s%%", phone))
		}
		return db
	}
}

func accountBranchesCondition(branchIDs []uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(branchIDs) > 0 {
			db = db.Where("card_account.branch_id in (?)", branchIDs)
		}
		return db
	}
}

func accountStatusCondition(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			db = db.Where("card_account.status = ?", status)
		}
		return db
	}
}

func memberIDCondition(memberID uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if memberID != uuid.Nil {
			db = db.Where("card_account.member_id = ?", memberID)
		}
		return db
	}
}

func memberIDInCondition(memberIDs []uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(memberIDs) != 0 {
			db = db.Where("card_account.member_id in (?)", memberIDs)
		}
		return db
	}
}
