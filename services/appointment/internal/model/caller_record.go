package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gorm.io/gorm"
)

// TableCallerRecord 来电记录表
type TableCallerRecord struct {
	ID         uuid.UUID         `gorm:"column:id"`
	MerchantID *uuid.UUID        `gorm:"column:merchant_id"`
	CallerID   *uuid.UUID        `gorm:"column:caller_id"`
	Phone      string            `gorm:"column:phone"`
	Operator   *uuid.UUID        `gorm:"column:operator"`
	CallAction *fields.StringArr `gorm:"column:call_action"`
	CallAt     *time.Time        `gorm:"column:call_at"`
	CreatedAt  *time.Time        `gorm:"column:created_at"`
	UpdatedAt  *time.Time        `gorm:"column:updated_at"`
}

// TableName 表名
func (TableCallerRecord) TableName() string {
	return "appointment.appointment_caller_records"
}

// CreatedAtGreateThan created_at大于
func CreatedAtGreateThan(t time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !t.IsZero() {
			return db.Where("created_at > ?", t)
		}
		return db
	}
}

// CallAtRangeCondition 最后来电时间
func CallAtRangeCondition(dateStart, dateEnd time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !dateStart.IsZero() && !dateEnd.IsZero() && !dateEnd.Before(dateStart) {
			return db.Where("call_at >= ? and call_at <= ?", dateStart, dateEnd)
		}
		return db
	}
}
