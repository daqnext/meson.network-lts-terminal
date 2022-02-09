package destinationMgr

import (
	"strconv"
	"strings"
	"time"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/imroc/req"
	"github.com/universe-30/UUtils/hash_util"
)

type DestMgr struct {
	CurrentDest string
	backupDest  map[string]struct{}

	isChecking bool
}

var echoServer *DestMgr

func Init() {
	echoServer = &DestMgr{
		backupDest: map[string]struct{}{},
	}

}

func GetSingleInstance() *DestMgr {
	return echoServer
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
func (d *DestMgr) GenBackupDest() {
	for i := 10; i <= 100; i = i + 10 {
		k := hash_util.MD5Hash([]byte(strconv.Itoa(i)))
		k = k[3:18]
		k = reverseString(k)
		k = "https://" + k + ".com"

		d.backupDest[k] = struct{}{}
	}
	if d.CurrentDest != "" {
		d.backupDest[d.CurrentDest] = struct{}{}
	}

}

func checkDestAvailable(targetUrl string) bool {
	r := req.New()
	r.SetTimeout(time.Duration(8) * time.Second)
	response, err := r.Get(targetUrl)
	if err != nil {
		return false
	}
	responseData := response.Response()
	responseStatusCode := responseData.StatusCode
	if responseStatusCode == 200 {
		//todo check back data is from meson

		return true
	}

	return false
}

func (d *DestMgr) SearchAvailableDest() {
	if d.isChecking {
		return
	}
	d.isChecking = true
	defer func() {
		d.isChecking = false
	}()

	for i := 0; i < 2; i++ {
		for k := range d.backupDest {
			url := k + "/api/health"
			checkResult := checkDestAvailable(url)
			if checkResult {
				d.CurrentDest = k
				return
			} else {
				time.Sleep(1 * time.Second)
			}
		}
		time.Sleep(2 * time.Second)
	}
	basic.Logger.Errorln("Network error. Please check network environment or download new version Terminal in https://meson.network")
}

func (d *DestMgr) GetDestHost() string {
	return d.CurrentDest
}

func (d *DestMgr) GetDestUrl(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return d.CurrentDest + path
}
