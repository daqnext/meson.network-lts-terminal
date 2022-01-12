package randomKeyMgr

import (
	"sync"

	"github.com/daqnext/meson.network-lts-terminal/tools"
	"github.com/universe-30/UJob"
	"github.com/universe-30/UUtils/rand_util"
)

type RandomKeyMgr struct {
	currentKey  string
	previousKey string
}

var randomKeyMgr *RandomKeyMgr
var once sync.Once

func Init() {
	//only run once
	once.Do(func() {
		randomKeyMgr = new()
	})
}

func GetSingleInstance() *RandomKeyMgr {
	Init()
	return randomKeyMgr
}

func new() *RandomKeyMgr {
	r := &RandomKeyMgr{}
	return r
}

func (r *RandomKeyMgr) ScheduleUpdateRandomKey() {
	UJob.Start(
		//job process
		func() {
			r.GenNewRandomKey()
		},
		//onPanic callback
		tools.PanicHandler,
		300,
		// job type
		// UJob.TYPE_PANIC_REDO  auto restart if panic
		// UJob.TYPE_PANIC_RETURN  stop if panic
		UJob.TYPE_PANIC_REDO,
		// check continue callback, the job will stop running if return false
		// the job will keep running if this callback is nil
		nil,
		// onFinish callback
		nil,
	)
}

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

func (r *RandomKeyMgr) CheckRandomKey(inputKey string) bool {
	if inputKey == r.currentKey || inputKey == r.previousKey {
		return true
	}
	return false
}

func (r *RandomKeyMgr) GetCurrentRandomKey() string {
	if r.currentKey == "" {
		r.GenNewRandomKey()
	}
	return r.currentKey
}
