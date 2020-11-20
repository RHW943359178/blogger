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
