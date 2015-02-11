package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Brand struct {
	Id                int
	Name              string `orm:"unique"`
	Brand_number      int    `orm:"unique"`
	Domain_limit      int    `orm:"default(0)"`
	Call_limit        int    `orm:"default(0)"`
	Reg_limit         int    `orm:"default(0)"`
	Active_reg        int    `orm:"default(0)"`
	Active_call       int    `orm:"default(0)"`
	Signal_encryption int    `orm:"default(0)"`
	Media_encryption  int    `orm:"default(0)"`
	Is_active         int    `orm:"default(0)"`
	Brand_header      string
	Brand_footer      string
	Balance_url       string
	Deleted           bool      `orm:"default(false)"`
	Created_at        time.Time `orm:"auto_now_add;type(datetime)"`
	Updated_at        time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Brand))
}
