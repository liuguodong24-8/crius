package rpc

import (
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// CreateMemberAddress 新增用户地址
func (s *Server) CreateMemberAddress(ctx context.Context, req *proto.MemberAddress) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("req", logger.MakeFields(req)).Info("CreateMemberAddress")

	conn := model.DatabaseConn().Begin()

	address := &model.TableMemberAddress{
		ID:         uuid.NewV4(),
		MemberID:   uuid.FromStringOrNil(req.MemberId),
		Phone:      req.Phone,
		PhoneCode:  req.PhoneCode,
		ProvinceID: req.ProvinceId,
		CityID:     req.CityId,
		DistrictID: req.DistrictId,
		Address:    req.Address,
		Name:       req.Name,
		IsDefault:  req.IsDefault,
	}

	if err := conn.Create(address).Error; err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("保存用户地址信息错误")
		conn.Rollback()
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("保存地址信息错误:%s", err.Error()),
		}, nil
	}

	// 修改时 设置当前为默认地址，更改其他的为非默认
	if address.IsDefault {
		if err := conn.Model(&model.TableMemberAddress{}).Scopes(util.ColumnEqualScope("member_id", address.MemberID), util.ColumnSymbolScope("id", "!=", address.ID)).Updates(map[string]interface{}{"is_default": false}).Error; nil != err {
			util.Logger.WithMetadata(ctx).WithError(err).Error("保存用户地址默认信息错误")
			conn.Rollback()
			return &proto.Response{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: fmt.Sprintf("保存地址信息错误:%s", err.Error()),
			}, nil
		}
	}
	conn.Commit()
	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// UpdateMemberAddress 修改用户地址
func (s *Server) UpdateMemberAddress(ctx context.Context, req *proto.MemberAddress) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("req", logger.MakeFields(req)).Info("UpdateMemberAddress")

	var address model.TableMemberAddress
	if err := model.DatabaseConn().Model(&model.TableMemberAddress{}).Scopes(util.ColumnEqualScope("id", req.Id)).First(&address).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.Response{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "",
			}, nil
		}

		util.Logger.WithMetadata(ctx).WithError(err).Error("查找地址信息错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查找地址信息错误:%s", err.Error()),
		}, nil
	}

	conn := model.DatabaseConn().Begin()

	address.Phone = req.Phone
	address.PhoneCode = req.PhoneCode
	address.ProvinceID = req.ProvinceId
	address.CityID = req.CityId
	address.DistrictID = req.DistrictId
	address.Address = req.Address
	address.Name = req.Name
	address.IsDefault = req.IsDefault

	if err := conn.Save(&address).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).WithFields("address", logger.MakeFields(address)).Error("修改地址信息错误")
		conn.Rollback()
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("修改地址信息错误:%s", err.Error()),
		}, nil
	}

	// 修改时 设置当前为默认地址，更改其他的为非默认
	if address.IsDefault {
		if err := conn.Model(&model.TableMemberAddress{}).Scopes(util.ColumnEqualScope("member_id", address.MemberID), util.ColumnSymbolScope("id", "!=", address.ID)).Updates(map[string]interface{}{"is_default": false}).Error; nil != err {
			util.Logger.WithMetadata(ctx).WithError(err).Error("修改地址信息默认错误")
			conn.Rollback()
			return &proto.Response{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: fmt.Sprintf("修改地址信息错误:%s", err.Error()),
			}, nil
		}
	}

	conn.Commit()

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// GetMemberAddress 获取用户地址
func (s *Server) GetMemberAddress(ctx context.Context, req *proto.GetMemberAddressRequest) (*proto.GetMemberAddressResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("req", logger.MakeFields(req)).Info("GetMemberAddressResponse")

	address, _, err := getMemberAddress(req.MemberId, false)
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询用户地址列表错误")
		return &proto.GetMemberAddressResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询用户地址列表错误:%s", err.Error()),
		}, nil
	}

	return &proto.GetMemberAddressResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         address,
	}, nil
}

// SetMemberAddressDefault 设置为默认地址
func (s *Server) SetMemberAddressDefault(ctx context.Context, req *proto.SetMemberAddressDefaultRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("req", logger.MakeFields(req)).Info("GetMemberAddressResponse")

	conn := model.DatabaseConn().Begin()
	// 设置
	if err := conn.Model(&model.TableMemberAddress{}).Where("member_id", req.MemberId).Updates(map[string]interface{}{"is_default": false}).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("修改默认地址错误")
		conn.Rollback()
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("修改错误:%s", err.Error()),
		}, nil
	}

	if err := conn.Model(&model.TableMemberAddress{}).Where("id", req.Id).Updates(map[string]interface{}{"is_default": true}).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("修改默认地址设置错误")
		conn.Rollback()
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("修改错误:%s", err.Error()),
		}, nil
	}

	conn.Commit()

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// GetMemberDefaultAddress 获取用户默认地址
func (s *Server) GetMemberDefaultAddress(ctx context.Context, req *proto.GetMemberDefaultAddressRequest) (*proto.GetMemberDefaultAddressResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("req", logger.MakeFields(req)).Info("GetMemberDefaultAddress")

	_, address, err := getMemberAddress(req.MemberId, true)
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("查询用户地址列表错误")
		return &proto.GetMemberDefaultAddressResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询用户地址列表错误:%s", err.Error()),
		}, nil
	}

	return &proto.GetMemberDefaultAddressResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         address,
	}, nil
}

// DeleteMemberAddress 删除用户地址
func (s *Server) DeleteMemberAddress(ctx context.Context, req *proto.DeleteMemberAddressRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("req", logger.MakeFields(req)).Info("GetMemberDefaultAddress")

	var address model.TableMemberAddress
	if err := model.DatabaseConn().Model(&model.TableMemberAddress{}).Scopes(util.ColumnEqualScope("id", req.Id)).First(&address).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.Response{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: fmt.Sprintf("[%s] 对应数据不存在", req.Id),
			}, nil
		}

		util.Logger.WithMetadata(ctx).WithError(err).Error("删除用户地址查询错误")

		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询错误:%s", err.Error()),
		}, nil
	}

	if err := model.DatabaseConn().Model(&address).Scopes(util.ColumnEqualScope("id", req.Id)).Delete(&req.Id).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("删除用户地址错误")

		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("删除错误:%s", err.Error()),
		}, nil
	}

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

func getMemberAddress(memberID string, dealDefault bool) ([]*proto.MemberAddress, *proto.MemberAddress, error) {
	if len(memberID) == 0 {
		return nil, nil, fmt.Errorf("[%s] 错误", memberID)
	}

	var items []model.TableMemberAddress
	if err := model.DatabaseConn().Model(&model.TableMemberAddress{}).Scopes(util.ColumnEqualScope("member_id", memberID)).Order("created_at desc").Find(&items).Error; nil != err {
		return nil, nil, fmt.Errorf("查询用户地址错误:%s", err.Error())
	}

	var data []*proto.MemberAddress
	var defaultAddress *proto.MemberAddress
	for _, item := range items {
		address := &proto.MemberAddress{
			Id:         item.ID.String(),
			MemberId:   item.MemberID.String(),
			Phone:      item.Phone,
			PhoneCode:  item.PhoneCode,
			ProvinceId: item.ProvinceID,
			CityId:     item.CityID,
			DistrictId: item.DistrictID,
			Address:    item.Address,
			Name:       item.Name,
			IsDefault:  item.IsDefault,
		}

		data = append(data, address)

		if item.IsDefault {
			defaultAddress = address
		}
	}

	if !dealDefault || data == nil {
		return data, nil, nil
	}

	// 如不存在默认，用创建时间第一条默认
	if defaultAddress == nil {
		idx := len(data) - 1
		first := data[idx]
		defaultAddress = first
		first.IsDefault = true
		data[idx] = first
	}

	return data, defaultAddress, nil
}
