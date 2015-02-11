package models

import (
	"github.com/astaxie/beego/orm"
)

type Brand_domain struct {
	Id        int
	Brand_id  int
	Domain_id int
}

func init() {
	orm.RegisterModel(new(Brand_domain))
}
