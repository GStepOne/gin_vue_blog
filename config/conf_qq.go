package config

type QQ struct {
	AppID    string `json:"app-id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"`
}

func (q QQ) GetPath() string {
	return ""
}
