package models

import (
	"time"
)

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id,select($any)" structs:"-"`
	CreatedAt time.Time `json:"created_at,select($any),omit(list)" structs:"-"`
	UpdatedAt time.Time `json:"updated_at,omit(list)"  structs:"-"`
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
	Page  int    `form:"page"  query:"page,default=1" json:"page"`
	Key   string `json:"key" form:"key" query:"key" `
	Limit int    `form:"limit,default=10" query:"limit,default=10" json:"limit"`
	Sort  string `form:"sort" query:"sort" json:"sort"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
