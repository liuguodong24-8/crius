package centreData

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gorm.io/gorm"
	"time"
)

type TableProductFinanceType struct {
	FinanceTypeID uuid.UUID  `gorm:"column:finance_type_id"`
	ErpCode       string     `gorm:"column:erp_code"`
	Grade         int8       `gorm:"column:grade"`
	Code          string     `gorm:"column:code"`
	TypeName      string     `gorm:"column:type_name"`
	ParentID      *uuid.UUID `gorm:"column:parent_id"`
	MerchantID    *uuid.UUID `gorm:"column:merchant_id"`
	CreateTime    *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime    *time.Time `json:"update_time" gorm:"column:update_time"`
	DeleteTime    *time.Time `json:"delete_time" gorm:"column:delete_time"`
}

func (TableProductFinanceType) TableName() string {
	return "centre_data.product_b_finance_type"
}

// GetFinanceTypes list
func GetFinanceTypes(scopes []func(db *gorm.DB) *gorm.DB) ([]TableProductFinanceType, error) {
	var data []TableProductFinanceType
	err := model.DatabaseConn().Scopes(scopes...).Order("create_time desc").Find(&data).Error
	return data, err
}
