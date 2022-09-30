package rpc

import (
	"context"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/wechat/internal/config"
	"gitlab.omytech.com.cn/micro-service/wechat/proto"
)

// Server server
type Server struct {
	proto.UnimplementedWechatServiceServer
	// 服务发现注册
	crius *crius.Client
}

// NewWechatServe 实例化
func NewWechatServe() (*Server, error) {
	client, err := crius.NewClient(context.Background(), crius.ClientConfig{Address: config.Setting.Crius.Address})
	if err != nil {
		return nil, fmt.Errorf("crius client失败:%s", err.Error())
	}

	return &Server{
		crius: client,
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
