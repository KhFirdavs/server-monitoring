package main

import (
	"log"

	"github.com/KhFirdavs/server-monitoring-go/internal/api"
	"github.com/KhFirdavs/server-monitoring-go/internal/database"
	"github.com/KhFirdavs/server-monitoring-go/internal/handler"
	"github.com/KhFirdavs/server-monitoring-go/internal/metrics"
	"github.com/KhFirdavs/server-monitoring-go/internal/models"
)

func main() {
	metrics.StartCollector()
	log.Println("Сбор метрик запущен...")

	/*Collect, err := metrics.CollectMetrics()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("CPU Usage: %.2f%%\n", Collect.CPUUsage)
	fmt.Printf("RAM Usage: %d MB\n", Collect.RAMUsed)
	fmt.Printf("RAM Total: %d MB\n", Collect.RAMTotal)
	fmt.Printf("Disk Used: %d GB\n", Collect.DiskUsed)
	fmt.Printf("Disk Total: %d GB\n", Collect.DiskTotal)
	fmt.Printf("Net Sent: %d KB\n", Collect.NetSent)
	fmt.Printf("Net Received: %d KB\n", Collect.NetRecv)

	/*cpuUsage, err := metrics.CPUPerc()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
	//
	RAMUsage, RAMTotal, err := metrics.RAMUsage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	RAMUsageMB := RAMUsage / 1024 / 1024
	RAMTotalMB := RAMTotal / 1024 / 1024
	fmt.Println("RAM Usage:", RAMUsageMB, "MB", "Total:", RAMTotalMB, "MB")
	//
	DiskUsed, DiskTotal, err := metrics.DiskUsage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	DiskUsedGb := DiskUsed / 1024 / 1024 / 1024
	DiskTotalGb := DiskTotal / 1024 / 1024 / 1024
	fmt.Println("Disk Used:", DiskUsedGb, "Gb", "Total:", DiskTotalGb, "Gb")
	//
	NetSent, NetRecv, err := metrics.NetUsage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	NetSentKb := NetSent / 1024
	NetRecvKb := NetRecv / 1024
	fmt.Println("Net Sent:", NetSentKb, "Kb", "Received:", NetRecvKb, "Kb")*/

	router := handler.NewRouter()

	srv := &api.Server{}
	if err := srv.ServerRun(router, "8080"); err != nil {
		log.Fatalf("Failed to run server: %s", err.Error())
	}

	db := database.NewConnectPostgres()
	if err := db.AutoMigrate(&models.Metrics{}); err != nil {
		log.Fatalf("Failed to migrate database: %s", err.Error())
	}
}
