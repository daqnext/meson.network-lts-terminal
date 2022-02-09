package signMgr

import "testing"

func Test_v(t *testing.T) {
	Init()
	GetSingleInstance().GetAndParsePublicKey()
}
