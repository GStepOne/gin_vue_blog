package qiniu

import (
	"blog/gin/global"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type UploadQiniu struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func GetQinNiuToken() string {
	accessKey := global.Config.QiNiu.AccessKey
	secretKey := global.Config.QiNiu.SecretKey
	bucket := global.Config.QiNiu.Bucket
	mac := auth.New(accessKey, secretKey)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	return putPolicy.UploadToken(mac)
}

func UploadImageQiniu(localFile string, filename string) (ok string, err error) {
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong //华东 其他见sdk文档
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	token := GetQinNiuToken()

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "jack gin blogs",
		},
	}
	err = formUploader.PutFile(context.Background(), &ret, token, filename, localFile, &putExtra)
	if err != nil {
		return "", err
	}
	fmt.Println(ret.Key, ret.Hash)

	return "ok", nil

}
