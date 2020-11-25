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
//  @Param commentCount body int true "评论数"Article
////  @Router /home/article/save [post]
//  @Success 200 {object} Response
func HandleArticleSave(c *gin.Context) {
	//fmt.Println(c.Request.Body, 123)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "hello",
	//})

	//	获取传入参数
	//c.PostForm()
	//fmt.Printf("%v\n", c.Request.Body)
	// 根据请求body创建一个json解析器实例
	//decoder := json.NewDecoder(c.Request.Body)
	////
	//////	用于存放参数key-value的map
	//var params map[string]interface{}
	////
	////// 解析参数 存入map
	//err := decoder.Decode(&params)
	//if err != nil {
	//	log.Fatalln("decoder failed, err: ", err)
	//	return
	//}
	////log.Println(decoder)
	//title := params["title"]
	//fmt.Println(title, 123)
	//if title == "" {
	//	log.Fatalln("文章标题不可为空")
	//	return
	//}
	//summary := params["summary"]
	//username := params["username"]
	//if username == "" {
	//	log.Fatalln("文章作者不可为空")
	//	return
	//}
	//var categoryId int64 //	声明categoryId变量
	//id := params["categoryId"]
	//if id == "" {
	//	log.Fatalln("所属分类不可为空")
	//	return
	//} else {
	//	categoryId, _ = strconv.ParseInt(string(id), 10, 64)
	//}
	//var viewCount, commentCount uint32
	//////view := c.PostForm("viewCount")
	//view, _ := strconv.ParseUint(params["viewCount"], 10, 64)
	//viewCount = uint32(view)
	////
	//comment, _ := strconv.ParseUint(params["commentCount"], 10, 64)
	//commentCount = uint32(comment)
	//content := params["content"]
	//if content == "" {
	//	log.Fatalln("文章内容不可为空")
	//	return
	//}
	////	初始化结构体
	//articleDetail := &model.ArticleDetail{
	//	Content: content,
	//	ArticleInfo: model.ArticleInfo{
	//		CategoryId:   categoryId,
	//		Summary:      summary,
	//		Title:        title,
	//		ViewCount:    viewCount,
	//		CommentCount: commentCount,
	//		Username:     username,
	//		CreateTime:   time.Now(),
	//	},
	//}
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
