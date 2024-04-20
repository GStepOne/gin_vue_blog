package images_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func (ImagesApi) ImageView(c *gin.Context) {
	var imageList []ImageResponse

	err := global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Find(&imageList).Error
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	res.OKWithData(imageList, c)
}
