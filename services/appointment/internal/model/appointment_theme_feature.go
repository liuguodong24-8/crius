package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// TableAppointmentThemeFeature 主题预约特色
type TableAppointmentThemeFeature struct {
	ID        uuid.UUID   `gorm:"column:id"`
	Name      string      `gorm:"column:name"`
	Weight    int32       `gorm:"column:weight"`
	Status    util.Status `gorm:"column:status"`
	Icon      string      `gorm:"column:icon"`
	CreatedAt *time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time  `gorm:"column:updated_at"`
}

// TableName 表名
func (TableAppointmentThemeFeature) TableName() string {
	return "appointment.appointment_theme_feature"
}

// CreateAppointmentThemeFeature 创建主题预约特色
func CreateAppointmentThemeFeature(t *TableAppointmentThemeFeature) error {
	return entity.Conn.Create(t).Error
}

// UpdateAppointmentThemeFeature 更新主题预约特色
func UpdateAppointmentThemeFeature(t *TableAppointmentThemeFeature) error {
	return entity.Conn.Scopes(util.ColumnEqualScope("id", t.ID)).Select("name", "weight", "icon", "status").Updates(t).Error
}

// UpdateAppointmentThemeFeatureStatus 更新主题预约特色状态
func UpdateAppointmentThemeFeatureStatus(id uuid.UUID, status util.Status) error {
	return entity.Conn.Model(&TableAppointmentThemeFeature{}).Scopes(util.ColumnEqualScope("id", id)).UpdateColumn("status", status.String()).Error
}

// ShowAppointmentThemeFeatureByName 名字查询主题预约特色
func ShowAppointmentThemeFeatureByName(name string) (*TableAppointmentThemeFeature, error) {
	feature := new(TableAppointmentThemeFeature)
	err := entity.Conn.Scopes(util.ColumnEqualScope("name", name)).Take(feature).Error
	return feature, err
}

// GetAppointmentThemeFeatures 获取主题预约特色列表
func GetAppointmentThemeFeatures(name string, status util.Status, offset, limit int32) ([]TableAppointmentThemeFeature, int64, error) {
	var total int64
	features := make([]TableAppointmentThemeFeature, 0)
	db := entity.Conn.Scopes(util.ColumnLikeScope("name", name), util.ColumnEqualScopeDefault("status", status.String()))
	if err := db.Model(&TableAppointmentThemeFeature{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return features, 0, nil
	}
	if err := db.Scopes(PagingCondition(offset, limit)).Order("weight desc").Find(&features).Error; err != nil {
		return nil, 0, err
	}
	return features, total, nil
}

// ShowAppointmentThemeFeature 获取主题预约特色
func ShowAppointmentThemeFeature(id uuid.UUID) (*TableAppointmentThemeFeature, error) {
	feature := new(TableAppointmentThemeFeature)
	err := entity.Conn.Scopes(util.ColumnEqualScope("id", id)).Take(feature).Error
	return feature, err
}
