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

//	ResponseArticleSave	保存文章信息到数据库
type ResponseArticle struct {
	Code    int    `json:"code"`    //	业务响应状态码
	Message string `json:"message"` //	提示信息
	Data    int64  `json:"data"`    //	数据
}

//	ResponseGetSingleArticle 根据id获取单个文章信息
type ResponseGetSingleArticle struct {
	Code    int                  `json:"code"`    //	业务响应状态码
	Message string               `json:"message"` //	提示信息
	Data    *model.ArticleDetail `json:"data"`    //	返回数据
}

//	ResponseUserInfo 用户校验返回用户信息
type ResponseUserInfo struct {
	Code    int            `json:"code"`    //	业务响应状态码
	Message string         `json:"message"` //	提示信息
	Data    *model.ResUser `json:"data"`    //	返回数据
}

//	ResponseUserArticle	返回这个用户所以的文章
type ResponseUserArticle struct {
	Code    int                  `json:"code"`    //	业务响应状态码
	Message string               `json:"message"` //	提示信息
	Data    []*model.UserArticle `json:"data"`    //	返回数据
}
