package rpc

import (
	"context"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/message/internal/model"
	"gitlab.omytech.com.cn/micro-service/message/proto"
	"gorm.io/gorm"
)

// WechatStat 微信统计
type WechatStat struct {
	BranchID     string `json:"branch_id"`
	Total        int64  `json:"total"`
	SuccessTotal int64  `json:"success_total"`
	FailTotal    int64  `json:"fail_total"`
}

// WechatStat 微信发送统计
func (s *Server) WechatStat(ctx context.Context, req *proto.WechatStatRequest) (*proto.WechatStatResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("WechatStat")

	metadata := pkgs.GetMetadata(ctx)
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))

	if len(req.MessageType) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("message_type", req.MessageType))
	}

	if len(req.BranchId) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("branch_id", req.BranchId))
	}

	if req.BeginDate > 0 {
		b := fields.UnixToDateTime(req.BeginDate)
		scopes = append(scopes, util.ColumnSymbolScope("created_at", ">", b.String()))
	}

	if req.EndDate > 0 {
		e := fields.UnixToDateTime(req.EndDate)
		scopes = append(scopes, util.ColumnSymbolScope("created_at", "<", e.String()))
	}

	column := `branch_id, count(1) as total, sum(case when status='success' then 1 else 0 end) as success_total,sum(case when status='fail' then 1 else 0 end) as fail_total`
	query := s.database.Conn.Model(&model.WechatStat{}).Scopes(scopes...).Group(`branch_id`).Select(column)

	var total int64
	if err := s.database.Conn.Table(`(?) as total`, query).Count(&total).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询微信模版消息记录总量失败")
		return &proto.WechatStatResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询总量失败:%s", err.Error()),
		}, nil
	}

	var stats []WechatStat
	var queryErr error
	if req.WithPage {
		queryErr = query.Scopes(util.PaginationScope(req.Offset, req.Limit)).Scan(&stats).Error
	} else {
		queryErr = query.Scan(&stats).Error
	}

	if queryErr != nil {
		util.Logger.WithMetadata(ctx).WithError(queryErr).Error("查询微信模版消息记录失败")
		return &proto.WechatStatResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询失败:%s", queryErr.Error()),
		}, nil
	}

	var data []*proto.WechatStat
	for _, v := range stats {
		data = append(data, toWechatStat(v))
	}

	return &proto.WechatStatResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.WechatStatData{
			Stats: data,
			Total: total,
		},
	}, nil
}

func toWechatStat(item WechatStat) *proto.WechatStat {
	return &proto.WechatStat{
		BranchId:     item.BranchID,
		Total:        item.Total,
		SuccessTotal: item.SuccessTotal,
		FailTotal:    item.FailTotal,
	}
}
