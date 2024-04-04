package digg_api

import (
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/redis_ser"
	"github.com/gin-gonic/gin"
)

func (DiggApi) DiggArticleView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	redis_ser.Digg(cr.ID)

	res.OKWithMessage("文章点赞成功", c)
}
