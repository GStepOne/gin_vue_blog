package article_api

import (
	"blog/gin/models/res"
	"blog/gin/service/es_ser"
	"github.com/gin-gonic/gin"
)

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	model, err := es_ser.CommonDetail(cr.ID)
	if err != nil {
		res.FailWithMessage("文章详情获取失败", c)
		return
	}

	res.OKWithData(model, c)
}
