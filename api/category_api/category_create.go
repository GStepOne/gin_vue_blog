package category_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type CategoryRequest struct {
	Label string `gorm:"size:32" binding:"required" maxLength:"20" msg:"请输入分类"  structs:"label" json:"label"`
}

func (CategoryApi) CategoryCreate(c *gin.Context) {
	var request CategoryRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		res.FailWithError(err, &request, c)
		return
	}

	var tag models.CategoryModel

	err = global.DB.Take(&tag, "label = ?", request.Label).Error
	if err == nil {
		res.FailWithMessage("分类已经存在", c)
		return
	}
	err = global.DB.Create(&models.CategoryModel{
		Label: request.Label,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加分类失败", c)
		return
	}

	res.OKWithMessage("添加分类成功", c)
}
