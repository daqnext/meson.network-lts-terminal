package tlsCertificateMgr

import (
	"sync"
)

type CertMgr struct {
	ChainFileUrl string
	KeyFileUrl   string
}

var certMgr *CertMgr
var once sync.Once

func Init() {
	//only run once
	once.Do(func() {
		certMgr = new()
	})
}

func GetSingleInstance() *CertMgr {
	Init()
	return certMgr
}

func new() *CertMgr {
	cm := &CertMgr{}
	return cm
}

func (c *CertMgr) getChainFile() {

}

func (c *CertMgr) getKeyFile() {

}
