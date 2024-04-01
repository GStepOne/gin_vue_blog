package flag

import (
	"blog/gin/global"
	"blog/gin/models/ctype"
	"blog/gin/service/user_ser"
	"fmt"
)

func CreateUser(permissions string) {
	//创建用户的逻辑
	// 用户名 昵称 密码 确认密码  邮箱
	var (
		username        string
		nickname        string
		password        string
		confirmPassword string
		email           string
	)

	fmt.Printf("创建用户请输入用户名:")
	fmt.Scan(&username)
	fmt.Printf("创建用户请输入昵称:")
	fmt.Scan(&nickname)
	fmt.Printf("创建用户请输入密码:")
	fmt.Scan(&password)
	fmt.Printf("创建用户再次请输入密码:")
	fmt.Scan(&confirmPassword)
	fmt.Printf("创建用户请输入邮箱:")
	fmt.Scan(&email)
	fmt.Printf("创建用户请输如角色:")
	fmt.Scanln(&permissions)

	//校验两次密码
	if password != confirmPassword {
		global.Log.Error("两次密码不一致，请重新输入")
		return
	}

	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	err := user_ser.UserService{}.CreateUser(username, nickname, password, role, email, "127.0.0.1")
	if err != nil {
		//存在了
		fmt.Println(err)
		global.Log.Error("创建用户失败", err)
		return
	}

	fmt.Println(username, nickname, password, confirmPassword, email)

	global.Log.Infof("用户%s创建成功", username)
}
