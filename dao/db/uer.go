package db

import (
	"blogger/model"
	"log"
)

//	插入用户保存
func InsertUserInfo(userInfo *model.User) (id int64, err error) {
	if userInfo == nil {
		log.Fatal("invalid userInfo")
		return
	}
	sql := `insert into user(user_id, username, password, email) values (?, ?, ?, ?)`
	ret, err := db.Exec(sql, userInfo.UserId, userInfo.Username, userInfo.Password, userInfo.Email)
	if err != nil {
		log.Fatalln("Exec to table user failed, err: ", err)
		return
	}
	id, err = ret.LastInsertId() //	新插入数据的id
	if err != nil {
		log.Fatalln("get lastInsertId failed", err)
		return
	}
	return
}

//	查询某个姓名用于前台校验
func ConditionSelect(condition string) (exist int) {
	var user model.User
	var err error
	sql := `select username from user where username=?`
	//	db.Get(&user, sql, condition) 方法没有值会返回错误
	if err = db.Get(&user, sql, condition); err != nil {
		exist = 0
		return
	}
	//if user.Username == "" {
	//	exist = false
	//} else {
	exist = 1
	//}
	return
}

//	登录状态校验
func ValidateLogin(username, password string) (status int, user *model.User) {
	//	status 代表的相应状态 0：数据库查无此人， 1：用户名或密码错误， 2：登录成功
	user = &model.User{}
	var err error
	sql := `select username, user_id, password, img_url from user where username=?`
	//	查无数据
	if err = db.Get(user, sql, username); err != nil {
		status = 0
		return
	}
	//	判断用户名和密码是否匹配
	if user.Password != password {
		status = 1
		return
	}
	status = 2
	return
}

//	插入用户头像
func InsertUserIcon(userId, imgUrl string) (row interface{}, err error) {
	sqlStr := `update from user set img_url=? where user_id = ?`
	row, err = db.Exec(sqlStr, imgUrl, userId)
	return
}

//	根据用户id获取用户信息
func GetUserInfo(userId string) (user *model.ResUser, err error) {
	user = &model.ResUser{}
	sqlStr := `select username, user_id, img_url, font_count, article_count from user where user_Id = ?`
	err = db.Get(user, sqlStr, userId)
	return
}

//	根据 userId 获取该作者总文章数和总字数
func GetAuthorInfo(userId string) (fontCount, articleCount int, err error) {
	data := struct {
		FontCount    int `db:"SUM(CHAR_LENGTH(content))" json:"fontCount"`
		ArticleCount int `db:"COUNT(*)" json:"articleCount"`
	}{}
	sqlStr := `select SUM(CHAR_LENGTH(content)), COUNT(*) from article where user_id = ?`
	err = db.Get(&data, sqlStr, userId)
	fontCount = data.FontCount
	articleCount = data.ArticleCount
	return
}
