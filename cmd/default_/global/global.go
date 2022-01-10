package global

import (
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/components"
	"github.com/go-redis/redis/v8"
	elasticSearch "github.com/olivere/elastic/v7"
	"github.com/universe-30/RedisSpr"
	"github.com/universe-30/UCache"
	"gorm.io/gorm"
)

//define your components here
var Cache *UCache.Cache
var HttpServer *components.EchoServer
var EsClient *elasticSearch.Client
var Redis *redis.ClusterClient
var SprMgr *RedisSpr.SprJobMgr
var DB *gorm.DB

func IniResources() {
	var err error

	Cache = components.NewUCache()

	HttpServer, err = components.NewEchoServer()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//EsClient, err = components.NewElasticSearch()
	//if err != nil {
	//	basic.Logger.Fatalln(err)
	//}

	Redis, err = components.NewRedis()
	if err != nil {
		//error as example
		basic.Logger.Errorln(err)
		//panic is suggest in production app
		//basic.Logger.Fatalln(err)
	}

	SprMgr, err = components.NewRedisSpr()
	if err != nil {
		//error as example
		basic.Logger.Errorln(err)
		//panic is suggest in production app
		//basic.Logger.Fatalln(err)
	}

	DB, _, err = components.NewDB()
	if err != nil {
		//error as example
		basic.Logger.Errorln(err)
		//panic is suggest in production app
		//basic.Logger.Fatalln(err)
	}
}

func ReleaseResources() {

}
