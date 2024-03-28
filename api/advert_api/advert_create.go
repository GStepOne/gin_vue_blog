package advert_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type AdvertRequest struct {
	Title  string `gorm:"size:32" binding:"required" maxLength:"20" msg:"请输入标题"  json:"title"`
	Href   string `json:"href" binding:"required,url" msg:"请输入跳转链接"`
	Images string `json:"images" binding:"required,url" msg:"请输入一个合法的图片地址"`
	IsShow bool   `json:"is_show" binding:"required"  msg:"请选择是否展示"`
}

func (AdvertApi) AdvertCreate(c *gin.Context) {
	var request AdvertRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res.FailWithError(err, &request, c)
		return
	}

	var advert models.AdvertModel

	err = global.DB.Take(&advert, "title = ?", request.Title).Error
	if err == nil {
		res.FailWithMessage("广告已经存在", c)
		return
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  request.Title,
		Href:   request.Href,
		Images: request.Images,
		IsShow: request.IsShow,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}

	res.OKWithMessage("添加广告成功", c)
}
