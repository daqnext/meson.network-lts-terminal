package signMgr

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/src/destinationMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/globalData"
	"github.com/daqnext/meson.network-lts-terminal/src/randomKeyMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/requestUtil"
)

var keyContent = `-----BEGIN meson_PublicKey-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA24jvMVpBLPrIKdiuXO7Z
nNmMbUCAflKjBWeEFQILruh0vA0MQ/QVFZ0ykgz+rHXpZ3DcwGSt5VNkK0lwcZL1
53wCP5TIu+CCS41KJ8j4CCLC0AP3/rYUM68hhAyA2Y9ELPbiZVXunEmChXi5FJHc
Xq1pqxGzIM2Z1TbADHqpJAogb0lm9cyTb0jvknoYrtBCBLJ0Yijw5ASZ64CZzCun
y2PGEOrrBMZK8g1j5GN5gnNi35R0RQyQrDS2Um6crOI2n6YwWLB6Uu5NumSSZmNn
cZ/UB/IOuZDDijBRwSRDbp/D2Dz+r6hYrtRtKMswd25rPawfDBKHcp1QnZ7T5Gw7
iQIDAQAB
-----END meson_PublicKey-----`

type SignMgr struct {
	PublicKey *rsa.PublicKey
}

var signMgr *SignMgr

func Init() {
	signMgr = &SignMgr{}
}

func GetSingleInstance() *SignMgr {
	return signMgr
}

func parsePublicKey(buf []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.New("publicKey error")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	v, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("interface to *rsa.PublicKey error")
	}
	return v, nil
}
func (s *SignMgr) GetAndParsePublicKey() error {
	//get sign PublicKey from server
	url := destinationMgr.GetSingleInstance().GetDestUrl("/api/terminal/publickey")
	resp, err := requestUtil.Get(url, nil, 30, globalData.Token)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != 200 {
		return errors.New("GetAndParsePublicKey response status error:" + resp.Response().Status)
	}

	buf := resp.Bytes()
	pubKey, err := parsePublicKey(buf)
	if err != nil {
		return err
	}
	s.PublicKey = pubKey
	return nil
}

func (s *SignMgr) validateSignature(signContent string, sign string) bool {
	if s.PublicKey == nil {
		err := s.GetAndParsePublicKey()
		if err != nil {
			basic.Logger.Errorln(err)
			return false
		}
	}

	hashed := sha256.Sum256([]byte(signContent))
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		basic.Logger.Debugln(err)
		return false
	}
	err = rsa.VerifyPKCS1v15(s.PublicKey, crypto.SHA256, hashed[:], sig)
	if err != nil {
		basic.Logger.Debugln("rsa2 public check sign failed.", "err", err)
		return false
	}
	return true
}

func (s *SignMgr) CheckRequestLegal(sign string) bool {
	currentKey, previousKey := randomKeyMgr.GetSingleInstance().GetRandomKey()
	pass := s.validateSignature(currentKey, sign)
	if !pass {
		pass = s.validateSignature(previousKey, sign)
	}

	return pass
}
