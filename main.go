package main

import (
	"blogger/controller"
	"blogger/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//	初始化数据库连接
	dns := "root:123456@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	//	加载静态文件
	router.Static("/static/", "./static")
	//	加载模板
	router.LoadHTMLGlob("views/*")
	//	访问主页
	router.GET("/", controller.IndexHandle)
	router.GET("/category", controller.CategoryList)
	_ = router.Run(":8001")
}
