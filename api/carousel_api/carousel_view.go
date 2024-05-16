package carousel_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type CarouselResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func (CarouselApi) CarouselView(c *gin.Context) {
	var carouselList []CarouselResponse

	err := global.DB.Model(models.CarouselModel{}).Select("id", "path", "name").Find(&carouselList).Error
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	res.OKWithData(carouselList, c)
}
