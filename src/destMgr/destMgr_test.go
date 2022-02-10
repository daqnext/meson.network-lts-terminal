package destMgr

import (
	"log"
	"testing"

	"github.com/daqnext/meson.network-lts-terminal/basic"
)

func init() {
	basic.InitUnitTestLogger()
	destMgr = &DestMgr{
		backupDest: map[string]struct{}{
			"coldcdn.com": struct{}{},
		},
	}
	destMgr.genBackupDest()
}

func Test_printBackupDest(t *testing.T) {
	for k := range destMgr.backupDest {
		log.Println(k)
	}
}

func Test_SearchAvailableDest(t *testing.T) {
	destMgr.SearchAvailableDest()
}
