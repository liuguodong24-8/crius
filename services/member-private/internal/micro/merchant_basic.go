package micro

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/member-private/internal/config"
	merchantBasicProto "gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"google.golang.org/grpc"
)

// MerchantBasicClient 微信client
type MerchantBasicClient struct {
	Entity merchantBasicProto.MerchantBasicServiceClient
}

// GetMerchantBasicServer 获取微信client
func GetMerchantBasicServer(ctx context.Context, client *crius.Client) (*MerchantBasicClient, error) {
	server, err := client.Discover(ctx, crius.DiscoverRequest{Name: config.Setting.Crius.MerchantBasic})
	if err != nil {
		return nil, fmt.Errorf("获取MerchantBasic server失败:%s", err.Error())
	}

	util.Logger.WithMetadata(ctx).WithFields("MerchantBasic server", logger.MakeFields(server)).Info("发现MerchantBasic服务")

	conn, connErr := grpc.Dial(fmt.Sprintf("%s:%d", server.IP, server.Port), grpc.WithInsecure())
	if connErr != nil {
		return nil, fmt.Errorf("实例化MerchantBasic服务失败:%s", connErr.Error())
	}

	return &MerchantBasicClient{Entity: merchantBasicProto.NewMerchantBasicServiceClient(conn)}, nil
}

// GetBranchLatelyBusinessResponse 获取门店最近一次营业日信息
type GetBranchLatelyBusinessResponse struct {
	BusinessDate  string
	BeginTime     string
	EndTime       string
	IsNextDay     bool
	BusinessBegin *time.Time
	BusinessEnd   *time.Time
}

// GetBranchLatelyBusiness 获取门店最近一次营业日
func (client *MerchantBasicClient) GetBranchLatelyBusiness(ctx context.Context, branchID string, dateTime int64) (*GetBranchLatelyBusinessResponse, error) {
	res, err := client.Entity.GetBranchLatelyBusiness(ctx, &merchantBasicProto.GetBranchLatelyBusinessRequest{
		BranchId: branchID,
		DateTime: dateTime,
	})

	util.Logger.WithMetadata(ctx).WithFields("GetBranchLatelyBusiness", logger.MakeFields(res)).Info("请求获取门店营业日")

	if err != nil {
		return nil, fmt.Errorf("请求获取门店营业信息错误:%s", err.Error())
	}

	if res.ErrorCode != pkgs.Success {
		return nil, fmt.Errorf("请求获取门店营业信息错误:code:%d, %s", res.ErrorCode, res.ErrorMessage)
	}

	if res.Data == nil {
		return nil, errors.New("请求获取门店营业信息错误, data nil")
	}

	response := &GetBranchLatelyBusinessResponse{
		BusinessDate: res.Data.BusinessDate[:10],
		BeginTime:    res.Data.BeginTime[:8],
		EndTime:      res.Data.EndTime[:8],
		IsNextDay:    res.Data.IsNextDay,
	}

	layout := `2006-01-02 15:04:05`

	if begin, err := time.ParseInLocation(layout, fmt.Sprintf("%s %s", response.BusinessDate, response.BeginTime), time.Local); nil == err {
		response.BusinessBegin = &begin
	}

	if end, err := time.ParseInLocation(layout, fmt.Sprintf("%s %s", response.BusinessDate, response.EndTime), time.Local); nil == err {
		response.BusinessEnd = &end
	}

	return response, nil
}
