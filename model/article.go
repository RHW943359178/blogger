package model

import "time"

//	id, category_id, content, title, view_count, comment_content, username, STATUS, summary, create_time, update_time

//	定义文章结构体
type ArticleInfo struct {
	Id         int64 `db:"id" json:"id"`
	CategoryId int64 `db:"category_id" json:"categoryId"`
	//	文章摘要
	Summary      string `db:"summary" json:"summary"`
	Title        string `db:"title" json:"title"`
	ViewCount    uint32 `db:"view_count" json:"viewCount"`
	CommentCount uint32 `db:"comment_count" json:"commentCount"`
	//	时间
	CreateTime time.Time `db:"create_time" json:"createTime"`
	Username   string    `db:"username" json:"username"`
}

//	用于文章详情页的实体
//	为了提升效率
type ArticleDetail struct {
	ArticleInfo
	//	文章内容
	Content string `db:"content"`
	Category
}

//	用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
