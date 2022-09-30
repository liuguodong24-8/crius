package rpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gitlab.omytech.com.cn/micro-service/appointment/util"
	"gorm.io/gorm"
)

// SaveAppointmentLock 保存预约锁
func (s *Server) SaveAppointmentLock(ctx context.Context, req *proto.SaveAppointmentLockRequest) (*proto.SaveAppointmentLockResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("SaveAppointmentLock")
	resp := &proto.SaveAppointmentLockResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	appointmentDate := time.Unix(int64(req.AppointmentDate), 0)
	id := uuid.FromStringOrNil(req.Id)
	appointmentLock := model.AppointmentLock{
		BranchID:        req.BranchId,
		RoomGroupID:     req.RoomGroupId,
		Way:             int8(req.Way),
		AppointmentDate: appointmentDate.Format("2006-01-02"),
		AppointmentAt:   req.AppointmentAt,
	}

	// id不存在则新增，存在则修改
	if id == uuid.Nil {
		id = uuid.NewV4()
	} else {
		score, err := s.redis.ZSCORE(model.AppointmentZsetLockKey, id.String())
		if err != nil && err != redigo.ErrNil {
			crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 获取锁过期时间redis错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "保存锁失败"
			return resp, nil
		}
		// 锁没有过期，更新锁过期时间
		if !util.AppointLockUpdateExpired(int64(score)) && err != redigo.ErrNil {
			if err := s.redis.ZADD(model.AppointmentZsetLockKey, map[int64]interface{}{
				time.Now().Add(util.AppointLockExpireTime).Add(util.AppointLockDelayTime).Unix(): id.String()}); err != nil {
				crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 更新锁redis错误:%v", err))
				resp.ErrorCode = pkgs.ErrInternal
				resp.ErrorMessage = "保存锁失败"
				return resp, nil
			}
			resp.Data = id.String()
			return resp, nil
		}
		id = uuid.NewV4()
	}

	tx := s.database.Conn.Begin()
	dateCounter, total, err := getDateCounter(tx, appointmentLock, merchantID)
	if err != nil {
		tx.Rollback()
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存锁失败"
		return resp, nil
	}

	// 锁定数量自增1
	if err := tx.Model(&model.TableAppointmentDateCounter{}).Scopes(crius.ColumnEqualScope("id", dateCounter.ID)).Update("appoint_num", gorm.Expr("appoint_num+1")).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 预约数量增加数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存锁失败"
		return resp, nil
	}

	var tabDateCounter model.TableAppointmentDateCounter
	if err := tx.Scopes(crius.ColumnEqualScope("id", dateCounter.ID)).Take(&tabDateCounter).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 预约数量查询数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存锁失败"
		return resp, nil
	}
	if tabDateCounter.AppointNum > total {
		tx.Rollback()
		resp.ErrorCode = pkgs.ErrNumberLimit
		resp.ErrorMessage = "预约房型数量达上限"
		return resp, nil
	}

	//创建数据
	if err := s.redis.HMSET(fmt.Sprintf(model.AppointmentHashLockKey, id.String()), appointmentLock); err != nil {
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 创建redis hash错误:%v", err))
		tx.Rollback()
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存锁失败"
		return resp, nil
	}

	if err := s.redis.ZADD(model.AppointmentZsetLockKey, map[int64]interface{}{
		time.Now().Add(util.AppointLockExpireTime).Add(util.AppointLockDelayTime).Unix(): id.String()}); err != nil {
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 创建redis zset错误:%v", err))
		tx.Rollback()
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存锁失败"
		return resp, nil
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 提交事务数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "保存锁失败"
		return resp, nil
	}

	resp.Data = id.String()
	return resp, nil
}

// DeleteAppointmentLock 删除预约锁
func (s *Server) DeleteAppointmentLock(ctx context.Context, req *proto.DeleteAppointmentLockRequest) (*proto.Response, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("DeleteAppointmentLock")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	lock := new(model.AppointmentLock)
	if err := s.redis.HGETALLSTRUCT(fmt.Sprintf(model.AppointmentHashLockKey, req.Id), lock); err != nil {
		crius.Logger.Error(fmt.Sprintf("DeleteAppointmentLock 获取锁数据redis错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除预约锁失败"
		return resp, nil
	}
	tx := s.database.Conn.Begin()
	if err := tx.Model(&model.TableAppointmentDateCounter{}).Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnEqualScope("branch_id", lock.BranchID),
		crius.ColumnEqualScope("room_group_id", lock.RoomGroupID), crius.ColumnEqualScope("appointment_time", time.Unix(int64(lock.AppointmentAt), 0)),
		model.WayAnd(lock.Way)).Update("appoint_num", gorm.Expr("appoint_num-1")).Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("DeleteAppointmentLock 预约数量自减数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除预约锁失败"
		return resp, nil
	}
	if n, err := s.redis.ZREM(model.AppointmentZsetLockKey, req.Id); err != nil || n == 0 {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("DeleteAppointmentLock 删除锁redis错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除预约锁失败"
		return resp, nil
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		crius.Logger.Error(fmt.Sprintf("DeleteAppointmentLock 提交事务失败数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "删除预约锁失败"
		return resp, nil
	}
	s.redis.DEL(fmt.Sprintf(model.AppointmentHashLockKey, req.Id))
	return resp, nil
}

func getDateCounter(tx *gorm.DB, lock model.AppointmentLock, merchantID uuid.UUID) (*model.TableAppointmentDateCounter, int32, error) {
	var dateCounter model.TableAppointmentDateCounter
	createDateCounter := false
	if err := tx.Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnEqualScope("branch_id", lock.BranchID), crius.ColumnEqualScope("appointment_time", time.Unix(int64(lock.AppointmentAt),
		0)), model.WayAnd(lock.Way), crius.ColumnEqualScope("room_group_id", lock.RoomGroupID)).Take(&dateCounter).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 查询错误:%v", err))
			return nil, 0, err
		}
		createDateCounter = true
	}
	total := dateCounter.Number

	//没有dateCounter数据，或者没有设置临时房型总数
	if createDateCounter || total == 0 {
		// 没有预约过，数据不存在
		var templateConfig model.TableAppointmentTemplateConfig
		templateID, err := getTemplateID(tx, merchantID, uuid.FromStringOrNil(lock.BranchID), lock.AppointmentDate)
		if err != nil {
			return nil, 0, err
		}

		if err := tx.Scopes(crius.ColumnEqualScope("template_id", templateID), crius.ColumnEqualScope("room_type_id", lock.RoomGroupID)).Take(&templateConfig).Error; err != nil {
			crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 获取模板配置数据库错误:%v", err))
			return nil, 0, err
		}

		// 获取设置房型总数
		findConfigNum := false
		for _, v := range []map[string]interface{}(*templateConfig.Configure) {
			if lock.Way&int8(v["way"].(float64)) == lock.Way && time.Unix(int64(lock.AppointmentAt), 0).Format("15:04:05") == v["time"].(string) {
				total = int32(v["num"].(float64))
				if createDateCounter {
					branchID := uuid.FromStringOrNil(lock.BranchID)
					appointmentAt := time.Unix(int64(lock.AppointmentAt), 0)
					appointmentDate, _ := time.Parse("2006-01-02", lock.AppointmentDate)
					dateCounter = model.TableAppointmentDateCounter{
						ID:              uuid.NewV4(),
						BranchID:        &branchID,
						MerchantID:      &merchantID,
						RoomGroupID:     templateConfig.RoomTypeID,
						Way:             int8(v["way"].(float64)),
						AppointmentDate: &appointmentDate,
						AppointmentTime: &appointmentAt,
						AppointNum:      0,
					}
					if err := tx.Create(&dateCounter).Error; err != nil {
						crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 创建date_counter数据库错误:%v", err))
						return nil, 0, err
					}
				}
				findConfigNum = true
				break
			}
		}
		if !findConfigNum {
			crius.Logger.Error("SaveAppointmentLock 预约方式不匹配，没有找到或创建date_counter")
			return nil, 0, errors.New("预约方式不匹配，没有找到或创建date_counter")
		}
	}
	return &dateCounter, total, nil
}

func getTemplateID(tx *gorm.DB, merchantID, branchID uuid.UUID, appointmentDate string) (uuid.UUID, error) {
	var calendars []model.TableTemplateCalendar
	var templateID uuid.UUID
	if err := tx.Scopes(crius.ColumnEqualScope("merchant_id", merchantID), crius.ColumnEqualScope("branch_id", branchID),
		crius.ColumnEqualScope("business_date", appointmentDate)).Find(&calendars).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("SaveAppointmentLock 获取模板日历数据库错误:%v", err))
		return uuid.Nil, err
	}
	if len(calendars) > 1 {
		for _, v := range calendars {
			if v.Category == model.CalendarHoliday {
				templateID = v.TemplateID
			}
		}
	} else if len(calendars) == 1 {
		templateID = calendars[0].TemplateID
	} else {
		// 没有数据
		crius.Logger.Error("SaveAppointmentLock 模板日历数据不存在")
		return uuid.Nil, errors.New("模板日历数据不存在")
	}
	return templateID, nil
}
