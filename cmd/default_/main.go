package default_

import (
	"net/http"
	"time"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/cmd/default_/controllers"
	"github.com/daqnext/meson.network-lts-terminal/cmd/default_/stopSingle"
	"github.com/daqnext/meson.network-lts-terminal/plugin/cache"
	"github.com/daqnext/meson.network-lts-terminal/src/destinationMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/diskFileMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/echoServer"
	"github.com/daqnext/meson.network-lts-terminal/src/randomKeyMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/signMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/statusMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/tlsCertificateMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/versionMgr"
	"github.com/daqnext/meson.network-lts-terminal/tools"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"

	"github.com/universe-30/USafeGo"
)

func initComponent() {
	cache.Init()
	destinationMgr.Init()
	randomKeyMgr.Init()
	signMgr.Init()
	statusMgr.Init()
	tlsCertificateMgr.Init()
	versionMgr.Init()

	err := echoServer.Init()
	if err != nil {
		basic.Logger.Fatalln("echoServer init error:", err)
	}

	err = diskFileMgr.Init()
	if err != nil {
		basic.Logger.Fatalln("diskFileMgr init error:", err)
	}

}

//func startDiskMgr() {
//	//disk check and sync provide folder
//	//read provide Info from config
//	provideFolder, err := configuartion.Config.GetProvideFolders()
//	if err != nil {
//		//basic.Logger.Fatalln("http server start failed")
//		//todo handle get provideFolder err
//	}
//	err = diskFileMgr.GetSingleInstance().AddProvideFolder(provideFolder)
//	if err != nil {
//		basic.Logger.Fatalln("Provide folder err:", err)
//	}
//	err = diskFileMgr.GetSingleInstance().CheckFolderSpace()
//	if err != nil {
//		basic.Logger.Fatalln("Check folder space err:", err)
//	}
//
//	missingFiles, errorFiles, err := diskFileMgr.GetSingleInstance().ScanDiskFileInDb()
//	if err != nil {
//		basic.Logger.Errorln(err)
//	}
//	if len(missingFiles) > 0 {
//		//send to server
//	}
//	if len(errorFiles) > 0 {
//		//send to server
//	}
//}

func startEchoServer() {
	//api
	controllers.DeclareApi()
	//start echo server
	go func() {
		chain, key := tlsCertificateMgr.GetSingleInstance().GetTlsCert()
		err := echoServer.GetSingleInstance().StartTLS(chain, key)
		if err != nil {
			if err == http.ErrServerClosed {
				basic.Logger.Debugln(err)
			} else {
				basic.Logger.Errorln(err)
			}
		}
	}()
	//wait for echo server start
	err := echoServer.GetSingleInstance().WaitForServerStart(true)
	if err != nil {
		basic.Logger.Fatalln("http server start failed")
	}
	basic.Logger.Infoln("echo server started on port:", echoServer.GetSingleInstance().Http_port)
}

func StartDefault(clictx *cli.Context) {
	color.Green(basic.Logo)

	// init component
	initComponent()

	//check destination

	//check version

	//check config
	//input missing config

	//public key for sign

	//domain and tls for echo server

	//diskMgr
	//startDiskMgr()

	//echo server
	startEchoServer()
	stopSingle.WaitingForStopSingle()

	//job after echo start success

	////////////
	////test////
	basic.Logger.Debugln("test part")
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
