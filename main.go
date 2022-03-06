package main

import (
	"fmt"
	_ "ssoer/routers"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
)

var DbName = "ssoer"
var DbUser = "ssoer"
var DbPassword = "ssoer"
var DbHost = "localhost"
var DbPort = "5432"

func init() {
	var dbConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		DbUser, DbPassword, DbHost, DbPort, DbName,
	)

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dbConnectionString)

	// orm.RegisterDataBase("default", "postgres",
	// 	fmt.Sprintf("user=%s "+
	// 		"password=%s "+
	// 		"host=%s "+
	// 		"port=%s "+
	// 		"dbname=%s "+
	// 		"sslmode=disable",
	// 		DbUser, DbPassword, DbHost, DbPort, DbName,
	// 	),
	// )
	orm.Debug = true
}

func main() {
	fmt.Println("-------------------------------------------------")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
