package flag

import sys_flag "flag"

type Option struct {
	DB bool
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	sys_flag.Parse()
	return Option{
		DB: *db,
	}
}

func IsWebStop(option Option) bool {
	return option.DB
}

func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
	}
}
