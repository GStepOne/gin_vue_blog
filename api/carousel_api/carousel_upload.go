package carousel_api

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
	"strings"
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

type FileUploadResponse struct {
	FileName  string          `json:"file_name"`
	IsSuccess bool            `json:"is_success"`
	Msg       string          `json:"msg"`
	ImageType ctype.ImageType `json:"image_type"`
	FilePath  string          `json:"file_path"`
}

func (CarouselApi) CarouselsSingleUploadView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println("我不信", err)
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

	fmt.Println("当前的file", file.Filename)
	var resList image.FileUploadResponse
	resList = service.ServiceApp.ImageService.ImageUploadService(file, true)
	fmt.Println("单张上传的轮播图的返回", resList)
	//不是7牛 本地保存一下
	if !global.Config.QiNiu.Enable {
		fmt.Println("文件路径非7牛", resList.FilePath)

		//有前缀的 干掉前缀
		if strings.HasPrefix(resList.FilePath, "/") {
			err = c.SaveUploadedFile(file, resList.FilePath[1:])
			resList.FilePath = global.Config.SiteInfo.Web + resList.FilePath
		} else {
			err = c.SaveUploadedFile(file, resList.FilePath)
			resList.FilePath = global.Config.SiteInfo.Web + "/" + resList.FilePath
		}

		if err != nil {
			global.Log.Error(err.Error())
			resList.Msg = err.Error()
			resList.IsSuccess = false
			res.FailWithMessage("轮播图保存失败", c)
			return
		}
	}

	res.OK(resList.FilePath, resList.Msg, c)

}
