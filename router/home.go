package router

import (
	"blogger/controller"
	"github.com/gin-gonic/gin"
)

//	访问主页面接口
func VisitHomeInterface(r *gin.Engine) {
	//	主页路由组
	group := r.Group("/home")

	/*
		获取所有分类列表
	*/
	group.GET("/category", controller.GetAllCategory)

	/**
	获取文章分类列表
	*/
	group.GET("/article", controller.GetAllArticleList)
}
