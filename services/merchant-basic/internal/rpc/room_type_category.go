package rpc

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// CreateRoomTypeCategory 创建
func (s *Server) CreateRoomTypeCategory(ctx context.Context, request *proto.CreateRoomTypeCategoryRequest) (*proto.CreateRoomTypeCategoryResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateRoomTypeCategory")
	resp := &proto.CreateRoomTypeCategoryResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	_, err := model.ShowRoomTypeCategoryByName(request.Name, merchantID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("CreateRoomTypeCategory 获取数据数据库错误:%v", err))
			resp.ErrorMessage = "创建预约分类失败"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	} else {
		crius.Logger.Error("CreateRoomTypeCategory 分类名称已存在")
		resp.ErrorMessage = "分类名称已存在"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	category := model.TableRoomTypeCategory{
		ID:         uuid.NewV4(),
		Name:       request.Name,
		Category:   int8(request.Category),
		Status:     crius.Status(request.Status),
		MerchantID: &merchantID,
	}

	if err := model.CreateRoomTypeCategory(category); err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateRoomTypeCategory 创建分类数据库错误:%v", err))
		resp.ErrorMessage = "创建分类失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	return resp, nil
}

// GetRoomTypeCategories 获取列表
func (s *Server) GetRoomTypeCategories(ctx context.Context, request *proto.GetRoomTypeCategoriesRequest) (*proto.GetRoomTypeCategoriesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetRoomTypeCategories")
	resp := &proto.GetRoomTypeCategoriesResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	category := model.TableRoomTypeCategory{
		Name:       request.Name,
		Status:     crius.Status(request.Status),
		MerchantID: &merchantID,
		Category:   int8(request.Category),
	}

	roomTypes, count, err := model.GetRoomTypeCategories(category, request.Offset, request.Limit)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetRoomTypeCategories 获取分类数据库错误:%v", err))
		resp.ErrorMessage = "获取分类失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	resp.Data = &proto.RoomTypeCategoriesData{
		Total: int32(count),
	}
	for _, v := range roomTypes {
		resp.Data.RoomTypeCategories = append(resp.Data.RoomTypeCategories, toProtoRoomTypeCategory(v))
	}

	return resp, nil
}

// UpdateRoomTypeCategory 更新
func (s *Server) UpdateRoomTypeCategory(ctx context.Context, request *proto.UpdateRoomTypeCategoryRequest) (*proto.UpdateRoomTypeCategoryResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateRoomTypeCategory")
	resp := &proto.UpdateRoomTypeCategoryResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	if category, err := model.ShowRoomTypeCategoryByName(request.Name, merchantID); err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("UpdateRoomTypeCategory 获取数据数据库错误:%v", err))
			resp.ErrorMessage = "更新预约分类失败"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	} else {
		// id不同则不是同一条数据
		if category.ID != id {
			crius.Logger.Error("UpdateRoomTypeCategory 分类名称已存在")
			resp.ErrorMessage = "分类名称已存在"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	category := model.TableRoomTypeCategory{
		ID:         id,
		Name:       request.Name,
		Category:   int8(request.Category),
		Status:     crius.Status(request.Status),
		MerchantID: &merchantID,
	}

	if err := model.UpdateRoomTypeCategory(category); err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateRoomTypeCategory 更新分类数据库错误:%v", err))
		resp.ErrorMessage = "更新分类失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	return resp, nil
}

func toProtoRoomTypeCategory(category model.TableRoomTypeCategory) *proto.RoomTypeCategory {
	return &proto.RoomTypeCategory{
		Id:       category.ID.String(),
		Name:     category.Name,
		Category: int32(category.Category),
		Status:   category.Status.String(),
	}
}
