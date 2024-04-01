package main

import (
	"blog/gin/core"
	"blog/gin/utils/jwt"
	"fmt"
)

func main() {
	core.InitCoreConf()
	token, err := jwt.GenToken(jwt.JwtPayload{
		UserId:   1,
		Role:     1,
		UserName: "jack",
		Nickname: "lee",
	})

	fmt.Println(token, err)

	claims, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImphY2siLCJuaWNrbmFtZSI6ImxlZSIsInJvbGUiOjEsInVzZXJJZCI6MSwiZXhwIjoxNzExOTg1MjU3LjQyNTI5OCwiaXNzIjoieHgifQ.CtHPgAQuE1lSIrTbP_QmhiDmmK-HHPFzgYkQV2xrgGY")
	fmt.Println(claims)
}
