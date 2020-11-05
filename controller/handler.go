package controller

import (
	"blogger/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//	访问主页的控制器
func IndexHandle(c *gin.Context) {
	//	从service取数据
	//	1.加载文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	//	2.加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	//	gin.H本质上是个map
	//var data map[string]interface{} = make(map[string]interface{}, 16)
	//data["article_list"] = articleRecordList
	//data["category"] = categoryList
	//c.JSON(http.StatusOK, data)

	c.JSON(http.StatusOK, gin.H{
		"article_list": articleRecordList,
		"category":     categoryList,
	})
}

//	获取分类数据
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	//	转成Int
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	//	根据分类id，获取文章列表
	articleRecordList, err := service.GetArticleRecordListById(categoryId, 0, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	//	再次加载所有分类数据，用于分类云显示
	categoryList, err := service.GetAllCategoryList()
	c.JSON(http.StatusOK, gin.H{
		"article_list": articleRecordList,
		"category":     categoryList,
	})
}
