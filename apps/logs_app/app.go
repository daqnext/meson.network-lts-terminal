package logs_app

import (
	fj "github.com/daqnext/fastjson"
	clitool "github.com/daqnext/meson.network-lts-terminal/cli"
	"github.com/urfave/cli/v2"
)

func StartLog(ConfigJson *fj.FastJson, CliContext *cli.Context) {

	num := CliContext.Int("num")
	if num == 0 {
		num = 20
	}

	onlyerr := CliContext.Bool("onlyerr")
	if onlyerr {
		clitool.LocalLogger.PrintLastN_ErrLogs(num)
	} else {
		clitool.LocalLogger.PrintLastN_AllLogs(num)
	}
}
