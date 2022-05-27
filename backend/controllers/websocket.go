package controllers

import (
	"backend/models"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

// Operations about object
type WebsocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

//web  全局变量
var Hub = models.NewHub()

// Get ...
// 方法为所在请求变量
func (this *WebsocketController) Get() {
	go Hub.Run()
	conn, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)

	if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	client := &models.Client{Hub: Hub, Conn: conn, Send: make(chan []byte, models.MaxMessageSize)}
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
