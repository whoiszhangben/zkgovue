package main

import (
	"backend/base"
	_ "backend/routers"
	"fmt"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	//"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func main() {
	err := base.Init_log()
	Logs := base.Logs
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//	Logs.Info("program start .\n")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	enc := beego.AppConfig.String("enc_mobile")
	Logs.Info("enc_mobile %s\n", enc)
	ver := beego.AppConfig.String("version")
	Logs.Info("version %s\n", ver)

	appSecret := "72c02927220db6980226c6f9b4001d90"
	encryptSecret, errd := base.EncodeDsn(appSecret)

	fmt.Println("加密后的appSecret:", encryptSecret)

	dsn := beego.AppConfig.String("dsn")
	dedsn, errd := base.DecodeDsn(dsn)
	fmt.Println(dedsn)
	if errd != nil {
		Logs.Error("decode: dsn failed! %s\n", errd)
		return
	}

	initDatabase("default", beego.AppConfig.String("driver"), dedsn)
	Logs.Info("init database default success.\n")
	go base.SysSignalHandleDeamon()
	ignoreStaticPath()
	beego.Run()
	//Logs.CloseLog()
}

func initDatabase(dbname string, driver string, dsn string) {
	//注册数据库
	orm.RegisterDataBase(dbname, driver, dsn)
}

func ignoreStaticPath() {
	//pattern 路由规则，可以根据一定的规则进行路由，如果你全匹配可以用"*"
	// beego.InsertFilter("*",beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url:", orpath)
	//如果请求url还有api字段，说明指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	if strings.Index(orpath, "test") >= 0 {

		return
	}
	fmt.Println("url地址:", "views"+ctx.Request.URL.Path)
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "views"+ctx.Request.URL.Path)
}
