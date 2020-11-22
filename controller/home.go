package controller

import (
	"blogger/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"

	//"log"
	"net/http"
	//"strconv"
)

/*
处理主页接口
*/

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

//	@Tags 获取所有的文章列表
//	@Accept application/json
//	@Produce application/json
//  @Param condition query string false "查询条件"
//  @Param categoryId query int false "分类id"
//  @Param pageNum query int true "页面数"
//  @Param pageSize query int true "页码范围"
//  @Success 200 {object} ResponseArticleList
//  @Router /home/article [get]
func GetAllArticleList(c *gin.Context) {
	var err error
	//	从service层取数据
	condition := c.Query("condition")
	//	categoryId
	fmt.Println("condition", condition)
	var categoryId []string
	if c.Query("categoryId") == "" { //	categoryId 判空处理
		categoryId = []string{}
	} else {
		categoryId = strings.Split(c.Query("categoryId"), ",")
	}
	fmt.Println("categoryId", categoryId)
	// pageNum 验证
	var pageNum int
	if c.Query("pageNum") == "" { //	如果为空，就默认是第一页
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Query("pageNum"))
	}
	if err != nil {
		log.Fatalln("err", err)
		return
	}
	//	pageSize 验证
	var pageSize int
	if c.Query("pageSize") == "" { //	如果为空，就默认是十条数据一页
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(c.Query("pageSize"))
	}
	if err != nil {
		log.Fatalln("err", err)
		return
	}
	articleList, err := service.GetArticleListByCondition(condition, categoryId, pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	dataMap := make(map[string]interface{})
	dataMap["list"] = articleList
	dataMap["count"] = len(articleList)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取文章列表成功！",
		"data":    dataMap,
	})

}

//	@Tags 文章保存
//	@Accept application/json
//	@Produce application/json
//  @Param title body string true "文章标题"
//  @Param summary body string false "文章梗概"
//  @Param categoryId body int true "所属分类id"
//  @Param content body string true "文章内容"
//  @Param username body string true "文章作者名称"
//  @Param viewCount body int true "浏览数"
//  @Param commentCount body int true "评论数"
//  @Success 200 {object} ResponseArticle
//  @Router /home/article/save [post]
func HandleArticleSave(c *gin.Context) {

}
