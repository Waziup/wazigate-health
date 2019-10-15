package health

import "time"

// Boottime is the time this device has booted.
func Boottime() time.Time {
	return prevStat.BootTime
}
