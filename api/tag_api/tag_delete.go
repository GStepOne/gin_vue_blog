package tag_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	//批量删除
	var query models.RemoveRequest
	err := c.ShouldBindJSON(&query)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var TagList []models.TagModel
	count := global.DB.Find(&TagList, query.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("标签不存在", c)
		return
	}
	global.DB.Delete(&TagList)
	res.OKWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)
	//绑一个钩子函数
}
