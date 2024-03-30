package main

import (
	"blog/gin/models"
	"fmt"
	"github.com/fatih/structs"
)

type AdvertView struct {
	models.MODEL `structs:"-"`
	Title        string `gorm:"size:32" binding:"required" max:"20" msg:"请输入标题"  json:"title" structs:"title"`
	Href         string `json:"href" binding:"required,url" msg:"请输入跳转链接" structs:"-"`
	Images       string `json:"images" binding:"required,url" msg:"请输入一个合法的图片地址"`
	IsShow       bool   `json:"is_show" binding:"required"  msg:"请选择是否展示" structs:"is_show"`
}

func main() {
	u1 := AdvertView{
		Title:  "xxx",
		Href:   "22",
		Images: "xxx",
		IsShow: true,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}
