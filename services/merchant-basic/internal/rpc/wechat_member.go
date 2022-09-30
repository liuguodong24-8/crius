package rpc

import (
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

//ShowWechatUser 获取微信用户信息
func (s *Server) ShowWechatUser(ctx context.Context, request *proto.ShowWechatUserRequest) (*proto.ShowWechatUserResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowWechatUser")
	resp := &proto.ShowWechatUserResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.Id)
	member, err := model.ShowWechatUser(id)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("ShowWechatUser 数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	resp.Data = toProtoWechatUser(member)
	return resp, nil
}

//ShowWechatUserByMember 获取微信用户信息
func (s *Server) ShowWechatUserByMember(ctx context.Context, request *proto.ShowWechatUserByMemberRequest) (*proto.ShowWechatUserByMemberResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowWechatUserByMember")
	resp := &proto.ShowWechatUserByMemberResponse{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(request.MemberId)
	member, err := model.ShowWechatUserByMember(id)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("ShowWechatUserByMember 数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	resp.Data = toProtoWechatUser(member)
	return resp, nil
}

// GetWechatUsers 批量查询微信用户
func (s *Server) GetWechatUsers(ctx context.Context, request *proto.GetWechatUsersRequest) (*proto.GetWechatUsersResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetWechatUsers")
	resp := &proto.GetWechatUsersResponse{
		ErrorCode: pkgs.Success,
	}

	if len(request.Ids) == 0 {
		crius.Logger.Error(fmt.Sprintf("GetWechatUsers 参数错误:%v", request))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}
	var ids []uuid.UUID
	for _, id := range request.Ids {
		ids = append(ids, uuid.FromStringOrNil(id))
	}

	members, err := model.GetWechatUsers(ids)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetWechatUsers 数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	var protoWechatUsers []*proto.WechatUserInfo
	for _, member := range members {
		protoWechatUsers = append(protoWechatUsers, &proto.WechatUserInfo{
			Id:         member.ID.String(),
			Nickname:   member.Nickname,
			Headimgurl: member.HeadImgURL,
			MemberId:   crius.UUIDToString(member.MemberID),
		})
	}
	resp.Data = protoWechatUsers
	return resp, nil
}

//CreateOrUpdateWechatUser 新建或更新
func (s *Server) CreateOrUpdateWechatUser(ctx context.Context, request *proto.CreateOrUpdateWechatUserRequest) (*proto.CreateOrUpdateWechatUserResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateOrUpdateWechatUser")
	resp := &proto.CreateOrUpdateWechatUserResponse{
		ErrorCode: pkgs.Success,
	}

	if request.User.Unionid == "" {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}
	//unionid, appid 唯一确定一个平台的微信用户
	user, err := model.ShowWechatUserByUnionID(request.User.Unionid, request.User.Appid)
	if err != nil && err != gorm.ErrRecordNotFound {
		crius.Logger.Error(fmt.Sprintf("ShowWechatUserByUnionID 数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	createOrUpdateMember := &model.TableWechatMember{
		Nickname:   request.User.Nickname,
		Sex:        int8(request.User.Sex),
		Province:   request.User.Province,
		City:       request.User.City,
		HeadImgURL: request.User.Headimgurl,
	}
	if memberID := uuid.FromStringOrNil(request.User.MemberId); memberID != uuid.Nil {
		//传了用户ID
		if user.MemberID != nil && !uuid.Equal(*user.MemberID, memberID) {
			//已绑定过
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "该手机号已绑定其他微信账号"
			return resp, nil
		}
		//授权更新的时候没有memberID，绑定会员的时候有
		createOrUpdateMember.MemberID = &memberID
		//查询member是否有wechat渠道
		member, err := model.ShowMember(memberID)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("ShowMember 数据库错误:%v", err))
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.ErrorCode = pkgs.ErrUnprocessableEntity
				resp.ErrorMessage = "用户不存在"
				return resp, nil
			}
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
		if !isOldChannel(model.MemberChannelWechat, *member.Channels) {
			//更新member的渠道
			var channels model.MemberChannelArr
			channels = append(*member.Channels, model.MemberChannelWechat)
			updateMember := model.TableMember{
				ID:       memberID,
				Channels: &channels,
			}
			_ = model.UpdateMember(&updateMember)
			//不处理，不影响主流程
		}
	}

	//一个memberID只能绑定一个(appid, unionid)
	if user.ID != uuid.Nil {
		//存在，更新
		if err = model.UpdateWechatUser(user.AppID, user.UnionID, createOrUpdateMember); err != nil {
			crius.Logger.Error(fmt.Sprintf("UpdateWechatUser 数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
	} else {
		createOrUpdateMember.ID = uuid.NewV4()
		createOrUpdateMember.AppID = request.User.Appid
		createOrUpdateMember.OpenID = request.User.Openid
		createOrUpdateMember.UnionID = request.User.Unionid
		if err = model.CreateWechatUser(createOrUpdateMember); err != nil {
			crius.Logger.Error(fmt.Sprintf("CreateWechatUser 数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
	}
	//重新查一遍
	user, err = model.ShowWechatUserByUnionID(request.User.Unionid, request.User.Appid)
	resp.Data = toProtoWechatUser(user)
	return resp, nil
}

func toProtoWechatUser(wechatUser *model.TableWechatMember) *proto.WechatUser {
	var memberID string
	if wechatUser.MemberID != nil {
		memberID = wechatUser.MemberID.String()
	}
	var birthday string
	if wechatUser.Member.Birthday != nil {
		birthday = wechatUser.Member.Birthday.Format(BirthdayFormat)
	}
	return &proto.WechatUser{
		Id:              wechatUser.ID.String(),
		Openid:          wechatUser.OpenID,
		Appid:           wechatUser.AppID,
		MemberId:        memberID,
		Nickname:        wechatUser.Nickname,
		Sex:             int32(wechatUser.Sex),
		Province:        wechatUser.Province,
		City:            wechatUser.City,
		Headimgurl:      wechatUser.HeadImgURL,
		Unionid:         wechatUser.UnionID,
		CreatedAt:       wechatUser.CreatedAt.String(),
		MemberPhone:     wechatUser.Member.Phone,
		MemberPhoneCode: wechatUser.Member.PhoneCode,
		MemberName:      wechatUser.Member.Name,
		MemberAvatar:    wechatUser.Member.Avatar,
		MemberGender:    int32(wechatUser.Member.Gender),
		MemberBirthday:  birthday,
	}
}
