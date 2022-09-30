package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TablePermission 数据库权限表
type TablePermission struct {
	ID         int32      `gorm:"column:id"`
	Permission string     `gorm:"column:permission"`
	Service    string     `gorm:"column:service"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
}

// TableName 权限表名
func (TablePermission) TableName() string {
	return "merchant_basic.permission"
}

// ShowPermissionExistsByID 查询权限是否存在
func ShowPermissionExistsByID(id int32) error {
	return entity.Conn.Unscoped().Where("id = ?", id).Take(&TablePermission{}).Error
}

// UpdatePermission 更新权限
func UpdatePermission(p TablePermission) error {
	return entity.Conn.Unscoped().Select("id", "permission", "service", "deleted_at").Where("id = ?", p.ID).Updates(&p).Error
}

// CreatePermission 新增权限
func CreatePermission(p TablePermission) error {
	return entity.Conn.Create(&p).Error
}

// GetPermissionsByStaffID 获取权限
func GetPermissionsByStaffID(id uuid.UUID, service string) ([]TablePermission, error) {
	var permissions []TablePermission
	db := entity.Conn
	if id != uuid.Nil {
		subQuery := db.Select("role_id").Model(&TableRole{}).Joins("inner join merchant_basic.mapping_staff_role as msr on role.id = msr.role_id").Where("msr.staff_id = ?", id)
		db = db.Model(&TablePermission{}).Distinct().Joins("inner join merchant_basic.mapping_role_permission on permission.id = mapping_role_permission.permission_id").
			Where("role_id in (?) and service = ?", subQuery, service)
	}
	err := db.Find(&permissions).Error
	return permissions, err
}

// GetPermissionsByID 获取权限
func GetPermissionsByID(id []int32) ([]TablePermission, error) {
	var permissions []TablePermission
	err := entity.Conn.Where("id in ?", id).Find(&permissions).Error
	return permissions, err
}

// SavePermissions 删除旧权限，新增权限
func SavePermissions(permissions []TablePermission, service string) error {
	tx := entity.Conn.Begin()
	err := tx.Where("service = ?", service).Delete(&TablePermission{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Create(&permissions).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
