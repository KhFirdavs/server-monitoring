package metrics

import (
	"log"
	"time"

	"github.com/KhFirdavs/server-monitoring-go/internal/database"
	"github.com/KhFirdavs/server-monitoring-go/internal/models"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type MetricsData struct {
	CPUUsage  float64 `json:"cpu_usage"`
	RAMUsed   uint64  `json:"ram_used_mb"`
	RAMTotal  uint64  `json:"ram_total_mb"`
	DiskUsed  uint64  `json:"disk_used_gb"`
	DiskTotal uint64  `json:"disk_total_gb"`
	NetSent   uint64  `json:"net_sent_kb"`
	NetRecv   uint64  `json:"net_recv_kb"`
}

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
func CollectMetrics() (MetricsData, error) {
	cpuUsage, err := CPUPerc()
	if err != nil {
		return MetricsData{}, err
	}
	ramUsage, ramTotal, err := RAMUsage()
	if err != nil {
		return MetricsData{}, err
	}
	diskUsed, diskTotal, err := DiskUsage()
	if err != nil {
		return MetricsData{}, err
	}
	netSent, netRecv, err := NetUsage()
	if err != nil {
		return MetricsData{}, err
	}
	return MetricsData{
		CPUUsage:  cpuUsage,
		RAMUsed:   ramUsage / 1024 / 1024,
		RAMTotal:  ramTotal / 1024 / 1024,
		DiskUsed:  diskUsed / 1024 / 1024 / 1024,
		DiskTotal: diskTotal / 1024 / 1024 / 1024,
		NetSent:   netSent / 1024,
		NetRecv:   netRecv / 1024,
	}, nil
}

func StartCollector() {
	db := database.NewConnectPostgres()

	go func() {
		for {
			time.Sleep(1 * time.Second)

			data, err := CollectMetrics()
			if err != nil {
				log.Println("Ошибка при сборе метрик:", err)
				continue
			}

			err = database.SaveMetricsToDB(db, &models.Metrics{
				CPUUsage:  data.CPUUsage,
				RAMUsed:   data.RAMUsed,
				RAMTotal:  data.RAMTotal,
				DiskUsed:  data.DiskUsed,
				DiskTotal: data.DiskTotal,
				NetSent:   data.NetSent,
				NetRecv:   data.NetRecv,
			})
			if err != nil {
				log.Println("Ошибка при сохранении метрик в базу:", err)
			} else {
				log.Println("Метрики успешно сохранены в базу")
			}
		}
	}()
}
