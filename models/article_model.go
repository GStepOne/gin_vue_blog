package models

import "blog/gin/models/ctype"

type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"`
	Abstract      string         `json:"abstract"`
	Content       string         `json:"content"`
	LookCount     int            `json:"look_count"`
	CommentCount  string         `json:"comment_count"`
	DiggCount     string         `json:"digg_count"`
	CollectsCount string         `json:"collects_count"`
	TagModels     []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"`
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`
	UserID        uint           `json:"user_id"`
	Category      string         `gorm:"size:20" json:"category"`
	Source        string         `json:"source"`
	Link          string         `json:"link"`
	Banner        BannerModel    `gorm:"foreignKey:BannerID" json:"-"`
	BannerID      uint           `json:"banner_id"`
	NickName      uint           `gorm:"size:42" json:"nick_name"`
	BannerPath    uint           `json:"banner_path"`
	Tags          ctype.Array    `json:"type:string;size:64" json:"tags"`
}
