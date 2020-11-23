# bndstat
[![Actions Status](https://github.com/robkingsbury/bndstat/workflows/Build/badge.svg)](https://github.com/robkingsbury/bndstat/actions)
[![Actions Status](https://github.com/robkingsbury/bndstat/workflows/Test/badge.svg)](https://github.com/robkingsbury/bndstat/actions)

A simple Go program that displays throughput stats for network interfaces.

## Quickstart

```
$ git clone https://github.com/robkingsbury/bndstat
$ cd bndstat
$ go build
$ ./bndstat
$ ./bndstat --help
```

## Current Version

```
$ bndstat --version
bndstat v0.4.3
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 9973dd6
Compiled: Mon 23 Nov 13:21:49 PST 2020
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
     694.78       33.68      29.97      699.91       0.50        0.41       2.27        0.00
     853.92       42.97      38.91      860.07       0.37        0.42       2.27        0.00
     746.98       27.23      24.25      753.08       0.37        0.42       2.22        0.00
     823.84       56.76      51.49      829.91       0.37        0.42       2.51        0.00
     732.99       40.84      36.40      738.61       0.37        0.42       2.27        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      31.57      921.49       0.37        0.42
      53.21      791.79       0.95        1.15
      29.11      973.08       0.75        0.81
      23.31      929.47       0.37        0.42
      29.97      557.78       0.37        0.42
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1123 13:22:11.065456   32631 bndstat.go:100] interval = 1.000000, count = 1
I1123 13:22:11.066045   32631 throughput.go:21] os is "linux"
I1123 13:22:11.066186   32631 throughput.go:33] running Reporter.Report() twice to prime the stats
I1123 13:22:11.066255   32631 throughput.go:35] prime 1
I1123 13:22:11.066616   32631 linux.go:113] found device eth0
I1123 13:22:11.066690   32631 linux.go:118] bytesRecvStr for eth0: 1482014574
I1123 13:22:11.066802   32631 linux.go:119] bytesTransStr for eth0: 1987710915
I1123 13:22:11.066878   32631 linux.go:113] found device eth1
I1123 13:22:11.066939   32631 linux.go:118] bytesRecvStr for eth1: 726390173352
I1123 13:22:11.067035   32631 linux.go:119] bytesTransStr for eth1: 533904435768
I1123 13:22:11.067106   32631 linux.go:113] found device eth2
I1123 13:22:11.067168   32631 linux.go:118] bytesRecvStr for eth2: 1302215811
I1123 13:22:11.067230   32631 linux.go:119] bytesTransStr for eth2: 1339553185
I1123 13:22:11.067302   32631 linux.go:113] found device wlan0
I1123 13:22:11.067361   32631 linux.go:118] bytesRecvStr for wlan0: 2480313862
I1123 13:22:11.067425   32631 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:22:11.067495   32631 linux.go:113] found device lo
I1123 13:22:11.067555   32631 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:22:11.067618   32631 linux.go:119] bytesTransStr for lo: 619794
I1123 13:22:11.067695   32631 linux.go:65] updating state for eth0
I1123 13:22:11.067784   32631 linux.go:65] updating state for eth1
I1123 13:22:11.067848   32631 linux.go:65] updating state for eth2
I1123 13:22:11.067912   32631 linux.go:65] updating state for wlan0
I1123 13:22:11.067972   32631 linux.go:65] updating state for lo
I1123 13:22:11.068060   32631 throughput.go:38] prime 2
I1123 13:22:11.068274   32631 linux.go:113] found device eth0
I1123 13:22:11.068341   32631 linux.go:118] bytesRecvStr for eth0: 1482014688
I1123 13:22:11.068404   32631 linux.go:119] bytesTransStr for eth0: 1987710915
I1123 13:22:11.068476   32631 linux.go:113] found device eth1
I1123 13:22:11.068537   32631 linux.go:118] bytesRecvStr for eth1: 726390173352
I1123 13:22:11.068600   32631 linux.go:119] bytesTransStr for eth1: 533904435890
I1123 13:22:11.068682   32631 linux.go:113] found device eth2
I1123 13:22:11.068744   32631 linux.go:118] bytesRecvStr for eth2: 1302215811
I1123 13:22:11.068807   32631 linux.go:119] bytesTransStr for eth2: 1339553185
I1123 13:22:11.068907   32631 linux.go:113] found device wlan0
I1123 13:22:11.068967   32631 linux.go:118] bytesRecvStr for wlan0: 2480313862
I1123 13:22:11.069030   32631 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:22:11.069136   32631 linux.go:113] found device lo
I1123 13:22:11.069231   32631 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:22:11.069294   32631 linux.go:119] bytesTransStr for lo: 619794
I1123 13:22:11.069367   32631 linux.go:65] updating state for eth0
I1123 13:22:11.069427   32631 linux.go:65] updating state for eth1
I1123 13:22:11.069487   32631 linux.go:65] updating state for eth2
I1123 13:22:11.069546   32631 linux.go:65] updating state for wlan0
I1123 13:22:11.069630   32631 linux.go:65] updating state for lo
I1123 13:22:11.069885   32631 linux.go:113] found device eth0
I1123 13:22:11.069955   32631 linux.go:118] bytesRecvStr for eth0: 1482014688
I1123 13:22:11.070019   32631 linux.go:119] bytesTransStr for eth0: 1987710915
I1123 13:22:11.070089   32631 linux.go:113] found device eth1
I1123 13:22:11.070150   32631 linux.go:118] bytesRecvStr for eth1: 726390173352
I1123 13:22:11.070212   32631 linux.go:119] bytesTransStr for eth1: 533904435890
I1123 13:22:11.070282   32631 linux.go:113] found device eth2
I1123 13:22:11.070341   32631 linux.go:118] bytesRecvStr for eth2: 1302215811
I1123 13:22:11.070467   32631 linux.go:119] bytesTransStr for eth2: 1339553185
I1123 13:22:11.070542   32631 linux.go:113] found device wlan0
I1123 13:22:11.070601   32631 linux.go:118] bytesRecvStr for wlan0: 2480313862
I1123 13:22:11.070664   32631 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:22:11.070733   32631 linux.go:113] found device lo
I1123 13:22:11.070795   32631 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:22:11.070856   32631 linux.go:119] bytesTransStr for lo: 619794
I1123 13:22:11.070928   32631 linux.go:65] updating state for eth0
I1123 13:22:11.070989   32631 linux.go:65] updating state for eth1
I1123 13:22:11.071049   32631 linux.go:65] updating state for eth2
I1123 13:22:11.071108   32631 linux.go:65] updating state for wlan0
I1123 13:22:11.071167   32631 linux.go:65] updating state for lo
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1123 13:22:12.072214   32631 linux.go:113] found device eth0
I1123 13:22:12.072298   32631 linux.go:118] bytesRecvStr for eth0: 1482153406
I1123 13:22:12.072346   32631 linux.go:119] bytesTransStr for eth0: 1987716982
I1123 13:22:12.072401   32631 linux.go:113] found device eth1
I1123 13:22:12.072455   32631 linux.go:118] bytesRecvStr for eth1: 726390178773
I1123 13:22:12.072500   32631 linux.go:119] bytesTransStr for eth1: 533904575487
I1123 13:22:12.072551   32631 linux.go:113] found device eth2
I1123 13:22:12.072594   32631 linux.go:118] bytesRecvStr for eth2: 1302215859
I1123 13:22:12.072637   32631 linux.go:119] bytesTransStr for eth2: 1339553239
I1123 13:22:12.072687   32631 linux.go:113] found device wlan0
I1123 13:22:12.072728   32631 linux.go:118] bytesRecvStr for wlan0: 2480314153
I1123 13:22:12.072773   32631 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:22:12.072820   32631 linux.go:113] found device lo
I1123 13:22:12.072862   32631 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:22:12.072901   32631 linux.go:119] bytesTransStr for lo: 619794
I1123 13:22:12.072957   32631 linux.go:65] updating state for eth0
I1123 13:22:12.072999   32631 linux.go:65] updating state for eth1
I1123 13:22:12.073035   32631 linux.go:65] updating state for eth2
I1123 13:22:12.073074   32631 linux.go:65] updating state for wlan0
I1123 13:22:12.073109   32631 linux.go:65] updating state for lo
I1123 13:22:12.073187   32631 table.go:58] rows = 40, tableLineCount = 2
I1123 13:22:12.073235   32631 table.go:69] tableLineCount = 2, rows-3 = 37
    1081.54       47.30      42.27     1088.40       0.37        0.42       2.27        0.00
```

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
