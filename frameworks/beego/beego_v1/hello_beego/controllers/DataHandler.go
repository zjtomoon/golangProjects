package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type user struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}

type ObjectController struct {
	beego.Controller
}

//获取Request body的内容

func (this *ObjectController) Post() {
}

//文件上传

type FormController struct {
	beego.Controller
}

func (c *FormController) Post() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err", err)
	}
	defer f.Close()
	c.SaveToFile("upload", "static/upload/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
}
