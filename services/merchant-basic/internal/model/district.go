package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
)

// TableDistrict 区域表
type TableDistrict struct {
	ID        uuid.UUID    `gorm:"column:id"`
	Name      string       `gorm:"column:name"`
	Code      int          `gorm:"column:code"`
	Status    crius.Status `gorm:"column:status"`
	CreatedAt *time.Time   `gorm:"column:created_at"`
	UpdatedAt *time.Time   `gorm:"column:updated_at"`
	LoadExtra string       `gorm:"-"` // 导入信息 不处理
}

// TableName 指定表名
func (TableDistrict) TableName() string {
	return "merchant_basic.district"
}

// CreateDistrict 创建
func CreateDistrict(district TableDistrict) error {
	return entity.Conn.Omit("code").Create(&district).Error
}

// UpdateDistrict 创建
func UpdateDistrict(district TableDistrict) error {
	return entity.Conn.Scopes(crius.ColumnEqualScope("id", district.ID)).Updates(&district).Error
}

// GetDistricts 获取列表
func GetDistricts(district TableDistrict, offset, limit int32) ([]TableDistrict, int64, error) {
	var districts []TableDistrict
	var count int64
	db := entity.Conn.Scopes(crius.ColumnLikeScope("name", district.Name), crius.ColumnEqualScopeDefault("status", district.Status.String()))
	err := db.Model(&TableDistrict{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return districts, 0, err
	}
	err = db.Scopes(pagingCondition(offset, limit)).Order("created_at desc").Find(&districts).Error
	return districts, count, err
}
