package model

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// Promotion 优惠方案组
type Promotion struct {
	ID         uuid.UUID        `json:"id" gorm:"column:id"`
	Name       string           `json:"name" gorm:"column:name"`
	MerchantID uuid.UUID        `json:"merchant_id" gorm:"column:merchant_id"`
	BeginAt    *fields.DateTime `json:"begin_at" gorm:"column:begin_at"`
	EndAt      *fields.DateTime `json:"end_at" gorm:"column:end_at"`
	Status     util.Status      `json:"status" gorm:"column:status"`
	BranchIds  fields.UUIDArr   `json:"branch_ids" gorm:"column:branch_ids"`
	Extra      *pkgs.Params     `json:"extra" gorm:"column:extra"`
	CreatedAt  fields.NullTime  `json:"created_at" form:"created_at" gorm:"column:created_at"`
	UpdatedAt  fields.NullTime  `json:"updated_at" form:"updated_at" gorm:"column:updated_at"`
}

// TableName ...
func (p *Promotion) TableName() string {
	return `member_private.promotions`
}
