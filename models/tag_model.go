package models

const ()

type TagModel struct {
	MODEL
	Title string `gorm:"size:16" json:"title'"`
	//Articles []ArticleModel `gorm:"many2many:article_tag_models" json:"-"` //关联该标签的文章列表
}
