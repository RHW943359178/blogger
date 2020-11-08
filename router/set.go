package router

import (
	"blogger/service"
	"github.com/gin-gonic/gin"
)

func LoadSet(e *gin.Engine) {
	e.GET("/set/getCategory", GetAllCategory)
}
