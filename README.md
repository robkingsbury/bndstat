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
bndstat v0.4.2
Rob Kingsbury
https://github.com/robkingsbury/bndstat
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

### Example

In this example, running on a raspberry pi that I am using as a router, eth0 is the hardline connection to wifi router, eth1 is my
primary internet provider, eth2 is my backup internet line and wlan0 is a connection to the wifi network:

```
$ bndstat 3 5
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
     624.19      240.77     235.62      629.79       0.37        0.41      16.07        0.00
     840.62     5823.62    5772.71      848.57       0.37        0.42       0.00        0.00
     754.07     2090.82    2070.02      761.05       0.37        0.55       0.71        0.00
     947.58     1173.05    1160.70      954.52       0.37        0.28       0.00        0.00
     629.30       51.50      47.76      634.72       0.37        0.42       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
     711.39     1119.57       0.37        0.42
      99.64      294.33       0.37        0.42
     206.09     2910.00       0.00        0.42
      33.08      757.87       0.95        0.73
      83.59     1335.66       0.37        0.42
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1102 08:30:06.150584    3527 bndstat.go:96] interval = 1.000000, count = 1
I1102 08:30:06.151319    3527 throughput.go:21] os is "linux"
I1102 08:30:06.151397    3527 throughput.go:33] running Reporter.Report() twice to prime the stats
I1102 08:30:06.151444    3527 throughput.go:35] prime 1
I1102 08:30:06.151874    3527 linux.go:113] found device eth0
I1102 08:30:06.151924    3527 linux.go:118] bytesRecvStr for eth0: 2979227700
I1102 08:30:06.151968    3527 linux.go:119] bytesTransStr for eth0: 1960737705
I1102 08:30:06.152022    3527 linux.go:113] found device eth1
I1102 08:30:06.152064    3527 linux.go:118] bytesRecvStr for eth1: 515467989925
I1102 08:30:06.152106    3527 linux.go:119] bytesTransStr for eth1: 301998180898
I1102 08:30:06.152156    3527 linux.go:113] found device eth2
I1102 08:30:06.152197    3527 linux.go:118] bytesRecvStr for eth2: 292710251
I1102 08:30:06.152237    3527 linux.go:119] bytesTransStr for eth2: 399900410
I1102 08:30:06.152287    3527 linux.go:113] found device wlan0
I1102 08:30:06.152328    3527 linux.go:118] bytesRecvStr for wlan0: 3796113093
I1102 08:30:06.152372    3527 linux.go:119] bytesTransStr for wlan0: 99353350
I1102 08:30:06.152421    3527 linux.go:113] found device lo
I1102 08:30:06.152463    3527 linux.go:118] bytesRecvStr for lo: 107348
I1102 08:30:06.152504    3527 linux.go:119] bytesTransStr for lo: 107348
I1102 08:30:06.152560    3527 linux.go:65] updating state for eth0
I1102 08:30:06.152604    3527 linux.go:65] updating state for eth1
I1102 08:30:06.152645    3527 linux.go:65] updating state for eth2
I1102 08:30:06.152686    3527 linux.go:65] updating state for wlan0
I1102 08:30:06.152725    3527 linux.go:65] updating state for lo
I1102 08:30:06.152797    3527 throughput.go:38] prime 2
I1102 08:30:06.153025    3527 linux.go:113] found device eth0
I1102 08:30:06.153072    3527 linux.go:118] bytesRecvStr for eth0: 2979227700
I1102 08:30:06.153115    3527 linux.go:119] bytesTransStr for eth0: 1960737705
I1102 08:30:06.153166    3527 linux.go:113] found device eth1
I1102 08:30:06.153227    3527 linux.go:118] bytesRecvStr for eth1: 515467989925
I1102 08:30:06.153271    3527 linux.go:119] bytesTransStr for eth1: 301998180898
I1102 08:30:06.153320    3527 linux.go:113] found device eth2
I1102 08:30:06.153361    3527 linux.go:118] bytesRecvStr for eth2: 292710251
I1102 08:30:06.153403    3527 linux.go:119] bytesTransStr for eth2: 399900410
I1102 08:30:06.153454    3527 linux.go:113] found device wlan0
I1102 08:30:06.153495    3527 linux.go:118] bytesRecvStr for wlan0: 3796113093
I1102 08:30:06.153539    3527 linux.go:119] bytesTransStr for wlan0: 99353350
I1102 08:30:06.153589    3527 linux.go:113] found device lo
I1102 08:30:06.153630    3527 linux.go:118] bytesRecvStr for lo: 107348
I1102 08:30:06.153668    3527 linux.go:119] bytesTransStr for lo: 107348
I1102 08:30:06.153720    3527 linux.go:65] updating state for eth0
I1102 08:30:06.153761    3527 linux.go:65] updating state for eth1
I1102 08:30:06.153800    3527 linux.go:65] updating state for eth2
I1102 08:30:06.153840    3527 linux.go:65] updating state for wlan0
I1102 08:30:06.153878    3527 linux.go:65] updating state for lo
I1102 08:30:06.154122    3527 linux.go:113] found device eth0
I1102 08:30:06.154169    3527 linux.go:118] bytesRecvStr for eth0: 2979227700
I1102 08:30:06.154209    3527 linux.go:119] bytesTransStr for eth0: 1960737705
I1102 08:30:06.154262    3527 linux.go:113] found device eth1
I1102 08:30:06.154304    3527 linux.go:118] bytesRecvStr for eth1: 515467989925
I1102 08:30:06.154345    3527 linux.go:119] bytesTransStr for eth1: 301998180898
I1102 08:30:06.154395    3527 linux.go:113] found device eth2
I1102 08:30:06.154435    3527 linux.go:118] bytesRecvStr for eth2: 292710251
I1102 08:30:06.154477    3527 linux.go:119] bytesTransStr for eth2: 399900410
I1102 08:30:06.154528    3527 linux.go:113] found device wlan0
I1102 08:30:06.154569    3527 linux.go:118] bytesRecvStr for wlan0: 3796113093
I1102 08:30:06.154610    3527 linux.go:119] bytesTransStr for wlan0: 99353350
I1102 08:30:06.154659    3527 linux.go:113] found device lo
I1102 08:30:06.154699    3527 linux.go:118] bytesRecvStr for lo: 107348
I1102 08:30:06.154740    3527 linux.go:119] bytesTransStr for lo: 107348
I1102 08:30:06.154794    3527 linux.go:65] updating state for eth0
I1102 08:30:06.154834    3527 linux.go:65] updating state for eth1
I1102 08:30:06.154873    3527 linux.go:65] updating state for eth2
I1102 08:30:06.154912    3527 linux.go:65] updating state for wlan0
I1102 08:30:06.155035    3527 linux.go:65] updating state for lo
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1102 08:30:07.155786    3527 linux.go:113] found device eth0
I1102 08:30:07.155866    3527 linux.go:118] bytesRecvStr for eth0: 2979258631
I1102 08:30:07.155912    3527 linux.go:119] bytesTransStr for eth0: 1960739790
I1102 08:30:07.155965    3527 linux.go:113] found device eth1
I1102 08:30:07.156005    3527 linux.go:118] bytesRecvStr for eth1: 515467991830
I1102 08:30:07.156044    3527 linux.go:119] bytesTransStr for eth1: 301998212201
I1102 08:30:07.156092    3527 linux.go:113] found device eth2
I1102 08:30:07.156129    3527 linux.go:118] bytesRecvStr for eth2: 292710299
I1102 08:30:07.156169    3527 linux.go:119] bytesTransStr for eth2: 399900464
I1102 08:30:07.156217    3527 linux.go:113] found device wlan0
I1102 08:30:07.156256    3527 linux.go:118] bytesRecvStr for wlan0: 3796113093
I1102 08:30:07.156294    3527 linux.go:119] bytesTransStr for wlan0: 99353350
I1102 08:30:07.156341    3527 linux.go:113] found device lo
I1102 08:30:07.156380    3527 linux.go:118] bytesRecvStr for lo: 107348
I1102 08:30:07.156417    3527 linux.go:119] bytesTransStr for lo: 107348
I1102 08:30:07.156468    3527 linux.go:65] updating state for eth0
I1102 08:30:07.156507    3527 linux.go:65] updating state for eth1
I1102 08:30:07.156543    3527 linux.go:65] updating state for eth2
I1102 08:30:07.156579    3527 linux.go:65] updating state for wlan0
I1102 08:30:07.156614    3527 linux.go:65] updating state for lo
I1102 08:30:07.156687    3527 table.go:58] rows = 40, tableLineCount = 2
I1102 08:30:07.156730    3527 table.go:69] tableLineCount = 2, rows-3 = 37
     241.24       16.26      14.86      244.15       0.37        0.42       0.00        0.00
```

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
