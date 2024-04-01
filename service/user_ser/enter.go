package user_ser

import (
	"blog/gin/global"
	"blog/gin/utils/jwt"
	"fmt"
	"time"
)

type UserService struct {
}

func (UserService) Logout(claims *jwt.CustomClaims, token string) error {
	//需要计算过期时间
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	err := global.Redis.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()
	return err
}
