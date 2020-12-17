package main

import (
	"blogger/dao/db"
	"blogger/dao/redis"
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

	//	初始化配置文件
	//conf := config.GetConf()

	//	初始化 redis 数据库
	rdn := "81.69.255.188:6379"
	//rdn := conf.RedisAddr
	err := redis.InitClient(rdn)
	if err != nil {
		panic(err)
	}
	//	初始化 mysql 数据库连接
	dns := "RHW:RHW943359178@tcp(81.69.255.188:3306)/blogger?parseTime=true&loc=Local"
	//dns := conf.MysqlAddr
	err = db.Init(dns)
	if err != nil {
		panic(err)
	}

	router.VisitHomeInterface(r)
	router.VisitUserInterface(r)
	router.VisitCommonInterface(r)

	//	生成 swagger 文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	_ = r.Run(":8000")
	//_ = r.Run(":" + conf.Port)
}
