package model

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// TableAppointmentTheme 主题
type TableAppointmentTheme struct {
	ID         uuid.UUID         `gorm:"column:id"`
	Name       string            `gorm:"column:name"`
	Color      string            `gorm:"column:color"`
	Weight     int32             `gorm:"column:weight"`
	Status     util.Status       `gorm:"column:status"`
	FeatureIDs *fields.UUIDArr   `gorm:"column:feature_ids"`
	Contents   *pkgs.ParamsArr   `gorm:"column:contents"`
	Style      string            `gorm:"column:style"`
	Images     *fields.StringArr `gorm:"column:images"`
	Video      string            `gorm:"column:video"`
	Details    *fields.StringArr `gorm:"column:details"`
	CategoryID *uuid.UUID        `gorm:"column:category_id"`
	CreatedAt  *time.Time        `gorm:"column:created_at"`
	UpdatedAt  *time.Time        `gorm:"column:updated_at"`
}

// AppointmentThemeRoomType 主题房型结构
type AppointmentThemeRoomType struct {
	TableAppointmentTheme
	PackageID     uuid.UUID       `gorm:"column:package_id"`
	PackageName   string          `gorm:"column:package_name"`
	Packages      *pkgs.ParamsArr `gorm:"column:packages"`
	Decoration    string          `gorm:"column:decoration"`
	Staffing      string          `gorm:"column:staffing"`
	CustomConfigs *pkgs.ParamsArr `gorm:"column:custom_configs"`
	RoomTypes     *pkgs.ParamsArr `gorm:"column:room_types"`
	CategoryName  string          `gorm:"column:category_name"`
}

// AppointmentThemeCategory 主题包含分类名
type AppointmentThemeCategory struct {
	TableAppointmentTheme
	CategoryName string `gorm:"column:category_name"`
}

// TableName 表名
func (TableAppointmentTheme) TableName() string {
	return "appointment.appointment_theme"
}

// ShowAppointmentThemeByName 根据名字查询主题
func ShowAppointmentThemeByName(name string) (*TableAppointmentTheme, error) {
	theme := new(TableAppointmentTheme)
	err := entity.Conn.Scopes(util.ColumnEqualScope("name", name)).Take(theme).Error
	return theme, err
}

// CreateAppointmentTheme 创建主题
func CreateAppointmentTheme(t *TableAppointmentTheme, p []TableAppointmentThemePackage) error {
	tx := entity.Conn.Begin()
	if err := tx.Create(t).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(p) != 0 {
		if err := tx.Create(&p).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// UpdateAppointmentTheme 更新主题
func UpdateAppointmentTheme(t *TableAppointmentTheme, p []TableAppointmentThemePackage) error {
	tx := entity.Conn.Begin()
	if err := tx.Scopes(util.ColumnEqualScope("id", t.ID)).Select("name", "color", "weight", "status", "feature_ids", "contents", "style", "images", "video", "details").UpdateColumns(t).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Scopes(util.ColumnEqualScope("theme_id", t.ID)).Delete(&TableAppointmentThemePackage{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(p) != 0 {
		if err := tx.Create(&p).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// UpdateAppointmentThemeStatus 更新主题状态
func UpdateAppointmentThemeStatus(id uuid.UUID, status util.Status) error {
	return entity.Conn.Model(&TableAppointmentTheme{}).Scopes(util.ColumnEqualScope("id", id)).UpdateColumn("status", status.String()).Error
}

// GetAppointmentThemes 获取主题列表
func GetAppointmentThemes(name string, status util.Status, id uuid.UUID, offset, limit int32) ([]TableAppointmentTheme, int64, error) {
	var total int64
	themes := make([]TableAppointmentTheme, 0)
	db := entity.Conn.Scopes(util.ColumnLikeScope("name", name), util.ColumnEqualScopeDefault("status", status.String()), util.ColumnEqualScopeDefault("category_id", id))
	if err := db.Model(&TableAppointmentTheme{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return themes, 0, nil
	}
	if err := db.Scopes(PagingCondition(offset, limit)).Order("weight desc").Find(&themes).Error; err != nil {
		return nil, 0, err
	}
	return themes, total, nil
}

// ShowAppointmentTheme 获取主题
func ShowAppointmentTheme(id uuid.UUID) (*TableAppointmentTheme, error) {
	theme := new(TableAppointmentTheme)
	err := entity.Conn.Scopes(util.ColumnEqualScope("id", id)).Take(theme).Error
	return theme, err
}

// ShowAppointmentThemeWithCategory 获取主题和分类名
func ShowAppointmentThemeWithCategory(id uuid.UUID) (*AppointmentThemeCategory, error) {
	themeCategory := new(AppointmentThemeCategory)
	err := entity.Conn.Model(&TableAppointmentTheme{}).Joins("inner join appointment.appointment_theme_category on appointment_theme.category_id = appointment_theme_category.id").
		Scopes(util.ColumnEqualScope("appointment_theme.id", id)).Select("appointment_theme.*", "appointment_theme_category.name as category_name").Take(themeCategory).Error
	return themeCategory, err
}

// GetAppointmentThemesByRoomType 根据房型获取主题套餐
func GetAppointmentThemesByRoomType(id, themeID uuid.UUID) ([]AppointmentThemeRoomType, error) {
	themes := make([]AppointmentThemeRoomType, 0)
	err := entity.Conn.Model(&TableAppointmentThemePackage{}).Joins("left join appointment.appointment_theme on appointment_theme_package.theme_id = appointment_theme.id").
		Joins("left join appointment.appointment_theme_category on appointment_theme.category_id = appointment_theme_category.id").
		Where(fmt.Sprintf(`appointment_theme_package.room_types @> '[{"id":"%s"}]'`, id.String())).Scopes(util.ColumnEqualScope("appointment_theme.status", util.StatusOpened.String()),
		util.ColumnEqualScopeDefault("appointment_theme.id", themeID)).Select("appointment_theme.*", "appointment_theme_package.name as package_name", "appointment_theme_package.decoration", "appointment_theme_package.staffing",
		"appointment_theme_package.custom_configs", "appointment_theme_package.room_types", "appointment_theme_package.id as package_id", "appointment_theme_category.name as category_name", "appointment_theme_package.packages").
		Order("appointment_theme_category.weight desc").Find(&themes).Error
	return themes, err
}
