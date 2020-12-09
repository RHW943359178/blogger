package router

import (
	"blogger/controller"
	"blogger/utils"
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
	group.POST("/article/save", utils.SessionMiddleware(), controller.HandleArticleSave)

	/**
	根据文章id获取单个文章信息
	*/
	group.GET("/getArticleById", controller.HandleGetSingleArticle)

	/**
	根据文章id更新文章信息
	*/
	group.POST("/updateArticleInfo", utils.SessionMiddleware(), controller.UpdateArticleInfo)

	/**
	根据用户 id 获取所有文章信息
	*/
	group.GET("/getArticleByUserId", utils.SessionMiddleware(), controller.GetAllArticleByUserId)

	/**
	根据用户 id 删除文章
	*/
	group.POST("/article/delete", utils.SessionMiddleware(), controller.DeleteArticle)

	/**
	查询用户其他文章列表
	*/
	group.POST("getOtherArticle", controller.GetOtherArticle)
}
