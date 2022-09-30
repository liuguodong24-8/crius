package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// TableRoomType 表结构
type TableRoomType struct {
	ID               uuid.UUID       `gorm:"column:id"`
	BranchID         *uuid.UUID      `gorm:"column:branch_id"`
	CategoryID       *uuid.UUID      `gorm:"column:category_id"`
	Name             string          `gorm:"column:name"`
	Category         int8            `gorm:"column:category"`
	Status           crius.Status    `gorm:"column:status"`
	MerchantID       *uuid.UUID      `gorm:"column:merchant_id"`
	CustomerMin      int8            `gorm:"column:customer_min"`
	CustomerMax      int8            `gorm:"column:customer_max"`
	Order            int32           `gorm:"column:order"`
	RoomTypeGroupIDs *fields.UUIDArr `gorm:"column:room_type_group_ids"`
	CreatedAt        *time.Time      `gorm:"column:created_at"`
	UpdatedAt        *time.Time      `gorm:"column:updated_at"`
	LoadExtra        string          `gorm:"-"` // 导入信息 不处理
}

// RoomType 返回数据结构
type RoomType struct {
	TableRoomType
	CategoryName string
}

// TableName 表名
func (TableRoomType) TableName() string {
	return "merchant_basic.room_type"
}

// ShowRoomTypeByName 查询单条
func ShowRoomTypeByName(name string, merchantID uuid.UUID) (*TableRoomType, error) {
	var roomType TableRoomType
	if name == "" {
		return nil, gorm.ErrRecordNotFound
	}
	err := entity.Conn.Debug().Scopes(crius.ColumnEqualScopeDefault("name", name), crius.ColumnEqualScope("merchant_id", merchantID)).Take(&roomType).Error
	return &roomType, err
}

// CreateRoomType 创建
func CreateRoomType(t TableRoomType) error {
	return entity.Conn.Create(&t).Error
}

// GetRoomTypes 获取列表
func GetRoomTypes(t TableRoomType, offset, limit int32) ([]RoomType, int64, error) {
	var count int64
	roomTypes := make([]RoomType, 0)
	db := entity.Conn.Select("room_type.*", "room_type_category.name as category_name").Model(&TableRoomType{}).Joins("left join merchant_basic.room_type_category on room_type.category_id = room_type_category.id")
	if t.BranchID != nil {
		db.Where("room_type.branch_id = ? and room_type.status = ? and room_type_category.status = ?", t.BranchID, StatusOpened, StatusOpened)
	}
	if t.CategoryID != nil {
		db.Where("room_type.category_id = ?", t.CategoryID)
	}
	if t.Name != "" {
		db = db.Where("room_type.name = ?", t.Name)
	}
	if t.Status != "" {
		db = db.Where("room_type.status = ?", t.Status)
	}

	db = db.Where("room_type.merchant_id = ?", t.MerchantID)
	err := db.Model(&TableRoomType{}).Count(&count).Error
	if err != nil {
		return nil, count, err
	}
	if count == 0 {
		return roomTypes, 0, nil
	}
	err = db.Scopes(pagingCondition(offset, limit)).Order(`"order" desc`).Find(&roomTypes).Error
	if err != nil {
		return nil, count, err
	}
	return roomTypes, count, nil
}

// GetRoomTypesByIDs 根据id获取房型列表
func GetRoomTypesByIDs(ids []uuid.UUID) ([]RoomType, error) {
	var roomTypes []RoomType
	var idsInterface []interface{}
	for _, id := range ids {
		idsInterface = append(idsInterface, id)
	}
	err := entity.Conn.Select("room_type.*", "room_type_category.name as category_name").Model(&TableRoomType{}).Joins("left join merchant_basic.room_type_category on room_type.category_id = room_type_category.id").
		Scopes(crius.ColumnInScope("room_type.id", idsInterface)).Order(`"order" desc`).Find(&roomTypes).Error
	return roomTypes, err
}

// UpdateRoomType 更新
func UpdateRoomType(t TableRoomType) error {
	return entity.Conn.Select("name", "status", "merchant_id", "customer_min", "customer_max", "order", "category_id", "room_type_group_ids").Scopes(crius.ColumnEqualScope("id", t.ID)).Updates(&t).Error
}

// UpdateRoomTypeStatus 更新状态
func UpdateRoomTypeStatus(id uuid.UUID, status crius.Status) error {
	return entity.Conn.Model(&TableRoomType{}).Scopes(crius.ColumnEqualScope("id", id)).UpdateColumn("status", status).Error
}

// ShowRoomType 获取房型
func ShowRoomType(id uuid.UUID) (*RoomType, error) {
	room := new(RoomType)
	err := entity.Conn.Select("room_type.*", "room_type_category.name as category_name").Model(&TableRoomType{}).
		Joins("left join merchant_basic.room_type_category on room_type.category_id = room_type_category.id").
		Where("room_type.id = ?", id).Take(room).Error
	return room, err
}
