package service

import (
	"blog/gin/service/image"
	"blog/gin/service/user_ser"
)

type ServiceGroup struct {
	ImageService image.ImageService
	UserService  user_ser.UserService
}

var ServiceApp = new(ServiceGroup)
