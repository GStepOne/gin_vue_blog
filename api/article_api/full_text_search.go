package article_api

import (
	"blog/gin/models"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) FullTextSearch(c *gin.Context) {
	var cr models.PageView

	_ = c.ShouldBindQuery(&cr)

}
