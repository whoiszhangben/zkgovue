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
		beego.NSNamespace("/valid",
			beego.NSInclude(
				&controllers.ValidWechatController{},
			)),
		beego.NSNamespace("/upload",
			beego.NSInclude(
				&controllers.UploadController{},
			)))
	ts := beego.NewNamespace("/api",
		beego.NSNamespace("/valid",
			beego.NSInclude(
				&controllers.ValidWechatController{},
			)))
	beego.AddNamespace(ns)
	beego.AddNamespace(ts)
	beego.Router("/api/websocket", &controllers.WebsocketController{})
	beego.Router("/api/callback", &controllers.CallbackController{})
	beego.Router("/api/redirect", &controllers.RedirectController{})
}
