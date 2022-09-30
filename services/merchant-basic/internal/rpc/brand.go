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

// CreateBrand 创建品牌
func (s *Server) CreateBrand(ctx context.Context, req *proto.CreateBrandRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreateBrand")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	staffID := pkgs.GetMetadata(ctx).StaffID
	_, err := model.ShowBrandByName(req.Name)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("CreateBrand 获取品牌数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "创建品牌失败"
			return resp, nil
		}
	} else {
		crius.Logger.Error(fmt.Sprintf("CreateBrand 品牌名已存在:%v", req.Name))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "品牌名已存在"
		return resp, nil
	}

	brand := model.TableBrand{
		ID:     uuid.NewV4(),
		Name:   req.Name,
		Status: crius.StringToStatus(req.Status),
		Order:  req.Order,
	}
	err = model.CreateBrand(brand)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateBrand 创建品牌数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建品牌失败"
		return resp, nil
	}
	after := pkgs.MakeParams(brand)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: brand.TableName(),
		After:             &after,
		TableID:           &brand.ID,
		Method:            "create",
	}

	return resp, nil
}

// UpdateBrand 更新品牌
func (s *Server) UpdateBrand(ctx context.Context, req *proto.UpdateBrandRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateBrand")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	staffID := pkgs.GetMetadata(ctx).StaffID
	id := uuid.FromStringOrNil(req.Id)

	brand, err := model.ShowBrandByName(req.Name)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("UpdateBrand 获取品牌数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "更新品牌失败"
			return resp, nil
		}
	} else {
		if brand.ID != id {
			crius.Logger.Error(fmt.Sprintf("UpdateBrand 品牌名已存在:%v", req.Name))
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "品牌名已存在"
			return resp, nil
		}
	}
	before := pkgs.MakeParams(brand)

	brand.Name = req.Name
	brand.Status = crius.StringToStatus(req.Status)
	brand.Order = req.Order
	brand.ID = id
	err = model.UpdateBrand(*brand)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBrand 更新品牌数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新品牌失败"
		return resp, nil
	}
	after := pkgs.MakeParams(brand)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: brand.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &brand.ID,
		Method:            "update",
	}

	return resp, nil
}

// UpdateBrandStatus 更新品牌状态
func (s *Server) UpdateBrandStatus(ctx context.Context, req *proto.UpdateBrandStatusRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateBrandStatus")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	staffID := pkgs.GetMetadata(ctx).StaffID
	brand, err := model.ShowBrand(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			crius.Logger.Error("UpdateBrandStatus 更新品牌状态数据不存在")
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "品牌不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdateBrandStatus 更新品牌状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新品牌状态失败"
		return resp, nil
	}
	before := pkgs.MakeParams(brand)
	err = model.UpdateBrandStatus(id, req.Status)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateBrandStatus 更新品牌状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新品牌状态失败"
		return resp, nil
	}
	brand.Status = crius.Status(req.Status)
	after := pkgs.MakeParams(brand)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: brand.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &brand.ID,
		Method:            "update",
	}

	return resp, nil
}

// GetBrands 获取品牌列表
func (s *Server) GetBrands(ctx context.Context, req *proto.GetBrandsRequest) (*proto.GetBrandsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetBrands")
	resp := &proto.GetBrandsResponse{
		ErrorCode: pkgs.Success,
	}

	brands, count, err := model.GetBrands(req.Name, req.Status, req.Offset, req.Limit)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBrands 获取品牌列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取品牌列表失败"
		return resp, nil
	}

	resp.Data = &proto.GetBrandsData{
		Total: int32(count),
	}

	for _, v := range brands {
		resp.Data.Brands = append(resp.Data.Brands, toProtoBrand(v))
	}

	return resp, nil
}

func toProtoBrand(brand model.TableBrand) *proto.Brand {
	var createdAt int32
	if brand.CreatedAt != nil {
		createdAt = int32(brand.CreatedAt.Unix())
	}
	return &proto.Brand{
		Id:        brand.ID.String(),
		Name:      brand.Name,
		Order:     brand.Order,
		Status:    brand.Status.String(),
		CreatedAt: createdAt,
	}
}
