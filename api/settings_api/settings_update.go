package settings_api

import (
	"blog/gin/config"
	"blog/gin/core"
	"blog/gin/global"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var ser SettingsUri
	err := c.ShouldBindUri(&ser)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	fmt.Println("当前的值1", ser.Name)
	switch ser.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = info

	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = info
	case "jwt":
		var info config.JWT
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.JWT = info

	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	default:
		res.FailWithMessage("未知的配置", c)
		return
	}

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OKWith(c)
}
