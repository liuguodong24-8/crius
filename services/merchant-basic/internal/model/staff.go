package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// StaffFullInfo 员工信息，包含所属门店，角色权限
type StaffFullInfo struct {
	Staff    TableStaff
	Branches []Branch
	Roles    []RoleInfo
}

// TableStaff staff表
type TableStaff struct {
	ID           uuid.UUID      `gorm:"column:id"`
	EmployeeCode string         `gorm:"column:employee_code"`
	Name         string         `gorm:"column:name"`
	Phone        string         `gorm:"column:phone"`
	PhoneCode    string         `gorm:"column:phone_code"`
	Gender       int8           `gorm:"column:gender"`
	Status       crius.Status   `gorm:"column:status"`
	Code         string         `gorm:"column:code"`
	EntryAt      *time.Time     `gorm:"column:entry_at"`
	QuitAt       *time.Time     `gorm:"column:quit_at"`
	Password     string         `gorm:"column:password"`
	Salt         string         `gorm:"column:salt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	CreatedAt    *time.Time     `gorm:"column:created_at"`
	UpdatedAt    *time.Time     `gorm:"column:updated_at"`
	Admin        bool           `gorm:"column:admin"`
	MerchantID   *uuid.UUID     `gorm:"column:merchant_id"`
	LoadExtra    string         `gorm:"-"` // 导入信息 不处理
}

// TableName 指定表名
func (TableStaff) TableName() string {
	return "merchant_basic.staff"
}

// TableMappingStaffBranch 员工门店中间表
type TableMappingStaffBranch struct {
	ID       uuid.UUID `gorm:"column:id"`
	StaffID  uuid.UUID `gorm:"column:staff_id"`
	BranchID uuid.UUID `gorm:"column:branch_id"`
}

// TableName 指定表名
func (TableMappingStaffBranch) TableName() string {
	return "merchant_basic.mapping_staff_branch"
}

// TableMappingStaffRole 员工角色中间表
type TableMappingStaffRole struct {
	ID      uuid.UUID `gorm:"column:id"`
	StaffID uuid.UUID `gorm:"column:staff_id"`
	RoleID  uuid.UUID `gorm:"column:role_id"`
}

// TableName 指定表名
func (TableMappingStaffRole) TableName() string {
	return "merchant_basic.mapping_staff_role"
}

// CreateStaff 新增员工
func CreateStaff(staff TableStaff, branches []TableMappingStaffBranch, roles []TableMappingStaffRole) error {
	tx := entity.Conn.Begin()
	if err := tx.Create(&staff).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(branches) != 0 {
		if err := tx.Create(&branches).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(roles) != 0 {
		if err := tx.Create(&roles).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// UpdateStaff 更新员工信息
func UpdateStaff(staff TableStaff, branches []TableMappingStaffBranch, roles []TableMappingStaffRole) error {
	tx := entity.Conn.Begin()
	if err := tx.Select("name", "phone", "phone_code", "gender", "entry_at", "quit_at").Scopes(crius.ColumnEqualScope("id", staff.ID)).Updates(&staff).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Scopes(crius.ColumnEqualScope("staff_id", staff.ID)).Delete(&TableMappingStaffBranch{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(branches) != 0 {
		if err := tx.Create(&branches).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Scopes(crius.ColumnEqualScope("staff_id", staff.ID)).Delete(&TableMappingStaffRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(roles) != 0 {
		if err := tx.Create(&roles).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// UpdateStaffStatus 更新员工状态
func UpdateStaffStatus(id uuid.UUID, status string) error {
	return entity.Conn.Model(&TableStaff{}).Where("id = ?", id).Update("status", status).Error
}

// GetStaffs 获取员工信息
func GetStaffs(staff TableStaff, branches []uuid.UUID, merchantID uuid.UUID, dateStart, dateEnd, offset, limit int32) ([]StaffFullInfo, int64, error) {
	staffs := make([]TableStaff, 0)
	staffFulls := make([]StaffFullInfo, 0)
	var count int64

	db := entity.Conn.Distinct().Model(&TableStaff{}).Scopes(crius.ColumnLikeScope("name", staff.Name),
		crius.ColumnEqualScopeDefault("gender", staff.Gender),
		crius.ColumnLikeScope("phone", staff.Phone),
		entryAtCondition(dateStart, dateEnd),
		crius.ColumnEqualScopeDefault("status", staff.Status.String()),
		crius.ColumnEqualScope("merchant_id", merchantID))

	if len(branches) != 0 {
		db = db.Joins("inner join merchant_basic.mapping_staff_branch on mapping_staff_branch.staff_id = staff.id").
			Where("mapping_staff_branch.branch_id in ?", branches)
	}
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return staffFulls, 0, nil
	}
	err := db.Scopes(pagingCondition(offset, limit)).Order("created_at DESC").Find(&staffs).Error
	if err != nil {
		return nil, 0, err
	}

	for _, v := range staffs {
		if v.Admin {
			staffFulls = append(staffFulls, StaffFullInfo{
				Staff: v,
			})
			continue
		}
		branches, roles, err := getBranchesRolesByStaffID(v.ID)
		if err != nil {
			return nil, 0, err
		}

		staffFulls = append(staffFulls, StaffFullInfo{
			Staff:    v,
			Branches: branches,
			Roles:    roles,
		})
	}

	return staffFulls, count, nil
}

// GetStaffsByRoleID 根据角色id获取员工
func GetStaffsByRoleID(roleID uuid.UUID) ([]TableStaff, error) {
	staffs := make([]TableStaff, 0)
	err := entity.Conn.Model(&TableStaff{}).Joins("inner join merchant_basic.mapping_staff_role on mapping_staff_role.staff_id = staff.id").
		Where("mapping_staff_role.role_id = ?", roleID).Find(&staffs).Error
	return staffs, err
}

// DeleteStaff 删除员工
func DeleteStaff(id uuid.UUID) error {
	return entity.Conn.Scopes(crius.ColumnEqualScope("id", id)).Delete(&TableStaff{}).Error
}

// ShowStaffByUsername 根据工号/编号获取用户信息
func ShowStaffByUsername(username string) (*StaffFullInfo, error) {
	if username == "" {
		return nil, gorm.ErrRecordNotFound
	}
	staff := new(TableStaff)
	err := entity.Conn.Scopes(usernameCondition(username)).Take(staff).Error
	if err != nil {
		return nil, err
	}

	//管理员不需要查询所属门店角色
	if staff.Admin {
		return &StaffFullInfo{Staff: *staff}, err
	}

	branches, roles, err := getBranchesRolesByStaffID(staff.ID)
	if err != nil {
		return nil, err
	}

	return &StaffFullInfo{
		Staff:    *staff,
		Branches: branches,
		Roles:    roles,
	}, nil
}

// ShowStaffByPhone 根据手机号查询用户信息
func ShowStaffByPhone(phone, phoneCode string) (*StaffFullInfo, error) {
	if phone == "" {
		return nil, gorm.ErrRecordNotFound
	}
	staff := new(TableStaff)
	err := entity.Conn.Scopes(phoneCondition(phone, phoneCode)).Take(staff).Error
	if err != nil {
		return nil, err
	}
	//管理员不需要查询所属门店角色
	if staff.Admin {
		return &StaffFullInfo{Staff: *staff}, err
	}

	branches, roles, err := getBranchesRolesByStaffID(staff.ID)
	if err != nil {
		return nil, err
	}

	return &StaffFullInfo{
		Staff:    *staff,
		Branches: branches,
		Roles:    roles,
	}, nil
}

// ShowStaffByID 根据ID获取单个员工信息
func ShowStaffByID(id uuid.UUID) (*StaffFullInfo, error) {
	staff := new(TableStaff)
	err := entity.Conn.Scopes(crius.ColumnEqualScope("id", id)).Take(staff).Error
	if err != nil {
		return nil, err
	}

	//管理员不需要查询所属门店角色
	if staff.Admin {
		return &StaffFullInfo{Staff: *staff}, err
	}

	branches, roles, err := getBranchesRolesByStaffID(id)
	if err != nil {
		return nil, err
	}

	staffFull := &StaffFullInfo{
		Staff:    *staff,
		Branches: branches,
		Roles:    roles,
	}
	return staffFull, err
}

// ShowBasicStaffByID 根据ID获取单个员工基本信息，不联表查询
func ShowBasicStaffByID(id uuid.UUID) (*TableStaff, error) {
	staff := new(TableStaff)
	err := entity.Conn.Scopes(crius.ColumnEqualScopeDefault("status", crius.StatusOpened.String()), quitCondition(), crius.ColumnEqualScope("id", id)).Take(staff).Error
	return staff, err
}

// UpdateStaffPassword 修改员工密码
func UpdateStaffPassword(id uuid.UUID, password string) error {
	return entity.Conn.Model(&TableStaff{}).Scopes(crius.ColumnEqualScope("id", id)).Update("password", password).Error
}

// ShowStaffExistsByCode 校验工号/编号是否存在
func ShowStaffExistsByCode(code string, merchantID uuid.UUID) error {
	if code == "" {
		return gorm.ErrRecordNotFound
	}
	return entity.Conn.Select("code", "employee_code").Unscoped().Scopes(usernameCondition(code), crius.ColumnEqualScope("merchant_id", merchantID)).Take(&TableStaff{}).Error
}

// ShowStaffExistsByPhone 根据手机号查询用户是否存在
func ShowStaffExistsByPhone(phone, phoneCode string, merchantID uuid.UUID) error {
	if phone == "" {
		return gorm.ErrRecordNotFound
	}
	return entity.Conn.Unscoped().Scopes(phoneCondition(phone, phoneCode), crius.ColumnEqualScope("merchant_id", merchantID)).Take(&TableStaff{}).Error
}

// getBranchesRolesByStaffID 根据员工id获取关联的门店和角色
func getBranchesRolesByStaffID(id uuid.UUID) ([]Branch, []RoleInfo, error) {
	branches := make([]Branch, 0)
	err := entity.Conn.Select("branch.*", "brand.name as brand_name").Model(&TableBranch{}).Joins("inner join merchant_basic.mapping_staff_branch on mapping_staff_branch.branch_id = branch.id").
		Joins("left join merchant_basic.brand on branch.brand_id = brand.id").Where("mapping_staff_branch.staff_id = ?", id).Scan(&branches).Error
	if err != nil {
		return nil, nil, err
	}

	roles := make([]RoleInfo, 0)
	// 子查询，根据员工id查询对应的角色id
	subQuery := entity.Conn.Select("role.id").Model(&TableRole{}).Joins("inner join merchant_basic.mapping_staff_role on mapping_staff_role.role_id = role.id").
		Where("mapping_staff_role.staff_id = ?", id)
	// 根据角色id查询角色信息和创建人名字
	err = entity.Conn.Select("staff.name as staff_name", "role.*").Model(&TableStaff{}).Joins("right join merchant_basic.role on role.staff_id = staff.id").Where("role.id in (?)", subQuery).Scan(&roles).Error
	if err != nil {
		return nil, nil, err
	}
	return branches, roles, nil
}

// scope条件查询-----------
func usernameCondition(username string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("code = ? or employee_code = ?", username, username)
	}
}

func entryAtCondition(dateStart, dateEnd int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if dateEnd >= dateStart && dateEnd > 0 {
			db = db.Where("entry_at >= ?", getTodayDate(int64(dateStart))).Where("entry_at < ?", getTomorrowDate(int64(dateEnd)))
		}
		return db
	}
}

func phoneCondition(phone, phoneCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if phone != "" {
			db = db.Where("phone = ?", phone)
		}
		if phoneCode != "" {
			db = db.Where("phone_code = ?", phoneCode)
		}
		return db
	}
}

func quitCondition() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("quit_at > ? or quit_at is null", time.Now())
		return db
	}
}
