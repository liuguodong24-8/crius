package centreData

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gorm.io/gorm"
	"time"
)

type TableBranchGoods struct {
	BranchGoodsID uuid.UUID  `gorm:"column:branch_goods_id"`
	BranchID      uuid.UUID  `gorm:"column:branch_id"`
	GoodsID       uuid.UUID  `gorm:"column:goods_id"`
	Price         int32      `gorm:"column:price"`
	MakeDuration  int32      `gorm:"column:make_duration"`
	UpDate        *time.Time `gorm:"column:up_date"`
	DownDate      *time.Time `gorm:"column:down_date"`
	CreateTime    *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime    *time.Time `json:"update_time" gorm:"column:update_time"`
	DeleteTime    *time.Time `json:"delete_time" gorm:"column:delete_time"`
}

func (TableBranchGoods) TableName() string {
	return "centre_data.branch_b_goods"
}

// GetBranchGoodsLowestPrice 最低价
func GetBranchGoodsLowestPrice(ids []uuid.UUID) ([]TableBranchGoods, error) {
	var goods []TableBranchGoods
	err := model.DatabaseConn().
		Select("goods_id, min(price) price").
		Where("goods_id in (?)", ids).
		Group("goods_id").
		Find(&goods).Error
	return goods, err
}

// ShowBranchGoodsLowestPrice 单个商品最低价
func ShowBranchGoodsLowestPrice(id uuid.UUID) (TableBranchGoods, error) {
	var goods TableBranchGoods
	err := model.DatabaseConn().
		Where("goods_id = ?", id).
		Order("price asc").
		First(&goods).Error
	return goods, err
}

// GetBranchGoods list
func GetBranchGoods(scopes []func(db *gorm.DB) *gorm.DB) ([]TableBranchGoods, error) {
	var data []TableBranchGoods
	err := model.DatabaseConn().Scopes(scopes...).Order("create_time desc").Find(&data).Error
	return data, err
}
