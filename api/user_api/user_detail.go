package user_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	models.UserModel
}

func (UserApi) UserDetailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	userId := claims.UserId

	var user models.UserModel
	err := global.DB.Take(&user, "id = ?", userId).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	res.OKWithData(user, c)

}
