package centreData

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gorm.io/gorm"
	"time"
)

type TablePromotionPackage struct {
	PackageID    uuid.UUID       `gorm:"column:package_id"`
	PackageName  string          `gorm:"column:package_name"`
	Code         string          `gorm:"column:code"`
	ImageURL     string          `gorm:"column:image_url"`
	PosCode      string          `gorm:"column:pos_code"`
	Simplify     string          `gorm:"column:simplify"`
	ActiveTypeID *uuid.UUID      `gorm:"column:active_type_id"`
	BeginDate    *time.Time      `gorm:"column:begin_date"`
	EndDate      *time.Time      `gorm:"column:end_date"`
	GoodsSet     *pkgs.ParamsArr `gorm:"column:goods_set"`
	BranchIDs    *fields.UUIDArr `gorm:"column:branch_ids"`
	MerchantID   *uuid.UUID      `gorm:"column:merchant_id"`
	CreateTime   *time.Time      `json:"create_time" gorm:"column:create_time"`
	UpdateTime   *time.Time      `json:"update_time" gorm:"column:update_time"`
	DeleteTime   *time.Time      `json:"delete_time" gorm:"column:delete_time"`
}

type PackageWithPrice struct {
	TablePromotionPackage
	Price int32
}

func (TablePromotionPackage) TableName() string {
	return "centre_data.promotion_b_package"
}

// CountPackages count
func CountPackages(scopes []func(db *gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := model.DatabaseConn().Model(&TablePromotionPackage{}).Scopes(scopes...).Count(&count).Error
	return count, err
}

// GetPackages list
func GetPackages(scopes []func(db *gorm.DB) *gorm.DB) ([]TablePromotionPackage, error) {
	var data []TablePromotionPackage
	err := model.DatabaseConn().Scopes(scopes...).Order("create_time desc").Find(&data).Error
	return data, err
}

// SearchPackages 搜索
func SearchPackages(nameOrPosCode string, merchantID uuid.UUID) ([]PackageWithPrice, error) {
	var data []PackageWithPrice
	db := model.DatabaseConn().Where("merchant_id = ?", merchantID)
	if nameOrPosCode != "" {
		db = db.Where("package_name like ?", fmt.Sprintf("%%%s%%", nameOrPosCode)).Or("pos_code like ?", fmt.Sprintf("%%%s%%", nameOrPosCode))
	}
	err := db.Order("create_time desc").Find(&data).Error
	return data, err
}

//ShowPackage show
func ShowPackage(id, branchID uuid.UUID) (*PackageWithPrice, error) {
	var data PackageWithPrice
	err := model.DatabaseConn().Raw(
		`SELECT
				pp.*,
				bp.price 
			FROM
				centre_data.promotion_b_package pp
			LEFT JOIN centre_data.branch_b_package bp ON bp.package_id = pp.package_id and bp.branch_id = ?
			WHERE
				pp.package_id = ?;`, branchID, id).
		First(&data).Error
	return &data, err
}
