package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"qbit/models"
	_ "strings"
)

type BrandController struct {
	beego.Controller
}

func (this *BrandController) Create_brand() {
	this.TplNames = "brand/add_brand.tpl"
}

func (this *BrandController) Save_brand() {
	if this.Ctx.Request.Method == "POST" {
		brand_name := this.GetString("b_name")
		brand := new(models.Brand)
		brand.Name = brand_name

		o := orm.NewOrm()
		_, err := o.Insert(brand)

		if err == nil {
			this.Redirect("/", 302)
			return
		} else {
			this.Redirect("/brand/new", 302)
		}

	}
}
