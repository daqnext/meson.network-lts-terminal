package sign

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/daqnext/meson.network-lts-terminal/cli"
	"time"
)

type SignMgr struct {
	PublicKey *rsa.PublicKey
}

func New() *SignMgr {
	return &SignMgr{}
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
	//url:=""+"/api/signkey"
	buf := []byte("data")
	pubKey, err := parsePublicKey(buf)
	if err != nil {
		return err
	}
	s.PublicKey = pubKey
	return nil
}

func (s *SignMgr) ValidateSignature(signContent string, sign string) bool {
	if s.PublicKey == nil {
		err := s.GetAndParsePublicKey()
		if err != nil {
			cli.LocalLogger.Errorln(err)
			return false
		}
	}

	hashed := sha256.Sum256([]byte(signContent))
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		cli.LocalLogger.Debugln(err)
		return false
	}
	err = rsa.VerifyPKCS1v15(s.PublicKey, crypto.SHA256, hashed[:], sig)
	if err != nil {
		cli.LocalLogger.Debugln("rsa2 public check sign failed.", "err", err)
		return false
	}
	return true
}

func (s *SignMgr) CheckRequestLegal(timeStamp int64, macAddr string, macSign string) bool {
	//make sure request is in 30s
	if time.Now().Unix() > timeStamp+30 {
		cli.LocalLogger.Debugln("request past due")
		return false
	}

	//if statemgr.State.MacAddr != macAddr {
	//	logger.Error("request mac address error")
	//	return false
	//}
	//
	//pass := ValidateSignature(statemgr.State.MacAddr, macSign)
	//if pass == false {
	//	logger.Error("ValidateSignature MacAddr fail")
	//	return false
	//}

	return true
}
