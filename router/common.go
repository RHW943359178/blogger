package router

import (
	"blogger/controller"
	"github.com/gin-gonic/gin"
)

//	访问公共服务接口
func VisitCommonInterface(r *gin.Engine) {
	//	路由组
	commonGroup := r.Group("/common")

	//	返回当前ip地址
	commonGroup.GET("/getServerIp", controller.GetServerIp)
}
