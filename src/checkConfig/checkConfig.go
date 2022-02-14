package checkConfig

import (
	"fmt"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/daqnext/meson.network-lts-terminal/src/destMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/versionMgr"
	"github.com/universe-30/UUtils/path_util"
)

func PreCheckConfig() error {
	//search dest
	destMgr.Init()
	err := destMgr.GetSingleInstance().SearchAvailableDest()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//check version
	versionMgr.Init()
	versionMgr.GetSingleInstance().CheckUpdate()

	//check token
	preCheckToken()

	//check port
	preCheckPort()

	//check folder
	preCheckFolder()

	return nil
}

func preCheckToken() {
	token, err := configuration.Config.GetString("token", "")
	if err != nil {
		fmt.Println("Get token form config err:", err)
	}
	originToken := token

	for {
		for {
			if token != "" {
				break
			}
			var inputToken string
			fmt.Printf("Please input your token: ")
			_, err := fmt.Scanln(&inputToken)
			if err != nil {
				fmt.Println("read input error:", err)
				continue
			}
			token = inputToken
		}

		err := CheckToken(token)
		if err != nil {

			token = ""
			continue
		}
		break
	}

	if originToken != token {
		//save config
		configuration.Config.Set("token", token)
		err = configuration.Config.WriteConfig()
		if err != nil {
			//todo handle error
		}
	}
}

func preCheckPort() {
	port, err := configuration.Config.GetInt("port", 0)
	if err != nil {
		fmt.Println("Get port form config err:", err)
	}
	originPort := port

	for {
		for {
			if port != 0 {
				break
			}
			var inputPort int
			fmt.Printf("Please input your port. strongly recommend use 443. \nPlease use a port >10000 if you cannot use port 443(default 443): ")
			_, err := fmt.Scanln(&inputPort)
			if err != nil {
				fmt.Println("input error:", err)
				continue
			}
			if inputPort == 0 {
				port = 443
				break
			}
			port = inputPort
		}

		err := CheckPort(port)
		if err != nil {

			port = 0
			continue
		}
		break
	}

	if originPort != port {
		//save config
		configuration.Config.Set("port", port)
		err = configuration.Config.WriteConfig()
		if err != nil {
			//todo handle error
		}
	}

}

func preCheckFolder() {
	provideFolder, err := configuration.Config.GetProvideFolders()
	if err != nil {
		fmt.Println("provide folder config err:", err)
		fmt.Println("Please reset folders")
	} else if len(provideFolder) == 0 {
		fmt.Println("provide folder not set, please input a folder path and size(GB) for meson file storage")
	}

	for {
		if len(provideFolder) == 0 {

			var pathStr string
			defaultPath := path_util.GetAbsPath("mesonfolder")
			fmt.Printf("Please input the folder path for meson(default <%s>)\n", defaultPath)
			fmt.Printf("path:")
			_, err := fmt.Scanln(&pathStr)
			if err != nil {
				fmt.Println("input error:", err)
				continue
			}
			if pathStr == "" {
				pathStr = defaultPath
			}

			//var newPath string
			//var size int
			_, _, provideFolder, err = HandleAddPath(pathStr)
			if err != nil {
				fmt.Println(err)
				continue
			}

			//save config
			configuration.Config.SetProvideFolders(provideFolder)
			err = configuration.Config.WriteConfig()
			if err != nil {
				//todo handle error
			}
			break
		}
	}
}
