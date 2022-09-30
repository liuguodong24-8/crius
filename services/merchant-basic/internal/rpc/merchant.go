package rpc

import (
	"context"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// UpdateMerchantLogo 更新商户logo
func (s *Server) UpdateMerchantLogo(ctx context.Context, request *proto.UpdateMerchantLogoRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateMerchantLogo")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	merchant := model.TableMerchant{
		ID:   merchantID,
		Logo: request.Logo,
	}
	if err := model.UpdateMerchantLogo(&merchant); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateMerchantLogo 更新商户信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新商户失败"
		return resp, nil
	}

	return resp, nil
}

// UpdateMerchantUserAgreement 更新商户用户协议
func (s *Server) UpdateMerchantUserAgreement(ctx context.Context, request *proto.UpdateMerchantUserAgreementRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateMerchantUserAgreement")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	merchant := model.TableMerchant{
		ID: merchantID,
		UserAgreement: &model.Agreement{
			Agreement:  request.UserAgreement,
			FileFormat: request.AgreementFileFormat,
		},
	}
	if err := model.UpdateMerchantUserAgreement(&merchant); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateMerchantUserAgreement 更新商户信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新商户失败"
		return resp, nil
	}

	return resp, nil
}

// ShowMerchant 查询商户信息
func (s *Server) ShowMerchant(ctx context.Context, request *proto.Empty) (*proto.ShowMerchantResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateMerchant")
	resp := &proto.ShowMerchantResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	merchant, err := model.ShowMerchant(merchantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "商户不存在"
			return resp, nil
		}
		util.Logger.Error(fmt.Sprintf("UpdateMerchant 更新商户信息数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新商户失败"
		return resp, nil
	}
	resp.Data = &proto.ShowMerchantResponse_Data{
		Logo: merchant.Logo,
		Name: merchant.Name,
	}
	if merchant.UserAgreement != nil {
		resp.Data.UserAgreement = merchant.UserAgreement.Agreement
		resp.Data.AgreementFileFormat = merchant.UserAgreement.FileFormat
	}

	return resp, nil
}
