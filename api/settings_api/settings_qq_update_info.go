package settings_api

import (
	"blog/gin/global"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsQQInfoUpdateView(c *gin.Context) {
	res.OKWithData(global.Config.QQ, c)
}
