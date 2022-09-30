package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"time"
)

// PointRule 积分规则
type PointRule struct {
	ID          uuid.UUID       `json:"id"`
	MerchantID  *uuid.UUID      `json:"merchant_id"`
	RuleName    string          `json:"rule_name"`
	GainRules   *Rules          `json:"gain_rules"`
	UseRules    *Rules          `json:"use_rules"`
	ValidityDay int32           `json:"validity_day"`
	BranchIDs   *fields.UUIDArr `json:"branch_ids" gorm:"column:branch_ids"`
	Status      util.Status     `json:"status"`
	Extra       *pkgs.Params    `json:"extra"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *time.Time      `json:"deleted_at"`
}

// TableName table name
func (PointRule) TableName() string {
	return `merchant_basic.point_rule`
}

// Rule 积分规则详情
type Rule struct {
	CategoryID uuid.UUID `json:"category_id"`
	Point      int64     `json:"point"`
	Fee        int64     `json:"fee"`
}

// Rules 积分规则
type Rules []Rule

// Value 将对象转换为数据库可存储类型
func (r *Rules) Value() (driver.Value, error) {
	if nil == r {
		return nil, nil
	}

	return json.Marshal(r)
}

// Scan 将数据库对象转换成可以使用的golang 属性
func (r *Rules) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("point rules error")
	}

	if err := json.Unmarshal(bytes, r); nil != err {
		return errors.New("point rules error json")
	}

	return nil
}

// GetBranchPointCategoryRuleRequest 获取门店积分类型对应规则
type GetBranchPointCategoryRuleRequest struct {
	MerchantID   string
	BranchID     string
	CategoryCode string
}

// BranchPointRule 门店积分规则
type BranchPointRule struct {
	Point       int64
	Fee         int64
	ValidityDay int32
}

// GetBranchPointCategoryRule 获取门店积分类型对应规则 获取规则、抵扣规则、错误
func GetBranchPointCategoryRule(req GetBranchPointCategoryRuleRequest) (gain *BranchPointRule, use *BranchPointRule, err error) {
	// 查找门店积分规则
	var rule PointRule
	err = entity.Conn.Model(&PointRule{}).Scopes(util.ColumnEqualScope("merchant_id", req.MerchantID), util.ArrayAnyScope("branch_ids", req.BranchID)).Order("updated_at desc").First(&rule).Error
	if err != nil {
		return
	}

	// 查找积分消费类型
	var category ConsumeCategory
	err = entity.Conn.Model(&ConsumeCategory{}).Scopes(util.ColumnEqualScope("merchant_id", req.MerchantID), util.ColumnEqualScope("code", req.CategoryCode)).First(&category).Error
	if err != nil {
		return
	}

	if rule.GainRules != nil {
		for _, v := range *rule.GainRules {
			if v.CategoryID == category.ID {
				gain = &BranchPointRule{
					Point:       v.Point,
					Fee:         v.Fee,
					ValidityDay: rule.ValidityDay,
				}
			}
		}
	}

	if rule.UseRules != nil {
		for _, v := range *rule.UseRules {
			if v.CategoryID == category.ID {
				use = &BranchPointRule{
					Point:       v.Point,
					Fee:         v.Fee,
					ValidityDay: rule.ValidityDay,
				}
			}
		}
	}

	return
}

//GetBranchPointRules 门店所有积分规则
func GetBranchPointRules(branchID uuid.UUID) ([]PointRule, error) {
	var rules []PointRule
	err := entity.Conn.Model(&PointRule{}).Scopes(util.ArrayAnyScope("branch_ids", branchID)).Order("updated_at desc").Find(&rules).Error
	return rules, err
}

// GetBranchPointGainRule 门店积分获取规则
func GetBranchPointGainRule(branchID uuid.UUID) (*PointRule, error) {
	var rule PointRule
	err := entity.Conn.Scopes(crius.ArrayAnyScope("branch_ids", branchID)).First(&rule).Error
	return &rule, err
}
