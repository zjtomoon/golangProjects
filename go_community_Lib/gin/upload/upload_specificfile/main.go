package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//有的用户上传文件需要限制上传文件的类型以及上传文件的大小，但是gin框架暂时没有这些函数(也有可能是我没找到)，
// 因此基于原生的函数写法自己写了一个可以限制大小以及文件类型的上传函数
func main() {
	r := gin.Default()
	r.POST("upload", func(c *gin.Context) {
		_,headers,err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("err when try to get file : %v",err)
		}
		//headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Context-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers,"./video/"+headers.Filename)
		c.String(http.StatusOK,headers.Filename)
	})
	r.Run()
}