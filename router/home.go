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

	/**
	插入文章
	*/
	group.POST("/article/save", controller.HandleArticleSave)

	/**
	根据文章id获取单个文章信息
	*/
	group.GET("/getArticleById", controller.HandleGetSingleArticle)
}
