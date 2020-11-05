package db

import "blogger/model"

//	插入文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	//	加个验证
	if article == nil {
		return
	}
	sqlStr := `insert into article(content, summary, title, username, category_id, view_count, comment_count) value(?,?,?,?,?,?,?)`
	result, err := db.Exec(sqlStr, article.Content, article.Summary, article.Title, article.Username,
		article.ViewCount, article.ArticleInfo.CategoryId, article.ViewCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

//	获取文章列表，分页
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	//	参数验证
	if pageNum <= 0 || pageSize <= 0 {
		return
	}
	//	时间降序排列
	sqlStr := `select id, summary, title, view_count, create_time, comment_count, username
				from article where status = 1 order by create_time desc limit ?, ?`
	err = db.Select(&articleList, sqlStr, pageNum, pageSize)
	if err != nil {
		return
	}
	return
}

//	根据文章id，查询单个文章
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail = &model.ArticleDetail{}
	if articleId < 0 {
		return
	}
	sqlStr := `select id, summary, title, view_count, content, create_time, comment_count, username, category_id
				from article where id = ? and status = 1`
	err = db.Get(articleDetail, sqlStr, articleId)
	return
}

//	根据分类id， 查询这一类的文章
func GetArticleListByCategoryId(categoryId, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	//	参数验证
	if pageNum <= 0 || pageSize <= 0 {
		return
	}
	sqlStr := `select id, summary, title, view_count, create_time, comment_count, username, category
				from article where status = 1 and category_id = ? order by create_time desc limit ?, ?`
	err = db.Select(&articleList, sqlStr, categoryId, pageNum, pageSize)
	return
}
