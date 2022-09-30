package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// TableMerchant 商户
type TableMerchant struct {
	ID            uuid.UUID  `gorm:"column:id"`
	Name          string     `gorm:"column:name"`
	Logo          string     `gorm:"column:logo"`
	UserAgreement *Agreement `gorm:"column:user_agreement"`
}

// Agreement 协议
type Agreement struct {
	Agreement  string `json:"agreement"`
	FileFormat string `json:"file_format"`
}

// Value 将对象转换为数据库可存储类型
func (r *Agreement) Value() (driver.Value, error) {
	if nil == r {
		return nil, nil
	}

	return json.Marshal(r)
}

// Scan 将数据库对象转换成可以使用的golang 属性
func (r *Agreement) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("growth rules error")
	}

	if err := json.Unmarshal(bytes, r); nil != err {
		return errors.New("growth rules error json")
	}
	return nil
}

// TableName 表名
func (TableMerchant) TableName() string {
	return "merchant_basic.merchant"
}

// UpdateMerchantLogo 更新商户logo
func UpdateMerchantLogo(merchant *TableMerchant) error {
	return entity.Conn.Scopes(util.ColumnEqualScope("id", merchant.ID)).Select("logo").Updates(merchant).Error
}

// UpdateMerchantUserAgreement 更新商户用户协议
func UpdateMerchantUserAgreement(merchant *TableMerchant) error {
	return entity.Conn.Scopes(util.ColumnEqualScope("id", merchant.ID)).Select("user_agreement").Updates(merchant).Error
}

// ShowMerchant 查询商户信息
func ShowMerchant(id uuid.UUID) (*TableMerchant, error) {
	merchant := new(TableMerchant)
	err := entity.Conn.Scopes(util.ColumnEqualScope("id", id)).Take(merchant).Error
	return merchant, err
}
