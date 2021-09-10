package db

import (
	_ "github.com/go-sql-driver/mysql"
	"golangProjects/my-demos/gin-blog/model"
)

//插入文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	//加个验证
	if article == nil {
		return
	}
	sqlstr := `insert into 
			article(content,summary,title,username,category_id,view_count,comment_count) values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlstr, article.Content, article.Summary, article.Title, article.Username,
		article.ArticleInfo.CategoryId, article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

//获取文章列表，做个分页
