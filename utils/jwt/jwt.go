package jwt

import (
	"blog/gin/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type JwtPayload struct {
	UserName string `json:"username"`
	Nickname string `json:"nickname"`
	Role     uint   `json:"role"`
	UserId   uint   `json:"userId"`
}

type CustomClaims struct {
	JwtPayload
	jwt.StandardClaims
}

func GenToken(user JwtPayload) (string, error) {

	var Secret = []byte(global.Config.JWT.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.JWT.Expires))),
			Issuer:    global.Config.JWT.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(Secret)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	var Secret = []byte(global.Config.JWT.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		global.Log.Error(fmt.Sprintf("token parse err:%s", err.Error()))
		return nil, err
	}

	//括号是断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
