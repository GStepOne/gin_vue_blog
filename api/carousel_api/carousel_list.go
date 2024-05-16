package carousel_api

import (
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/common"
	"github.com/gin-gonic/gin"
)

func (CarouselApi) CarouselListView(c *gin.Context) {
	var carouselList []models.CarouselModel
	var page models.PageView
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	carouselList, count, _ := common.ComList(models.CarouselModel{}, common.Option{
		PageView: page,
		Debug:    true,
		Likes:    []string{"name"},
	})

	res.OkWithList(carouselList, count, c)
}
