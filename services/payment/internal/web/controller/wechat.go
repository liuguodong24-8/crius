package controller

import (
	"context"
	"io/ioutil"
	"net/http"

	wechatPay "github.com/fideism/golang-wechat/pay"
	wechatPayConfig "github.com/fideism/golang-wechat/pay/config"
	"github.com/gin-gonic/gin"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/payment/internal/payment"
)

// WechatPayNotify 微信支付回调
func WechatPayNotify(c *gin.Context) {
	ctx := pkgs.SleuthContext()
	appid := c.Param("appid")
	body, err := ioutil.ReadAll(c.Request.Body)

	util.Logger.WithSleuthContext(ctx).WithFields("request", logger.Fields{
		"appid": appid,
		"body":  string(body),
	}).Info("微信支付回调, begin")

	defer c.Request.Body.Close()
	if err != nil {
		util.Logger.WithSleuthContext(ctx).WithError(err).Error("微信支付回调，解析失败")
		c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": "body error"})
		return
	}

	pay := getWechatPay(appid)
	response, responseErr := pay.GetNotify().AnalysisPayNotify(body)
	if nil != responseErr {
		util.Logger.WithSleuthContext(ctx).WithError(responseErr).Error("微信支付回调，xml解析失败")
		c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": "parse error"})
		return
	}
	util.Logger.WithSleuthContext(ctx).WithFields("wechat notify", logger.MakeFields(response)).Info("微信支付回调, 解析信息")

	if err := payment.WechatPayNotify(ctx, response); nil != err {
		util.Logger.WithSleuthContext(ctx).WithError(err).Error("微信支付回调，处理回调失败")
		c.XML(http.StatusOK, gin.H{"return_code": "SUCCESS", "return_msg": "deal error"})
		return
	}

	c.XML(http.StatusOK, gin.H{"return_code": "SUCCESS", "return_msg": "OK"})
}

// WechatRefundNotify 微信退款回调
func WechatRefundNotify(c *gin.Context) {
	ctx := pkgs.SleuthContext()
	appid := c.Param("appid")
	body, err := ioutil.ReadAll(c.Request.Body)
	util.Logger.WithSleuthContext(ctx).WithFields("request", logger.Fields{
		"appid": appid,
		"body":  string(body),
	}).Info("微信退款回调, begin")

	defer c.Request.Body.Close()
	if err != nil {
		util.Logger.WithSleuthContext(ctx).WithError(err).Error("微信退款回调，解析失败")
		c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": "body error"})
		return
	}

	pay, wechatErr := getWechatRefund(ctx, appid)
	if wechatErr != nil {
		util.Logger.WithSleuthContext(ctx).WithError(err).Error("微信退款回调，获取微信支付配置")
		c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": "grpc wechat error"})
		return
	}
	refundResponse, refundError := pay.GetNotify().AnalysisRefundNotify(body)
	if refundError != nil {
		util.Logger.WithSleuthContext(ctx).WithError(refundError).Error("微信退款回调,xml解析失败")
		c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": "parse error"})
		return
	}

	util.Logger.WithSleuthContext(ctx).WithFields("wechat refund notify", logger.MakeFields(refundResponse)).Info("微信退款回调, 解析结果")

	if err := payment.WechatRefundNotify(ctx, &refundResponse.Detail); nil != err {
		util.Logger.WithSleuthContext(ctx).WithError(err).Error("微信退款回调，处理回调失败")
		c.XML(http.StatusOK, gin.H{"return_code": "SUCCESS", "return_msg": "deal error"})
		return
	}

	c.XML(http.StatusOK, gin.H{"return_code": "FAIL", "return_msg": "test error"})
}

func getWechatPay(appid string) *wechatPay.Pay {
	// todo 是否查询对应信息
	return wechatPay.NewPay(&wechatPayConfig.Config{
		Sandbox: false,
		AppID:   appid,
		MchID:   "",
		Key:     "",
	})
}

func getWechatRefund(ctx context.Context, appid string) (*wechatPay.Pay, error) {
	wechat, err := payment.GetWechatPayment(ctx, getWebCache(), appid)
	if err != nil {
		return nil, err
	}

	return wechatPay.NewPay(&wechatPayConfig.Config{
		Sandbox: false,
		AppID:   appid,
		MchID:   wechat.MchID,
		Key:     wechat.PrivateKey,
	}), nil
}
