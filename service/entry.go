package service

import "blog/gin/service/image"

type ServiceGroup struct {
	ImageService image.ImageService
}

var ServiceApp = new(ServiceGroup)
