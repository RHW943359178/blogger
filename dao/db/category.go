package db

import (
	"blogger/model"
	"github.com/jmoiron/sqlx"
	"log"
)

//	添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := `insert into category(category_name, category_no) value(?, ?)`
	result, err := db.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		log.Fatal("exec sql failed, err: ", err)
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

//	获取单个文章分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := `select id, category_name, category_no from category where id = ?`
	err = db.Get(category, sqlStr, id)
	if err != nil {
		log.Fatalln("get data failed, err: ", err)
		return
	}
	return
}

//	获取多个文章分类
func GetCategoryList(categoryIds []int64) (categories []*model.Category, err error) {
	//	构建 sql 语句
	log.Println("categoryIds: ", categoryIds)
	sqlStr, args, err := sqlx.In(`select id, category_name, category_no from category where id in(?)`, categoryIds)
	if err != nil {
		log.Fatalln("sqlx in failed, err: ", err)
		return
	}
	//	语句
	err = db.Select(&categories, sqlStr, args...)
	if err != nil {
		return
	}
	return
}

//	获取所有文章分类
func GetAllCategory() (categories []*model.Category, err error) {
	sqlStr := `select id, category_name, category_no from category order by category_no asc`
	err = db.Select(&categories, sqlStr)
	if err != nil {
		log.Fatalln("get data failed, err: ", err)
		return
	}
	return
}
