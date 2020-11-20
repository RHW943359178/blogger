package controller

import "blogger/model"

// _ResponseCategoryList 全部分类接口相应数据
type ResponseCategoryList struct {
	Code    int               `json:"code"`    // 业务响应状态码
	Message string            `json:"message"` // 提示信息
	Data    []*model.Category `json:"data"`    // 数据
}

//	ResponseArticleList 根据关键字，分类id，页码数，页码大小获取相应的文章数据
type ResponseArticleList struct {
	Code    int                  `json:"code"`    // 业务响应状态码
	Message string               `json:"message"` // 提示信息
	Data    []*model.ArticleInfo `json:"data"`    // 数据
}