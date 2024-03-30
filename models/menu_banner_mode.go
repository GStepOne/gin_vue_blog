package models

type MenuBannerModel struct {
	MenuId     uint        `json:"menu_id"`
	MenuModel  MenuModel   `gorm:"foreignKey:MenuId" json:"-"`
	ImageId    uint        `json:"image_id"`
	ImageModel BannerModel `gorm:"foreignKey:ImageId"` //foreignKey
	Sort       int         `gorm:"size:10" json:"sort"`
}
