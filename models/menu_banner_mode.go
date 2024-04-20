package models

type MenuBannerModel struct {
	MenuId      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuId" json:"-"`
	BannerId    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerId" json:"-"` //foreignKey
	Sort        int         `gorm:"size:10" json:"sort"`
}
