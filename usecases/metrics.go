package usecases

import (
	"github.com/riomhaire/systeminfo/entities"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// Metrics - export Data structure form of the metrics
func Metrics() entities.SystemInformation {
	info := entities.SystemInformation{}

	vmem, err := mem.VirtualMemory()
	if err == nil {
		info.Memory = *vmem
	}

	hostdata, cerr := host.Info()
	if cerr == nil {
		info.Host = *hostdata
	}

	loaddata, lerr := load.Avg()
	if lerr == nil {
		info.Load = *loaddata
	}

	info.CPU.CPUCount, _ = cpu.Counts(false)
	cpuLoad, cperr := cpu.Percent(0, false)
	if cperr == nil {
		info.CPU.CPULoad = cpuLoad[0]
	}

	// OK look up physical storage/partition info
	disks, derr := disk.Partitions(false)
	if derr == nil {
		// OK we have something - make space
		info.Storage = make([]disk.UsageStat, len(disks))
		// Range through the partitions and retrieve the data
		for idx, partition := range disks {
			diskInfo, err := disk.Usage(partition.Mountpoint)
			if err == nil {
				info.Storage[idx] = *diskInfo
			}

		}
	}

	// OK lookup network stats
	info.Network, _ = net.IOCounters(false)

	return info
}
