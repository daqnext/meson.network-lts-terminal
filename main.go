package main

import (
	"os"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/cmd"
)

func main() {
	basic.InitLogger()

	//config app to run
	errRun := cmd.ConfigCmd().Run(os.Args)
	if errRun != nil {
		basic.Logger.Panicln(errRun)
	}
}
