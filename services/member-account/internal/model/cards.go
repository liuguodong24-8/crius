package model

import (
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	cutil "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// ErrBeyondPrimaryAccountBalance 开副卡金额大于主卡余额错误
var ErrBeyondPrimaryAccountBalance = errors.New("金额大于主卡账户余额")

// ErrAccountAbnormal 账户状态异常
var ErrAccountAbnormal = errors.New("账户状态异常")

// TableCard 卡数据库结构
type TableCard struct {
	ID             uuid.UUID       `gorm:"column:id"`
	Category       CardCategory    `gorm:"column:category"`
	SubCategory    CardSubCategory `gorm:"column:sub_category"`
	Code           string          `gorm:"column:code"`
	PrimaryID      *uuid.UUID      `gorm:"column:primary_id"`
	CreateBranchID *uuid.UUID      `gorm:"column:create_branch_id"`
	OpenBranchID   *uuid.UUID      `gorm:"column:open_branch_id"`
	CreateStaffID  *uuid.UUID      `gorm:"column:create_staff_id"`
	AccountIDs     *fields.UUIDArr `gorm:"column:account_ids"`
	MerchantID     *uuid.UUID      `gorm:"column:merchant_id"`
	Status         CardStatus      `gorm:"column:status"`
	OpenedAt       *time.Time      `gorm:"column:opened_at"`
	Password       string          `gorm:"column:password"`
	OpenStaffID    *uuid.UUID      `gorm:"column:open_staff_id"`
	Extra          *pkgs.Params    `gorm:"column:extra"`
	CreatedAt      *time.Time      `gorm:"column:created_at"`
	UpdatedAt      *time.Time      `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`
	MemberID       *uuid.UUID      `gorm:"column:member_id"`
}

// TableName 表名
func (TableCard) TableName() string {
	return "member_account.cards"
}

type (
	// CardStatus 卡状态
	CardStatus string
	// CardCategory 卡类型
	CardCategory string
	// CardSubCategory 卡子类型
	CardSubCategory string
)

const (
	// CardStatusInit 初始化状态卡
	CardStatusInit CardStatus = "init"
	// CardStatusActive 激活状态卡
	CardStatusActive CardStatus = "active"
	// CardStatusLost 挂失状态卡
	CardStatusLost CardStatus = "lost"
	// CardStatusCancelling 注销审核中状态卡
	CardStatusCancelling CardStatus = "cancelling"
	// CardStatusCancelled 注销状态卡
	CardStatusCancelled CardStatus = "cancelled"

	// CardCategoryMember 会员卡
	CardCategoryMember CardCategory = "member"
	// CardCategoryGift 礼品卡
	CardCategoryGift CardCategory = "gift"

	// CardSubCategoryBlank 不记名账户
	CardSubCategoryBlank CardSubCategory = "blank"
	// CardSubCategoryPrimary 主卡账户
	CardSubCategoryPrimary CardSubCategory = "primary"
	// CardSubCategorySecondary 副卡账户
	CardSubCategorySecondary CardSubCategory = "secondary"

	//CardActionLost 挂失
	CardActionLost = "lost"
	// CardActionFind 找回
	CardActionFind = "find"
	// CardActionCancel 注销
	CardActionCancel = "cancel"
	// CardActionCancelExamine 注销审核
	CardActionCancelExamine = "cancel_examine"
)

// CreateCard 创建卡
func CreateCard(card TableCard) error {
	return entity.Conn.Create(&card).Error
}

// ShowCard 根据id查询卡
func ShowCard(id, merchantID uuid.UUID) (*TableCard, error) {
	card := new(TableCard)
	err := entity.Conn.Scopes(idCondition(id), merchantIDCondition(merchantID)).Take(card).Error
	return card, err
}

// ShowCardByCode 根据code查询卡
func ShowCardByCode(code string, merchantID uuid.UUID) (*TableCard, error) {
	card := new(TableCard)
	err := entity.Conn.Scopes(cardCodeCondition(code), merchantIDCondition(merchantID)).Take(card).Error
	return card, err
}

//GetMemberCardsByMemberID 查询用户的所有卡
func GetMemberCardsByMemberID(memberID uuid.UUID) (*[]TableCard, error) {
	var cards []TableCard
	err := entity.Conn.Where("member_id = ?", memberID).Find(&cards).Error
	return &cards, err
}

// ShowPrimaryCardBySecondaryID 根据副卡id查主卡
func ShowPrimaryCardBySecondaryID(id, merchantID uuid.UUID) (*TableCard, error) {
	card := new(TableCard)
	subQuery := entity.Conn.Select("parimary_id").Model(&TableCard{}).Scopes(idCondition(id), merchantIDCondition(merchantID))
	err := entity.Conn.Where("id = ?", subQuery).Take(card).Error
	return card, err
}

// RechargeCard 激活/充值 卡
func RechargeCard(card TableCard, account TableCardAccount, bill TableAccountBill, products, packages []TableProductPackage) error {
	tx := entity.Conn.Begin()
	// createAt为空，则是创建账户
	if account.CreatedAt == nil {
		if err := tx.Create(&account).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		accountMap := make(map[string]interface{})
		accountMap["base_value"] = gorm.Expr("base_value+?", account.BaseValue)
		accountMap["gift_value"] = gorm.Expr("gift_value+?", account.GiftValue)
		if err := tx.Model(&TableCardAccount{}).Scopes(cutil.ColumnEqualScope("id", account.ID)).Updates(accountMap).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Updates(&card).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Omit("account").Create(&bill).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// ActiveSecondaryCard 激活副卡
func ActiveSecondaryCard(tx *gorm.DB, card TableCard, primaryID uuid.UUID, account TableCardAccount, primaryBill, secondaryBill TableAccountBill, consumeBills []TableConsumeBills, bills []TableAccountBill) error {
	err := tx.Updates(&card).Error
	if err != nil {
		return err
	}
	primaryMap := make(map[string]interface{})
	primaryMap["base_value"] = gorm.Expr("base_value-?", account.BaseValue)
	primaryMap["gift_value"] = gorm.Expr("gift_value-?", account.GiftValue)
	if err = tx.Model(&TableCardAccount{}).Scopes(idCondition(primaryID)).Updates(primaryMap).Error; err != nil {
		return err
	}
	if err = tx.Create(&account).Error; err != nil {
		return err
	}
	if err = tx.Omit("account").Create(&primaryBill).Error; err != nil {
		return err
	}
	if err = tx.Omit("account").Create(&secondaryBill).Error; err != nil {
		return err
	}
	if err = tx.Create(&consumeBills).Error; err != nil {
		return err
	}
	for i := 0; i < len(bills); i++ {
		if err := tx.Save(&bills[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

// ActiveSecondaryCardTransaction 激活副卡事务
func ActiveSecondaryCardTransaction(tabCard *TableCard, primaryAccount, cardAccount *TableCardAccount, primaryBill, secondaryBill *TableAccountBill, primaryAccountID uuid.UUID, rechargeValue int32) error {
	rechargeTotal := rechargeValue
	tx := entity.Conn.Begin()
	bills, err := GetRechargeBillsByAccountID(tx, primaryAccountID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("ActiveSecondaryCard 查询充值账户充值流水数据库错误:%v", err))
		tx.Rollback()
		return err
	}
	consumeBillsMap := make(map[uuid.UUID]TableConsumeBills)

	for i := 0; i < len(bills); i++ {
		if bills[i].GiftValueLeft == 0 {
			continue
		}
		if bills[i].GiftValueLeft < rechargeValue {
			consumeBillsMap[bills[i].ID] = TableConsumeBills{
				ID:            uuid.NewV4(),
				GiftValue:     bills[i].GiftValueLeft,
				BillID:        &primaryBill.ID,
				ConsumeBillID: &bills[i].ID,
			}
			rechargeValue -= bills[i].GiftValueLeft
			bills[i].GiftValueLeft = 0
		} else {
			consumeBillsMap[bills[i].ID] = TableConsumeBills{
				ID:            uuid.NewV4(),
				GiftValue:     rechargeValue,
				BillID:        &primaryBill.ID,
				ConsumeBillID: &bills[i].ID,
			}
			bills[i].GiftValueLeft -= rechargeValue
			rechargeValue = 0
			break
		}
	}
	// 写入账单、副卡账户 本金 赠金
	secondaryBill.GiftValue, secondaryBill.GiftValueLeft = rechargeTotal-rechargeValue, rechargeTotal-rechargeValue
	secondaryBill.BaseValue, secondaryBill.BaseValueLeft = rechargeValue, rechargeValue
	cardAccount.BaseValue, cardAccount.GiftValue = secondaryBill.BaseValue, secondaryBill.GiftValue
	primaryBill.BaseValue, primaryBill.GiftValue = secondaryBill.BaseValue, secondaryBill.GiftValue
	primaryAccount.BaseValue, primaryAccount.GiftValue = primaryAccount.BaseValue-secondaryBill.BaseValue, primaryAccount.GiftValue-secondaryBill.GiftValue

	if rechargeValue != 0 {
		for i := 0; i < len(bills); i++ {
			if bills[i].BaseValueLeft == 0 {
				continue
			}
			if bills[i].BaseValueLeft < rechargeValue {
				if consumeBill, ok := consumeBillsMap[bills[i].ID]; !ok {
					consumeBillsMap[bills[i].ID] = TableConsumeBills{
						ID:            uuid.NewV4(),
						BaseValue:     bills[i].BaseValueLeft,
						BillID:        &primaryBill.ID,
						ConsumeBillID: &bills[i].ID,
					}
				} else {
					consumeBill.BaseValue = bills[i].BaseValueLeft
					consumeBillsMap[bills[i].ID] = consumeBill
				}
				rechargeValue -= bills[i].BaseValueLeft
				bills[i].BaseValueLeft = 0
			} else {
				if consumeBill, ok := consumeBillsMap[bills[i].ID]; !ok {
					consumeBillsMap[bills[i].ID] = TableConsumeBills{
						ID:            uuid.NewV4(),
						BaseValue:     rechargeValue,
						BillID:        &primaryBill.ID,
						ConsumeBillID: &bills[i].ID,
					}
				} else {
					consumeBill.BaseValue = rechargeValue
					consumeBillsMap[bills[i].ID] = consumeBill
				}
				bills[i].BaseValueLeft -= rechargeValue
				rechargeValue = 0
				break
			}
		}
	}

	if rechargeValue != 0 {
		cutil.Logger.Error("ActiveSecondaryCard 账户金额与充值流水不匹配")
		tx.Rollback()
		return ErrBeyondPrimaryAccountBalance
	}

	consumeBills := make([]TableConsumeBills, 0)
	for _, v := range consumeBillsMap {
		consumeBills = append(consumeBills, v)
	}

	primaryBill.AfterAccount, secondaryBill.AfterAccount = primaryAccount, cardAccount
	err = ActiveSecondaryCard(tx, *tabCard, primaryAccountID, *cardAccount, *primaryBill, *secondaryBill, consumeBills, bills)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("ActiveSecondaryCard 更新卡激活状态错误:%v", err))
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// UpdateCard 更新卡
func UpdateCard(card TableCard) error {
	return entity.Conn.Scopes(idCondition(card.ID)).Updates(&card).Error
}

// GetCardsByAccountID 根据账户id查询卡
func GetCardsByAccountID(accountID uuid.UUID) ([]*TableCard, error) {
	var cards []*TableCard
	//实现同一账户下多张卡按照最近使用时间排序
	subQuery := entity.Conn.
		Model(&TableAccountBill{}).
		Select("card_id, split_part(string_agg(created_at::VARCHAR,',' ORDER BY created_at DESC),',',1) bill_created_at").
		Where("account_id = ?", accountID).
		Group("card_id")
	err := entity.Conn.Distinct().
		Select("cards.*, tmp.bill_created_at").
		Where("account_id = ?", accountID).
		Joins("left join (?) tmp on tmp.card_id = cards.id", subQuery).
		Order("bill_created_at DESC nulls last").Find(&cards).Error
	return cards, err
}

// GetCards 查询卡列表
func GetCards(branchIDs []uuid.UUID, category, status string, offset, limit int32) ([]TableCard, error) {
	var cards []TableCard
	err := entity.Conn.Scopes(
		branchIDsCondition(branchIDs),
		cardSubCategoryCondition(category),
		cardStatusCondition(status),
		pagingCondition(offset, limit),
	).Order("created_at desc").
		Find(&cards).Error
	return cards, err
}

// CountCards 计算卡数量
func CountCards(branchIDs []uuid.UUID, category, status string) (int64, error) {
	var count int64
	err := entity.Conn.Model(&TableCard{}).Scopes(
		branchIDsCondition(branchIDs),
		cardSubCategoryCondition(category),
		cardStatusCondition(status),
	).Count(&count).Error
	return count, err
}

func cardCodeCondition(code string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("cards.code = ?", code)
	}
}

func cardStatusCondition(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			db = db.Where("cards.status = ?", status)
		}
		return db
	}
}

func branchIDsCondition(branchIDs []uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(branchIDs) > 0 {
			db = db.Where("cards.open_branch_id in (?)", branchIDs)
		}
		return db
	}
}

func cardSubCategoryCondition(category string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if category != "" {
			db = db.Where("cards.sub_category = ?", category)
		}
		return db
	}
}

// SearchCards 卡查询
func SearchCards(scopes []func(db *gorm.DB) *gorm.DB) ([]TableCard, error) {
	var cards []TableCard
	err := entity.Conn.Scopes(scopes...).Order("created_at desc").Limit(50).Find(&cards).Error
	return cards, err
}
