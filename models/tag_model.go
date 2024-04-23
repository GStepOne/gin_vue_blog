package models

const ()

type TagModel struct {
	MODEL
	Title string `gorm:"size:16" json:"title"`
	Value string `gorm:"size:30" json:"value"`
}
