package rpc

import (
	"context"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/basic/util"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/basic/internal/config"

	"gitlab.omytech.com.cn/micro-service/basic/proto"
)

// Server server
type Server struct {
	proto.UnimplementedBasicServiceServer
	// 服务发现注册
	crius *crius.Client
}

// NewBasicServe 实例化
func NewBasicServe() (*Server, error) {
	fmt.Println("new basic service")

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
	util.Logger.WithFields("config", logger.MakeFields(cfg)).Info("basic 心跳")
	return s.crius.Heart(ctx, crius.HeartRequest{
		Name:   cfg.Name,
		Desc:   cfg.Desc,
		IP:     cfg.IP,
		Port:   cfg.Port,
		Weight: cfg.Weight,
	})
}

// Hello 测试
func (s *Server) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("hello")
	fmt.Println(ctx)
	fmt.Println(req)
	return &proto.HelloResponse{Message: "hello world"}, nil
}
