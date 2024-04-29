package middleware

import (
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/service/redis_ser"
	"blog/gin/utils/jwt"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", context)
			context.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", context)
			context.Abort()
		}

		//判断是否在redis中
		boolean := redis_ser.CheckLogout(token)
		if boolean {
			res.FailWithMessage("token已失效", context)
			context.Abort()
			return
		}

		context.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", context)
			context.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", context)
			context.Abort()
		}

		if claims.Role != uint(ctype.PermissionAdmin) {
			res.FailWithMessage("权限错误", context)
			context.Abort()
			return
		}

		context.Set("claims", claims)
	}
}
