package health

import (
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

// MemStat holds information about he total memory and the used memory.
type MemStat struct {
	// Number of bytes used.
	Used uint64 `json:"used"`
	// Total memory available..
	All uint64 `json:"all"`
}

// Mem returns disk usage imformation.
func Mem() MemStat {
	mem, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Fatalf("Err Can not read memory: %v", err)
	}
	total := mem.MemTotal
	free := mem.MemFree + mem.Buffers + mem.Cached
	return MemStat{total-free, total}
}
