package chat_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/ctype"
	"blog/gin/models/res"
	"blog/gin/utils"
	"encoding/json"
	"fmt"
	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"time"
)

var ConnGroupMap = map[string]ChatUser{}

const (
	InRoomMsg  ctype.MessageType = 1
	TextMsg    ctype.MessageType = 2
	VideoMsg   ctype.MessageType = 3
	VoiceMsg   ctype.MessageType = 4
	ImageMsg   ctype.MessageType = 5
	SystemMsg  ctype.MessageType = 6
	OutRoomMsg ctype.MessageType = 7
)

type GroupRequest struct {
	Content     string            `json:"content"`
	MessageType ctype.MessageType `json:"message_type"`
}

type GroupResponse struct {
	NickName    string            `json:"nick_name"`
	Avatar      string            `json:"avatar"`
	Content     string            `json:"content"`
	MessageType ctype.MessageType `json:"message_type"`
	Date        time.Time         `json:"date"`         //消息的时间
	OnlineCount int               `json:"online_count"` //在线人数
}

type ChatUser struct {
	Conn     *websocket.Conn
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (ChatApi) ChatGroup(c *gin.Context) {
	var upGreader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			//鉴权 true 表示放行 false 表示拦截
			return true
		},
	}
	//表示升级为websocket
	conn, err := upGreader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	addr := conn.RemoteAddr().String()
	nickname := randomname.GenerateName()
	avatar := fmt.Sprintf("uploads/avatar/%s.png", nickname)
	chatUser := ChatUser{
		Conn:     conn,
		Nickname: nickname,
		Avatar:   avatar,
	}
	ConnGroupMap[addr] = chatUser

	//需要生成昵称，映射头像地址
	global.Log.Infof("%s 🔗成功", addr)
	for {
		_, p, err := conn.ReadMessage()
		//进行参数绑定
		if err != nil {
			//用户断开聊天
			SendGroupMessage(conn, GroupResponse{
				Content:     fmt.Sprintf("%s 离开聊天室", chatUser.Nickname),
				MessageType: OutRoomMsg,
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap) - 1,
				NickName:    chatUser.Nickname,
				Avatar:      chatUser.Avatar,
			})
			break
		}

		var request GroupRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			global.Log.Infof("聊天内容不合法:%s", p)
			SendMessage(addr, GroupResponse{
				Content:     "参数绑定失败",
				MessageType: SystemMsg,
			})
			continue
		}

		switch request.MessageType {
		case TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				global.Log.Error("聊天内容为空")
				SendMessage(addr, GroupResponse{
					Content:     "聊天内容为空",
					MessageType: SystemMsg,
				})
				continue
			}
			SendGroupMessage(conn, GroupResponse{
				NickName:    chatUser.Nickname,
				Avatar:      chatUser.Avatar,
				Content:     request.Content,
				MessageType: TextMsg,
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap),
			})
		case InRoomMsg:
			request.Content = fmt.Sprintf("%s 进入聊天室", chatUser.Nickname)
			SendGroupMessage(conn, GroupResponse{
				NickName:    chatUser.Nickname,
				Avatar:      chatUser.Avatar,
				Content:     request.Content,
				MessageType: InRoomMsg,
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap),
			})
		default:
			SendMessage(addr, GroupResponse{
				Content:     "聊天内容错误",
				MessageType: SystemMsg,
			})
		}
	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// SendGroupMessage 发给每一个链接
func SendGroupMessage(conn *websocket.Conn, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	_addr := conn.RemoteAddr().String()
	ip, addr := getIpAndAddr(_addr)

	global.DB.Create(&models.ChatModel{
		NickName:    response.NickName,
		Avatar:      response.Avatar,
		Content:     response.Content,
		MessageType: response.MessageType,
		IP:          ip,
		Addr:        addr,
		IsGroup:     true,
	})
	for _, user := range ConnGroupMap {
		user.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// SendMessage 单独给某个用户发信息
func SendMessage(addr string, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	ChatUser := ConnGroupMap[addr]
	ip, addr := getIpAndAddr(addr)
	global.DB.Create(&models.ChatModel{
		NickName:    response.NickName,
		Avatar:      response.Avatar,
		Content:     response.Content,
		MessageType: response.MessageType,
		IP:          ip,
		Addr:        addr,
		IsGroup:     false,
	})
	ChatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
}

func getIpAndAddr(addr string) (string, userAddr string) {
	ips := strings.Split(addr, ":")
	addr = utils.GetAddr(ips[0])
	return ips[0], addr
}
