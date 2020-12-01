package router

import (
	"blogger/controller"
	"blogger/utils"
	"github.com/gin-gonic/gin"
)

//	访问用户相关接口
func VisitUserInterface(r *gin.Engine) {
	//	用户相关理由组
	userGroup := r.Group("/user")
	/**
	用户信息保存接口
	*/
	userGroup.POST("/save", controller.HandleSaveUserInfo)

	/**
	查询用户名是否已经注册
	*/
	userGroup.GET("username/select", controller.HandleConditionSelect)
	/**
	用户登录校验
	*/
	userGroup.POST("login/validate", utils.SessionMiddleware(), controller.ValidateLoginStatus)
}
