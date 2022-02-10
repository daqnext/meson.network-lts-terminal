package basic

import (
	"os"

	"github.com/fatih/color"
	"github.com/universe-30/LogrusULog"
	"github.com/universe-30/ULog"
	"github.com/universe-30/UUtils/path_util"
)

var Logger ULog.Logger

func InitLogger() {
	var llerr error
	Logger, llerr = LogrusULog.New(path_util.GetAbsPath("logs"), 2, 20, 30)

	if llerr != nil {
		color.Set(color.FgRed)
		defer color.Unset()
		panic("Error:" + llerr.Error())
	}
}

func InitUnitTestLogger() {
	var llerr error
	Logger, llerr = LogrusULog.New(path_util.GetAbsPath("unitTestLogs"), 2, 20, 30)
	l := ULog.ParseLogLevel("DEBU")
	Logger.SetLevel(l)
	Logger.SetOutput(os.Stdout)

	if llerr != nil {
		color.Set(color.FgRed)
		defer color.Unset()
		panic("Error:" + llerr.Error())
	}
}
