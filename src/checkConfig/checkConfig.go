package checkConfig

import (
	"fmt"
	"log"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
)

func checkConfig() {
	//token
	token, err := configuration.Config.GetString("token", "")
	if err != nil {
		basic.Logger.Errorln("Get token form config err:", err)
	}
	if token == "" {
		fmt.Println("Can not find your token. Please login https://meson.network")
		fmt.Printf("Please input your token: ")
		_, err := fmt.Scanln(&token)
		if err != nil {
			log.Fatalln("read input token error:", err)
		}
	}

	//provide folder

	//port

}
