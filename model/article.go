package model

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
	//	创建时间
	CreateTime interface{} `db:"create_time" json:"createTime"`
	UpdateTime interface{} `db:"update_time" json:"updateTime"`
	Username   string      `db:"username" json:"username"`
	UserId     string      `db:"user_id" json:"userId"`
	OpenFlag   int         `db:"open_flag" json:"openFlag"`
}

//	用于文章详情页的实体
//	为了提升效率
type ArticleDetail struct {
	ArticleInfo
	//	文章内容
	Content string `db:"content" json:"content"`
}

//	用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}

//	文章保存时绑定的结构体
type ArticleBind struct {
	CategoryId int64 `json:"categoryId"`
	//	文章摘要
	Summary      string `json:"summary"`
	Title        string `json:"title"`
	ViewCount    uint32 `json:"viewCount"`
	CommentCount uint32 `json:"commentCount"`
	//	时间
	//CreateTime time.Time `json:"createTime"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

//	与用户相关的文章信息
type UserArticle struct {
	Id         int64       `db:"id" json:"id"`
	CategoryId int64       `db:"category_id" json:"categoryId"`
	Title      string      `db:"title" json:"title"`
	CreateTime interface{} `db:"create_time" json:"createTime"`
}
