package flag

import (
	"blog/gin/models"
	"fmt"
)

// 生成es的doc
func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()
	fmt.Println("创建文章es索引成功")
}
