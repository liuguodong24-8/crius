package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.omytech.com.cn/micro-service/Crius/internal/web/controller"
)

// Router ...
var Router *gin.Engine

// Init ...
func Init() {
	Router = gin.Default()

	Router.GET(`/hello`, controller.Hello)
}
