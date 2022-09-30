package rpc

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

//ShowGrowthConfig 查询
func (s *Server) ShowGrowthConfig(ctx context.Context, request *proto.Empty) (*proto.ShowGrowthConfigResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("ShowGrowthConfig params", logger.MakeFields(request)).Info("ShowGrowthConfig")
	resp := &proto.ShowGrowthConfigResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	config, err := model.ShowGrowConfigByMerchantID(merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("ShowGrowConfigByMerchantID 数据库错误, %v", err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}

	resp.Data = &proto.GrowthConfig{
		Name:  config.Name,
		Top:   config.Top,
		Rules: config.Rules.Slice(),
	}
	return resp, nil
}

// SaveGrowthConfig 保存
func (s *Server) SaveGrowthConfig(ctx context.Context, request *proto.SaveGrowthConfigRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("SaveGrowthConfig params", logger.MakeFields(request)).Info("SaveGrowthConfig")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	config := model.TableGrowthConfig{
		ID:         uuid.NewV4(),
		MerchantID: merchantID,
		Name:       request.Config.Name,
		Top:        request.Config.Top,
	}
	if len(request.Config.Rules) > 0 {
		config.Rules = (*fields.StringArr)(&request.Config.Rules)
	}
	if err := model.CreateOrUpdateGrowthConfig(config); err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateOrUpdateGrowthConfig 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
	}
	return resp, nil
}
