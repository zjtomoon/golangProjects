package main

import "github.com/gin-gonic/gin"

//routes group是为了管理一些相同的URL

func main() {
	//1.创建路由
	//默认使用了2个中间件Logger(),Recovery()
	r := gin.Default()
	//路由组1 ，处理GET请求

}
