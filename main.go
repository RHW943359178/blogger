package main

import (
	"blogger/controller"
	"blogger/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//	初始化数据库连接
	dns := "RHW:RHW943359178@tcp(81.69.255.188:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	//	访问主页
	router.GET("/", controller.IndexHandle)
	router.GET("/category", controller.CategoryList)
	_ = router.Run(":8001")
}
