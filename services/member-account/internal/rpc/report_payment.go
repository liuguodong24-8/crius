package rpc

import (
	"context"
	"fmt"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/member-account/internal/model"
	"gitlab.omytech.com.cn/micro-service/member-account/proto"
	"gorm.io/gorm"
)

// ReportPayment 收款汇总，开卡充值收款方式汇总
func (s *Server) ReportPayment(ctx context.Context, req *proto.ReportPaymentRequest) (*proto.ReportPaymentResponse, error) {
	defer util.CatchException()

	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ReportPayment")

	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnInScope("change_type", []interface{}{model.BillTypeOpen, model.BillTypeRecharge, model.BillTypeNobody}))
	if len(req.BranchId) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("branch_id", req.BranchId))
	}

	if req.BeginTime > 0 {
		scopes = append(scopes, util.ColumnSymbolScope("created_at", ">=", time.Unix(req.BeginTime, 0)))
	}

	if req.EndTime > 0 {
		scopes = append(scopes, util.ColumnSymbolScope("created_at", "<=", time.Unix(req.EndTime, 0)))
	}

	stats, err := searchPaymentStat(scopes)
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询错误")
		return &proto.ReportPaymentResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询错误：%s", err.Error()),
		}, nil
	}

	return &proto.ReportPaymentResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         stats,
	}, nil
}

type paymentStats struct {
	Wechat struct {
		Open     int32
		Recharge int32
	}
	Card struct {
		Open     int32
		Recharge int32
	}
	Cash struct {
		Open     int32
		Recharge int32
	}
	Alipay struct {
		Open     int32
		Recharge int32
	}
}

func (p paymentStats) toProtoPaymentStat() []*proto.ReportPaymentResponse_Payment {
	var res []*proto.ReportPaymentResponse_Payment
	res = append(res, &proto.ReportPaymentResponse_Payment{
		Way:         "wechat",
		WayDesc:     "微信",
		OpenFee:     p.Wechat.Open,
		RechargeFee: p.Wechat.Recharge,
	})

	res = append(res, &proto.ReportPaymentResponse_Payment{
		Way:         "card",
		WayDesc:     "刷卡",
		OpenFee:     p.Card.Open,
		RechargeFee: p.Card.Recharge,
	})

	res = append(res, &proto.ReportPaymentResponse_Payment{
		Way:         "cash",
		WayDesc:     "现金",
		OpenFee:     p.Cash.Open,
		RechargeFee: p.Cash.Recharge,
	})

	res = append(res, &proto.ReportPaymentResponse_Payment{
		Way:         "alipay",
		WayDesc:     "支付宝",
		OpenFee:     p.Alipay.Open,
		RechargeFee: p.Alipay.Recharge,
	})

	return res
}

func searchPaymentStat(scopes []func(db *gorm.DB) *gorm.DB) ([]*proto.ReportPaymentResponse_Payment, error) {
	conn := model.DatabaseConn().Conn
	//rows, err := conn.Model(&model.TableAccountBill{}).Select([]string{
	//	`"payments"->'card' as card_fee`,
	//	`"payments"->'cash' as cash_fee`,
	//	`"payments"->'wechat' as wechat_fee`,
	//	`"payments"->'alipay' as alipay_fee`,
	//	"change_category",
	//}).Rows()

	rows, err := conn.Model(&model.TableAccountBill{}).Scopes(scopes...).Select([]string{"payments", "change_type"}).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stat paymentStats

	for rows.Next() {
		var row struct {
			Payments   *model.Payments `json:"payments"`
			ChangeType string          `json:"change_type"`
		}

		if err := conn.ScanRows(rows, &row); nil != err {
			return nil, err
		}

		if row.Payments == nil {
			continue
		}

		// wechat
		if row.Payments.Wechat > 0 {
			open, recharge := dealPaymentStatBase(model.BillType(row.ChangeType), row.Payments.Wechat)
			stat.Wechat.Open += open
			stat.Wechat.Recharge += recharge
		}

		if row.Payments.Card > 0 {
			open, recharge := dealPaymentStatBase(model.BillType(row.ChangeType), row.Payments.Card)
			stat.Card.Open += open
			stat.Card.Recharge += recharge
		}

		if row.Payments.Cash > 0 {
			open, recharge := dealPaymentStatBase(model.BillType(row.ChangeType), row.Payments.Cash)
			stat.Cash.Open += open
			stat.Cash.Recharge += recharge
		}

		if row.Payments.Alipay > 0 {
			open, recharge := dealPaymentStatBase(model.BillType(row.ChangeType), row.Payments.Alipay)
			stat.Alipay.Open += open
			stat.Alipay.Recharge += recharge
		}
	}

	return stat.toProtoPaymentStat(), nil
}

func dealPaymentStatBase(category model.BillType, fee int32) (open int32, recharge int32) {
	if category == model.BillTypeRecharge {
		recharge = fee
	}

	if category == model.BillTypeOpen || category == model.BillTypeNobody {
		open = fee
	}

	return
}
