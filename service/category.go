package service

import (
	"blogger/dao/db"
	"blogger/model"
	"log"
)

//	获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategory()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}
