package core

import (
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"log"
)

func InitAddrDB() *geoip2.DBReader {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		log.Fatal("ip 地址数据库加载失败", err)
	}
	return db
}
