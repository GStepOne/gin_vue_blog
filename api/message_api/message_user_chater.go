package message_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type MessageUserChaterRequest struct {
	UserId int `json:"user_id" query:"user_id" form:"user_id" uri:"user_id"`
}

func (MessageApi) MessageUserChaterView(c *gin.Context) {
	var cr MessageUserChaterRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	response, err := getMessageUsers(cr.UserId)
	if err != nil {
		res.FailWithMessage("查询聊天记录失败", c)
		return
	}

	res.OkWithList(response, int64(len(response)), c)
}

func getMessageUsers(userId int) ([]MessageUserResponse, error) {
	var response []MessageUserResponse

	err := global.DB.Debug().Model(models.MessageModel{}).Group("user_id").
		Select("count(id) as count,rev_user_id as user_id,rev_user_nick_name as nickname,rev_user_avatar as avatar").
		Where("send_user_id = ?", userId).Find(&response).Error

	return response, err
}
