package images_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/plugins/qiniu"
	"blog/gin/service"
	"blog/gin/service/image"
	"blog/gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"os"
	"path"
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

func (ImagesApi) ImagesMultiUploadViewEver(c *gin.Context) {
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

	var resList []FileUploadResponse
	for _, file := range fileList {
		fileName := file.Filename
		nameList := strings.Split(fileName, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1])
		fmt.Println(suffix)
		fmt.Println(utils.InList(suffix, WhiteImageList))
		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}

		filePath := path.Join(basePath, file.Filename)
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定大小，设定大小为: %dMB,当前大小为%.2fMB", global.Config.Upload.Size, size),
			})
			continue
		}

		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "success",
		})
		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		//读取md5的值
		byteFile, err := io.ReadAll(fileObj)
		imageHash := utils.Md5(byteFile)
		fmt.Println(imageHash)
		//去数据库中查询是否存在
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
		if err == nil {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "图片找到了",
			})
			continue
		}
		imageType := ctype.Local
		//上传到7牛
		if global.Config.QiNiu.Enable {
			//上传到七牛去
			imageType = ctype.QiNiu
			filePathIn7Cow, err := qiniu.UploadImageQiniu(byteFile, fileName, global.Config.QiNiu.Prefix)
			fmt.Println("7niu", filePathIn7Cow)
			if err != nil {
				resList = append(resList, FileUploadResponse{
					FileName:  fileName,
					FilePath:  filePathIn7Cow,
					Msg:       "上传失败",
					ImageType: imageType,
					IsSuccess: false,
				})
			} else {
				resList = append(resList, FileUploadResponse{
					FileName:  fileName,
					FilePath:  filePathIn7Cow,
					Msg:       "上传成功",
					ImageType: imageType,
					IsSuccess: true,
				})
			}
		}

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}
		//图片入库
		global.DB.Create(&models.BannerModel{
			Path:      filePath,
			Hash:      imageHash,
			Name:      fileName,
			ImageType: imageType,
		})

	}

	res.OKWithData(resList, c)

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
