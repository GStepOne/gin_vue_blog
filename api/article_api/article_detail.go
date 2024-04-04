package article_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/es_ser"
	"blog/gin/service/redis_ser"
	"github.com/gin-gonic/gin"
)

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	redis_ser.ArticleLook(cr.ID)

	model, err := es_ser.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("文章详情获取失败", c)
		return
	}

	res.OKWithData(model, c)
}

func (ArticleApi) ArticleDetailByTitle(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	model, err := es_ser.CommonDetailByKeyWord(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OKWithData(model, c)
}
