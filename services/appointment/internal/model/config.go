package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TableAppointmentConfig 预约参数
type TableAppointmentConfig struct {
	ID                       uuid.UUID  `json:"id" gorm:"column:id"`
	MerchantID               uuid.UUID  `json:"merchant_id" gorm:"column:merchant_id"`
	KeepTime                 int16      `json:"keep_time" gorm:"column:keep_time"`
	RemindTime               float32    `json:"remind_time" gorm:"column:remind_time"`
	OrderLimit               int16      `json:"order_limit" gorm:"column:order_limit"`
	RoomNumWarn              int16      `json:"room_num_warn" gorm:"column:room_num_warn"`
	PaymentTime              int16      `json:"payment_time" gorm:"column:payment_time"`
	CancelTime               float32    `json:"cancel_time" gorm:"column:cancel_time"`
	RefundPercentBefore      float32    `json:"refund_percent_before" gorm:"column:refund_percent_before"`
	RefundPercentAfter       float32    `json:"refund_percent_after" gorm:"column:refund_percent_after"`
	BreachMonths             int16      `json:"breach_months" gorm:"column:breach_months"`
	BreachTotal              int16      `json:"breach_total" gorm:"column:breach_total"`
	DecorationFee            string     `json:"decoration_fee" gorm:"column:decoration_fee"`
	ThemeKeepTime            int16      `json:"theme_keep_time" gorm:"column:theme_keep_time"`
	ThemeCancelTime          float32    `json:"theme_cancel_time" gorm:"column:theme_cancel_time"`
	ThemeRefundPercentBefore float32    `json:"theme_refund_percent_before" gorm:"column:theme_refund_percent_before"`
	ThemeRefundPercentAfter  float32    `json:"theme_refund_percent_after" gorm:"column:theme_refund_percent_after"`
	CreatedAt                *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName ...
func (t TableAppointmentConfig) TableName() string {
	return `appointment.appointment_config`
}

// AppointmentConfigKey 预约配置cache key
const AppointmentConfigKey = "appointment:config:merchant_id:%s"
