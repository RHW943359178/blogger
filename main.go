package main

import (
	"blogger/dao/db"
	_ "blogger/docs"
	"blogger/router"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//	@title blogger项目接口文档
//	@version 1.0
//	@description blogger是一个开源的golang + vue项目，用于平时记录心得以及一些衍生功能

//	@host 8000
//	@BaseUrl localhost:8080
func main() {
	r := gin.Default()
	//	初始化数据库连接
	dns := "RHW:RHW943359178@tcp(81.69.255.188:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	router.VisitHomeInterface(r)
	router.VisitUserInterface(r)

	//	生成 swagger 文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	_ = r.Run(":8000")
}
