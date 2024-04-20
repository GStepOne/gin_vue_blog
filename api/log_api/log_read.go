package log_api

import (
	"blog/gin/global"
	"blog/gin/models/res"
	"blog/gin/plugins/logstash"
	"github.com/gin-gonic/gin"
)

type LogReadRequest struct {
	ID uint `json:"id" query:"id" form:"id"`
}

func (LogApi) LogReadView(c *gin.Context) {
	var cr LogReadRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var loginDataModel = logstash.LogStashModel{}
	err = global.DB.Debug().First(&loginDataModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("数据不存在", c)
		return
	}

	loginDataModel.ReadStatus = 1

	// 使用 Updates 方法来更新指定字段
	err = global.DB.Debug().Model(&loginDataModel).Updates(map[string]interface{}{"read_status": 1}).Error
	if err != nil {
		res.FailWithMessage("数据更新失败", c)
		return
	}

	res.OKWithMessage("已经更新为已读", c)
	return
}
