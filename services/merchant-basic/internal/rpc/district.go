package rpc

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// CreateDistrict 创建
func (s *Server) CreateDistrict(ctx context.Context, request *proto.CreateDistrictRequest) (*proto.CreateDistrictResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateDistrict")
	resp := &proto.CreateDistrictResponse{
		ErrorCode: pkgs.Success,
	}

	district := model.TableDistrict{
		ID:     uuid.NewV4(),
		Name:   request.Name,
		Status: crius.Status(request.Status),
	}
	if err := model.CreateDistrict(district); err != nil {
		util.Logger.Error(fmt.Sprintf("CreateDistrict 创建区域数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建区域失败"
	}

	return resp, nil
}

// UpdateDistrict 更新
func (s *Server) UpdateDistrict(ctx context.Context, request *proto.UpdateDistrictRequest) (*proto.UpdateDistrictResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateDistrict")
	resp := &proto.UpdateDistrictResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	district := model.TableDistrict{
		ID:     id,
		Name:   request.Name,
		Status: crius.Status(request.Status),
	}
	if err := model.UpdateDistrict(district); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateDistrict 更新区域数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新区域失败"
	}

	return resp, nil
}

// GetDistricts 获取列表
func (s *Server) GetDistricts(ctx context.Context, request *proto.GetDistrictsRequest) (*proto.GetDistrictsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetDistricts")
	resp := &proto.GetDistrictsResponse{
		ErrorCode: pkgs.Success,
	}

	district := model.TableDistrict{
		Name:   request.Name,
		Status: crius.Status(request.Status),
	}
	districts, count, err := model.GetDistricts(district, request.Offset, request.Limit)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetDistricts 获取区域列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取区域列表失败"
	}
	resp.Data = &proto.GetDistrictsData{
		Total: int32(count),
	}
	for _, v := range districts {
		resp.Data.Districts = append(resp.Data.Districts, toProtoDistrict(v))
	}

	return resp, nil
}

func toProtoDistrict(d model.TableDistrict) *proto.District {
	return &proto.District{
		Id:     d.ID.String(),
		Name:   d.Name,
		Code:   fmt.Sprintf("%03d", d.Code),
		Status: d.Status.String(),
	}
}
