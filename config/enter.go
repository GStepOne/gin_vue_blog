package config

type config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Mysql  `yaml:"logger"`
	System System `yaml:"system"`
}
