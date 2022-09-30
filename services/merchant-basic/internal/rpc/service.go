package rpc

import (
	"context"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/config"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Server server
type Server struct {
	proto.MerchantBasicServiceServer
	// 服务发现注册
	crius *crius.Client
}

// NewMerchantBasicServe 实例化
func NewMerchantBasicServe() (*Server, error) {

	client, err := crius.NewClient(context.Background(), crius.ClientConfig{Address: config.Setting.Crius.Address})
	if err != nil {
		return nil, err
	}

	return &Server{
		crius: client,
	}, nil
}

// Heart 心跳
func (s *Server) Heart(ctx context.Context) error {
	cfg := config.Setting.App
	// util.Logger.WithFields("config", logger.MakeFields(cfg)).Info("basic 心跳")
	return s.crius.Heart(ctx, crius.HeartRequest{
		Name:   cfg.Name,
		Desc:   cfg.Desc,
		IP:     cfg.IP,
		Port:   cfg.Port,
		Weight: cfg.Weight,
	})
}
