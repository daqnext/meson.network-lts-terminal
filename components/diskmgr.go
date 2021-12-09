package components

import (
	"github.com/daqnext/LocalLog/log"
	"github.com/daqnext/diskmgr"
	"github.com/daqnext/downloadmgr"
)

func InitDiskMgr(
	folderInfo []diskmgr.ProvideFolderInfo,
	onSuccess func(task *downloadmgr.Task),
	onFail func(task *downloadmgr.Task),
	onCancel func(task *downloadmgr.Task),
	onDownloading func(task *downloadmgr.Task),
	onLowSpeed func(task *downloadmgr.Task),
	absDbFolder string,
	localLogger *log.LocalLog,
) (*diskmgr.DiskMgr, error) {
	//////// init diskMgr   //////////////////////
	return diskmgr.New(
		folderInfo,
		onSuccess,
		onFail,
		onCancel,
		onDownloading,
		onLowSpeed,
		absDbFolder,
		localLogger,
	)
}
