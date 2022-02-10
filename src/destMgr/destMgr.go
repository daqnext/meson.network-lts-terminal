package destMgr

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/daqnext/meson.network-lts-terminal/src/requestUtil"
	"github.com/universe-30/UUtils/hash_util"
)

type DestMgr struct {
	CurrentDest string
	backupDest  map[string]struct{}

	isChecking bool
}

var destMgr *DestMgr

func Init() {
	destMgr = &DestMgr{
		backupDest: map[string]struct{}{
			"coldcdn.com": struct{}{}, //todo running server host
		},
	}
	dest, _ := configuration.Config.GetString("dest", "")
	destMgr.CurrentDest = dest
	destMgr.genBackupDest()
}

func GetSingleInstance() *DestMgr {
	return destMgr
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
func (d *DestMgr) genBackupDest() {
	for i := 10; i <= 100; i = i + 10 {
		k := hash_util.MD5Hash([]byte(strconv.Itoa(i)))
		k = k[3:18]
		k = reverseString(k)
		k = k + ".com"

		d.backupDest[k] = struct{}{}
	}
	if d.CurrentDest != "" {
		d.backupDest[d.CurrentDest] = struct{}{}
	}

}

func checkDestAvailable(targetUrl string) bool {
	basic.Logger.Debugln("checking url:", targetUrl)
	response, err := requestUtil.Get(targetUrl, nil, 8, "")
	if err != nil {
		basic.Logger.Debugln("checkDestAvailable err:", err)
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

func (d *DestMgr) SearchAvailableDest() error {
	if d.isChecking {
		return nil
	}
	d.isChecking = true
	defer func() {
		d.isChecking = false
	}()

	if d.CurrentDest != "" {
		url := "https://" + d.CurrentDest + "/api/health"
		checkResult := checkDestAvailable(url)
		if checkResult {
			return nil
		}
	}

	for i := 0; i < 2; i++ {
		for k := range d.backupDest {
			url := "https://" + k + "/api/health"
			checkResult := checkDestAvailable(url)
			if checkResult {
				d.CurrentDest = k
				return nil
			} else {
				time.Sleep(1 * time.Second)
			}
		}
		time.Sleep(2 * time.Second)
	}
	basic.Logger.Errorln("Network error. Please check network environment or download new version Terminal in https://meson.network")
	return errors.New("no available dest")
}

func (d *DestMgr) GetDestHost() string {
	return d.CurrentDest
}

func (d *DestMgr) GetDestUrl(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return "https://" + d.CurrentDest + path
}
