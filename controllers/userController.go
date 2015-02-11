package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type UserController struct {
	beego.Controller
}

type Demouser struct {
	Name  string
	Email string
}

func (this *UserController) Authuser() {
	req_uri := this.Ctx.Request.RequestURI
	params := strings.Split(req_uri, "=")
	db_data := strings.Split(params[1], "|")
	beego.Info(db_data)
	user := Demouser{db_data[0], db_data[1]}
	this.Data["json"] = &user
	this.ServeJson()

}
