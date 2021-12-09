package service_app

import (
	"fmt"
	fj "github.com/daqnext/fastjson"
	mesoncli "github.com/daqnext/meson.network-lts-terminal/cli"
	"github.com/takama/daemon"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime"
)

const (
	// name of the service
	name        = "meson"
	description = "meson terminal"
)

// Service has embedded daemon
type Service struct {
	daemon.Daemon
}

func (service *Service) Manage() (string, error) {

	usage := "Usage: myservice install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}
	return "", nil
}

func RunServerCmd(ConfigJson *fj.FastJson, CliContext *cli.Context) {
	subCmds := CliContext.Command.Names()
	if len(subCmds) == 0 {
		log.Println("no sub command")
		return
	}
	kind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		kind = daemon.UserAgent
	}
	srv, err := daemon.New(name, description, kind)
	if err != nil {
		mesoncli.LocalLogger.Fatalln("run daemon error:", err)
	}
	service := &Service{srv}
	action := subCmds[0]
	log.Println(action)
	if action == "config" {
		log.Println(CliContext.String("token"))
	}
	//return

	var status string
	var e error
	switch action {
	case "install":
		status, e = service.Install()
		log.Println("cmd install")
	case "remove":
		service.Stop()
		status, e = service.Remove()
		log.Println("cmd remove")
	case "start":
		status, e = service.Start()
		log.Println("cmd start")
	case "stop":
		status, e = service.Stop()
		log.Println("cmd stop")
	case "restart":
		service.Stop()
		status, e = service.Start()
		log.Println("cmd restart")
	case "status":
		status, e = service.Status()
		log.Println("cmd status")
	case "config":
		log.Println("cmd config")
		s, err := service.Stop()
		if err != nil {
			log.Println(err)
			//return
		}
		//log.Println(s)
		s, err = service.Start()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(s)
	default:
		log.Println("no sub command")
		return
	}

	if e != nil {
		fmt.Println(status, "\nError: ", e)
		os.Exit(1)
	}
	fmt.Println(status)
}
