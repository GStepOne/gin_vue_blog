package utils

import (
	"fmt"
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"github.com/gin-gonic/gin"
	"net"
)

var db *geoip2.DBReader

func init() {
	//写到这里很快
	db, _ = geoip2db.NewGeoipDbByStatik()
}

func GetAddrByGin(c *gin.Context) (ip, addr string) {
	ip = c.ClientIP()
	addr = GetAddr(ip)
	return ip, addr
}
func GetAddr(ip string) string {
	defer db.Close()
	parseIp := net.ParseIP(ip)
	if IsIntranetIP(parseIp) {
		return "内网地址"
	}

	record, err := db.City(parseIp)
	if err != nil {
		return "错误的ip地址"
	}

	var province string
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}

	city := record.City.Names["zh-CN"]

	return fmt.Sprintf("%s-%s", province, city)
}

// IsIntranetIP 是不是内联网
func IsIntranetIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}
	//192.168
	//10.
	//172.16-172.31
	//169.254
	ip4 := ip.To4()
	if ip4 == nil {
		return true
	}

	return ip4[0] == 192 && ip4[1] == 168 ||
		(ip4[0] == 172 && ip4[1] >= 16 || (ip4[0] == 172 && ip4[0] <= 31)) ||
		ip4[0] == 10 || (ip4[0] == 169 && ip4[1] == 254)

}
