package logstash

import (
	"time"
)

type LogStashModel struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	IP         string    `gorm:"size:32" json:"ip"`
	Addr       string    `gorm:"size:64" json:"addr"`
	Level      Level     `gorm:"size:4" json:"level"`
	Content    string    `gorm:"size:128" json:"content"`
	UserID     uint      `json:"user_id"`
	Title      string    `json:"title"`
	ReadStatus uint      `json:"read_status"`
	UserName   string    `json:"user_name"`
}
