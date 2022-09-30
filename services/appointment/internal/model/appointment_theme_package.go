package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// TableAppointmentThemePackage 主题套餐
type TableAppointmentThemePackage struct {
	ID            uuid.UUID       `gorm:"column:id"`
	Name          string          `gorm:"column:name"`
	Packages      *pkgs.ParamsArr `gorm:"column:packages"`
	Decoration    string          `gorm:"column:decoration"`
	Staffing      string          `gorm:"column:staffing"`
	CustomConfigs *pkgs.ParamsArr `gorm:"column:custom_configs"`
	RoomTypes     *pkgs.ParamsArr `gorm:"column:room_types"`
	ThemeID       *uuid.UUID      `gorm:"column:theme_id"`
	CreatedAt     *time.Time      `gorm:"column:created_at"`
	UpdatedAt     *time.Time      `gorm:"column:updated_at"`
}

// TableName 表名
func (TableAppointmentThemePackage) TableName() string {
	return "appointment.appointment_theme_package"
}

// GetAppointmentThemePackagesByThemeID 根据主题id查询套餐
func GetAppointmentThemePackagesByThemeID(id uuid.UUID) ([]TableAppointmentThemePackage, error) {
	packages := make([]TableAppointmentThemePackage, 0)
	err := entity.Conn.Scopes(util.ColumnEqualScope("theme_id", id)).Find(&packages).Error
	return packages, err
}

// ShowAppointmentThemePackage 获取主题套餐
func ShowAppointmentThemePackage(id uuid.UUID) (*TableAppointmentThemePackage, error) {
	p := new(TableAppointmentThemePackage)
	err := entity.Conn.Scopes(util.ColumnEqualScope("id", id)).Take(p).Error
	return p, err
}
