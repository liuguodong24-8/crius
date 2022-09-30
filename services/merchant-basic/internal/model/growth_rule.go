package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"time"
)

// TableGrowthRule 表结构
type TableGrowthRule struct {
	ID         uuid.UUID       `gorm:"column:id"`
	Name       string          `gorm:"column:name"`
	GrowthGain *GrowthRules    `gorm:"column:growth_gain"`
	ExpireDay  int32           `gorm:"column:expire_day"`
	Branches   *fields.UUIDArr `gorm:"column:branches"`
	Status     string          `gorm:"column:status"`
	MerchantID uuid.UUID       `gorm:"column:merchant_id"`
	CreatedAt  *time.Time      `gorm:"column:created_at"`
	UpdatedAt  *time.Time      `gorm:"column:updated_at"`
}

// TableName 表名
func (TableGrowthRule) TableName() string {
	return "merchant_basic.growth_rule"
}

// GrowthRule 成长值规则详情
type GrowthRule struct {
	ConsumeTypeID uuid.UUID `json:"consume_type"`
	Cost          int64     `json:"cost"`
}

// GrowthRules 成长值规则
type GrowthRules []GrowthRule

// Value 将对象转换为数据库可存储类型
func (r *GrowthRules) Value() (driver.Value, error) {
	if nil == r {
		return nil, nil
	}

	return json.Marshal(r)
}

// Scan 将数据库对象转换成可以使用的golang 属性
func (r *GrowthRules) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("growth rules error")
	}

	if err := json.Unmarshal(bytes, r); nil != err {
		return errors.New("growth rules error json")
	}
	return nil
}

// CreateGrowthRule 新建
func CreateGrowthRule(t TableGrowthRule) error {
	return entity.Conn.Create(&t).Error
}

// CountGrowthRules count
func CountGrowthRules(name, status string, merchantID uuid.UUID) (int64, error) {
	var count int64
	err := entity.Conn.Model(&TableGrowthRule{}).Scopes(
		crius.ColumnLikeScope("name", name),
		crius.ColumnEqualScope("merchant_id", merchantID),
		crius.ColumnEqualScopeDefault("status", status),
	).Count(&count).Error
	return count, err
}

// GetGrowthRules 列表
func GetGrowthRules(name, status string, merchantID uuid.UUID, offset, limit int32) ([]TableGrowthRule, error) {
	var rules []TableGrowthRule
	err := entity.Conn.Model(&TableGrowthRule{}).Scopes(
		crius.ColumnLikeScope("name", name),
		crius.ColumnEqualScope("merchant_id", merchantID),
		crius.ColumnEqualScopeDefault("status", status),
		crius.PaginationScope(offset, limit),
	).Order("created_at desc").Find(&rules).Error
	return rules, err
}

// ShowGrowthRule 详情
func ShowGrowthRule(id uuid.UUID) (TableGrowthRule, error) {
	var rule TableGrowthRule
	err := entity.Conn.Model(&TableGrowthRule{}).Scopes(
		crius.ColumnEqualScope("id", id),
	).First(&rule).Error
	return rule, err
}

// UpdateGrowthRule 更新
func UpdateGrowthRule(rule TableGrowthRule) error {
	return entity.Conn.Model(&TableGrowthRule{}).Where("id", rule.ID).Updates(rule).Error
}

// GetHasGrowthRuleBranches 有规则的门店
func GetHasGrowthRuleBranches(id, merchantID uuid.UUID) ([]TableGrowthRule, error) {
	var rules []TableGrowthRule
	err := entity.Conn.Select("Branches").Scopes(
		crius.ColumnEqualScope("merchant_id", merchantID),
		crius.ColumnNotEqualScope("id", id),
	).Find(&rules).Error
	return rules, err
}

// GetBranchGrowthRule 门店规则
func GetBranchGrowthRule(branchID uuid.UUID) (*TableGrowthRule, error) {
	var rule TableGrowthRule
	err := entity.Conn.Scopes(crius.ArrayAnyScope("branches", branchID)).First(&rule).Error
	return &rule, err
}
