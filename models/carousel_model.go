package models

import (
	"blog/gin/global"
	"blog/gin/models/ctype"
	"gorm.io/gorm"
	"os"
)

type CarouselModel struct {
	MODEL
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `gorm:"size:38" json:"name"`
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"`
}

// BeforeDelete 在同一个事务中更新数据
func (b *CarouselModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 检查文件是否存在
		if _, err := os.Stat(b.Path[1:]); err == nil {
			// 文件存在，执行删除
			err = os.Remove(b.Path[1:])
			if err != nil {
				global.Log.Error(err)
			}
		} else if os.IsNotExist(err) {
			// 文件不存在，可能已经被删除
			global.Log.Infof("文件 %s 不存在，无需删除", b.Path)
		} else {
			// 其他错误，记录日志
			global.Log.Errorf("无法访问文件 %s: %v", b.Path, err)
		}
	}
	return
}
