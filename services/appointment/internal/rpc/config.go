package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	mqMessage "gitlab.omytech.com.cn/micro-service/Crius/pkgs/message"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gorm.io/gorm"
)

// UpdateAppointmentConfig 修改配置信息
func (s *Server) UpdateAppointmentConfig(ctx context.Context, req *proto.AppointmentConfig) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentConfig")
	metadata := pkgs.GetMetadata(ctx)
	var config model.TableAppointmentConfig
	if err := s.database.Conn.Model(&model.TableAppointmentConfig{}).Scopes(util.ColumnEqualScope("merchant_id", metadata.MerchantID)).First(&config).Error; nil != err {
		// 不存在记录 保存
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config = toTableAppointmentConfig(req)
			config.MerchantID = metadata.MerchantID
			if err := s.database.Conn.Create(&config).Error; nil != err {
				util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("保存预约配置失败")
				return &proto.Response{
					ErrorCode:    pkgs.ErrInternal,
					ErrorMessage: fmt.Sprintf("保存数据失败:%s", err.Error()),
				}, nil
			}
			after := pkgs.MakeParams(config)
			go s.SaveSnapshot(ctx, model.TableSnapshot{
				SnapShotTableName: config.TableName(),
				TableID:           &config.ID,
				Method:            model.CreateMethod,
				After:             &after,
			})

			return &proto.Response{
				ErrorCode:    pkgs.Success,
				ErrorMessage: "",
			}, nil
		}

		// 抛出错误
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("修改预约配置，查询数据库失败")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, nil
	}

	// 修改配置信息
	config.KeepTime = int16(req.KeepTime)
	config.RemindTime = req.RemindTime
	config.OrderLimit = int16(req.OrderLimit)
	config.RoomNumWarn = int16(req.RoomNumWarn)
	config.PaymentTime = int16(req.PaymentTime)
	config.CancelTime = req.CancelTime
	config.RefundPercentBefore = req.RefundPercentBefore
	config.RefundPercentAfter = req.RefundPercentAfter
	config.BreachMonths = int16(req.BreachMonths)
	config.BreachTotal = int16(req.BreachTotal)
	config.DecorationFee = req.DecorationFee
	config.ThemeKeepTime = int16(req.ThemeKeepTime)
	config.ThemeCancelTime = req.ThemeCancelTime
	config.ThemeRefundPercentBefore = req.ThemeRefundPercentBefore
	config.ThemeRefundPercentAfter = req.ThemeRefundPercentAfter
	if err := s.database.Conn.Save(&config).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("更新预约配置失败")
		return &proto.Response{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("更新数据库失败:%s", err.Error()),
		}, nil
	}

	before := pkgs.MakeParams(config)
	after := pkgs.MakeParams(req)
	go s.SaveSnapshot(ctx, model.TableSnapshot{
		SnapShotTableName: config.TableName(),
		TableID:           &config.ID,
		Method:            model.UpdateMethod,
		Before:            &before,
		After:             &after,
	})

	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_AppointmentConfig,
	}
	go s.PublishDataChangeEvent(message)
	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

func toTableAppointmentConfig(req *proto.AppointmentConfig) model.TableAppointmentConfig {
	return model.TableAppointmentConfig{
		ID:                       uuid.NewV4(),
		KeepTime:                 int16(req.KeepTime),
		RemindTime:               req.RemindTime,
		OrderLimit:               int16(req.OrderLimit),
		RoomNumWarn:              int16(req.RoomNumWarn),
		PaymentTime:              int16(req.PaymentTime),
		CancelTime:               req.CancelTime,
		RefundPercentBefore:      req.RefundPercentBefore,
		RefundPercentAfter:       req.RefundPercentBefore,
		BreachMonths:             int16(req.BreachMonths),
		BreachTotal:              int16(req.BreachTotal),
		DecorationFee:            req.DecorationFee,
		ThemeKeepTime:            int16(req.ThemeKeepTime),
		ThemeCancelTime:          req.ThemeCancelTime,
		ThemeRefundPercentBefore: req.ThemeRefundPercentBefore,
		ThemeRefundPercentAfter:  req.ThemeRefundPercentAfter,
	}
}

// GetAppointmentConfig 获取预约配置信息
func (s *Server) GetAppointmentConfig(ctx context.Context, req *proto.Empty) (*proto.GetAppointmentConfigResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentConfig")
	metadata := pkgs.GetMetadata(ctx)
	var config model.TableAppointmentConfig
	if err := s.database.Conn.Model(&model.TableAppointmentConfig{}).Scopes(util.ColumnEqualScope("merchant_id", metadata.MerchantID)).First(&config).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.GetAppointmentConfigResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "没有该商户对应预约配置信息",
			}, nil
		}

		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("获取预约配置，查询数据库失败")
		return &proto.GetAppointmentConfigResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("数据库查询错误:%s", err.Error()),
		}, nil
	}

	return &proto.GetAppointmentConfigResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         toAppointmentConfigResponse(config),
	}, nil
}

func toAppointmentConfigResponse(config model.TableAppointmentConfig) *proto.AppointmentConfig {
	return &proto.AppointmentConfig{
		KeepTime:                 int32(config.KeepTime),
		RemindTime:               config.RemindTime,
		OrderLimit:               int32(config.OrderLimit),
		RoomNumWarn:              int32(config.RoomNumWarn),
		PaymentTime:              int32(config.PaymentTime),
		CancelTime:               config.CancelTime,
		RefundPercentBefore:      config.RefundPercentBefore,
		RefundPercentAfter:       config.RefundPercentAfter,
		BreachMonths:             int32(config.BreachMonths),
		BreachTotal:              int32(config.BreachTotal),
		DecorationFee:            config.DecorationFee,
		ThemeKeepTime:            int32(config.ThemeKeepTime),
		ThemeCancelTime:          config.ThemeCancelTime,
		ThemeRefundPercentBefore: config.ThemeRefundPercentBefore,
		ThemeRefundPercentAfter:  config.ThemeRefundPercentAfter,
	}
}

func (s *Server) getAppointmentConfig(merchantID uuid.UUID) (*model.TableAppointmentConfig, error) {
	config := new(model.TableAppointmentConfig)

	// 获取缓存中的配置
	bs, err := s.cache.Get(fmt.Sprintf(model.AppointmentConfigKey, merchantID.String()))
	if err == nil {
		err = json.Unmarshal(bs, config)
		if err == nil {
			return config, nil
		}
	}
	// 获取数据库配置
	if err := s.database.Conn.Scopes(crius.ColumnEqualScope("merchant_id", merchantID)).Take(config).Error; err != nil {
		return nil, err
	}
	// 反写配置入缓存
	bs, _ = json.Marshal(config)
	s.cache.Set(fmt.Sprintf(model.AppointmentConfigKey, merchantID.String()), bs)
	return config, nil
}
