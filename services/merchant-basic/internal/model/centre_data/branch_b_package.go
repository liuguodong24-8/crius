package centreData

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gorm.io/gorm"
	"time"
)

type TableBranchPackage struct {
	BranchPackageID uuid.UUID  `gorm:"column:branch_package_id"`
	BranchID        uuid.UUID  `gorm:"column:branch_id"`
	PackageID       uuid.UUID  `gorm:"column:package_id"`
	Price           int32      `gorm:"column:price"`
	BeginDate       *time.Time `gorm:"column:begin_date"`
	EndDate         *time.Time `gorm:"column:end_date"`
	CreateTime      *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime      *time.Time `json:"update_time" gorm:"column:update_time"`
	DeleteTime      *time.Time `json:"delete_time" gorm:"column:delete_time"`
}

func (TableBranchPackage) TableName() string {
	return "centre_data.branch_b_package"
}

// GetBranchPackagesLowestPrice 最低价
func GetBranchPackagesLowestPrice(ids []uuid.UUID) ([]TableBranchPackage, error) {
	var packages []TableBranchPackage
	err := model.DatabaseConn().
		Select("package_id, min(price) price").
		Where("package_id in (?)", ids).
		Group("package_id").
		Find(&packages).Error
	return packages, err
}

// ShowBranchPackageLowestPrice 单个商品最低价
func ShowBranchPackageLowestPrice(id uuid.UUID) (TableBranchPackage, error) {
	var goods TableBranchPackage
	err := model.DatabaseConn().
		Where("package_id = ?", id).
		Order("price asc").
		First(&goods).Error
	return goods, err
}

// GetBranchPackages list
func GetBranchPackages(scopes []func(db *gorm.DB) *gorm.DB) ([]TableBranchPackage, error) {
	var data []TableBranchPackage
	err := model.DatabaseConn().Scopes(scopes...).Order("create_time desc").Find(&data).Error
	return data, err
}
