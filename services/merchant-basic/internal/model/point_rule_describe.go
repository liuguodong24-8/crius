package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
)

// PointRuleDescribe 积分规则说明
type PointRuleDescribe struct {
	ID            uuid.UUID         `json:"id"`
	MerchantID    *uuid.UUID        `json:"merchant_id"`
	GraphicDetail *fields.StringArr `json:"graphic_detail"`
	Extra         *pkgs.Params      `json:"extra"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

// TableName table name
func (PointRuleDescribe) TableName() string {
	return `merchant_basic.point_rule_describe`
}
