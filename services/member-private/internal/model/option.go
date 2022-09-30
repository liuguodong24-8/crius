package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gitlab.omytech.com.cn/micro-service/Crius/util"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
)

// Param 待确认套餐活动优惠券
type Param struct {
	ID    uuid.UUID `json:"id"`
	Code  string    `json:"code"`
	Title string    `json:"title"`
	Count int32     `json:"count"`
	Price int32     `json:"price"`
	Unit  string    `json:"unit"`
}

// Params ...
type Params []Param

// Value 序列化数据到数据库
func (p Params) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 反序列化数据到结构体
func (p *Params) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("scan error")
	}
	if err := json.Unmarshal(bytes, p); nil != err {
		return errors.New("json解析失败")
	}

	return nil
}

// Option 优惠方案
type Option struct {
	ID            uuid.UUID       `json:"id" gorm:"column:id"`
	Name          string          `json:"name" gorm:"column:name"`
	MerchantID    uuid.UUID       `json:"merchant_id" gorm:"column:merchant_id"`
	PromotionID   uuid.UUID       `json:"promotion_id" gorm:"column:promotion_id"`
	RechargeValue int64           `json:"recharge_value" gorm:"column:recharge_value"`
	BaseValue     int64           `json:"base_value" gorm:"column:base_value"`
	GiftValue     int64           `json:"gift_value" gorm:"column:gift_value"`
	Describe      string          `json:"describe" gorm:"describe"`
	Status        util.Status     `json:"status" gorm:"column:status"`
	Products      *Params         `json:"products" gorm:"column:products"`
	Packages      *Params         `json:"packages" gorm:"column:packages"`
	Tickets       *Params         `json:"tickets" gorm:"tickets"`
	TagID         *uuid.UUID      `json:"tag_id" gorm:"tag_id"`
	Extra         *pkgs.Params    `json:"extra" gorm:"column:extra"`
	CreatedAt     fields.NullTime `json:"created_at" form:"created_at" gorm:"column:created_at"`
	UpdatedAt     fields.NullTime `json:"updated_at" form:"updated_at" gorm:"column:updated_at"`
}

// TableName ...
func (o *Option) TableName() string {
	return `member_private.promotion_options`
}
