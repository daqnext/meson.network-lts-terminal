package tlsCertificateMgr

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"time"

	"github.com/daqnext/meson-msg"
	"github.com/daqnext/meson.network-lts-terminal/src/destMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/globalData"
	"github.com/daqnext/meson.network-lts-terminal/src/requestUtil"
)

type CertMgr struct {
	ChainFileUrl    string
	KeyFileUrl      string
	ChainByte       []byte
	KeyByte         []byte
	CurrentCertHash uint32
}

var certMgr *CertMgr

func Init() {
	if certMgr != nil {
		return
	}
	certMgr = &CertMgr{}
}

func GetSingleInstance() *CertMgr {
	return certMgr
}

func (c *CertMgr) CheckCertHash() (needUpdate bool, err error) {
	if c.CurrentCertHash == 0 {
		return true, nil
	}

	url := destMgr.GetSingleInstance().GetDestUrl("/api/cert/terminaldomain/checkhash")
	resp, err := requestUtil.Get(url, nil, 30, globalData.Token)
	if err != nil {
		return false, err
	}
	if resp.Response().StatusCode != 200 {
		return false, errors.New("CheckCertHash response status error:" + resp.Response().Status)
	}

	var certMsg meson_msg.CertMsg
	err = resp.ToJSON(&certMsg)
	if err != nil {
		return false, err
	}

	if certMsg.Hash != 0 && certMsg.Hash == c.CurrentCertHash {
		return true, nil
	}
	return false, nil
}

func validateCert(chain []byte, key []byte) error {
	//load tls file
	cert, err := tls.X509KeyPair(chain, key)
	if err != nil {
		return err
	}

	//parse
	c, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return err
	}

	//check pastDue
	if time.Now().Unix() > c.NotAfter.Unix() {
		return errors.New("cert past due")
	}

	return nil
}

func (c *CertMgr) GetAndParseCert() error {
	//get chain from server
	url := destMgr.GetSingleInstance().GetDestUrl("/api/cert/terminaldomain/cert")
	resp, err := requestUtil.Get(url, nil, 30, globalData.Token)
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != 200 {
		return errors.New("GetAndParseCert response status error:" + resp.Response().Status)
	}

	var certMsg meson_msg.CertMsg
	err = resp.ToJSON(&certMsg)
	if err != nil {
		return err
	}

	err = validateCert(certMsg.Chain, certMsg.Key)
	if err != nil {
		return err
	}

	c.ChainByte = certMsg.Chain
	c.KeyByte = certMsg.Key
	c.CurrentCertHash = certMsg.Hash
	return nil
}

func (c *CertMgr) GetTlsCert() (chain []byte, key []byte) {
	return c.getTlsChain(), c.getTlsKey()
}

func (c *CertMgr) getTlsChain() []byte {
	//for test https://local.shoppynext.com
	return []byte(`-----BEGIN CERTIFICATE-----
MIIGOzCCBSOgAwIBAgIRAM0948NAA9Jn+KT5yZlbTYkwDQYJKoZIhvcNAQELBQAw
gY8xCzAJBgNVBAYTAkdCMRswGQYDVQQIExJHcmVhdGVyIE1hbmNoZXN0ZXIxEDAO
BgNVBAcTB1NhbGZvcmQxGDAWBgNVBAoTD1NlY3RpZ28gTGltaXRlZDE3MDUGA1UE
AxMuU2VjdGlnbyBSU0EgRG9tYWluIFZhbGlkYXRpb24gU2VjdXJlIFNlcnZlciBD
QTAeFw0yMTEyMTQwMDAwMDBaFw0yMjEyMTQyMzU5NTlaMBsxGTAXBgNVBAMMECou
c2hvcHB5bmV4dC5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDy
G2E6sQ1rCu8iUz8wnTSZH8wUChHpzNeiW5e2ssMjCW28aJRYw37WY8wLQpCupk5i
5CbeAWfB5N7FUAnmIlhH35zkFxpsikne8RGyMWMpC679tbdL5/dAoOlWihuwLbl/
VvwZapMlp+tvfVO9tfBGeiYKWbYJ1BbKo12PGe7m/SE48x/pURLhfVxMIBn1WBJR
mEzTQCetpmIlVwca6S5F2pUqYejUcZedzoVAX7+gwyx62UcDBPW9S9oz/FoGL4B4
GvJrvfu9Hj+kXoH95j77Ig1i1URl03bQO15UJDzov2fW+E/SQWSS7ejfECEnjOkQ
U5J6dWdt2w5k2U2C+xqDAgMBAAGjggMDMIIC/zAfBgNVHSMEGDAWgBSNjF7EVK2K
4Xfpm/mbBeG4AY1h4TAdBgNVHQ4EFgQUCRFrSaR6c7d9j6oRiOMDaEeHMPwwDgYD
VR0PAQH/BAQDAgWgMAwGA1UdEwEB/wQCMAAwHQYDVR0lBBYwFAYIKwYBBQUHAwEG
CCsGAQUFBwMCMEkGA1UdIARCMEAwNAYLKwYBBAGyMQECAgcwJTAjBggrBgEFBQcC
ARYXaHR0cHM6Ly9zZWN0aWdvLmNvbS9DUFMwCAYGZ4EMAQIBMIGEBggrBgEFBQcB
AQR4MHYwTwYIKwYBBQUHMAKGQ2h0dHA6Ly9jcnQuc2VjdGlnby5jb20vU2VjdGln
b1JTQURvbWFpblZhbGlkYXRpb25TZWN1cmVTZXJ2ZXJDQS5jcnQwIwYIKwYBBQUH
MAGGF2h0dHA6Ly9vY3NwLnNlY3RpZ28uY29tMCsGA1UdEQQkMCKCECouc2hvcHB5
bmV4dC5jb22CDnNob3BweW5leHQuY29tMIIBfwYKKwYBBAHWeQIEAgSCAW8EggFr
AWkAdgBGpVXrdfqRIDC1oolp9PN9ESxBdL79SbiFq/L8cP5tRwAAAX22wsEwAAAE
AwBHMEUCIAkRBAM//DxB9wZUo+OxrbEwBiZCLcIsSeqDSTXRNtGmAiEA2xQBTmfK
4X0dliwCVCMZJneHZyaPvoZWhdFp9UgK1eMAdgBByMqx3yJGShDGoToJQodeTjGL
GwPr60vHaPCQYpYG9gAAAX22wsFDAAAEAwBHMEUCIHAKIbSZYjzaPyuXfFtp5GIY
dxi6SlhxqsmM9RFfbMKlAiEA+WdCMAg+q0Swu816Vi26bo1ze+vlJlR2wygG+SzT
oF8AdwApeb7wnjk5IfBWc59jpXflvld9nGAK+PlNXSZcJV3HhAAAAX22wsESAAAE
AwBIMEYCIQD3i19I5eLkgOIOsWj+yK8oliio5A0pT+2eVX9KAZqnRwIhAPtki1QS
UyZpcYknACWJv2GePkcuYkugH3bke5aiLS/zMA0GCSqGSIb3DQEBCwUAA4IBAQBg
IbQCiaNShj66X/mCaWP/6jR+MauYItfdd+sp9PmKe9KZmsFQHhhkvmA6CpilU9+q
F7VgCkzWMyd9Ay4tvWxzlxhf7S5PKR/1BYdi1oGtG3jWteGXypO2MIVzPd7IqS7g
duLsWuddkI33Urd7pLmQhRPlPpdh0hVwgmnL89PJiyYwJ9CfHLg+eFBEjth8YR/y
IML+r3ev9GxggwDyDiVz/Gc8Md12BkbKinj8FU+fj6QQVzi8DK1ivaFkUN9qPQOT
6iBiIBoZUDWgVX6BHMOHtS7lVf+4qxFAABxQhFApgqUvnU1ELSmObYZClRw9t9lz
VpuJiijQX6QhF7+HgpJN
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIGEzCCA/ugAwIBAgIQfVtRJrR2uhHbdBYLvFMNpzANBgkqhkiG9w0BAQwFADCB
iDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCk5ldyBKZXJzZXkxFDASBgNVBAcTC0pl
cnNleSBDaXR5MR4wHAYDVQQKExVUaGUgVVNFUlRSVVNUIE5ldHdvcmsxLjAsBgNV
BAMTJVVTRVJUcnVzdCBSU0EgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMTgx
MTAyMDAwMDAwWhcNMzAxMjMxMjM1OTU5WjCBjzELMAkGA1UEBhMCR0IxGzAZBgNV
BAgTEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4GA1UEBxMHU2FsZm9yZDEYMBYGA1UE
ChMPU2VjdGlnbyBMaW1pdGVkMTcwNQYDVQQDEy5TZWN0aWdvIFJTQSBEb21haW4g
VmFsaWRhdGlvbiBTZWN1cmUgU2VydmVyIENBMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEA1nMz1tc8INAA0hdFuNY+B6I/x0HuMjDJsGz99J/LEpgPLT+N
TQEMgg8Xf2Iu6bhIefsWg06t1zIlk7cHv7lQP6lMw0Aq6Tn/2YHKHxYyQdqAJrkj
eocgHuP/IJo8lURvh3UGkEC0MpMWCRAIIz7S3YcPb11RFGoKacVPAXJpz9OTTG0E
oKMbgn6xmrntxZ7FN3ifmgg0+1YuWMQJDgZkW7w33PGfKGioVrCSo1yfu4iYCBsk
Haswha6vsC6eep3BwEIc4gLw6uBK0u+QDrTBQBbwb4VCSmT3pDCg/r8uoydajotY
uK3DGReEY+1vVv2Dy2A0xHS+5p3b4eTlygxfFQIDAQABo4IBbjCCAWowHwYDVR0j
BBgwFoAUU3m/WqorSs9UgOHYm8Cd8rIDZsswHQYDVR0OBBYEFI2MXsRUrYrhd+mb
+ZsF4bgBjWHhMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8ECDAGAQH/AgEAMB0G
A1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAbBgNVHSAEFDASMAYGBFUdIAAw
CAYGZ4EMAQIBMFAGA1UdHwRJMEcwRaBDoEGGP2h0dHA6Ly9jcmwudXNlcnRydXN0
LmNvbS9VU0VSVHJ1c3RSU0FDZXJ0aWZpY2F0aW9uQXV0aG9yaXR5LmNybDB2Bggr
BgEFBQcBAQRqMGgwPwYIKwYBBQUHMAKGM2h0dHA6Ly9jcnQudXNlcnRydXN0LmNv
bS9VU0VSVHJ1c3RSU0FBZGRUcnVzdENBLmNydDAlBggrBgEFBQcwAYYZaHR0cDov
L29jc3AudXNlcnRydXN0LmNvbTANBgkqhkiG9w0BAQwFAAOCAgEAMr9hvQ5Iw0/H
ukdN+Jx4GQHcEx2Ab/zDcLRSmjEzmldS+zGea6TvVKqJjUAXaPgREHzSyrHxVYbH
7rM2kYb2OVG/Rr8PoLq0935JxCo2F57kaDl6r5ROVm+yezu/Coa9zcV3HAO4OLGi
H19+24rcRki2aArPsrW04jTkZ6k4Zgle0rj8nSg6F0AnwnJOKf0hPHzPE/uWLMUx
RP0T7dWbqWlod3zu4f+k+TY4CFM5ooQ0nBnzvg6s1SQ36yOoeNDT5++SR2RiOSLv
xvcRviKFxmZEJCaOEDKNyJOuB56DPi/Z+fVGjmO+wea03KbNIaiGCpXZLoUmGv38
sbZXQm2V0TP2ORQGgkE49Y9Y3IBbpNV9lXj9p5v//cWoaasm56ekBYdbqbe4oyAL
l6lFhd2zi+WJN44pDfwGF/Y4QA5C5BIG+3vzxhFoYt/jmPQT2BVPi7Fp2RBgvGQq
6jG35LWjOhSbJuMLe/0CjraZwTiXWTb2qHSihrZe68Zk6s+go/lunrotEbaGmAhY
LcmsJWTyXnW0OMGuf1pGg+pRyrbxmRE1a6Vqe8YAsOf4vmSyrcjC8azjUeqkk+B5
yOGBQMkKW+ESPMFgKuOXwIlCypTPRpgSabuY0MLTDXJLR27lk8QyKGOHQ+SwMj4K
00u/I5sUKUErmgQfky3xxzlIPK1aEn8=
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFgTCCBGmgAwIBAgIQOXJEOvkit1HX02wQ3TE1lTANBgkqhkiG9w0BAQwFADB7
MQswCQYDVQQGEwJHQjEbMBkGA1UECAwSR3JlYXRlciBNYW5jaGVzdGVyMRAwDgYD
VQQHDAdTYWxmb3JkMRowGAYDVQQKDBFDb21vZG8gQ0EgTGltaXRlZDEhMB8GA1UE
AwwYQUFBIENlcnRpZmljYXRlIFNlcnZpY2VzMB4XDTE5MDMxMjAwMDAwMFoXDTI4
MTIzMTIzNTk1OVowgYgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpOZXcgSmVyc2V5
MRQwEgYDVQQHEwtKZXJzZXkgQ2l0eTEeMBwGA1UEChMVVGhlIFVTRVJUUlVTVCBO
ZXR3b3JrMS4wLAYDVQQDEyVVU0VSVHJ1c3QgUlNBIENlcnRpZmljYXRpb24gQXV0
aG9yaXR5MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAgBJlFzYOw9sI
s9CsVw127c0n00ytUINh4qogTQktZAnczomfzD2p7PbPwdzx07HWezcoEStH2jnG
vDoZtF+mvX2do2NCtnbyqTsrkfjib9DsFiCQCT7i6HTJGLSR1GJk23+jBvGIGGqQ
Ijy8/hPwhxR79uQfjtTkUcYRZ0YIUcuGFFQ/vDP+fmyc/xadGL1RjjWmp2bIcmfb
IWax1Jt4A8BQOujM8Ny8nkz+rwWWNR9XWrf/zvk9tyy29lTdyOcSOk2uTIq3XJq0
tyA9yn8iNK5+O2hmAUTnAU5GU5szYPeUvlM3kHND8zLDU+/bqv50TmnHa4xgk97E
xwzf4TKuzJM7UXiVZ4vuPVb+DNBpDxsP8yUmazNt925H+nND5X4OpWaxKXwyhGNV
icQNwZNUMBkTrNN9N6frXTpsNVzbQdcS2qlJC9/YgIoJk2KOtWbPJYjNhLixP6Q5
D9kCnusSTJV882sFqV4Wg8y4Z+LoE53MW4LTTLPtW//e5XOsIzstAL81VXQJSdhJ
WBp/kjbmUZIO8yZ9HE0XvMnsQybQv0FfQKlERPSZ51eHnlAfV1SoPv10Yy+xUGUJ
5lhCLkMaTLTwJUdZ+gQek9QmRkpQgbLevni3/GcV4clXhB4PY9bpYrrWX1Uu6lzG
KAgEJTm4Diup8kyXHAc/DVL17e8vgg8CAwEAAaOB8jCB7zAfBgNVHSMEGDAWgBSg
EQojPpbxB+zirynvgqV/0DCktDAdBgNVHQ4EFgQUU3m/WqorSs9UgOHYm8Cd8rID
ZsswDgYDVR0PAQH/BAQDAgGGMA8GA1UdEwEB/wQFMAMBAf8wEQYDVR0gBAowCDAG
BgRVHSAAMEMGA1UdHwQ8MDowOKA2oDSGMmh0dHA6Ly9jcmwuY29tb2RvY2EuY29t
L0FBQUNlcnRpZmljYXRlU2VydmljZXMuY3JsMDQGCCsGAQUFBwEBBCgwJjAkBggr
BgEFBQcwAYYYaHR0cDovL29jc3AuY29tb2RvY2EuY29tMA0GCSqGSIb3DQEBDAUA
A4IBAQAYh1HcdCE9nIrgJ7cz0C7M7PDmy14R3iJvm3WOnnL+5Nb+qh+cli3vA0p+
rvSNb3I8QzvAP+u431yqqcau8vzY7qN7Q/aGNnwU4M309z/+3ri0ivCRlv79Q2R+
/czSAaF9ffgZGclCKxO/WIu6pKJmBHaIkU4MiRTOok3JMrO66BQavHHxW/BBC5gA
CiIDEOUMsfnNkjcZ7Tvx5Dq2+UUTJnWvu6rvP3t3O9LEApE9GQDTF1w52z97GA1F
zZOFli9d31kWTz9RvdVFGD/tSo7oBmF0Ixa1DVBzJ0RHfxBdiSprhTEUxOipakyA
vGp4z7h/jnZymQyd/teRCBaho1+V
-----END CERTIFICATE-----
`)
}

func (c *CertMgr) getTlsKey() []byte {
	//for test https://local.shoppynext.com
	return []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA8hthOrENawrvIlM/MJ00mR/MFAoR6czXoluXtrLDIwltvGiU
WMN+1mPMC0KQrqZOYuQm3gFnweTexVAJ5iJYR9+c5BcabIpJ3vERsjFjKQuu/bW3
S+f3QKDpVoobsC25f1b8GWqTJafrb31TvbXwRnomClm2CdQWyqNdjxnu5v0hOPMf
6VES4X1cTCAZ9VgSUZhM00AnraZiJVcHGukuRdqVKmHo1HGXnc6FQF+/oMMsetlH
AwT1vUvaM/xaBi+AeBrya737vR4/pF6B/eY++yINYtVEZdN20DteVCQ86L9n1vhP
0kFkku3o3xAhJ4zpEFOSenVnbdsOZNlNgvsagwIDAQABAoIBAHTDPa4nq5f/avs7
+NLCTpa23h6gCKmgcDLOR6oGJtRj/LeMHFd+2VIeSU746HxCmrVY4WHafabmcYXp
pwyFbdwj/S+H8Od9/kh2LRmqde9awid8Fw2VrwpNUJu6+cWLF9ZRJj4/xv23MXXQ
lK1+7IWg7W26fG/bCK7sAPg+ApkFdXcwqw935sR0xLw1NZFH4Io/4X5dbRNEpC6u
fAkRvl2CJeV0ViJQBpwdA0YjMXmgjXBZYtZbjJL9gbavKrcVEwjX0mYdvXeg3Ax8
MEZh39RNjxkLtSwmgCLlex7zx7cPAVV4prHUe7dTOpV08lq4I/izACTZRXg5LjQU
goJE5GkCgYEA90phQTreaTPbWcQWNXH9gYlyPSom0TyYkXOaDAS6wuntEQCdVG+6
g4mHZrrsDnpCl0iIB0f9ZjzANS1zMIvipOGd7qYqFO3noo67u6622Io5fhAitBWK
M7qms461BXVsph1jOgoOFVG9rcbtndchNY20Ydex6C3PrwPql4BN+IcCgYEA+qJD
eyuTstogg5SZ859sRb9PCdnqRrLfYJY+Qj5+gvNodoTN1zeG4x/36U56HjYTzFrQ
74P3WJEwcm00ph5VGiH0TUJAdMQQ2TAaXYODvxBkWVO6TAUgL/M75Y/RFdCO/US1
BRwDJ1/YaLR3bejyKb1RpZwNzqEL6WpzoGtHGSUCgYAldCiCvq3M9UO1ttM+SQOC
SRT6WpYCftEExqOASn8W4mM7fgJWNY3kOkI9tvXlw3KugxfHMooYn/7kjvuxUJ6F
Jn7LFHOvM9Evd7rLVEzxQw4uH7eB2vlRmGWDMIwORZitGCpdMgSsCfNWjJiUnW60
M8AsSYTyi223Ljqrs40bpQKBgCdNbLGS6s1gITshAWdPt6XUUyujTXaatCasSMUQ
kbwtOVNkjfbS0UcqizC9yq6UIlSoZR06H352/hbjgx8NoKDBdFLtMbhdypqUTX4e
knlSs7nCRHOJVjvOs7TS8aGvG80hihVsCB6TjBcXPacxoU/kTTpgF1YwsPKAa/Na
/0I1AoGADVUByMLPdjfLf/waz81Qub8G39GkRxZX/saSZQBZOwFQ8Njq3dF/qAbp
KF7wauwAVpbN144XD9Ne6LF2HeARAoU40xTmhyQS9a1DBmEKnCEaPZgh1EhrVwVd
05rEflq8Cg7CqPj0akHnnDIg0X9Krg2QiRIhX2uv3JtKjEK/JLc=
-----END RSA PRIVATE KEY-----
`)
}
