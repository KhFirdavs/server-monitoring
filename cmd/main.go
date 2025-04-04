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
	RAMUsageMB := RAMUsage / 1024 / 1024
	RAMTotalMB := RAMTotal / 1024 / 1024
	fmt.Println("RAM Usage:", RAMUsageMB, "MB", "Total:", RAMTotalMB, "MB")

	DiskUsed, DiskTotal, err := metrics.DiskUsage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	DiskUsedGb := DiskUsed / 1024 / 1024 / 1024
	DiskTotalGb := DiskTotal / 1024 / 1024 / 1024
	fmt.Println("Disk Used:", DiskUsedGb, "Gb", "Total:", DiskTotalGb, "Gb")

	NetSent, NetRecv, err := metrics.NetUsage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	NetSentKb := NetSent / 1024
	NetRecvKb := NetRecv / 1024
	fmt.Println("Net Sent:", NetSentKb, "Kb", "Received:", NetRecvKb, "Kb")
}
