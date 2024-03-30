package flag

import (
	sys_flag "flag"
	"fmt"
)

type Option struct {
	DB   bool
	User string //- u admin -u user
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

func IsWebStop(option Option) bool {
	return option.DB
}

func IsCreateUser(option Option) string {
	return option.User
}
func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
	}
	fmt.Println(option.User, option.User == "")

	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}

	sys_flag.Usage()
}
