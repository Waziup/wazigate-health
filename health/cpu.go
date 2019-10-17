package health

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

var prevStat *linuxproc.Stat

func init() {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal(err)
	}
	prevStat = stat
}

// CPU returns the CPU usage (since the last call to CPU) as value 0~1.
func CPU() float32 {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatalf("Err Can not read CPU: %v", err)
	}
	u := usage(stat.CPUStatAll, prevStat.CPUStatAll)
	prevStat = stat
	return u
}

// CPUTemp return sthe CPU temperature as °C value.
func CPUTemp() float32 {

	data, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		log.Fatalf("Err Can not read CPU temperature: %v", err)
	}
	str := strings.TrimSpace(string(data))
	temp, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatalf("Err Can not parse CPU temperature: %v", err)
	}
	return float32(temp / 1000)
}

func usage(curr, prev linuxproc.CPUStat) float32 {

	prevIdle := prev.Idle + prev.IOWait
	idle := curr.Idle + curr.IOWait

	prevNonIdle := prev.User + prev.Nice + prev.System + prev.IRQ + prev.SoftIRQ + prev.Steal
	nonIdle := curr.User + curr.Nice + curr.System + curr.IRQ + curr.SoftIRQ + curr.Steal

	prevTotal := prevIdle + prevNonIdle
	total := idle + nonIdle

	usage := (float32(total-prevTotal) - float32(idle-prevIdle)) / float32(total-prevTotal)

	return usage
}
