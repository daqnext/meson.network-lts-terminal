package versionMgr

import (
	"sync"
)

const version = "3.0.0"

type VersionMgr struct {
	CurrentVersion string
}

var versionMgr *VersionMgr
var once sync.Once

func GetSingleInstance() *VersionMgr {
	once.Do(func() {
		versionMgr = new()
	})
	return versionMgr
}

func new() *VersionMgr {
	v := &VersionMgr{
		CurrentVersion: version,
	}

	return v
}

func (v *VersionMgr) CheckUpdate() {

}
