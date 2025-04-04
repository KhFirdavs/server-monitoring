package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
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

//func DiskUsage() (float64, error) {
