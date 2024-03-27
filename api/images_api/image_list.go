package images_api

import (
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageListView(c *gin.Context) {
	var imageList []models.BannerModel
	var page models.PageView
	err := c.ShouldBindQuery(&page)
	if err != nil {
		fmt.Println(err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	imageList, count, _ := common.ComList(models.BannerModel{}, common.Option{
		PageView: page,
		Debug:    true,
	})

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	res.OKWithData(gin.H{"count": count, "list": imageList}, c)
}
