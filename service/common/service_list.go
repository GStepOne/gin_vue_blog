package common

import (
	"blog/gin/global"
	"blog/gin/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageView
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	//先查一个总数
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	count = DB.Debug().Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	DB.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list)

	return list, count, err
}
