package versionMgr

import (
	"github.com/daqnext/meson.network-lts-common/mesonUtils"
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/src/destMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/requestUtil"
)

const version = "3.0.0"

type VersionMgr struct {
	CurrentVersion string
}

var versionMgr *VersionMgr

func Init() {
	if versionMgr != nil {
		return
	}
	versionMgr = &VersionMgr{
		CurrentVersion: version,
	}
}

func GetSingleInstance() *VersionMgr {
	return versionMgr
}

func (v *VersionMgr) GetVersion() string {
	return v.CurrentVersion
}

func (v *VersionMgr) getTerminalVersionFromServer() (latestVersion string, allowVersion string, err error) {
	//check is there new version or not
	basic.Logger.Debugln("Check Version...")
	url := destMgr.GetSingleInstance().GetDestUrl("/api/common/terminalversion")
	resp, err := requestUtil.Get(url, nil, 15, "")
	if err != nil {
		basic.Logger.Errorln("GetTerminalVersionFromServer request error", err)
		return "", "", err
	}
	var i int
	err = resp.ToJSON(&i)
	if err != nil {
		basic.Logger.Errorln("GetTerminalVersionFromServer resp.ToJSON error", err)
		return "", "", err
	}

	return "", "", nil
}

func (v *VersionMgr) IsLatestVersion() (bool, error) {
	latestVersion, _, err := v.getTerminalVersionFromServer()
	if err != nil {
		return true, err
	}

	r := mesonUtils.VersionCompare(v.CurrentVersion, latestVersion)
	if r >= 0 {
		return true, nil
	}
	return false, nil
}

func (v *VersionMgr) CheckUpdate() {
	isLatestVersion, _ := v.IsLatestVersion()
	if isLatestVersion {
		return
	}

	//download new version
}
