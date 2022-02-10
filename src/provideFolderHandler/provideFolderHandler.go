package provideFolderHandler

import (
	"fmt"

	"github.com/daqnext/diskmgr/folderMgr"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/urfave/cli/v2"
)

func handleAddPath(newFolderPath string) (modified bool, err error) {

	provideFolders := []folderMgr.FolderConfig{}
	provide_folder := configuration.Config.Get("provide_folder", nil)
	iArray, ok := provide_folder.([]interface{})
	if !ok {

		//...
		return false
	}
	folderPath := new
	//check path legal
	//...

	for _, v := range iArray {
		m := v.(map[string]interface{})
		if m["abs_path"].(string) == newFolderPath {
			fmt.Println("path already exist")
			return false
		}
		pf := folderMgr.FolderConfig{
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

	pf := folderMgr.FolderConfig{
		AbsPath: folderPath,
		SizeGB:  size,
	}
	provideFolders = append(provideFolders, pf)
	configuration.Config.Set("provide_folder", provideFolders)
	fmt.Println("new folder added:", folderPath, "size", size, "GB")
	return true

}

func handleRemovePath(clictx *cli.Context) (modified bool) {

	provideFolders := []folderMgr.FolderConfig{}
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
		pf := folderMgr.FolderConfig{
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
