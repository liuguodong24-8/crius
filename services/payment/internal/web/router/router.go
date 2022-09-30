package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.omytech.com.cn/micro-service/payment/internal/web/controller"
)

// Router ...
var Router *gin.Engine

// Init ...
func Init() {
	Router = gin.Default()
	// 支付回调
	Router.Any(`/api/v1/wechat/:appid/pay/notify`, controller.WechatPayNotify)
	// 退款回调
	Router.Any(`/api/v1/wechat/:appid/refund/notify`, controller.WechatRefundNotify)
}
