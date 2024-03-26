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
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	fmt.Println("before", global.Config)
	global.Config.SiteInfo = cr
	fmt.Println(cr)

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}

	fmt.Println("after", global.Config)

	res.OKWith(c)
}
