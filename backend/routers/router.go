package routers

import (
	"backend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/info",
			beego.NSInclude(
				&controllers.InfoController{},
			)),
		// beego.NSNamespace("/websocket",
		// 	beego.NSInclude(
		// 		&controllers.WebsocketController{},
		// 	)),
		beego.NSNamespace("/upload",
			beego.NSInclude(
				&controllers.UploadController{},
			)))
	beego.AddNamespace(ns)
	beego.Router("/api/websocket", &controllers.WebsocketController{})
}
