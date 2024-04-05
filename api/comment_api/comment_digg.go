package comment_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/redis_ser"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CommentIDRequest struct {
	ID uint `json:"id" uri:"id"  form:"id" binding:"required"`
}

func (CommentApi) CommentDigg(c *gin.Context) {
	var CommentId CommentIDRequest
	err := c.ShouldBindUri(&CommentId)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, CommentId.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", c)
		return
	}

	//增加文章评论点赞数量
	err = redis_ser.NewCommentDigg().Set(fmt.Sprintf("%d", CommentId.ID))

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("评论点赞失败", c)
		return
	}

	res.OKWithMessage("评论点赞成功", c)
}
