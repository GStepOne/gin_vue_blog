package images_api

import (
	"blog/gin/global"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/service"
	"blog/gin/service/image"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
)

var (
	WhiteImageList = []string{
		"jpg",
		"jpeg",
		"png",
		"gif",
		"bmp",
		"tiff",
		"svg",
		"webp",
	}
)

/**
** ImageUploadView 上传单个图片
 */
//func (ImagesApi) ImageUploadView(c *gin.Context) {
//	fileHeader, err := c.FormFile("image")
//	if err != nil {
//		res.FailWithMessage(err.Error(), c)
//		return
//	}
//
//	fmt.Println(fileHeader.Header)
//	fmt.Println(fileHeader.Size)
//	fmt.Println(fileHeader.Filename)
//}

type FileUploadResponse struct {
	FileName  string          `json:"file_name"`
	IsSuccess bool            `json:"is_success"`
	Msg       string          `json:"msg"`
	ImageType ctype.ImageType `json:"image_type"`
	FilePath  string          `json:"file_path"`
}

func (ImagesApi) ImagesMultiUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
			return
		}
	}

	var resList []image.FileUploadResponse
	for _, file := range fileList {
		ServiceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		fmt.Println("上传返回值", ServiceRes)
		if ServiceRes.IsSuccess == false {
			global.Log.Error(ServiceRes.Msg)
			resList = append(resList, ServiceRes)
			continue
		}
		resList = append(resList, ServiceRes)
		//不是7牛 本地保存一下
		if !global.Config.QiNiu.Enable {
			fmt.Println("文件路径 非7牛", ServiceRes.FilePath)
			err = c.SaveUploadedFile(file, ServiceRes.FilePath)
			if err != nil {
				global.Log.Error(err.Error())
				ServiceRes.Msg = err.Error()
				ServiceRes.IsSuccess = false
				resList = append(resList, ServiceRes)
				continue
			}
		}
	}

	res.OKWithData(resList, c)

}

func (ImagesApi) ImagesSingleUploadView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	//路径不存在创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
			return
		}
	}

	var resList image.FileUploadResponse
	resList = service.ServiceApp.ImageService.ImageUploadService(file)
	//不是7牛 本地保存一下
	if !global.Config.QiNiu.Enable {
		fmt.Println("文件路径 非7牛", resList.FilePath)
		err = c.SaveUploadedFile(file, resList.FilePath)
		if err != nil {
			global.Log.Error(err.Error())
			resList.Msg = err.Error()
			resList.IsSuccess = false
		}
		resList.FilePath = "/" + resList.FilePath
	}

	res.OK(resList.FilePath, resList.Msg, c)

}
