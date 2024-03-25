package models

// 评论表
type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  //子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` //父级评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                               //父评论
	Content            string          `gorm:"size:256" json:"content"`                         //评论内容
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`              //点赞
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`           //评论数
	Article            ArticleModel    `gorm:"foreignKey:ArticleID"json:"article"`              //评论文章
	ArticleID          uint            `json:"article_id"`                                      //评论文章id
	User               UserModel       `json:"user"`                                            //评论人
	UserID             uint            `gorm:"size:10" json:"user_id"`                          //评论内容
}
