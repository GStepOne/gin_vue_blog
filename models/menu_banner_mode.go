package models

type MenuBannerModel struct {
	MenuId    uint      `json:"menu_id"`
	MenuModel MenuModel `gorm:"foreignKey:MenuId" json:"-"`
	//gorm:"foreignKey:关联表的结构体字段;references:当前表的结构体字段;`
	ImageId    uint        `json:"image_id"`
	ImageModel BannerModel `gorm:"foreignKey:ImageId"` //foreignKey
	Sort       int         `gorm:"size:10" json:"sort"`
}
