package rpc

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	wechatPay "github.com/fideism/golang-wechat/pay"
	"github.com/fideism/golang-wechat/pay/base"
	wechatPayConfig "github.com/fideism/golang-wechat/pay/config"
	wechatUtil "github.com/fideism/golang-wechat/util"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/payment/internal/config"
	"gitlab.omytech.com.cn/micro-service/payment/internal/model"
	"gitlab.omytech.com.cn/micro-service/payment/internal/payment"
	"gitlab.omytech.com.cn/micro-service/payment/proto"
	"gorm.io/gorm"
)

// WechatUnifiedOrder 统一下单
func (s *Server) WechatUnifiedOrder(ctx context.Context, req *proto.WechatUnifiedOrderRequest) (*proto.WechatUnifiedOrderResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("UnifiedOrder")
	if len(req.BranchId) == 0 || len(req.Int64Map) == 0 || len(req.StringMap) == 0 {
		util.Logger.WithMetadata(ctx).Error("微信下单，参数错误")
		return &proto.WechatUnifiedOrderResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	wechatPayment, err := payment.GetBranchWechatPayment(ctx, s.cache, s.crius, req.BranchId)
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("微信下单,获取微信支付配置失败")
		return &proto.WechatUnifiedOrderResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("获取微信配置失败:%s", err.Error()),
		}, nil
	}

	wechat, _ := getWechat(wechatPayment)
	params := toUnifiedOrderParams(req)
	params.Set("sub_mch_id", wechatPayment.SubMchID)
	if !params.Exists("total_fee") || !params.Exists("trade_type") {
		util.Logger.WithMetadata(ctx).Error("微信下单，参数错误")
		return &proto.WechatUnifiedOrderResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}
	metadata := pkgs.GetMetadata(ctx)
	totalFee, _ := strconv.Atoi(params.GetString("total_fee"))
	requestParams := pkgs.MakeParams(req)
	trade := model.WechatTrade{
		ID:            uuid.NewV4(),
		MerchantID:    metadata.MerchantID,
		BranchID:      uuid.FromStringOrNil(req.BranchId),
		TotalFee:      totalFee,
		TradeType:     model.TradeType(params.GetString("trade_type")),
		NotifyURL:     params.GetString("notify_url"),
		RequestParams: &requestParams,
		TradeStatus:   model.TradeStatusInit,
		NotifyState:   0,
	}
	if !params.Exists("out_trade_no") {

		no := strings.Replace(trade.ID.String(), "-", "", -1)
		trade.OutTradeNo = no
		params.Set("out_trade_no", no)
	} else {
		trade.OutTradeNo = params.GetString("out_trade_no")
	}

	// 重新设置 服务 通知回调地址
	params.Set("notify_url", fmt.Sprintf("%s/api/v1/wechat/%s/pay/notify", config.Setting.Web.Domain, wechatPayment.AppID))
	wechatRequest := pkgs.MakeParams(params)
	trade.WechatRequest = &wechatRequest

	util.Logger.WithFields("params", logger.MakeFields(params)).Info("debug")
	payResponse, payErr := wechat.GetOrder().Unify(params)
	wechatResponse := pkgs.MakeParams(payResponse)

	trade.WechatResponse = &wechatResponse

	if err := s.database.Conn.Create(&trade).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("wechat trade", logger.MakeFields(trade)).WithError(err).Error("微信下单,保存微信交易信息失败")
		return &proto.WechatUnifiedOrderResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("保存微信下单记录失败:%s", err.Error()),
		}, nil
	}
	util.Logger.WithMetadata(ctx).WithFields("wechat trade", logger.MakeFields(trade)).Info("微信下单,微信交易信息")

	if payErr != nil {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(params)).WithError(payErr).Error("微信下单,微信下单失败")
		return &proto.WechatUnifiedOrderResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("微信下单失败:%s", payErr.Error()),
		}, nil
	}

	tradeType := payResponse.Data.GetString("trade_type")
	if tradeType == "JSAPI" {
		payResponse = makeJsApiPaySign(wechatPayment, payResponse)
	}

	return &proto.WechatUnifiedOrderResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.WechatUnifiedOrderResponseData{
			WechatTradeId:  trade.ID.String(),
			WechatResponse: wechatUtilParamsToWechatResponse(payResponse),
		},
	}, nil
}

func makeJsApiPaySign(cfg payment.WechatPayment, res *base.Response) *base.Response {
	response := res
	prepay := fmt.Sprintf("prepay_id=%s", res.Data.GetString("prepay_id"))
	t := time.Now().Unix()
	response.Data.Set("timestamp", t)
	response.Data.Set("sign_type", "MD5")

	// 重新加密生成sdk pay sign
	response.Data.Set("pay_sign", base.SignParamsMD5(cfg.PrivateKey, wechatUtil.Params{
		"appId":     cfg.AppID,
		"timeStamp": t,
		"nonceStr":  res.Data.GetString("nonce_str"),
		"package":   prepay,
		"signType":  "MD5",
	}))

	response.Data.Set("package", prepay)

	return response
}

func toUnifiedOrderParams(req *proto.WechatUnifiedOrderRequest) wechatUtil.Params {
	params := wechatUtil.Params{}

	for k, v := range req.StringMap {
		params.Set(k, v)
	}

	for k, v := range req.Int64Map {
		params.Set(k, v)
	}

	return params
}

func toRefundParams(req *proto.WechatRefundRequest) wechatUtil.Params {
	params := wechatUtil.Params{}

	for k, v := range req.StringMap {
		params.Set(k, v)
	}

	for k, v := range req.Int64Map {
		params.Set(k, v)
	}

	return params
}

func getWechat(cfg payment.WechatPayment) (*wechatPay.Pay, wechatPayConfig.Cert) {
	return wechatPay.NewPay(&wechatPayConfig.Config{
			Sandbox: false,
			AppID:   cfg.AppID,
			MchID:   cfg.MchID,
			Key:     cfg.PrivateKey,
		}), wechatPayConfig.Cert{
			Content: cfg.CertContent,
		}
}

// WechatOrderQuery 订单查询
func (s *Server) WechatOrderQuery(ctx context.Context, req *proto.WechatOrderQueryRequest) (*proto.WechatOrderQueryResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("WechatOrderQuery")
	if len(req.WechatTradeId) == 0 {
		util.Logger.WithMetadata(ctx).Error("微信订单查询，参数错误")
		return &proto.WechatOrderQueryResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	branchWechatTrade := s.getBranchWechatTrade(ctx, req.WechatTradeId)
	if branchWechatTrade.Error != nil {
		util.Logger.WithMetadata(ctx).Error(branchWechatTrade.Error.Error())
		return &proto.WechatOrderQueryResponse{
			ErrorCode:    branchWechatTrade.ErrorCode,
			ErrorMessage: branchWechatTrade.ErrorMessage,
		}, nil
	}

	wechat, _ := getWechat(branchWechatTrade.WechatPayment)

	wechatResponse, wechatErr := wechat.GetOrder().Query(wechatUtil.Params{
		"out_trade_no": branchWechatTrade.Trade.OutTradeNo,
		"sub_mch_id":   branchWechatTrade.WechatPayment.SubMchID,
	})

	queryResponse := pkgs.MakeParams(wechatResponse)
	if err := s.database.Conn.Model(&branchWechatTrade.Trade).Updates(model.WechatTrade{
		QueryContent: &queryResponse,
	}).Error; nil != err {
		//更新失败，只记录日志
		util.Logger.WithMetadata(ctx).WithError(err).Error("微信订单查询，保存查询结果失败")
	}

	if wechatErr != nil {
		util.Logger.WithMetadata(ctx).WithError(wechatErr).Error("微信订单查询,请求微信错误")
		return &proto.WechatOrderQueryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("微信订单查询:%s", wechatErr.Error()),
		}, nil
	}

	data := wechatUtilParamsToWechatResponse(wechatResponse)
	util.Logger.WithMetadata(ctx).WithFields("wechat response", logger.MakeFields(data)).Info("WechatOrderQuery")
	return &proto.WechatOrderQueryResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         data,
	}, nil
}

// WechatCloseOrder 订单关闭
func (s *Server) WechatCloseOrder(ctx context.Context, req *proto.WechatCloseOrderRequest) (*proto.WechatCloseOrderResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("WechatCloseOrder")
	if len(req.WechatTradeId) == 0 || len(req.BranchId) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Error("微信关闭订单，参数错误")
		return &proto.WechatCloseOrderResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	branchWechatTrade := s.getBranchWechatTrade(ctx, req.WechatTradeId)
	if branchWechatTrade.Error != nil {
		util.Logger.WithMetadata(ctx).Error(branchWechatTrade.Error.Error())
		return &proto.WechatCloseOrderResponse{
			ErrorCode:    branchWechatTrade.ErrorCode,
			ErrorMessage: branchWechatTrade.ErrorMessage,
		}, nil
	}

	wechat, _ := getWechat(branchWechatTrade.WechatPayment)

	if branchWechatTrade.Trade.TradeStatus != model.TradeStatusInit {
		util.Logger.WithMetadata(ctx).Error("微信关闭订单，订单状态不允许操作")
		return &proto.WechatCloseOrderResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "订单状态错误",
		}, nil
	}

	wechatResponse, wechatErr := wechat.GetOrder().Close(wechatUtil.Params{
		"out_trade_no": branchWechatTrade.Trade.OutTradeNo,
		"sub_mch_id":   branchWechatTrade.WechatPayment.SubMchID,
	})

	if wechatErr != nil {
		util.Logger.WithMetadata(ctx).WithError(wechatErr).Error("微信关闭订单,请求微信错误")
		return &proto.WechatCloseOrderResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("微信关闭订单:%s", wechatErr.Error()),
		}, nil
	}

	data := wechatUtilParamsToWechatResponse(wechatResponse)
	util.Logger.WithMetadata(ctx).WithFields("wechat response", logger.MakeFields(data)).Info("WechatCloseOrder")
	trade := branchWechatTrade.Trade
	if data.ReturnCode == `SUCCESS` && data.ReturnMsg == `OK` {
		if err := s.database.Conn.Model(&trade).Updates(model.WechatTrade{
			TradeStatus: model.TradeStatusClose,
		}).Error; nil != err {
			return &proto.WechatCloseOrderResponse{
				ErrorCode:    pkgs.ErrInternal,
				ErrorMessage: fmt.Sprintf("微信关闭订单,修改订单状态错误:%s", err.Error()),
			}, nil
		}
	}

	return &proto.WechatCloseOrderResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         data,
	}, nil
}

// WechatRefund 订单退款
func (s *Server) WechatRefund(ctx context.Context, req *proto.WechatRefundRequest) (*proto.WechatRefundResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("WechatRefund")
	if len(req.WechatTradeId) == 0 || len(req.BranchId) == 0 {
		util.Logger.WithMetadata(ctx).Error("微信订单退款，参数错误")
		return &proto.WechatRefundResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	branchWechatTrade := s.getBranchWechatTrade(ctx, req.WechatTradeId)
	if branchWechatTrade.Error != nil {
		util.Logger.WithMetadata(ctx).Error(branchWechatTrade.Error.Error())
		return &proto.WechatRefundResponse{
			ErrorCode:    branchWechatTrade.ErrorCode,
			ErrorMessage: branchWechatTrade.ErrorMessage,
		}, nil
	}

	// 组装退款数据
	params := toRefundParams(req)
	params.Set("sub_mch_id", branchWechatTrade.WechatPayment.SubMchID)
	params.Set("out_trade_no", branchWechatTrade.Trade.OutTradeNo)
	params.Set("total_fee", branchWechatTrade.Trade.TotalFee)
	var refundFee int
	if !params.Exists("refund_fee") {
		refundFee = branchWechatTrade.Trade.TotalFee
	} else {
		fee, _ := strconv.Atoi(params.GetString("refund_fee"))
		refundFee = fee
	}

	metadata := pkgs.GetMetadata(ctx)
	requestParams := pkgs.MakeParams(req)
	refund := model.WechatRefund{
		ID:            uuid.NewV4(),
		MerchantID:    metadata.MerchantID,
		BranchID:      uuid.FromStringOrNil(req.BranchId),
		WechatTradeID: branchWechatTrade.Trade.ID,
		RefundFee:     refundFee,
		NotifyURL:     params.GetString("notify_url"),
		RequestParams: &requestParams,
		RefundStatus:  model.RefundStatusInit,
	}
	if !params.Exists("out_refund_no") {
		no := strings.Replace(refund.ID.String(), "-", "", -1)
		refund.OutRefundNo = no
		params.Set("out_refund_no", no)
	} else {
		refund.OutRefundNo = params.GetString("out_refund_no")
	}
	// 重新设置 服务 通知回调地址
	params.Set("notify_url", fmt.Sprintf("%s/api/v1/wechat/%s/refund/notify", config.Setting.Web.Domain, branchWechatTrade.WechatPayment.AppID))
	wechatRequest := pkgs.MakeParams(params)
	refund.WechatRequest = &wechatRequest

	wechat, cert := getWechat(branchWechatTrade.WechatPayment)
	refundResponse, refundErr := wechat.GetRefund().Refund(params, cert)
	wechatResponse := pkgs.MakeParams(refundResponse)
	refund.WechatResponse = &wechatResponse

	if err := s.database.Conn.Create(&refund).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("wechat refund", logger.MakeFields(refund)).WithError(err).Error("微信订单退款,保存微信退款信息失败")
		return &proto.WechatRefundResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("保存微信退款记录失败:%s", err.Error()),
		}, nil
	}

	util.Logger.WithMetadata(ctx).WithFields("wechat refund", logger.MakeFields(refund)).Info("微信订单退款,微信退款信息")

	if refundErr != nil {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(params)).WithError(refundErr).Error("微信订单退款,微信退款失败")
		return &proto.WechatRefundResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("微信订单退款:%s", refundErr.Error()),
		}, nil
	}

	return &proto.WechatRefundResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data: &proto.WechatRefundResponseData{
			WechatRefundId: refund.ID.String(),
			WechatResponse: wechatUtilParamsToWechatResponse(refundResponse),
		},
	}, nil
}

// WechatRefundQuery 退款查询
func (s *Server) WechatRefundQuery(ctx context.Context, req *proto.WechatRefundQueryRequest) (*proto.WechatRefundQueryResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("request", logger.MakeFields(req)).Info("WechatRefundQuery")
	if len(req.WechatRefundId) == 0 || len(req.BranchId) == 0 {
		util.Logger.WithMetadata(ctx).Error("微信退款查询，参数错误")
		return &proto.WechatRefundQueryResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	wechatPayment, err := payment.GetBranchWechatPayment(ctx, s.cache, s.crius, req.BranchId)
	if err != nil {
		util.Logger.WithMetadata(ctx).WithError(err).Error("微信退款查询,获取微信支付配置失败")
		return &proto.WechatRefundQueryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("获取微信配置失败:%s", err.Error()),
		}, nil
	}

	var refund model.WechatRefund
	if err := s.database.Conn.Model(&model.WechatRefund{}).Scopes(util.ColumnEqualScope("id", req.WechatRefundId)).First(&refund).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("微信退款查询,数据库查询错误")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &proto.WechatRefundQueryResponse{
				ErrorCode:    pkgs.ErrNotFound,
				ErrorMessage: "查询对应退款信息错误",
			}, nil
		}

		return &proto.WechatRefundQueryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: "数据库查询错误",
		}, nil
	}

	wechat, _ := getWechat(wechatPayment)

	wechatResponse, wechatErr := wechat.GetRefund().Query(wechatUtil.Params{
		"out_refund_no": refund.OutRefundNo,
		"sub_mch_id":    wechatPayment.SubMchID,
	})

	queryResponse := pkgs.MakeParams(wechatResponse)
	if err := s.database.Conn.Model(&refund).Updates(model.WechatRefund{
		QueryContent: &queryResponse,
	}).Error; nil != err {
		//更新失败，只记录日志
		util.Logger.WithMetadata(ctx).WithError(err).Error("微信退款查询，保存查询结果失败")
	}

	if wechatErr != nil {
		util.Logger.WithMetadata(ctx).WithError(wechatErr).Error("微信退款查询,请求微信错误")
		return &proto.WechatRefundQueryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("微信退款查询:%s", wechatErr.Error()),
		}, nil
	}

	data := wechatUtilParamsToWechatResponse(wechatResponse)
	util.Logger.WithMetadata(ctx).WithFields("wechat response", logger.MakeFields(data)).Info("WechatRefundQuery")
	return &proto.WechatRefundQueryResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         data,
	}, nil
}

func wechatUtilParamsToWechatResponse(response *base.Response) *proto.WechatResponse {
	if response == nil {
		return nil
	}

	item := map[string]string{}
	for k, v := range response.Data {
		item[k] = wechatUtil.InterfaceToString(v)
	}

	return &proto.WechatResponse{
		ReturnCode: response.ReturnCode,
		ReturnMsg:  response.ReturnMsg,
		Detail:     response.Detail,
		Data:       item,
	}
}

type branchWechat struct {
	WechatPayment payment.WechatPayment
	Trade         model.WechatTrade
	ErrorCode     int32
	ErrorMessage  string
	Error         error
}

// getBranchWechatTrade 统一获取门店支付配置 交易配置
func (s *Server) getBranchWechatTrade(ctx context.Context, tradeID string) branchWechat {
	res := branchWechat{}
	// 查询交易信息
	var trade model.WechatTrade
	if err := s.database.Conn.Model(&model.WechatTrade{}).Where("id = ?", tradeID).First(&trade).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.Error = fmt.Errorf("微信订单退款，数据库查询错误")
			res.ErrorCode = pkgs.ErrNotFound
			res.ErrorMessage = "查询对应订单信息错误"
			return res
		}

		res.Error = fmt.Errorf("查询数据库错误:%s", err.Error())
		res.ErrorCode = pkgs.ErrUnprocessableEntity
		res.ErrorMessage = "查询对应订单信息错误"

		return res
	}

	res.Trade = trade

	// 查询门店微信配置信息
	wechatPayment, err := payment.GetBranchWechatPayment(ctx, s.cache, s.crius, trade.BranchID.String())
	if err != nil {
		res.Error = fmt.Errorf("获取微信支付配置失败:%s", err.Error())
		res.ErrorCode = pkgs.ErrInternal
		res.ErrorMessage = "获取微信支付配置失败"

		return res
	}

	res.WechatPayment = wechatPayment
	res.Error = nil

	return res
}
