package models

type MenuImageModel struct {
	MenuID     uint        `json:"menu_id"`
	MenuModel  MenuModel   `json:"foreignKey:MenuID"`
	ImageId    uint        `json:"image_id"`
	ImageModel BannerModel `gorm:"foreignKey:ImageID"`
	Sort       int         `gorm:"size:10" json:"sort"`
}
