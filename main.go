package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "qbit/models"
	_ "qbit/routers"
)

func init() {
	/* Read the config like Rails */

	db_user := beego.AppConfig.String("db_username")
	db_passwd := beego.AppConfig.String("db_password")
	db_dbname := beego.AppConfig.String("db_database")
	db_charset := beego.AppConfig.String("db_charset")
	if db_charset == "" {
		db_charset = "utf8"
	}
	db_url := fmt.Sprintf("%s:%s@/%s?%s", db_user, db_passwd, db_dbname, db_charset)
	beego.Info(db_url)
	/*===============================================================*/

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", db_url)
	orm.Debug = true // for test env

	// alias db name
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true
	beego.Info(beego.AppConfig.String("db"))

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error(err)
	}

}
func main() {
	beego.Run()
}
