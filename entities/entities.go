package entities

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type CPUInformation struct {
	CPUCount int     `json:"cpus"`
	CPULoad  float64 `json:"cpuLoad"`
}

type SystemInformation struct {
	Memory  mem.VirtualMemoryStat `json:"memory"`
	Host    host.InfoStat         `json:"host"`
	Load    load.AvgStat          `json:"load"`
	CPU     CPUInformation        `json:"cpu"`
	Storage []disk.UsageStat      `json:"storage"`
	Network []net.IOCountersStat  `json:"network"`
}
