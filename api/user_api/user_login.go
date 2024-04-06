package user_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/plugins/logstash"
	"blog/gin/utils"
	"blog/gin/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func (UserApi) EmailLogin(c *gin.Context) {
	var EmailLoginRequest EmailLoginRequest
	err := c.ShouldBindJSON(&EmailLoginRequest)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &EmailLoginRequest, c)
		return
	}
	log := logstash.NewLogByGin(c)
	var userModel models.UserModel
	err = global.DB.Debug().Take(&userModel, "user_name = ? or email=?", EmailLoginRequest.Username, EmailLoginRequest.Username).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("用户不存在", c)
		log.Warn(fmt.Sprintf("%s用户名不存在", EmailLoginRequest.Username))
		return
	}

	//校验密码
	fmt.Println(userModel.Password)
	fmt.Println(EmailLoginRequest.Password)

	isCheck := utils.CheckPwd(userModel.Password, EmailLoginRequest.Password)

	if !isCheck {
		global.Log.Warn("用户名密码错误")
		log.Warn(fmt.Sprintf("%s 用户名密码错误", EmailLoginRequest.Username))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}

	token, err := jwt.GenToken(jwt.JwtPayload{
		Nickname: userModel.NickName,
		Role:     uint(userModel.Role),
		UserId:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", c)
		log.Error(fmt.Sprintf("token生成失败%s", err.Error()))
		return
	}

	ip, addr := utils.GetAddrByGin(c)
	log = logstash.New(ip, token)
	log.Info("登录成功")

	//记录一下登录的数据
	global.DB.Create(&models.LoginDataModel{
		UserId:    userModel.ID,
		IP:        userModel.IP,
		Nickname:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})

	res.OKWithData(token, c)
}
