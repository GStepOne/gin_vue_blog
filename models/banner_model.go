package models

import (
	"blog/gin/global"
	"blog/gin/models/ctype"
	"gorm.io/gorm"
	"os"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `gorm:"size:38" json:"name"`
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"`
	//ArticleModels []ArticleModel `gorm:"foreignKey:CoverID" json:"-"`
}

// 在同一个事务中更新数据
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		err = os.Remove(b.Path)
		global.Log.Error(err)
		return err
	}
	return
}
