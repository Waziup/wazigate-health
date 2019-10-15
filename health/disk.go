package health

import (
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

// DiskStat holds information about the disk size and usage.
type DiskStat struct {
	// Number of bytes used.
	Used uint64 `json:"used"`
	// Total size of the disk.
	All uint64 `json:"all"`
}

// Disk returns disk usage imformation.
func Disk() DiskStat {
	disk, err := linuxproc.ReadDisk("/")
	if err != nil {
		log.Fatalf("Err Can not read disk: %v", err)
	}
	return DiskStat{disk.Used, disk.All}
}
