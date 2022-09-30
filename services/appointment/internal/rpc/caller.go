package rpc

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gorm.io/gorm"
)

// GetCallers 获取来电用户列表
func (s *Server) GetCallers(ctx context.Context, req *proto.GetCallersRequest) (*proto.GetCallersResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetCallers")
	resp := &proto.GetCallersResponse{
		ErrorCode: pkgs.Success,
	}

	var start, end time.Time
	var count int64
	var callers []model.TableCaller
	if req.DateStart != 0 && req.DateEnd != 0 {
		start = time.Unix(int64(req.DateStart), 0)
		end = time.Unix(int64(req.DateEnd), 0)
	}

	db := s.database.Conn.Scopes(crius.ColumnEqualScopeDefault("caller_name", req.Name), crius.ColumnEqualScopeDefault("phone", req.Phone), crius.ColumnEqualScopeDefault("phone_suffix", req.PhoneSuffix),
		model.LastCallAtRangeCondition(start, end), model.BlackCondition(int8(req.IsBlack)))
	if err := db.Model(&model.TableCaller{}).Count(&count).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetCallers 获取来电列表数量数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取来电用户列表失败"
		return resp, nil
	}
	if count == 0 {
		resp.Data = &proto.GetCallersData{Total: 0}
		return resp, nil
	}

	if err := db.Scopes(model.PagingCondition(req.Offset, req.Limit)).Order("created_at desc").Find(&callers).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetCallers 获取来电列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取来电用户列表失败"
		return resp, nil
	}

	resp.Data = &proto.GetCallersData{
		Total: int32(count),
	}

	for _, v := range callers {
		resp.Data.Callers = append(resp.Data.Callers, toProtoCaller(v))
	}

	return resp, nil
}

// ShowCallerByPhone 获取来电用户
func (s *Server) ShowCallerByPhone(ctx context.Context, req *proto.ShowCallerByPhoneRequest) (*proto.ShowCallerByPhoneResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowCallerByPhone")
	resp := &proto.ShowCallerByPhoneResponse{
		ErrorCode: pkgs.Success,
	}

	caller := new(model.TableCaller)
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("phone", req.Phone)).Take(caller).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "未找到来电用户"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowCallerByPhone 获取来电用户数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取来电用户失败"
		return resp, nil
	}

	resp.Data = toProtoCaller(*caller)

	return resp, nil
}

// UpdateCaller 更新用户信息
func (s *Server) UpdateCaller(ctx context.Context, req *proto.UpdateCallerRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateCaller")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	tags := make(pkgs.ParamsArr, 0)
	for _, v := range req.Tags {
		m := make(map[string]interface{})
		m["tag"] = v.Tag
		m["color"] = v.Color
		tags = append(tags, m)
	}
	caller := model.TableCaller{
		ID:         id,
		CallerName: req.Name,
		Gender:     int8(req.Gender),
		Tags:       &tags,
	}
	if err := s.database.Conn.Updates(&caller).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateCaller 更新用户信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新用户信息失败"
		return resp, nil
	}
	return resp, nil
}

// UpdateCallerBlack 更新用户拉黑状态
func (s *Server) UpdateCallerBlack(ctx context.Context, req *proto.UpdateCallerBlackRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateCallerBlack")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	caller := model.TableCaller{
		IsBlcak: req.IsBlack,
	}
	if req.IsBlack {
		caller.BlackReason = req.Reason
	}
	if err := s.database.Conn.Select("is_black", "black_reason").Scopes(crius.ColumnEqualScope("id", id)).Updates(caller).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateCallerBlack 更新用户信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新用户信息失败"
		return resp, nil
	}

	return resp, nil
}

func toProtoCaller(c model.TableCaller) *proto.Caller {
	var merchantID, lastOperator string
	var lastCallAt int32
	var tags []*proto.Tag
	if c.MerchantID != nil {
		merchantID = c.MerchantID.String()
	}
	if c.LastOperator != nil {
		lastOperator = c.LastOperator.String()
	}
	if c.LastCallAt != nil {
		lastCallAt = int32(c.LastCallAt.Unix())
	}

	if c.Tags != nil {
		for _, v := range *c.Tags {
			tags = append(tags, &proto.Tag{Tag: v["tag"].(string), Color: v["color"].(string)})
		}
	}
	return &proto.Caller{
		Id:             c.ID.String(),
		MerchantId:     merchantID,
		PhoneCode:      c.PhoneCode,
		Phone:          c.Phone,
		PhoneTail:      c.PhoneTail,
		PhoneSuffix:    c.PhoneSuffix,
		Name:           c.CallerName,
		Gender:         int32(c.Gender),
		IsBlack:        c.IsBlcak,
		BlackReason:    c.BlackReason,
		LastCallAt:     lastCallAt,
		LastOperator:   lastOperator,
		LastCallAction: c.LastCallAction.Slice(),
		Tags:           tags,
	}
}
