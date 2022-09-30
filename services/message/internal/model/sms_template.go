package model

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// SmsTemplate 短信模版
type SmsTemplate struct {
	ID          uuid.UUID     `json:"id" gorm:"primary_key"`
	MerchantID  uuid.UUID     `json:"merchant_id"`
	Sign        string        `json:"sign"`
	Name        string        `json:"name"`
	Category    string        `json:"category"`
	CategoryKey string        `json:"category_key"`
	Content     string        `json:"content"`
	Extra       *pkgs.Params  `json:"extra"`
	Status      util.Status   `json:"status"`
	CreatedAt   pkgs.NullTime `json:"created_at" form:"created_at"` // 统一定义时间字段处理，格式化
	UpdatedAt   pkgs.NullTime `json:"updated_at" form:"updated_at"` // 统一定义时间字段处理，格式化
}

// TableName 表名
func (s SmsTemplate) TableName() string {
	return "message.sms_templates"
}
