package daemonService

import (
	"os"
	"runtime"

	"github.com/daqnext/daemon"
	"github.com/daqnext/meson.network-lts-terminal/basic"
)

const (
	// name of the service
	name        = "meson"
	description = "meson terminal"
)

var systemDConfig = `[Unit]
Description={{.Description}}
Requires={{.Dependencies}}
After=network.target nss-lookup.target
[Service]
PIDFile=/var/run/{{.Name}}.pid
ExecStartPre=/bin/rm -f /var/run/{{.Name}}.pid
ExecStart={{.Path}} {{.Args}}
Restart=always
RestartSec=7
[Install]
WantedBy=multi-user.target
`

type Service struct {
	daemon.Daemon
}

var service *Service

func Init() {
	if service != nil {
		return
	}
	service = newDaemonService()
}

func GetSingleInstance() *Service {
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
	s := &Service{srv}

	if _, err := os.Stat("/run/systemd/system"); err == nil {
		err := s.SetTemplate(systemDConfig)
		if err != nil {
			basic.Logger.Fatalln("MesonService SetTemplate error", "err", err)
		}
	}

	return s
}
