package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	cutil "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// TableBranchTag 表结构
type TableBranchTag struct {
	ID            uuid.UUID       `gorm:"column:id"`
	Name          string          `gorm:"column:name"`
	Branches      *fields.UUIDArr `gorm:"column:branches"`
	CreateStaffID *uuid.UUID      `gorm:"column:create_staff_id"`
	Status        crius.Status    `gorm:"column:status"`
	Extra         *pkgs.Params    `gorm:"column:extra"`
	MerchantID    *uuid.UUID      `gorm:"column:merchant_id"`
	CreatedAt     *time.Time      `gorm:"column:created_at"`
	UpdatedAt     *time.Time      `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at"`
	LoadExtra     string          `gorm:"-"` // 导入信息 不处理
}

// BranchTag 返回表结构
type BranchTag struct {
	ID            *uuid.UUID      `gorm:"column:id"`
	Name          string          `gorm:"column:name"`
	Branches      *fields.UUIDArr `gorm:"column:branches"`
	CreateStaffID *uuid.UUID      `gorm:"column:create_staff_id"`
	Status        cutil.Status    `gorm:"column:status"`
	Extra         *pkgs.Params    `gorm:"column:extra"`
	MerchantID    uuid.UUID       `gorm:"column:merchant_id"`
	CreatedAt     *time.Time      `gorm:"column:created_at"`
	UpdatedAt     *time.Time      `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at"`
	StaffName     string          `gorm:"column:staff_name"`
}

// TableName 表名
func (TableBranchTag) TableName() string {
	return "merchant_basic.branch_tag"
}

// DirectlyBranchTag 直营标签ID
func DirectlyBranchTag() uuid.UUID {
	return uuid.FromStringOrNil(`5260754d-e100-4fec-8f66-ff9d2a5e3545`)
}

// CreateBranchTag 创建标签
func CreateBranchTag(tag TableBranchTag) error {
	return entity.Conn.Create(&tag).Error
}

// UpdateBranchTag 更新标签
func UpdateBranchTag(tag TableBranchTag) error {
	return entity.Conn.Scopes(crius.ColumnEqualScope("id", tag.ID)).Updates(&tag).Error
}

// UpdateBranchTagStatus 更新标签状态
func UpdateBranchTagStatus(id uuid.UUID, status string) error {
	return entity.Conn.Model(&TableBranchTag{}).Scopes(crius.ColumnEqualScope("id", id)).Update("status", status).Error
}

// GetBranchTags 获取标签列表
func GetBranchTags(tag TableBranchTag, dateStart, dateEnd, offset, limit int32) ([]BranchTag, int64, error) {
	var count int64
	tags := make([]BranchTag, 0)
	db := entity.Conn.Model(&TableBranchTag{}).Select("branch_tag.*", "staff.name as staff_name").Joins("left join merchant_basic.staff on branch_tag.create_staff_id = staff.id")
	if tag.Name != "" {
		db = db.Where("branch_tag.name = ?", tag.Name)
	}
	if tag.Status != "" {
		db = db.Where("branch_tag.status = ?", tag.Status)
	}
	if dateEnd >= dateStart && dateEnd > 0 {
		db = db.Where("branch_tag.created_at >= ?", getTodayDate(int64(dateStart))).Where("branch_tag.created_at < ?", getTomorrowDate(int64(dateEnd)))
	}
	if len(tag.Branches.Slice()) > 0 {
		ids := make([]interface{}, len(tag.Branches.Slice()))
		for i := range tag.Branches.Slice() {
			ids[i] = tag.Branches.Slice()[i]
		}
		db = db.Scopes(crius.ArrayOverlapScope("branch_tag.branches", "uuid", ids))
	}

	err := db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return tags, 0, nil
	}
	err = db.Order("updated_at desc").Scan(&tags).Error
	if err != nil {
		return nil, 0, err
	}

	return tags, count, nil
}

// GetBranchTagsByIDs 根据id列表获取标签列表
func GetBranchTagsByIDs(ids []uuid.UUID) ([]BranchTag, error) {
	tags := make([]BranchTag, 0)
	db := entity.Conn.Model(&TableBranchTag{}).Select("branch_tag.*", "staff.name as staff_name").Joins("left join merchant_basic.staff on branch_tag.create_staff_id = staff.id")
	err := db.Where("branch_tag.id in ?", ids).Find(&tags).Error
	return tags, err
}

// ShowBranchTagByName 根据名称查询标签
func ShowBranchTagByName(name string, merchantID uuid.UUID) (*TableBranchTag, error) {
	tag := new(TableBranchTag)
	if name == "" {
		return nil, gorm.ErrRecordNotFound
	}
	err := entity.Conn.Scopes(crius.ColumnEqualScopeDefault("name", name), crius.ColumnEqualScope("merchant_id", merchantID)).Take(&tag).Error
	return tag, err
}

// ShowBranchTag 查询单条标签
func ShowBranchTag(id uuid.UUID) (*BranchTag, error) {
	tag := new(BranchTag)
	err := entity.Conn.Model(&TableBranchTag{}).Select("branch_tag.*", "staff.name as staff_name").
		Joins("left join merchant_basic.staff on branch_tag.create_staff_id = staff.id").
		Where("branch_tag.id = ?", id).Take(&tag).Error
	return tag, err
}
