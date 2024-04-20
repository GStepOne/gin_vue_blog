package log_api

import (
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/plugins/logstash"
	"blog/gin/service/common"
	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	models.PageView
	Level  logstash.Level `json:"level" query:"level" form:"level"`
	Status *bool          `json:"status" query:"status" form:"status"`
}

func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if cr.Sort == "" {
		cr.Sort = "created_at desc"
	}
	list, count, _ := common.ComList(logstash.LogStashModel{
		Level: cr.Level,
	}, common.Option{
		PageView: cr.PageView,
		Debug:    true,
		Likes:    []string{"content", "addr"},
	})

	res.OkWithList(list, count, c)
	return
}
