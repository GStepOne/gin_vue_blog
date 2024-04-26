package models

import (
	"blog/gin/models/ctype"
	"time"
)

type ChatModel struct {
	MODEL       `json:"model"`
	CreatedAt   time.Time         `json:"created_at" structs:"-"`
	NickName    string            `gorm:"size:15" json:"nick_name"`
	Avatar      string            `gorm:"size:128" json:"avatar"`
	Content     string            `gorm:"size:256" json:"content"`
	MessageType ctype.MessageType `gorm:"size:4" json:"message_type"`
	IP          string            `json:"ip"`
	Addr        string            `gorm:"size:64" json:"addr"`
	IsGroup     bool              `json:"is_group"` //是否是群组消息
}
