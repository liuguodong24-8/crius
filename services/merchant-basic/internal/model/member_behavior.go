package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
)

//TableMemberBehavior 会员行为表
type TableMemberBehavior struct {
	ID         uuid.UUID  `gorm:"column:id"`
	MemberID   uuid.UUID  `gorm:"column:member_id"`
	Behavior   string     `gorm:"column:behavior"`
	StaffID    *uuid.UUID `gorm:"column:staff_id"`
	BranchID   *uuid.UUID `gorm:"column:branch_id"`
	MerchantID uuid.UUID  `gorm:"column:merchant_id"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
}

//TableName 表名
func (TableMemberBehavior) TableName() string {
	return "merchant_basic.member_behavior"
}

//CreateMemberBehavior 新建
func CreateMemberBehavior(table *TableMemberBehavior) error {
	return entity.Conn.Create(&table).Error
}

//GetMemberBehaviors 用户行为列表
func GetMemberBehaviors(memberID uuid.UUID, offset, limit int32) ([]TableMemberBehavior, error) {
	var behaviors []TableMemberBehavior
	err := entity.Conn.Scopes(
		crius.ColumnEqualScope("member_id", memberID),
		pagingCondition(offset, limit),
	).Order("created_at desc").
		Find(&behaviors).Error

	return behaviors, err
}

//GetMemberBehaviorsCount 用户行为数
func GetMemberBehaviorsCount(memberID uuid.UUID) (int64, error) {
	var count int64
	err := entity.Conn.Model(&TableMemberBehavior{}).Scopes(
		crius.ColumnEqualScope("member_id", memberID),
	).Count(&count).Error
	return count, err
}
