package status

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
	"testing"
	"time"
)

func Test_cpu(t *testing.T) {
	percent, err := cpu.Percent(5*time.Second, false)
	if err != nil || len(percent) <= 0 {
		log.Println("failed to get cup usage", "err", err)
	}
}
