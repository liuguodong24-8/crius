package rpc

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	centreData "gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model/centre_data"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// GetRoomTypeGroups 获取房型组
func (s *Server) GetRoomTypeGroups(ctx context.Context, request *proto.Empty) (*proto.GetRoomTypeGroupsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetRoomTypeGroups")
	resp := &proto.GetRoomTypeGroupsResponse{
		ErrorCode: pkgs.Success,
	}

	roomTypeGroups, err := centreData.GetRoomTypeGroups()
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetRoomTypeGroups 查询房型组数据库错误:%v", err))
		resp.ErrorMessage = "数据库"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	var roomTypeGroupsProto []*proto.RoomTypeGroup
	if len(roomTypeGroups) > 0 {
		for _, group := range roomTypeGroups {
			roomTypeGroupsProto = append(roomTypeGroupsProto, toProtoRoomTypeGroup(group))
		}
	}

	resp.Data = roomTypeGroupsProto
	return resp, nil
}

// GetRoomTypeGroupsByRoomTypeID 根据room_type_id获取房型组
func (s *Server) GetRoomTypeGroupsByRoomTypeID(ctx context.Context, request *proto.GetRoomTypeGroupsByRoomTypeIDRequest) (*proto.GetRoomTypeGroupsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetRoomTypeGroupsByRoomTypeID")
	resp := &proto.GetRoomTypeGroupsResponse{
		ErrorCode: pkgs.Success,
	}

	roomTypeID := uuid.FromStringOrNil(request.RoomTypeId)
	roomTypeGroups, err := centreData.GetRoomTypeGroupsByRoomTypeID(roomTypeID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetRoomTypeGroupsByRoomTypeID 查询房型组数据库错误:%v", err))
		resp.ErrorMessage = "数据库"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	var roomTypeGroupsProto []*proto.RoomTypeGroup
	if len(roomTypeGroups) > 0 {
		for _, group := range roomTypeGroups {
			roomTypeGroupsProto = append(roomTypeGroupsProto, toProtoRoomTypeGroup(group))
		}
	}

	resp.Data = roomTypeGroupsProto
	return resp, nil
}

func toProtoRoomTypeGroup(v centreData.TableMetaRoomTypeGroup) *proto.RoomTypeGroup {
	return &proto.RoomTypeGroup{
		Id:   v.RoomTypeGroupID.String(),
		Name: v.RoomTypeGroupName,
	}
}
