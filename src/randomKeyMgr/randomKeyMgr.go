package randomKeyMgr

import (
	"github.com/daqnext/meson.network-lts-terminal/tools"
	"github.com/universe-30/UJob"
	"github.com/universe-30/UUtils/rand_util"
)

type RandomKeyMgr struct {
	currentKey  string
	previousKey string
}

var randomKeyMgr *RandomKeyMgr

func Init() {
	if randomKeyMgr != nil {
		return
	}
	randomKeyMgr = &RandomKeyMgr{}
}

func GetSingleInstance() *RandomKeyMgr {
	return randomKeyMgr
}

//ScheduleUpdateRandomKey start a background job. This job will generate new random key periodically
func (r *RandomKeyMgr) ScheduleUpdateRandomKey() {
	UJob.Start(
		//job process
		func() {
			r.GenNewRandomKey()
		},
		//onPanic callback
		tools.PanicHandler,
		300,
		UJob.TYPE_PANIC_REDO,
		nil,
		nil,
	)
}

//GenNewRandomKey generate a new random key as currentKey, and the old one will be record as previousKey
func (r *RandomKeyMgr) GenNewRandomKey() string {
	randKey := rand_util.GenRandStr(10)
	if r.currentKey == "" {
		r.previousKey = randKey
	} else {
		r.previousKey = r.currentKey
	}
	r.currentKey = randKey
	return randKey
}

//CheckRandomKey check the given randomKey is equal to the current key or previous key
func (r *RandomKeyMgr) CheckRandomKey(inputKey string) bool {
	if inputKey == r.currentKey || inputKey == r.previousKey {
		return true
	}
	return false
}

func (r *RandomKeyMgr) GetRandomKey() (currentKey string, previousKey string) {
	if r.currentKey == "" {
		r.GenNewRandomKey()
	}
	return r.currentKey, r.previousKey
}
