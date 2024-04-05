package redis_ser

import (
	"blog/gin/global"
	"blog/gin/utils"
	"time"
)

const prefix = "logout_"

// logout 针对注销操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(prefix+token, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
