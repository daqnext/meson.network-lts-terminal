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

func Init() {
	//only run once
	once.Do(func() {
		versionMgr = new()
	})
}

func GetSingleInstance() *VersionMgr {
	Init()
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
