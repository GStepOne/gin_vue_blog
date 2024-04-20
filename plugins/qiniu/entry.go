package qiniu

import (
	"blog/gin/global"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
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

func GetCfg() storage.Config {
	cfg := storage.Config{}
	//空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	return cfg
}

func UploadImageQiniu(data []byte, filename string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Enable {
		return "", errors.New("请启用七牛云上传")
	}
	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置accessKey及secretKey")
	}
	uploadSize := float64(len(data)) / 1024 / 1024
	if uploadSize > q.Size {
		return "", errors.New("请配置accessKey及secretKey")
	}
	token := GetQinNiuToken()
	fmt.Println(token)
	cfg := GetCfg()
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置

	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}

	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s%s.png", prefix, now, filename)

	dataLen := int64(len(data))
	err = formUploader.Put(context.Background(), &ret, token, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		global.Log.Error("七牛上传错误", err.Error())
		return "", err
	}

	return fmt.Sprintf("%s/%s", q.CDN, ret.Key), nil

}

//http://sazz982o7.hd-bkt.clouddn.com/gin/20240328184835upworkavatar.jpg.png?
//e=1711623243&token=P-8AgS1WqT-BgVSDf1o3oPzs-892QbbNuHectlVo:1QUZjiVcV54OaZuQLZq1DfyVQBU=
