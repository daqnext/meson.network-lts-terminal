package config

import (
	"github.com/urfave/cli/v2"
)

//get all config flags
func GetFlags() (allflags []cli.Flag) {
	return []cli.Flag{
		&cli.IntFlag{Name: "port", Required: false},
		&cli.StringFlag{Name: "log_level", Required: false},
		&cli.StringFlag{Name: "token", Required: false},
		&cli.StringFlag{Name: "dest", Required: false},
		&cli.StringFlag{Name: "addpath", Required: false},
		&cli.StringFlag{Name: "removepath", Required: false},
	}
}
