package config

import (
	meson_msg "github.com/daqnext/meson-msg"
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/daqnext/meson.network-lts-terminal/src/destMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/provideFolderHandler"
	"github.com/daqnext/meson.network-lts-terminal/src/requestUtil"
	"github.com/urfave/cli/v2"
)

func ConfigSetting(clictx *cli.Context) {

	//if clictx.IsSet("port") {
	//	configModify = handleAddPath(clictx)
	//}
	//
	//if clictx.IsSet("log_level") {
	//	configModify = handleAddPath(clictx)
	//}
	//
	//if clictx.IsSet("token") {
	//	clictx.String("token")
	//	configModify = handleAddPath(clictx)
	//}
	//
	//if clictx.IsSet("dest") {
	//	configModify = handleAddPath(clictx)
	//}

	if clictx.IsSet("addpath") {
		addPath(clictx)
	}

	if clictx.IsSet("removepath") {
		removePath(clictx)
	}

	//if configModify {
	//	err := configuration.Config.WriteConfig()
	//	if err != nil {
	//		color.Red("config save error:", err)
	//		return
	//	}
	//	fmt.Println("config modified new config")
	//	fmt.Println(configuration.Config.GetConfigAsString())
	//}
}

func setPort(clictx *cli.Context) {
	newPort := clictx.Int("port")
	//pre check token

	destMgr.Init()
	err := destMgr.GetSingleInstance().SearchAvailableDest()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//if fail
	basic.Logger.Fatalln("err")

	//if pass
	configuration.Config.Set("port", newPort)
	configuration.Config.WriteConfig()

}

func setLogLevel(clictx *cli.Context) {
	newLevel := clictx.String("log_level")
	//pre check token

	destMgr.Init()
	err := destMgr.GetSingleInstance().SearchAvailableDest()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//if fail
	basic.Logger.Fatalln("err")

	//if pass
	configuration.Config.Set("log_level", newLevel)
	configuration.Config.WriteConfig()
}

func setToken(clictx *cli.Context) {
	newToken := clictx.String("token")
	//pre check token

	destMgr.Init()
	err := destMgr.GetSingleInstance().SearchAvailableDest()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	url := destMgr.GetSingleInstance().GetDestUrl("/api/tokenvalidator")

	resp, err := requestUtil.Post(url, nil, meson_msg.ValidateToken{Token: newToken}, 15, "")
	if err != nil {
		basic.Logger.Fatalln("network err")
	}

	_ = resp

	//if fail
	basic.Logger.Fatalln("invalid token, please login https://meson.network")

	//if pass
	configuration.Config.Set("token", newToken)
	configuration.Config.WriteConfig()
}

func setDest(clictx *cli.Context) {
	newDest := clictx.String("dest")
	//pre check dest

	//if fail
	basic.Logger.Fatalln("error")

	//if pass
	configuration.Config.Set("dest", newDest)
	configuration.Config.WriteConfig()
}

func addPath(clictx *cli.Context) {
	pathToAdd := clictx.String("addpath")
	err := provideFolderHandler.HandleAddPath(pathToAdd)
	if err != nil {
		basic.Logger.Fatalln(err)
	}
}

func removePath(clictx *cli.Context) {
	pathToRemove := clictx.String("removepath")
	err := provideFolderHandler.HandleRemovePath(pathToRemove)
	if err != nil {
		basic.Logger.Fatalln(err)
	}
}

//func handleAddPath(clictx *cli.Context) (modified bool) {
//
//	//read exist path from config
//	provideFolder, err := configuration.Config.GetProvideFolders()
//	if err == configuration.ErrProvideFolderType || err == configuration.ErrProvideFolderContent {
//		fmt.Println("the exist provide folder configuration is invalid, it will be deleted")
//	}
//
//	//transfer path to abs path
//	folderToAdd := clictx.String("addpath")
//	folderToAdd = filepath.Clean(folderToAdd)
//	if !filepath.IsAbs(folderToAdd) {
//		folderToAdd = path_util.GetAbsPath(folderToAdd)
//	}
//
//	for _, v := range provideFolder {
//		if v.AbsPath == folderToAdd {
//			return errors.New(fmt.Sprintf("The path <%s> is already exist", folderToAdd))
//		}
//	}
//
//	//input size
//	var size int
//	for {
//		fmt.Println("Please input provider folder size(For example if you provide 100GB, please input 100, only support integer).")
//		fmt.Printf("40GB disk space is the minimum. Please make sure you have enough free space:")
//		_, err := fmt.Scanln(&size)
//		if err != nil {
//			fmt.Println("read input size error:", err)
//			continue
//		}
//		if size < 20 {
//			fmt.Errorf("minimum size is 40 GB")
//			continue
//		}
//		break
//	}
//
//	//diskMgr check folder size
//
//	newFolder := folderMgr.FolderConfig{
//		AbsPath: folderToAdd,
//		SizeGB:  size,
//	}
//
//	provideFolder = append(provideFolder, newFolder)
//
//	configuration.Config.SetProvideFolders(provideFolder)
//	err = configuration.Config.SafeWriteConfig()
//	if err != nil {
//
//	}
//	fmt.Println("new folder added:", folderToAdd, "size:", size, "GB")
//	return nil
//
//}
//
//func handleRemovePath(clictx *cli.Context) (modified bool) {
//
//	provideFolders := []ProvideFolder{}
//	provide_folder := configuration.Config.Get("provide_folder", nil)
//	iArray, ok := provide_folder.([]interface{})
//	if !ok {
//
//		//...
//		return
//	}
//	value := clictx.String("removepath")
//	removed := false
//	for _, v := range iArray {
//		m := v.(map[string]interface{})
//		if m["abs_path"].(string) == value {
//			removed = true
//			continue
//		}
//		pf := ProvideFolder{
//			AbsPath: m["abs_path"].(string),
//			SizeGB:  int(m["size_GB"].(float64)),
//		}
//		provideFolders = append(provideFolders, pf)
//	}
//
//	if removed {
//		configuration.Config.Set("provide_folder", provideFolders)
//		fmt.Println("path removed:", value)
//		return true
//	} else {
//		fmt.Println("removepath failed, path not exist")
//		return false
//	}
//
//}
