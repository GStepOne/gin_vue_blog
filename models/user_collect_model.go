package models

import "time"

// 收藏文章表
type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"size:32"`
	//ArticleModel uint      `gorm:"foreignKey:ArticleID"`
	CreatedAt time.Time `gorm:"created_at"`
}
