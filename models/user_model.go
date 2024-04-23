package models

import (
	"blog/gin/models/ctype"
)

type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36" json:"nick_name,select(comment)"`
	UserName   string           `gorm:"size:36" json:"user_name"`
	Password   string           `gorm:"size:64" json:"-"`
	Avatar     string           `gorm:"size:256" json:"avatar,select(comment)"`
	Email      string           `gorm:"size:128" json:"email"`
	Tel        string           `gorm:"size:18" json:"tel"`
	Addr       string           `gorm:"size:64" json:"addr,select(comment)"`
	Token      string           `gorm:"size:64" json:"token"`
	Sign       string           `gorm:"size:255" json:"sign"`
	Link       string           `gorm:"size:255" json:"link"`
	IP         string           `gorm:"size:20" json:"ip,select(comment)"`
	Role       ctype.Role       `gorm:"size:4;default:1" json:"role"`
	RoleId     int              `gorm:"-" json:"role_id"`
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`
}
