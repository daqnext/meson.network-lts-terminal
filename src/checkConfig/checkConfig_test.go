package checkConfig

import (
	"log"
	"testing"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
)

func init() {
	basic.InitUnitTestLogger()
	config, err := configuration.ReadConfig("/Users/bruce/workspace/go/project/meson.network-lts-terminal/configs/dev.json")
	if err != nil {
		log.Println(err)
	}
	configuration.Config = config
}

func Test_preCheckToken(t *testing.T) {

	preCheckToken() //fmt.Scanln can't run in test
}
