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
)

// GetPermissions 获取权限列表
func (s *Server) GetPermissions(ctx context.Context, request *proto.GetPermissionsRequest) (*proto.GetPermissionsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetPermissions")
	resp := &proto.GetPermissionsResponse{
		ErrorCode: pkgs.Success,
	}

	permissions, err := model.GetPermissionsByStaffID(uuid.FromStringOrNil(request.Id), request.Service)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetPermissions 获取权限列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取权限列表失败"
		return resp, nil
	}

	for _, v := range permissions {
		resp.Data = append(resp.Data, &proto.PermissionInfo{
			Id:         v.ID,
			Permission: v.Permission,
			Service:    v.Service,
		})
	}

	return resp, nil
}

// CreatePermissions CreatePermissions
func (s *Server) CreatePermissions(ctx context.Context, request *proto.CreatePermissionsRequest) (*proto.CreatePermissionsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreatePermissions")
	resp := &proto.CreatePermissionsResponse{
		ErrorCode: pkgs.Success,
	}

	if len(request.Permissions) == 0 ||
		request.Service == "" {
		crius.Logger.Error(fmt.Sprintf("CreatePermissions 请求参数错误"))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	permissions := make([]model.TablePermission, 0)
	for _, v := range request.Permissions {
		permissions = append(permissions, model.TablePermission{
			ID:         v.Id,
			Permission: v.Permission,
			Service:    request.Service,
		})
	}
	err := model.SavePermissions(permissions, request.Service)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreatePermissions 存储权限数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建权限失败"
		return resp, nil
	}

	return resp, nil
}
