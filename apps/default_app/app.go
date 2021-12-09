package default_app

import (
	fj "github.com/daqnext/fastjson"
	_ "github.com/daqnext/meson.network-lts-terminal/apps/default_app/controllers"
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/global"
	servercli "github.com/daqnext/meson.network-lts-terminal/cli"
	"github.com/urfave/cli/v2"
	"time"
)

func startJobs() {
	// start BGJob ////////

	//heartBeat

	//uploadStatus

	//scanDisk

	//scanExpireFiles

	//checkTlsCertificate

	//updateRandomKey
}

func StartDefault(ConfigJson *fj.FastJson, CliContext *cli.Context) {
	global.Wg.Add(1)
	defer func() {
		servercli.LocalLogger.Infoln("StartDefault closed , start to ReleaseResource()")
		global.ReleaseResource()
	}()

	//check config

	//check destination

	//check version

	//check sign

	//start EchoServer
	//start the http server
	go func() {
		err := global.EchoServer.Start()
		if err != nil {
			servercli.LocalLogger.Infoln(err)
		}
	}()
	//wait for echo server start
	err := global.EchoServer.Echo.WaitForServerStart(true)
	if err != nil {
		servercli.LocalLogger.Fatalln("http server start failed")
	}

	//login in server

	//wait request from server

	//start diskMgr

	//

	startJobs()

	//test log
	go func() {
		count := 0
		for {
			count++
			if count == 5 {
				global.EchoServer.Shutdown()
			}
			if count >= 60 {
				global.Wg.Done()
			}
			time.Sleep(time.Second)
			servercli.LocalLogger.Infoln("default app test running")
		}
	}()

	global.Wg.Wait()

	//stop log
	servercli.LocalLogger.Infoln("default app finish")
}

func init() {

}
