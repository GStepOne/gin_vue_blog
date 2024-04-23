package category_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type CategoryResponseData struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (CategoryApi) CategoryListViewNoPage(c *gin.Context) {
	var response []CategoryResponseData
	err := global.DB.Model(&models.CategoryModel{}).Debug().Select("id,label,value").Find(&response).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("获取分类失败", c)
		return
	}

	res.OKWithData(response, c)
}
