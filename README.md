# systeminfo
Exports basic system info and metrics as a service.

The service exports on port 3333 (can be changed via -port option) at the '/metrics' endpoint information similar to:


```json
{
	"memory": {
		"total": 1001472000,
		"available": 874307584,
		"used": 127164416,
		"usedPercent": 12.697750511247444,
		"free": 142577664,
		"active": 354631680,
		"inactive": 433770496,
		"wired": 0,
		"buffers": 59011072,
		"cached": 700710912,
		"writeback": 0,
		"dirty": 16384,
		"writebacktmp": 0,
		"shared": 6619136,
		"slab": 57503744,
		"pagetables": 1146880,
		"swapcached": 20480
	},
	"host": {
		"hostname": "bugsbunny",
		"uptime": 920325,
		"bootTime": 1516253243,
		"procs": 179,
		"os": "linux",
		"platform": "raspbian",
		"platformFamily": "debian",
		"platformVersion": "9.3",
		"kernelVersion": "4.9.59-v7+",
		"virtualizationSystem": "",
		"virtualizationRole": "",
		"hostid": "9e789dee-23d1-4283-b6da-d2081860c36d"
	},
	"load": {
		"load1": 0.08,
		"load5": 0.02,
		"load15": 0.01
	},
	"cpu": {
		"cpus": 4,
		"cpuLoad": 7.158351409456117
	}
}

```
If the accept type is 'application/json' and :

```
# HELP systeminfo_memory_bytes How much memory.
# TYPE systeminfo_memory_bytes gauge
systeminfo_memory_bytes 33420267520

# HELP systeminfo_memory_available_bytes How much memory available.
# TYPE systeminfo_memory_available_bytes gauge
systeminfo_memory_available_bytes 27293855744

# HELP systeminfo_memory_used_percentage How much memory used as percentage.
# TYPE systeminfo_memory_used_percentage gauge
systeminfo_memory_used_percentage 18.331426498407634


# HELP systeminfo_host_uptime How long been on
# TYPE systeminfo_host_uptime counter
systeminfo_host_uptime 990830

# HELP systeminfo_host_procs How Many Processes
# TYPE systeminfo_host_procs guage
systeminfo_procs 560


# HELP systeminfo_load_1 Load Last Minute
# TYPE systeminfo_load_1 guage
systeminfo_load_1 0.12

# HELP systeminfo_load_5 Load Last 5 Minute
# TYPE systeminfo_load_5 guage
systeminfo_load_5 0.08

# HELP systeminfo_load_15 Load Last 15 Minute
# TYPE systeminfo_load_15 guage
systeminfo_load_15 0.02


# HELP systeminfo_cpu_cores Number or Cores
# TYPE systeminfo_cpu_cores guage
systeminfo_cpu_cores 8

# HELP systeminfo_cpu_load Number or Cores
# TYPE systeminfo_cpu_load guage
systeminfo_cpu_load 1.0370978545021885


```
If type is 'text/plain'.

A very basic but functional service for a machine.
