package version

const version = "3.0.0"

type VersionMgr struct {
	CurrentVersion string
}

func New() *VersionMgr {
	v := &VersionMgr{
		CurrentVersion: version,
	}

	return v
}

func (v *VersionMgr) CheckUpdate() {

}
