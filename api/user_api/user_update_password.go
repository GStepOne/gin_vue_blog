package user_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/utils"
	"blog/gin/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Password       string `json:"password" binding:"required" msg:"用户密码必填"`
	OriginPassword string `json:"origin_password" binding:"required" msg:"用户原密码必填"`
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	var cr UserRequest

	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var user models.UserModel

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	//查询用户是否存在
	err := global.DB.Take(&user, claims.UserId).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	//验证密码是否一样
	fmt.Println("原来密码", user.Password)
	if !utils.CheckPwd(user.Password, cr.OriginPassword) {
		res.FailWithMessage("原来密码不正确", c)
		return
	}
	hashPwd := utils.HashPwd(cr.Password)

	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error("修改密码错误", err)
		res.OKWithMessage("修改密码错误", c)
		return
	}
	res.OKWithMessage("密码修改成功", c)
	return
}
