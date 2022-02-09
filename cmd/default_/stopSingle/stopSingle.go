package stopSingle

import (
	"github.com/daqnext/meson.network-lts-terminal/basic"
)

var StopChannel = make(chan struct{})

func StopTerminal() {
	StopChannel <- struct{}{}
}

func WaitingForStopSingle() {
	<-StopChannel
	basic.Logger.Infoln("Terminal stop")
}
