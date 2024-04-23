package tag_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type TagListView struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type TagResponseData struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (TagApi) TagListViewNoPage(c *gin.Context) {
	var response []TagResponseData
	err := global.DB.Model(&models.TagModel{}).Select("id,title as label,value").Find(&response).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("获取标签失败", c)
	}

	res.OKWithData(response, c)
}
