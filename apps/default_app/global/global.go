package global

import (
	"github.com/daqnext/BGJOB_GO/bgjob"
	SPR_go "github.com/daqnext/SPR-go"
	gofastcache "github.com/daqnext/go-fast-cache"
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/manager/destination"
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/manager/randomkey"
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/manager/sign"
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/manager/status"
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/manager/tlscertificate"
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/manager/version"
	"github.com/daqnext/meson.network-lts-terminal/cli"
	"github.com/daqnext/meson.network-lts-terminal/components"
	"github.com/labstack/echo/v4/middleware"
	"sync"
)

var GLOBAL_INIT_FINISHED bool

var StopChan = make(chan bool)
var Wg = sync.WaitGroup{}

var EchoServer *components.EchoServer
var SpMgr *SPR_go.SprJobMgr
var BGJobM *bgjob.JobManager
var LocalCache *gofastcache.LocalCache

//
var DestMgr *destination.DestMgr
var RandKeyMgr *randomkey.RandomKeyMgr
var SignMgr *sign.SignMgr
var StatusMgr *status.StatusMgr
var CertMgr *tlscertificate.CertMgr
var VersionMgr *version.VersionMgr

func init() {
	if !cli.AppIsActive(cli.APP_NAME_DEFAULT) {
		return
	}

	//init your global components
	//logger
	err := components.InitLocalLog(cli.LocalLogger, cli.AppToDO.ConfigJson)
	if err != nil {
		panic(err.Error())
	}

	//localCache
	LocalCache = components.InitFastCache(cli.LocalLogger)

	//BGJob
	components.InitSmartRoutine()
	BGJobM = components.InitBGJobs(cli.LocalLogger)

	cli.LocalLogger.Info("init system .....")
	////////////ini more components config as you need///////////////////

	//echo http server
	EchoServer, err = components.InitEchoServer(cli.LocalLogger, cli.AppToDO.ConfigJson)
	if err != nil {
		cli.LocalLogger.Fatal(err.Error())
	}
	EchoServer.Echo.Use(middleware.Recover())
	EchoServer.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "=> ${time_rfc3339_nano} | ${method} | ${remote_ip} | ${host} | ${uri} | ${status} | ${bytes_out} B | ${latency_human} |${error}\n",
		Output: cli.LocalLogger.Out,
	}))

	//

	cli.LocalLogger.Info("=========== end of init system ==================")

	GLOBAL_INIT_FINISHED = true
}

func StopNode() {
	StopChan <- true
}

func ReleaseResource() {
	if EchoServer != nil {
		EchoServer.Close()
	}
}
