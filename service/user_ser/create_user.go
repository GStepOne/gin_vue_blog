package user_ser

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/utils"
	"errors"
	"fmt"
)

const Avatar = "/uploads/avatar/default_avatar.jpg"

func (UserService) CreateUser(username, nickname, password string, role ctype.Role, email string, ip string) error {
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", username).Error
	if err == nil {
		//存在了
		fmt.Println(err)
		global.Log.Error("用户名已经存在，请重新输入")
		return errors.New("用户名已经存在")
	}

	// 对密码hash
	hashPwd := utils.HashPwd(password)
	//头像
	//avatar := "/uploads/avatar/default_avatar.jpg"

	err = global.DB.Create(&models.UserModel{
		UserName:   username,
		NickName:   nickname,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		SignStatus: ctype.SignEmail,
		IP:         ip,
		Addr:       "内网",
		Avatar:     Avatar,
	}).Error

	if err != nil {
		global.Log.Error(err)
		return errors.New("创建失败")
	}

	global.Log.Infof("用户%s创建成功", username)
	return nil
}
