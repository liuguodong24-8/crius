package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// TableRoomTypeCategory 表结构
type TableRoomTypeCategory struct {
	ID         uuid.UUID    `gorm:"column:id"`
	Name       string       `gorm:"column:name"`
	Category   int8         `gorm:"column:category"`
	Status     crius.Status `gorm:"column:status"`
	MerchantID *uuid.UUID   `gorm:"column:merchant_id"`
	CreatedAt  *time.Time   `gorm:"column:created_at"`
	UpdatedAt  *time.Time   `gorm:"column:updated_at"`
	LoadExtra  string       `gorm:"-"` // 导入信息 不处理
}

// TableName 表名
func (TableRoomTypeCategory) TableName() string {
	return "merchant_basic.room_type_category"
}

// ShowRoomTypeCategoryByName 查询单条
func ShowRoomTypeCategoryByName(name string, merchantID uuid.UUID) (*TableRoomTypeCategory, error) {
	var category TableRoomTypeCategory
	if name == "" {
		return nil, gorm.ErrRecordNotFound
	}
	err := entity.Conn.Scopes(crius.ColumnEqualScopeDefault("name", name), crius.ColumnEqualScope("merchant_id", merchantID)).Take(&category).Error
	return &category, err
}

// CreateRoomTypeCategory 创建
func CreateRoomTypeCategory(category TableRoomTypeCategory) error {
	return entity.Conn.Create(&category).Error
}

// GetRoomTypeCategories 获取列表
func GetRoomTypeCategories(category TableRoomTypeCategory, offset, limit int32) ([]TableRoomTypeCategory, int64, error) {
	categories := make([]TableRoomTypeCategory, 0)
	var count int64

	db := entity.Conn.Scopes(crius.ColumnLikeScope("name", category.Name), crius.ColumnEqualScopeDefault("status", category.Status.String()), crius.ColumnEqualScopeDefault("category", category.Category), crius.ColumnEqualScope("merchant_id", *category.MerchantID))
	err := db.Model(&TableRoomTypeCategory{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return categories, 0, nil
	}
	err = db.Scopes(pagingCondition(offset, limit)).Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}
	return categories, count, nil
}

// UpdateRoomTypeCategory 更新
func UpdateRoomTypeCategory(category TableRoomTypeCategory) error {
	return entity.Conn.Scopes(crius.ColumnEqualScope("id", category.ID)).Updates(&category).Error
}
