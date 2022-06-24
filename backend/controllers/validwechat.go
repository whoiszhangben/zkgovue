package controllers

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/astaxie/beego"
)

const (
	token = "whoiszhangben"
)

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

// ValidWechatController ...
type ValidWechatController struct {
	beego.Controller
}

// Get ...
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (c *ValidWechatController) Get() {
	// 获取到请求的参数
	fmt.Println("当前请求的URI:", c.Ctx.Request.RequestURI)
	fmt.Println("signature:", c.Ctx.Input.Query("signature"))
	fmt.Println("timestamp:", c.Ctx.Input.Query("timestamp"))
	fmt.Println("nonce:", c.Ctx.Input.Query("nonce"))
	fmt.Println("echostr:", c.Ctx.Input.Query("echostr"))
	c.Ctx.WriteString(c.Ctx.Input.Query("echostr"))
}
