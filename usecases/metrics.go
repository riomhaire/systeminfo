package usecases

import (
	"github.com/riomhaire/systeminfo/entities"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

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

	return info
}
