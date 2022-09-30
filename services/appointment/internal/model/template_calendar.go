package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// CalendarCategory 模版日历类型
type CalendarCategory string

const (
	// CalendarHoliday 节假日
	CalendarHoliday CalendarCategory = `holiday`
	// CalendarNormal 普通
	CalendarNormal CalendarCategory = `normal`
)

// String ...
func (c CalendarCategory) String() string {
	return string(c)
}

// StringToCalendarCategory 字符串转type
func StringToCalendarCategory(s string) CalendarCategory {
	return CalendarCategory(s)
}

// TableTemplateCalendar 模版日历
type TableTemplateCalendar struct {
	ID           uuid.UUID        `json:"id" gorm:"column:id"`
	MerchantID   uuid.UUID        `json:"merchant_id" gorm:"column:merchant_id"`
	BranchID     uuid.UUID        `json:"branch_id" gorm:"column:branch_id"`
	BusinessDate fields.DateTime  `json:"business_date" gorm:"column:business_date"`
	TemplateID   uuid.UUID        `json:"template_id" gorm:"column:template_id"`
	Category     CalendarCategory `json:"category" gorm:"column:category"`
	ThemeIDs     *fields.UUIDArr  `json:"theme_ids" gorm:"column:theme_ids"`
	CreatedAt    *time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    *time.Time       `json:"updated_at" gorm:"column:updated_at"`
}

// TableName ...
func (t TableTemplateCalendar) TableName() string {
	return `appointment.appointment_template_calendar`
}

// ShowTemplateCalendarByBranchIDDate 根据门店id和日期查询日历模板
func ShowTemplateCalendarByBranchIDDate(branchID uuid.UUID, date time.Time) (*TableTemplateCalendar, error) {
	calendar := new(TableTemplateCalendar)
	err := entity.Conn.Scopes(util.ColumnEqualScope("branch_id", branchID), util.ColumnEqualScope("business_date", date)).Take(calendar).Error
	return calendar, err
}
