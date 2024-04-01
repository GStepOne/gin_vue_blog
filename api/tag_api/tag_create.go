package tag_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	Title string `gorm:"size:32" binding:"required" maxLength:"20" msg:"请输入标签"  structs:"title" json:"title"`
}

func (TagApi) TagCreate(c *gin.Context) {
	var request TagRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res.FailWithError(err, &request, c)
		return
	}

	var tag models.TagModel

	err = global.DB.Take(&tag, "title = ?", request.Title).Error
	if err == nil {
		res.FailWithMessage("标签已经存在", c)
		return
	}
	err = global.DB.Create(&models.TagModel{
		Title: request.Title,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加标签失败", c)
		return
	}

	res.OKWithMessage("添加标签成功", c)
}
