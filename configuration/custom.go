package configuration

import (
	"errors"

	"github.com/daqnext/diskmgr/folderMgr"
)

var ErrProvideFolderNotSet = errors.New("provide_folder not find in config")
var ErrProvideFolderType = errors.New("provide_folder in config type error")
var ErrProvideFolderContent = errors.New("provide_folder content type error")

//example read provide_folder
func (c *VConfig) GetProvideFolders() ([]folderMgr.FolderConfig, error) {
	key := "provide_folder"
	if !c.Viper.IsSet(key) {
		return nil, ErrProvideFolderNotSet
	}

	provide_folder := c.Viper.Get("provide_folder")
	iArray, ok := provide_folder.([]interface{})
	if !ok {
		return nil, ErrProvideFolderType
	}

	provideFolders := []folderMgr.FolderConfig{}
	for _, v := range iArray {
		m, ok := v.(map[string]interface{})
		if !ok {
			return nil, ErrProvideFolderContent
		}
		pf := folderMgr.FolderConfig{
			AbsPath: m["abs_path"].(string),
			SizeGB:  int(m["size_GB"].(float64)),
		}
		provideFolders = append(provideFolders, pf)
	}
	return provideFolders, nil
}

func SetProvideFolders(pf []folderMgr.FolderConfig) {
	Config.Set("provide_folder", pf)
}
