package main

import (
	_ "github.com/udistrital/resoluciones_crud_v2/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	auditoria "github.com/udistrital/utils_oas/auditoria"
	"github.com/udistrital/utils_oas/customerrorv2"
	"github.com/udistrital/utils_oas/xray"
)

func main() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+":"+
		beego.AppConfig.String("PGpass")+"@"+
		beego.AppConfig.String("PGhost")+":"+
		beego.AppConfig.String("PGport")+"/"+
		beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschema")+"")

	AllowedOrigins := []string{"*.udistrital.edu.co"}
	if beego.BConfig.RunMode == "dev" {
		AllowedOrigins = []string{"*"}
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: AllowedOrigins,
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	xray.InitXRay()
	beego.ErrorController(&customerrorv2.CustomErrorController{})
	apistatus.Init()
	auditoria.InitMiddleware()
	beego.Run()
}
