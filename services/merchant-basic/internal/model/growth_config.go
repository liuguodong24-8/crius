package model

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gorm.io/gorm/clause"
)

// TableGrowthConfig 表结构
type TableGrowthConfig struct {
	ID         uuid.UUID         `gorm:"column:id"`
	MerchantID uuid.UUID         `gorm:"column:merchant_id"`
	Name       string            `gorm:"column:name"`
	Top        uint32            `gorm:"column:top"`
	Rules      *fields.StringArr `gorm:"column:rules"`
}

// TableName 表名
func (TableGrowthConfig) TableName() string {
	return "merchant_basic.growth_config"
}

// ShowGrowConfigByMerchantID 根据商户查询成长值配置
func ShowGrowConfigByMerchantID(merchantID uuid.UUID) (TableGrowthConfig, error) {
	var config TableGrowthConfig
	err := entity.Conn.Where("merchant_id", merchantID).First(&config).Error
	return config, err
}

// CreateOrUpdateGrowthConfig 设置成长值配置
func CreateOrUpdateGrowthConfig(config TableGrowthConfig) error {
	return entity.Conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "merchant_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "top", "rules"}),
	}).Create(&config).Error
}
