package diskFileMgr

import (
	"github.com/daqnext/diskmgr"
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/daqnext/meson.network-lts-terminal/tools"
	"github.com/universe-30/UUtils/path_util"
)

var dm diskmgr.IDiskMgr

const CheckLimitGB = 5 //todo set 40 in production
const BottomSizeGB = 5

func Init() error {

	var err error
	dm, err = new()
	if err != nil {
		return err
	}
	return nil
}

func GetSingleInstance() diskmgr.IDiskMgr {
	return dm
}

func new() (diskmgr.IDiskMgr, error) {

	//read provide Info from config
	provideFolder, err := configuration.Config.GetProvideFolders()
	if err != nil {
		return nil, err
	}

	d := diskmgr.New(basic.Logger, tools.PanicHandler)

	err = d.StartUp(
		provideFolder,
		path_util.GetAbsPath("db"),
		onDownloadSuccess,
		onDownloadFail,
		onDownloadCancel,
		onDownloading,
		onDownloadSlowSpeed,
		onCachedFileDeleted,
		onDownloadingFileDeleted,
		onFileMissing,
	)

	if err != nil {
		return nil, err
	}

	return d, nil
}
