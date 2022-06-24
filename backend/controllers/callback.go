package controllers

import (
	"backend/utils"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// CallbackController ...
type CallbackController struct {
	beego.Controller
}

// ResAccessToken 获取用户授权access_token的返回结果
type ResAccessToken struct {
	utils.CommonError

	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

//UserInfo 用户授权获取到用户信息
type UserInfo struct {
	utils.CommonError

	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

const (
	redirectOauthURL      = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	accessTokenURL        = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	refreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	userInfoURL           = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	checkAccessTokenURL   = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
	AppID                 = "wx7c3a1cbaa1be0098"
	AppSecret             = "72c02927220db6980226c6f9b4001d90"
)

// Get ...
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (c *CallbackController) Get() {
	code := c.Ctx.Input.Query("code")
	fmt.Println("临时授权码code:", code)
	// 根据code获取access_token
	accessToken, err := utils.GetUserAccessToken(code)
	if err != nil {
		return
	}
	user, err := utils.GetUserInfo(accessToken.AccessToken, accessToken.OpenID)
	if err != nil {
		return
	}
	c.Data["json"] = user
	c.ServeJSON()

	// token := utils.GetAccessToken()
	// c.Data["json"] = token
	// c.ServeJSON()
}

//CheckAccessToken 检验access_token是否有效
func (c *CallbackController) CheckAccessToken(accessToken, openID string) (b bool, err error) {
	urlStr := fmt.Sprintf(checkAccessTokenURL, accessToken, openID)
	var response []byte
	response, err = utils.HTTPGet(urlStr)
	if err != nil {
		return
	}
	var result utils.CommonError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		b = false
		return
	}
	b = true
	return
}
