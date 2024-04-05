package models

import (
	"time"
)

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id,select($any)" structs:"-"`
	CreatedAt time.Time `json:"created_at,select($any)" structs:"-"`
	UpdatedAt time.Time `json:"-"  structs:"-"`
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

type PageView struct {
	Page  int    `form:"page" default:"1"`
	Key   string `form:"key"`
	Limit int    `form:"limit" default:"15" query:"limit,default:15"`
	Sort  string `form:"sort"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
