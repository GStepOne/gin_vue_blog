package user_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4"`
	NickName string     `json:"nick_name"`
	UserId   uint       `json:"user_id"`
}

func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var user models.UserModel
	err := global.DB.Take(&user, cr.UserId).Error
	if err != nil {
		res.FailWithMessage("用户id错误,用户不存在", c)
		return
	}

	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		res.FailWithMessage("用户修改失败", c)
		return
	}

	res.OKWithMessage("用户修改成功", c)
}
