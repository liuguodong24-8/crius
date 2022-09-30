package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// TableBranchBusiness 表结构
type TableBranchBusiness struct {
	ID         uuid.UUID              `gorm:"column:id"`
	BranchID   *uuid.UUID             `gorm:"column:branch_id"`
	BeginDate  *time.Time             `gorm:"column:begin_date"`
	EndDate    *time.Time             `gorm:"column:end_date"`
	Weeks      *fields.Int8Arr        `gorm:"column:weeks"`
	BeginTime  *fields.LocalTime      `gorm:"column:begin_time"`
	EndTime    *fields.LocalTime      `gorm:"column:end_time"`
	IsNextDay  bool                   `gorm:"column:is_next_day"`
	MerchantID *uuid.UUID             `gorm:"column:merchant_id"`
	CreatedAt  *time.Time             `gorm:"column:created_at"`
	UpdatedAt  *time.Time             `gorm:"column:updated_at"`
	Status     string                 `gorm:"column:status"`
	Category   BranchBusinessCategory `gorm:"column:category"`
	LoadExtra  string                 `gorm:"-"` // 导入信息 不处理
}

// TableName 表名
func (TableBranchBusiness) TableName() string {
	return "merchant_basic.branch_business"
}

// BranchBusinessCategory ...
type BranchBusinessCategory string

const (
	//BranchBusinessCategoryNormal 营业时间
	BranchBusinessCategoryNormal BranchBusinessCategory = "normal"

	//BranchBusinessCategorySpecial 特殊营业时间
	BranchBusinessCategorySpecial BranchBusinessCategory = "special"
)

// GetBranchBusiness 获取列表
func GetBranchBusiness(business TableBranchBusiness, offset, limit int32) ([]TableBranchBusiness, int64, error) {
	var businesses []TableBranchBusiness
	var count int64
	db := entity.Conn.Scopes(crius.ColumnEqualScopeDefault("branch_id", *business.BranchID), crius.ColumnEqualScope("merchant_id", *business.MerchantID), crius.ColumnEqualScopeDefault("status", business.Status),
		crius.ColumnEqualScopeDefault("category", business.Category))
	if err := db.Model(&TableBranchBusiness{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return businesses, 0, nil
	}

	if err := db.Scopes(pagingCondition(offset, limit)).Find(&businesses).Error; err != nil {
		return nil, 0, err
	}
	return businesses, count, nil
}

// ShowBranchBusinessByBranchIDDate 根据门店id日期查询
func ShowBranchBusinessByBranchIDDate(branchID, merchantID uuid.UUID, date time.Time) (*TableBranchBusiness, error) {
	business := new(TableBranchBusiness)
	if err := entity.Conn.Scopes(crius.ColumnEqualScopeDefault("branch_id", branchID), crius.ColumnEqualScope("merchant_id", merchantID), branchBusinessDateCondition(date), crius.ColumnEqualScopeDefault("status", crius.StatusOpened.String())).
		Order("created_at desc").Take(business).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	} else {
		return business, nil
	}

	err := entity.Conn.Scopes(crius.ColumnEqualScopeDefault("branch_id", branchID), crius.ColumnEqualScope("merchant_id", merchantID), crius.ArrayAnyScope("weeks", int(date.Weekday()))).Take(business).Error
	return business, err
}

// CreateBranchBusiness 创建营业时间
func CreateBranchBusiness(business TableBranchBusiness) error {
	return entity.Conn.Create(&business).Error
}

// UpdateBranchBusinessSpecial 更新特殊营业时间
func UpdateBranchBusinessSpecial(business TableBranchBusiness) error {
	return entity.Conn.Scopes(crius.ColumnEqualScope("id", business.ID)).Select("begin_date", "end_date", "begin_time", "end_time", "status", "is_next_day").Updates(&business).Error
}

// UpdateBranchBusinessNormal 更新普通营业时间
func UpdateBranchBusinessNormal(businesses []TableBranchBusiness, branchID uuid.UUID) error {
	tx := entity.Conn.Begin()
	if err := tx.Scopes(crius.ColumnEqualScope("branch_id", branchID), crius.ColumnEqualScope("category", string(BranchBusinessCategoryNormal))).
		Delete(&TableBranchBusiness{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&businesses).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// UpdateBranchBusinessStatus 更新营业时间状态
func UpdateBranchBusinessStatus(id uuid.UUID, status string) error {
	return entity.Conn.Model(&TableBranchBusiness{}).Scopes(crius.ColumnEqualScope("id", id)).Update("status", status).Error
}

func branchBusinessDateCondition(date time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("begin_date <= ? and end_date >= ?", date, date)
	}
}
