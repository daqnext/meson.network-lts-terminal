package config_app

import (
	fj "github.com/daqnext/fastjson"
	clitool "github.com/daqnext/meson.network-lts-terminal/cli"
	"github.com/daqnext/utils/path_util"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
)

func ConfigSetting(ConfigJson *fj.FastJson, CliContext *cli.Context) {
	dev := CliContext.Bool("dev")
	var defaultConfigPath string
	if dev {
		defaultConfigPath = path_util.GetAbsPath("configs/dev/" + "default.json")
	} else {
		defaultConfigPath = path_util.GetAbsPath("configs/pro/" + "default.json")
	}

	var config *fj.FastJson
	var err error
	config, err = fj.NewFromFile(defaultConfigPath)
	if err != nil {
		log.Println(err)
		dir := filepath.Dir(defaultConfigPath)
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			//log error

			return
		}
		config = fj.NewFromString("")
	}

	//modify token
	newToken := CliContext.String("token")
	log.Println("newToken:", newToken)
	if newToken != "" {

		//logInfo if success
	}

	//modify port
	newPort := CliContext.Int("port")
	clitool.LocalLogger.Infoln("newPort:", newPort)
	if newPort != 0 {
		//check port
		//not in chrome block list

		config.SetInt(newPort, "http_port")

		//logInfo if success
	}

	//modify logLevel

	//add provide folder
	newFolder := CliContext.String("addfolder")
	log.Println("newFolder:", newFolder)

	//modify provide folder size

	log.Println(config.GetContentAsString())
	os.WriteFile(defaultConfigPath, config.GetContent(), 0666)

	//log info
	//restart terminal to use new config
}
