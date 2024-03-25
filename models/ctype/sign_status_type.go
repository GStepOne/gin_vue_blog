package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1
	SingGitee SignStatus = 2
	SignEmail SignStatus = 3
	SignWeiBo SignStatus = 4
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SingGitee:
		str = "Gitee"
	case SignEmail:
		str = "Email"
	case SignWeiBo:
		str = "Weibo"
	default:
		str = "其他"
	}

	return str
}
