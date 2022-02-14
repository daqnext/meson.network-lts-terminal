package config

import (
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/daqnext/meson.network-lts-terminal/src/checkConfig"
	"github.com/urfave/cli/v2"
)

func ConfigSetting(clictx *cli.Context) {

	changeSetting := false

	if clictx.IsSet("port") {
		setPort(clictx)
		changeSetting = true
	}

	if clictx.IsSet("log_level") {
		setLogLevel(clictx)
		changeSetting = true
	}

	if clictx.IsSet("token") {
		setToken(clictx)
		changeSetting = true
	}

	if clictx.IsSet("dest") {
		setDest(clictx)
		changeSetting = true
	}

	if clictx.IsSet("addpath") {
		addPath(clictx)
		changeSetting = true
	}

	if clictx.IsSet("removepath") {
		removePath(clictx)
		changeSetting = true
	}

	if changeSetting {
		basic.Logger.Infoln("new configuration will take effect after restart")
	}

}

func setPort(clictx *cli.Context) {
	newPort := clictx.Int("port")
	//pre check port
	err := checkConfig.CheckPort(newPort)
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//if pass
	configuration.Config.Set("port", newPort)
	err = configuration.Config.WriteConfig()
	if err != nil {
		basic.Logger.Fatalln(err)
	}
	basic.Logger.Infoln("new port:", newPort)
}

func setLogLevel(clictx *cli.Context) {
	newLevel := clictx.String("log_level")

	//pre check log
	err := checkConfig.CheckLogLevel(newLevel)
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//if pass
	configuration.Config.Set("log_level", newLevel)
	err = configuration.Config.WriteConfig()
	if err != nil {
		basic.Logger.Fatalln(err)
	}
	basic.Logger.Infoln("new log level:", newLevel)
}

func setToken(clictx *cli.Context) {
	newToken := clictx.String("token")
	//pre check token
	err := checkConfig.CheckToken(newToken)
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//if pass
	configuration.Config.Set("token", newToken)
	err = configuration.Config.WriteConfig()
	if err != nil {
		basic.Logger.Fatalln(err)
	}
	basic.Logger.Infoln("new token:", newToken)
}

func setDest(clictx *cli.Context) {
	newDest := clictx.String("dest")
	//pre check dest
	err := checkConfig.CheckDest(newDest)
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//if pass
	configuration.Config.Set("dest", newDest)
	err = configuration.Config.WriteConfig()
	if err != nil {
		basic.Logger.Fatalln(err)
	}
	basic.Logger.Infoln("new dest:", newDest)
}

func addPath(clictx *cli.Context) {
	pathToAdd := clictx.String("addpath")
	newPath, sizeGB, provideFolder, err := checkConfig.HandleAddPath(pathToAdd)
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	configuration.Config.SetProvideFolders(provideFolder)
	err = configuration.Config.WriteConfig()
	if err != nil {
		basic.Logger.Fatalln(err)
	}
	basic.Logger.Infoln("new folder added:", newPath, "size:", sizeGB, "GB")
}

func removePath(clictx *cli.Context) {
	pathToRemove := clictx.String("removepath")
	removedPath, provideFolder, err := checkConfig.HandleRemovePath(pathToRemove)
	if err != nil {
		basic.Logger.Fatalln(err)
	}
	configuration.Config.SetProvideFolders(provideFolder)
	err = configuration.Config.WriteConfig()
	if err != nil {
		basic.Logger.Fatalln(err)
	}
	basic.Logger.Infoln("path removed:", removedPath)
}
