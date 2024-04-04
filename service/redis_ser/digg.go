package redis_ser

import (
	"blog/gin/global"
	"strconv"
)

const diffPrefix = "digg_"

func Digg(id string) error {
	num, _ := global.Redis.HGet(diffPrefix, id).Int()
	num++
	err := global.Redis.HSet(diffPrefix, id, 1).Err()

	return err
}

// 获取点赞数
func GetDigg(id string) int {
	num, _ := global.Redis.HGet(diffPrefix, id).Int()
	return num
}

//同步点赞数据到es

// 取出点赞数据
func GetDiggInfo() map[string]int {
	maps := global.Redis.HGetAll(diffPrefix).Val()
	var DiggInfo = map[string]int{}
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func DiggClear() {
	//直接删除key ？？
	global.Redis.Del(diffPrefix)
}
