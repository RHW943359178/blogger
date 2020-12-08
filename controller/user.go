package controller

import (
	session "blogger/cookie_session"
	"blogger/model"
	"blogger/service"
	"blogger/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/**
处理用户信息相关接口
*/

//	@Tags 用户信息保存
//	@Accept application/json
//	@Produce application/json
//  @Param userId body string true "用户id"
//  @Param username body string true "用户姓名"
//  @Param email body string true "邮箱"
//  @Param password body string true "邮箱"
//  @Success 200 {object} int64
//  @Router /user/save [post]
//	用户信息保存并获取插入id
func HandleSaveUserInfo(c *gin.Context) {
	var userBind model.User
	err := c.ShouldBind(&userBind)
	if err != nil {
		log.Fatalln("bind user failed, err: ", err)
		return
	}
	//	绑定生成的guid值
	userBind.UserId = utils.UniqueId()
	//	从service层取数据
	insertId, err := service.InsertUserInfo(&userBind)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("get insertId id from service failed, err: %v\n", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "用户注册成功！",
		"data":    insertId,
	})
}

//	@Tags 查询数据库是否有已经注册
//	@Param condition query string true "用户姓名"
//  @Success 200 {object} string
//  @Router /user/username/select [get]
func HandleConditionSelect(c *gin.Context) {
	condition := c.Query("condition")
	if condition == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "无效字段",
		})
		return
	}
	//	从service层中取数据
	//exist, err := service.ConditionSelect(condition)
	exist := service.ConditionSelect(condition)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"message": fmt.Sprintf("get exist from service failed, err:  %v\n", err),
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "返回成功",
		"data":    exist,
	})
}

//	@Tags 用户登录校验
//	@Param username body string true "用户名"
//	@Param password body string true "密码"
//  @Success 200 {object} ResponseUserInfo
//  @Router /user/login/validate [post]
func ValidateLoginStatus(c *gin.Context) {
	var userBind model.User
	if err := c.ShouldBind(&userBind); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("bind model failed, err: %v\n", err),
		})
		return
	}
	//	校验用户名参数
	if userBind.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户名不可为空",
		})
		return
	}
	if userBind.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码不可为空",
		})
		return
	}
	//	从service层取数据
	status, user := service.ValidateStatus(&userBind)
	//	声明一个消息
	var message string
	//var count int
	//	声明返回的数据
	data := make(map[string]interface{}, 0)
	data["status"] = status
	data["user"] = user
	if data["status"] == 0 {
		message = "该用户未注册，请重试"
	} else if data["status"] == 1 {
		message = "用户名或密码错误"
	} else {
		message = "登录成功！"
		//	获取 sessionId
		sessionID := c.MustGet("sessionID").(string)
		//	获取 options 配置
		options := c.MustGet("options").(session.Options)
		//	获取 session 实例
		redis := c.MustGet("session").(session.Session)
		//	设置 redis 库的数据结构为：用户 session: 用户信息 userInfo
		userInfo := make(map[string]string)
		userInfo["username"] = user.Username
		userInfo["userId"] = user.UserId

		redis.Set("userInfo", userInfo)
		data["session"] = sessionID
		c.SetCookie(session.SessionCookieName, sessionID, options.MaxAge, options.Path, options.Domain, options.Secure, options.HttpOnly)
	}
	//	构建返回消息 map
	resData := map[string]interface{}{
		"code":    200,
		"message": message,
		"data":    data,
	}
	c.JSON(http.StatusOK, resData)
}

//	@Tags 用户上传头像
func HandleImgUpload(c *gin.Context) {
	//	验证 session 值并从数据库匹配
	user := utils.UnauthorizedMethod(c)
	//	单个文件
	file, err := c.FormFile("icon")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	//	给 icon 分配一个id值
	iconId := utils.UniqueId()
	dst := fmt.Sprintf("C:/user/icon/%s", iconId)
	//	上传文件到指定的目录
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	// 同步数据到数据库
	_, err = service.UpdateUserImg(user["userId"], iconId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "图片上传成功",
		"data":    iconId,
	})
	return
}
