package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// SnapshotMethod 操作类型
type SnapshotMethod string

// String ...
func (s SnapshotMethod) String() string {
	return string(s)
}

const (
	// UpdateMethod 修改
	UpdateMethod SnapshotMethod = `update`

	// CreateMethod 新增
	CreateMethod SnapshotMethod = `create`
)

// Snapshot 快照表
type Snapshot struct {
	ID                uuid.UUID      `gorm:"column:id"`
	StaffID           uuid.UUID      `gorm:"column:staff_id"`
	SleuthCode        string         `gorm:"column:sleuth_code"`
	SnapShotTableName string         `gorm:"column:table_name"`
	TableID           uuid.UUID      `gorm:"column:table_id"`
	Method            SnapshotMethod `gorm:"column:method"`
	Before            *pkgs.Params   `gorm:"column:before"`
	After             *pkgs.Params   `gorm:"column:after"`
	CreatedAt         *time.Time     `gorm:"created_at"`
}

// TableName 数据库表名
func (s Snapshot) TableName() string {
	return "member_private.snapshot"
}
