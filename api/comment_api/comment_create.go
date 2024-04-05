package comment_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/es_ser"
	"blog/gin/service/redis_ser"
	"blog/gin/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入内容"`
	ParentCommentID *uint  `json:"parent_comment_id"`
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	//1 文章是否存在 2
	//article := models.ArticleModel{}
	_, err = es_ser.CommonDetail(cr.ArticleID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}
	if cr.ParentCommentID != nil {
		//子评论,给父评论加1
		var parentComment models.CommentModel
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error

		if err != nil {
			res.FailWithMessage("父评论不存在", c)
			return
		}
		//判断父评论的文章是否与当前文章一致
		if parentComment.ArticleID != cr.ArticleID {
			res.FailWithMessage("评论文章不一致", c)
			return
		}
		//给父评论加一
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + ?", 1))
	}

	//添加评论
	global.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserId,
	})
	//文章存在的话
	redis_ser.NewCommentCount().Set(cr.ArticleID)
	res.OKWithMessage("文章评论成功", c)
	fmt.Println(claims.UserId)

}
