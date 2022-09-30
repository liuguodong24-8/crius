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

// StatSmsHistory 统计消息发送
func (s *Server) StatSmsHistory(ctx context.Context, req *proto.StatSmsHistoryRequest) (*proto.StatSmsHistoryResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("StatSmsHistory")
	metadata := pkgs.GetMetadata(ctx)
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))
	if len(req.BranchId) > 0 {
		scopes = append(scopes, util.ColumnEqualScope(`branch_id`, req.BranchId))
	}
	if len(req.BranchIds) > 0 {
		scopes = append(scopes, util.ColumnInScope(`branch_id`, fields.StringArr(req.BranchIds).ToInterfaceArr()))
	}
	if len(req.MessageType) > 0 {
		scopes = append(scopes, util.ColumnEqualScope(`message_type`, req.MessageType))
	}
	if len(req.BeginDate) > 0 {
		scopes = append(scopes, util.ColumnSymbolScope(`created_at`, `>=`, req.BeginDate+` 00:00:00`))
	}
	if len(req.EndDate) > 0 {
		scopes = append(scopes, util.ColumnSymbolScope(`created_at`, `<`, req.EndDate+` 23:59:59`))
	}

	var total int64
	var result []*proto.StatSmsHistory

	statQuery := s.database.Conn.Model(&model.SmsStat{}).Scopes(scopes...).Group(`branch_id`).Select(
		`branch_id, count(branch_id) as total, sum(case when status = 'success' then 1 else 0 end) as succeed, sum(case when status = 'fail' then 1 else 0 end) as failured`,
	)

	s.database.Conn.Table(`(?) as stat`, statQuery).Count(&total)
	if total == 0 {
		return &proto.StatSmsHistoryResponse{
			ErrorCode:    pkgs.Success,
			ErrorMessage: "",
		}, nil
	}

	var p []func(db *gorm.DB) *gorm.DB
	if req.WithPage {
		p = append(p, util.PaginationScope(req.Offset, req.Limit))
	}
	orderBy := "total desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}
	if err := s.database.Conn.Table(`(?) as stat`, statQuery).Scopes(p...).Order(orderBy).Find(&result).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("统计消息发送历史列表错误")
		return &proto.StatSmsHistoryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据查询错误:%s", err.Error()),
		}, nil
	}

	return &proto.StatSmsHistoryResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.StatSmsHistoryData{
			Stats: result,
			Total: total,
		},
	}, nil
}

// ListSmsHistory 消息发送列表
func (s *Server) ListSmsHistory(ctx context.Context, req *proto.ListSmsHistoryRequest) (*proto.ListSmsHistoryResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("StatSmsHistory")
	metadata := pkgs.GetMetadata(ctx)
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))
	if len(req.Ids) > 0 {
		scopes = append(scopes, util.ColumnInScope(`id`, fields.StringArr(req.Ids).ToInterfaceArr()))
	}
	if len(req.BranchId) > 0 {
		scopes = append(scopes, util.ColumnEqualScope(`branch_id`, req.BranchId))
	}
	if len(req.MessageType) > 0 {
		scopes = append(scopes, util.ColumnEqualScope(`message_type`, req.MessageType))
	}
	if len(req.BeginDate) > 0 {
		scopes = append(scopes, util.ColumnSymbolScope(`created_at`, `>=`, req.BeginDate+` 00:00:00`))
	}
	if len(req.EndDate) > 0 {
		scopes = append(scopes, util.ColumnSymbolScope(`created_at`, `<`, req.EndDate+` 23:59:59`))
	}
	if len(req.SmsStatus) > 0 {
		scopes = append(scopes, util.ColumnEqualScope(`status`, req.SmsStatus))
	}
	if len(req.Phone) > 0 {
		scopes = append(scopes, util.ColumnLikeScope(`phone`, req.Phone))
	}

	var total int64
	var result []model.SmsStat

	s.database.Conn.Model(&model.SmsStat{}).Scopes(scopes...).Count(&total)
	if total == 0 {
		return &proto.ListSmsHistoryResponse{
			ErrorCode:    pkgs.Success,
			ErrorMessage: "",
		}, nil
	}
	if req.WithPage {
		scopes = append(scopes, util.PaginationScope(req.Offset, req.Limit))
	}
	orderBy := "created_at desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}
	if err := s.database.Conn.Model(&model.SmsStat{}).Scopes(scopes...).Order(orderBy).Find(&result).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("消息发送历史详情列表错误")
		return &proto.ListSmsHistoryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据查询错误:%s", err.Error()),
		}, nil
	}

	data := make([]*proto.DetailSmsHistory, 0)
	for _, v := range result {
		data = append(data, toProtoDetailSmsHistory(v))
	}

	return &proto.ListSmsHistoryResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ListSmsHistoryData{
			Histories: data,
			Total:     total,
		},
	}, nil
}

func toProtoDetailSmsHistory(item model.SmsStat) (result *proto.DetailSmsHistory) {
	history := proto.DetailSmsHistory{
		Id:        item.ID.String(),
		AreaCode:  item.AreaCode,
		Phone:     item.Phone,
		Sign:      item.Sign,
		Content:   item.Content,
		Status:    string(item.Status),
		CreatedAt: item.CreatedAt.Time.Unix(),
	}

	return &history
}
