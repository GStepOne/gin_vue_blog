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

//	type UserList struct {
//		models.UserModel
//
// }
func (UserApi) UserListView(c *gin.Context) {
	var pageView models.PageView
	if err := c.ShouldBindQuery(&pageView); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//token := c.Request.Header.Get("token")
	//if token == "" {
	//	res.FailWithMessage("未携带token", c)
	//	return
	//}
	//claims, err := jwt.ParseToken(token)
	//if err != nil {
	//	res.FailWithMessage("token 错误", c)
	//	return
	//}
	//fmt.Println("claims", claims)
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims) //断言

	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageView: pageView,
		Debug:    true,
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
