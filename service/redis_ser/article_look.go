package redis_ser

import (
	"blog/gin/global"
	"strconv"
)

const lookPrefix = "look_"

func ArticleLook(id string) error {
	num, _ := global.Redis.HGet(lookPrefix, id).Int()
	num++
	err := global.Redis.HSet(lookPrefix, id, 1).Err()

	return err
}

// 获取点赞数
func GetLook(id string) int {
	num, _ := global.Redis.HGet(lookPrefix, id).Int()
	return num
}

//同步点赞数据到es

// 取出点赞数据
func GetLookInfo() map[string]int {
	maps := global.Redis.HGetAll(lookPrefix).Val()
	var LookInfo = map[string]int{}
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		LookInfo[id] = num
	}
	return LookInfo
}

func LookClear() {
	//直接删除key ？？
	global.Redis.Del(lookPrefix)
}
