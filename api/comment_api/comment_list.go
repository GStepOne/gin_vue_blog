package comment_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/service/redis_ser"
	"fmt"
	"github.com/gin-gonic/gin"
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
	info := redis_ser.NewCommentDigg().GetInfo()

	// 递归查询子评论
	for _, model := range RootCommentList {
		fmt.Println("root", fmt.Sprintf("%d", model.ID))
		digg := info[fmt.Sprintf("%d", model.ID)]
		model.DiggCount += digg

		subComments := findSubComment(model)
		// 循环遍历子评论
		for _, subComment := range subComments {
			subDigg := info[fmt.Sprintf("%d", model.ID)]
			subComment.DiggCount += subDigg
		}
		model.SubComments = subComments
	}

	// 返回结果
	res.OKWithData(RootCommentList, c)
}

// findSubComment 递归查询子评论
func findSubComment(model *models.CommentModel) []*models.CommentModel {
	global.DB.Preload("SubComments.User").Order("id desc").Take(&model)
	for _, sub := range model.SubComments {
		sub.SubComments = findSubComment(sub)
	}

	return model.SubComments
}

func TotalSubComment(model *models.CommentModel) (int, []uint) {

	dataArr := findSubComment(model)
	subCommentsCount := 0
	var ids []uint
	for _, comment := range dataArr {
		subCommentsCount += len(comment.SubComments)
		ids = append(ids, comment.ID)
	}

	return subCommentsCount, ids
}
