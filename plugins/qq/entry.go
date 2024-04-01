package qq

import (
	"blog/gin/global"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type QQInfo struct {
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Avatar   string `json:"figure_url_qq"` //头像大图
	OpenId   string `json:"open_id"`
}

type QQLogin struct {
	appId       string
	appKey      string
	redirect    string
	code        string
	accessToken string
	openId      string
}

func NewQQLogin(code string) (qqInfo QQInfo, err error) {
	qqLogin := &QQLogin{
		appId:    global.Config.QQ.AppID,
		appKey:   global.Config.QQ.Key,
		redirect: global.Config.QQ.Redirect,
		code:     code,
	}

	err = qqLogin.GetAccessToken()
	if err != nil {
		return qqInfo, err
	}

	err = qqLogin.GetOpenID()

	if err != nil {
		return qqInfo, err
	}

	qqInfo, err = qqLogin.GetUserInfo()
	if err != nil {
		return qqInfo, err
	}

	qqInfo.OpenId = qqLogin.openId
	return qqInfo, nil
}

func (q *QQLogin) GetAccessToken() error {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_qq", q.appId)
	params.Add("client_secret", q.appKey)
	params.Add("code", q.code)
	params.Add("redirect_url", q.redirect)

	u, err := url.Parse("https://graph.qq.com/oauth2.0/token")
	if err != nil {
		return err
	}
	u.RawQuery = params.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	qs, err := parseQs(res.Body)
	if err != nil {
		return err
	}

	q.accessToken = qs[`access_token`][0]
	return nil

}

func (q *QQLogin) GetOpenID() error {
	u, err := url.Parse(fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s", q.accessToken))
	if err != nil {
		return err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}

	defer res.Body.Close()
	openId, err := getOpenId(res.Body)
	if err != nil {
		return err
	}
	q.openId = openId
	return nil
}

func (q *QQLogin) GetUserInfo() (qqInfo QQInfo, err error) {
	params := url.Values{}
	params.Add("access_token", q.accessToken)
	params.Add("oauth_consumer_key", q.appId)
	params.Add("openid", q.openId)
	u, err := url.Parse("https://graph.qq.com/user/get_user_info")
	if err != nil {
		return qqInfo, err
	}
	u.RawQuery = params.Encode()
	res, err := http.Get(u.String())
	data, err := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &qqInfo)
	defer res.Body.Close()
	if err != nil {
		return qqInfo, err
	}

	return qqInfo, nil
}

func parseQs(r io.Reader) (val map[string][]string, err error) {
	val, err = url.ParseQuery(readAll(r))
	if err != nil {
		return val, err
	}
	return val, nil
}

func getOpenId(r io.Reader) (string, error) {
	body := readAll(r)
	start := strings.Index(body, `"openid":"`) + len(`"openid":"`)
	if start == -1 {
		return "", fmt.Errorf("openid not found")
	}
	end := strings.Index(body[start:], `"`)
	if end == -1 {
		return "", fmt.Errorf("openid not found")
	}
	return body[start : start+end], nil

}

func readAll(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
