package centreData

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gorm.io/gorm"
	"time"
)

type TableProductGoods struct {
	GoodsID        uuid.UUID         `gorm:"column:goods_id"`
	PosCode        string            `gorm:"column:pos_code"`
	Code           string            `gorm:"column:code"`
	CnName         string            `gorm:"column:cn_name"`
	Simplify       string            `gorm:"column:simplify"`
	FinanceTypeID  *uuid.UUID        `gorm:"column:finance_type_id"`
	OperateTypeID  *uuid.UUID        `gorm:"column:operate_type_id"`
	ErpCode        string            `gorm:"column:erp_code"`
	Alias          string            `gorm:"column:alias"`
	EnName         string            `gorm:"column:en_name"`
	Crafts         int8              `gorm:"column:crafts"`
	GuidePrice     int32             `gorm:"column:guide_price"`
	BarCode        string            `gorm:"column:bar_code"`
	IsDiscountable int8              `gorm:"column:is_discountable"`
	SaleUnitID     *uuid.UUID        `gorm:"column:sale_unit_id"`
	CheckUnitID    *uuid.UUID        `gorm:"column:check_unit_id"`
	UnitRelation   int32             `gorm:"column:unit_relation"`
	Address        string            `gorm:"column:address"`
	Vendor         string            `gorm:"column:vendor"`
	Content        string            `gorm:"column:content"`
	UpDate         *time.Time        `gorm:"column:up_date"`
	DownDate       *time.Time        `gorm:"column:down_date"`
	Description    string            `gorm:"column:description"`
	BranchIDs      *fields.UUIDArr   `gorm:"column:branch_ids"`
	ImageURL       string            `gorm:"column:image_url"`
	MakeDuration   int32             `gorm:"column:make_duration"`
	WorkshopID     *uuid.UUID        `gorm:"column:workshop_id"`
	ImageUrls      *fields.StringArr `gorm:"column:image_urls"`
	Quick          *fields.Int8Arr   `gorm:"column:quick"`
	MerchantID     *uuid.UUID        `gorm:"column:merchant_id"`
	CreateTime     *time.Time        `json:"create_time" gorm:"column:create_time"`
	UpdateTime     *time.Time        `json:"update_time" gorm:"column:update_time"`
	DeleteTime     *time.Time        `json:"delete_time" gorm:"column:delete_time"`
}

type GoodsWithPrice struct {
	TableProductGoods
	Price    int32
	UnitName string
}

func (TableProductGoods) TableName() string {
	return "centre_data.product_b_goods"
}

// CountGoods count
func CountGoods(scopes []func(db *gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := model.DatabaseConn().Model(&TableProductGoods{}).Scopes(scopes...).Count(&count).Error
	return count, err
}

// GetGoods list
func GetGoods(scopes []func(db *gorm.DB) *gorm.DB) ([]TableProductGoods, error) {
	var data []TableProductGoods
	err := model.DatabaseConn().Scopes(scopes...).Order("create_time desc").Find(&data).Error
	return data, err
}

// SearchGoods 搜索
func SearchGoods(nameOrPosCode string, merchantID uuid.UUID) ([]GoodsWithPrice, error) {
	var data []GoodsWithPrice
	db := model.DatabaseConn().Select("product_b_goods.*, unit.unit_name").Where("merchant_id = ?", merchantID)
	if nameOrPosCode != "" {
		db = db.Where("cn_name like ?", fmt.Sprintf("%%%s%%", nameOrPosCode)).Or("pos_code like ?", fmt.Sprintf("%%%s%%", nameOrPosCode))
	}
	err := db.Joins("left join centre_data.product_b_unit unit on unit.unit_id = product_b_goods.sale_unit_id").Order("create_time desc").Find(&data).Error
	return data, err
}

//ShowGoods show
func ShowGoods(id, branchID uuid.UUID) (*GoodsWithPrice, error) {
	var data GoodsWithPrice
	err := model.DatabaseConn().Raw(
		`SELECT
				pg.*,
				bg.price 
			FROM
				centre_data.product_b_goods pg
				LEFT JOIN centre_data.branch_b_goods bg ON pg.goods_id = bg.goods_id and bg.branch_id = ?
			WHERE
				pg.goods_id = ?`, branchID, id).
		First(&data).Error
	return &data, err
}
