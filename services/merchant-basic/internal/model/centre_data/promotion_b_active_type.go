package centreData

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gorm.io/gorm"
	"time"
)

type TablePromotionActiveType struct {
	ActiveTypeID uuid.UUID  `gorm:"column:active_type_id"`
	Grade        int8       `gorm:"column:grade"`
	Code         string     `gorm:"column:code"`
	TypeName     string     `gorm:"column:type_name"`
	ParentID     *uuid.UUID `gorm:"column:parent_id"`
	MerchantID   *uuid.UUID `gorm:"column:merchant_id"`
	CreateTime   *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime   *time.Time `json:"update_time" gorm:"column:update_time"`
	DeleteTime   *time.Time `json:"delete_time" gorm:"column:delete_time"`
}

func (TablePromotionActiveType) TableName() string {
	return "centre_data.promotion_b_active_type"
}

// CountActiveType count
func CountActiveType(scopes []func(db *gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := model.DatabaseConn().Model(&TablePromotionActiveType{}).Scopes(scopes...).Count(&count).Error
	return count, err
}

// GetActiveTypes list
func GetActiveTypes(scopes []func(db *gorm.DB) *gorm.DB) ([]TablePromotionActiveType, error) {
	var data []TablePromotionActiveType
	err := model.DatabaseConn().Scopes(scopes...).Order("create_time desc").Find(&data).Error
	return data, err
}
