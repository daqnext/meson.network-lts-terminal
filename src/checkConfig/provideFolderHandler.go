package checkConfig

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/daqnext/diskmgr"
	"github.com/daqnext/diskmgr/folderMgr"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/daqnext/meson.network-lts-terminal/src/diskFileMgr"
	"github.com/universe-30/UUtils/path_util"
)

func HandleAddPath(newFolderPath string) (newPath string, sizeGB int, provideFolder []folderMgr.FolderConfig, err error) {
	//read exist path from config
	provideFolder, err = configuration.Config.GetProvideFolders()
	if err == configuration.ErrProvideFolderType || err == configuration.ErrProvideFolderContent {
		fmt.Println("the exist provide folder configuration is invalid, it will be deleted")
	}

	//transfer path to abs path
	folderToAdd := filepath.Clean(newFolderPath)
	if !filepath.IsAbs(folderToAdd) {
		folderToAdd = path_util.GetAbsPath(folderToAdd)
	}

	for _, v := range provideFolder {
		if v.AbsPath == folderToAdd {
			return "", 0, nil, errors.New(fmt.Sprintf("The path <%s> is already exist", folderToAdd))
		}
	}

	//input size
	var size int
	for {
		fmt.Println("Please input provider folder size(For example if you provide 100GB, please input 100, only support integer).")
		fmt.Printf("%d GB disk space is the minimum. Please make sure you have enough free space:", diskFileMgr.BottomSizeGB)
		_, err := fmt.Scanln(&size)
		if err != nil {
			fmt.Println("read input size error:", err)
			continue
		}
		if size < diskFileMgr.BottomSizeGB {
			fmt.Printf("minimum size is %d GB\n", diskFileMgr.BottomSizeGB)
			continue
		}
		break
	}

	//check folder size
	err = diskmgr.CheckFolder(folderToAdd, size, diskFileMgr.CheckLimitGB, diskFileMgr.BottomSizeGB)
	if err != nil {
		return "", 0, nil, err
	}

	newFolder := folderMgr.FolderConfig{
		AbsPath: folderToAdd,
		SizeGB:  size,
	}

	provideFolder = append(provideFolder, newFolder)

	return folderToAdd, size, provideFolder, nil
}

func HandleRemovePath(pathToRemove string) (removedPath string, provideFolder []folderMgr.FolderConfig, err error) {
	//read exist path from config
	provideFolder, err = configuration.Config.GetProvideFolders()
	if err == configuration.ErrProvideFolderType || err == configuration.ErrProvideFolderContent {
		fmt.Println("the exist provide folder configuration is invalid, it will be deleted")
	}

	//transfer path to abs path
	folderToRemove := filepath.Clean(pathToRemove)
	if !filepath.IsAbs(folderToRemove) {
		folderToRemove = path_util.GetAbsPath(folderToRemove)
	}

	exist := false
	for i, v := range provideFolder {
		if v.AbsPath == folderToRemove {
			exist = true
			provideFolder = append(provideFolder[:i], provideFolder[i+1:]...)
			break
		}
	}

	if !exist {
		return "", nil, errors.New(fmt.Sprintf("The path <%s> is not exist", folderToRemove))
	}

	return folderToRemove, provideFolder, nil
}
