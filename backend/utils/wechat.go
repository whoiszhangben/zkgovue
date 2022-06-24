package utils

import (
	"backend/base"

	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	redirectOauthURL      = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	accessTokenURL        = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	refreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	userInfoURL           = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	checkAccessTokenURL   = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
	baseAccessToken       = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

// AccessToken 获取公众号的accessToken
type AccessToken struct {
	CommonError

	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

// WebAuthAccessToken 获取用户授权access_token的返回结果
type WebAuthAccessToken struct {
	CommonError

	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

//UserInfo 用户授权获取到用户信息
type UserInfo struct {
	CommonError

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

func GetAccessToken() AccessToken {
	var res = AccessToken{}
	appId := beego.AppConfig.String("appId")
	appSecret, err := base.DecodeDsn(beego.AppConfig.String("appSecret"))
	fmt.Println(appId, appSecret)
	if err != nil {
		fmt.Println("解密appSecret失败", err)
		return res
	}
	accessTokenURLFormat := fmt.Sprintf(baseAccessToken, appId, appSecret)
	binaryRes, err := HTTPGet(accessTokenURLFormat)
	if err != nil {
		fmt.Println("获取token失败", err)
		return res
	}
	err = json.Unmarshal(binaryRes, &res)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return res
	}
	return res
}

func GetUserAccessToken(code string) (result WebAuthAccessToken, err error) {
	var res = WebAuthAccessToken{}
	appId := beego.AppConfig.String("appId")
	appSecret, err := base.DecodeDsn(beego.AppConfig.String("appSecret"))
	if err != nil {
		fmt.Println("解密appSecret失败", err)
		return res, err
	}
	accessTokenURLFormat := fmt.Sprintf(accessTokenURL, appId, appSecret, code)
	binaryRes, err := HTTPGet(accessTokenURLFormat)
	if err != nil {
		fmt.Println("获取token失败", err)
		return res, err
	}
	err = json.Unmarshal(binaryRes, &res)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return res, err
	}
	return res, nil
}

func GetUserInfo(accessToken string, openId string) (UserInfo, error) {
	var res = UserInfo{}
	userinfoURLFormat := fmt.Sprintf(userInfoURL, accessToken, openId)
	binaryRes, err := HTTPGet(userinfoURLFormat)
	if err != nil {
		fmt.Println("获取用户信息失败", err)
		return res, err
	}
	err = json.Unmarshal(binaryRes, &res)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return res, err
	}
	return res, nil
}
