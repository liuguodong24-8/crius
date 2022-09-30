package rpc

import (
	"context"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// CreateRole 创建角色
func (s *Server) CreateRole(ctx context.Context, request *proto.CreateRoleRequest) (*proto.CreateRoleResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateRole")
	resp := &proto.CreateRoleResponse{
		ErrorCode: pkgs.Success,
	}
	operaterID := pkgs.GetMetadata(ctx).StaffID
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if request.Name == "" ||
		(request.Property != 1 && request.Property != 2) ||
		len(request.Permissions) == 0 ||
		operaterID == uuid.Nil ||
		merchantID == uuid.Nil {
		crius.Logger.Error("CreateRole rpc请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	role := model.TableRole{
		ID:         uuid.NewV4(),
		Name:       request.Name,
		Status:     model.StatusOpened,
		Property:   int8(request.Property),
		StaffID:    &operaterID,
		MerchantID: &merchantID,
	}

	mappingPermissions := make([]model.TableMappingRolePermission, len(request.Permissions))
	permissionIDs := make([]int32, len(request.Permissions))
	for i := range request.Permissions {
		mappingPermissions[i].ID = uuid.NewV4()
		mappingPermissions[i].PermissionID = request.Permissions[i]
		mappingPermissions[i].RoleID = role.ID
		permissionIDs[i] = request.Permissions[i]
	}

	if _, _, err := model.ShowRoleByName(role.Name, merchantID); err != nil {
		if err != gorm.ErrRecordNotFound { //查询数据报错且不是没有找到数据错误
			crius.Logger.Error(fmt.Sprintf("CreateRole 校验用户输入角色名称数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "校验角色名称错误"
			return resp, nil
		}
	} else { //正确查询到数据
		crius.Logger.Error("CreateRole 用户输入角色名称已经存在")
		resp.ErrorMessage = "角色名称已存在"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	err := model.CreateRole(role, mappingPermissions)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateRole 新建角色错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "新建角色错误"
		return resp, nil
	}

	roleInfo := model.RoleInfo{TableRole: role}
	staffFull, err := model.ShowStaffByID(operaterID)
	if err == nil {
		roleInfo.StaffName = staffFull.Staff.Name
	}

	permissions, err := model.GetPermissionsByID(permissionIDs)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateRole 快照查询权限数据库错误:%v", err))
		resp.ErrorCode = 0
		return resp, nil
	}

	m := make(map[string]interface{})
	m["role"] = roleInfo
	m["permission"] = permissions
	staffID := pkgs.GetMetadata(ctx).StaffID
	after := pkgs.MakeParams(m)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: model.TableRole{}.TableName(),
		After:             &after,
		TableID:           &role.ID,
		Method:            "create",
	}

	return resp, nil
}

// UpdateRole 更新角色信息
func (s *Server) UpdateRole(ctx context.Context, request *proto.UpdateRoleRequest) (*proto.UpdateRoleResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateRole")
	resp := &proto.UpdateRoleResponse{
		ErrorCode: pkgs.Success,
	}
	roleID := uuid.FromStringOrNil(request.Id)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if request.Name == "" ||
		(request.Property != 1 && request.Property != 2) ||
		len(request.Permissions) == 0 ||
		roleID == uuid.Nil ||
		merchantID == uuid.Nil {
		crius.Logger.Error("UpdateRole rpc请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	role := model.TableRole{
		ID:       roleID,
		Name:     request.Name,
		Property: int8(request.Property),
	}

	mappingPermissions := make([]model.TableMappingRolePermission, len(request.Permissions))
	afterPermissionIDs := make([]int32, len(request.Permissions))
	for i := range request.Permissions {
		mappingPermissions[i].ID = uuid.NewV4()
		mappingPermissions[i].PermissionID = request.Permissions[i]
		mappingPermissions[i].RoleID = roleID
		afterPermissionIDs[i] = request.Permissions[i]
	}

	var tabRole *model.RoleInfo
	var tabPermissions []model.TableMappingRolePermission
	var err error
	if tabRole, tabPermissions, err = model.ShowRoleByName(role.Name, merchantID); err != nil {
		if err != gorm.ErrRecordNotFound { //查询数据报错且不是没有找到数据错误
			crius.Logger.Error(fmt.Sprintf("UpdateRole 校验用户输入角色名称数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "校验角色名称错误"
			return resp, nil
		}
	} else { //正确查询到数据
		if tabRole.ID != roleID {
			crius.Logger.Error("UpdateRole 用户输入角色名称已经存在")
			resp.ErrorMessage = "角色名称已存在"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	err = model.UpdateRole(role, mappingPermissions)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateRole 更新角色数据库错误:%v", err))
		resp.ErrorMessage = "更新角色错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	//获取修改之前的权限
	beforePermissionIDs := make([]int32, len(tabPermissions))
	for i := range tabPermissions {
		beforePermissionIDs[i] = tabPermissions[i].PermissionID
	}
	beforePermissions, err := model.GetPermissionsByID(beforePermissionIDs)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateRole 快照查询权限数据库错误:%v", err))
		return resp, nil
	}
	beforeMap := make(map[string]interface{})
	beforeMap["role"] = *tabRole
	beforeMap["permission"] = beforePermissions

	//获取修改之后的权限
	afterPermissions, err := model.GetPermissionsByID(afterPermissionIDs)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateRole 快照查询权限数据库错误:%v", err))
		return resp, nil
	}
	tabRole.Name = request.Name
	tabRole.Property = int8(request.Property)
	afterMap := make(map[string]interface{})
	afterMap["role"] = *tabRole
	afterMap["permission"] = afterPermissions
	staffID := pkgs.GetMetadata(ctx).StaffID
	before := pkgs.MakeParams(beforeMap)
	after := pkgs.MakeParams(afterMap)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: model.TableRole{}.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &role.ID,
		Method:            "update",
	}

	return resp, nil
}

// UpdateRoleStatus 更新角色状态
func (s *Server) UpdateRoleStatus(ctx context.Context, request *proto.UpdateRoleStatusRequest) (*proto.UpdateRoleStatusResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateRoleStatus")
	resp := &proto.UpdateRoleStatusResponse{
		ErrorCode: pkgs.Success,
	}
	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil || (request.Status != model.StatusOpened && request.Status != model.StatusClosed) {
		crius.Logger.Error("UpdateRoleStatus rpc请求参数错误")
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	tabRole, _, err := model.ShowRoleByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "角色不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdateRoleStatus 查询角色数据库错误:%v", err))
		resp.ErrorMessage = "更新角色状态错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	err = model.UpdateRoleStatus(id, request.Status)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateRoleStatus 更新角色状态数据库错误:%v", err))
		resp.ErrorMessage = "更新角色状态错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	beforeMap := make(map[string]interface{})
	beforeMap["role"] = *tabRole

	tabRole.Status = crius.Status(request.Status)
	afterMap := make(map[string]interface{})
	afterMap["role"] = *tabRole
	staffID := pkgs.GetMetadata(ctx).StaffID
	before := pkgs.MakeParams(beforeMap)
	after := pkgs.MakeParams(afterMap)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: model.TableRole{}.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &id,
		Method:            "update_status",
	}

	return resp, nil
}

// GetRoles 获取角色列表
func (s *Server) GetRoles(ctx context.Context, request *proto.GetRolesRequest) (*proto.GetRolesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetRoles")
	resp := &proto.GetRolesResponse{
		ErrorCode: pkgs.Success,
	}

	staffID := uuid.FromStringOrNil(request.StaffId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	status := strings.ToLower(request.Status)
	if (status != model.StatusOpened && status != model.StatusClosed && status != "") || merchantID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("GetRoles 请求参数错误 status:%v, merchantID:%v", status, merchantID))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}
	if staffID != uuid.Nil {
		staff, err := model.ShowBasicStaffByID(staffID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				crius.Logger.Error(fmt.Sprintf("GetRoles 查询员工不存在,id:%v", staffID))
				resp.ErrorCode = pkgs.ErrNotFound
				resp.ErrorMessage = "当前登录员工不存在"
				return resp, nil
			}
			crius.Logger.Error(fmt.Sprintf("GetRoles 查询员工数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "查询员工信息错误"
			return resp, nil
		}
		if staff.Admin {
			staffID = uuid.Nil
		}
	}

	roles, count, err := model.GetRoles(request.Name, status, staffID, merchantID, request.Offset, request.Limit)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetRoles 查询角色信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "查询角色失败"
		return resp, nil
	}
	resp.Data = &proto.RolesData{Total: int32(count)}
	for _, v := range roles {
		resp.Data.Roles = append(resp.Data.Roles, &proto.RoleInfo{
			Id:        v.ID.String(),
			Name:      v.Name,
			Status:    v.Status.String(),
			Property:  int32(v.Property),
			StaffId:   crius.UUIDToString(v.StaffID),
			StaffName: v.StaffName,
		})
	}
	return resp, nil
}

// ShowRole 查询单个角色信息
func (s *Server) ShowRole(ctx context.Context, request *proto.ShowRoleRequest) (*proto.ShowRoleResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetRoles")
	resp := &proto.ShowRoleResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil {
		crius.Logger.Error("ShowRole rpc请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	role, permissions, err := model.ShowRoleByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "角色不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("查询角色信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "查询角色失败"
		return resp, nil
	}
	resp.Data = &proto.RolePermissionInfo{}
	resp.Data.Role = &proto.RoleInfo{
		Id:        role.ID.String(),
		Name:      role.Name,
		Status:    role.Status.String(),
		Property:  int32(role.Property),
		StaffId:   crius.UUIDToString(role.StaffID),
		StaffName: role.StaffName,
	}
	for _, v := range permissions {
		resp.Data.Permissions = append(resp.Data.Permissions, v.PermissionID)
	}

	return resp, nil
}

// DeleteRole 删除角色
func (s *Server) DeleteRole(ctx context.Context, request *proto.DeleteRoleRequest) (*proto.DeleteRoleResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("DeleteRole")
	resp := &proto.DeleteRoleResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil {
		crius.Logger.Error("DeleteRole rpc请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	tabRole, tabPermissions, err := model.ShowRoleByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "角色不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("DeleteRole 查询角色数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除角色失败"
		return resp, nil
	}

	err = model.DeleteRole(id)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("DeleteRole 删除角色数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除角色失败"
		return resp, nil
	}

	beforePermissionIDs := make([]int32, len(tabPermissions))
	for i := range tabPermissions {
		beforePermissionIDs[i] = tabPermissions[i].PermissionID
	}
	beforePermissions, err := model.GetPermissionsByID(beforePermissionIDs)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("DeleteRole 快照查询权限数据库错误:%v", err))
		return resp, nil
	}
	beforeMap := make(map[string]interface{})
	beforeMap["role"] = *tabRole
	beforeMap["permission"] = beforePermissions
	before := pkgs.MakeParams(beforeMap)
	staffID := pkgs.GetMetadata(ctx).StaffID
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: model.TableRole{}.TableName(),
		Before:            &before,
		TableID:           &id,
		Method:            "delete",
	}

	return resp, nil
}

// GetRoleHistories 获取角色修改历史
func (s *Server) GetRoleHistories(ctx context.Context, request *proto.GetRoleHistoriesRequest) (*proto.GetRoleHistoriesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetRoleHistories")
	resp := &proto.GetRoleHistoriesResponse{
		ErrorCode: pkgs.Success,
	}
	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("GetRoleHistories 请求参数错误:%v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	snapshots, err := model.GetSnapshots(id)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetRoleHistories 获取快照数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取角色修改历史失败"
		return resp, nil
	}

	role, _, err := model.ShowRoleByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "角色不存在"
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("GetRoleHistories 获取角色数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取角色失败"
		return resp, nil
	}

	resp.Data = &proto.RoleHistoriesData{Role: &proto.RoleInfo{Id: role.ID.String(), Name: role.Name, Status: role.Status.String(),
		Property: int32(role.Property), StaffId: crius.UUIDToString(role.StaffID), StaffName: role.StaffName}}
	for _, v := range snapshots {
		resp.Data.Snapshots = append(resp.Data.Snapshots, &proto.Snapshot{
			StaffId:   v.StaffID.String(),
			StaffName: v.StaffName,
			Before:    v.Before.JSON(),
			After:     v.After.JSON(),
			CreatedAt: int32(v.CreatedAt.Unix()),
			Method:    v.Method,
		})
	}
	return resp, nil
}

func toProtoRole(role *model.RoleInfo) *proto.RoleInfo {
	return &proto.RoleInfo{
		Id:        role.ID.String(),
		Name:      role.Name,
		Status:    role.Status.String(),
		Property:  int32(role.Property),
		StaffId:   crius.UUIDToString(role.StaffID),
		StaffName: role.StaffName,
	}
}
