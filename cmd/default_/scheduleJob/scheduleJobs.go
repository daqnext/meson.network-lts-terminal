package scheduleJob

import (
	"github.com/daqnext/meson.network-lts-terminal/src/randomKeyMgr"
	"github.com/daqnext/meson.network-lts-terminal/tools"
)

func heartBeat() {

}

func sendTerminalStatus() {

}

func scanDisk() {

}

func scanExpireFiles() {

}

func checkTlsCertificate() {

}

func checkPublicKey() {

}

func StartJobs() {
	// start BGJob ////////

	//heartBeat

	//uploadStatus

	//scanDisk

	//scanExpireFiles

	//checkTlsCertificate

	//checkPublicKey

	//updateRandomKey
	randomKeyMgr.GetSingleInstance().ScheduleUpdateRandomKey()

	//uploadPanic
	tools.ScheduleUploadPanic()
}
