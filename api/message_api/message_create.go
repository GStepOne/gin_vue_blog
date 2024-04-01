package message_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
)

type MessageRequest struct {
	SendUserId uint   `json:"send_user_id" binding:"required" msg:"发送方"`
	RevUserId  uint   `json:"rev_user_id" binding:"required" msg:"接受方"`
	Content    string `json:"content" binding:"required"`
}

// 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var sendUser, revUser models.UserModel
	err = global.DB.Take(&sendUser, cr.SendUserId).Error
	if err != nil {
		res.FailWithMessage("发送方不存在", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if claims.UserId != cr.SendUserId {
		res.FailWithMessage("当前发送人不是当前登陆者", c)
		return
	}

	err = global.DB.Take(&revUser, cr.RevUserId).Error
	if err != nil {
		res.FailWithMessage("接受方不存在", c)
		return
	}

	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserId,
		SendUserAvatar:   sendUser.Avatar,
		SendUserNickName: sendUser.NickName,

		RevUserID:       cr.RevUserId,
		RevUserAvatar:   revUser.Avatar,
		RevUserNickName: revUser.NickName,

		IsRead:  false,
		Content: cr.Content,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("消息发送失败", c)
		return
	}

	res.OKWithMessage("发送消息成功", c)
}
