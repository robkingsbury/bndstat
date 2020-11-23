# bndstat
[![Actions Status](https://github.com/robkingsbury/bndstat/workflows/Build/badge.svg)](https://github.com/robkingsbury/bndstat/actions)
[![Actions Status](https://github.com/robkingsbury/bndstat/workflows/Test/badge.svg)](https://github.com/robkingsbury/bndstat/actions)

A simple Go program that displays throughput stats for network interfaces.

## Quickstart

Assuming you have your Go environment and path set up in a
[standard way](https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies),
this should get you started.

```
$ git clone https://github.com/robkingsbury/bndstat
$ cd bndstat/util
$ ./install.sh
$ bndstat
$ bndstat --help
```

## Current Version

```
$ bndstat --version
bndstat (untagged version)
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 4479bb9
Compiled: Mon 23 Nov 13:28:14 PST 2020
Build Host: bender
```

## Usage

```
$ bndstat --help
Usage: bndstat [option]... [interval [count]]
Output the average throughput of network devices over a time interval.

Interval and count have the same behavior as the options of the same
name. However, when both an option and the non-option arg are present,
the value specified by the option takes precedence.

Interval is specified as a float for both the option and arg.

Options [default]:
  --interval=seconds    Number of seconds between updates [1.0]
  --count=num           Number of updates to print, any num less than one
                          will output infinite updates [0]
  --devices=list        Comma separated list of devices to output; if empty,
                          all non-loopback devices are included [empty]
  --unit=string         Specify the output unit; is one of bps, kbps, mbps
                          or tbps; input is not case sensitive [kbps]
  --showunit            Show the --unit used in the output header [false]
  --version             Print version information [false]
  --helpfull            List all available options [false]
```

### Examples

In this example, running on a raspberry pi that I am using as a router, eth0 is the hardline connection to wifi router, eth1 is my
primary internet provider, eth2 is my backup internet line and wlan0 is a connection to the wifi network:

```
$ bndstat 3 5
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
     763.22       72.58      29.61      765.67       0.37        0.42       1.51        0.00
     714.91       36.92      28.56      715.25       0.37        0.42       2.51        0.00
    1422.51     2280.85    2237.27     1435.40       0.25        0.28      16.90        0.00
    2159.93     2746.71    2710.01     2180.72       0.37        0.42      49.82        0.00
    1402.60      197.03     188.29     1413.15       0.37        0.42       4.99        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
     861.18      508.35       0.37        0.42
    1166.09     1112.78       0.37        0.42
      37.46     1219.94       0.37        0.42
     676.05      530.48       1.33        1.55
     117.67      387.08       0.37        0.42
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1123 13:28:35.298751     459 bndstat.go:100] interval = 1.000000, count = 1
I1123 13:28:35.299531     459 throughput.go:21] os is "linux"
I1123 13:28:35.299586     459 throughput.go:33] running Reporter.Report() twice to prime the stats
I1123 13:28:35.299629     459 throughput.go:35] prime 1
I1123 13:28:35.300056     459 linux.go:113] found device eth0
I1123 13:28:35.300107     459 linux.go:118] bytesRecvStr for eth0: 1523136864
I1123 13:28:35.300151     459 linux.go:119] bytesTransStr for eth0: 2005601804
I1123 13:28:35.300202     459 linux.go:113] found device eth1
I1123 13:28:35.300242     459 linux.go:118] bytesRecvStr for eth1: 726407418286
I1123 13:28:35.300283     459 linux.go:119] bytesTransStr for eth1: 533945785486
I1123 13:28:35.300335     459 linux.go:113] found device eth2
I1123 13:28:35.300378     459 linux.go:118] bytesRecvStr for eth2: 1302236597
I1123 13:28:35.300420     459 linux.go:119] bytesTransStr for eth2: 1339577093
I1123 13:28:35.300472     459 linux.go:113] found device wlan0
I1123 13:28:35.300513     459 linux.go:118] bytesRecvStr for wlan0: 2480636495
I1123 13:28:35.300556     459 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:28:35.300606     459 linux.go:113] found device lo
I1123 13:28:35.300647     459 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:28:35.300690     459 linux.go:119] bytesTransStr for lo: 619794
I1123 13:28:35.300765     459 linux.go:65] updating state for eth0
I1123 13:28:35.300809     459 linux.go:65] updating state for eth1
I1123 13:28:35.300850     459 linux.go:65] updating state for eth2
I1123 13:28:35.300890     459 linux.go:65] updating state for wlan0
I1123 13:28:35.300930     459 linux.go:65] updating state for lo
I1123 13:28:35.301000     459 throughput.go:38] prime 2
I1123 13:28:35.301225     459 linux.go:113] found device eth0
I1123 13:28:35.301272     459 linux.go:118] bytesRecvStr for eth0: 1523136864
I1123 13:28:35.301314     459 linux.go:119] bytesTransStr for eth0: 2005601804
I1123 13:28:35.301365     459 linux.go:113] found device eth1
I1123 13:28:35.301407     459 linux.go:118] bytesRecvStr for eth1: 726407418286
I1123 13:28:35.301449     459 linux.go:119] bytesTransStr for eth1: 533945785486
I1123 13:28:35.301498     459 linux.go:113] found device eth2
I1123 13:28:35.301539     459 linux.go:118] bytesRecvStr for eth2: 1302236597
I1123 13:28:35.301581     459 linux.go:119] bytesTransStr for eth2: 1339577093
I1123 13:28:35.301632     459 linux.go:113] found device wlan0
I1123 13:28:35.301671     459 linux.go:118] bytesRecvStr for wlan0: 2480636495
I1123 13:28:35.301712     459 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:28:35.301761     459 linux.go:113] found device lo
I1123 13:28:35.301802     459 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:28:35.301844     459 linux.go:119] bytesTransStr for lo: 619794
I1123 13:28:35.301897     459 linux.go:65] updating state for eth0
I1123 13:28:35.301936     459 linux.go:65] updating state for eth1
I1123 13:28:35.301975     459 linux.go:65] updating state for eth2
I1123 13:28:35.302012     459 linux.go:65] updating state for wlan0
I1123 13:28:35.302056     459 linux.go:65] updating state for lo
I1123 13:28:35.302298     459 linux.go:113] found device eth0
I1123 13:28:35.302345     459 linux.go:118] bytesRecvStr for eth0: 1523136864
I1123 13:28:35.302385     459 linux.go:119] bytesTransStr for eth0: 2005601804
I1123 13:28:35.302435     459 linux.go:113] found device eth1
I1123 13:28:35.302476     459 linux.go:118] bytesRecvStr for eth1: 726407418286
I1123 13:28:35.302518     459 linux.go:119] bytesTransStr for eth1: 533945785486
I1123 13:28:35.302567     459 linux.go:113] found device eth2
I1123 13:28:35.302671     459 linux.go:118] bytesRecvStr for eth2: 1302236597
I1123 13:28:35.302717     459 linux.go:119] bytesTransStr for eth2: 1339577093
I1123 13:28:35.302769     459 linux.go:113] found device wlan0
I1123 13:28:35.302810     459 linux.go:118] bytesRecvStr for wlan0: 2480636495
I1123 13:28:35.302852     459 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:28:35.302901     459 linux.go:113] found device lo
I1123 13:28:35.302942     459 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:28:35.302981     459 linux.go:119] bytesTransStr for lo: 619794
I1123 13:28:35.303035     459 linux.go:65] updating state for eth0
I1123 13:28:35.303073     459 linux.go:65] updating state for eth1
I1123 13:28:35.303112     459 linux.go:65] updating state for eth2
I1123 13:28:35.303151     459 linux.go:65] updating state for wlan0
I1123 13:28:35.303189     459 linux.go:65] updating state for lo
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1123 13:28:36.303920     459 linux.go:113] found device eth0
I1123 13:28:36.304027     459 linux.go:118] bytesRecvStr for eth0: 1523376937
I1123 13:28:36.304075     459 linux.go:119] bytesTransStr for eth0: 2005987285
I1123 13:28:36.304131     459 linux.go:113] found device eth1
I1123 13:28:36.304176     459 linux.go:118] bytesRecvStr for eth1: 726407798839
I1123 13:28:36.304219     459 linux.go:119] bytesTransStr for eth1: 533946028228
I1123 13:28:36.304269     459 linux.go:113] found device eth2
I1123 13:28:36.304312     459 linux.go:118] bytesRecvStr for eth2: 1302236645
I1123 13:28:36.304355     459 linux.go:119] bytesTransStr for eth2: 1339577147
I1123 13:28:36.304406     459 linux.go:113] found device wlan0
I1123 13:28:36.304448     459 linux.go:118] bytesRecvStr for wlan0: 2480636786
I1123 13:28:36.304492     459 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:28:36.304542     459 linux.go:113] found device lo
I1123 13:28:36.304583     459 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:28:36.304625     459 linux.go:119] bytesTransStr for lo: 619794
I1123 13:28:36.304681     459 linux.go:65] updating state for eth0
I1123 13:28:36.304722     459 linux.go:65] updating state for eth1
I1123 13:28:36.304761     459 linux.go:65] updating state for eth2
I1123 13:28:36.304798     459 linux.go:65] updating state for wlan0
I1123 13:28:36.304836     459 linux.go:65] updating state for lo
I1123 13:28:36.304915     459 table.go:58] rows = 40, tableLineCount = 2
I1123 13:28:36.304962     459 table.go:69] tableLineCount = 2, rows-3 = 37
    1872.49     3006.63    2968.19     1893.31       0.37        0.42       2.27        0.00
```

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
