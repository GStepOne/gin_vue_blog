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
	"github.com/gin-gonic/gin"
)

func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("qq登录code 为空", c)
		return
	}
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		global.Log.Error("登录QQ报错", err.Error())
		res.FailWithMessage(err.Error(), c)
		return
	}

	openId := qqInfo.OpenId
	var user models.UserModel
	err = global.DB.Take(&user, "token=?", openId).Error
	ip, addr := utils.GetAddrByGin(c)
	if err != nil {
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   "QQ_" + openId, //qq登录
			Password:   utils.HashPwd(random.RandStr(16)),
			Avatar:     qqInfo.Avatar,
			Addr:       addr,
			Token:      qqInfo.OpenId,
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
			IP:         ip,
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

	global.DB.Create(&models.LoginDataModel{
		UserId:    user.ID,
		IP:        ip,
		Nickname:  user.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignQQ,
	})

	res.OKWithData(token, c)

}

func (UserApi) QQLoginPath(c *gin.Context) {
	path := global.Config.QQ.GetPath()
	if len(path) > 0 {
		res.OKWithData(path, c)
		return
	}

	res.OKWithMessage("qq登录的路径获取失败", c)
}
