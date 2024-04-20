package message_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type MessageRecordRequest struct {
	SendUserId uint `json:"send_user_id" query:"send_user_id" form:"send_user_id"`
	RevUserId  uint `json:"rev_user_id" query:"rev_user_id" form:"rev_user_id"`
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithError(err, &cr, c)
		return
	}

	var _messageList []models.MessageModel

	fmt.Println("wtf", cr)
	//var MessageList = make([]models.MessageModel, 0) //这里是防止接口返回null

	err = global.DB.Debug().Order("created_at desc").Find(&_messageList,
		"(send_user_id = ? and rev_user_id= ?) or (rev_user_id = ? and send_user_id= ?)",
		cr.SendUserId, cr.RevUserId, cr.SendUserId, cr.RevUserId).Error
	if err != nil {
		res.FailWithMessage("聊天记录获取失败", c)
		return
	}

	res.OkWithList(_messageList, int64(len(_messageList)), c)

}
