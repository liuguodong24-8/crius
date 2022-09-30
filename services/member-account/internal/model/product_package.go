package model

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"time"
)

// TableProductPackage 表结构
type TableProductPackage struct {
	ID               uuid.UUID              `gorm:"column:id"`
	ProductPackageID uuid.UUID              `gorm:"column:product_package_id"`
	Code             string                 `gorm:"column:code"`
	Number           int32                  `gorm:"column:number"`
	Price            int32                  `gorm:"column:price"`
	Title            string                 `gorm:"column:title"`
	BillID           *uuid.UUID             `gorm:"column:bill_id"`
	Left             int32                  `gorm:"column:left"`
	Category         ProductPackageCategory `gorm:"column:category"`
	CreatedAt        *time.Time             `gorm:"column:created_at"`
	UpdatedAt        *time.Time             `gorm:"column:updated_at"`
	AccountBill      TableAccountBill       `gorm:"foreignkey:BillID"`
}

// TableName 表名
func (TableProductPackage) TableName() string {
	return "member_account.bill_product_package"
}

type ProductPackageItem struct {
	ID     uuid.UUID
	Number int32
}

// ProductPackageCategory 商品套餐类别
type ProductPackageCategory string

const (
	// ProductPackageCategoryProduct 商品
	ProductPackageCategoryProduct ProductPackageCategory = "product"
	// ProductPackageCategoryPackage 套餐
	ProductPackageCategoryPackage ProductPackageCategory = "package"
)

// GetProductPackagesByBillID 根据bill_id获取商品/套餐
func GetProductPackagesByBillID(billID uuid.UUID) ([]TableProductPackage, error) {
	pp := make([]TableProductPackage, 0)
	err := entity.Conn.Scopes(util.ColumnEqualScope("bill_id", billID)).Find(&pp).Error
	return pp, err
}
