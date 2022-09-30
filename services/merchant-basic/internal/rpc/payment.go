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

// GetBranchWithSubMchID 列表
func (s *Server) GetBranchWithSubMchID(ctx context.Context, request *proto.GetBranchWithSubMchIDRequest) (*proto.GetBranchWithSubMchIDResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBranchWithSubMchID")
	resp := &proto.GetBranchWithSubMchIDResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	count, err := model.GetBranchesWithSubMchIDCount(request.ProvinceId, request.CityId, request.DistrictId, request.BranchName, merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranchesWithSubMchIDCount 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	var branchProto []*proto.BranchSubMchID
	if count != 0 {
		branches, err := model.GetBranchesWithSubMchID(request.ProvinceId, request.CityId, request.DistrictId, request.BranchName, merchantID, request.Offset, request.Limit)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("GetBranchesWithSubMchID 数据库错误:%+v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
		if branches != nil {
			for _, b := range *branches {
				branch := &proto.BranchSubMchID{
					BranchId:   b.ID.String(),
					BranchName: b.Name,
					SubMchId:   b.SubMchID,
				}
				branchProto = append(branchProto, branch)
			}
		}
	}
	respData := proto.BranchSubMchIDData{
		Branches: branchProto,
		Total:    int32(count),
	}
	resp.Data = &respData
	return resp, nil
}

// SetBranchSubMchID 设置门店subMchID
func (s *Server) SetBranchSubMchID(ctx context.Context, request *proto.SetBranchSubMchIDRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("SetBranchSubMchID")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}
	branchID := uuid.FromStringOrNil(request.BranchId)
	if branchID == uuid.Nil {
		crius.Logger.Error(fmt.Sprintf("SetBranchSubMchID 参数错误:%+v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}
	err := model.SetBranchSubMchID(branchID, request.SubMchId)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("SetBranchSubMchID 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	return resp, nil
}

// GetWechatPaySetting 获取支付设置
func (s *Server) GetWechatPaySetting(ctx context.Context, request *proto.GetWechatPaySettingRequest) (*proto.GetWechatPaySettingResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetWechatPaySettingRequest")
	resp := &proto.GetWechatPaySettingResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	setting, err := model.GetWechatPaySetting(merchantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("GetWechatPaySetting 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	resp.Data = &proto.WechatPaySetting{
		MerchantId:           setting.MerchantID.String(),
		AppId:                setting.AppID,
		MchId:                setting.MchID,
		HeadquartersSubMchId: setting.HeadquartersSubMchID,
		PrivateKey:           setting.PrivateKey,
		CertFilename:         setting.CertFilename,
		CertContent:          setting.CertContent,
	}

	return resp, nil
}

// GetWechatPaySettingByAppID 根据appid获取支付设置
func (s *Server) GetWechatPaySettingByAppID(ctx context.Context, request *proto.GetWechatPaySettingByAppIDRequest) (*proto.GetWechatPaySettingResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetWechatPaySettingByAppID")
	resp := &proto.GetWechatPaySettingResponse{
		ErrorCode: pkgs.Success,
	}
	appID := request.AppId

	setting, err := model.GetWechatPaySettingByAppID(appID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("GetWechatPaySettingByAppID 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	resp.Data = &proto.WechatPaySetting{
		MerchantId:           setting.MerchantID.String(),
		AppId:                setting.AppID,
		MchId:                setting.MchID,
		HeadquartersSubMchId: setting.HeadquartersSubMchID,
		PrivateKey:           setting.PrivateKey,
		CertFilename:         setting.CertFilename,
		CertContent:          setting.CertContent,
	}

	return resp, nil
}

// SetWechatPaySetting 保存支付设置
func (s *Server) SetWechatPaySetting(ctx context.Context, request *proto.SetWechatPaySettingRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("SetWechatPaySettingRequest")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if merchantID == uuid.Nil || request.AppId == "" || request.MchId == "" ||
		request.PrivateKey == "" {
		crius.Logger.Error(fmt.Sprintf("SetWechatPaySettingRequest 参数错误:%+v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}
	setting := model.TableWechatPay{
		MerchantID:           merchantID,
		AppID:                request.AppId,
		MchID:                request.MchId,
		HeadquartersSubMchID: request.HeadquartersSubMchId,
		PrivateKey:           request.PrivateKey,
		CertFilename:         request.CertFilename,
		CertContent:          request.CertContent,
	}
	if err := model.SaveOrCreateWechatPaySetting(&setting); err != nil {
		crius.Logger.Error(fmt.Sprintf("SetWechatPaySettingRequest 数据库错误:%+v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	return resp, nil
}

// GetBranchWechatPaymentSetting 根据branch_id获取微信支付配置
func (s *Server) GetBranchWechatPaymentSetting(ctx context.Context, request *proto.GetBranchWechatPaymentSettingRequest) (*proto.GetBranchWechatPaymentSettingResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBranchWechatPaymentSetting")
	resp := &proto.GetBranchWechatPaymentSettingResponse{
		ErrorCode: pkgs.Success,
	}
	branchID := uuid.FromStringOrNil(request.BranchId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if branchID == uuid.Nil {
		//总部子商户ID
		setting, err := model.GetWechatPaySetting(merchantID)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("GetWechatPaySetting 数据库错误:%+v", err))
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resp.ErrorCode = pkgs.ErrUnprocessableEntity
				resp.ErrorMessage = "商户支付配置不存在"
				return resp, nil
			}
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "数据库错误"
			return resp, nil
		}
		resp.Data = &proto.BranchWechatPaymentSetting{
			MerchantId:           setting.MerchantID.String(),
			HeadquartersSubMchId: setting.HeadquartersSubMchID,
			AppId:                setting.AppID,
			MchId:                setting.MchID,
			PrivateKey:           setting.PrivateKey,
			CertFilename:         setting.CertFilename,
			CertContent:          setting.CertContent,
		}
		return resp, nil
	}

	setting, err := model.GetBranchWechatPaymentSetting(branchID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranchWechatPaymentSetting 数据库错误:%+v", err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "门店不存在"
			return resp, nil
		}
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}

	resp.Data = &proto.BranchWechatPaymentSetting{
		BranchId:             setting.BranchID.String(),
		MerchantId:           setting.MerchantID.String(),
		HeadquartersSubMchId: setting.HeadquartersSubMchID,
		AppId:                setting.AppID,
		MchId:                setting.MchID,
		SubMchId:             setting.SubMchID,
		PrivateKey:           setting.PrivateKey,
		CertFilename:         setting.CertFilename,
		CertContent:          setting.CertContent,
	}

	return resp, nil
}
