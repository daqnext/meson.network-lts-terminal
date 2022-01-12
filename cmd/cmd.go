package cmd

import (
	"os"
	"strings"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/cmd/config"
	"github.com/daqnext/meson.network-lts-terminal/cmd/default_"
	"github.com/daqnext/meson.network-lts-terminal/cmd/log"
	"github.com/daqnext/meson.network-lts-terminal/cmd/service"
	"github.com/daqnext/meson.network-lts-terminal/configuartion"
	"github.com/universe-30/ULog"
	"github.com/universe-30/UUtils/path_util"
	"github.com/urfave/cli/v2"
)

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_LOG = "logs"
const CMD_NAME_SERVICE = "service"
const CMD_NAME_CONFIG = "config"

////////config to do cmd ///////////
func ConfigCmd() *cli.App {
	//check is dev or pro
	isDev := false
	for index, arg := range os.Args {
		s := strings.ToLower(arg)
		if s == "-dev=true" {
			isDev = true
			os.Args = append(os.Args[:index], os.Args[index+1:]...)
			break
		}
	}
	conferr := iniConfig(isDev)
	if conferr != nil {
		basic.Logger.Panicln(conferr)
	}

	return &cli.App{
		Action: func(clictx *cli.Context) error {
			path_util.ExEPathPrintln()
			logerr := setLoggerLevel()
			if logerr != nil {
				return logerr
			}
			default_.StartDefault(clictx)
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:  CMD_NAME_LOG,
				Usage: "print all logs",
				Flags: log.GetFlags(),
				Action: func(clictx *cli.Context) error {
					path_util.ExEPathPrintln()
					log.StartLog(clictx)
					return nil
				},
			},
			{
				Name:  CMD_NAME_CONFIG,
				Usage: "config command",
				Subcommands: []*cli.Command{
					//show config
					{
						Name:  "show",
						Usage: "show configs",
						Action: func(clictx *cli.Context) error {
							logerr := setLoggerLevel()
							if logerr != nil {
								return logerr
							}
							return nil
						},
					},
					//set config
					{
						Name:  "set",
						Usage: "set config",
						Flags: config.GetFlags(),
						Action: func(clictx *cli.Context) error {
							path_util.ExEPathPrintln()
							logerr := setLoggerLevel()
							if logerr != nil {
								return logerr
							}
							config.ConfigSetting(clictx)
							return nil
						},
					},
				},
			},
			{
				Name:  CMD_NAME_SERVICE,
				Usage: "service command",
				Subcommands: []*cli.Command{
					//service install
					{
						Name:  "install",
						Usage: "install meson node in service",
						Action: func(clictx *cli.Context) error {
							logerr := setLoggerLevel()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service remove
					{
						Name:  "remove",
						Usage: "remove meson node from service",
						Action: func(clictx *cli.Context) error {
							logerr := setLoggerLevel()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service start
					{
						Name:  "start",
						Usage: "run",
						Action: func(clictx *cli.Context) error {
							logerr := setLoggerLevel()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service stop
					{
						Name:  "stop",
						Usage: "stop",
						Action: func(clictx *cli.Context) error {
							logerr := setLoggerLevel()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service restart
					{
						Name:  "restart",
						Usage: "restart",
						Action: func(clictx *cli.Context) error {
							logerr := setLoggerLevel()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service status
					{
						Name:  "status",
						Usage: "show process status",
						Action: func(clictx *cli.Context) error {
							logerr := setLoggerLevel()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
				},
			},
		},
	}
}

////////end config to do app ///////////
func readDefaultConfig(isDev bool) (*configuartion.VConfig, string, error) {
	var defaultConfigPath string
	if isDev {
		basic.Logger.Infoln("======== using dev mode ========")
		defaultConfigPath = path_util.GetAbsPath("configs/dev.json")
	} else {
		basic.Logger.Infoln("======== using pro mode ========")
		defaultConfigPath = path_util.GetAbsPath("configs/pro.json")
	}

	basic.Logger.Infoln("config file:", defaultConfigPath)

	config, err := configuartion.ReadConfig(defaultConfigPath)
	if err != nil {
		basic.Logger.Errorln("no pro.json under /configs folder , use --dev=true to run dev mode")
		return nil, "", err
	} else {
		return config, defaultConfigPath, nil
	}
}

func iniConfig(isDev bool) error {
	//path_util.ExEPathPrintln()
	////read default config
	config, _, err := readDefaultConfig(isDev)
	if err != nil {
		return err
	}
	basic.Logger.Infoln("======== start of config ========")
	configs, _ := config.GetConfigAsString()
	basic.Logger.Infoln(configs)
	basic.Logger.Infoln("======== end  of  config ========")
	configuartion.Config = config
	return nil
}

func setLoggerLevel() error {
	logLevel := "INFO"
	if configuartion.Config != nil {
		var err error
		logLevel, err = configuartion.Config.GetString("local_log_level", "INFO")
		if err != nil {
			return err
		}
	}

	l := ULog.ParseLogLevel(logLevel)
	basic.Logger.SetLevel(l)
	return nil
}
