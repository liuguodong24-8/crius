package rpc

import (
	"context"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/message/internal/cache"
	"gitlab.omytech.com.cn/micro-service/message/proto"
)

// ListOfficialLink 公众号链接列表
func (s *Server) ListOfficialLink(ctx context.Context, req *proto.Empty) (*proto.ListOfficialLinkResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).Info("ListOfficialLink")

	content, err := cache.GetWechatLink(ctx, s.cache)
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error(err.Error())
		return &proto.ListOfficialLinkResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: "读取配置文件错误",
		}, nil
	}

	var links []*proto.ListOfficialLinkResponse_OfficialLink

	for _, v := range content.Official {
		links = append(links, &proto.ListOfficialLinkResponse_OfficialLink{
			Url:  v.URL,
			Name: v.Name,
		})
	}

	return &proto.ListOfficialLinkResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         links,
	}, nil
}

// ListMiniprogramLink 小程序链接地址列表
func (s *Server) ListMiniprogramLink(ctx context.Context, req *proto.Empty) (*proto.ListMiniprogramLinkResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).Info("ListMiniprogramLink")
	content, err := cache.GetWechatLink(ctx, s.cache)
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error(err.Error())
		return &proto.ListMiniprogramLinkResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: "读取配置文件错误",
		}, nil
	}

	var links []*proto.ListMiniprogramLinkResponse_MiniprogramLink

	for _, v := range content.Miniprogram {
		links = append(links, &proto.ListMiniprogramLinkResponse_MiniprogramLink{
			Url:  v.URL,
			Name: v.Name,
		})
	}

	return &proto.ListMiniprogramLinkResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         links,
	}, nil
}
