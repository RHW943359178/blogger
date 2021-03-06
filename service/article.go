package service

import (
	"blogger/dao/db"
	"blogger/model"
	"log"
)

//	获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	////	1.获取文章列表
	//articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
	//if len(articleInfoList) <= 0 {
	//	return
	//}
	////	2.获取文章对应的分类 (多个)
	//categoryIds := getCategoryIds(articleInfoList)
	//categoryList, err := db.GetCategoryList(categoryIds)
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
	////	返回页面，做聚合
	////	遍历所有文章
	//for _, article := range articleInfoList {
	//	//	根据当前文章，生成结构体
	//	articleRecord := &model.ArticleRecord{
	//		ArticleInfo: *article,
	//	}
	//	//	文章取出分类id
	//	categoryId := article.CategoryId
	//	//	遍历分类列表
	//	for _, category := range categoryList {
	//		if categoryId == category.CategoryId {
	//			articleRecord.Category = *category
	//			break
	//		}
	//	}
	//	articleRecordList = append(articleRecordList, articleRecord)
	//}
	return
}

//	根据多个文章的id， 获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
	//	遍历文章，得到每个文章
	for _, article := range articleInfoList {
		//	从当前文章取出分类id
		categoryId := article.CategoryId
		//	去重， 防止重复
		for _, id := range ids {
			//	看当前id是否存在
			if id != categoryId {
				ids = append(ids, categoryId)
			}
		}
	}
	return
}

//	根据分类id， 获取该类文章和他们对应的分类信息
func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//	1.获取文章列表
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("articleInfoList: ", articleInfoList)
	if len(articleInfoList) <= 0 {
		return
	}
	//	2.获取文章对应的分类 (多个)
	categoryIds := getCategoryIds(articleInfoList)
	log.Println("categoryIds123: ", categoryIds)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		log.Fatalln(err)
		return
	}
	//	返回页面，做聚合
	//	遍历所有文章
	for _, article := range articleInfoList {
		//	根据当前文章，生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		//	文章取出分类id
		categoryId := article.CategoryId
		//	遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

//	根据条件获取相应的文章列表（分类，页码）
func GetArticleListByCondition(condition string, categoryId []string, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	//	数据库的顺序从0开始
	pageNum -= 1
	//	从 dao 层取出数据
	articleList, err = db.GetArticleList(condition, categoryId, pageNum, pageSize)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

//	保存用户的文章信息
func ArticleSave(article *model.ArticleDetail) (articleId int64, err error) {
	articleId, err = db.InsertArticle(article)
	if err != nil {
		log.Fatalln("insert article failed", err)
		return
	}
	return
}

//	根据文章id获取单个文章信息
func GetArticleInfoById(articleId int64) (article *model.ArticleDetail, err error, errFlag int) {
	//	错误标志
	errFlag = 0
	//	从数据库获取数据
	article, err = db.GetArticleDetail(articleId)
	if err != nil {
		log.Println("get data from database failed, err: ", err)
		return nil, err, 0
	}
	//	每次返回成功就修改一次数据的view_count的状态
	_, err = db.UpdateViewCount(article)
	if err != nil {
		errFlag = 1
		log.Fatalln("update article view_count failed, err: ", err)
	}
	return
}

//	根据用户 id 获取全部文章信息
func GetArticleListByUserId(userId string, pageSize, pageNum int) (articleList []*model.UserArticle, err error) {
	//	从数据库取数据
	articleList, err = db.GetArticleByUserId(userId, pageSize, pageNum)
	//	如果从数据库取出的数据为空，也因该返回给前端空的 切片 而不是 null
	if articleList == nil {
		articleList = []*model.UserArticle{}
	}
	log.Println(articleList, 123)
	if err != nil {
		log.Println("get data from database failed, err: ", err)
	}
	return
}

//	根据用户 id 修改文章信息
func UpdateArticleInfo(article *model.ArticleDetail) (row interface{}, err error) {
	//	从 db 层取数据
	row, err = db.UpdateArticleInfo(article)
	if err != nil {
		log.Fatalln("get data from database failed, err: ", err)
		return
	}
	return
}

//	根据用户 id 删除文章信息
func DeleteArticle(article *model.ArticleDetail) (row interface{}, err error) {
	var articleId = article.Id
	//	从 db 层取数据
	row, err = db.DeleteArticle(articleId)
	if err != nil {
		log.Fatalln("get data from db failed, err: ", err)
		return
	}
	return
}

//	根据userId获取其他文章
func GetOtherArticle(userId string, articleId int64, pageSize, pageNum int) (articleList []*model.UserArticle, err error) {
	//	从db层取数据
	articleList, err = db.GetOtherArticle(userId, articleId, pageSize, pageNum)
	if err != nil {
		log.Println("get other from db failed, err: ", err)
	}
	return
}

//	根据 categoryId 获取推荐文章
func GetRecommendArticle(categoryId int64, num int) (articleList []*model.UserArticle, err error) {
	//	验证 num 参数
	if num < 1 || num > 10 {
		return
	}
	articleList, err = db.GetRecommendArticle(categoryId, num)
	if err != nil {
		log.Println("get recommend article from db failed, err: ", err)
	}
	return
}
