package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// RefundStatus 交易状态
type RefundStatus string

const (
	// RefundStatusInit 初始
	RefundStatusInit RefundStatus = `init`
	// RefundStatusSuccess 成功
	RefundStatusSuccess RefundStatus = `success`
	// RefundStatusFail 失败
	RefundStatusFail RefundStatus = `fail`
)

// WechatRefund 微信交易
type WechatRefund struct {
	ID             uuid.UUID    `json:"id" gorm:"column:id"`
	MerchantID     uuid.UUID    `json:"merchant_id" gorm:"column:merchant_id"`
	BranchID       uuid.UUID    `json:"branch_id" gorm:"column:branch_id"`
	WechatTradeID  uuid.UUID    `json:"wechat_trade_id" gorm:"column:wechat_trade_id"`
	RefundFee      int          `json:"refund_fee" gorm:"column:refund_fee"`
	RefundID       string       `json:"refund_id" gorm:"column:refund_id"`
	OutRefundNo    string       `json:"out_refund_no" gorm:"column:out_refund_no"`
	NotifyURL      string       `json:"notify_url" gorm:"column:notify_url"`
	RequestParams  *pkgs.Params `json:"request_params" gorm:"column:request_params"`
	WechatRequest  *pkgs.Params `json:"wechat_request" gorm:"column:wechat_request"`
	WechatResponse *pkgs.Params `json:"wechat_response" gorm:"column:wechat_response"`
	NotifyContent  *pkgs.Params `json:"notify_content" gorm:"column:notify_content"`
	QueryContent   *pkgs.Params `json:"query_content" gorm:"column:query_content"`
	RefundStatus   RefundStatus `json:"refund_status" gorm:"column:refund_status"`
	NotifyState    int8         `json:"notify_state" gorm:"column:notify_state"`
	Extra          *pkgs.Params `json:"extra" gorm:"column:extra"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

// TableName 表名
func (WechatRefund) TableName() string {
	return "payment.wechat_refunds"
}
