package rpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/member-private/internal/micro"
	"gitlab.omytech.com.cn/micro-service/member-private/internal/model"
	"gitlab.omytech.com.cn/micro-service/member-private/proto"
	"gorm.io/gorm"
)

// CreateStaffShift 新增班次
func (s *Server) CreateStaffShift(ctx context.Context, req *proto.CreateStaffShiftRequest) (*proto.CreateStaffShiftResponse, error) {
	defer util.CatchException()

	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("CreateStaffShift")

	begin, err := s.getStaffShiftBeginTime(ctx, getShiftTimeRequest{
		BranchId:  req.BranchId,
		ShiftTime: req.ShiftTime,
	})
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("获取上次交班信息错误")
		return &proto.CreateStaffShiftResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("获取上次交班信息错误:%s", err.Error()),
		}, nil
	}
	metadata := pkgs.GetMetadata(ctx)

	shift := model.StaffShift{
		ID:         uuid.NewV4(),
		MerchantID: pkgs.GetMetadata(ctx).MerchantID,
		BranchID:   uuid.FromStringOrNil(req.BranchId),
		StaffID:    metadata.StaffID,
		BeginTime:  *begin,
		EndTime:    time.Now(),
	}

	if err := s.database.Conn.Create(&shift).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("保存交班信息错误")
		return &proto.CreateStaffShiftResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("保存交班信息错误:%s", err.Error()),
		}, nil
	}

	after := pkgs.MakeParams(shift)
	go s.SaveSnapshot(ctx, model.Snapshot{
		StaffID:           metadata.StaffID,
		SnapShotTableName: shift.TableName(),
		TableID:           shift.ID,
		Method:            model.CreateMethod,
		After:             &after,
	})

	return &proto.CreateStaffShiftResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.CreateStaffShiftResponse_Shift{
			BeginTime: shift.BeginTime.Unix(),
			EndTime:   shift.EndTime.Unix(),
		},
	}, nil
}

// GetStaffShiftTime 获取员工当前班次起始时间
func (s *Server) GetStaffShiftTime(ctx context.Context, req *proto.GetStaffShiftTimeRequest) (*proto.GetStaffShiftTimeResponse, error) {
	defer util.CatchException()

	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("CreateStaffShift")

	begin, err := s.getStaffShiftBeginTime(ctx, getShiftTimeRequest{
		BranchId: req.BranchId,
	})

	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("获取上次交班信息错误")
		return &proto.GetStaffShiftTimeResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("获取上次交班信息错误:%s", err.Error()),
		}, nil
	}

	return &proto.GetStaffShiftTimeResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.GetStaffShiftTimeResponse_Shift{
			BeginTime: begin.Unix(),
			EndTime:   time.Now().Unix(),
		},
	}, nil
}

type getShiftTimeRequest struct {
	BranchId  string
	ShiftTime int64
}

func (s *Server) getStaffShiftBeginTime(ctx context.Context, req getShiftTimeRequest) (*time.Time, error) {
	//todo metadata再次传递获取不到值
	metadata := pkgs.GetMetadata(ctx)
	// 第一次 不存在 使用最近一次门店营业开始时间
	// 存在 使用上次交班时间
	var shift model.StaffShift
	if err := s.database.Conn.Model(&model.StaffShift{}).Scopes(util.ColumnEqualScope("branch_id", req.BranchId), util.ColumnEqualScope("staff_id", metadata.StaffID)).Order("created_at desc").First(&shift).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return s.getBranchLatelyBusinessBegin(pkgs.MetadataContent(metadata), req)
		}

		return nil, fmt.Errorf("查询上次交班时间错误:%s", err.Error())
	}

	return &shift.CreatedAt, nil
}

func (s *Server) getBranchLatelyBusinessBegin(ctx context.Context, req getShiftTimeRequest) (*time.Time, error) {
	client, err := micro.GetMerchantBasicServer(ctx, s.crius)
	if err != nil {
		return nil, err
	}

	business, businessErr := client.GetBranchLatelyBusiness(ctx, req.BranchId, req.ShiftTime)
	if businessErr != nil {
		return nil, businessErr
	}

	if business.BusinessBegin == nil {
		return nil, errors.New("门店营业日信息错误")
	}

	return business.BusinessBegin, nil
}

// ListStaffShift 班次列表
func (s *Server) ListStaffShift(ctx context.Context, req *proto.ListStaffShiftRequest) (*proto.ListStaffShiftResponse, error) {
	defer util.CatchException()

	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("ListStaffShift")

	metadata := pkgs.GetMetadata(ctx)
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))
	scopes = append(scopes, util.ColumnEqualScope("staff_id", metadata.StaffID))
	if len(req.BranchId) == 0 {
		scopes = append(scopes, util.ColumnEqualScope("branch_id", req.BranchId))
	}

	if req.BeginAt > 0 {
		scopes = append(scopes, util.ColumnSymbolScope("created_at", ">=", time.Unix(req.BeginAt, 0)))
	} else {
		t := time.Now().AddDate(0, 0, -7)
		scopes = append(scopes, util.ColumnSymbolScope("created_at", ">=", t))
	}

	if req.EndAt > 0 {
		scopes = append(scopes, util.ColumnSymbolScope("created_at", "<=", time.Unix(req.EndAt, 0)))
	}
	orderBy := "created_at desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}

	var shifts []model.StaffShift
	if err := s.database.Conn.Model(&model.StaffShift{}).Scopes(scopes...).Order(orderBy).Find(&shifts).Error; nil != err {
		return &proto.ListStaffShiftResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询交班列表错误:%s", err.Error()),
		}, nil
	}

	data := make([]*proto.ListStaffShiftResponse_StaffShift, 0)
	for _, v := range shifts {
		data = append(data, &proto.ListStaffShiftResponse_StaffShift{
			Id:        v.ID.String(),
			BeginTime: v.BeginTime.Unix(),
			EndTime:   v.EndTime.Unix(),
			CreatedAt: v.CreatedAt.Unix(),
		})
	}

	return &proto.ListStaffShiftResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         data,
	}, nil
}
