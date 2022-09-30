package model

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// WechatStatus 短信模版状态
type WechatStatus string

const (
	// WechatStatusSuccess 成功
	WechatStatusSuccess WechatStatus = "success"
	// WechatStatusFail 失败
	WechatStatusFail WechatStatus = "fail"
)

// WechatStat 微信模版记录
type WechatStat struct {
	ID             uuid.UUID     `json:"id" gorm:"primary_key"`
	MerchantID     *uuid.UUID    `json:"merchant_id"`
	BranchID       *uuid.UUID    `json:"branch_id"`
	MessageType    string        `json:"message_type"`
	System         string        `json:"system"` // 调用来源
	MemberID       *uuid.UUID    `json:"member_id"`
	MemberWechatID *uuid.UUID    `json:"member_wechat_id"`
	Request        *pkgs.Params  `json:"request"`
	WechatResponse *pkgs.Params  `json:"wechat_response"`
	Status         WechatStatus  `json:"status"` // 状态
	Extra          *pkgs.Params  `json:"extra"`
	CreatedAt      pkgs.NullTime `json:"created_at" form:"created_at"` // 统一定义时间字段处理，格式化
	UpdatedAt      pkgs.NullTime `json:"updated_at" form:"updated_at"` // 统一定义时间字段处理，格式化
}

// TableName 表名
func (w WechatStat) TableName() string {
	return "message.wechat_stats"
}
