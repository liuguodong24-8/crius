package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// TradeType 下单类型
type TradeType string

const (
	// TradeTypeMWEB H5下单
	TradeTypeMWEB TradeType = `MWEB`
)

// TradeStatus 交易状态
type TradeStatus string

const (
	// TradeStatusInit 初始
	TradeStatusInit TradeStatus = `init`
	// TradeStatusSuccess 成功
	TradeStatusSuccess TradeStatus = `success`
	// TradeStatusFail 失败
	TradeStatusFail TradeStatus = `fail`
	// TradeStatusClose 关闭
	TradeStatusClose TradeStatus = `close`
)

// WechatTrade 微信交易
type WechatTrade struct {
	ID             uuid.UUID    `json:"id" gorm:"column:id"`
	MerchantID     uuid.UUID    `json:"merchant_id" gorm:"column:merchant_id"`
	BranchID       uuid.UUID    `json:"branch_id" gorm:"column:branch_id"`
	TotalFee       int          `json:"total_fee" gorm:"column:total_fee"`
	TradeType      TradeType    `json:"trade_type" gorm:"column:trade_type"`
	TransactionID  string       `json:"transaction_id" gorm:"column:transaction_id"`
	OutTradeNo     string       `json:"out_trade_no" gorm:"column:out_trade_no"`
	NotifyURL      string       `json:"notify_url" gorm:"column:notify_url"`
	RequestParams  *pkgs.Params `json:"request_params" gorm:"column:request_params"`
	WechatRequest  *pkgs.Params `json:"wechat_request" gorm:"column:wechat_request"`
	WechatResponse *pkgs.Params `json:"wechat_response" gorm:"column:wechat_response"`
	NotifyContent  *pkgs.Params `json:"notify_content" gorm:"column:notify_content"`
	QueryContent   *pkgs.Params `json:"query_content" gorm:"column:query_content"`
	TradeStatus    TradeStatus  `json:"trade_status" gorm:"column:trade_status"`
	NotifyState    int8         `json:"notify_state" gorm:"column:notify_state"`
	Extra          *pkgs.Params `json:"extra" gorm:"column:extra"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

// TableName 表名
func (WechatTrade) TableName() string {
	return "payment.wechat_trades"
}
