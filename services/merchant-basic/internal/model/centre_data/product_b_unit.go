package centreData

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TableProductUnit 售卖单位
type TableProductUnit struct {
	UnitID     uuid.UUID  `json:"unit_id" gorm:"column:unit_id"`
	UnitName   string     `json:"unit_name" gorm:"column:unit_name"`
	CreateTime *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time"`
	DeleteTime *time.Time `json:"delete_time" gorm:"column:delete_time"`
}

// TableName ...
func (TableProductUnit) TableName() string {
	return `centre_data.product_b_unit`
}
