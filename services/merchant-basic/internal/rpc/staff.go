package rpc

import (
	"context"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/util"
	"gorm.io/gorm"
)

// CreateStaff CreateStaff
func (s *Server) CreateStaff(ctx context.Context, request *proto.CreateStaffRequest) (*proto.CreateStaffResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateStaff")
	resp := &proto.CreateStaffResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID

	entryAt := time.Unix(int64(request.Staff.EntryAt), 0)
	staff := model.TableStaff{
		ID:           uuid.NewV4(),
		EmployeeCode: request.Staff.EmployeeCode,
		Name:         request.Staff.Name,
		Phone:        request.Staff.Phone,
		PhoneCode:    request.Staff.PhoneCode,
		Gender:       int8(request.Staff.Gender),
		Status:       model.StatusOpened,
		EntryAt:      &entryAt,
		Salt:         util.GenerateSalt(),
		MerchantID:   &merchantID,
	}

	staff.Password = util.GetMD5(fmt.Sprintf("%s:%s", util.DefaultPassword, staff.Salt))

	// 校验手机号唯一性
	if err := model.ShowStaffExistsByPhone(request.Staff.Phone, request.Staff.PhoneCode, merchantID); err != nil {
		if err != gorm.ErrRecordNotFound { //查询数据报错且不是没有找到数据错误
			crius.Logger.Error(fmt.Sprintf("CreateStaff 校验用户输入手机号数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "校验用户输入手机号错误"
			return resp, nil
		}
	} else { //正确查询到数据
		crius.Logger.Error("CreateStaff 用户输入手机号已经存在")
		resp.ErrorMessage = "用户输入手机号已经存在"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	// 校验工号/编号唯一性
	if err := model.ShowStaffExistsByCode(staff.EmployeeCode, merchantID); err != nil {
		if err != gorm.ErrRecordNotFound { //查询数据报错且不是没有找到数据错误
			crius.Logger.Error(fmt.Sprintf("CreateStaff 校验用户输入工号数据库错误:%v", err))
			resp.ErrorMessage = "校验用户输入工号错误"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	} else { //正确查询到数据
		crius.Logger.Error("CreateStaff 用户输入工号已经存在")
		resp.ErrorMessage = "用户输入工号已经存在"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	for i := 0; i < 10; i++ {
		// 获取数据库员工code序列号
		num := model.GetStaffSequence()
		if num < 0 {
			resp.ErrorMessage = "用户编号生成失败"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
		// 生成员工code
		code := util.GenerateStaffCode(num)
		// 查询员工code是否已存在
		if err := model.ShowStaffExistsByCode(code, merchantID); err != nil {
			if err != gorm.ErrRecordNotFound { //查询数据报错且不是没有找到数据错误
				crius.Logger.Error(fmt.Sprintf("CreateStaff 校验员工生成code数据库错误:%v", err))
				resp.ErrorMessage = "校验用户生成编号失败"
				resp.ErrorCode = pkgs.ErrInternal
				return resp, nil
			}
			staff.Code = code
			break
		}
		// 此处员工code已存在，重新生成code
	}

	if staff.Code == "" {
		crius.Logger.Error("CreateStaff 员工code生成10次已存在，创建员工失败")
		resp.ErrorMessage = "生成用户编号已存在"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	mappingBranches := make([]model.TableMappingStaffBranch, 0)
	for _, v := range request.Branches {
		if b := uuid.FromStringOrNil(v); b != uuid.Nil {
			mappingBranches = append(mappingBranches, model.TableMappingStaffBranch{
				ID:       uuid.NewV4(),
				StaffID:  staff.ID,
				BranchID: b,
			})
		} else {
			crius.Logger.Error(fmt.Sprintf("CreateStaff rpc请求门店错误:%v", request.Branches))
			resp.ErrorMessage = "门店ID错误"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	mappingRoles := make([]model.TableMappingStaffRole, 0)
	for _, v := range request.Roles {
		if b := uuid.FromStringOrNil(v); b != uuid.Nil {
			mappingRoles = append(mappingRoles, model.TableMappingStaffRole{
				ID:      uuid.NewV4(),
				StaffID: staff.ID,
				RoleID:  b,
			})
		} else {
			crius.Logger.Error(fmt.Sprintf("CreateStaff rpc请求角色错误:%v", request.Roles))
			resp.ErrorMessage = "角色ID错误"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	err := model.CreateStaff(staff, mappingBranches, mappingRoles)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateStaff 新增员工数据库操作错误:%v", err))
		resp.ErrorMessage = "新增员工错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	//快照
	m := make(map[string]interface{})
	m["staff"] = staff
	m["branches"] = mappingBranches
	m["roles"] = mappingRoles
	after := pkgs.MakeParams(m)
	staffID := pkgs.GetMetadata(ctx).StaffID
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: staff.TableName(),
		After:             &after,
		TableID:           &staff.ID,
		Method:            "create",
	}

	return resp, nil
}

// UpdateStaff UpdateStaff
func (s *Server) UpdateStaff(ctx context.Context, request *proto.UpdateStaffRequest) (*proto.UpdateStaffResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateStaff")
	resp := &proto.UpdateStaffResponse{
		ErrorCode: pkgs.Success,
	}

	// 判断参数合法性
	if request.Staff == nil {
		crius.Logger.Error("UpdateStaff rpc请求参数nil")
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	id := uuid.FromStringOrNil(request.Staff.Id)
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	entryAt := time.Unix(int64(request.Staff.EntryAt), 0)
	staff := model.TableStaff{
		ID:        id,
		Name:      request.Staff.Name,
		Phone:     request.Staff.Phone,
		PhoneCode: request.Staff.PhoneCode,
		Gender:    int8(request.Staff.Gender),
		EntryAt:   &entryAt,
	}

	staffFullInfo, err := model.ShowStaffByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "员工不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdateStaff 获取员工信息数据库操作错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新员工错误"
		return resp, nil
	}

	// 校验手机号唯一性
	if err := model.ShowStaffExistsByPhone(request.Staff.Phone, request.Staff.PhoneCode, merchantID); err != nil {
		if err != gorm.ErrRecordNotFound { //查询数据报错且不是没有找到数据错误
			crius.Logger.Error(fmt.Sprintf("CreateStaff 校验用户输入手机号数据库错误:%v", err))
			resp.ErrorMessage = "校验用户输入手机号错误"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	} else { //正确查询到数据
		// 有一个字段不等，手机号则改变了
		if staffFullInfo.Staff.Phone != request.Staff.Phone || staffFullInfo.Staff.PhoneCode != request.Staff.PhoneCode {
			crius.Logger.Error("CreateStaff 用户输入手机号已经存在")
			resp.ErrorMessage = "用户输入手机号已经存在"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	quitAt := time.Unix(int64(request.Staff.QuitAt), 0)
	if request.Staff.QuitAt != 0 {
		staff.QuitAt = &quitAt
	}

	mappingBranches := make([]model.TableMappingStaffBranch, 0)
	for _, v := range request.Branches {
		if b := uuid.FromStringOrNil(v); b != uuid.Nil {
			mappingBranches = append(mappingBranches, model.TableMappingStaffBranch{
				ID:       uuid.NewV4(),
				StaffID:  staff.ID,
				BranchID: b,
			})
		} else {
			crius.Logger.Error(fmt.Sprintf("UpdateStaff rpc请求门店错误:%v", request.Branches))
			resp.ErrorMessage = "门店ID错误"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	mappingRoles := make([]model.TableMappingStaffRole, 0)
	for _, v := range request.Roles {
		if b := uuid.FromStringOrNil(v); b != uuid.Nil {
			mappingRoles = append(mappingRoles, model.TableMappingStaffRole{
				ID:      uuid.NewV4(),
				StaffID: staff.ID,
				RoleID:  b,
			})
		} else {
			crius.Logger.Error(fmt.Sprintf("UpdateStaff rpc请求角色错误:%v", request.Roles))
			resp.ErrorMessage = "角色ID错误"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	err = model.UpdateStaff(staff, mappingBranches, mappingRoles)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateStaff 更新员工信息数据库操作错误:%v", err))
		resp.ErrorMessage = "更新员工错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	//快照数据组装
	before := pkgs.MakeParams(toSnapshotStaff(staffFullInfo, nil, nil))
	staffFullInfo.Staff.EmployeeCode = request.Staff.EmployeeCode
	staffFullInfo.Staff.Name = request.Staff.Name
	staffFullInfo.Staff.Phone = request.Staff.Phone
	staffFullInfo.Staff.PhoneCode = request.Staff.PhoneCode
	staffFullInfo.Staff.Gender = int8(request.Staff.Gender)
	staffFullInfo.Staff.EntryAt = &entryAt
	after := pkgs.MakeParams(toSnapshotStaff(staffFullInfo, mappingBranches, mappingRoles))
	staffID := pkgs.GetMetadata(ctx).StaffID
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: staff.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &staff.ID,
		Method:            "update",
	}

	return resp, nil
}

// UpdateStaffStatus 更新员工状态
func (s *Server) UpdateStaffStatus(ctx context.Context, request *proto.UpdateStaffStatusRequest) (*proto.UpdateStaffStatusResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateStaffStatus")
	resp := &proto.UpdateStaffStatusResponse{
		ErrorCode: pkgs.Success,
	}
	request.Status = strings.ToLower(request.Status)
	if len(request.Id) == 0 || (request.Status != model.StatusOpened && request.Status != model.StatusClosed) {
		crius.Logger.Error("UpdateStaffStatus rpc请求参数错误")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("UpdateStaffStatus rpc请求参数错误, id:%v", request.Id))
		resp.ErrorMessage = "请求参数id错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	staffFullInfo, err := model.ShowStaffByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "员工不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdateStaffStatus 获取员工信息数据库操作错误:%v", err))
		resp.ErrorMessage = "更新员工状态错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	err = model.UpdateStaffStatus(id, request.Status)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdateStaffStatus 更新员工状态数据库错误:%v", err))
		resp.ErrorMessage = "更新员工状态错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	//组装快照
	before := pkgs.MakeParams(toSnapshotStaff(staffFullInfo, nil, nil))
	staffFullInfo.Staff.Status = crius.Status(request.Status)
	after := pkgs.MakeParams(toSnapshotStaff(staffFullInfo, nil, nil))
	staffID := pkgs.GetMetadata(ctx).StaffID
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: staffFullInfo.Staff.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &staffFullInfo.Staff.ID,
		Method:            "update",
	}

	return resp, nil
}

// GetStaffsByRoleID GetStaffsByRoleID
func (s *Server) GetStaffsByRoleID(ctx context.Context, request *proto.GetStaffsByRoleIDRequest) (*proto.GetStaffsByRoleIDResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetStaffs")
	resp := &proto.GetStaffsByRoleIDResponse{
		ErrorCode: pkgs.Success,
	}

	roleID := uuid.FromStringOrNil(request.RoleId)
	if roleID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("GetStaffsByRoleID 请求参数错误:%v", request.RoleId))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	staffs, err := model.GetStaffsByRoleID(roleID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetStaffsByRoleID 获取员工信息数据库操作错误:%v", err))
		resp.ErrorMessage = "获取员工信息错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	for i := range staffs {
		resp.Data = append(resp.Data, toProtoStaff(&staffs[i]))
	}
	return resp, nil
}

// GetStaffs GetStaffs
func (s *Server) GetStaffs(ctx context.Context, request *proto.GetStaffsRequest) (*proto.GetStaffsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetStaffs")
	resp := &proto.GetStaffsResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID

	id := pkgs.GetMetadata(ctx).StaffID
	limitBranches := make(map[uuid.UUID]bool, 0)
	admin := true
	// 查询用户所属门店
	staffFull, err := model.ShowStaffByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("GetStaffs 查询员工不存在,id:%v", id))
			resp.ErrorMessage = "当前登录员工不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("GetStaffs 查询员工数据库错误:%v", err))
		resp.ErrorMessage = "查询员工信息错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if !staffFull.Staff.Admin {
		for _, v := range staffFull.Branches {
			limitBranches[v.ID] = true
		}
		admin = false
	}

	staff := model.TableStaff{
		Name:   request.Name,
		Gender: int8(request.Gender),
		Phone:  request.Phone,
		Status: crius.Status(request.Status),
	}

	branches := getStaffBranches(admin, request.Branches, limitBranches)
	if !admin && len(branches) == 0 {
		resp.Data = &proto.StaffsData{Total: 0}
		return resp, nil
	}

	staffFulls, total, err := model.GetStaffs(staff, branches, merchantID, request.DateStart, request.DateEnd, request.Offset, request.Limit)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetStaffs 获取员工信息数据库操作错误:%v", err))
		resp.ErrorMessage = "获取员工信息错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	resp.Data = &proto.StaffsData{Total: int32(total)}

	for _, staffFull := range staffFulls {
		branches := make([]*proto.BranchInfo, 0)
		roles := make([]*proto.RoleInfo, 0)
		for i := range staffFull.Branches {
			branches = append(branches, toProtoBranch(&staffFull.Branches[i]))
		}

		for i := range staffFull.Roles {
			roles = append(roles, toProtoRole(&staffFull.Roles[i]))
		}

		resp.Data.Staffs = append(resp.Data.Staffs, &proto.StaffFullInfo{
			Staff:    toProtoStaff(&staffFull.Staff),
			Branches: branches,
			Roles:    roles,
		})

	}
	return resp, nil
}

// DeleteStaff DeleteStaff
func (s *Server) DeleteStaff(ctx context.Context, request *proto.DeleteStaffRequest) (*proto.DeleteStaffResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("DeleteStaff")
	resp := &proto.DeleteStaffResponse{
		ErrorCode: pkgs.Success,
	}
	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("DeleteStaff rpc请求参数错误，员工id:%v", request.Id))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	staffFullInfo, err := model.ShowStaffByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("DeleteStaff 获取员工信息数据库操作错误:%v", err))
			resp.ErrorMessage = "员工不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("DeleteStaff 获取员工信息数据库操作错误:%v", err))
		resp.ErrorMessage = "删除员工错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	err = model.DeleteStaff(id)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("DeleteStaff 删除员工信息数据库错误:%v", err))
		resp.ErrorMessage = "删除员工错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	staffID := pkgs.GetMetadata(ctx).StaffID
	before := pkgs.MakeParams(toSnapshotStaff(staffFullInfo, nil, nil))
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: model.TableStaff{}.TableName(),
		Before:            &before,
		TableID:           &id,
		Method:            "delete",
	}
	return resp, nil
}

// ShowStaff ShowStaff
func (s *Server) ShowStaff(ctx context.Context, request *proto.ShowStaffRequest) (*proto.ShowStaffResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowStaff")
	resp := &proto.ShowStaffResponse{
		ErrorCode: pkgs.Success,
	}

	staffID := uuid.FromStringOrNil(request.Id)
	if staffID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("ShowStaff rpc请求参数错误，员工id:%v", request.Id))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}
	staffFull, err := model.ShowStaffByID(staffID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			crius.Logger.Error("ShowStaff 员工不存在")
			resp.ErrorMessage = "员工不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowStaff 查询数据库错误:%v", err))
		resp.ErrorMessage = "查询员工错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	branches := make([]*proto.BranchInfo, 0)
	for i := range staffFull.Branches {
		branches = append(branches, toProtoBranch(&staffFull.Branches[i]))
	}

	roles := make([]*proto.RoleInfo, 0)
	for i := range staffFull.Roles {
		roles = append(roles, toProtoRole(&staffFull.Roles[i]))
	}

	resp.Data = &proto.StaffFullInfo{
		Staff:    toProtoStaff(&staffFull.Staff),
		Branches: branches,
		Roles:    roles,
	}
	return resp, nil
}

// ShowStaffByPhone ShowStaffByPhone
func (s *Server) ShowStaffByPhone(ctx context.Context, request *proto.ShowStaffByPhoneRequest) (*proto.ShowStaffByPhoneResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowStaff")
	resp := &proto.ShowStaffByPhoneResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if request.Phone == "" ||
		request.PhoneCode == "" ||
		merchantID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("ShowStaffByPhone rpc请求参数错误:%v, %v", request.Phone, request.PhoneCode))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}
	staffFull, err := model.ShowStaffByPhone(request.Phone, request.PhoneCode)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			crius.Logger.Error("ShowStaffByPhone 员工不存在")
			resp.ErrorMessage = "员工不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowStaffByPhone 查询数据库错误:%v", err))
		resp.ErrorMessage = "查询员工错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	// 不是管理员，且商户ID不对，用户不存在
	if !staffFull.Staff.Admin && *staffFull.Staff.MerchantID != merchantID {
		crius.Logger.Error(fmt.Sprintf("ShowStaffByPhone 用户商户ID错误:%v", crius.UUIDToString(staffFull.Staff.MerchantID)))
		resp.ErrorMessage = "用户不存在"
		resp.ErrorCode = pkgs.ErrNotFound
		return resp, nil
	}

	branches := make([]*proto.BranchInfo, 0)
	for i := range staffFull.Branches {
		branches = append(branches, toProtoBranch(&staffFull.Branches[i]))
	}

	roles := make([]*proto.RoleInfo, 0)
	for i := range staffFull.Roles {
		roles = append(roles, toProtoRole(&staffFull.Roles[i]))
	}

	resp.Data = &proto.StaffFullInfo{
		Staff:    toProtoStaff(&staffFull.Staff),
		Branches: branches,
		Roles:    roles,
	}
	return resp, nil
}

// SignIn SignIn
func (s *Server) SignIn(ctx context.Context, request *proto.SignInRequest) (*proto.SignInResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("SignIn")
	resp := &proto.SignInResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if request.Username == "" || len(request.Password) < 6 || len(request.Password) > 20 || merchantID == uuid.Nil {
		crius.Logger.Error("SignIn 账号密码参数错误")
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}
	staffFull, err := model.ShowStaffByUsername(request.Username)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("SignIn 查找用户数据库错误:%v", err))
			resp.ErrorMessage = "查找用户错误"
			resp.ErrorCode = pkgs.ErrInternal
		} else {
			crius.Logger.Error("SignIn 用户不存在")
			resp.ErrorMessage = "用户不存在"
			resp.ErrorCode = pkgs.ErrNotFound
		}
		return resp, nil
	}
	// 不是管理员，且商户ID不对，用户不存在
	if !staffFull.Staff.Admin && *staffFull.Staff.MerchantID != merchantID {
		crius.Logger.Error(fmt.Sprintf("SignIn 用户商户ID错误:%v", crius.UUIDToString(staffFull.Staff.MerchantID)))
		resp.ErrorMessage = "用户不存在"
		resp.ErrorCode = pkgs.ErrNotFound
		return resp, nil
	}
	if staffFull.Staff.Password != util.GetMD5(fmt.Sprintf("%s:%s", request.Password, staffFull.Staff.Salt)) {
		crius.Logger.Error("SignIn 校验用户密码失败")
		resp.ErrorMessage = "密码错误"
		resp.ErrorCode = pkgs.ErrPassword
		return resp, nil
	}

	branches := make([]*proto.BranchInfo, 0)
	for i := range staffFull.Branches {
		branches = append(branches, toProtoBranch(&staffFull.Branches[i]))
	}

	roles := make([]*proto.RoleInfo, 0)
	for i := range staffFull.Roles {
		roles = append(roles, toProtoRole(&staffFull.Roles[i]))
	}

	resp.Data = &proto.SignInData{StaffFull: &proto.StaffFullInfo{
		Staff:    toProtoStaff(&staffFull.Staff),
		Branches: branches,
		Roles:    roles,
	}}
	if request.Password == util.DefaultPassword {
		resp.Data.DefaultPassword = true
	}

	return resp, nil
}

// UpdatePassword 更新密码
func (s *Server) UpdatePassword(ctx context.Context, request *proto.UpdatePasswordRequest) (*proto.UpdatePasswordResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdatePassword")
	resp := &proto.UpdatePasswordResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil || len(request.NewPassword) < 6 || len(request.NewPassword) > 20 {
		crius.Logger.Error("UpdatePassword rpc请求参数错误")
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}
	for _, v := range []byte(request.NewPassword) {
		// 密码校验，数字，大写字母，小写字母ASCII码
		if !((v >= 48 && v <= 57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122)) {
			crius.Logger.Error("UpdatePassword rpc请求参数错误，密码包含非数字，字母的字符")
			resp.ErrorMessage = "密码格式错误"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			return resp, nil
		}
	}

	staffFull, err := model.ShowStaffByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "用户不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("UpdatePassword 查询用户数据库错误:%v", err))
		resp.ErrorMessage = "查询用户错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	// 校验旧密码
	if staffFull.Staff.Password != util.GetMD5(fmt.Sprintf("%s:%s", request.OldPassword, staffFull.Staff.Salt)) {
		crius.Logger.Error("UpdatePassword 校验用户旧密码失败")
		resp.ErrorMessage = "原密码错误"
		resp.ErrorCode = pkgs.ErrPassword
		return resp, nil
	}

	err = model.UpdateStaffPassword(id, util.GetMD5(fmt.Sprintf("%s:%s", request.NewPassword, staffFull.Staff.Salt)))
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("UpdatePassword 修改密码数据库错误:%v", err))
		resp.ErrorMessage = "修改密码错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	before := pkgs.MakeParams(toSnapshotStaff(staffFull, nil, nil))
	staffFull.Staff.Password = util.GetMD5(fmt.Sprintf("%s:%s", request.NewPassword, staffFull.Staff.Salt))
	after := pkgs.MakeParams(toSnapshotStaff(staffFull, nil, nil))
	staffID := pkgs.GetMetadata(ctx).StaffID
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: model.TableStaff{}.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &id,
		Method:            "update",
	}

	return resp, nil
}

// ResetPassword 重置密码
func (s *Server) ResetPassword(ctx context.Context, request *proto.ResetPasswordRequest) (*proto.ResetPasswordResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ResetPassword")
	resp := &proto.ResetPasswordResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil {
		crius.Logger.Error("ResetPassword rpc请求参数错误")
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	staffFull, err := model.ShowStaffByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "用户不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ResetPassword 查询用户数据库错误:%v", err))
		resp.ErrorMessage = "查询用户错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	err = model.UpdateStaffPassword(id, util.GetMD5(fmt.Sprintf("%s:%s", util.DefaultPassword, staffFull.Staff.Salt)))
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("ResetPassword 修改密码数据库错误:%v", err))
		resp.ErrorMessage = "重置密码错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	before := pkgs.MakeParams(toSnapshotStaff(staffFull, nil, nil))
	staffFull.Staff.Password = util.GetMD5(fmt.Sprintf("%s:%s", util.DefaultPassword, staffFull.Staff.Salt))
	after := pkgs.MakeParams(toSnapshotStaff(staffFull, nil, nil))
	staffID := pkgs.GetMetadata(ctx).StaffID
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: model.TableStaff{}.TableName(),
		Before:            &before,
		After:             &after,
		TableID:           &id,
		Method:            "update",
	}

	return resp, nil
}

func toProtoStaff(staff *model.TableStaff) *proto.StaffInfo {
	var entryAt, quitAt, createdAt int32
	if staff.EntryAt != nil {
		entryAt = int32(staff.EntryAt.Unix())
	}
	if staff.QuitAt != nil {
		quitAt = int32(staff.QuitAt.Unix())
	}
	if staff.CreatedAt != nil {
		createdAt = int32(staff.CreatedAt.Unix())
	}
	return &proto.StaffInfo{
		Name:         staff.Name,
		Phone:        staff.Phone,
		PhoneCode:    staff.PhoneCode,
		Gender:       int32(staff.Gender),
		Status:       staff.Status.String(),
		Code:         staff.Code,
		EntryAt:      entryAt,
		QuitAt:       quitAt,
		CreatedAt:    createdAt,
		Id:           staff.ID.String(),
		EmployeeCode: staff.EmployeeCode,
		Admin:        staff.Admin,
	}
}

func toSnapshotStaff(staffFull *model.StaffFullInfo, branches []model.TableMappingStaffBranch, roles []model.TableMappingStaffRole) map[string]interface{} {
	m := make(map[string]interface{})
	m["staff"] = staffFull.Staff

	if len(branches) == 0 {
		branches := make([]model.TableMappingStaffBranch, 0)
		for _, v := range staffFull.Branches {
			branches = append(branches, model.TableMappingStaffBranch{
				StaffID:  staffFull.Staff.ID,
				BranchID: v.ID,
			})
		}
		m["branches"] = branches
	} else {
		m["branches"] = branches
	}

	if len(roles) == 0 {
		roles := make([]model.TableMappingStaffRole, 0)
		for _, v := range staffFull.Roles {
			roles = append(roles, model.TableMappingStaffRole{
				StaffID: staffFull.Staff.ID,
				RoleID:  v.ID,
			})
		}
		m["roles"] = roles
	} else {
		m["roles"] = roles
	}
	return m
}

func getStaffBranches(admin bool, reqBranches []string, limitBranches map[uuid.UUID]bool) []uuid.UUID {
	branches := make([]uuid.UUID, 0)
	if admin { //管理员
		for _, v := range reqBranches {
			branchID := uuid.FromStringOrNil(v)
			if branchID != uuid.Nil {
				branches = append(branches, branchID)
			}
		}
	} else {
		// 非管理员
		if len(reqBranches) != 0 { // 请求门店筛选不为空，结合员工所属门店组合筛选条件
			for _, v := range reqBranches {
				branchID := uuid.FromStringOrNil(v)
				if branchID != uuid.Nil && limitBranches[branchID] {
					branches = append(branches, branchID)
				}
			}
		} else { // 请求门店筛选为空，取员工所属门店
			for branchID := range limitBranches {
				branches = append(branches, branchID)
			}
		}
	}
	return branches
}
