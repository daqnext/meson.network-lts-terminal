package main

import (
	_ "github.com/daqnext/meson.network-lts-terminal/apps/default_app/global"
	"github.com/daqnext/meson.network-lts-terminal/apps/service_app"
	"github.com/daqnext/meson.network-lts-terminal/cli"

	configapp "github.com/daqnext/meson.network-lts-terminal/apps/config_app"
	defaultapp "github.com/daqnext/meson.network-lts-terminal/apps/default_app"
	logsapp "github.com/daqnext/meson.network-lts-terminal/apps/logs_app"
)

func main() {
	switch cli.AppToDO.AppName {
	case cli.APP_NAME_LOG:
		logsapp.StartLog(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	case cli.APP_NAME_SERVICE:
		service_app.RunServerCmd(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	case cli.APP_NAME_CONFIG:
		configapp.ConfigSetting(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	default:
		cli.LocalLogger.Infoln("======== start default app ===")
		defaultapp.StartDefault(cli.AppToDO.ConfigJson, cli.AppToDO.CliContext)
	}
}
