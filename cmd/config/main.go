package config

import (
	"fmt"

	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func ConfigSetting(clictx *cli.Context) {

	configModify := false

	for _, v := range stringConfParams {
		if clictx.IsSet(v) {
			newValue := clictx.String(v)
			configuration.Config.Set(v, newValue)
			configModify = true
		}
	}

	for _, v := range float64ConfParams {
		if clictx.IsSet(v) {
			newValue := clictx.Float64(v)
			configuration.Config.Set(v, newValue)
			configModify = true
		}
	}

	for _, v := range boolConfPrams {
		if clictx.IsSet(v) {
			newValue := clictx.Bool(v)
			configuration.Config.Set(v, newValue)
			configModify = true
		}
	}

	if clictx.IsSet("addpath") {
		configModify = handleAddPath(clictx)
	}

	if clictx.IsSet("removepath") {
		configModify = handleRemovePath(clictx)
	}

	if configModify {
		err := configuration.Config.WriteConfig()
		if err != nil {
			color.Red("config save error:", err)
			return
		}
		fmt.Println("config modified new config")
		fmt.Println(configuration.Config.GetConfigAsString())
	}
}

type ProvideFolder struct {
	AbsPath string `json:"abs_path""`
	SizeGB  int    `json:"size_GB"`
}

func handleAddPath(clictx *cli.Context) (modified bool) {

	provideFolders := []ProvideFolder{}
	provide_folder := configuration.Config.Get("provide_folder", nil)
	iArray, ok := provide_folder.([]interface{})
	if !ok {

		//...
		return false
	}
	folderPath := clictx.String("addpath")
	//check path legal
	//...

	for _, v := range iArray {
		m := v.(map[string]interface{})
		if m["abs_path"].(string) == folderPath {
			fmt.Println("path already exist")
			return false
		}
		pf := ProvideFolder{
			AbsPath: m["abs_path"].(string),
			SizeGB:  int(m["size_GB"].(float64)),
		}
		provideFolders = append(provideFolders, pf)
	}

	//input size
	var size int

	fmt.Printf("Please input provider folder size: ")
	_, err := fmt.Scanln(&size)
	if err != nil {
		fmt.Errorf("read input size error: %s", err.Error())
		return false
	}
	if size < 20 {
		fmt.Errorf("minimum size is 20 GB")
		return false
	}

	pf := ProvideFolder{
		AbsPath: folderPath,
		SizeGB:  size,
	}
	provideFolders = append(provideFolders, pf)
	configuration.Config.Set("provide_folder", provideFolders)
	fmt.Println("new folder added:", folderPath, "size", size, "GB")
	return true

}

func handleRemovePath(clictx *cli.Context) (modified bool) {

	provideFolders := []ProvideFolder{}
	provide_folder := configuration.Config.Get("provide_folder", nil)
	iArray, ok := provide_folder.([]interface{})
	if !ok {

		//...
		return
	}
	value := clictx.String("removepath")
	removed := false
	for _, v := range iArray {
		m := v.(map[string]interface{})
		if m["abs_path"].(string) == value {
			removed = true
			continue
		}
		pf := ProvideFolder{
			AbsPath: m["abs_path"].(string),
			SizeGB:  int(m["size_GB"].(float64)),
		}
		provideFolders = append(provideFolders, pf)
	}

	if removed {
		configuration.Config.Set("provide_folder", provideFolders)
		fmt.Println("path removed:", value)
		return true
	} else {
		fmt.Println("removepath failed, path not exist")
		return false
	}

}
