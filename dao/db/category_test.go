package db

import "testing"

func init() {
	//	parseTime=true 将mysql中时间类型，自动解析为 go 结构体中的时间类型
	//	不加报错
	dns := "RHW:RHW943359178@tcp(81.69.255.188:3306)/blogger?parseTime=True"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

//	获取单个分类信息
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category: %#v\n", category)
}

//	获取多个分类信息
func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _, category := range categoryList {
		t.Logf("id: %d, category: %#v\n", category.CategoryId, category)
	}
}

//	获取全部分类
func TestGetAllCategory(t *testing.T) {
	categoryList, err := GetAllCategory()
	if err != nil {
		panic(err)
	}
	for _, category := range categoryList {
		t.Logf("id: %d, category: %#v\n", category.CategoryId, category)
	}
}
