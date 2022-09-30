package rpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/message"

	"gitlab.omytech.com.cn/micro-service/merchant-basic/util"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// BirthdayFormat 日期格式
const BirthdayFormat = "2006-01-02"

//CreateMember 创建会员 只开放给商户平台的新建，可能的情况（1.手机号不存在 2.手机号存在且是来电同步过来的数据）
func (s *Server) CreateMember(ctx context.Context, request *proto.CreateMemberRequest) (*proto.CreateMemberResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateMember")
	resp := &proto.CreateMemberResponse{
		ErrorCode: pkgs.Success,
	}

	// 判断参数合法性
	if request.Member == nil {
		crius.Logger.Error("CreateMember rpc请求参数nil")
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "rpc请求参数nil"
		return resp, nil
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	firstBranchID := uuid.FromStringOrNil(request.Member.FirstBranchId)
	staffID := pkgs.GetMetadata(ctx).StaffID

	// 校验手机号唯一性
	member, err := model.ShowMemberByAccuratePhone(request.Member.PhoneCode, request.Member.Phone, merchantID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		//查询数据报错且不是没有找到数据错误
		crius.Logger.Error(fmt.Sprintf("CreateMember 校验用户输入手机号数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	branch, err := model.ShowBranchByID(firstBranchID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranch 数据库查询错误:%v, %v", request.Member, err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库查询错误"
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "对应门店不存在"
		}
		return resp, nil
	}

	createMember := model.TableMember{
		Name:          request.Member.Name,
		Phone:         request.Member.Phone,
		PhoneCode:     request.Member.PhoneCode,
		Gender:        int8(request.Member.Gender),
		Avatar:        request.Member.Avatar,
		City:          branch.CityID,
		FirstBranchID: &firstBranchID,
		StaffID:       &staffID,
		MerchantID:    &merchantID,
		FirstBrand:    branch.BrandID,
	}
	if request.Member.Birthday != "" {
		birthday, _ := time.ParseInLocation(BirthdayFormat, request.Member.Birthday, time.Local)
		createMember.Birthday = &birthday
	}
	var memberID uuid.UUID

	if member.ID == uuid.Nil {
		//没有记录，新增
		createMember.ID = uuid.NewV4()
		memberID = createMember.ID
		createMember.FirstChannel = model.MemberChannel(request.Member.Channel)
		channels := model.MemberChannelArr{model.MemberChannel(request.Member.Channel)}
		createMember.Channels = &channels
		err = model.CreateMember(createMember)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("CreateMember 创建会员数据库操作错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
		resp.Data = createMember.ID.String()
	} else {
		//有记录,覆盖来电渠道的用户数据
		updateMember := createMember
		if member.FirstBranchID != nil {
			//门店不为空，部分数据覆盖
			updateMember = model.TableMember{
				Name:      createMember.Name,
				PhoneCode: createMember.PhoneCode,
				Phone:     createMember.Phone,
				Avatar:    createMember.Avatar,
				Gender:    createMember.Gender,
				Birthday:  createMember.Birthday,
			}
		}

		channels := member.Channels
		if !isOldChannel(model.MemberChannel(request.Member.Channel), *channels) {
			*channels = append(*channels, model.MemberChannel(request.Member.Channel))
		}
		updateMember.ID = member.ID
		memberID = member.ID
		updateMember.Channels = channels
		err = model.UpdateMember(&updateMember)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("UpdateMember 数据库操作错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
		resp.Data = updateMember.ID.String()
	}

	tasks := []message.TaskCategory{
		message.BindPhone,
	}
	if request.Member.Name != "" {
		tasks = append(tasks, message.FillName)
	}
	if request.Member.Birthday != "" {
		tasks = append(tasks, message.FillBirthday)
	}
	if request.Member.Gender != 0 {
		tasks = append(tasks, message.FillSex)
	}
	if request.Member.Avatar != "" {
		tasks = append(tasks, message.UploadPhoto)
	}
	util.PublishMemberTasks(tasks, memberID)
	return resp, nil
}

// CreateOrUpdateCallingMember 同步来电用户信息
func (s *Server) CreateOrUpdateCallingMember(ctx context.Context, request *proto.CreateOrUpdateCallingMemberRequest) (*proto.CreateOrUpdateCallingMemberResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateOrUpdateCallingMember")
	resp := &proto.CreateOrUpdateCallingMemberResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	//先查询手机会员是否存在
	member, err := model.ShowMemberByAccuratePhone(request.PhoneCode, request.Phone, merchantID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		crius.Logger.Error(fmt.Sprintf("ShowMemberByAccuratePhone 数据库错误:%v", request))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	var memberID uuid.UUID
	if member.ID == uuid.Nil {
		staffID := pkgs.GetMetadata(ctx).StaffID
		//新用户，插入
		channels := &model.MemberChannelArr{model.MemberChannel(request.Channel)}
		memberTable := model.TableMember{
			ID:           uuid.NewV4(),
			Name:         request.Name,
			Phone:        request.Phone,
			PhoneCode:    request.PhoneCode,
			Gender:       int8(request.Gender),
			FirstChannel: model.MemberChannel(request.Channel),
			Channels:     channels,
			MerchantID:   &merchantID,
			StaffID:      &staffID,
		}
		memberID = memberTable.ID
		firstBranchID := uuid.FromStringOrNil(request.BranchId)
		//门店ID存在时处理
		if firstBranchID != uuid.Nil {
			branch, err := model.ShowBranchByID(firstBranchID)
			if err != nil {
				crius.Logger.Error(fmt.Sprintf("CreateOrUpdateCallingMember 查询门店错误:%v", err))
				resp.ErrorCode = pkgs.ErrInternal
				resp.ErrorMessage = "数据库错误"
				return resp, nil
			}
			memberTable.City = branch.CityID
			memberTable.FirstBranchID = &firstBranchID
			memberTable.FirstBrand = branch.BrandID
		}

		err = model.CreateMember(memberTable)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("CreateMember 数据库错误:%v", request))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
	} else {
		memberID = member.ID
		//判断来源
		channels := member.Channels
		if isMemberFromNoCallChannel(*channels) {
			//退出 不更新信息
			crius.Logger.Info(fmt.Sprintf("因会员渠道问题不更新信息:%v", channels))
			return resp, nil
		}
		//可覆盖 区分来电列表编辑还是新增预约
		if request.CanOverwrite {
			channels := member.Channels
			if !isOldChannel(model.MemberChannel(request.Channel), *channels) {
				*channels = append(*channels, model.MemberChannel(request.Channel))
			}
			memberTable := model.TableMember{
				ID:        member.ID,
				Name:      request.Name,
				Phone:     request.Phone,
				PhoneCode: request.PhoneCode,
				Gender:    int8(request.Gender),
				Channels:  channels,
			}
			firstBranchID := uuid.FromStringOrNil(request.BranchId)
			//现有数据门店不存在且新数据门店ID存在时处理
			if member.FirstBranchID == nil && firstBranchID != uuid.Nil {
				branch, err := model.ShowBranchByID(firstBranchID)
				if err != nil {
					crius.Logger.Error(fmt.Sprintf("CreateOrUpdateCallingMember 查询门店错误:%v", err))
					resp.ErrorCode = pkgs.ErrInternal
					resp.ErrorMessage = "数据库错误"
					return resp, nil
				}
				memberTable.City = branch.CityID
				memberTable.FirstBranchID = &firstBranchID
				memberTable.FirstBrand = branch.BrandID
			}
			err := model.UpdateMember(&memberTable)
			if err != nil {
				crius.Logger.Error(fmt.Sprintf("UpdateMember 数据库错误:%v", request))
				resp.ErrorCode = pkgs.ErrInternal
				resp.ErrorMessage = "数据库错误"
				return resp, nil
			}
		}
	}
	tasks := []message.TaskCategory{
		message.BindPhone,
	}
	if request.Name != "" {
		tasks = append(tasks, message.FillName)
	}
	if request.Gender != 0 {
		tasks = append(tasks, message.FillSex)
	}
	util.PublishMemberTasks(tasks, memberID)
	resp.Data = memberID.String()
	return resp, nil
}

//CreateWechatMember 只开发给微信手机号注册
func (s *Server) CreateWechatMember(ctx context.Context, request *proto.CreateWechatMemberRequest) (*proto.CreateWechatMemberResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateWechatMember")
	resp := &proto.CreateWechatMemberResponse{
		ErrorCode: pkgs.Success,
	}

	// 判断参数合法性
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if request.Phone == "" ||
		request.PhoneCode == "" ||
		merchantID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("CreateWechatMember rpc请求参数错误:%v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "rpc请求参数错误"
		return resp, nil
	}

	// 业务上有前置判断，此处不做手机号唯一校验
	createMember := model.TableMember{
		ID:           uuid.NewV4(),
		Name:         request.Name,
		Phone:        request.Phone,
		PhoneCode:    request.PhoneCode,
		Gender:       int8(request.Gender),
		Avatar:       request.Avatar,
		MerchantID:   &merchantID,
		FirstChannel: model.MemberChannelWechat,
		Channels:     &model.MemberChannelArr{model.MemberChannelWechat},
	}
	if request.Birthday != "" {
		birthday, _ := time.ParseInLocation(BirthdayFormat, request.Birthday, time.Local)
		createMember.Birthday = &birthday
	}
	err := model.CreateMember(createMember)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateMember 创建会员数据库操作错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	//任务
	tasks := []message.TaskCategory{
		message.BindPhone,
	}
	if request.Name != "" {
		tasks = append(tasks, message.FillName)
	}
	if request.Birthday != "" {
		tasks = append(tasks, message.FillBirthday)
	}
	if request.Gender != 0 {
		tasks = append(tasks, message.FillSex)
	}
	if request.Avatar != "" {
		tasks = append(tasks, message.UploadPhoto)
	}
	util.PublishMemberTasks(tasks, createMember.ID)

	resp.Data = createMember.ID.String()
	return resp, nil
}

//GetMembers 会员列表
func (s *Server) GetMembers(ctx context.Context, request *proto.GetMembersRequest) (*proto.GetMembersResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetMembers")
	resp := &proto.GetMembersResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if merchantID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("GetMembers merchant_id missing"))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	//门店范围
	branchIDs := pkgs.GetMetadata(ctx).BranchIDs
	firstBrand := uuid.FromStringOrNil(request.FirstBrand)

	scopes := []func(db *gorm.DB) *gorm.DB{
		crius.ColumnEqualScopeDefault("phone", request.Phone),
		crius.ColumnEqualScopeDefault("name", request.Name),
		crius.ColumnEqualScopeDefault("gender", int8(request.Gender)),
		crius.ColumnEqualScopeDefault("first_channel", request.Channel),
		crius.ColumnInScopeDefault("first_branch_id", branchIDs.Slice()),
		crius.ColumnEqualScope("merchant_id", merchantID),
		crius.ColumnEqualScopeDefault("first_brand", firstBrand),
	}
	if request.InChannel != "" {
		scopes = append(scopes, crius.ArrayAnyScope("channels", model.MemberChannel(request.InChannel)))
	}
	count, err := model.CountMembers(scopes)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("model CountMembers 查询数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "查询数据库错误"
		return resp, nil
	}
	var protoMembers []*proto.MemberInfo
	if count > 0 {
		scopes = append(scopes, crius.PaginationScope(request.Offset, request.Limit))
		members, err := model.GetMembers(scopes)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("model GetMembers 查询数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "查询数据库错误"
			return resp, nil
		}

		if len(members) > 0 {
			for _, member := range members {
				protoMembers = append(protoMembers, toMemberWithBehaviorCountProto(&member))
			}
		}
	}

	resp.Data = &proto.MembersData{Total: int32(count), Members: protoMembers}
	return resp, nil
}

// GetMembersByIDs 批量通过id查询会员信息
func (s *Server) GetMembersByIDs(ctx context.Context, request *proto.GetMembersByIDsRequest) (*proto.GetMembersByIDsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetMembersByIDs")
	resp := &proto.GetMembersByIDsResponse{
		ErrorCode: pkgs.Success,
	}
	var ids []uuid.UUID
	for _, id := range request.Ids {
		idTmp := uuid.FromStringOrNil(id)
		if idTmp != uuid.Nil {
			ids = append(ids, idTmp)
		}
	}
	if len(request.Ids) == 0 {
		crius.Logger.Error(fmt.Sprintf("GetMembersByIDs"))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	members, err := model.GetMembersByIDs(ids)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetMembersByIDs 查询数据库错误 %+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "查询数据库错误"
		return resp, nil
	}
	var protoMembers []*proto.MemberInfo
	if len(members) > 0 {
		for _, member := range members {
			protoMembers = append(protoMembers, toMemberProto(&member))
		}
	}

	resp.Data = protoMembers
	return resp, nil
}

//ShowMember 根据id查询会员信息
func (s *Server) ShowMember(ctx context.Context, request *proto.ShowMemberRequest) (*proto.ShowMemberResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowMember")
	resp := &proto.ShowMemberResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	if id == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("ShowMember 参数错误:%+v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}

	member, err := model.ShowMember(id)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("model.ShowMember 数据库查询错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库查询错误"
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "未找到对应用户"
		}
		return resp, nil
	}

	resp.Data = toMemberProto(member)
	return resp, nil
}

// ShowMemberByAccuratePhone 手机号精确查询用户
func (s *Server) ShowMemberByAccuratePhone(ctx context.Context, request *proto.ShowMemberByAccuratePhoneRequest) (*proto.ShowMemberByAccuratePhoneResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowMemberByAccuratePhone")
	resp := &proto.ShowMemberByAccuratePhoneResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if request.Phone == "" || merchantID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("ShowMemberByAccuratePhone 参数错误:%+v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}

	member, err := model.ShowMemberByAccuratePhone(request.PhoneCode, request.Phone, merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("model.ShowMemberByAccuratePhone 数据库查询错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库查询错误"
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "未找到对应用户"
		}
		return resp, nil
	}

	resp.Data = toMemberProto(member)
	return resp, nil

}

// UpdateMember 更新会员信息
func (s *Server) UpdateMember(ctx context.Context, request *proto.UpdateMemberRequest) (*proto.UpdateMemberResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateMember")
	resp := &proto.UpdateMemberResponse{
		ErrorCode: pkgs.Success,
	}
	memberID := uuid.FromStringOrNil(request.MemberId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if memberID == uuid.Nil ||
		(request.Name == "" && request.PhoneCode == "" && request.Phone == "" &&
			request.Gender == 0 && request.Avatar == "" && request.Birthday == "") {
		crius.Logger.Error(fmt.Sprintf("UpdateMember 参数错误:%+v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}

	_, err := model.ShowMemberByAccuratePhone(request.PhoneCode, request.Phone, merchantID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		crius.Logger.Error(fmt.Sprintf("ShowMemberByAccuratePhone 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	if err == nil {
		//手机号已存在
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "手机号已存在"
		return resp, nil
	}

	updateMember := model.TableMember{
		ID:     memberID,
		Name:   request.Name,
		Phone:  request.Phone,
		Gender: int8(request.Gender),
		Avatar: request.Avatar,
	}
	if request.Birthday != "" {
		birthday, _ := time.ParseInLocation(BirthdayFormat, request.Birthday, time.Local)
		updateMember.Birthday = &birthday
	}

	err = model.UpdateMember(&updateMember)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("model.UpdateMember 数据库处理错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	tasks := []message.TaskCategory{
		message.BindPhone,
	}
	if request.Name != "" {
		tasks = append(tasks, message.FillName)
	}
	if request.Birthday != "" {
		tasks = append(tasks, message.FillBirthday)
	}
	if request.Gender != 0 {
		tasks = append(tasks, message.FillSex)
	}
	if request.Avatar != "" {
		tasks = append(tasks, message.UploadPhoto)
	}
	util.PublishMemberTasks(tasks, memberID)
	return resp, nil
}

//UpdateMemberBranchInfo 更新会员的所属门店信息
func (s *Server) UpdateMemberBranchInfo(ctx context.Context, request *proto.UpdateMemberBranchInfoRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateMemberBranchInfo")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}
	memberID := uuid.FromStringOrNil(request.MemberId)
	branchID := uuid.FromStringOrNil(request.BranchId)

	member, err := model.ShowMember(memberID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("mShowMember 数据库错误:%+v", err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "用户不存在"
			return resp, nil
		}
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	if member.FirstBranchID == nil {
		//更新门店信息
		branch, err := model.ShowBranchByID(branchID)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("ShowBranchByID 数据库错误:%+v", err))
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.ErrorCode = pkgs.ErrUnprocessableEntity
				resp.ErrorMessage = "门店不存在"
				return resp, nil
			}
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
		updateMember := model.TableMember{
			ID:            memberID,
			FirstBranchID: &branchID,
			FirstBrand:    branch.BrandID,
			City:          branch.CityID,
		}

		err = model.UpdateMember(&updateMember)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("UpdateMember 数据库错误:%+v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
	}

	return resp, nil
}

// GetBirthdayMembers 获取生日的会员
func (s *Server) GetBirthdayMembers(ctx context.Context, request *proto.GetBirthdayMembersRequest) (*proto.GetBirthdayMembersResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBirthdayMembers")
	resp := &proto.GetBirthdayMembersResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if merchantID == uuid.Nil || request.Birthday == "" {
		crius.Logger.Error(fmt.Sprintf("GetBirthdayMembers 参数错误:%+v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}
	members, err := model.GetBirthdayMembers(request.Birthday, merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBirthdayMembers 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	var membersProto []*proto.MemberInfo
	if len(members) > 0 {
		for _, m := range members {
			membersProto = append(membersProto, toMemberProto(&m))
		}
	}
	resp.Data = membersProto
	return resp, nil
}

// GetMembersByPhoneSuffix 根据手机尾号查询会员
func (s *Server) GetMembersByPhoneSuffix(ctx context.Context, request *proto.GetMembersByPhoneSuffixRequest) (*proto.GetMembersByPhoneSuffixResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetMembersByPhoneSuffix")
	resp := &proto.GetMembersByPhoneSuffixResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if merchantID == uuid.Nil || request.PhoneSuffix == "" || len(request.PhoneSuffix) != 4 {
		crius.Logger.Error(fmt.Sprintf("GetMembersByPhoneSuffix 参数错误:%+v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}
	members, err := model.GetMembersByPhoneSuffix(request.PhoneSuffix, merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetMembersByPhoneSuffix 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	var membersProto []*proto.MemberInfo
	if len(members) > 0 {
		for _, m := range members {
			membersProto = append(membersProto, toMemberProto(&m))
		}
	}
	resp.Data = membersProto
	return resp, nil
}

// CountCouponMemberIDs count
func (s *Server) CountCouponMemberIDs(ctx context.Context, request *proto.GetCouponMemberIDsRequest) (*proto.CountCouponMemberIDsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CountCouponMemberIDs")
	resp := &proto.CountCouponMemberIDsResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	scopes := []func(db *gorm.DB) *gorm.DB{
		crius.ColumnEqualScope("merchant_id", merchantID),
		crius.ColumnInScopeDefault("city", request.CityCodes),
		crius.ColumnInScopeDefault("first_branch_id", request.BranchIds),
		crius.ColumnEqualScopeDefault("gender", request.Gender),
		crius.ColumnInScopeDefault("phone", request.Phones),
		crius.ColumnSymbolScope("created_at", ">", time.Unix(0, request.CreateAt*1e6)),
	}

	count, err := model.CountCouponMemberIDs(scopes)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CountCouponMemberIDs 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	resp.Data = count
	return resp, nil
}

// GetCouponMemberIDs 推送券对象
func (s *Server) GetCouponMemberIDs(ctx context.Context, request *proto.GetCouponMemberIDsRequest) (*proto.GetCouponMemberIDsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetCouponMemberIDs")
	resp := &proto.GetCouponMemberIDsResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	scopes := []func(db *gorm.DB) *gorm.DB{
		crius.ColumnEqualScope("merchant_id", merchantID),
		crius.ColumnInScopeDefault("city", request.CityCodes),
		crius.ColumnInScopeDefault("first_branch_id", request.BranchIds),
		crius.ColumnEqualScopeDefault("gender", request.Gender),
		crius.ColumnInScopeDefault("phone", request.Phones),
		crius.ColumnSymbolScope("created_at", ">", time.Unix(0, request.CreateAt*1e6)),
		crius.PaginationScope(0, request.Limit),
	}

	members, err := model.GetCouponMemberIDs(scopes)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetCouponMemberIDs 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	var protoMembers []*proto.CouponMember
	for _, m := range members {
		pm := &proto.CouponMember{
			Id:        m.ID.String(),
			CreatedAt: m.CreatedAt.UnixNano() / 1e6,
		}
		protoMembers = append(protoMembers, pm)
	}

	resp.Data = protoMembers
	return resp, nil
}

// SearchMember 会员查找
func (s *Server) SearchMember(ctx context.Context, request *proto.SearchMemberRequest) (*proto.SearchMemberResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("SearchMember")
	resp := &proto.SearchMemberResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	scopes := []func(*gorm.DB) *gorm.DB{
		crius.ColumnEqualScope("merchant_id", merchantID),
		crius.ColumnLikeScope("name", request.Name),
		crius.ColumnLikeScope("phone", request.Phone),
	}

	members, err := model.SearchMember(scopes)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("SearchMember 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	var protoMembers []*proto.MemberInfo
	if len(members) > 0 {
		for _, member := range members {
			protoMembers = append(protoMembers, toMemberProto(&member))
		}
	}

	resp.Data = protoMembers
	return resp, nil
}

func toMemberProto(member *model.TableMember) *proto.MemberInfo {
	var birthday string
	if member.Birthday != nil {
		birthday = member.Birthday.Format(BirthdayFormat)
	}

	return &proto.MemberInfo{
		Id:            member.ID.String(),
		Name:          member.Name,
		Phone:         member.Phone,
		PhoneCode:     member.PhoneCode,
		Gender:        int32(member.Gender),
		Avatar:        member.Avatar,
		CityCode:      member.City,
		Code:          member.Code,
		FirstBranchId: crius.UUIDToString(member.FirstBranchID),
		StaffId:       crius.UUIDToString(member.StaffID),
		CreatedAt:     int32(member.CreatedAt.Unix()),
		Channel:       string(member.FirstChannel),
		Birthday:      birthday,
		FirstBrand:    crius.UUIDToString(member.FirstBrand),
	}
}

func toMemberWithBehaviorCountProto(member *model.MemberWithBehavior) *proto.MemberInfo {
	var birthday string
	if member.Birthday != nil {
		birthday = member.Birthday.Format(BirthdayFormat)
	}

	return &proto.MemberInfo{
		Id:            member.ID.String(),
		Name:          member.Name,
		Phone:         member.Phone,
		PhoneCode:     member.PhoneCode,
		Gender:        int32(member.Gender),
		Avatar:        member.Avatar,
		CityCode:      member.City,
		Code:          member.Code,
		FirstBranchId: crius.UUIDToString(member.FirstBranchID),
		StaffId:       crius.UUIDToString(member.StaffID),
		CreatedAt:     int32(member.CreatedAt.Unix()),
		Channel:       string(member.FirstChannel),
		Birthday:      birthday,
		BehaviorCount: uint32(member.BehaviorCount),
		FirstBrand:    crius.UUIDToString(member.FirstBrand),
	}
}

// isMemberFromNoCallChannel 会员是否包含非来电的渠道
func isMemberFromNoCallChannel(channels model.MemberChannelArr) bool {
	for _, ch := range channels {
		if ch == model.MemberChannelOpenCard || ch == model.MemberChannelWechat {
			return true
		}
	}
	return false
}

// 是否是已存在的渠道
func isOldChannel(channel model.MemberChannel, channels model.MemberChannelArr) bool {
	for _, ch := range channels {
		if ch == channel {
			return true
		}
	}
	return false
}
