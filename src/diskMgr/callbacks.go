package diskMgr

import (
	"github.com/daqnext/downloadmgr"
)

//download callback
func onDownloadSuccess(task *downloadmgr.Task) {

}
func onDownloadFail(task *downloadmgr.Task) {

}
func onDownloadCancel(task *downloadmgr.Task) {

}
func onDownloading(task *downloadmgr.Task) {

}
func onDownloadSlowSpeed(task *downloadmgr.Task) {

}

//cache file issue callback
func onCachedFileDeleted(deletedFile []string) {

}
func onDownloadingFileDeleted(deletedFile []string) {

}
func onFileMissing(missingFile []string) {

}
