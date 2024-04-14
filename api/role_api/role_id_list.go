package role_api

import (
	"blog/gin/models/res"
	"github.com/gin-gonic/gin"
)

type OptionResponse struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

func (RoleApi) RoleIDListView(c *gin.Context) {

	res.OKWithData([]OptionResponse{
		{"管理员", 1},
		{"普通用户", 2},
		{"游客", 3}},
		c)
}
