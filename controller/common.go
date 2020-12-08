package controller

import (
	"blogger/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
处理公告服务相关接口
*/

//	@Tags 获取服务器ip和端口
func GetServerIp(c *gin.Context) {
	dir := utils.GetRootDir()
	//ip := net.IP{}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": dir,
		"data":    dir,
	})
	fmt.Println("dir: ", dir)
}
