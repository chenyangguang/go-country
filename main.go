package main

import (
	_ "github.com/chenyangguang/go-country/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// init ...
func init() {
	maxIdle := 20
	maxConn := 20
	rdbErr := orm.RegisterDataBase("default", "mysql", "country:countrY3334#@/go_db?charset=utf8", maxIdle, maxConn)
	if rdbErr != nil {
		beego.Error(rdbErr)
	}

	orm.Debug = true
	err := orm.RunSyncdb("default", false, true) // autocreate table
	if err != nil {
		beego.Error(err)
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
