package rpc

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// CreateBranchTag 创建标签
func (s *Server) CreateBranchTag(ctx context.Context, request *proto.CreateBranchTagRequest) (*proto.CreateBranchTagResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateBranchTag")
	resp := &proto.CreateBranchTagResponse{
		ErrorCode: pkgs.Success,
	}

	operatorID := pkgs.GetMetadata(ctx).StaffID
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if request.Name == "" || operatorID == uuid.Nil || merchantID == uuid.Nil {
		crius.Logger.Error("CreateBranchTag 请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	_, err := model.ShowBranchTagByName(request.Name, merchantID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("CreateBranchTag 查询标签数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "创建标签失败"
			return resp, nil
		}
	} else {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "标签名已存在"
		return resp, nil
	}

	branchIds := make(fields.UUIDArr, 0)
	for _, v := range request.BranchIds {
		id := uuid.FromStringOrNil(v)
		if id != uuid.Nil {
			branchIds = append(branchIds, id)
		} else {
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "请求参数错误"
			return resp, nil
		}
	}

	tag := model.TableBranchTag{
		ID:            uuid.NewV4(),
		Name:          request.Name,
		Branches:      &branchIds,
		CreateStaffID: &operatorID,
		Status:        model.StatusOpened,
		MerchantID:    &merchantID,
	}

	err = model.CreateBranchTag(tag)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateBranchTag 创建标签数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建标签失败"
		return resp, nil
	}
	return resp, nil
}

// UpdateBranchTag 更新标签
func (s *Server) UpdateBranchTag(ctx context.Context, request *proto.UpdateBranchTagRequest) (*proto.UpdateBranchTagResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateBranchTag")
	resp := &proto.UpdateBranchTagResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	id := uuid.FromStringOrNil(request.Id)
	if request.Name == "" || merchantID == uuid.Nil || id == uuid.Nil {
		crius.Logger.Error("UpdateBranchTag 请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	tabTag, err := model.ShowBranchTagByName(request.Name, merchantID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("UpdateBranchTag 查询标签数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "更新标签失败"
			return resp, nil
		}
	} else if tabTag.ID != id {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "标签名已存在"
		return resp, nil
	}

	branchIds := make(fields.UUIDArr, 0)
	for _, v := range request.BranchIds {
		id := uuid.FromStringOrNil(v)
		if id != uuid.Nil {
			branchIds = append(branchIds, id)
		} else {
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "请求参数错误"
			return resp, nil
		}
	}

	tag := model.TableBranchTag{
		ID:       id,
		Name:     request.Name,
		Branches: &branchIds,
	}

	err = model.UpdateBranchTag(tag)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBranchTag 更新标签数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新标签失败"
		return resp, nil
	}
	return resp, nil
}

// UpdateBranchTagStatus 更新标签状态
func (s *Server) UpdateBranchTagStatus(ctx context.Context, request *proto.UpdateBranchTagStatusRequest) (*proto.UpdateBranchTagStatusResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateBranchTagStatus")
	resp := &proto.UpdateBranchTagStatusResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	id := uuid.FromStringOrNil(request.Id)
	if (request.Status != model.StatusOpened && request.Status != model.StatusClosed) || merchantID == uuid.Nil || id == uuid.Nil {
		crius.Logger.Error("UpdateBranchTagStatus 请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	err := model.UpdateBranchTagStatus(id, request.Status)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBranchTagStatus 更新标签数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新标签失败"
		return resp, nil
	}
	return resp, nil
}

// GetBranchTags 获取标签列表
func (s *Server) GetBranchTags(ctx context.Context, request *proto.GetBranchTagsRequest) (*proto.GetBranchTagsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBranchTags")
	resp := &proto.GetBranchTagsResponse{
		ErrorCode: pkgs.Success,
	}

	if request.Status != "" && request.Status != model.StatusClosed && request.Status != model.StatusOpened {
		crius.Logger.Error("GetBranchTags 请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}
	branchIDs := make(fields.UUIDArr, len(request.BranchIds))
	for i := range request.BranchIds {
		branchIDs[i] = uuid.FromStringOrNil(request.BranchIds[i])
		if branchIDs[i] == uuid.Nil {
			crius.Logger.Error(fmt.Sprintf("GetBranchTags 请求参数错误:%v", request.BranchIds))
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "请求参数错误"
			return resp, nil
		}
	}

	tag := model.TableBranchTag{
		Name:     request.Name,
		Status:   crius.Status(request.Status),
		Branches: &branchIDs,
	}
	tags, count, err := model.GetBranchTags(tag, request.DateStart, request.DateEnd, request.Offset, request.Limit)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranchTags 获取标签数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取标签失败"
		return resp, nil
	}
	resp.Data = &proto.GetBranchTagsData{Total: int32(count)}

	for i := range tags {
		resp.Data.BranchTags = append(resp.Data.BranchTags, toProtoBranchTag(&tags[i]))
	}
	return resp, nil
}

// GetBranchTagsByIDs 根据id列表获取标签列表
func (s *Server) GetBranchTagsByIDs(ctx context.Context, request *proto.GetBranchTagsByIDsRequest) (*proto.GetBranchTagsByIDsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBranchTagsByIDs")
	resp := &proto.GetBranchTagsByIDsResponse{
		ErrorCode: pkgs.Success,
	}

	if len(request.Ids) == 0 {
		crius.Logger.Error("GetBranchTagsByIDs 请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	ids := make([]uuid.UUID, len(request.Ids))
	for i := range request.Ids {
		ids[i] = uuid.FromStringOrNil(request.Ids[i])
		if ids[i] == uuid.Nil {
			crius.Logger.Error(fmt.Sprintf("GetBranchTagsByIDs 请求参数错误:%v", request.Ids))
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "请求参数错误"
			return resp, nil
		}
	}

	tags, err := model.GetBranchTagsByIDs(ids)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranchTagsByIDs 获取标签数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取标签失败"
		return resp, nil
	}

	for i := range tags {
		resp.Data = append(resp.Data, toProtoBranchTag(&tags[i]))
	}
	return resp, nil
}

// ShowBranchTag 根据id列表获取标签列表
func (s *Server) ShowBranchTag(ctx context.Context, request *proto.ShowBranchTagRequest) (*proto.ShowBranchTagResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowBranchTag")
	resp := &proto.ShowBranchTagResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil {
		crius.Logger.Error("ShowBranchTag 请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	tag, err := model.ShowBranchTag(id)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("ShowBranchTag 获取标签数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取标签失败"
		return resp, nil
	}

	resp.Data = toProtoBranchTag(tag)
	return resp, nil
}

func toProtoBranchTag(tag *model.BranchTag) *proto.BranchTagData {
	branches := tag.Branches.ToStringArr()
	return &proto.BranchTagData{
		Id:            tag.ID.String(),
		Name:          tag.Name,
		Branches:      branches,
		CreateStaffId: crius.UUIDToString(tag.CreateStaffID),
		StaffName:     tag.StaffName,
		Status:        tag.Status.String(),
		CreatedAt:     int32(tag.CreatedAt.Unix()),
		UpdatedAt:     int32(tag.UpdatedAt.Unix()),
	}
}
