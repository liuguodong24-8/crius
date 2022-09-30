package rpc

import (
	"context"
	"fmt"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	cutil "gitlab.omytech.com.cn/micro-service/Crius/util"
	basicProto "gitlab.omytech.com.cn/micro-service/basic/proto"
	"gitlab.omytech.com.cn/micro-service/member-account/internal/config"
	"gitlab.omytech.com.cn/micro-service/member-account/proto"
	memberExtension "gitlab.omytech.com.cn/micro-service/member-extension/proto"
	private "gitlab.omytech.com.cn/micro-service/member-private/proto"
	merchantBasic "gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"google.golang.org/grpc"
)

// Server server
type Server struct {
	proto.UnimplementedMemberAccountServerServer
	// 服务发现注册
	crius *crius.Client
}

var (
	// member-private服务
	memberPrivateClient private.MemberPrivateServerClient
	// merchant-basic服务
	merchantBasicClient   merchantBasic.MerchantBasicServiceClient
	basicClient           basicProto.BasicServiceClient
	memberExtensionClient memberExtension.ExtensionServerClient
)

// NewMemberAccountServe 实例化
func NewMemberAccountServe() (*Server, error) {

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
	//cutil.Logger.WithFields("config", logger.MakeFields(cfg)).Info("basic 心跳")
	return s.crius.Heart(ctx, crius.HeartRequest{
		Name:   cfg.Name,
		Desc:   cfg.Desc,
		IP:     cfg.IP,
		Port:   cfg.Port,
		Weight: cfg.Weight,
	})
}

func (s *Server) memberPrivate() private.MemberPrivateServerClient {
	if memberPrivateClient == nil {
		resp, err := s.crius.Discover(context.Background(), crius.DiscoverRequest{Name: config.Setting.Crius.MemberPrivate})
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("member_private 服务发现错误err:%v, resp:%v", err, resp))
			return nil
		}
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", resp.IP, resp.Port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("member_private rpc连接错误err:%v, ip:port:%v", err, fmt.Sprintf("%s:%d", resp.IP, resp.Port)))
			return nil
		}
		memberPrivateClient = private.NewMemberPrivateServerClient(conn)
	}
	return memberPrivateClient
}

func (s *Server) merchantBasic() merchantBasic.MerchantBasicServiceClient {
	if merchantBasicClient == nil {
		resp, err := s.crius.Discover(context.Background(), crius.DiscoverRequest{Name: config.Setting.Crius.MerchantBasic})
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("merchantBasic 服务发现错误err:%v, resp:%v", err, resp))
			return nil
		}
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", resp.IP, resp.Port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("merchantBasic rpc连接错误err:%v, ip:port:%v", err, fmt.Sprintf("%s:%d", resp.IP, resp.Port)))
			return nil
		}
		merchantBasicClient = merchantBasic.NewMerchantBasicServiceClient(conn)
	}
	return merchantBasicClient
}

func (s *Server) memberExtension() memberExtension.ExtensionServerClient {
	if memberExtensionClient == nil {
		resp, err := s.crius.Discover(context.Background(), crius.DiscoverRequest{Name: config.Setting.Crius.MemberExtension})
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("memberExtension 服务发现错误err:%v, resp:%v", err, resp))
			return nil
		}
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", resp.IP, resp.Port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("memberExtension rpc连接错误err:%v, ip:port:%v", err, fmt.Sprintf("%s:%d", resp.IP, resp.Port)))
			return nil
		}
		memberExtensionClient = memberExtension.NewExtensionServerClient(conn)
	}
	return memberExtensionClient
}
func (s *Server) basic() basicProto.BasicServiceClient {
	if basicClient == nil {
		resp, err := s.crius.Discover(context.Background(), crius.DiscoverRequest{Name: config.Setting.Crius.Basic})
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("basic 服务发现错误err:%v, resp:%v", err, resp))
			return nil
		}

		conn, connErr := grpc.Dial(fmt.Sprintf("%s:%d", resp.IP, resp.Port), grpc.WithInsecure())
		if connErr != nil {
			cutil.Logger.Error(fmt.Sprintf("basic rpc连接错误err:%v, ip:port:%v", err, fmt.Sprintf("%s:%d", resp.IP, resp.Port)))
			return nil
		}

		basicClient = basicProto.NewBasicServiceClient(conn)
	}
	return basicClient
}
