package models

// 评论表
type CommentModel struct {
	MODEL              `json:",select(comment)"`
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments,select(comment)"`         //子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"parent_comment_model,select(comment)"` //父级评论
	ParentCommentID    *uint           `json:"parent_comment_id,select(comment)"`                                      //父评论
	Content            string          `gorm:"size:256" json:"content,select(comment)"`                                //评论内容
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count,select(comment)"`                     //点赞
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count,select(comment)"`                  //评论数
	ArticleID          string          `gorm:"size:32" json:"article_id,select(comment)"`                              //评论文章id
	User               UserModel       `json:"user,select(comment)"`                                                   //评论人
	UserID             uint            `json:"user_id,select(comment)"`                                                //评论内容
}
