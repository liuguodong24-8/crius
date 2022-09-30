package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TableMemberAddress 用户地址
type TableMemberAddress struct {
	ID         uuid.UUID `json:"id"`
	MemberID   uuid.UUID `json:"member_id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	PhoneCode  string    `json:"phone_code"`
	ProvinceID string    `json:"province_id"`
	CityID     string    `json:"city_id"`
	DistrictID string    `json:"district_id"`
	Address    string    `json:"address"`
	IsDefault  bool      `json:"is_default"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName table name
func (TableMemberAddress) TableName() string {
	return `merchant_basic.member_address`
}
