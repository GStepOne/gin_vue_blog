package message_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type MessageUserRequest struct {
	models.PageView
	Nickname string `json:"nickname" query:"nickname" form:"nickname"`
}

type MessageUserResponse struct {
	Avatar   string `json:"avatar"`
	Count    uint   `json:"count"`
	Nickname string `json:"nickname"`
	UserId   string `json:"user_id"`
}

func (MessageApi) MessageUserView(c *gin.Context) {
	var cr MessageUserRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var count int64
	var response []MessageUserResponse

	// 查询总数
	countResult := global.DB.Debug().Model(models.MessageModel{}).Group("send_user_id")
	if cr.Nickname != "" {
		countResult = countResult.Where("send_user_nick_name LIKE ?", "%"+cr.Nickname+"%")
	}
	countResult.Count(&count)

	// 查询数据
	query := global.DB.Debug().Model(models.MessageModel{}).Group("send_user_id").Select("count(id) as count,send_user_id as user_id,send_user_nick_name as nickname,send_user_avatar as avatar").Limit(cr.Limit).Offset((cr.Page - 1) * cr.Limit)
	if cr.Nickname != "" {
		query = query.Where("send_user_nick_name LIKE ?", "%"+cr.Nickname+"%")
	}

	result := query.Find(&response)

	if result.Error != nil {
		res.FailWithMessage("查询聊天记录失败", c)
		return
	}

	res.OkWithList(response, count, c)
}
