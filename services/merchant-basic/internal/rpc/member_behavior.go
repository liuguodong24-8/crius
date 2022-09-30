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

//CreateMemberBehavior 记录用户行为
func (s *Server) CreateMemberBehavior(ctx context.Context, request *proto.CreateMemberBehaviorRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CreateMemberBehavior")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	// 判断参数合法性
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	branchID := uuid.FromStringOrNil(request.BranchId)
	staffID := pkgs.GetMetadata(ctx).StaffID
	memberID := uuid.FromStringOrNil(request.MemberId)
	if request.Behavior == "" ||
		memberID == uuid.Nil ||
		merchantID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("CreateMemberBehavior rpc请求参数错误:%v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "rpc请求参数错误"
		return resp, nil
	}

	tableMemberBehavior := model.TableMemberBehavior{
		ID:         uuid.NewV4(),
		MemberID:   memberID,
		Behavior:   request.Behavior,
		StaffID:    &staffID,
		BranchID:   &branchID,
		MerchantID: merchantID,
	}

	err := model.CreateMemberBehavior(&tableMemberBehavior)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateMemberBehavior 数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	return resp, nil
}

//GetMemberBehaviors 查询用户行为
func (s *Server) GetMemberBehaviors(ctx context.Context, request *proto.GetMemberBehaviorsRequest) (*proto.GetMemberBehaviorsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetMemberBehaviors")
	resp := &proto.GetMemberBehaviorsResponse{
		ErrorCode: pkgs.Success,
	}

	memberID := uuid.FromStringOrNil(request.MemberId)
	if memberID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("GetMemberBehaviors rpc请求参数错误:%v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "rpc请求参数错误"
		return resp, nil
	}

	count, err := model.GetMemberBehaviorsCount(memberID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetMemberBehaviorsCount 数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	var protoBehaviors []*proto.MemberBehavior
	if count > 0 {
		behaviors, err := model.GetMemberBehaviors(memberID, request.Offset, request.Limit)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("GetMemberBehaviors 数据库错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
		if len(behaviors) > 0 {
			for _, behavior := range behaviors {
				protoBehaviors = append(protoBehaviors, toProtoMemberBehaviors(&behavior))
			}
		}
	}
	resp.Data = &proto.MemberBehaviors{
		MemberBehaviors: protoBehaviors,
		Total:           int32(count),
	}
	return resp, nil
}

func toProtoMemberBehaviors(memberBehavior *model.TableMemberBehavior) *proto.MemberBehavior {
	return &proto.MemberBehavior{
		Id:        memberBehavior.ID.String(),
		MemberId:  memberBehavior.MemberID.String(),
		Behavior:  memberBehavior.Behavior,
		StaffId:   crius.UUIDToString(memberBehavior.StaffID),
		BranchId:  crius.UUIDToString(memberBehavior.BranchID),
		CreatedAt: int32(memberBehavior.CreatedAt.Unix()),
	}
}
