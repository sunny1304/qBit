package routers

import (
	"github.com/astaxie/beego"
	"qbit/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/brand/new", new(controllers.BrandController), "get:Create_brand")
	beego.Router("/brand/save", new(controllers.BrandController), "post:Save_brand")
}
