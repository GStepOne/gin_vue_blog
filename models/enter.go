package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

type PageInfo struct {
	Page  int      `json:"page" form:"page"`
	Limit int      `json:"limit" form:"limit"`
	Key   string   `json:"key" form:"key"`
	Sort  string   `json:"sort" form:"key"`
	Like  []string `json:"like" form:"like"`
}

type ResponseList struct {
	List  any `json:"list"`
	Count any `json:"count"`
}
