package models

import (
	"github.com/astaxie/beego/orm"
)

type Brand struct {
	Id                int
	Name              string
	Brand_number      int
	Domain_limit      int
	Call_limit        int
	Reg_limit         int
	Active_reg        int
	Active_call       int
	Signal_encryption int
	Media_encryption  int
	Is_active         bool
	Brand_header      string
	Brand_footer      string
	Balance_url       string
}

func init() {
	orm.RegisterModel(new(Brand))
}
