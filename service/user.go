package service

import (
	"blogger/dao/db"
	"blogger/model"
	"log"
)

//	插入保存用户的信息
func InsertUserInfo(userInfo *model.User) (insertId int64, err error) {
	//	从db层获取数据
	insertId, err = db.InsertUserInfo(userInfo)
	if err != nil {
		log.Fatalln("get insertId from db failed, err: ", err)
		return
	}
	return
}

//	获取用户表中是否包含某个字段
func ConditionSelect(condition string) (exist int) {
	//	从 db 层取数据
	exist = db.ConditionSelect(condition)
	//if err != nil {
	//	log.Fatal("get condition from db.ConditionSelect failed, err: ", err)
	//	return
	//}
	return
}

//	校验用户登录状态
func ValidateStatus(user *model.User) (status int, resUser *model.ResUser) {
	//	从db层取数据
	//user = &model.User{}
	status, user = db.ValidateLogin(user.Username, user.Password)
	if status == 0 || status == 1 {
		return status, nil
	} else {
		//	初始化结构体
		resUser = &model.ResUser{}
		resUser.Username = user.Username
		resUser.UserId = user.UserId
		resUser.ImgUrl = user.ImgUrl
		return
	}
}

//	插入用户头像
func UpdateUserImg(userId, imgUrl string) (row interface{}, err error) {
	//	从db层取数据
	row, err = db.InsertUserIcon(userId, imgUrl)
	if err != nil {
		log.Fatalln("get img from db failed, err: ", err)
		return
	}
	return
}

//	获取用户信息
func GetUserInfo(userId string) (user *model.ResUser, err error) {
	//	从db层取数据
	user, err = db.GetUserInfo(userId)
	if err != nil {
		log.Println("get userInfo from db failed, err: ", err)
	}
	return
}

//	获取用户的总文章数和总字数
func GetUserArticleInfo(userId string) (fontCount, articleCount int, err error) {
	//	从 db 层取数据
	fontCount, articleCount, err = db.GetAuthorInfo(userId)
	if err != nil {
		log.Println("get article_count, font_count from db failed, err: ", err)
	}
	return
}
