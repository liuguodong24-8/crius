package crius

import (
	"context"

	"gitlab.omytech.com.cn/micro-service/Crius/proto"
	"google.golang.org/grpc"
)

// ClientConfig 实例client配置
type ClientConfig struct {
	Address string
}

// HeartRequest 心跳请求
type HeartRequest struct {
	Name   string
	Desc   string
	IP     string
	Port   int64
	Weight int32
}

// DiscoverRequest 发现请求
type DiscoverRequest struct {
	Name string
}

// DiscoverServersRequest 批量发现服务请求
type DiscoverServersRequest struct {
	Names []string
}

// Server 服务
type Server struct {
	Name   string
	Desc   string
	IP     string
	Port   int64
	Weight int32
}

// Interface interface
type Interface interface {
	Heart(ctx context.Context, req HeartRequest) error
	Discover(ctx context.Context, req DiscoverRequest) (*Server, error)
	DiscoverServers(ctx context.Context, req DiscoverServersRequest) (*[]Server, error)
}

var _ Interface = (*Client)(nil)

// Client client
type Client struct {
	client proto.CriusServiceClient
}

// NewClient 实例化client
func NewClient(ctx context.Context, cfg ClientConfig) (*Client, error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		client: proto.NewCriusServiceClient(conn),
	}, nil
}

// Heart 心跳
func (c *Client) Heart(ctx context.Context, req HeartRequest) (err error) {
	_, err = c.client.Heart(ctx, &proto.Service{
		Name:   req.Name,
		Desc:   req.Desc,
		Ip:     req.IP,
		Port:   req.Port,
		Weight: req.Weight,
	})

	return
}

// Discover 发现
func (c *Client) Discover(ctx context.Context, req DiscoverRequest) (*Server, error) {
	s, err := c.client.Discover(ctx, &proto.DiscoverRequest{Name: req.Name})

	if err != nil {
		return nil, err
	}

	return &Server{
		Name:   s.Name,
		Desc:   s.Desc,
		IP:     s.Ip,
		Port:   s.Port,
		Weight: s.Weight,
	}, nil
}

// DiscoverServers 批量发现
func (c *Client) DiscoverServers(ctx context.Context, req DiscoverServersRequest) (*[]Server, error) {
	res, err := c.client.DiscoverServers(ctx, &proto.DiscoverServersRequest{Names: req.Names})
	if err != nil {
		return nil, err
	}

	var servers []Server

	for _, s := range res.Services {
		servers = append(servers, Server{
			Name:   s.Name,
			Desc:   s.Desc,
			IP:     s.Ip,
			Port:   s.Port,
			Weight: s.Weight,
		})
	}

	return &servers, nil
}
