package common

import (
	"blog/gin/global"
	"blog/gin/models"
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	models.PageView
	Debug bool
	Likes []string
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	//先查一个总数
	fmt.Println("key", option)
	fmt.Println("key1", option.Limit)
	fmt.Println("key2", option.PageView.Limit)
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	DB = DB.Where(model)
	if len(option.Likes) > 0 {
		for index, column := range option.Likes {
			if index == 0 {
				DB.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
				continue
			}
			DB.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
		}
	}
	//增加model
	count = DB.Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	//query := DB.Where(model)

	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
