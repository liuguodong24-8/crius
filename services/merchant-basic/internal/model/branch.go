package model

import (
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/util"

	crius "gitlab.omytech.com.cn/micro-service/Crius/util"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gorm.io/gorm"
)

// TableBranch branch表
type TableBranch struct {
	ID                uuid.UUID         `gorm:"column:id"`
	Name              string            `gorm:"column:name"`
	ProvinceID        string            `gorm:"column:province_id"`
	CityID            string            `gorm:"column:city_id"`
	DistrictID        string            `gorm:"column:district_id"`
	Address           string            `gorm:"column:address"`
	ContactPhone      string            `gorm:"column:contact_phone"`
	Code              string            `gorm:"column:code"`
	Longitude         float32           `gorm:"column:longitude"`
	Latitude          float32           `gorm:"column:latitude"`
	Status            crius.Status      `gorm:"column:status"`
	MerchantID        *uuid.UUID        `gorm:"column:merchant_id"`
	OpenedAt          *time.Time        `gorm:"column:opened_at"`
	Photo             *fields.StringArr `gorm:"column:photo"`
	Parking           string            `gorm:"column:parking"`
	AreaID            *uuid.UUID        `gorm:"column:area_id"`
	Weight            int32             `gorm:"column:weight"`
	Domain            string            `gorm:"column:domain"`
	BizType           int8              `gorm:"column:biz_type"`
	Alias             string            `gorm:"column:alias"`
	Simplify          string            `gorm:"column:simplify"`
	BusinessStatus    string            `gorm:"column:business_status"`
	Location          string            `gorm:"column:location"`
	AuthorizationSalt string            `json:"authorization_salt"`
	SignatureSalt     string            `json:"signature_salt"`
	SubMchID          string            `json:"sub_mch_id"`
	Extra             *pkgs.Params      `gorm:"column:extra"`
	LoadExtra         string            `gorm:"-"` // 导入信息 不处理
	CreatedAt         *time.Time        `gorm:"column:created_at"`
	UpdatedAt         *time.Time        `gorm:"column:updated_at"`
	BrandID           *uuid.UUID        `gorm:"column:brand_id"`
}

// Branch 门店查询结构
type Branch struct {
	TableBranch
	BrandName string
}

// TableName 指定表名
func (TableBranch) TableName() string {
	return "merchant_basic.branch"
}

// BranchWithMch 门店商户号
type BranchWithMch struct {
	ID       uuid.UUID `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	SubMchID string    `gorm:"column:sub_mch_id"`
}

// BranchWechatPayment 带门店的微信支付设置
type BranchWechatPayment struct {
	BranchID             uuid.UUID
	MerchantID           uuid.UUID
	AppID                string
	MchID                string
	SubMchID             string
	HeadquartersSubMchID string
	PrivateKey           string
	CertFilename         string
	CertContent          []byte
}

// CreateBranch 新增门店
func CreateBranch(branch TableBranch, businesses []TableBranchBusiness) error {
	tx := entity.Conn.Begin()
	if err := tx.Create(&branch).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&TableBranchAppointment{
		BranchID:               &branch.ID,
		OpenAppointment:        1,
		AppointmentGranularity: 15,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&businesses).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// UpdateBranch 更新门店信息
func UpdateBranch(branch TableBranch, businesses []TableBranchBusiness) error {
	tx := entity.Conn.Begin()
	if err := tx.Select("id", "name", "province_id", "city_id", "district_id", "address", "contact_phone", "latitude", "location",
		"longitude", "parking", "weight", "domain", "biz_type", "business_status", "alias", "simplify", "area_id", "photo", "opened_at", "brand_id").
		Scopes(crius.ColumnEqualScope("id", branch.ID)).Updates(&branch).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Scopes(crius.ColumnEqualScopeDefault("branch_id", branch.ID), crius.ColumnEqualScope("merchant_id", branch.MerchantID)).Delete(&TableBranchBusiness{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(businesses) != 0 {
		if err := tx.Create(&businesses).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// UpdateBranchStatus 更新门店状态
func UpdateBranchStatus(id uuid.UUID, status string) error {
	return entity.Conn.Model(&TableBranch{}).Where("id = ?", id).Update("status", status).Error
}

// GetBranches 获取门店信息
func GetBranches(branch TableBranch, businessStatus []string, brandID, staffID, merchantID uuid.UUID, offset, limit int32) ([]Branch, int64, error) {
	branches := make([]Branch, 0)
	var count int64

	db := entity.Conn.Model(&TableBranch{})
	if staffID != uuid.Nil {
		db = db.Joins("inner join merchant_basic.mapping_staff_branch on mapping_staff_branch.branch_id = branch.id").
			Where("mapping_staff_branch.staff_id = ?", staffID)
	}
	db = db.Joins("left join merchant_basic.brand on branch.brand_id = brand.id").Scopes(
		crius.ColumnLikeScope("branch.name", branch.Name), crius.ColumnEqualScopeDefault("branch.status", branch.Status.String()), crius.ColumnInScopeDefault("business_status", businessStatus),
		districtCondition(branch.ProvinceID, branch.CityID, branch.DistrictID), crius.ColumnEqualScope("branch.merchant_id", merchantID), crius.ColumnEqualScopeDefault("brand.id", brandID))
	if branch.AreaID != nil {
		db = db.Scopes(crius.ColumnEqualScopeDefault("area_id", *branch.AreaID))
	}

	err := db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return branches, 0, nil
	}

	err = db.Select("branch.*", "brand.name as brand_name").Scopes(pagingCondition(offset, limit)).Order("weight DESC").Find(&branches).Error
	if err != nil {
		return nil, 0, err
	}

	return branches, count, err
}

// GetBranchesByIDs 根据id获取门店信息
func GetBranchesByIDs(ids []uuid.UUID, status string, businessStatus []string) ([]Branch, error) {
	var branches []Branch
	err := entity.Conn.Select("branch.*", "brand.name as brand_name").Joins("left join merchant_basic.brand on branch.brand_id = brand.id").
		Scopes(crius.ColumnEqualScopeDefault("branch.status", status), crius.ColumnInScopeDefault("business_status", businessStatus), crius.ColumnInScopeDefault("branch.id", ids)).Order("weight DESC").Find(&branches).Error
	return branches, err
}

// ShowBranchByID 根据ID获取单个门店信息
func ShowBranchByID(id uuid.UUID) (*Branch, error) {
	branch := new(Branch)
	err := entity.Conn.Select("branch.*", "brand.name as brand_name").Joins("left join merchant_basic.brand on branch.brand_id = brand.id").Scopes(crius.ColumnEqualScope("branch.id", id)).Take(branch).Error
	return branch, err
}

// DeleteBranch 删除门店
func DeleteBranch(id uuid.UUID) error {
	return entity.Conn.Scopes(crius.ColumnEqualScope("id", id)).Delete(&TableBranch{}).Error
}

// ShowBranchExistsByCode 根据编号查询门店是否存在
func ShowBranchExistsByCode(code string, merchantID uuid.UUID) error {
	if code == "" {
		return gorm.ErrRecordNotFound
	}
	return entity.Conn.Select("code").Unscoped().Scopes(crius.ColumnEqualScopeDefault("code", code), crius.ColumnEqualScope("merchant_id", merchantID)).Take(&TableBranch{}).Error
}

// ShowBranchByName 根据名字查询门店
func ShowBranchByName(name string, merchantID uuid.UUID) (*TableBranch, error) {
	branch := new(TableBranch)
	if name == "" {
		return nil, gorm.ErrRecordNotFound
	}
	err := entity.Conn.Unscoped().Scopes(crius.ColumnEqualScopeDefault("name", name), crius.ColumnEqualScope("merchant_id", merchantID)).Take(branch).Error
	return branch, err
}

// GetBranchesWithSubMchIDCount count
func GetBranchesWithSubMchIDCount(provinceID, cityID, districtID, branchName string, merchantID uuid.UUID) (int64, error) {
	var count int64
	err := entity.Conn.Model(TableBranch{}).Scopes(
		areaCondition(provinceID, cityID, districtID),
		crius.ColumnLikeScope("name", branchName),
		crius.ColumnEqualScopeDefault("status", util.StatusOpened.String()),
		crius.ColumnEqualScope("merchant_id", merchantID),
	).Count(&count).Error
	return count, err
}

// GetBranchesWithSubMchID 根据city查询门店配置sub_mch_id的情况
func GetBranchesWithSubMchID(provinceID, cityID, districtID, branchName string, merchantID uuid.UUID, offset, limit int32) (*[]BranchWithMch, error) {
	var branchWithMch []BranchWithMch
	err := entity.Conn.Model(TableBranch{}).Scopes(
		areaCondition(provinceID, cityID, districtID),
		crius.ColumnLikeScope("name", branchName),
		crius.ColumnEqualScopeDefault("status", util.StatusOpened.String()),
		crius.ColumnEqualScope("merchant_id", merchantID),
		pagingCondition(offset, limit),
	).Find(&branchWithMch).Error

	return &branchWithMch, err
}

// SetBranchSubMchID 设置门店商户号
func SetBranchSubMchID(branchID uuid.UUID, subMchID string) error {
	return entity.Conn.Model(TableBranch{}).Scopes(crius.ColumnEqualScope("id", branchID)).Update("sub_mch_id", subMchID).Error
}

// GetBranchWechatPaymentSetting 获取门店微信支付信息
func GetBranchWechatPaymentSetting(branchID uuid.UUID) (BranchWechatPayment, error) {
	var branchWechatPayment BranchWechatPayment
	err := entity.Conn.Model(&TableBranch{}).
		Select("branch.id branch_id, branch.sub_mch_id, wp.*").
		Joins("left join merchant_basic.wechat_payment wp on branch.merchant_id = wp.merchant_id ").
		Where("branch.id = ?", branchID).
		First(&branchWechatPayment).Error
	return branchWechatPayment, err
}

// scope条件查询-----------

func areaCondition(provinceID, cityID, districtID string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if districtID != "" {
			return db.Where("district_id = ?", districtID)
		}
		if cityID != "" {
			return db.Where("city_id = ?", cityID)
		}
		if provinceID != "" {
			return db.Where("province_id = ?", provinceID)
		}
		return db
	}
}

func districtCondition(provinceID, cityID, districtID string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if provinceID != "" {
			db = db.Where("province_id = ?", provinceID)
			if cityID != "" {
				db = db.Where("city_id = ?", cityID)
				if districtID != "" {
					db = db.Where("district_id = ?", districtID)
				}
			}
		}
		return db
	}
}
