package image

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/plugins/qiniu"
	"blog/gin/utils"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

type FileUploadResponse struct {
	FileName  string          `json:"file_name"`
	IsSuccess bool            `json:"is_success"`
	Msg       string          `json:"msg"`
	ImageType ctype.ImageType `json:"image_type"`
	FilePath  string          `json:"file_path"`
}

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

func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	fmt.Println("文件名", fileName)
	imageType := ctype.Local
	res.IsSuccess = false
	//白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return res
	}

	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	res.FilePath = filePath

	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，设定大小为: %dMB,当前大小为%.2fMB", global.Config.Upload.Size, size)
		return res
	}

	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	//读取md5的值
	byteFile, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteFile)
	//去数据库中查询是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		res.Msg = "图片已存在"
		res.FilePath = bannerModel.Path
		return res
	}

	if global.Config.QiNiu.Enable {
		//上传到七牛去
		imageType = ctype.QiNiu
		filePathIn7Cow, err := qiniu.UploadImageQiniu(byteFile, fileName, global.Config.QiNiu.Prefix)
		fmt.Println(err)
		if err != nil {
			global.Log.Error("七牛上传失败:", err)
			res.Msg = err.Error()
			res.ImageType = imageType
		} else {
			res.IsSuccess = true
			res.ImageType = imageType
			res.Msg = "上传成功"
			res.FilePath = filePathIn7Cow
		}
	}

	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: ctype.QiNiu,
	})

	return res
}
