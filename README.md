# systeminfo
Exports basic system info and metrics as a service.

The service exports on port 3333 (can be changed via -port option) at the '/metrics' endpoint information similar to:


```
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

A very basic but functional service for a machine.
