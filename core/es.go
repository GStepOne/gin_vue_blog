package core

import (
	"blog/gin/global"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := global.Config.Es.URL()
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.Es.User, global.Config.Es.Password),
	)
	if err != nil {
		logrus.Fatalf("es链接失败%s", err.Error())
	}

	return client
}
