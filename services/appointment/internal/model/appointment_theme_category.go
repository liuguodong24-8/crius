package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// TableAppointmentThemeCategory 主题预约分类
type TableAppointmentThemeCategory struct {
	ID        uuid.UUID   `gorm:"column:id"`
	Name      string      `gorm:"column:name"`
	Weight    int32       `gorm:"column:weight"`
	Status    util.Status `gorm:"column:status"`
	CreatedAt *time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time  `gorm:"column:updated_at"`
}

// TableName 表名
func (TableAppointmentThemeCategory) TableName() string {
	return "appointment.appointment_theme_category"
}

// CreateAppointmentThemeCategory 创建主题预约分类
func CreateAppointmentThemeCategory(t *TableAppointmentThemeCategory) error {
	return entity.Conn.Create(t).Error
}

// UpdateAppointmentThemeCategory 更新主题预约分类
func UpdateAppointmentThemeCategory(t *TableAppointmentThemeCategory) error {
	return entity.Conn.Scopes(util.ColumnEqualScope("id", t.ID)).Select("name", "weight", "status").Updates(t).Error
}

// UpdateAppointmentThemeCategoryStatus 更新主题预约分类状态
func UpdateAppointmentThemeCategoryStatus(id uuid.UUID, status util.Status) error {
	return entity.Conn.Model(&TableAppointmentThemeCategory{}).Scopes(util.ColumnEqualScope("id", id)).UpdateColumn("status", status.String()).Error
}

// ShowAppointmentThemeCategoryByName 名字查询主题预约分类
func ShowAppointmentThemeCategoryByName(name string) (*TableAppointmentThemeCategory, error) {
	category := new(TableAppointmentThemeCategory)
	err := entity.Conn.Scopes(util.ColumnEqualScope("name", name)).Take(category).Error
	return category, err
}

// GetAppointmentThemeCategories 获取主题预约分类列表
func GetAppointmentThemeCategories(name string, status util.Status, offset, limit int32) ([]TableAppointmentThemeCategory, int64, error) {
	var total int64
	categories := make([]TableAppointmentThemeCategory, 0)
	db := entity.Conn.Scopes(util.ColumnLikeScope("name", name), util.ColumnEqualScopeDefault("status", status.String()))
	if err := db.Model(&TableAppointmentThemeCategory{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return categories, 0, nil
	}
	if err := db.Scopes(PagingCondition(offset, limit)).Order("weight desc").Find(&categories).Error; err != nil {
		return nil, 0, err
	}
	return categories, total, nil
}

// ShowAppointmentThemeCategory 查询主题预约分类
func ShowAppointmentThemeCategory(id uuid.UUID) (*TableAppointmentThemeCategory, error) {
	category := new(TableAppointmentThemeCategory)
	err := entity.Conn.Scopes(util.ColumnEqualScope("id", id)).Take(category).Error
	return category, err
}
