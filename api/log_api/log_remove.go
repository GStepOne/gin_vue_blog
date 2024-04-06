package log_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/plugins/logstash"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (LogApi) LogRemoveListView(c *gin.Context) {
	var query models.RemoveRequest
	//批量删除
	err := c.ShouldBindJSON(&query)

	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		fmt.Println(err)
		return
	}
	var logStashList []logstash.LogStashModel
	count := global.DB.Find(&logStashList, query.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("日志不存在", c)
		return
	}
	global.DB.Delete(&logStashList)
	res.OKWithMessage(fmt.Sprintf("共删除 %d 个日志", count), c)
}
