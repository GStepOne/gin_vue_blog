package message_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

type Message struct {
	SendUserId       uint      `json:"send_user_id"`
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserId        uint      `json:"rev_user_id"`
	RevUserNickname  string    `json:"rev_user_nickname"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`
	CreatedAt        time.Time `json:"created_at"`    //最新的消息时间
	MessageCount     int       `json:"message_count"` //消息条数
}
type MessageGroup map[uint]*Message

var messageGroup = MessageGroup{}

const (
	userID       = 1
	userNickname = "sansan"
)

func (MessageApi) MessageListView(c *gin.Context) {
	//var cr models.PageView
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var messageList []models.MessageModel
	var messageGroup = MessageGroup{}
	var messages []Message

	err := global.DB.Order("created_at asc").Find(&messageList, "send_user_id = ? or rev_user_id =?", claims.UserId, claims.UserId).Error
	if err != nil {
		res.FailWithMessage("获取聊天记录失败", c)
		return
	}

	for _, model := range messageList {
		message := Message{
			SendUserId:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserId:        model.RevUserID,
			RevUserNickname:  model.RevUserNickName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}
		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}

	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	res.OKWithData(messageGroup, c)
}
