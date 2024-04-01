package user_ser

import (
	"blog/gin/service/redis_ser"
	"blog/gin/utils/jwt"
	"time"
)

func (UserService) Logout(claims *jwt.CustomClaims, token string) error {
	//需要计算过期时间
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	err := redis_ser.Logout(token, diff)
	//err := global.Redis.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()
	return err
}
