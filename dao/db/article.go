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
	sqlStr := `insert into article(content, summary, title, user_id, username, category_id, view_count, open_flag, comment_count) value(?,?,?,?,?,?,?,?,?)`
	result, err := db.Exec(sqlStr, article.Content, article.ArticleInfo.Summary, article.ArticleInfo.Title, article.ArticleInfo.UserId, article.ArticleInfo.Username,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CategoryId, article.ArticleInfo.OpenFlag, article.ArticleInfo.ViewCount)
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
	//	判断category的值，如果为0则查全部分类的文章文章参数为空
	log.Println(len(categoryId), 12)
	var sqlStr string
	if len(categoryId) != 0 {
		sqlStr, args, err := sqlx.In(`select
					id, summary, category_id, title, view_count, create_time, update_time, comment_count, username
				from article
				where category_id in (?) and title like concat('%', ?, '%')
				order by create_time desc limit ?, ?`, categoryId, condition, pageNum, pageSize)
		if err != nil {
			log.Fatalln("err", err)
			//return
		}
		err = db.Select(&articleList, sqlStr, args...)
	} else {
		sqlStr = `select
					id, summary, category_id, title, view_count, create_time, update_time, comment_count, username
				from article
				where title like concat('%', ?, '%')
				order by create_time desc limit ?, ?`
		err = db.Select(&articleList, sqlStr, condition, pageNum, pageSize)
	}
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
	sqlStr := `select id, summary, title, view_count, content, create_time, update_time, comment_count, user_id, username, category_id, open_flag
				from article where id = ?`
	err = db.Get(articleDetail, sqlStr, articleId)
	//if err != nil {
	//	log.Fatalln("db.Get failed, err: ", err)
	//	return
	//}
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

//	根据用户id， 查询该用户id下所有的文章
func GetArticleByUserId(userId string, pageSize, pageNum int) (articleList []*model.UserArticle, err error) {
	if pageSize == 0 && pageNum == 0 {
		sqlStr := `select id, category_id, create_time, title from article where user_id = ?`
		err = db.Select(&articleList, sqlStr, userId)
	} else {
		sqlStr := `select id, category_id, create_time, title, summary from article where user_id = ? limit ?, ?`
		err = db.Select(&articleList, sqlStr, userId, (pageNum-1)*pageSize, pageSize)
	}
	return
}

//	根据用户 id 页码范围查询其他文章
func GetOtherArticle(userId string, articleId int64, pageSize, pageNum int) (articleList []*model.UserArticle, err error) {
	sqlStr := `select id, view_count,  title from article  where user_id = ? and id != ? limit ?, ?`
	err = db.Select(&articleList, sqlStr, userId, articleId, (pageNum-1)*pageSize, pageSize)
	//err = db.Select(&articleList, sqlStr, userId, articleId, 30, 40)
	return
}

//	根据 categoryId 获取推荐文章阅读
func GetRecommendArticle(categoryId int64, num int) (articleList []*model.UserArticle, err error) {
	sqlStr := `select id, view_count, title from article where category_id = ? ORDER BY RAND() limit ?`
	err = db.Select(&articleList, sqlStr, categoryId, num)
	return
}

//	根据文章id修改文章信息
func UpdateArticleInfo(article *model.ArticleDetail) (row interface{}, err error) {
	sqlStr := `update article set content=?, summary=?, title=?, category_id=?, open_flag=? where id = ?`
	row, err = db.Exec(sqlStr, article.Content, article.Summary, article.Title, article.CategoryId, article.OpenFlag, article.Id)
	return
}

//	根据文章id删除文章信息
func DeleteArticle(articleId int64) (row interface{}, err error) {
	sqlStr := `delete from article where id = ?`
	row, err = db.Exec(sqlStr, articleId)
	return
}

//	修改文章预览数量
func UpdateViewCount(article *model.ArticleDetail) (row interface{}, err error) {
	if article.Id != 0 {
		article.ViewCount++
	}
	sqlStr := `update article set view_count=? where id = ?`
	row, err = db.Exec(sqlStr, article.ViewCount, article.Id)
	return
}
