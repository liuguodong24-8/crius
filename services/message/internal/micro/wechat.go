package micro

import (
	"context"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/message/internal/config"
	wechatProto "gitlab.omytech.com.cn/micro-service/wechat/proto"
	"google.golang.org/grpc"
)

// WechatClient 微信client
type WechatClient struct {
	Entity wechatProto.WechatServiceClient
}

// GetWechatServer 获取微信client
func GetWechatServer(ctx context.Context, client *crius.Client) (*WechatClient, error) {
	server, err := client.Discover(ctx, crius.DiscoverRequest{Name: config.Setting.Crius.Wechat})
	if err != nil {
		return nil, fmt.Errorf("获取wechat server失败:%s", err.Error())
	}

	util.Logger.WithMetadata(ctx).WithFields("wechat server", logger.MakeFields(server)).Info("发现wechat服务")

	conn, connErr := grpc.Dial(fmt.Sprintf("%s:%d", server.IP, server.Port), grpc.WithInsecure())
	if connErr != nil {
		return nil, fmt.Errorf("实例化wechat服务失败:%s", connErr.Error())
	}

	return &WechatClient{Entity: wechatProto.NewWechatServiceClient(conn)}, nil
}
