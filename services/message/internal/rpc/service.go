package rpc

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/db"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	thirdParty "gitlab.omytech.com.cn/micro-service/Crius/pkgs/third_party"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/message/internal/cache"
	"gitlab.omytech.com.cn/micro-service/message/internal/config"
	"gitlab.omytech.com.cn/micro-service/message/internal/model"
	"gitlab.omytech.com.cn/micro-service/message/proto"
)

// Server server
type Server struct {
	proto.UnimplementedMessageServiceServer
	// 服务发现注册
	crius *crius.Client
	// 数据库
	database *db.Entity
	// 缓存
	cache *cache.Entity
}

// NewMessageServe 实例化
func NewMessageServe() (*Server, error) {

	client, err := crius.NewClient(context.Background(), crius.ClientConfig{Address: config.Setting.Crius.Address})
	if err != nil {
		return nil, fmt.Errorf("crius client失败:%s", err.Error())
	}

	entity, entityErr := model.DatabaseConnection()
	if entityErr != nil {
		return nil, fmt.Errorf("数据库连接错误:%s", entityErr.Error())
	}

	cacheEntity, cacheErr := cache.NewEntity()
	if cacheErr != nil {
		return nil, fmt.Errorf("缓存实例失败:%s", cacheErr.Error())
	}

	return &Server{
		crius:    client,
		database: entity,
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

// GetShortURL 获取百度短连接
func (s *Server) GetShortURL(ctx context.Context, req *proto.GetShortURLRequest) (*proto.GetShortURLResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("GetShortURL")

	baidu := thirdParty.NewBaidu()
	result, err := baidu.ShortURL(req.Url, thirdParty.ValidityLongTerm)
	util.Logger.WithFields("baidu response", logger.MakeFields(result)).Info("获取百度短链返回")
	if err != nil {
		util.Logger.WithError(err).Error("请求获取百度短链失败")
		return &proto.GetShortURLResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("获取百度短链失败:%s", err.Error()),
		}, nil
	}

	data := &proto.GetShortURLResponse_Data{
		Url: result.URL,
	}

	return &proto.GetShortURLResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         data,
	}, nil
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
