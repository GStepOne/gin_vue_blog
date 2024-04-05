package comment_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `json:"article_id"`
}

// CommentListView 处理获取文章评论列表的请求
func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 查询根评论
	var RootCommentList []*models.CommentModel
	global.DB.Preload("User").Where("article_id = ? and parent_comment_id is null", cr.ArticleID).Find(&RootCommentList)

	// 递归查询子评论
	for _, model := range RootCommentList {
		model.SubComments = findSubComment(*model)
	}

	res.OKWithData(filter.Select("comment", RootCommentList), c)
}

// findSubComment 递归查询子评论
func findSubComment(model models.CommentModel) []models.CommentModel {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		sub.SubComments = findSubComment(sub)
	}

	return model.SubComments
}
