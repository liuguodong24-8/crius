package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// ReportBillPromotion 详情
type ReportBillPromotion struct {
	ID                  uuid.UUID `json:"id"`
	BranchID            uuid.UUID `json:"branch_id"`
	BillID              uuid.UUID `json:"bill_id"`
	CardID              uuid.UUID `json:"card_id"`
	CardCode            string    `json:"card_code"`
	PromotionOptionID   uuid.UUID `json:"promotion_option_id"`
	PromotionOptionName string    `json:"promotion_option_name"`
	RechargeValue       int32     `json:"recharge_value"`
	Total               int       `json:"total"`
	ChangeType          string    `json:"change_type"`
	ChangeValue         int32     `json:"change_value"`
	CreatedAt           time.Time `json:"created_at"`
}

// TableName ...
func (r ReportBillPromotion) TableName() string {
	return `member_account.report_promotion_stat`
}

// SaveBillPromotion 保存流水方案详情
func SaveBillPromotion(bill TableAccountBill) {
	// 没有选择优惠信息，跳过
	if nil == bill.PromotionOptions {
		return
	}

	for _, option := range *bill.PromotionOptions {
		detail := &ReportBillPromotion{
			ID:                  uuid.NewV4(),
			BranchID:            *bill.BranchID,
			BillID:              bill.ID,
			CardID:              *bill.CardID,
			CardCode:            bill.CardCode,
			PromotionOptionID:   option.ID,
			PromotionOptionName: option.Name,
			RechargeValue:       option.RechargeValue,
			Total:               option.Count,
			ChangeType:          string(bill.ChangeType),
			ChangeValue:         bill.ChangeValue,
		}

		if err := entity.Conn.Create(&detail).Error; nil != err {
			util.Logger.WithError(err).WithFields("bill detail", logger.MakeFields(detail)).Error("保存账户流水详情错误")
			return
		}
	}
}

// PromotionStat 优惠方案汇总统计
type PromotionStat struct {
	PromotionOptionID   uuid.UUID `json:"promotion_option_id"`
	PromotionOptionName string    `json:"promotion_option_name"`
	Total               int32     `json:"total"`
	OpenTotal           int32     `json:"open_total"`
	RechargeTotal       int32     `json:"recharge_total"`
	TotalRechargeValue  int32     `json:"total_recharge_value"`
}

// PromotionStatRequest 优惠方案汇总查询条件
type PromotionStatRequest struct {
	Scopes   []func(db *gorm.DB) *gorm.DB
	WithPage bool
	Offset   int32
	Limit    int32
}

// SearchPromotionStats 查询优惠方案汇总信息
func SearchPromotionStats(req PromotionStatRequest) ([]PromotionStat, int64, error) {
	var total int64
	var stats []PromotionStat

	columns := fmt.Sprintf(`promotion_option_id,
		promotion_option_name, count(1) as total, 
		sum(total*recharge_value) as total_recharge_value,
		sum(case when change_type in ('%s', '%s') then 1 else 0 end) as open_total,
		sum(case when change_type='%s' then 1 else 0 end) as recharge_total`,
		BillTypeOpen, BillTypeNobody, BillTypeRecharge)

	query := entity.Conn.Model(&ReportBillPromotion{}).Scopes(req.Scopes...).Group(`promotion_option_id,promotion_option_name`).Select(columns)

	if err := entity.Conn.Table(`(?) as stat`, query).Count(&total).Error; nil != err {
		return nil, 0, err
	}

	if req.WithPage {
		req.Scopes = append(req.Scopes, util.PaginationScope(req.Offset, req.Limit))
	}

	err := entity.Conn.Model(&ReportBillPromotion{}).Scopes(req.Scopes...).Group(`promotion_option_id,promotion_option_name`).Select(columns).Order("total desc").Find(&stats).Error

	return stats, total, err
}
