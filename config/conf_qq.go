package config

import "fmt"

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"`
	AppKey   string `json:"app_key" yaml:"app_key"`
	Redirect string `json:"redirect" yaml:"redirect"`
}

func (q QQ) GetPath() string {

	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
