package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	cutil "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

const (
	billCodeSeqKey      = "bill_code_seq_%s"
	areaCodeBranchIDKey = "area_code:branch_id_%s"
)

// TableAccountBill 账单表
type TableAccountBill struct {
	ID                  uuid.UUID              `gorm:"column:id"`
	BillCode            string                 `gorm:"column:bill_code"`
	AccountID           *uuid.UUID             `gorm:"column:account_id"`
	CardID              *uuid.UUID             `gorm:"column:card_id"`
	CardCode            string                 `gorm:"column:card_code"`
	BranchID            *uuid.UUID             `gorm:"column:branch_id"`
	ChangeValue         int32                  `gorm:"column:change_value"`
	ChangeCategory      BillCategory           `gorm:"column:change_category"`
	ChangeType          BillType               `gorm:"column:change_type"`
	BaseValue           int32                  `gorm:"column:base_value"`
	GiftValue           int32                  `gorm:"column:gift_value"`
	Payments            *Payments              `gorm:"column:payments"`
	AfterAccount        *TableCardAccount      `gorm:"column:after_account"`
	StaffID             *uuid.UUID             `gorm:"column:staff_id"`
	OperatorComment     string                 `gorm:"column:operator_comment"`
	MerchantID          *uuid.UUID             `gorm:"column:merchant_id"`
	BaseValueLeft       int32                  `gorm:"column:base_value_left"`
	GiftValueLeft       int32                  `gorm:"column:gift_value_left"`
	PromotionOptions    *PromotionOptions      `json:"promotion_options" gorm:"column:promotion_options"`
	CreatedAt           *time.Time             `gorm:"column:created_at"`
	UpdatedAt           *time.Time             `gorm:"column:updated_at"`
	Extra               *pkgs.Params           `gorm:"column:extra"`
	ProductPackageBills []*TableProductPackage `gorm:"foreignkey:bill_id"`
	PosBillID        *uuid.UUID        `gorm:"column:pos_bill_id"`
	Status           BillStatus        `gorm:"column:status"`
}

// PromotionOption 充值开卡选择具体方案数量
type PromotionOption struct {
	ID            uuid.UUID                       `json:"id"`
	Name          string                          `json:"name"`
	Count         int                             `json:"count"`
	RechargeValue int32                           `json:"recharge_value"`
	BaseValue     int32                           `json:"base_value"`
	GiftValue     int32                           `json:"gift_value"`
	Products      *fields.ProductPackageTicketArr `json:"products"`
	Packages      *fields.ProductPackageTicketArr `json:"packages"`
	Tickets       *fields.ProductPackageTicketArr `json:"tickets"`
}

type (
	// BillCategory 账单类型
	BillCategory string
	// BillType 账单行为
	BillType string
	// BillStatus 账单状态
	BillStatus string
)

const (
	// BillCategoryRecharge 账单充值
	BillCategoryRecharge BillCategory = "recharge"
	// BillCategoryConsume 账单消费
	BillCategoryConsume BillCategory = "consume"
	// BillCategoryChange 修改余额
	BillCategoryChange BillCategory = "change"

	// BillTypeOpen 开卡
	BillTypeOpen BillType = "open"
	// BillTypeSub 开副卡
	BillTypeSub BillType = "sub"
	// BillTypeNobody 开不记名卡
	BillTypeNobody BillType = "nobody"
	// BillTypeRecharge 充值
	BillTypeRecharge BillType = "recharge"
	// BillTypeTransfer 副卡划账消费
	BillTypeTransfer BillType = "transfer"
	// BillTypeChange 修改余额
	BillTypeChange BillType = "change"
	// BillTypeDeduction 扣款
	BillTypeDeduction BillType = "deduction"
	// BillTypeReplace 补卡
	BillTypeReplace BillType = "replace"
	// BillTypeConsume 消费
	BillTypeConsume BillType = "consume"
	// BillTypeRefund 退款
	BillTypeRefund BillType = "refund"
	// BillStatusSuccess 成功
	BillStatusSuccess BillStatus = "success"
	// BillStatusCancel 取消
	BillStatusCancel BillStatus = "cancel"
)

// PromotionOptions 充值开卡选择具体方案数量
type PromotionOptions []PromotionOption

// Value 将对象转换为数据库可存储类型
func (j *PromotionOptions) Value() (driver.Value, error) {
	if nil == j {
		return nil, nil
	}

	return json.Marshal(j)
}

// Scan 将数据库对象转换成可以使用的golang 属性
func (j *PromotionOptions) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("promotion_options error")
	}

	if err := json.Unmarshal(bytes, j); nil != err {
		return errors.New("promotion_options json")
	}

	return nil
}

// TableName 表名
func (TableAccountBill) TableName() string {
	return "member_account.account_bills"
}

// BeforeCreate 创建之前生成账单号
func (t *TableAccountBill) BeforeCreate(tx *gorm.DB) error {
	if len(t.BillCode) > 0 {
		return nil
	}

	billCode, err := ShowBillCodeSeq()
	if err != nil {
		return err
	}
	t.BillCode = strconv.FormatInt(billCode, 10)
	return nil
}

// CreateAccountBill 创建
func CreateAccountBill(accountBill TableAccountBill) error {
	return entity.Conn.Create(&accountBill).Error
}

// ShowBillCodeSeq 获取BillCode自增序列
func ShowBillCodeSeq() (int64, error) {
	key := fmt.Sprintf(billCodeSeqKey, time.Now().Format("20060102"))
	code, err := redisEntity.INCR(key)
	if err != nil {
		return 0, err
	}
	redisEntity.EXPIRE(key, 24*3600)
	return code, nil
}

// SetAreaCode 根据branchid设置地区code缓存
func SetAreaCode(id uuid.UUID, code string) error {
	key := fmt.Sprintf(areaCodeBranchIDKey, id.String())
	err := redisEntity.SET(key, code)
	if err != nil {
		return err
	}
	//暂定缓存一天
	err = redisEntity.EXPIRE(key, 3600*24)
	return err
}

// GetAreaCode 根据branchid获取地区code缓存
func GetAreaCode(id uuid.UUID) (string, error) {
	key := fmt.Sprintf(areaCodeBranchIDKey, id.String())
	return redisEntity.GET(key)
}

// ShowBill 查询账单
func ShowBill(billID uuid.UUID) (*TableAccountBill, error) {
	var bill TableAccountBill
	err := entity.Conn.Scopes(idCondition(billID)).Take(&bill).Error
	return &bill, err
}

// ShowBillByCode 查询账单
func ShowBillByCode(billCode string) (*TableAccountBill, error) {
	var bill TableAccountBill
	err := entity.Conn.Preload("ProductPackageBills").Where("bill_code = ?", billCode).First(&bill).Error
	return &bill, err
}

// GetBills 获取账单列表
func GetBills(accountID uuid.UUID, changeCategory BillCategory, offset, limit int32) ([]TableAccountBill, error) {
	bills := make([]TableAccountBill, 0)
	db := entity.Conn.Distinct().Where("account_id = ?", accountID).Scopes(
		billCategoryCondition(changeCategory),
		pagingCondition(offset, limit))

	err := db.Order("created_at DESC").Find(&bills).Error
	if err != nil {
		return nil, err
	}
	return bills, nil
}

// CountBills 计算账单数量
func CountBills(accountID uuid.UUID, category BillCategory) (int64, error) {
	var count int64
	err := entity.Conn.Model(&TableAccountBill{}).Where("account_id = ?", accountID).Scopes(
		billCategoryCondition(category)).Count(&count).Error
	return count, err
}

// GetRechargeBillsByAccountID 根据账户id查询充值流水
func GetRechargeBillsByAccountID(tx *gorm.DB, id uuid.UUID) ([]TableAccountBill, error) {
	db := entity.Conn
	if tx != nil {
		db = tx
	}
	bills := make([]TableAccountBill, 0)
	err := db.Scopes(cutil.ColumnEqualScope("account_id", id), billValueNotEqual0()).Order("created_at asc").Find(&bills).Error
	return bills, err
}

func billCategoryCondition(category BillCategory) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch string(category) {
		case string(BillCategoryConsume):
			db = db.Where("change_category = ? AND change_type <> ?", BillCategoryConsume, BillTypeDeduction)
		case string(BillCategoryRecharge):
			db = db.Where("change_category = ?", BillCategoryRecharge)
		case string(BillCategoryChange):
			db = db.Where("change_category = ?", BillCategoryChange)
		case string(BillTypeDeduction):
			db = db.Where("change_type = ?", BillTypeDeduction)
		}
		return db
	}
}

func billValueNotEqual0() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("base_value_left != 0 or gift_value_left != 0")
	}
}

// GetReportBillsRequest 报表查询条件
type GetReportBillsRequest struct {
	Scopes   []func(db *gorm.DB) *gorm.DB
	WithPage bool
	Offset   int32
	Limit    int32
}

// GetReportBills 报表查询
func GetReportBills(req GetReportBillsRequest) ([]TableAccountBill, int64, error) {
	var total int64
	var items []TableAccountBill

	if err := entity.Conn.Model(&TableAccountBill{}).Scopes(req.Scopes...).Count(&total).Error; nil != err {
		return nil, 0, err
	}

	if req.WithPage {
		req.Scopes = append(req.Scopes, cutil.PaginationScope(req.Offset, req.Limit))
	}

	err := entity.Conn.Model(&TableAccountBill{}).Scopes(req.Scopes...).Order("created_at desc").Find(&items).Error

	return items, total, err
}

// GetAccountBillsByPosBillID 根据pos账单id查询流水
func GetAccountBillsByPosBillID(posBillID uuid.UUID, status BillStatus) ([]TableAccountBill, error) {
	bills := make([]TableAccountBill, 0)
	err := entity.Conn.Scopes(util.ColumnEqualScope("pos_bill_id", posBillID), util.ColumnEqualScope("status", string(status))).Find(&bills).Error
	return bills, err
}

// RefundBill 退款
func RefundBill(tx *gorm.DB, bills []TableAccountBill, refundBIlls []TableAccountBill) error {
	for _, v := range bills {
		if err := tx.Model(&TableAccountBill{}).Scopes(util.ColumnEqualScope("id", v.ID)).UpdateColumn("status", BillStatusCancel).Error; err != nil {
			return err
		}
		m := make(map[string]interface{})
		m["base_value"] = gorm.Expr("base_value+?", v.BaseValue)
		m["gift_value"] = gorm.Expr("gift_value+?", v.GiftValue)
		if err := tx.Model(&TableCardAccount{}).Scopes(util.ColumnEqualScope("id", *v.AccountID)).Updates(&m).Error; err != nil {
			return err
		}
	}
	return tx.Create(&refundBIlls).Error
}
