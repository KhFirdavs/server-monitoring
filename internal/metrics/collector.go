package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func CPUPerc() (float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}
func RAMUsage() (uint64, uint64, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}
	return v.Used, v.Total, nil
}

func DiskUsage() (uint64, uint64, error) {
	v, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return v.Used, v.Total, nil
}
func NetUsage() (uint64, uint64, error) {
	v, err := net.IOCounters(false)
	if err != nil {
		return 0, 0, err
	}
	return v[0].BytesSent, v[0].BytesRecv, nil
}
