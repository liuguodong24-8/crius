package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// WechatTemplate 微信模版
type WechatTemplate struct {
	ID              uuid.UUID              `json:"id" gorm:"primary_key"`
	MerchantID      uuid.UUID              `json:"merchant_id"`
	TemplateName    string                 `json:"template_name"`
	TemplateCode    string                 `json:"template_code"`
	TemplateContent *WechatTemplateContent `json:"template_content"`
	OfficialLink    string                 `json:"official_link"`
	MiniprogramLink string                 `json:"miniprogram_link"`
	Category        string                 `json:"category"`
	CategoryKey     string                 `json:"category_key"`
	Extra           *pkgs.Params           `json:"extra"`
	Status          util.Status            `json:"status"`
	CreatedAt       pkgs.NullTime          `json:"created_at" form:"created_at"` // 统一定义时间字段处理，格式化
	UpdatedAt       pkgs.NullTime          `json:"updated_at" form:"updated_at"` // 统一定义时间字段处理，格式化
}

// TableName 表名
func (w WechatTemplate) TableName() string {
	return "message.wechat_templates"
}

// WechatTemplateContentBase 基础配置
type WechatTemplateContentBase struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// WechatTemplateContentDetail 详情配置
type WechatTemplateContentDetail struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Color string `json:"color"`
}

// WechatTemplateContent 模版内容
type WechatTemplateContent struct {
	First  WechatTemplateContentBase     `json:"first"`
	Detail []WechatTemplateContentDetail `json:"detail"`
	Remark WechatTemplateContentBase     `json:"remark"`
}

//Value 转换json存储
func (w WechatTemplateContent) Value() (driver.Value, error) {
	return json.Marshal(w)
}

//Scan 将数据库对象转换成可以使用的golang 属性
func (w *WechatTemplateContent) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("解析微信模版内容错误")
	}

	if err := json.Unmarshal(bytes, w); nil != err {
		return errors.New("template_content json解析失败")
	}

	return nil
}
