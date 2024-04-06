package data_api

import (
	"blog/gin/global"
	"blog/gin/models"
	"blog/gin/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type DataCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
type DateCountResponse struct {
	DateList  []string `json:"date_list"`
	LoginData []int    `json:"login_data"`
	SignData  []int    `json:"sign_data"`
}

func (DataApi) SevenLogin(c *gin.Context) {
	var loginDateCount, signDateCount []DataCount
	global.DB.Model(models.LoginDataModel{}).
		Where("date_sub(curdate(),interval 7 day) <= created_at").
		Select("date_format(created_at,'%Y-%m-%d') as date", "count(id) as count").
		Group("date").Scan(&loginDateCount)

	global.DB.Model(models.UserModel{}).
		Where("date_sub(curdate(),interval 7 day) <= created_at").
		Select("date_format(created_at,'%Y-%m-%d') as date", "count(id) as count").
		Group("date").Scan(&signDateCount)

	var loginDateMap = map[string]int{}

	for _, v := range loginDateCount {
		loginDateMap[v.Date] = v.Count
	}

	var loginCountMap = map[string]int{}
	for _, v := range signDateCount {
		loginCountMap[v.Date] = v.Count
	}

	now := time.Now()

	var dataList []string
	var loginCountList, signCountList []int

	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02")
		loginCount := loginCountMap[day]
		signCount := loginDateMap[day]
		dataList = append(dataList, day)

		loginCountList = append(loginCountList, loginCount)
		signCountList = append(signCountList, signCount)
	}

	fmt.Println(loginDateCount, signDateCount)

	res.OKWithData(DateCountResponse{
		DateList:  dataList,
		LoginData: loginCountList,
		SignData:  signCountList,
	}, c)
}
