package model

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// TableWechatPay 表结构
type TableWechatPay struct {
	MerchantID           uuid.UUID  `gorm:"column:merchant_id"`
	AppID                string     `gorm:"column:app_id"`
	MchID                string     `gorm:"column:mch_id"`
	HeadquartersSubMchID string     `gorm:"column:headquarters_sub_mch_id"` //总部子商户号
	PrivateKey           string     `gorm:"column:private_key"`
	CertFilename         string     `gorm:"column:cert_filename"`
	CertContent          []byte     `gorm:"column:cert_content"`
	CreatedAt            *time.Time `gorm:"column:created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at"`
}

// TableName TableName
func (t *TableWechatPay) TableName() string {
	return "merchant_basic.wechat_payment"
}

// GetWechatPaySetting 获取商户的支付配置
func GetWechatPaySetting(merchantID uuid.UUID) (*TableWechatPay, error) {
	var setting TableWechatPay
	err := entity.Conn.Scopes(crius.ColumnEqualScope("merchant_id", merchantID)).First(&setting).Error
	return &setting, err
}

// GetWechatPaySettingByAppID 获取微信支付信息
func GetWechatPaySettingByAppID(appID string) (*TableWechatPay, error) {
	var setting TableWechatPay
	err := entity.Conn.Scopes(crius.ColumnEqualScope("app_id", appID)).First(&setting).Error
	return &setting, err
}

// SaveOrCreateWechatPaySetting 保存微信支付信息
func SaveOrCreateWechatPaySetting(setting *TableWechatPay) error {
	var dbData TableWechatPay
	if err := entity.Conn.Scopes(crius.ColumnEqualScope("merchant_id", setting.MerchantID)).First(&dbData).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//inert
			return entity.Conn.Create(setting).Error
		}
		return err
	}
	//save
	return entity.Conn.Scopes(crius.ColumnEqualScope("merchant_id", setting.MerchantID)).Updates(setting).Error
}
