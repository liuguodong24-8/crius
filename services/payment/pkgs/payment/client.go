package payment

import (
	"context"

	"gitlab.omytech.com.cn/micro-service/payment/proto"
	"google.golang.org/grpc"
)

// ClientConfig 实例化配置
type ClientConfig struct {
	Address string
}

// Response 统一返回
type Response struct {
	ErrorCode    int32
	ErrorMessage string
}

// Interface 定义接口
type Interface interface {
	WechatUnifiedOrder(ctx context.Context, req WechatUnifiedOrderRequest) (WechatUnifiedOrderResponse, error)
	WechatOrderQuery(ctx context.Context, req WechatOrderQueryRequest) (WechatOrderQueryResponse, error)
	WechatCloseOrder(ctx context.Context, req WechatCloseOrderRequest) (WechatCloseOrderResponse, error)
	WechatRefund(ctx context.Context, req WechatRefundRequest) (WechatRefundResponse, error)
}

var _ Interface = (*Client)(nil)

// Client 实例化对象
type Client struct {
	client proto.PaymentServerClient
}

// NewClient 实例化
func NewClient(ctx context.Context, cfg ClientConfig) (*Client, error) {
	grpc.WithBlock()
	conn, err := grpc.Dial(cfg.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		client: proto.NewPaymentServerClient(conn),
	}, nil
}
