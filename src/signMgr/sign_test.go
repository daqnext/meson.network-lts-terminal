package signMgr

import (
	"log"
	"testing"
)

func Test_v(t *testing.T) {
	Init()
	err := GetSingleInstance().GetAndParsePublicKey()
	if err != nil {
		log.Println(err)
	}
}
