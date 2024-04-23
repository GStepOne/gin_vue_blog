package flag

import (
	sys_flag "flag"
	"fmt"
	structs2 "github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string //- u admin -u user
	ES   string //-es create -es delete 删除索引
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	es := sys_flag.String("es", "", "elastic操作")
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

func IsWebStop(option Option) (f bool) {
	maps := structs2.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

func IsCreateUser(option Option) string {
	return option.User
}

func IsCreateIndex(option Option) string {
	return option.ES
}

func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
		return
	}
	fmt.Println(option.User, option.User == "")

	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}

	if option.ES == "create" {
		EsCreateIndex()
		return
	}

	if option.ES == "delete" {
		//EsCreateIndex()
		return
	}

	sys_flag.Usage()
}
