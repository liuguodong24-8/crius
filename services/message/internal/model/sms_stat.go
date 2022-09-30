package model

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// SmsStatus 短信状态
type SmsStatus string

const (
	// SmsStatusSuccess 成功
	SmsStatusSuccess SmsStatus = "success"
	// SmsStatusFail 失败
	SmsStatusFail SmsStatus = "fail"
)

// SmsStat 短信记录
type SmsStat struct {
	ID          uuid.UUID     `json:"id" gorm:"primary_key"`
	MerchantID  *uuid.UUID    `json:"merchant_id"`
	BranchID    *uuid.UUID    `json:"branch_id"`
	MessageType string        `json:"message_type"`
	AreaCode    string        `json:"area_code"`
	Phone       string        `json:"phone"`
	Sign        string        `json:"sign"`
	System      string        `json:"system"` // 调用来源
	Status      SmsStatus     `json:"status"` // 状态
	Content     string        `json:"content"`
	Extra       *pkgs.Params  `json:"extra"`
	CreatedAt   pkgs.NullTime `json:"created_at" form:"created_at"` // 统一定义时间字段处理，格式化
	UpdatedAt   pkgs.NullTime `json:"updated_at" form:"updated_at"` // 统一定义时间字段处理，格式化
}

// TableName 表名
func (s SmsStat) TableName() string {
	return "message.sms_stats"
}
