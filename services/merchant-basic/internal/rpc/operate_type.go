package rpc

import (
	"context"
	"fmt"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	centreData "gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model/centre_data"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// GetOperateTypes list
func (s *Server) GetOperateTypes(ctx context.Context, request *proto.Empty) (*proto.GetOperateTypesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetOperateTypes")
	resp := &proto.GetOperateTypesResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	data, err := centreData.GetOperateTypes([]func(db *gorm.DB) *gorm.DB{crius.ColumnEqualScope("merchant_id", merchantID)})
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetOperateTypes 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}

	var protoTypes []*proto.OperateType
	if len(data) > 0 {
		for _, v := range data {
			protoTypes = append(protoTypes, toProtoOperateType(v))
		}
	}

	resp.Data = protoTypes

	return resp, nil
}

func toProtoOperateType(v centreData.TableProductOperateType) *proto.OperateType {
	return &proto.OperateType{
		Id:       v.OperateTypeID.String(),
		Name:     v.TypeName,
		ParentId: crius.UUIDToString(v.ParentID),
	}
}
