package user_api

import (
	"blog/gin/global"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/service/user_ser"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nickname" binding:"required" msg:"请输入昵称"`
	UserName string     `json:"username" binding:"required" msg:"请输入用户名"`
	Password string     `json:"password" binding:"required" msg:"请输入密码"`
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &err, c)
		return
	}

	err = user_ser.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("用户创建失败", c)
		return
	}

	res.OKWithMessage("用户创建成功", c)
}
