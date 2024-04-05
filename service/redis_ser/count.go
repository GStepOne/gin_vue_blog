package redis_ser

import (
	"blog/gin/global"
	"strconv"
)

type CountDB struct {
	Index string //前缀
}

// 设置某个数据
func (c CountDB) Set(id string) error {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	num++
	err := global.Redis.HSet(c.Index, id, num).Err()

	return err
}

// 获取某个数据
func (c CountDB) Get(id string) int {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	return num
}

// 清空
func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}

// 取出数据
func (c CountDB) GetInfo() map[string]int {
	maps := global.Redis.HGetAll(c.Index).Val()
	var info = map[string]int{}
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		info[id] = num
	}
	return info
}

// 设置某个数据
func (c CountDB) SetCount(id string, num int) error {
	oldNum, _ := global.Redis.HGet(c.Index, id).Int()
	newNum := oldNum + num
	err := global.Redis.HSet(c.Index, id, newNum).Err()

	return err
}
