package health

import (
	"bytes"
	"io/ioutil"
	"log"
)

// TelemetryStat holds various device data.
type TelemetryStat struct {
	OSVersion   string `json:"osVersion"`
	OSVersionID string `json:"osVersionID"`
	OSName      string `json:"osName"`
	OSID        string `json:"osID"`
}

// Telemetry returns some general information about this device.
func Telemetry() TelemetryStat {
	buf, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		log.Fatalf("Err Can not read telemetry: %v", err)
	}
	var stat TelemetryStat
	i := 0
	for i != -1 {
		i = bytes.IndexRune(buf, '\n')
		line := buf
		if i != -1 {
			line = buf[:i]
		}
		j := bytes.IndexRune(line, '=')
		if j != -1 {
			key := line[:j]
			value := line[j+1:]
			if len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
				value = value[1 : len(value)-1]
			}
			switch string(key) {
			case "NAME":
				stat.OSName = string(value)
			case "VERSION":
				stat.OSVersion = string(value)
			case "VERSION_ID":
				stat.OSVersionID = string(value)
			case "ID":
				stat.OSID = string(value)
			}
		}
		buf = buf[i+1:]
	}
	return stat
}
