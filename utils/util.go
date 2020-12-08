package utils

import (
	session "blogger/cookie_session"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const base_format = "2006-01-02 15:04:05"

func GetRootDir() (rootPath string) {
	exePath := os.Args[0]
	rootPath = filepath.Dir(exePath)
	return rootPath
}

/**
验证登录公用方法
*/
func UnauthorizedMethod(c *gin.Context) (user map[string]string) {
	sessionID := c.MustGet("sessionID").(string)
	redis := c.MustGet("session").(session.Session)
	userInfo, err := redis.GetData(sessionID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户未登陆！",
			"data":    nil,
		})
		return
	}
	//	将字符串转为map
	tmpUser := make(map[string]string, 0)
	err = json.Unmarshal([]byte(userInfo), &tmpUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": fmt.Sprintf("unmarshal data from redis failed, err: %v\n", err),
		})
		return
	}
	return tmpUser
}

/**
数据库时间格式转换
*/
func DataTimeFormat(t string) (sTime string, err error) {
	//	参数如果为空
	//if t == nil {
	//
	//}
	format, err := time.Parse(base_format, t)
	if err != nil {
		log.Fatalln("parse time failed, err: ", err)
		return
	}
	sTime = format.Format(base_format)
	//DefaultTimeLoc := time.Local
	//loginTime, err := time.ParseInLocation(base_format, t, DefaultTimeLoc)
	return
}
