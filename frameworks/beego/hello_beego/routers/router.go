package routers

import (
    "github.com/astaxie/beego"
    "hello_beego/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
