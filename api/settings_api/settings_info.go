package settings_api

import (
	"blog/gin/global"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		emailInfo := global.Config.Email
		emailInfo.Password = "******"
		res.OKWithData(emailInfo, c)
	case "qq":
		QQInfo := global.Config.QQ
		QQInfo.AppKey = "******"
		res.OKWithData(global.Config.QQ, c)
	case "qiniu":
		QiniuInfo := global.Config.QiNiu
		QiniuInfo.SecretKey = "******"
		res.OKWithData(global.Config.QiNiu, c)
	case "jwt":
		JwtInfo := global.Config.JWT
		JwtInfo.Secret = "******"
		res.OKWithData(global.Config.JWT, c)
	default:
		res.FailWithMessage("没有对应的方法", c)
	}

}
