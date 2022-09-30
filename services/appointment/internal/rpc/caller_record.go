package rpc

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gorm.io/gorm"
)

// GetCallerRecords 获取来电记录列表
func (s *Server) GetCallerRecords(ctx context.Context, req *proto.GetCallerRecordsRequest) (*proto.GetCallerRecordsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetCallerRecords")
	resp := &proto.GetCallerRecordsResponse{
		ErrorCode: pkgs.Success,
	}

	var start, end time.Time
	var count int64
	var records []model.TableCallerRecord
	if req.DateStart != 0 && req.DateEnd != 0 {
		start, end = time.Unix(int64(req.DateStart), 0), time.Unix(int64(req.DateEnd), 0)
	}

	db := s.database.Conn.Scopes(model.CallAtRangeCondition(start, end), crius.ColumnEqualScopeDefault("phone", req.Phone))
	if err := db.Model(&model.TableCallerRecord{}).Count(&count).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetCallerRecords 获取来电记录列表数量数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取来电记录列表失败"
		return resp, nil
	}
	if count == 0 {
		resp.Data = &proto.GetCallerRecordsData{Total: 0}
		return resp, nil
	}

	if err := db.Scopes(model.PagingCondition(req.Offset, req.Limit)).Order("created_at desc").Find(&records).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetCallerRecords 获取来电记录列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取来电记录列表失败"
		return resp, nil
	}

	resp.Data = &proto.GetCallerRecordsData{Total: int32(count)}
	for _, v := range records {
		resp.Data.Records = append(resp.Data.Records, toProtoCallerRecord(v))
	}
	return resp, nil
}

// CreateCallerRecord 创建来电记录
func (s *Server) CreateCallerRecord(ctx context.Context, req *proto.CreateCallerRecordRequest) (*proto.CreateCallerRecordResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreateCallerRecord")
	resp := &proto.CreateCallerRecordResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	now := time.Now()

	var count int64
	if err := s.database.Conn.Model(&model.TableCallerRecord{}).Scopes(crius.ColumnEqualScope("phone", req.Phone), model.CreatedAtGreateThan(time.Now().Add(-time.Hour))).Count(&count).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateCallerRecord 获取一个小时内来电数量数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建来电记录失败"
		return resp, nil
	}

	caller := new(model.TableCaller)
	err := s.database.Conn.Scopes(crius.ColumnEqualScope("phone", req.Phone)).Take(caller).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		crius.Logger.Error(fmt.Sprintf("CreateCallerRecord 获取来电用户数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建来电记录失败"
		return resp, nil
	}

	tx := s.database.Conn.Begin()
	if err == gorm.ErrRecordNotFound {
		caller.ID = uuid.NewV4()
		caller.MerchantID = &merchantID
		caller.Phone = req.Phone
		caller.CallerName = req.Name
		caller.Gender = int8(req.Gender)
		if len(req.Phone) >= 4 {
			caller.PhoneSuffix = string(req.Phone[len(req.Phone)-4:])
		} else {
			caller.PhoneSuffix = req.Phone
		}
		caller.PhoneTail = req.Phone[len(req.Phone)-1:]
		caller.LastCallAt = &now
		if err := tx.Create(caller).Error; err != nil {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("CreateCallerRecord 创建来电用户数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "创建来电记录失败"
			return resp, nil
		}
	} else {
		caller.LastCallAt = &now
		selects := make([]string, 0)
		selects = append(selects, "last_call_at", "last_call_action", "last_operator")
		if req.Name != "" {
			caller.CallerName = req.Name
			selects = append(selects, "caller_name")
		}
		if req.Gender != 0 {
			caller.Gender = int8(req.Gender)
			selects = append(selects, "gender")
		}
		if err := tx.Scopes(crius.ColumnEqualScope("phone", req.Phone)).Select(selects).
			Updates(caller).Error; err != nil {
			tx.Rollback()
			crius.Logger.Error(fmt.Sprintf("CreateCallerRecord 更新来电用户数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "创建来电记录失败"
			return resp, nil
		}
	}

	callerRecord := model.TableCallerRecord{
		ID:         uuid.NewV4(),
		MerchantID: &merchantID,
		CallerID:   &caller.ID,
		Phone:      req.Phone,
		CallAt:     &now,
	}
	if err := tx.Create(&callerRecord).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CreateCallerRecord 创建来电记录数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建来电记录失败"
		return resp, nil
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("CreateCallerRecord 事务提交错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建来电记录失败"
		return resp, nil
	}
	resp.Data = &proto.CreateCallerRecordData{
		CallerId:       caller.ID.String(),
		CallerRecordId: callerRecord.ID.String(),
		Count:          int32(count) + 1,
	}

	return resp, nil
}

// UpdateCallerRecordAction 更新来电记录操作
func (s *Server) UpdateCallerRecordAction(ctx context.Context, req *proto.UpdateCallerRecordActionRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateCallerRecordAction")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	staffID := pkgs.GetMetadata(ctx).StaffID
	id := uuid.FromStringOrNil(req.Id)

	callerRecord := new(model.TableCallerRecord)
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("id", id)).Take(callerRecord).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "未找到来电记录"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdateCallerRecordAction 获取来电操作数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新来电记录操作失败"
		return resp, nil
	}

	callerRecord.Operator = &staffID
	if callerRecord.CallAction == nil {
		callerRecord.CallAction = &fields.StringArr{}
	}
	*callerRecord.CallAction = append(*callerRecord.CallAction, req.Action)
	tx := s.database.Conn.Begin()
	if err := tx.Updates(callerRecord).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("UpdateCallerRecordAction 更新来电操作数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新来电记录操作失败"
		return resp, nil
	}

	caller := model.TableCaller{
		LastOperator:   &staffID,
		LastCallAction: callerRecord.CallAction,
	}
	if err := tx.Scopes(crius.ColumnEqualScope("id", callerRecord.CallerID)).Updates(&caller).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("UpdateCallerRecordAction 更新来电用户操作数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新来电记录操作失败"
		return resp, nil
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("UpdateCallerRecordAction 事务提交错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新来电记录操作失败"
		return resp, nil
	}
	return resp, nil
}

func toProtoCallerRecord(r model.TableCallerRecord) *proto.CallerRecord {
	var merchantID, callerID, operator string
	var callAt int32
	if r.MerchantID != nil {
		merchantID = r.MerchantID.String()
	}
	if r.CallerID != nil {
		callerID = r.CallerID.String()
	}
	if r.Operator != nil {
		operator = r.Operator.String()
	}
	if r.CallAt != nil {
		callAt = int32(r.CallAt.Unix())
	}
	return &proto.CallerRecord{
		Id:         r.ID.String(),
		MerchantId: merchantID,
		CallerId:   callerID,
		Phone:      r.Phone,
		Operator:   operator,
		CallAction: r.CallAction.Slice(),
		CallAt:     callAt,
	}
}
