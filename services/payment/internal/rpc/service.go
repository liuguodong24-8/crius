package rpc

import (
	"context"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/db"
	"gitlab.omytech.com.cn/micro-service/payment/internal/cache"
	"gitlab.omytech.com.cn/micro-service/payment/internal/config"
	"gitlab.omytech.com.cn/micro-service/payment/internal/model"
	"gitlab.omytech.com.cn/micro-service/payment/proto"
)

// Server server
type Server struct {
	proto.UnimplementedPaymentServerServer
	// 服务发现注册
	crius *crius.Client
	// 数据库
	database *db.Entity
	// 缓存
	cache *cache.Entity
}

// NewPaymentServe 实例化
func NewPaymentServe() (*Server, error) {
	client, err := crius.NewClient(context.Background(), crius.ClientConfig{Address: config.Setting.Crius.Address})
	if err != nil {
		return nil, fmt.Errorf("crius client失败:%s", err.Error())
	}

	dbEntity, dbErr := model.DatabaseConnection()
	if dbErr != nil {
		return nil, fmt.Errorf("数据库连接错误:%s", dbErr.Error())
	}

	cacheEntity, cacheErr := cache.NewEntity()
	if cacheErr != nil {
		return nil, fmt.Errorf("缓存实例失败:%s", cacheErr.Error())
	}

	return &Server{
		crius:    client,
		database: dbEntity,
		cache:    cacheEntity,
	}, nil
}

// Heart 心跳
func (s *Server) Heart(ctx context.Context) error {
	cfg := config.Setting.App
	return s.crius.Heart(ctx, crius.HeartRequest{
		Name:   cfg.Name,
		Desc:   cfg.Desc,
		IP:     cfg.IP,
		Port:   cfg.Port,
		Weight: cfg.Weight,
	})
}
