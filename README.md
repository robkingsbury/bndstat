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
    1251.53    33617.86   33331.59     1268.24       0.25        0.42       3.79        0.00
     697.79       35.90      32.42      703.50       0.37        0.28       1.84        0.00
     569.79       34.88      31.10      574.71       0.25        0.42      58.12        0.00
     891.77       60.72      55.66      899.37       0.37        0.28      12.84        0.00
     860.71       32.02      27.48      867.36       0.37        0.42       2.32        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      23.88      366.96       0.37        0.42
      29.72      242.19       0.37        0.42
      33.16     1166.56       0.00        0.42
      28.87     1309.89       0.75        0.39
      23.36      338.26       0.95        1.16
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1101 16:48:01.878568    1316 bndstat.go:96] interval = 1.000000, count = 1
I1101 16:48:01.879423    1316 throughput.go:21] os is "linux"
I1101 16:48:01.879570    1316 throughput.go:33] running Reporter.Report() twice to prime the stats
I1101 16:48:01.879641    1316 throughput.go:35] prime 1
I1101 16:48:01.880016    1316 linux.go:113] found device eth0
I1101 16:48:01.880086    1316 linux.go:118] bytesRecvStr for eth0: 1128195486
I1101 16:48:01.880152    1316 linux.go:119] bytesTransStr for eth0: 609303733
I1101 16:48:01.880226    1316 linux.go:113] found device eth1
I1101 16:48:01.880286    1316 linux.go:118] bytesRecvStr for eth1: 509896225016
I1101 16:48:01.880403    1316 linux.go:119] bytesTransStr for eth1: 295797812730
I1101 16:48:01.880474    1316 linux.go:113] found device eth2
I1101 16:48:01.880536    1316 linux.go:118] bytesRecvStr for eth2: 289863229
I1101 16:48:01.880598    1316 linux.go:119] bytesTransStr for eth2: 396610830
I1101 16:48:01.880671    1316 linux.go:113] found device wlan0
I1101 16:48:01.880731    1316 linux.go:118] bytesRecvStr for wlan0: 3710283538
I1101 16:48:01.880794    1316 linux.go:119] bytesTransStr for wlan0: 99351616
I1101 16:48:01.880895    1316 linux.go:113] found device lo
I1101 16:48:01.880955    1316 linux.go:118] bytesRecvStr for lo: 107348
I1101 16:48:01.881018    1316 linux.go:119] bytesTransStr for lo: 107348
I1101 16:48:01.881090    1316 linux.go:65] updating state for eth0
I1101 16:48:01.881155    1316 linux.go:65] updating state for eth1
I1101 16:48:01.881215    1316 linux.go:65] updating state for eth2
I1101 16:48:01.881278    1316 linux.go:65] updating state for wlan0
I1101 16:48:01.881338    1316 linux.go:65] updating state for lo
I1101 16:48:01.881425    1316 throughput.go:38] prime 2
I1101 16:48:01.881637    1316 linux.go:113] found device eth0
I1101 16:48:01.881705    1316 linux.go:118] bytesRecvStr for eth0: 1128195486
I1101 16:48:01.881768    1316 linux.go:119] bytesTransStr for eth0: 609303733
I1101 16:48:01.881855    1316 linux.go:113] found device eth1
I1101 16:48:01.881916    1316 linux.go:118] bytesRecvStr for eth1: 509896225016
I1101 16:48:01.881979    1316 linux.go:119] bytesTransStr for eth1: 295797812730
I1101 16:48:01.882049    1316 linux.go:113] found device eth2
I1101 16:48:01.882107    1316 linux.go:118] bytesRecvStr for eth2: 289863229
I1101 16:48:01.882170    1316 linux.go:119] bytesTransStr for eth2: 396610830
I1101 16:48:01.882239    1316 linux.go:113] found device wlan0
I1101 16:48:01.882298    1316 linux.go:118] bytesRecvStr for wlan0: 3710283538
I1101 16:48:01.882361    1316 linux.go:119] bytesTransStr for wlan0: 99351616
I1101 16:48:01.882428    1316 linux.go:113] found device lo
I1101 16:48:01.882488    1316 linux.go:118] bytesRecvStr for lo: 107348
I1101 16:48:01.882548    1316 linux.go:119] bytesTransStr for lo: 107348
I1101 16:48:01.882622    1316 linux.go:65] updating state for eth0
I1101 16:48:01.882687    1316 linux.go:65] updating state for eth1
I1101 16:48:01.882749    1316 linux.go:65] updating state for eth2
I1101 16:48:01.882810    1316 linux.go:65] updating state for wlan0
I1101 16:48:01.882870    1316 linux.go:65] updating state for lo
I1101 16:48:01.883108    1316 linux.go:113] found device eth0
I1101 16:48:01.883175    1316 linux.go:118] bytesRecvStr for eth0: 1128197142
I1101 16:48:01.883237    1316 linux.go:119] bytesTransStr for eth0: 609303733
I1101 16:48:01.883308    1316 linux.go:113] found device eth1
I1101 16:48:01.883368    1316 linux.go:118] bytesRecvStr for eth1: 509896225016
I1101 16:48:01.883429    1316 linux.go:119] bytesTransStr for eth1: 295797814402
I1101 16:48:01.883498    1316 linux.go:113] found device eth2
I1101 16:48:01.883557    1316 linux.go:118] bytesRecvStr for eth2: 289863229
I1101 16:48:01.883620    1316 linux.go:119] bytesTransStr for eth2: 396610830
I1101 16:48:01.883692    1316 linux.go:113] found device wlan0
I1101 16:48:01.883752    1316 linux.go:118] bytesRecvStr for wlan0: 3710283538
I1101 16:48:01.883815    1316 linux.go:119] bytesTransStr for wlan0: 99351616
I1101 16:48:01.883883    1316 linux.go:113] found device lo
I1101 16:48:01.883942    1316 linux.go:118] bytesRecvStr for lo: 107348
I1101 16:48:01.884003    1316 linux.go:119] bytesTransStr for lo: 107348
I1101 16:48:01.884075    1316 linux.go:65] updating state for eth0
I1101 16:48:01.884134    1316 linux.go:65] updating state for eth1
I1101 16:48:01.884195    1316 linux.go:65] updating state for eth2
I1101 16:48:01.884323    1316 linux.go:65] updating state for wlan0
I1101 16:48:01.884387    1316 linux.go:65] updating state for lo
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1101 16:48:02.885425    1316 linux.go:113] found device eth0
I1101 16:48:02.885509    1316 linux.go:118] bytesRecvStr for eth0: 1128253569
I1101 16:48:02.885555    1316 linux.go:119] bytesTransStr for eth0: 609312247
I1101 16:48:02.885615    1316 linux.go:113] found device eth1
I1101 16:48:02.885659    1316 linux.go:118] bytesRecvStr for eth1: 509896233208
I1101 16:48:02.885701    1316 linux.go:119] bytesTransStr for eth1: 295797869927
I1101 16:48:02.885751    1316 linux.go:113] found device eth2
I1101 16:48:02.885792    1316 linux.go:118] bytesRecvStr for eth2: 289863277
I1101 16:48:02.885836    1316 linux.go:119] bytesTransStr for eth2: 396610884
I1101 16:48:02.885896    1316 linux.go:113] found device wlan0
I1101 16:48:02.885936    1316 linux.go:118] bytesRecvStr for wlan0: 3710283538
I1101 16:48:02.885977    1316 linux.go:119] bytesTransStr for wlan0: 99351616
I1101 16:48:02.886032    1316 linux.go:113] found device lo
I1101 16:48:02.886075    1316 linux.go:118] bytesRecvStr for lo: 107348
I1101 16:48:02.886117    1316 linux.go:119] bytesTransStr for lo: 107348
I1101 16:48:02.886168    1316 linux.go:65] updating state for eth0
I1101 16:48:02.886209    1316 linux.go:65] updating state for eth1
I1101 16:48:02.886249    1316 linux.go:65] updating state for eth2
I1101 16:48:02.886290    1316 linux.go:65] updating state for wlan0
I1101 16:48:02.886328    1316 linux.go:65] updating state for lo
I1101 16:48:02.886406    1316 table.go:58] rows = 40, tableLineCount = 2
I1101 16:48:02.886452    1316 table.go:69] tableLineCount = 2, rows-3 = 37
     439.92       66.38      63.87      432.88       0.37        0.42       0.00        0.00
```

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.

## Current Version

```
$ bndstat --version
bndstat v0.4.1
Rob Kingsbury
https://github.com/robkingsbury/bndstat
```
