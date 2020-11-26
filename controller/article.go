package controller

import (
	"blogger/model"
	"blogger/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"time"

	//"log"
	"net/http"
	//"strconv"
)

/*
处理文章相关接口
*/

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
//  @Param commentCount body int true "评论数"Article
//  @Router /home/article/save [post]
//  @Success 200 {object} int64
func HandleArticleSave(c *gin.Context) {
	var articleBind model.ArticleBind
	err := c.ShouldBind(&articleBind)
	if err != nil {
		log.Fatalln("struct bind failed, err: ", err)
		return
	}
	//	验证标题参数
	if articleBind.Title == "" {
		log.Fatalln("标题参数为空")
		return
	}
	//	验证作者姓名参数
	if articleBind.Username == "" {
		log.Fatalln("文章作者参数为空")
		return
	}
	//	验证文章内容参数
	if articleBind.Content == "" {
		log.Fatalln("文章作者参数为空")
		return
	}
	//	初始化文章具体信息结构体
	articleDetail := &model.ArticleDetail{
		Content: articleBind.Content,
		ArticleInfo: model.ArticleInfo{
			CategoryId:   articleBind.CategoryId,
			Summary:      articleBind.Summary,
			Title:        articleBind.Title,
			ViewCount:    articleBind.ViewCount,
			CommentCount: articleBind.CommentCount,
			Username:     articleBind.Username,
			CreateTime:   time.Now(),
		},
	}
	//	从service层取数数据
	insertId, err := service.ArticleSave(articleDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "保存文章成功！",
		"data":    insertId,
	})
}

//	@Tags 根据id获取单个文章信息
//  @Param articleId query int64 true "文章id"
//  @Router /home/getArticleById [get]
//  @Success 200 {object} ResponseGetSingleArticle
func HandleGetSingleArticle(c *gin.Context) {
	//	获取参数
	id := c.Query("articleId")
	fmt.Println("id: ", id)
	if id == "" {
		log.Fatalln("文章参数为空")
		return
	}
	articleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Fatalln("parse id to int failed, err: ", err)
		return
	}
	//	从service层取数据
	article, err := service.GetArticleInfoById(articleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "根据id获取文章信息成功",
		"data":    article,
	})

}
