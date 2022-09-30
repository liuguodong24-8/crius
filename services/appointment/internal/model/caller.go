package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gorm.io/gorm"
)

// TableCaller 来电用户表
type TableCaller struct {
	ID             uuid.UUID         `gorm:"column:id"`
	MerchantID     *uuid.UUID        `gorm:"column:merchant_id"`
	PhoneCode      string            `gorm:"column:phone_code"`
	Phone          string            `gorm:"column:phone"`
	PhoneTail      string            `gorm:"column:phone_tail"`
	PhoneSuffix    string            `gorm:"column:phone_suffix"`
	CallerName     string            `gorm:"column:caller_name"`
	Tags           *pkgs.ParamsArr   `gorm:"column:tags"`
	Gender         int8              `gorm:"column:gender"`
	IsBlcak        bool              `gorm:"column:is_black"`
	BlackReason    string            `gorm:"column:black_reason"`
	LastCallAt     *time.Time        `gorm:"column:last_call_at"`
	LastOperator   *uuid.UUID        `gorm:"column:last_operator"`
	LastCallAction *fields.StringArr `gorm:"column:last_call_action"`
	CreatedAt      *time.Time        `gorm:"column:created_at"`
	UpdatedAt      *time.Time        `gorm:"column:updated_at"`
}

// TableName 表名
func (TableCaller) TableName() string {
	return "appointment.appointment_caller"
}

// BlackCondition black
func BlackCondition(black int8) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if black == 1 {
			return db.Where("is_black = ?", true)
		} else if black == 2 {
			return db.Where("is_black = ?", false)
		}
		return db
	}
}

// LastCallAtRangeCondition 最后来电时间
func LastCallAtRangeCondition(dateStart, dateEnd time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !dateStart.IsZero() && !dateEnd.IsZero() && !dateEnd.Before(dateStart) {
			return db.Where("last_call_at >= ? and last_call_at <= ?", dateStart, dateEnd)
		}
		return db
	}
}
