package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// TableSnapshot 快照表
type TableSnapshot struct {
	ID                uuid.UUID    `gorm:"column:id"`
	StaffID           *uuid.UUID   `gorm:"column:staff_id"`
	SleuthCode        string       `gorm:"column:sleuth_code"`
	SnapShotTableName string       `gorm:"column:table_name"`
	TableID           *uuid.UUID   `gorm:"column:table_id"`
	Method            string       `gorm:"column:method"`
	CreatedAt         *time.Time   `gorm:"created_at"`
	Before            *pkgs.Params `gorm:"column:before"`
	After             *pkgs.Params `gorm:"column:after"`
}

// SnapShotInfo 快照表查询信息
type SnapShotInfo struct {
	StaffID   uuid.UUID
	StaffName string
	Before    pkgs.Params
	After     pkgs.Params
	CreatedAt time.Time
	Method    string
}

// TableName 数据库表名
func (t TableSnapshot) TableName() string {
	return "member_account.snapshot"
}

// CreateSnapshot 创建快照
func CreateSnapshot(snapshot TableSnapshot) error {
	return entity.Conn.Create(&snapshot).Error
}
