package rpc

import (
	"context"
	"fmt"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/wechat/internal/service"
	"gitlab.omytech.com.cn/micro-service/wechat/proto"
)

// OfficialGetAuthURL 公众号 获取网页授权URL
func (s *Server) OfficialGetAuthURL(ctx context.Context, req *proto.OfficialAuthURLRequest) (*proto.OfficialAuthURLResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("OfficialGetAuthURL")

	official, _, officialErr := service.NewOfficial(req.Channel)
	if officialErr != nil {
		util.Logger.WithMetadata(ctx).WithError(officialErr).Error("获取公众号配置错误")
		return &proto.OfficialAuthURLResponse{
			ErrorCode:    pkgs.ErrNotFound,
			ErrorMessage: fmt.Sprintf("获取公众号配置错误:%s", officialErr.Error()),
		}, nil
	}

	auth := official.GetOauth()

	redirect, err := auth.GetRedirectURL(req.Url, req.Scope, req.State)
	util.Logger.WithMetadata(ctx).WithFields("wechat response", logger.Fields{
		"redirect": redirect,
	}).Info("微信返回")
	if err != nil {
		return &proto.OfficialAuthURLResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("请求微信失败:%s", err.Error()),
		}, nil
	}

	return &proto.OfficialAuthURLResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.OfficialAuthURLResponse_Data{
			Url: redirect,
		},
	}, nil
}

// OfficialJsSdk 公众号sdk
func (s *Server) OfficialJsSdk(ctx context.Context, req *proto.OfficialJsSdkRequest) (*proto.OfficialJsSdkResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("OfficialJsSdk")

	official, officialConfig, officialErr := service.NewOfficial(req.Channel)
	if officialErr != nil {
		util.Logger.WithMetadata(ctx).WithError(officialErr).Error("获取公众号配置错误")
		return &proto.OfficialJsSdkResponse{
			ErrorCode:    pkgs.ErrNotFound,
			ErrorMessage: fmt.Sprintf("获取公众号配置错误:%s", officialErr.Error()),
		}, nil
	}

	js := official.GetJs()
	response, err := js.GetConfig(req.Url)
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("获取公众号jssdk错误")
		return &proto.OfficialJsSdkResponse{
			ErrorCode:    pkgs.ErrNotFound,
			ErrorMessage: fmt.Sprintf("获取公众号jssdk错误:%s", err.Error()),
		}, nil
	}
	return &proto.OfficialJsSdkResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.OfficialJsSdkResponse_Data{
			Appid:     officialConfig.AppID,
			NonceStr:  response.NonceStr,
			Timestamp: response.Timestamp,
			Signature: response.Signature,
		},
	}, nil
}

// OfficialCodeGetUser 公众号 code换取user信息
func (s *Server) OfficialCodeGetUser(ctx context.Context, req *proto.OfficialCodeGetUserRequest) (*proto.OfficialCodeGetUserResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("OfficialCodeGetUser")

	official, officialConfig, officialErr := service.NewOfficial(req.Channel)
	if officialErr != nil {
		util.Logger.WithMetadata(ctx).WithError(officialErr).Error("获取公众号配置错误")
		return &proto.OfficialCodeGetUserResponse{
			ErrorCode:    pkgs.ErrNotFound,
			ErrorMessage: fmt.Sprintf("获取公众号配置错误:%s", officialErr.Error()),
		}, nil
	}

	auth := official.GetOauth()

	token, tokenErr := auth.GetUserAccessToken(req.Code)
	util.Logger.WithMetadata(ctx).WithFields("wechat token response", logger.MakeFields(token)).Info("微信返回")
	if tokenErr != nil {
		util.Logger.WithMetadata(ctx).Error(fmt.Sprintf("wechat token response err:%s", tokenErr.Error()))
		return &proto.OfficialCodeGetUserResponse{
			ErrorCode:    pkgs.ErrNotFound,
			ErrorMessage: fmt.Sprintf("请求微信token错误:%s", tokenErr.Error()),
		}, nil
	}

	if token.ErrCode != service.ErrcodeSuccess || len(token.AccessToken) == 0 {
		return &proto.OfficialCodeGetUserResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("error_code:%d, error_message:%s", token.ErrCode, token.ErrMsg),
		}, nil
	}

	// 暂时不做  refresh_token 刷新
	_ = auth.Cache.Set(openidCacheKey(token.OpenID), token.AccessToken, time.Hour*2)

	return &proto.OfficialCodeGetUserResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.OfficialCodeGetUserResponse_BaseWechatUser{
			Appid:  officialConfig.AppID,
			Openid: token.OpenID,
		},
	}, nil
}

// OfficialOpenidGetUser 公众号 openid获取详细用户信息
func (s *Server) OfficialOpenidGetUser(ctx context.Context, req *proto.OfficialOpenidGetUserRequest) (*proto.OfficialOpenidGetUserResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("OfficialCodeGetUser")

	official, officialConfig, officialErr := service.NewOfficial(req.Channel)
	if officialErr != nil {
		util.Logger.WithMetadata(ctx).WithError(officialErr).Error("获取公众号配置错误")
		return &proto.OfficialOpenidGetUserResponse{
			ErrorCode:    pkgs.ErrNotFound,
			ErrorMessage: fmt.Sprintf("获取公众号配置错误:%s", officialErr.Error()),
		}, nil
	}

	auth := official.GetOauth()

	accessToken := auth.Cache.Get(openidCacheKey(req.Openid))
	// 过期，暂时不刷新，直接返回
	// 后续考虑通过refresh_token刷新
	if accessToken == nil {
		return &proto.OfficialOpenidGetUserResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: "缓存token错误",
		}, nil
	}

	res, err := auth.GetUserInfo(accessToken.(string), req.Openid)
	util.Logger.WithMetadata(ctx).WithFields("wechat user response", logger.MakeFields(res)).Info("微信返回")

	if err != nil || res.ErrCode != service.ErrcodeSuccess {
		var message string
		if err != nil {
			message = fmt.Sprintf("请求微信失败:%s", err.Error())
		} else {
			message = fmt.Sprintf("error_code:%d, error_message:%s", res.ErrCode, res.ErrMsg)
		}
		return &proto.OfficialOpenidGetUserResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: message,
		}, nil
	}

	// "errcode":40163,"errmsg":"code been used, hints
	// errcode":48001,"errmsg":"api unauthorized

	return &proto.OfficialOpenidGetUserResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.WechatUser{
			Appid:      officialConfig.AppID,
			Follow:     true,
			Openid:     res.OpenID,
			Nickname:   res.Nickname,
			Sex:        res.Sex,
			Province:   res.Province,
			City:       res.City,
			Headimgurl: res.HeadImgURL,
			Privilege:  res.Privilege,
			Unionid:    res.Unionid,
		},
	}, nil
}

func openidCacheKey(src string) string {
	return fmt.Sprintf("wechat-server-openid:%s", src)
}
