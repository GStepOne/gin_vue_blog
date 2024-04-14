package user_api

import (
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/service/common"
	"blog/gin/utils/desens"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	models.PageView
	Role ctype.Role `json:"role" query:"role" form:"role"`
}

func (UserApi) UserListView(c *gin.Context) {
	var pageView UserListRequest
	if err := c.ShouldBindQuery(&pageView); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims) //断言

	model := models.UserModel{}
	role := int(pageView.Role)
	if role != 0 {
		model = models.UserModel{
			Role: pageView.Role,
		}
	}

	var likes []string
	if pageView.Key != "" {
		likes = []string{"user_name", "nick_name"}
	}

	list, count, _ := common.ComList(model, common.Option{
		PageView: pageView.PageView,
		Debug:    true,
		Likes:    likes,
	})

	var users []models.UserModel
	for _, user := range list {
		//只有超级管理员才能看到用户昵称
		if ctype.Role(claims.Role) == ctype.PermissionVisitor {
			//管理员
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationMail(user.Email)
		users = append(users, user)
	}

	res.OkWithList(users, count, c)
}
