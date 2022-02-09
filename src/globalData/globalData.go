package globalData

import (
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
)

var Token string

func Init() {
	var err error
	Token, err = configuration.Config.GetString("token", "")
	if err != nil {
		basic.Logger.Fatalln("get token form config error:", err)
	}
	if Token == "" {
		basic.Logger.Fatalln("get token form config error:", err)
	}
	//ServerHost, err = configuration.Config.GetString("server_host", "http://center.coldcdn.com")
	//if err != nil {
	//	basic.Logger.Fatalln("get server_host form config error:", err)
	//}
	//PublicIpfsGateway, err = configuration.Config.GetString("public_ipfs_gateway", "https://cloudflare-ipfs.com/ipfs/")
	//if err != nil {
	//	basic.Logger.Fatalln("get public_ipfs_gateway form config error:", err)
	//}
	//ExpireTimeHour, err = configuration.Config.GetInt("expire_time_hour", 48)
	//if err != nil {
	//	basic.Logger.Fatalln("get expire_time_hour form config error:", err)
	//}
}
