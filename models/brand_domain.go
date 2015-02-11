package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Brand_domain struct {
	Id         int
	Brand_id   int
	Domain_id  int
	Created_at time.Time `orm:"auto_now_add; type(datetime)"`
	Updated_at time.Time `orm: "auto_now; type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Brand_domain))
}
