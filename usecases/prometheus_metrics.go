package usecases

import (
	"bytes"
	"fmt"
)

func PrometheusMetrics() string {
	metrics := Metrics() // Get the raw metrics
	var buffer bytes.Buffer

	buffer.WriteString("# HELP systeminfo_memory_bytes How much memory.\n")
	buffer.WriteString("# TYPE systeminfo_memory_bytes gauge\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_memory_bytes %v\n", metrics.Memory.Total))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString("# HELP systeminfo_memory_available_bytes How much memory available.\n")
	buffer.WriteString("# TYPE systeminfo_memory_available_bytes gauge\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_memory_available_bytes %v\n", metrics.Memory.Available))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString("# HELP systeminfo_memory_used_percentage How much memory used as percentage.\n")
	buffer.WriteString("# TYPE systeminfo_memory_used_percentage gauge\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_memory_used_percentage %v\n", metrics.Memory.UsedPercent))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString("# HELP systeminfo_host_uptime How long been on\n")
	buffer.WriteString("# TYPE systeminfo_host_uptime counter\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_host_uptime %v\n", metrics.Host.Uptime))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString("# HELP systeminfo_host_procs How Many Processes\n")
	buffer.WriteString("# TYPE systeminfo_host_procs guage\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_procs %v\n", metrics.Host.Procs))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString("# HELP systeminfo_load_1 Load Last Minute\n")
	buffer.WriteString("# TYPE systeminfo_load_1 guage\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_load_1 %v\n", metrics.Load.Load1))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString("# HELP systeminfo_load_5 Load Last 5 Minute\n")
	buffer.WriteString("# TYPE systeminfo_load_5 guage\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_load_5 %v\n", metrics.Load.Load5))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString("# HELP systeminfo_load_15 Load Last 15 Minute\n")
	buffer.WriteString("# TYPE systeminfo_load_15 guage\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_load_15 %v\n", metrics.Load.Load15))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString("# HELP systeminfo_cpu_cores Number or Cores\n")
	buffer.WriteString("# TYPE systeminfo_cpu_cores counter\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_cpu_cores %v\n", metrics.CPU.CPUCount))
	buffer.WriteString(fmt.Sprintf("\n"))
	buffer.WriteString("# HELP systeminfo_cpu_load load\n")
	buffer.WriteString("# TYPE systeminfo_cpu_load guage\n")
	buffer.WriteString(fmt.Sprintf("systeminfo_cpu_load %v\n", metrics.CPU.CPULoad))
	buffer.WriteString(fmt.Sprintf("\n"))

	return buffer.String()
}