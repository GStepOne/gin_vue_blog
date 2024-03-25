package models

type SysModel struct {
	MODEL
	Version   string `gorm:"size:32" json:"version"`
	SiteTitle string `gorm:"size:32" json:"site_title"`
	SiteBeiAn string `gorm:"size:32" json:"site_bei_an"`
}
