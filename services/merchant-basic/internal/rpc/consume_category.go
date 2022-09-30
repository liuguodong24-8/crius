package rpc

import (
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// CreateConsumeCategory 新增消费类型
func (s *Server) CreateConsumeCategory(ctx context.Context, req *proto.CreateConsumeCategoryRequest) (*proto.Response, error) {
	defer util.CatchException()

	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("CreateConsumeCategory")
	metadata := pkgs.GetMetadata(ctx)

	exists, err := s.checkConsumeCategory(ctx, req.Category, req.Code, "")
	if !exists && err != nil {
		util.Logger.WithError(err).Error("新增积分分类，校验错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("新增积分分类，校验错误:%s", err.Error()),
		}, nil
	}

	if exists && err != nil {
		return &proto.Response{
			ErrorCode:    pkgs.ErrCheck,
			ErrorMessage: err.Error(),
		}, nil
	}

	if req.IsRoomFee {
		_, err := model.ShowConsumeCategoryRoomFee()
		if err != nil && err != gorm.ErrRecordNotFound {
			return &proto.Response{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: fmt.Sprintf("新增积分分类，校验错误:%s", err.Error()),
			}, nil
		}
		if err == nil {
			return &proto.Response{
				ErrorCode:    pkgs.ErrUnprocessableEntity,
				ErrorMessage: fmt.Sprintf("房费类型已存在", err.Error()),
			}, nil
		}
	}

	operatorTypes, err := fields.StringArrToUUIDArr(req.OperatorTypes)
	if err != nil {
		return &proto.Response{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: fmt.Sprintf("新增积分分类，参数错误:%s", err.Error()),
		}, nil
	}
	activeTypes, err := fields.StringArrToUUIDArr(req.ActiveTypes)
	if err != nil {
		return &proto.Response{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: fmt.Sprintf("新增积分分类，参数错误:%s", err.Error()),
		}, nil
	}
	category := model.ConsumeCategory{
		ID:            uuid.NewV4(),
		MerchantID:    &metadata.MerchantID,
		Category:      req.Category,
		Code:          req.Code,
		Status:        util.StringToStatus(req.Status),
		OperatorTypes: &operatorTypes,
		ActiveTypes:   &activeTypes,
		IsRoomFee:     req.IsRoomFee,
	}

	if err := model.DatabaseConn().Create(&category).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("保存积分设置类型错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("保存数据错误:%s", err.Error()),
		}, nil
	}

	after := pkgs.MakeParams(category)

	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &metadata.StaffID,
		SleuthCode:        metadata.SleuthCode,
		SnapShotTableName: category.TableName(),
		After:             &after,
		TableID:           &category.ID,
		Method:            "create",
	}

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// ListConsumeCategory 积分类型列表
func (s *Server) ListConsumeCategory(ctx context.Context, req *proto.ListConsumeCategoryRequest) (*proto.ListConsumeCategoryResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("ListConsumeCategory")
	metadata := pkgs.GetMetadata(ctx)

	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", metadata.MerchantID))
	if len(req.Category) > 0 {
		scopes = append(scopes, util.ColumnLikeScope("category", req.Category))
	}
	if len(req.Status) > 0 {
		scopes = append(scopes, util.ColumnEqualScope("status", req.Status))
	}

	var total int64
	var items []model.ConsumeCategory

	model.DatabaseConn().Model(&model.ConsumeCategory{}).Scopes(scopes...).Count(&total)
	// 是否需要分页
	if req.WithPage {
		scopes = append(scopes, util.PaginationScope(req.Offset, req.Limit))
	}
	orderBy := "created_at desc"
	if len(req.OrderBy) > 0 {
		orderBy = req.OrderBy
	}

	if err := model.DatabaseConn().Model(&model.ConsumeCategory{}).Scopes(scopes...).Order(orderBy).Find(&items).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("积分类型列表查询错误")
		return &proto.ListConsumeCategoryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询数据错误:%s", err.Error()),
		}, nil
	}

	var data []*proto.ConsumeCategory
	for _, item := range items {
		data = append(data, toProtoCategory(item))
	}

	return &proto.ListConsumeCategoryResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.ListConsumeCategoryResponse_Data{
			Data:  data,
			Total: total,
		},
	}, nil
}

// GetConsumeCategoryTypes 获取积分类型分组列表
func (s *Server) GetConsumeCategoryTypes(ctx context.Context, req *proto.Empty) (*proto.GetConsumeCategoryTypesResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("GetConsumeCategoryTypes")
	metadata := pkgs.GetMetadata(ctx)

	categories, err := model.GetConsumeCategoryTypes(metadata.MerchantID)
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error(fmt.Sprintf("积分类型分组列表查询错误:%v", err))
		return &proto.GetConsumeCategoryTypesResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询数据错误:%s", err.Error()),
		}, nil
	}

	var activeTypes, operatorTypes []string
	isRoomFee := false
	for i := range categories {
		actives := categories[i].ActiveTypes.ToStringArr()
		operators := categories[i].OperatorTypes.ToStringArr()
		if len(actives) != 0 {
			for _, v := range actives {
				activeTypes = append(activeTypes, v)
			}
		}
		if len(operators) != 0 {
			for _, v := range operators {
				operatorTypes = append(operatorTypes, v)
			}
		}
		if categories[i].IsRoomFee {
			isRoomFee = true
		}
	}

	return &proto.GetConsumeCategoryTypesResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.GetConsumeCategoryTypesResponse_Data{
			OperatorTypes: operatorTypes,
			ActiveTypes:   activeTypes,
			IsRoomFee:     isRoomFee,
		},
	}, nil
}

func toProtoCategory(category model.ConsumeCategory) *proto.ConsumeCategory {
	return &proto.ConsumeCategory{
		Id:            category.ID.String(),
		Category:      category.Category,
		Code:          category.Code,
		Status:        category.Status.String(),
		CreatedAt:     category.CreatedAt.Unix(),
		OperatorTypes: category.OperatorTypes.ToStringArr(),
		ActiveTypes:   category.ActiveTypes.ToStringArr(),
		IsRoomFee:     category.IsRoomFee,
	}
}

// UpdateConsumeCategory 修改积分类型
func (s *Server) UpdateConsumeCategory(ctx context.Context, req *proto.UpdateConsumeCategoryRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("UpdateConsumeCategory")

	operatorTypes, err := fields.StringArrToUUIDArr(req.OperatorTypes)
	if err != nil {
		return &proto.Response{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: fmt.Sprintf("新增积分分类，参数错误:%s", err.Error()),
		}, nil
	}
	activeTypes, err := fields.StringArrToUUIDArr(req.ActiveTypes)
	if err != nil {
		return &proto.Response{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: fmt.Sprintf("新增积分分类，参数错误:%s", err.Error()),
		}, nil
	}
	var category model.ConsumeCategory
	if err := model.DatabaseConn().Model(&model.ConsumeCategory{}).Scopes(util.ColumnEqualScope("id", req.Id)).First(&category).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.Response{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: fmt.Sprintf("积分类型对应数据不存在:%s", req.Id),
			}, nil
		}

		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询积分类型错误:%s", err.Error()),
		}, nil
	}
	if !category.IsRoomFee && req.IsRoomFee {
		_, err := model.ShowConsumeCategoryRoomFee()
		if err != nil && err != gorm.ErrRecordNotFound {
			return &proto.Response{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: fmt.Sprintf("新增积分分类，校验错误:%s", err.Error()),
			}, nil
		}
		if err == nil {
			return &proto.Response{
				ErrorCode:    pkgs.ErrUnprocessableEntity,
				ErrorMessage: fmt.Sprintf("房费类型已存在", err.Error()),
			}, nil
		}
	}

	exists, err := s.checkConsumeCategory(ctx, req.Category, req.Code, category.ID.String())
	if !exists && err != nil {
		util.Logger.WithError(err).Error("修改积分分类，校验错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("修改积分分类，校验错误:%s", err.Error()),
		}, nil
	}

	if exists && err != nil {
		return &proto.Response{
			ErrorCode:    pkgs.ErrCheck,
			ErrorMessage: err.Error(),
		}, nil
	}

	before := pkgs.MakeParams(category)
	category.Status = util.StringToStatus(req.Status)
	category.Category = req.Category
	category.Code = req.Code
	category.OperatorTypes = &operatorTypes
	category.ActiveTypes = &activeTypes
	category.IsRoomFee = req.IsRoomFee
	after := pkgs.MakeParams(req)

	if err := model.DatabaseConn().Save(&category).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("修改保存积分类型错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("修改保存积分类型错误:%s", err.Error()),
		}, nil
	}

	metadata := pkgs.GetMetadata(ctx)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &metadata.StaffID,
		SleuthCode:        metadata.SleuthCode,
		SnapShotTableName: category.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &category.ID,
		Method:            "update",
	}

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// UpdateConsumeCategoryStatus 修改积分类型状态
func (s *Server) UpdateConsumeCategoryStatus(ctx context.Context, req *proto.UpdateStatusRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("UpdateConsumeCategoryStatus")
	var category model.ConsumeCategory
	if err := model.DatabaseConn().Model(&model.ConsumeCategory{}).Scopes(util.ColumnEqualScope("id", req.Id)).First(&category).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.Response{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: fmt.Sprintf("积分类型对应数据不存在:%s", req.Id),
			}, nil
		}

		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询积分类型错误:%s", err.Error()),
		}, nil
	}

	before := pkgs.MakeParams(category)
	category.Status = util.StringToStatus(req.Status)
	after := pkgs.MakeParams(req)

	if err := model.DatabaseConn().Save(&category).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("修改积分类型状态错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("修改积分类型错误:%s", err.Error()),
		}, nil
	}

	metadata := pkgs.GetMetadata(ctx)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &metadata.StaffID,
		SleuthCode:        metadata.SleuthCode,
		SnapShotTableName: category.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &category.ID,
		Method:            "update",
	}

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// ShowConsumeCategory 积分类型详情
func (s *Server) ShowConsumeCategory(ctx context.Context, req *proto.ShowConsumeCategoryRequest) (*proto.ShowConsumeCategoryResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("ShowConsumeCategory")
	var category model.ConsumeCategory
	if err := model.DatabaseConn().Model(&model.ConsumeCategory{}).Scopes(util.ColumnEqualScope("id", req.Id)).First(&category).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.ShowConsumeCategoryResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: fmt.Sprintf("积分类型对应数据不存在:%s", req.Id),
			}, nil
		}

		return &proto.ShowConsumeCategoryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询积分类型错误:%s", err.Error()),
		}, nil
	}

	return &proto.ShowConsumeCategoryResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         toProtoCategory(category),
	}, nil
}

func (s *Server) checkConsumeCategory(ctx context.Context, category, code, ignoreID string) (bool, error) {
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, util.ColumnEqualScope("merchant_id", pkgs.GetMetadata(ctx).MerchantID))
	if len(ignoreID) > 0 {
		scopes = append(scopes, util.ColumnSymbolScope("id", "!=", ignoreID))
	}

	var total int64
	// 判断类型名称
	if err := model.DatabaseConn().Model(&model.ConsumeCategory{}).Scopes(scopes...).Where("category", category).Count(&total).Error; nil != err {
		return false, err
	}

	if total > 0 {
		return true, errors.New("类型名称已存在")
	}

	// 判断类型编码
	if err := model.DatabaseConn().Model(&model.ConsumeCategory{}).Scopes(scopes...).Where("code", code).Count(&total).Error; nil != err {
		return false, err
	}

	if total > 0 {
		return true, errors.New("类型编码已存在")
	}

	return false, nil
}
