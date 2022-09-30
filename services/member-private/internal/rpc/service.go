package rpc

import (
	"context"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"google.golang.org/grpc"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/db"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/member-private/internal/config"
	"gitlab.omytech.com.cn/micro-service/member-private/internal/model"
	"gitlab.omytech.com.cn/micro-service/member-private/proto"
	merchantBasic "gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// Server server
type Server struct {
	proto.UnimplementedMemberPrivateServerServer
	// 服务发现注册
	crius *crius.Client
	// 数据库
	database *db.Entity
}

var merchantBasicClient merchantBasic.MerchantBasicServiceClient

// NewMemberPrivateServe 实例化
func NewMemberPrivateServe() (*Server, error) {
	client, err := crius.NewClient(context.Background(), crius.ClientConfig{Address: config.Setting.Crius.Address})
	if err != nil {
		return nil, fmt.Errorf("crius client失败:%s", err.Error())
	}

	entity, err := model.DatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("数据库连接错误:%s", err.Error())
	}

	return &Server{
		crius:    client,
		database: entity,
	}, nil
}

func (s *Server) merchantBasic() merchantBasic.MerchantBasicServiceClient {
	if merchantBasicClient == nil {
		resp, err := s.crius.Discover(context.Background(), crius.DiscoverRequest{Name: config.Setting.Crius.MerchantBasic})
		if err != nil {
			util.Logger.Error(fmt.Sprintf("merchantBasic发现服务错误:%v, 服务名:%v", err, config.Setting.Crius.MerchantBasic))
			return nil
		}
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", resp.IP, resp.Port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			util.Logger.Error(fmt.Sprintf("merchantBasic rpc连接错误err:%v, ip:port:%v", err, fmt.Sprintf("%s:%d", resp.IP, resp.Port)))
			return nil
		}
		merchantBasicClient = merchantBasic.NewMerchantBasicServiceClient(conn)
	}
	return merchantBasicClient
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

// SaveSnapshot 保存快照
func (s *Server) SaveSnapshot(ctx context.Context, snapshot model.Snapshot) {
	snapshot.ID = uuid.NewV4()
	metadata := pkgs.GetMetadata(ctx)
	snapshot.SleuthCode = metadata.SleuthCode
	if err := s.database.Conn.Create(&snapshot).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(snapshot)).Info("保存快照信息错误")
	}
}
