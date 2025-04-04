package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func CPUPerc() (float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
}
