package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// ConsumeCategory 消费类型
type ConsumeCategory struct {
	ID            uuid.UUID       `json:"id"`
	MerchantID    *uuid.UUID      `json:"merchant_id"`
	Category      string          `json:"category"`
	Code          string          `json:"code"`
	Status        util.Status     `json:"status"`
	Extra         *pkgs.Params    `json:"extra"`
	OperatorTypes *fields.UUIDArr `json:"operator_types" gorm:"column:operator_types"`
	ActiveTypes   *fields.UUIDArr `json:"active_types" gorm:"column:active_types"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     *time.Time      `json:"deleted_at"`
	IsRoomFee     bool            `gorm:"column:is_room_fee"`
}

// TableName table name
func (ConsumeCategory) TableName() string {
	return `merchant_basic.consume_category`
}

//GetConsumeCategoriesByCodes code换id
func GetConsumeCategoriesByCodes(codes []string, merchantID uuid.UUID) ([]ConsumeCategory, error) {
	var categories []ConsumeCategory
	err := entity.Conn.Select("id,code").
		Scopes(
			util.ColumnEqualScope("merchant_id", merchantID),
			util.ColumnEqualScope("status", StatusOpened),
		).
		Where("code in ?", codes).
		Find(&categories).Error
	return categories, err
}

// GetConsumeCategoryTypes 获取消费分类分组id
func GetConsumeCategoryTypes(merchantID uuid.UUID) ([]ConsumeCategory, error) {
	categories := make([]ConsumeCategory, 0)
	err := entity.Conn.Select("id", "operator_types", "active_types", "category", "is_room_fee").Scopes(util.ColumnEqualScope("merchant_id", merchantID)).Find(&categories).Error
	return categories, err
}

// ShowConsumeCategoryRoomFee 获取房费消费类型
func ShowConsumeCategoryRoomFee() (*ConsumeCategory, error) {
	category := new(ConsumeCategory)
	err := entity.Conn.Scopes(util.ColumnEqualScope("is_room_fee", true)).Take(category).Error
	return category, err
}
