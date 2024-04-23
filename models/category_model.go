package models

type CategoryModel struct {
	MODEL
	Label string `gorm:"size:16" json:"label"`
	Value string `gorm:"size:30" json:"value"`
}
