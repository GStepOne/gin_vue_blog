package config

type Config struct {
	Mysql     Mysql     `yaml:"mysql"`
	Logger    Logger    `yaml:"logger"`
	System    System    `yaml:"system"`
	SiteInfo  SiteInfo  `yaml:"site_info"`
	QQ        QQ        `yaml:"qq"`
	JWT       JWT       `yaml:"jwt"`
	QiNiu     QiNiu     `yaml:"qi_niu"`
	Email     Email     `yaml:"email"`
	Upload    Upload    `yaml:"upload"`
	Redis     Redis     `yaml:"redis"`
	Es        Es        `yaml:"es"`
	ChatGroup ChatGroup `yaml:"chat_group"`
}
