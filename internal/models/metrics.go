package models

type Metrics struct {
	CPUUsage  float64 `json:"cpu_usage"`
	RAMUsed   uint64  `json:"ram_used_mb"`
	RAMTotal  uint64  `json:"ram_total_mb"`
	DiskUsed  uint64  `json:"disk_used_gb"`
	DiskTotal uint64  `json:"disk_total_gb"`
	NetSent   uint64  `json:"net_sent_kb"`
	NetRecv   uint64  `json:"net_recv_kb"`
}
