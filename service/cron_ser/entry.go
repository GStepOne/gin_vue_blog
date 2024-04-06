package cron_ser

import (
	"blog/gin/global"
	"github.com/robfig/cron/v3"
	"time"
)

func CronInt() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	//秒 分 时 日 月 周
	Cron.AddFunc("0 0 0 * * *", SyncArticleData)
	Cron.AddFunc("0 1 0 * * *", syncCommentData)

	global.Log.Infof("今日文章数据同步已经完成:日期%s", time.Now().Format("2006-01-02 15:03:04"))
}
