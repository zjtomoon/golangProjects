package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/beego/i18n"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//c.Ctx.WriteString("hello")

	//session控制
	v := c.GetSession("asta")
	if v == nil {
		c.SetSession("asta", int(1))
		c.Data["num"] = 0
	} else {
		c.SetSession("asta", v.(int)+1)
		c.Data["num"] = v.(int)
	}

	//Flash数据
	//它主要用于在两个逻辑间传递临时数据，flash 中存放的所有数据会在紧接着的下一个逻辑中调用后清除。
	// 一般用于传递提示和错误消息。它适合 Post/Redirect/Get 模式
	flash := beego.ReadFromRequest(&c.Controller)
	if _, ok := flash.Data["notice"]; ok {
		c.TplName = "set_success.html"
	} else if _, ok = flash.Data["error"]; ok {
		c.TplName = "set_error.html"
	} else {
		//c.Data["list"] = GetInfo()
		c.TplName = "setting_list.html"
	}
}

//获取参数

func (this *MainController) Post() {
	jsoninfo := this.GetString("jsoninfo")
	if jsoninfo == "" {
		this.Ctx.WriteString("jsoninfo is empty")
		return
	}

	//如果你需要的数据可能是其他类型的，例如是 int 类型而不是 int64，那么你需要这样处理：
	//id := this.Input().Get("id")
	//intid, err := strconv.Atoi(id)

	u := user{}
	if err := this.ParseForm(&u); err != nil {

	}

}

type AddController struct {
	beego.Controller
}

func (this *AddController) Prepare() {

}

func (this *AddController) Get() {
	this.Data["content"] = "value"
	this.Layout = "admin/layout.html"
	this.TplName = "admin/add.tpl"
}

func (this *AddController) Post() {

}

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	beego.Controller
}

type BaseAdminRouter struct {
}

//提前终止运行

type RController struct {
	beego.Controller
}

func (this *RController) Prepare() {
	this.Data["json"] = map[string]interface{}{"name": "astaxie"}
	this.ServeJSON()
	this.StopRun()
}
