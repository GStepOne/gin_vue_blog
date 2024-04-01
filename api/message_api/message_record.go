package message_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
)

type MessageRecordRequest struct {
	UserId uint `json:"user_id"`
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	userId := claims.UserId

	var _messageList []models.MessageModel

	var MessageList = make([]models.MessageModel, 0) //这里是防止接口返回null

	err = global.DB.Order("created_at desc").Find(&_messageList, "send_user_id = ? or rev_user_id= ?", userId, userId).Error
	if err != nil {
		res.FailWithMessage("聊天记录获取失败", c)
		return
	}

	for _, model := range _messageList {
		if model.RevUserID == cr.UserId || model.SendUserID == cr.UserId {
			MessageList = append(MessageList, model)
		}
	}

	//点开消息，改为消息已读

	res.OKWithData(MessageList, c)

}
