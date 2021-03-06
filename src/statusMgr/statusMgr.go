package statusMgr

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"time"

	meson_msg "github.com/daqnext/meson-msg"
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/src/diskMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/echoServer"
	"github.com/daqnext/meson.network-lts-terminal/src/versionMgr"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type StatusMgr struct {
	Status *meson_msg.TerminalStatesMsg
}

var statusMgr *StatusMgr

func Init() {
	if statusMgr != nil {
		return
	}
	statusMgr = &StatusMgr{}
}

func GetSingleInstance() *StatusMgr {
	return statusMgr
}

func RunCommand(cmdstring string, args ...string) (string, error) {
	cmd := exec.Command(cmdstring, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), err
}

//func getMainMacAddress() (string, error) {
//	ifas, err := sysnet.Interfaces()
//	if err != nil {
//		return "", err
//	}
//
//	ans := ""
//	ansIndex := 1024
//
//	for _, ifa := range ifas {
//		// fmt.Printf("%+v %+v\n", ifa, uint(ifa.Flags))
//
//		// Flags(19) means `up|broadcast|multicast`
//		if ifa.Flags == sysnet.Flags(19) && ifa.Index < ansIndex {
//			ans = ifa.HardwareAddr.String()
//			ansIndex = ifa.Index
//		}
//	}
//	return ans, nil
//}

//func calAverageNetSpeed() {
//	go func() {
//		for true {
//			if n, err := net.IOCounters(false); err == nil && len(n) > 0 {
//				for i, _ := range n {
//					if i >= 3 {
//						break
//					}
//					sent := n[i].BytesSent
//					recv := n[i].BytesRecv
//					if netBytesRecv[i] != 0 && netBytesSent[i] != 0 {
//						//State.NetInRate = (recv - netBytesRecv) / uint64(s.config.statsReportPeriod.Milliseconds()/1000)
//						//State.NetOutRate = (sent - netBytesSent) / uint64(s.config.statsReportPeriod.Milliseconds()/1000)
//						NetInRate := (recv - netBytesRecv[i]) / uint64(5)
//						NetOutRate := (sent - netBytesSent[i]) / uint64(5)
//						State.NetInMbs[i] = float64(NetInRate*8) / float64(1e6)
//						State.NetOutMbs[i] = float64(NetOutRate*8) / float64(1e6)
//						//fmt.Println(State.NetInMbs,"Mbs")
//						//fmt.Println(State.NetOutMbs,"Mbs")
//					}
//					netBytesRecv[i] = recv
//					netBytesSent[i] = sent
//				}
//			}
//			time.Sleep(time.Second * 5)
//		}
//	}()
//}

func getMachineSetupTime() string {
	switch runtime.GOOS {
	case "linux":
		result, err := RunCommand("/bin/bash", "-c", "ls -lact --full-time /etc | tail -1 |awk '{print $6,$7}'")
		if err != nil {
			basic.Logger.Debugln("aws ec2 run command err", "err", err)
			return "linux unknown"
		}
		basic.Logger.Debugln("machine setup time", "time", result)
		return result
	case "windows":
		return "windows unknown"
	case "darwin":
		return "darwin unknown"
	}
	return "unknown"
}

func (s *StatusMgr) GetMachineStatus() {
	if s.Status == nil {
		s.Status = &meson_msg.TerminalStatesMsg{}
	}

	//unchanged
	if s.Status.OS == "" {
		if h, err := host.Info(); err == nil {
			s.Status.OS = fmt.Sprintf("%v:%v(%v):%v", h.OS, h.Platform, h.PlatformFamily, h.PlatformVersion)
		}
	}

	if s.Status.CPU == "" {
		if c, err := cpu.Info(); err == nil {
			s.Status.CPU = c[0].ModelName
		}
	}

	if s.Status.MachineSetupTime == "" {
		s.Status.MachineSetupTime = getMachineSetupTime()
	}

	if s.Status.Port == 0 {
		s.Status.Port = echoServer.GetSingleInstance().Http_port
	}

	//need update data
	//memory
	if v, err := mem.VirtualMemory(); err == nil {
		s.Status.MemTotal = int64(v.Total)
		s.Status.MemAvailable = int64(v.Available)
	}

	//cpu usage
	if percent, err := cpu.Percent(time.Second, false); err != nil || len(percent) <= 0 {
		basic.Logger.Debugln("failed to get cup usage", "err", err)
	} else {
		s.Status.CpuUsage = percent[0]
	}

	s.Status.Version = versionMgr.GetSingleInstance().CurrentVersion

	//disk
	total, used, _ := diskMgr.GetSingleInstance().GetSpaceInfo()
	s.Status.CdnDiskTotal = total
	s.Status.CdnDiskUsed = used

}
