package model

//	`id`, `category_name `, `category_no`, `create_time`, `update_time`

//	定义分类结构体
type Category struct {
	CategoryId   int64  `db:"category_id" json:"categoryId"`
	CategoryName string `db:"category_name" json:"categoryName"`
	CategoryNo   int    `db:"category_no" json:"categoryNo"`
	Color        string `db:"color" json:"color"`
}
