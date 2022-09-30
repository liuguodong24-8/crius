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

// ReportPromotion 开卡充值汇总
func (s *Server) ReportPromotion(ctx context.Context, req *proto.ReportBillDetailRequest) (*proto.ReportPromotionResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("ReportOpenRecharge")

	stats, total, err := model.SearchPromotionStats(model.PromotionStatRequest{
		Scopes:   reportBillDetailScopes(req),
		WithPage: req.WithPage,
		Offset:   req.Offset,
		Limit:    req.Limit,
	})

	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询开卡充值，优惠方案汇总错误")
		return &proto.ReportPromotionResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询开卡充值，优惠方案汇总错误:%s", err.Error()),
		}, nil
	}

	var items []*proto.ReportPromotionResponse_Report

	for _, item := range stats {
		items = append(items, &proto.ReportPromotionResponse_Report{
			PromotionOptionId:   item.PromotionOptionID.String(),
			PromotionOptionName: item.PromotionOptionName,
			Total:               item.Total,
			OpenTotal:           item.OpenTotal,
			RechargeTotal:       item.RechargeTotal,
			TotalValue:          item.TotalRechargeValue,
		})
	}

	return &proto.ReportPromotionResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ReportPromotionResponse_Data{
			Data:  items,
			Total: total,
		},
	}, nil
}

// ReportBillDetail 开卡详情
func (s *Server) ReportBillDetail(ctx context.Context, req *proto.ReportBillDetailRequest) (*proto.ReportBillDetailResponse, error) {
	defer util.CatchException()

	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("ReportOpenDetail")

	items, total, err := model.GetReportBills(model.GetReportBillsRequest{
		Scopes:   reportBillDetailScopes(req),
		WithPage: req.WithPage,
		Offset:   req.Offset,
		Limit:    req.Limit,
	})

	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询列表错误")
		return &proto.ReportBillDetailResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询报表列表错误:%s", err.Error()),
		}, nil
	}

	var response []*proto.ReportBillDetailResponse_Report

	for _, item := range items {
		payments := &proto.Payments{}
		if item.Payments != nil {
			payments.Wechat = item.Payments.Wechat
			payments.Cash = item.Payments.Cash
			payments.Card = item.Payments.Card
			payments.Alipay = item.Payments.Alipay
		}

		var options []*proto.ReportBillDetailResponse_Report_PromotionOption
		if item.PromotionOptions != nil {
			for _, o := range *item.PromotionOptions {
				options = append(options, &proto.ReportBillDetailResponse_Report_PromotionOption{
					PromotionOptionId:   o.ID.String(),
					PromotionOptionName: o.Name,
					Count:               int32(o.Count),
					RechargeValue:       o.RechargeValue,
				})
			}
		}

		response = append(response, &proto.ReportBillDetailResponse_Report{
			Id:               item.ID.String(),
			CardId:           util.UUIDToString(item.CardID),
			CardCode:         item.CardCode,
			Payments:         payments,
			PromotionOptions: options,
			RechargeValue:    item.ChangeValue,
		})
	}

	return &proto.ReportBillDetailResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ReportBillDetailResponse_Data{
			Data:  response,
			Total: total,
		},
	}, nil
}

func makeInterfaceSlice(items []string) []interface{} {
	var res []interface{}
	for _, item := range items {
		res = append(res, item)
	}

	return res
}

func reportBillDetailScopes(req *proto.ReportBillDetailRequest) []func(db *gorm.DB) *gorm.DB {
	var scopes []func(db *gorm.DB) *gorm.DB
	if len(req.ReportType) > 0 {
		scopes = append(scopes, util.ColumnInScope("change_type", makeInterfaceSlice(req.ReportType)))
	} else {
		scopes = append(scopes, util.ColumnInScope("change_type", []interface{}{model.BillTypeOpen, model.BillTypeRecharge, model.BillTypeNobody}))
	}
	if len(req.BranchId) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("branch_id", req.BranchId))
	}

	if req.BeginTime > 0 {
		scopes = append(scopes, util.ColumnSymbolScope("created_at", ">=", time.Unix(req.BeginTime, 0)))
	}

	if req.EndTime > 0 {
		scopes = append(scopes, util.ColumnSymbolScope("created_at", "<=", time.Unix(req.EndTime, 0)))
	}

	return scopes
}
