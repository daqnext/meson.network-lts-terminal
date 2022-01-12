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

func GetSingleInstance() *CertMgr {
	once.Do(func() {
		certMgr = new()
	})
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