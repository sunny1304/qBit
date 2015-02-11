package routers

import (
	"github.com/astaxie/beego"
	"qbit/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", new(controllers.UserController), "get:Authuser")
}
