package daemonService

import (
	"runtime"
	"sync"

	"github.com/daqnext/daemon"
	"github.com/daqnext/meson.network-lts-terminal/basic"
)

const (
	// name of the service
	name        = "meson"
	description = "meson terminal"
)

type Service struct {
	daemon.Daemon
}

var service *Service
var once sync.Once

func Init() {
	//only run once
	once.Do(func() {
		service = newDaemonService()
	})
}

func GetSingleInstance() *Service {
	Init()
	return service
}

func newDaemonService() *Service {
	kind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		kind = daemon.UserAgent
	}
	srv, err := daemon.New(name, description, kind)
	if err != nil {
		basic.Logger.Fatalln("run daemon error:", err)
	}
	return &Service{srv}
}
