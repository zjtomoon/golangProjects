package db

import (
	"gin-blog/model"
	"testing"
	"time"
)

func init() {
	// parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	// 不加报错
	dns := "root:123@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

//测插入文章
func TestInsertArticle(t *testing.T) {
	//构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "abc fdlksafjdlajflk fdjlasjfdkljwa"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "5qi go"
	article.ArticleInfo.Username = "sun"
	article.ArticleInfo.Summary = "abc fd"
	article.ArticleInfo.ViewCount = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		return
	}
	t.Logf("articleId : %d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		return
	}
	t.Logf("article : %d\n", len(articleList))
}

func TestGetArticleDetail(t *testing.T) {
	arr, err := GetArticleDetail(1)
	if err != nil {
		return
	}
	t.Logf("article %#v\n", arr)
}
