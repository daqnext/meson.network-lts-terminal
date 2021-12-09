package randomkey

import "github.com/daqnext/utils/rand_util"

type RandomKeyMgr struct {
	currentKey  string
	previousKey string
}

func New() *RandomKeyMgr {
	r := &RandomKeyMgr{}
	return r
}

func (r *RandomKeyMgr) GenNewRandomKey() string {
	randKey := rand_util.GenRandStr(10)
	r.previousKey = r.currentKey
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
