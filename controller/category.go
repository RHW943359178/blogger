package controller

import (
	"blogger/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//	@Tags 获取所有的分类
//	@Accept application/json
//	@Produce application/json
//  @Success 200 {object} ResponseCategoryList
//  @Router /home/category [get]
func GetAllCategory(c *gin.Context) {
	//	从service层取数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取分类信息成功",
		"data":    categoryList,
	})
}
