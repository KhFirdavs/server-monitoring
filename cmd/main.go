package main

import (
	"fmt"

	"github.com/KhFirdavs/server-monitoring-go/internal/metrics"
)

func main() {
	cpuUsage, err := metrics.CPUPerc()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)

	RAMUsage, RAMTotal, err := metrics.RAMUsage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	RAMUsage = RAMUsage / 1024 / 1024
	RAMTotal = RAMTotal / 1024 / 1024
	fmt.Println("RAM Usage:", RAMUsage, "MB", "Total:", RAMTotal, "MB")
}
