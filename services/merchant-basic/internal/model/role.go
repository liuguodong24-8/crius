package model

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// TableRole 角色表结构
type TableRole struct {
	ID         uuid.UUID      `gorm:"column:id"`
	Name       string         `gorm:"column:name"`
	Status     crius.Status   `gorm:"column:status"`
	Property   int8           `gorm:"column:property"`
	CreatedAt  *time.Time     `gorm:"column:created_at"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
	StaffID    *uuid.UUID     `gorm:"column:staff_id"`
	MerchantID *uuid.UUID     `gorm:"column:merchant_id"`
}

// RoleInfo 角色联合员工名称结构
type RoleInfo struct {
	TableRole
	StaffName string
}

// TableMappingRolePermission 角色权限关系表
type TableMappingRolePermission struct {
	ID           uuid.UUID `grom:"column:id"`
	RoleID       uuid.UUID `gorm:"role_id"`
	PermissionID int32     `gorm:"permission_id"`
}

// TableName 角色表表名
func (TableRole) TableName() string {
	return "merchant_basic.role"
}

// TableName 角色权限关系表名
func (TableMappingRolePermission) TableName() string {
	return "merchant_basic.mapping_role_permission"
}

// CreateRole 创建角色
func CreateRole(role TableRole, mappingPermissions []TableMappingRolePermission) error {
	tx := entity.Conn.Begin()
	err := tx.Create(&role).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Create(&mappingPermissions).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// UpdateRole 更新角色
func UpdateRole(role TableRole, mappingPermissions []TableMappingRolePermission) error {
	tx := entity.Conn.Begin()
	err := tx.Scopes(crius.ColumnEqualScope("id", role.ID)).Updates(&role).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Scopes(crius.ColumnEqualScope("role_id", role.ID)).Delete(&TableMappingRolePermission{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Create(&mappingPermissions).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// UpdateRoleStatus 更新角色状态
func UpdateRoleStatus(id uuid.UUID, status string) error {
	return entity.Conn.Model(&TableRole{}).Where("id = ?", id).Update("status", status).Error
}

// GetRoles 获取角色
func GetRoles(name, status string, staffID, merchantID uuid.UUID, offset, limit int32) ([]RoleInfo, int64, error) {
	var roles []RoleInfo
	var count int64
	db := entity.Conn.Debug().Model(&TableRole{}).Select("merchant_basic.role.*", "merchant_basic.staff.name as staff_name").
		Joins("left join merchant_basic.staff on role.staff_id = staff.id").Where("role.merchant_id = ?", merchantID)
	if name != "" {
		db = db.Where("role.name like ?", fmt.Sprintf("%%%s%%", name))
	}
	if status != "" {
		db = db.Where("role.status = ?", status)
	}
	if staffID != uuid.Nil {
		db.Where("(staff.id = ? and property = ?) or property = ?", staffID, 2, 1)
	}
	err := db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return roles, 0, nil
	}
	err = db.Scopes(pagingCondition(offset, limit)).Order("created_at DESC").Find(&roles).Error
	return roles, count, err
}

// DeleteRole 删除角色
func DeleteRole(id uuid.UUID) error {
	return entity.Conn.Where("id = ?", id).Delete(&TableRole{}).Error
}

// ShowRoleByName 根据名字查询角色
func ShowRoleByName(name string, merchantID uuid.UUID) (*RoleInfo, []TableMappingRolePermission, error) {
	if name == "" {
		return nil, nil, gorm.ErrRecordNotFound
	}
	role := new(RoleInfo)
	err := entity.Conn.Model(&TableRole{}).Select("merchant_basic.role.*", "merchant_basic.staff.name as staff_name").
		Joins("left join merchant_basic.staff on role.staff_id = staff.id").Where("role.name = ? and role.merchant_id = ?", name, merchantID).Take(&role).Error
	if err != nil {
		return nil, nil, err
	}

	permissions := make([]TableMappingRolePermission, 0)
	if err := entity.Conn.Scopes(crius.ColumnEqualScope("role_id", role.ID)).Find(&permissions).Error; err != nil {
		return nil, nil, err
	}
	return role, permissions, nil
}

// ShowRoleByID 根据ID查询角色
func ShowRoleByID(id uuid.UUID) (*RoleInfo, []TableMappingRolePermission, error) {
	role := new(RoleInfo)
	err := entity.Conn.Model(&TableRole{}).Select("merchant_basic.role.*", "merchant_basic.staff.name as staff_name").
		Joins("left join merchant_basic.staff on role.staff_id = staff.id").Where("role.id = ?", id).Take(&role).Error
	if err != nil {
		return nil, nil, err
	}

	permissions := make([]TableMappingRolePermission, 0)
	if err := entity.Conn.Scopes(crius.ColumnEqualScope("role_id", id)).Find(&permissions).Error; err != nil {
		return nil, nil, err
	}
	return role, permissions, nil
}
