package versionMgr

const version = "3.0.0"

type VersionMgr struct {
	CurrentVersion string
}

var versionMgr *VersionMgr

func Init() {
	versionMgr = &VersionMgr{
		CurrentVersion: version,
	}
}

func GetSingleInstance() *VersionMgr {
	return versionMgr
}

func (v *VersionMgr) CheckUpdate() {

}
