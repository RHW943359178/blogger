package controller

import (
	"blogger/model"
	"blogger/service"
	"blogger/utils"
	//"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
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
	var categoryId []string
	if c.Query("categoryId") == "" { //	categoryId 判空处理
		categoryId = []string{}
	} else {
		categoryId = strings.Split(c.Query("categoryId"), ",")
	}
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
	var articleBind model.ArticleDetail
	err := c.ShouldBind(&articleBind)
	if err != nil {
		log.Fatalln("struct bind failed, err: ", err)
		return
	}
	//	验证标题参数
	if articleBind.Title == "" || articleBind.Username == "" || articleBind.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求参数有有误",
		})
		return
	}
	//	验证 session 值并从数据库匹配
	//sessionID := c.MustGet("sessionID").(string)
	//redis := c.MustGet("session").(session.Session)
	//userInfo, err := redis.GetData(sessionID)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"code":    401,
	//		"message": "用户未登陆！",
	//		"data":    nil,
	//	})
	//}
	////	将字符串转为map
	//tmpUser := make(map[string]string, 0)
	//err = json.Unmarshal([]byte(userInfo), &tmpUser)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"code":    500,
	//		"message": fmt.Sprintf("unmarshal data from redis failed, err: %v\n", err),
	//	})
	//}
	tmpUser := utils.UnauthorizedMethod(c)
	//	将前端传入的 userID和username 传到数据库
	articleBind.UserId = tmpUser["userId"]
	articleBind.Username = tmpUser["username"]
	//fmt.Println("userInfo: ", userInfo)

	//	初始化文章具体信息结构体
	//articleDetail := &model.ArticleDetail{
	//	Content: articleBind.Content,
	//	ArticleInfo: model.ArticleInfo{
	//		CategoryId:   articleBind.CategoryId,
	//		Summary:      articleBind.Summary,
	//		Title:        articleBind.Title,
	//		ViewCount:    articleBind.ViewCount,
	//		CommentCount: articleBind.CommentCount,
	//		Username:     articleBind.Username,
	//		CreateTime:   time.Now(),
	//	},
	//}
	//	从service层取数数据
	insertId, err := service.ArticleSave(&articleBind)
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
	if id == "" {
		log.Fatalln("文章参数为空")
		return
	}
	articleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	//	从service层取数据
	article, err, errFlag := service.GetArticleInfoById(articleId)
	if err != nil {
		if errFlag == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": err,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "根据id获取文章信息成功",
		"data":    article,
	})

}

//	@Tags 根据用户 id 查询该类目下所有的文章信息
//	@Accept application/json
//	@Produce application/json
//  @Router /home/getArticleByUserId [get]
//  @Success 200 {object} ResponseUserArticle
func GetAllArticleByUserId(c *gin.Context) {
	//	定义参数类型
	var (
		err  error
		size int
		num  int
		data []*model.UserArticle
	)

	//	从请求中获取参数，有则为无登录请求，无则为登录时请求
	userId := c.Query("userId")     //	用户id
	pageSize := c.Query("pageSize") //	页面范围
	pageNum := c.Query("pageNum")   //	页面数

	if userId == "" && pageSize == "" && pageNum == "" {
		//	验证 session 值并从数据库匹配
		tmpUser := utils.UnauthorizedMethod(c)
		//	从service层获取数据
		data, err = service.GetArticleListByUserId(tmpUser["userId"], 0, 0)
	} else {
		size, err = strconv.Atoi(pageSize)
		num, err = strconv.Atoi(pageNum)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		data, err = service.GetArticleListByUserId(userId, size, num)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "根据用户id获取文章信息成功",
		"data":    data,
	})
}

//	@Tags 根据文章id修改文章内容
//	@Accept application/json
//	@Produce application/json
//  @Param id body int64 true "文章id"
//  @Param content body string true "文章内容"
//  @Param summary body string false "文章内容梗概"
//  @Param title query string true "文章标题"
//  @Param categoryId body int64 true "分类id"
//  @Param openFlag body int true "公开标志"
//  @Router /home/updateArticleInfo [post]
//  @Success 200 {object} ResponseUserArticle
func UpdateArticleInfo(c *gin.Context) {
	//	验证 session 值并从数据库匹配
	_ = utils.UnauthorizedMethod(c)
	var articleBind *model.ArticleDetail
	err := c.ShouldBind(&articleBind)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	//	验证传入参数
	if articleBind.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "文章id不可为空",
		})
		return
	}
	if articleBind.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "文章内容不可为空",
		})
		return
	}
	if articleBind.CategoryId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "文章分类不可为空",
		})
		return
	}
	//	从service层获取数据
	row, err := service.UpdateArticleInfo(articleBind)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改文章信息成功",
		"data":    row,
	})
}

//	@Tags 根据文章id删除文章内容
//	@Accept application/json
//	@Produce application/json
//  @Param id query int64 true "文章id"
//  @Router /home/article/delete [post]
//  @Success 200 {object} row int
func DeleteArticle(c *gin.Context) {
	//	验证 session 值并从数据库匹配
	_ = utils.UnauthorizedMethod(c)
	var articleBind *model.ArticleDetail
	err := c.ShouldBind(&articleBind)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	if articleBind.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "文章id参数异常",
		})
		return
	}
	//	从server层取数据
	row, err := service.DeleteArticle(articleBind)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "文章删除成功",
		"data":    row,
	})
}

//	@Tags 根据用户 id 查询该用户的其他文章
//	@Accept application/json
//	@Produce application/json
//  @Router /home/getOtherArticle [post]
//  @Success 200 {object} ResponseUserArticle
func GetOtherArticle(c *gin.Context) {
	//	绑定参数结构体
	otherArticle := struct {
		UserId    string `json:"userId"`
		ArticleId int64  `json:"articleId"`
		PageNum   int    `json:"pageNum"`
		PageSize  int    `json:"pageSize"`
	}{}
	err := c.ShouldBind(&otherArticle)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err,
		})
		return
	}
	//	从service层取数据
	articleList, err := service.GetOtherArticle(otherArticle.UserId, otherArticle.ArticleId, otherArticle.PageSize, otherArticle.PageNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取其他文章列表成功",
		"data":    articleList,
	})
	return

}

//	@Tags 根据 categoryId 查询推荐文章
//	@Accept application/json
//	@Produce application/json
//  @Router /home/getRecommendArticle [post]
//  @Success 200 {object} ResponseUserArticle
func GetRecommendArticle(c *gin.Context) {
	//	分类id
	category := c.Query("categoryId")
	//	推荐文章条数
	num := c.Query("num")
	// 验证参数
	if category == "" || num == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
		})
		return
	}
	var (
		categoryId  int64
		number      int
		err         error
		articleList []*model.UserArticle
	)
	//	将参数转成指定类型
	categoryId, err = strconv.ParseInt(category, 10, 64)
	number, err = strconv.Atoi(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	//	从 service 层取数据
	articleList, err = service.GetRecommendArticle(categoryId, number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取推荐文章列表成功",
		"data":    articleList,
	})
}
