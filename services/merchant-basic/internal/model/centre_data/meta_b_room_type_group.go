package centreData

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
)

// TableMetaRoomTypeGroup 房型组
type TableMetaRoomTypeGroup struct {
	RoomTypeGroupID   uuid.UUID  `gorm:"column:room_type_group_id"`
	RoomTypeGroupName string     `gorm:"column:room_type_group_name"`
	RoomTypeIDs       string     `gorm:"column:room_type_ids"`
	CustomerCountMin  int32      `gorm:"column:customer_count_min"`
	CustomerCountMax  int32      `gorm:"column:customer_count_max"`
	Weight            int32      `gorm:"column:weight"`
	ShortName         string     `gorm:"column:short_name"`
	Type              string     `gorm:"column:type"`
	CreateTime        *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime        *time.Time `json:"update_time" gorm:"column:update_time"`
	DeleteTime        *time.Time `json:"delete_time" gorm:"column:delete_time"`
}

// TableName 表名
func (TableMetaRoomTypeGroup) TableName() string {
	return "centre_data.meta_b_room_type_group"
}

// GetRoomTypeGroups ...
func GetRoomTypeGroups() ([]TableMetaRoomTypeGroup, error) {
	var data []TableMetaRoomTypeGroup
	err := model.DatabaseConn().Order("create_time desc").Find(&data).Error
	return data, err
}

// GetRoomTypeGroupsByRoomTypeID ...
func GetRoomTypeGroupsByRoomTypeID(id uuid.UUID) ([]TableMetaRoomTypeGroup, error) {
	var data []TableMetaRoomTypeGroup
	err := model.DatabaseConn().Scopes(util.ArrayAnyScope("room_type_ids", id)).Find(&data).Error
	return data, err
}
