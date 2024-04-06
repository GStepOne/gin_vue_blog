package main

import (
	"fmt"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"net"
)

func main() {
	//加载数据库
	db, _ := geoip2db.NewGeoipDbByStatik()
	defer db.Close()
	record, _ := db.City(net.ParseIP("115.35.95.90"))
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)

	fmt.Println("中文名字", record.Subdivisions[0].Names["zh-CN"])

}

func GetAddr(ip string) string {
	db, _ := geoip2db.NewGeoipDbByStatik()
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
