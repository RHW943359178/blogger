package db

import (
	"blogger/model"
	"github.com/jmoiron/sqlx"
	"log"
)

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

//	获取文章列表(分页，分类)
func GetArticleList(condition string, categoryId []string, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	//	返回值切片初始化
	articleList = make([]*model.ArticleInfo, 0)
	//	参数验证
	//log.Println(len(categoryId), 12)
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	//	判断category的值，如果为0则查全部分类的文章
	log.Println(len(categoryId), 12)
	var sqlStr string
	if len(categoryId) != 0 {
		sqlStr, args, err := sqlx.In(`select
					id, summary, category_id, title, view_count, create_time, comment_count, username
				from article
				where category_id in (?) and title like concat('%', ?, '%')
				order by create_time desc limit ?, ?`, categoryId, condition, pageNum, pageSize)
		if err != nil {
			log.Fatalln("err", err)
			//return
		}
		//sqlStr = `select
		//			id, summary, category_id, title, view_count, create_time, comment_count, username
		//		from article
		//		where category_id in (?) and title like concat('%', ?) or title like concat(?, '%')
		//		order by create_time desc limit ?, ?`
		err = db.Select(&articleList, sqlStr, args...)
	} else {
		sqlStr = `select
					id, summary, category_id, title, view_count, create_time, comment_count, username
				from article
				where title like concat('%', ?, '%')
				order by create_time desc limit ?, ?`
		err = db.Select(&articleList, sqlStr, condition, pageNum, pageSize)
	}
	log.Println(articleList, "&articleList")
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
