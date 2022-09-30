package rpc

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// CreateRoomType 创建
func (s *Server) CreateRoomType(ctx context.Context, request *proto.CreateRoomTypeRequest) (*proto.CreateRoomTypeResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateRoomType")
	resp := &proto.CreateRoomTypeResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(request.BranchId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	categoryID := uuid.FromStringOrNil(request.CategoryId)
	roomTypeGroupIDs, err := fields.StringArrToUUIDArr(request.RoomTypeGroupIds)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateRoomType roomTypeGroupIDs参数错误:%v", err))
		resp.ErrorMessage = "参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}
	_, err = model.ShowRoomTypeByName(request.Name, merchantID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("CreateRoomType 获取数据数据库错误:%v", err))
			resp.ErrorMessage = "创建预约房型失败"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	} else {
		crius.Logger.Error("CreateRoomType 房型名称已存在")
		resp.ErrorMessage = "预约房型名称已存在"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	roomType := model.TableRoomType{
		ID:          uuid.NewV4(),
		Name:        request.Name,
		CategoryID:  &categoryID,
		Status:      crius.Status(request.Status),
		MerchantID:  &merchantID,
		CustomerMin: int8(request.CustomerMin),
		CustomerMax: int8(request.CustomerMax),
		Order:       request.Order,
	}

	if branchID != uuid.Nil {
		roomType.BranchID = &branchID
	}

	if len(roomTypeGroupIDs) != 0 {
		roomType.RoomTypeGroupIDs = &roomTypeGroupIDs
	}

	if err := model.CreateRoomType(roomType); err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateRoomType 创建预约房型数据库错误:%v", err))
		resp.ErrorMessage = "创建预约房型失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	return resp, nil
}

// GetRoomTypes 获取列表
func (s *Server) GetRoomTypes(ctx context.Context, request *proto.GetRoomTypesRequest) (*proto.GetRoomTypesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetRoomTypes")
	resp := &proto.GetRoomTypesResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(request.BranchId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	categoryID := uuid.FromStringOrNil(request.CategoryId)
	roomType := model.TableRoomType{
		Name:       request.Name,
		Status:     crius.Status(request.Status),
		MerchantID: &merchantID,
	}
	if branchID != uuid.Nil {
		roomType.BranchID = &branchID
	}
	if categoryID != uuid.Nil {
		roomType.CategoryID = &categoryID
	}

	roomTypes, count, err := model.GetRoomTypes(roomType, request.Offset, request.Limit)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetRoomTypes 获取预约房型数据库错误:%v", err))
		resp.ErrorMessage = "获取预约房型列表失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	resp.Data = &proto.RoomTypesData{
		Total: int32(count),
	}
	for _, v := range roomTypes {
		resp.Data.RoomTypes = append(resp.Data.RoomTypes, toProtoRoomType(v))
	}

	return resp, nil
}

// GetRoomTypesByIDs 根据id获取列表
func (s *Server) GetRoomTypesByIDs(ctx context.Context, request *proto.GetRoomTypesByIDsRequest) (*proto.GetRoomTypesByIDsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetRoomTypesByIDs")
	resp := &proto.GetRoomTypesByIDsResponse{
		ErrorCode: pkgs.Success,
	}

	ids := make([]uuid.UUID, 0)
	for _, v := range request.Ids {
		ids = append(ids, uuid.FromStringOrNil(v))
	}
	roomTypes, err := model.GetRoomTypesByIDs(ids)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetRoomTypesByIDs 获取预约房型数据库错误:%v", err))
		resp.ErrorMessage = "获取预约房型列表失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	for _, v := range roomTypes {
		resp.Data = append(resp.Data, toProtoRoomType(v))
	}

	return resp, nil
}

// ShowRoomType 获取房型
func (s *Server) ShowRoomType(ctx context.Context, request *proto.ShowRoomTypeRequest) (*proto.ShowRoomTypeResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowRoomType")
	resp := &proto.ShowRoomTypeResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)

	roomType, err := model.ShowRoomType(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "预约房型未找到"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowRoomType 获取预约房型数据库错误:%v", err))
		resp.ErrorMessage = "获取预约房型失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	resp.Data = toProtoRoomType(*roomType)
	return resp, nil
}

// UpdateRoomType 更新
func (s *Server) UpdateRoomType(ctx context.Context, request *proto.UpdateRoomTypeRequest) (*proto.UpdateRoomTypeResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateRoomType")
	resp := &proto.UpdateRoomTypeResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	categoryID := uuid.FromStringOrNil(request.CategoryId)
	roomTypeGroupIDs, err := fields.StringArrToUUIDArr(request.RoomTypeGroupIds)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateRoomType roomTypeGroupIDs参数错误:%v", err))
		resp.ErrorMessage = "参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	if roomType, err := model.ShowRoomTypeByName(request.Name, merchantID); err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("UpdateRoomType 获取数据数据库错误:%v", err))
			resp.ErrorMessage = "更新预约房型失败"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	} else {
		// id不同则不是同一条数据
		if roomType.ID != id {
			crius.Logger.Error("UpdateRoomType 房型名称已存在")
			resp.ErrorMessage = "房型名称已存在"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	roomType := model.TableRoomType{
		ID:          id,
		Name:        request.Name,
		Status:      crius.Status(request.Status),
		MerchantID:  &merchantID,
		CustomerMin: int8(request.CustomerMin),
		CustomerMax: int8(request.CustomerMax),
		Order:       request.Order,
	}
	if categoryID != uuid.Nil {
		roomType.CategoryID = &categoryID
	}
	if len(roomTypeGroupIDs) != 0 {
		roomType.RoomTypeGroupIDs = &roomTypeGroupIDs
	}
	if err := model.UpdateRoomType(roomType); err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateRoomType 更新预约房型数据库错误:%v", err))
		resp.ErrorMessage = "更新预约房型失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	return resp, nil
}

// UpdateRoomTypeStatus 更新
func (s *Server) UpdateRoomTypeStatus(ctx context.Context, request *proto.UpdateRoomTypeStatusRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateRoomTypeStatus")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	if err := model.UpdateRoomTypeStatus(id, crius.Status(request.Status)); err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateRoomTypeStatus 更新预约房型状态数据库错误:%v", err))
		resp.ErrorMessage = "更新预约房型状态失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	return resp, nil
}

func toProtoRoomType(t model.RoomType) *proto.RoomType {
	return &proto.RoomType{
		Id:               t.ID.String(),
		BranchId:         util.UUIDToString(t.BranchID),
		CategoryId:       util.UUIDToString(t.CategoryID),
		Name:             t.Name,
		Status:           t.Status.String(),
		CategoryName:     t.CategoryName,
		CustomerMax:      int32(t.CustomerMax),
		CustomerMin:      int32(t.CustomerMin),
		Order:            int32(t.Order),
		RoomTypeGroupIds: t.RoomTypeGroupIDs.ToStringArr(),
	}
}
