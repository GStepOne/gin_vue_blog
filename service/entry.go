package service

import (
	"blog/gin/service/image"
	"blog/gin/service/redis_ser"
	"blog/gin/service/user_ser"
)

type ServiceGroup struct {
	ImageService image.ImageService
	UserService  user_ser.UserService
	RedisService redis_ser.RedisService
}

var ServiceApp = new(ServiceGroup)
