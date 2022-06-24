package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

const (
	oAuthURL    = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	redirectURL = "http://dhloa.cloudpivot.cn/wechat/api/callback"
	appId       = "wx7c3a1cbaa1be0098"
	chkType     = "snsapi_userinfo"
)

// RedirectController ...
type RedirectController struct {
	beego.Controller
}

// Get ...
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (c *RedirectController) Get() {
	urlStr := fmt.Sprintf(oAuthURL, appId, redirectURL, chkType, "aaaa")
	fmt.Println(urlStr)
	c.Redirect(urlStr, 302)
}
