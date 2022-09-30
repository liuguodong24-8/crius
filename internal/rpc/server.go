package rpc

import (
	"context"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/internal/config"
	"gitlab.omytech.com.cn/micro-service/Crius/internal/micro"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/redis"
	"gitlab.omytech.com.cn/micro-service/Crius/proto"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// Server 服务实例
type Server struct {
	proto.UnimplementedCriusServiceServer
	register *micro.Register
	discover *micro.Discover
}

// NewCriusServe 实例化服务对象
func NewCriusServe() (*Server, error) {
	cfg := config.Setting.Redis
	rds, err := redis.NewEntity(redis.Config{
		IP:       cfg.IP,
		Port:     cfg.Port,
		Password: cfg.Password,
		Database: cfg.Database,
	})

	if err != nil {
		return nil, fmt.Errorf("redis error:%s", err.Error())
	}

	return &Server{
		register: micro.NewRegister(rds),
		discover: micro.NewDiscover(rds),
	}, nil
}

// Heart 心跳
func (s *Server) Heart(ctx context.Context, req *proto.Service) (*proto.Empty, error) {
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(&req)).Debug("服务心跳信息[begin]")

	cfg := micro.RegisterRequest{
		IP:     req.Ip,
		Port:   req.Port,
		Name:   req.Name,
		Desc:   req.Desc,
		Weight: req.Weight,
	}
	if err := s.register.Register(cfg); nil != err {
		util.Logger.WithMetadata(ctx).Error(fmt.Sprintf("服务注册失败:%s", err.Error()))
		return &proto.Empty{}, err
	}

	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Debug("服务心跳信息[done]")

	return &proto.Empty{}, nil
}

// Discover 服务发现
func (s *Server) Discover(ctx context.Context, req *proto.DiscoverRequest) (*proto.Service, error) {
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Debug("服务发现[begin]")

	res, err := s.discover.Discover(micro.DiscoverRequest{Name: req.Name})

	if err != nil {
		util.Logger.WithMetadata(ctx).Error(fmt.Sprintf("服务发现错误:%s", err.Error()))
		return nil, err
	}

	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(res)).Debug("服务发现[end]")

	return &proto.Service{
		Name:   res.Name,
		Desc:   res.Desc,
		Ip:     res.IP,
		Port:   res.Port,
		Weight: res.Weight,
	}, nil
}

// DiscoverServers 批量发现服务
func (s *Server) DiscoverServers(ctx context.Context, req *proto.DiscoverServersRequest) (*proto.DiscoverServersResponse, error) {
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Debug("批量服务发现[begin]")

	var servers []*proto.Service

	res, err := s.discover.DiscoverServers(micro.DiscoverServersRequest{Names: req.Names})
	if err != nil {
		util.Logger.WithMetadata(ctx).Error(fmt.Sprintf("批量服务发现错误:%s", err.Error()))
		return nil, err
	}

	for _, s := range res.Servers {
		servers = append(servers, &proto.Service{
			Name:   s.Name,
			Desc:   s.Desc,
			Ip:     s.IP,
			Port:   s.Port,
			Weight: s.Weight,
		})
	}

	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(servers)).Debug("批量服务发现[end]")

	return &proto.DiscoverServersResponse{Services: servers}, nil
}
