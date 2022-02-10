package preCheckConfig

import (
	"fmt"
	"strconv"

	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/daqnext/meson.network-lts-terminal/src/destMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/diskFileMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/echoServer"
	"github.com/daqnext/meson.network-lts-terminal/src/versionMgr"
)

func PreCheckConfig() error {
	destMgr.Init()
	versionMgr.Init()
	diskFileMgr.Init()

	//search dest
	err := destMgr.GetSingleInstance().SearchAvailableDest()
	if err != nil {
		return err
	}

	//check version
	versionMgr.GetSingleInstance().CheckUpdate()

	//check token
	checkToken()

	//check port
	checkPort()

	//check folder

	return nil

}

func checkToken() {
	token, _ := configuration.Config.GetString("token", "")
	//if err != nil {
	//	basic.Logger.Errorln("Get token form config err:", err)
	//}
	for {
		if token != "" {
			//check token on server

			//if pass
			pass := true
			if pass {
				break
			} else {
				fmt.Println("invalid token")
				token = ""
			}
		}

		fmt.Println("Can not find your token. Please login https://meson.network")
		fmt.Printf("Please input your token: ")
		_, err := fmt.Scanln(&token)
		if err != nil {
			fmt.Println("read input token error:", err)
		}
	}
}

func checkPort() {
	port, _ := configuration.Config.GetInt("port", 0)
	//if err != nil {
	//	return errors.New("port [int] in config error," + err.Error())
	//}
	for {
		var myport string
		for {
			if port == 0 {
				fmt.Printf("Please input your port. strongly recommend use 443. \nPlease use a port >10000 if you cannot use port 443(default 443): ")
				_, err := fmt.Scanln(&myport)
				if err != nil {
					fmt.Println("input error:", err)
					continue
				}
				if myport == "" {
					port = 443
					break
				}
				port, err = strconv.Atoi(myport)
				if err != nil {
					fmt.Println("input error:", err)
					continue
				}
				break
			}
		}

		_, exist := echoServer.DisablePortMap[port]
		if exist || (port < 10000 && port != 443) {
			fmt.Printf("Can not use port [%d], please use 443 or >10000")
			continue
		}
		break
	}
}

func checkFolder() {
	provideFolder, err := configuration.Config.GetProvideFolders()
	if err != nil {
		fmt.Println("provide folder config err:", err)
		fmt.Println("Please reset folders")
	} else if len(provideFolder) == 0 {
		fmt.Println("provide folder not set, please input a folder path and size(GB) for meson file storage")
	}

	var pathStr string
	fmt.Println("Please input the folder path for meson:")

}
