package user_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/plugins/qq"
	"blog/gin/utils"
	"blog/gin/utils/jwt"
	"blog/gin/utils/random"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("qq登录code 为空", c)
		return
	}
	fmt.Println(code)

	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	openId := qqInfo.OpenId
	var user models.UserModel
	err = global.DB.Take(&user, "token=?", openId).Error
	if err != nil {
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   "QQ_" + openId, //qq登录
			Password:   utils.HashPwd(random.RandStr(16)),
			Avatar:     qqInfo.Avatar,
			Addr:       "内网ip",
			Token:      qqInfo.OpenId,
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&user).Error

		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("注册失败", c)
			return
		}
	}

	token, err := jwt.GenToken(jwt.JwtPayload{
		Nickname: user.NickName,
		Role:     uint(user.Role),
		UserId:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", c)
		return
	}

	res.OKWithData(token, c)

}
