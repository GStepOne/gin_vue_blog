package user_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Link     string `json:"link"`
	Sign     string `json:"sign"`
}

func (UserApi) UserUpdateInfoView(c *gin.Context) {
	var cr UserInfo
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	userId := claims.UserId
	var user models.UserModel
	err := global.DB.Take(&user, userId).Error
	if err != nil {
		res.FailWithMessage("用户id错误,用户不存在", c)
		return
	}

	err = global.DB.Model(&user).Updates(map[string]any{
		"avatar":    cr.Avatar,
		"nick_name": cr.NickName,
		"link":      cr.Link,
		"sign":      cr.Sign,
	}).Error

	if err != nil {
		res.FailWithMessage("用户修改失败", c)
		return
	}

	res.OKWithMessage("用户修改成功", c)
}
