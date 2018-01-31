# systeminfo
Exports basic system info and metrics as a service.

## Installation

Installation is simple - just type at the command line:

```bash
go get github.com/riomhaire/systeminfo/infrastructure/application/systeminfo
```

## Building

If you have make install building and installation consists of the following command in the 'github.com/riomhaire/systeminfo' directory:

```bash
make
```

You should get something like:

```bash
Cleaning
Downloading Dependencies
Compiling Apps
--- systeminfo server
Done Compiling Apps
Running Unit Tests
?       github.com/riomhaire/systeminfo/entities        [no test files]
?       github.com/riomhaire/systeminfo/infrastructure/application/systeminfo   [no test files]
?       github.com/riomhaire/systeminfo/infrastructure/application/systeminfo/bootstrap [no test files]
?       github.com/riomhaire/systeminfo/infrastructure/web/rest [no test files]
ok      github.com/riomhaire/systeminfo/usecases        0.006s
Done
outpost :: github.com/riomhaire/systeminfo ‹feature/add_disk_network_metrics*› »
```
The 'systeminfo' application for your environment should be present in the root directory.

## What it Does.

The SystemInfo service is a wrapper around the excellent 'github.com/shirou/gopsutil' library and exposes its metrics as a rest and prometheus complient http service.

The service exports on port 3333 (can be changed via -port option) at the '/metrics' endpoint information similar to:


```json
{
	"memory": {
		"total": 8251195392,
		"available": 914178048,
		"used": 7337017344,
		"usedPercent": 88.92065931578415,
		"free": 162189312,
		"active": 5539987456,
		"inactive": 1694236672,
		"wired": 0,
		"buffers": 53530624,
		"cached": 1123106816,
		"writeback": 0,
		"dirty": 1564672,
		"writebacktmp": 0,
		"shared": 271060992,
		"slab": 283553792,
		"pagetables": 184344576,
		"swapcached": 29601792
	},
	"host": {
		"hostname": "deref",
		"uptime": 763842,
		"bootTime": 1516659876,
		"procs": 543,
		"os": "linux",
		"platform": "linuxmint",
		"platformFamily": "debian",
		"platformVersion": "18.3",
		"kernelVersion": "4.4.0-112-generic",
		"virtualizationSystem": "kvm",
		"virtualizationRole": "host",
		"hostid": "e82892e8-0e51-4613-8877-9f60e4ced996"
	},
	"load": {
		"load1": 2.61,
		"load5": 2.7,
		"load15": 2.72
	},
	"cpu": {
		"cpus": 8,
		"cpuLoad": 19.2319142601709
	},
	"storage": [
		{
			"path": "/",
			"fstype": "ext2/ext3",
			"total": 236616474624,
			"free": 26115538944,
			"used": 198457831424,
			"usedPercent": 83.87320947933286,
			"inodesTotal": 14688256,
			"inodesUsed": 2017904,
			"inodesFree": 12670352,
			"inodesUsedPercent": 13.738213713050754
		},
		{
			"path": "/snap/core/3604",
			"fstype": "squashfs",
			"total": 87949312,
			"free": 0,
			"used": 87949312,
			"usedPercent": 100,
			"inodesTotal": 13647,
			"inodesUsed": 13647,
			"inodesFree": 0,
			"inodesUsedPercent": 100
		},
		{
			"path": "/snap/core/3748",
			"fstype": "squashfs",
			"total": 87949312,
			"free": 0,
			"used": 87949312,
			"usedPercent": 100,
			"inodesTotal": 13648,
			"inodesUsed": 13648,
			"inodesFree": 0,
			"inodesUsedPercent": 100
		},
		{
			"path": "/snap/conjure-up/924",
			"fstype": "squashfs",
			"total": 91226112,
			"free": 0,
			"used": 91226112,
			"usedPercent": 100,
			"inodesTotal": 12263,
			"inodesUsed": 12263,
			"inodesFree": 0,
			"inodesUsedPercent": 100
		},
		{
			"path": "/snap/conjure-up/919",
			"fstype": "squashfs",
			"total": 91226112,
			"free": 0,
			"used": 91226112,
			"usedPercent": 100,
			"inodesTotal": 12259,
			"inodesUsed": 12259,
			"inodesFree": 0,
			"inodesUsedPercent": 100
		},
		{
			"path": "/snap/conjure-up/859",
			"fstype": "squashfs",
			"total": 82182144,
			"free": 0,
			"used": 82182144,
			"usedPercent": 100,
			"inodesTotal": 12207,
			"inodesUsed": 12207,
			"inodesFree": 0,
			"inodesUsedPercent": 100
		},
		{
			"path": "/boot",
			"fstype": "ext2/ext3",
			"total": 495560704,
			"free": 142810112,
			"used": 327165952,
			"usedPercent": 66.01934926624044,
			"inodesTotal": 124928,
			"inodesUsed": 318,
			"inodesFree": 124610,
			"inodesUsedPercent": 0.254546618852459
		},
		{
			"path": "/boot/efi",
			"fstype": "msdos",
			"total": 535805952,
			"free": 532238336,
			"used": 3567616,
			"usedPercent": 0.6658410543375226,
			"inodesTotal": 0,
			"inodesUsed": 0,
			"inodesFree": 0,
			"inodesUsedPercent": 0
		},
		{
			"path": "/var/lib/docker/aufs",
			"fstype": "ext2/ext3",
			"total": 236616474624,
			"free": 26115538944,
			"used": 198457831424,
			"usedPercent": 83.87320947933286,
			"inodesTotal": 14688256,
			"inodesUsed": 2017904,
			"inodesFree": 12670352,
			"inodesUsedPercent": 13.738213713050754
		},
		{
			"path": "/snap/core/3887",
			"fstype": "squashfs",
			"total": 85327872,
			"free": 0,
			"used": 85327872,
			"usedPercent": 100,
			"inodesTotal": 12821,
			"inodesUsed": 12821,
			"inodesFree": 0,
			"inodesUsedPercent": 100
		}
	],
	"network": [
		{
			"name": "all",
			"bytesSent": 8102684507,
			"bytesRecv": 18884073707,
			"packetsSent": 13991777,
			"packetsRecv": 21063684,
			"errin": 0,
			"errout": 0,
			"dropin": 0,
			"dropout": 0,
			"fifoin": 0,
			"fifoout": 0
		}
	]
}
```
If the accept type is 'application/json' and :

```
# HELP systeminfo_memory_bytes How much memory.
# TYPE systeminfo_memory_bytes gauge
systeminfo_memory_bytes 8251195392

# HELP systeminfo_memory_available_bytes How much memory available.
# TYPE systeminfo_memory_available_bytes gauge
systeminfo_memory_available_bytes 974381056

# HELP systeminfo_memory_used_bytes How much memory available.
# TYPE systeminfo_memory_used_bytes gauge
systeminfo_memory_used_bytes 7276814336

# HELP systeminfo_memory_used_percentage How much memory used as percentage.
# TYPE systeminfo_memory_used_percentage gauge
systeminfo_memory_used_percentage 88.19103160561781


# HELP systeminfo_host_uptime How long been on
# TYPE systeminfo_host_uptime counter
systeminfo_host_uptime 764990

# HELP systeminfo_host_procs How Many Processes
# TYPE systeminfo_host_procs guage
systeminfo_procs 547


# HELP systeminfo_load_1 Load Last Minute
# TYPE systeminfo_load_1 guage
systeminfo_load_1 2.54

# HELP systeminfo_load_5 Load Last 5 Minute
# TYPE systeminfo_load_5 guage
systeminfo_load_5 2.59

# HELP systeminfo_load_15 Load Last 15 Minute
# TYPE systeminfo_load_15 guage
systeminfo_load_15 2.62


# HELP systeminfo_cpu_cores Number or Cores
# TYPE systeminfo_cpu_cores counter
systeminfo_cpu_cores 8

# HELP systeminfo_cpu_load load
# TYPE systeminfo_cpu_load guage
systeminfo_cpu_load 31.08458744078956


# HELP systeminfo_network_bytes_sent Network bytes sent
# TYPE systeminfo_network_bytes_sent counter
systeminfo_network_bytes_sent 8172723733

# HELP systeminfo_network_bytes_received Network bytes received
# TYPE systeminfo_network_bytes_received counter
systeminfo_network_bytes_received 18927111005

# HELP systeminfo_network_packets_sent Network packets sent
# TYPE systeminfo_network_packets_sent counter
systeminfo_network_packets_sent 14048251

# HELP systeminfo_network_packets_received Network packets received
# TYPE systeminfo_network_packets_received counter
systeminfo_network_packets_received 21114441



```
If type is 'text/plain'.

A very basic but functional service for a machine.

### Usage with Prometheus and Grafana

We are not going to get into how to set up prometheus and grafana since there are many good sites on the wen which will take you through this. The only thing to do is to tell prometheus to access the machine and port (usually port 3333) where systeminfo is running. All the variables exported to prometheus start with 'systeminfo_' and can be gleaned from the above example.

One simple example of a grafina dashboard is:

![Example Grafina dashboard](https://github.com/riomhaire/systeminfo/blob/master/docs/images/grafana.png)

Hope this project is of some interest and use.

