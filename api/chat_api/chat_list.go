package chat_api

import (
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/common"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ChatApi) ChatListView(c *gin.Context) {
	var cr models.PageView
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	cr.Sort = "created_at asc"
	list, count, _ := common.ComList(models.ChatModel{IsGroup: true}, common.Option{
		PageView: cr,
	})

	res.OkWithList(filter.Omit("list", list), count, c)
}
