package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
)

// TableBrand 品牌表结构
type TableBrand struct {
	ID        uuid.UUID    `gorm:"column:id"`
	Name      string       `gorm:"column:name"`
	Status    crius.Status `gorm:"column:status"`
	Order     int32        `gorm:"column:order"`
	CreatedAt *time.Time   `gorm:"column:created_at"`
	UpdatedAt *time.Time   `gorm:"column:updated_at"`
}

// TableName 表名
func (TableBrand) TableName() string {
	return "merchant_basic.brand"
}

// ShowBrandByName 根据名字查询品牌
func ShowBrandByName(name string) (*TableBrand, error) {
	brand := new(TableBrand)
	err := entity.Conn.Scopes(crius.ColumnEqualScope("name", name)).Take(brand).Error
	return brand, err
}

// ShowBrand 查询品牌
func ShowBrand(id uuid.UUID) (*TableBrand, error) {
	brand := new(TableBrand)
	err := entity.Conn.Scopes(crius.ColumnEqualScope("id", id)).Take(brand).Error
	return brand, err
}

// CreateBrand 创建品牌
func CreateBrand(brand TableBrand) error {
	return entity.Conn.Create(&brand).Error
}

// UpdateBrand 更新品牌
func UpdateBrand(brand TableBrand) error {
	return entity.Conn.Select("name", "status", "order").Updates(&brand).Error
}

// UpdateBrandStatus 更新品牌状态
func UpdateBrandStatus(id uuid.UUID, status string) error {
	return entity.Conn.Model(&TableBrand{}).Scopes(crius.ColumnEqualScope("id", id)).Update("status", status).Error
}

// GetBrands 获取品牌列表
func GetBrands(name, status string, offset, limit int32) ([]TableBrand, int64, error) {
	var count int64
	var brands []TableBrand
	db := entity.Conn.Scopes(crius.ColumnLikeScope("name", name), crius.ColumnEqualScopeDefault("status", status))
	if err := db.Model(&TableBrand{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return brands, 0, nil
	}
	if err := db.Scopes(pagingCondition(offset, limit)).Order(`"order" desc`).Find(&brands).Error; err != nil {
		return nil, 0, err
	}
	return brands, count, nil
}
