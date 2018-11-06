package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "suanli/routers"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	host := beego.AppConfig.String("db_host")
	port := beego.AppConfig.String("db_port")
	username := beego.AppConfig.String("db_user")
	password := beego.AppConfig.String("db_pass")
	dbName := beego.AppConfig.String("db_name")
	orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbName+"?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)
	//orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

func main() {
	beego.Run()
}

