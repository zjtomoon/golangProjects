package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("tem/*") //加载模板文件
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{"title":"我是测试","ce":"123456"})
	})
	r.Run()
}
