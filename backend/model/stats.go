package model

type Stats struct {
	Hostname    string  `json:"hostname"`
	CpuUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
}
