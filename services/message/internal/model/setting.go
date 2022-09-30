package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/cyrnicolase/nulls"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// MessageSetting 发送设置
type MessageSetting struct {
	ID               uuid.UUID       `json:"id" gorm:"primary_key"`
	MerchantID       uuid.UUID       `json:"merchant_id"`
	MessageType      string          `json:"message_type"`
	TriggerType      string          `json:"trigger_type"`
	AdvanceHour      float64         `json:"advance_hour"`
	SmsTemplateID    nulls.UUID      `json:"sms_template_id"`
	WechatTemplateID nulls.UUID      `json:"wechat_template_id"`
	SpecialSetting   *SpecialSetting `json:"special_setting"`
	CcList           *CcList         `json:"cc_list"`
	SpecialBranches  *fields.UUIDArr `json:"special_branches"`
	Extra            *pkgs.Params    `json:"extra"`
	Status           util.Status     `json:"status"`
	CreatedAt        pkgs.NullTime   `json:"created_at" form:"created_at"` // 统一定义时间字段处理，格式化
	UpdatedAt        pkgs.NullTime   `json:"updated_at" form:"updated_at"` // 统一定义时间字段处理，格式化
}

// TableName 表名
func (s MessageSetting) TableName() string {
	return "message.message_setting"
}

// CcList 抄送列表
type CcList []Cc

// Cc 抄送
type Cc struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

// Value 返回数据库可识别类型
func (p CcList) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan ...
func (p *CcList) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("cc list value")
	}

	if err := json.Unmarshal(bytes, p); nil != err {
		return errors.New("cc list value error")
	}

	return nil
}

// SpecialSetting 特殊设置
type SpecialSetting []RangeSetting

// RangeSetting 区间特殊设置
type RangeSetting struct {
	Begin              fields.DateTime `json:"begin"`
	End                fields.DateTime `json:"end"`
	SmsTemplateID      nulls.UUID      `json:"sms_template_id"`
	SmsTemplateName    string          `json:"sms_template_name"`
	WechatTemplateID   nulls.UUID      `json:"wechat_template_id"`
	WechatTemplateName string          `json:"wechat_template_name"`
}

// Value 返回数据库可识别类型
func (p SpecialSetting) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan ...
func (p *SpecialSetting) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("special setting value")
	}

	if err := json.Unmarshal(bytes, p); nil != err {
		return errors.New("special setting value error")
	}

	return nil
}
