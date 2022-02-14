package service

import (
	"fmt"
	"log"
	"os"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/daqnext/meson.network-lts-terminal/plugin/daemonService"
	"github.com/daqnext/meson.network-lts-terminal/src/checkConfig"
	"github.com/urfave/cli/v2"
)

func RunServiceCmd(clictx *cli.Context) {
	//check command
	subCmds := clictx.Command.Names()
	if len(subCmds) == 0 {
		basic.Logger.Fatalln("no sub command")
		return
	}

	action := subCmds[0]
	daemonService.Init()
	compDeamon := daemonService.GetSingleInstance()

	var status string
	var e error
	switch action {
	case "install":
		status, e = compDeamon.Install()
		basic.Logger.Debugln("cmd install")
	case "remove":
		compDeamon.Stop()
		status, e = compDeamon.Remove()
		basic.Logger.Debugln("cmd remove")
	case "start":
		//check config
		checkConfig.PreCheckConfig()
		status, e = compDeamon.Start()
		basic.Logger.Debugln("cmd start")
	case "stop":
		status, e = compDeamon.Stop()
		basic.Logger.Debugln("cmd stop")
	case "restart":
		compDeamon.Stop()
		//check config
		checkConfig.PreCheckConfig()
		status, e = compDeamon.Start()
		basic.Logger.Debugln("cmd restart")
	case "status":
		status, e = compDeamon.Status()
		basic.Logger.Debugln("cmd status")
	default:
		basic.Logger.Debugln("no sub command")
		return
	}

	if e != nil {
		fmt.Println(status, "\nError: ", e)
		os.Exit(1)
	}
	fmt.Println(status)
}

func checkAndSetConfig() {
	//token
	token, err := configuration.Config.GetString("token", "")
	if err != nil {
		fmt.Println("Get token form config err:", err)
	}
	if token == "" {
		fmt.Println("Can not find your token. Please login https://meson.network")
		fmt.Printf("Please enter your token: ")
		_, err := fmt.Scanln(&token)
		if err != nil {
			log.Fatalln("read input token error")
		}
	}

	//verify token in server

	//port
	port, err := configuration.Config.GetInt("port", 0)
	if err != nil {
		fmt.Println("Get port form config err:", err)
	}
	if port == 0 {
		fmt.Printf("Please enter your port, 443 is the most recommended(default 443): ")
		var myport string
		_, err := fmt.Scanln(&myport)
		if err != nil {
			//UsingPort = "19091"
			fmt.Println("read input port error,server will be run in port:443.You can modify this value with command: ./meson config set -port=xxx")
		}
	}
}
