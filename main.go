package main

import (
	"os"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/cmd"
)

// for swagger
// @title           Meson Terminal API
// @version         1.0
// @description     meson terminal's api
// @termsOfService  https://meson.network

// @contact.name    Meson Support
// @contact.url     https://meson.network
// @contact.email   contact@meson.network

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            spec-00-02-015-bchkakakbxxxxxx-019-thisisthebinddomain.mesontrackingdomain.com
// @schemes         https

func main() {
	basic.InitLogger()

	//config app to run
	errRun := cmd.ConfigCmd().Run(os.Args)
	if errRun != nil {
		basic.Logger.Panicln(errRun)
	}
}
