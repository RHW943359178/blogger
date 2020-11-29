package db

import (
	"blogger/model"
	"testing"
	"time"
)

func init() {
	//	parseTime=true 将mysql中时间类型，自动解析为 go 结构体中的时间类型
	//	不加报错
	dns := "RHW:RHW943359178@tcp(81.69.255.188:3306)/blogger?parseTime=True"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

//	测试插入文章
func TestInsertArticle(t *testing.T) {
	//	构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 1
	article.Content = "ahsawosxjawpal this is text this is a text this is a text"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "go practice"
	article.ArticleInfo.Username = "rhw"
	article.ArticleInfo.Summary = "ahsawosxjawpal"
	article.ArticleInfo.ViewCount = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("articleId: %d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList("5qi", []string{"1"}, 0, 10)
	if err != nil {
		t.Log("err: ", err)
		return
	}
	t.Logf("article: %d\n", len(articleList))
}

//	根据文章id，查询单个文章
func TestGetArticleDetail(t *testing.T) {
	detail, err := GetArticleDetail(8)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("detail: %#v\n", detail)
}
