package util

import (
	"context"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/basic/internal/config"
	"runtime/debug"
)

// Logger 日志entity
var Logger Entity

// Entity ...
type Entity struct {
	*logger.Entity
}

// RegisterLogger 注册日志模块
func RegisterLogger() error {
	cfg := config.Setting.Log
	entity, err := logger.NewLoggerEntity(logger.Config{
		Channel:    cfg.Channel,
		Level:      logger.Level(cfg.Level),
		OutputFile: cfg.Output,
		WithStack:  cfg.Stack,
	})
	if err != nil {
		return err
	}

	Logger = Entity{entity}

	return nil
}

// WithSleuthCode 日志记录调用链信息
func (e *Entity) WithSleuthCode(ctx context.Context) *Entity {
	entity := e.WithFields("sleuth_code", logger.Fields{
		"code": pkgs.GetContentSleuthCode(ctx),
	})

	return &Entity{entity}
}

// CatchException 捕获异常panic情况
func CatchException() {
	errs := recover()
	if nil == errs {
		return
	}

	Logger.WithFields("exception", logger.Fields{
		"exception stack": string(debug.Stack()),
	}).Error("捕获异常信息")
}
