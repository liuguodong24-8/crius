package model

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// TableBranchAppointment 门店预约表
type TableBranchAppointment struct {
	OpenAppointment        int8              `gorm:"column:open_appointment"`
	AppointmentGranularity int8              `gorm:"column:appointment_granularity"`
	VR                     *pkgs.ParamsArr   `gorm:"column:vr"`
	Video                  *fields.StringArr `gorm:"column:video"`
	Environment            *fields.StringArr `gorm:"column:environment"`
	Meal                   *fields.StringArr `gorm:"column:meal"`
	Price                  *fields.StringArr `gorm:"column:price"`
	RoomTypes              *pkgs.ParamsArr   `gorm:"column:room_types"`
	Hot                    bool              `gorm:"column:hot"`
	BranchID               *uuid.UUID        `gorm:"column:branch_id"`
}

// TableName 表名
func (TableBranchAppointment) TableName() string {
	return "merchant_basic.branch_appointment"
}

// SaveBranchAppointment 更新
func SaveBranchAppointment(appointment TableBranchAppointment) error {
	tx := entity.Conn.Begin()
	if err := tx.Select("open_appointment", "appointment_granularity", "vr", "video", "environment", "meal", "price", "hot", "branch_id").
		Scopes(crius.ColumnEqualScopeDefault("branch_id", *appointment.BranchID)).Updates(&appointment).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ShowBranchAppointment 查询单条
func ShowBranchAppointment(branchID uuid.UUID) (*TableBranchAppointment, error) {
	appointment := new(TableBranchAppointment)
	if branchID == uuid.Nil {
		return nil, gorm.ErrRecordNotFound
	}
	err := entity.Conn.Scopes(crius.ColumnEqualScopeDefault("branch_id", branchID)).Take(appointment).Error
	return appointment, err
}

// UpdateBranchAppointmentRoomType 更新门店关联信息
func UpdateBranchAppointmentRoomType(appointment TableBranchAppointment) error {
	return entity.Conn.Model(&TableBranchAppointment{}).Scopes(crius.ColumnEqualScopeDefault("branch_id", *appointment.BranchID)).Updates(&appointment).Error
}
