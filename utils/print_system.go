package utils

import "blog/gin/global"

func PrintSystem() {
	ip := global.Config.System.Host
	port := global.Config.System.Port
	if ip == "0.0.0.0" {
		ipList := GetIPList()
		for _, i := range ipList {
			global.Log.Infof("程序gvb server 运行在 %s:%d", i, port)
		}
	} else {
		global.Log.Infof("程序gvb server 运行在 %s:%d", ip, port)
	}
}
