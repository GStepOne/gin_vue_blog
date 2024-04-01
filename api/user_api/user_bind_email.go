package user_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"blog/gin/plugins/email"
	"blog/gin/utils"
	"blog/gin/utils/jwt"
	"blog/gin/utils/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password" msg:"请输入密码"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	//1 用户绑定邮箱，第一次输入是邮箱，
	//2 后台会给这个邮箱发验证码
	//3 第二次 用户输入邮箱的验证码和密码
	//4 完成绑定
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	session := sessions.Default(c)
	//说明是第一次，后台发验证码
	if cr.Code == nil {
		//生成验证码
		verifyCode := random.Code(4)
		session.Set("valid_code", verifyCode)
		err = session.Save() //一定要save
		err = email.NewCode().Send(cr.Email, "你的验证码是:"+verifyCode)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("邮箱发送失败", c)
			return
		}
		res.OKWithMessage("验证码已发送到邮箱,请注意查收", c)
		return
	}

	code := session.Get("valid_code")
	if *cr.Code != code {
		res.FailWithMessage("验证码不正确", c)
		return
	}

	var user models.UserModel
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	err = global.DB.Take(&user, claims.UserId).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": utils.HashPwd(cr.Password),
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}

	res.OKWithMessage("绑定邮箱成功", c)
	return

}
