package model

import "database/sql"

//	定义用户信息结构体
type User struct {
	Id       int            `db:"id"`
	UserId   string         `db:"user_id" form:"userId" json:"userId"`
	Username string         `db:"username" form:"username" json:"username"`
	Password string         `db:"password" form:"password" json:"password"`
	Email    string         `db:"email" form:"email" json:"email"`
	ImgUrl   sql.NullString `db:"img_url" form:"imgUrl" json:"imgUrl"`
}

//	定义校验成功后返回结构体
type ResUser struct {
	UserId       string         `db:"user_id" json:"userId"`
	Username     string         `db:"username" json:"username"`
	ImgUrl       sql.NullString `db:"img_url" json:"imgUrl"`
	ArticleCount int            `db:"article_count" json:"articleCount"`
	FontCount    int            `db:"font_count" json:"fontCount"`
}
