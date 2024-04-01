package user_api

import (
	"blog/gin/global"
	"blog/gin/models/res"
	"blog/gin/service"
	"blog/gin/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	fmt.Println(claims)
	token := c.Request.Header.Get("token")
	//需要计算过期时间
	err := service.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}

	res.OKWithMessage("注销成功", c)
}
