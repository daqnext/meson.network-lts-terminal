package daemonservice

import "github.com/takama/daemon"

var SystemDConfig = `[Unit]
Description={{.Description}}
Requires={{.Dependencies}}
After={{.Dependencies}}
[Service]
PIDFile=/var/run/{{.Name}}.pid
ExecStartPre=/bin/rm -f /var/run/{{.Name}}.pid
ExecStart={{.Path}} {{.Args}}
Restart=always
[Install]
WantedBy=multi-user.target
`

// Service is the daemon service struct
type Service struct {
	daemon.Daemon
}

const (
	// name of the service
	name        = "meson"
	description = "meson terminal"
)

var MesonService *Service

func ServiceInstall() {

}

func ServiceRemove() {

}

func ServiceStart() {

}

func ServiceStop() {

}

func ServiceStatus() {

}
