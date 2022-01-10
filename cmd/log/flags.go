package log

import (
	"github.com/urfave/cli/v2"
)

//get default_ config flags
func GetFlags() (allflags []cli.Flag) {
	return []cli.Flag{
		&cli.IntFlag{Name: "num", Required: false},
		&cli.BoolFlag{Name: "onlyerr", Required: false},
	}
}
