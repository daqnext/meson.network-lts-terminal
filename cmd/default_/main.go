package default_

import (
	"time"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/cmd/default_/controllers"
	"github.com/daqnext/meson.network-lts-terminal/tools"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"

	"github.com/universe-30/USafeGo"
)

func startJobs() {
	// start BGJob ////////

	//heartBeat

	//uploadStatus

	//scanDisk

	//scanExpireFiles

	//checkTlsCertificate

	//checkPublicKey

	//updateRandomKey
	//global.RandKeyMgr.ScheduleUpdateRandomKey()
}

func StartDefault(clictx *cli.Context) {
	color.Green(basic.Logo)

	// init resourc
	//global.InitResources()
	//defer func() {
	//	global.ReleaseResources()
	//}()

	//api
	controllers.RunApi()

	//safeGo
	USafeGo.Go(
		//process
		func(args ...interface{}) {
			basic.Logger.Debugln("example of USafeGo")
			time.Sleep(10 * time.Second)
		},
		//onPanic callback
		tools.PanicHandler)

	for i := 0; i < 10; i++ {
		basic.Logger.Infoln("running")
		time.Sleep(1 * time.Second)
	}

	//httpServer example
	//global.HttpServer.GET("/test", func(context echo.Context) error {
	//	return context.String(200, "test success")
	//})

	//test log
	go func() {
		count := 0
		for {
			count++
			if count == 10 {
				//global.EchoServer.Shutdown()
				go func() {
					//err := global.EchoServer.Restart()
					//if err != nil {
					//	servercli.LocalLogger.Infoln(err)
					//}

					//servercli.LocalLogger.Debugln("try to colse")
					//global.EchoServer.Shutdown()
					//global.ReleaseResource()
					//os.Exit(0)

					//global.EchoServer.Shutdown()
					//time.Sleep(5 * time.Second)
					//global.EchoServer.StartTLS([]byte(chain), []byte(key))
				}()
			}
			if count >= 60 {

			}
			time.Sleep(time.Second)
			//servercli.LocalLogger.Debugln("default app test running")
		}
	}()

}
