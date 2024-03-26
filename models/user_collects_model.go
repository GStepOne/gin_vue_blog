package models

import "time"

// 收藏文章表
type UserCollectsModel struct {
	UserID       uint      `gorm:"primaryKey"`
	UserModel    UserModel `gorm:"foreignKey:UserID"`
	ArticleID    uint      `gorm:"primaryKey"`
	ArticleModel uint      `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
