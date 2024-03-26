package config

type QiNiu struct {
	AccessKey string `mapstructure:"access_key" json:"access_key" yaml:"access_key"`
	SecretKey string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	Bucket    string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	CDN       string `mapstructure:"cdn" json:"cdn" yaml:"cdn"`    //访问图片的地址的前缀
	Zone      string `mapstructure:"zone" json:"zone" yaml:"zone"` //存储的地区  华东 华北
	Size      string `mapstructure:"size" json:"size" yaml:"size"` //
}
