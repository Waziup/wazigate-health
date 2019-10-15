package health

import (
	"log"

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
