package payment

import (
	"context"
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"

	wechatNotify "github.com/fideism/golang-wechat/pay/notify"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	merchantBasic "gitlab.omytech.com.cn/micro-service/merchant-basic/pkgs/merchant-basic"
	"gitlab.omytech.com.cn/micro-service/payment/internal/cache"
	"gitlab.omytech.com.cn/micro-service/payment/internal/config"
	"gitlab.omytech.com.cn/micro-service/payment/internal/model"
)

// WechatPayment 微信支付参数
type WechatPayment struct {
	AppID        string `json:"app_id"`
	MchID        string `json:"mch_id"`
	SubMchID     string `json:"sub_mch_id"`
	PrivateKey   string `json:"private_key"`
	CertFilename string `json:"cert_filename"`
	CertContent  []byte `json:"cert_content"`
}

// CacheValue 缓存value
func (w WechatPayment) CacheValue() ([]byte, error) {
	return json.Marshal(w)
}

// CacheScan 缓存scan
func (w *WechatPayment) CacheScan(src []byte) error {
	return json.Unmarshal(src, w)
}

// GetWechatPayment 获取支付配置
func GetWechatPayment(ctx context.Context, cache *cache.Entity, appid string) (WechatPayment, error) {
	cacheKey := fmt.Sprintf("wechat_payment_%s", appid)
	item, cacheErr := cache.Get(cacheKey)
	var wechat WechatPayment
	// 读取缓存成功
	if nil == cacheErr {
		// 读取成功 映射成功
		if err := wechat.CacheScan(item); nil == err {
			return wechat, nil
		}
	}
	res, grpcErr := toGrpcWechatPayment(ctx, appid)
	util.Logger.WithSleuthContext(ctx).WithFields("grpc response", logger.Fields{
		"appid": appid,
		"res":   res,
	}).Info("获取支付配置")

	if grpcErr != nil {
		return res, fmt.Errorf("请求grpc获取支付配置错误:%s", grpcErr.Error())
	}

	cacheVal, valErr := res.CacheValue()
	if valErr != nil {
		// 内部记录日志，不对外暴露
		util.Logger.WithSleuthContext(ctx).WithFields("payment", logger.MakeFields(res)).WithError(valErr).Error("获取支付配置缓存val失败")
		return res, nil
	}

	if err := cache.Set(cacheKey, cacheVal); nil != err {
		// 内部记录日志，不对外暴露
		util.Logger.WithSleuthContext(ctx).WithFields("payment", logger.MakeFields(res)).WithError(valErr).Error("设置支付配置缓存失败")
	}

	return res, nil

}

// toGrpcWechatPayment grpc appid获取支付配置
func toGrpcWechatPayment(ctx context.Context, appid string) (WechatPayment, error) {
	var res WechatPayment
	metadata := pkgs.NewMetadataSleuthContent(fmt.Sprintf("%v", ctx.Value("sleuth_code")))
	client, clientErr := crius.NewClient(metadata, crius.ClientConfig{Address: config.Setting.Crius.Address})
	if clientErr != nil {
		return res, clientErr
	}

	server, err := client.Discover(metadata, crius.DiscoverRequest{Name: config.Setting.Crius.MerchantBasic})
	util.Logger.WithSleuthContext(ctx).WithFields("merchant basic", logger.MakeFields(server)).Info("获取商户基础服务配置")
	if err != nil {
		util.Logger.WithSleuthContext(ctx).WithFields("merchant basic", logger.MakeFields(server)).WithError(err).Info("获取商户基础服务配置")
		return res, err
	}
	basicClient, basicErr := merchantBasic.NewClient(metadata, merchantBasic.ClientConfig{Address: fmt.Sprintf("%s:%d", server.IP, server.Port)})
	if basicErr != nil {
		return res, basicErr
	}

	payment, paymentErr := basicClient.GetWechatPaymentSetting(metadata, merchantBasic.GetWechatPaymentSettingRequest{AppID: appid})
	if paymentErr != nil {
		return res, paymentErr
	}

	if payment.ErrorCode != pkgs.Success {
		return res, fmt.Errorf("请求获取门店支付配置返回错误:%s", payment.ErrorMessage)
	}

	return WechatPayment{
		AppID:      payment.Data.AppID,
		MchID:      payment.Data.MchID,
		PrivateKey: payment.Data.PrivateKey,
		//CertFilename: payment.Data.CertFilename, // 暂时不用 不处理
		//CertContent:  payment.Data.CertContent, // 暂时不用 不处理
	}, nil
}

// GetBranchWechatPayment 获取门店微信配置
func GetBranchWechatPayment(ctx context.Context, cache *cache.Entity, client *crius.Client, branchID string) (WechatPayment, error) {
	cacheKey := fmt.Sprintf("branch_wechat_payment_%s", branchID)
	item, cacheErr := cache.Get(cacheKey)
	var wechat WechatPayment

	// 读取缓存成功
	if nil == cacheErr {
		// 读取成功 映射成功
		if err := wechat.CacheScan(item); nil == err {
			return wechat, nil
		}
	}

	// 其余情况 都需要重新获取，并刷新缓存
	res, grpcErr := toGrpcBranchWechatPayment(ctx, client, branchID)
	util.Logger.WithMetadata(ctx).WithFields("grpc", logger.Fields{
		"branch_id": branchID,
		"res":       res,
	}).Info("获取门店微信支付配置")
	if grpcErr != nil {
		return res, fmt.Errorf("请求grpc获取门店支付配置失败:%s", grpcErr.Error())
	}

	cacheVal, valErr := res.CacheValue()
	if valErr != nil {
		// 内部记录日志，不对外暴露
		util.Logger.WithMetadata(ctx).WithFields("payment", logger.MakeFields(res)).WithError(valErr).Error("获取支付配置缓存val失败")
		return res, nil
	}

	if err := cache.Set(cacheKey, cacheVal); nil != err {
		// 内部记录日志，不对外暴露
		util.Logger.WithMetadata(ctx).WithFields("payment", logger.MakeFields(res)).WithError(valErr).Error("设置支付配置缓存失败")
	}

	return res, nil
}

// toGrpcBranchWechatPayment grpc获取门店支付配置
func toGrpcBranchWechatPayment(ctx context.Context, client *crius.Client, branchID string) (WechatPayment, error) {
	var res WechatPayment
	server, err := client.Discover(ctx, crius.DiscoverRequest{Name: config.Setting.Crius.MerchantBasic})
	util.Logger.WithMetadata(ctx).WithFields("merchant basic", logger.MakeFields(server)).Info("获取商户基础服务配置")
	if err != nil {
		util.Logger.WithMetadata(ctx).WithFields("merchant basic", logger.MakeFields(server)).WithError(err).Info("获取商户基础服务配置")
		return res, err
	}
	basicClient, basicErr := merchantBasic.NewClient(ctx, merchantBasic.ClientConfig{Address: fmt.Sprintf("%s:%d", server.IP, server.Port)})
	if basicErr != nil {
		return res, basicErr
	}

	payment, paymentErr := basicClient.GetBranchWechatPaymentSetting(ctx, merchantBasic.GetBranchWechatPaymentSettingRequest{BranchID: uuid.FromStringOrNil(branchID)})
	if paymentErr != nil {
		return res, paymentErr
	}

	if payment.ErrorCode != pkgs.Success {
		return res, fmt.Errorf("请求获取门店支付配置返回错误:%s", payment.ErrorMessage)
	}

	return WechatPayment{
		AppID:        payment.Data.AppID,
		MchID:        payment.Data.MchID,
		SubMchID:     payment.Data.SubMchID,
		PrivateKey:   payment.Data.PrivateKey,
		CertFilename: payment.Data.CertFilename,
		CertContent:  payment.Data.CertContent,
	}, nil
}

// WechatPayNotify 微信支付回调处理
func WechatPayNotify(ctx context.Context, response *wechatNotify.PayNotify) error {
	dbEntity, dbErr := model.DatabaseConnection()
	if dbErr != nil {
		return dbErr
	}

	var trade model.WechatTrade
	if err := dbEntity.Conn.Model(&model.WechatTrade{}).Scopes(util.ColumnEqualScope("out_trade_no", response.OutTradeNo)).First(&trade).Error; nil != err {
		return err
	}

	util.Logger.WithSleuthContext(ctx).WithFields("trade", logger.MakeFields(trade)).Info("微信支付回调,订单信息")

	if trade.TradeStatus != model.TradeStatusInit {
		util.Logger.WithSleuthContext(ctx).Info("微信支付回调,订单状态异常，跳过处理")
		return nil
	}

	if response.TotalFee != fmt.Sprintf("%d", trade.TotalFee) {
		util.Logger.WithSleuthContext(ctx).Info("微信支付回调,订单金额不匹配")
		return fmt.Errorf("订单金额不匹配")
	}

	params := pkgs.MakeParams(response)
	params.Set("wechat_trade_id", trade.ID)
	update := model.WechatTrade{
		NotifyContent: &params,
		TradeStatus:   model.TradeStatusFail,
		NotifyState:   1,
		TransactionID: response.TransactionID,
	}

	if response.ResultCode == "SUCCESS" && response.ReturnCode == "SUCCESS" {
		update.TradeStatus = model.TradeStatusSuccess
	}

	// 同步通知 下发回调失败 等待微信下次通知处理
	body, code, err := HTTPPost(ctx, trade.NotifyURL, params)
	util.Logger.WithSleuthContext(ctx).WithFields("http notify", logger.Fields{
		"body": string(body),
		"code": code,
	}).Error("微信支付回调，通知下发失败")
	if nil != err || code > 200 {
		update.NotifyState = 0
	}

	if err := dbEntity.Conn.Model(&trade).Updates(update).Error; nil != err {
		return fmt.Errorf("保存微信通知结果失败:%s", err.Error())
	}

	return nil
}

// WechatRefundNotify 退款回调
func WechatRefundNotify(ctx context.Context, response *wechatNotify.RefundNotifyDetail) error {
	dbEntity, dbErr := model.DatabaseConnection()
	if dbErr != nil {
		return dbErr
	}

	var refund model.WechatRefund
	if err := dbEntity.Conn.Model(&model.WechatRefund{}).Scopes(util.ColumnEqualScope("out_refund_no", response.OutRefundNo)).First(&refund).Error; nil != err {
		return err
	}

	util.Logger.WithSleuthContext(ctx).WithFields("refund", logger.MakeFields(refund)).Info("微信退款回调，退款信息")
	if refund.RefundStatus != model.RefundStatusInit {
		util.Logger.WithSleuthContext(ctx).Info("微信退款回调,订单状态异常，跳过处理")
		return nil
	}

	if response.SettlementRefundFee != fmt.Sprintf("%d", refund.RefundFee) {
		util.Logger.WithSleuthContext(ctx).Info("微信退款回调,订单金额不匹配")
		return fmt.Errorf("订单金额不匹配")
	}

	params := pkgs.MakeParams(response)
	params.Set("wechat_refund_id", refund.ID)
	update := model.WechatRefund{
		RefundID:      response.RefundId,
		NotifyContent: &params,
		RefundStatus:  model.RefundStatusFail,
		NotifyState:   1,
	}
	if response.RefundStatus == "SUCCESS" {
		update.RefundStatus = model.RefundStatusSuccess
	}

	// 同步通知 下发回调失败 等待微信下次通知处理
	body, code, err := HTTPPost(ctx, refund.NotifyURL, params)
	util.Logger.WithSleuthContext(ctx).WithFields("http notify", logger.Fields{
		"body": string(body),
		"code": code,
	}).Error("微信退款回调，通知下发失败")
	if nil != err || code > 200 {
		update.NotifyState = 0
	}

	if err := dbEntity.Conn.Model(&refund).Updates(update).Error; nil != err {
		return fmt.Errorf("保存微信退款通知结果失败:%s", err.Error())
	}

	return nil
}
