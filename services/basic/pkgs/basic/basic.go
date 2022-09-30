package basic

import (
	"context"

	"gitlab.omytech.com.cn/micro-service/basic/proto"
	"google.golang.org/grpc"
)

// ClientConfig 实例化配置
type ClientConfig struct {
	Address string
}

// HelloRequest hello request
type HelloRequest struct {
	Message string
}

// HelloResponse hello response
type HelloResponse struct {
	Message string
}

// Interface 定义接口
type Interface interface {
	Hello(ctx context.Context, req HelloRequest) (*HelloResponse, error)
}

var _ Interface = (*Client)(nil)

// Client 实例化对象
type Client struct {
	client proto.BasicServiceClient
}

// NewClient 实例化
func NewClient(ctx context.Context, cfg ClientConfig) (*Client, error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		client: proto.NewBasicServiceClient(conn),
	}, nil
}

// Hello hello 调试
func (c *Client) Hello(ctx context.Context, req HelloRequest) (*HelloResponse, error) {
	res, err := c.client.Hello(ctx, &proto.HelloRequest{Message: req.Message})

	if err != nil {
		return nil, err
	}

	return &HelloResponse{
		Message: res.Message,
	}, nil
}
