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
bndstat v0.4.3
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 405be28
Compiled: Mon 23 Nov 13:32:08 PST 2020
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
     993.30      734.46     722.59     1001.55       0.37        0.42       2.27        0.00
     766.20      889.78     876.53      774.18       0.37        0.42       3.03        0.00
     861.49       65.51      60.55      868.23       0.25        0.42       1.52        0.00
     982.19       25.99      22.82      988.49       0.57        0.53       3.03        0.00
    1119.13       96.04      42.42     1122.81       0.55        0.67       2.27        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
    3194.33     1792.26       0.37        0.42
      23.52     1114.54       0.37        0.42
      14.67      299.59       0.37        0.42
      22.92      924.23       0.37        0.42
     115.38      545.23       0.37        0.42
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1123 13:32:29.315711     602 bndstat.go:100] interval = 1.000000, count = 1
I1123 13:32:29.316509     602 throughput.go:21] os is "linux"
I1123 13:32:29.316597     602 throughput.go:33] running Reporter.Report() twice to prime the stats
I1123 13:32:29.316641     602 throughput.go:35] prime 1
I1123 13:32:29.317029     602 linux.go:113] found device eth0
I1123 13:32:29.317080     602 linux.go:118] bytesRecvStr for eth0: 1552042335
I1123 13:32:29.317126     602 linux.go:119] bytesTransStr for eth0: 2032297714
I1123 13:32:29.317178     602 linux.go:113] found device eth1
I1123 13:32:29.317219     602 linux.go:118] bytesRecvStr for eth1: 726433632295
I1123 13:32:29.317262     602 linux.go:119] bytesTransStr for eth1: 533974922348
I1123 13:32:29.317313     602 linux.go:113] found device eth2
I1123 13:32:29.317355     602 linux.go:118] bytesRecvStr for eth2: 1302248895
I1123 13:32:29.317397     602 linux.go:119] bytesTransStr for eth2: 1339591257
I1123 13:32:29.317449     602 linux.go:113] found device wlan0
I1123 13:32:29.317491     602 linux.go:118] bytesRecvStr for wlan0: 2480986621
I1123 13:32:29.317533     602 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:32:29.317583     602 linux.go:113] found device lo
I1123 13:32:29.317625     602 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:32:29.317667     602 linux.go:119] bytesTransStr for lo: 619794
I1123 13:32:29.317723     602 linux.go:65] updating state for eth0
I1123 13:32:29.317767     602 linux.go:65] updating state for eth1
I1123 13:32:29.317809     602 linux.go:65] updating state for eth2
I1123 13:32:29.317850     602 linux.go:65] updating state for wlan0
I1123 13:32:29.317890     602 linux.go:65] updating state for lo
I1123 13:32:29.317962     602 throughput.go:38] prime 2
I1123 13:32:29.318186     602 linux.go:113] found device eth0
I1123 13:32:29.318233     602 linux.go:118] bytesRecvStr for eth0: 1552042335
I1123 13:32:29.318276     602 linux.go:119] bytesTransStr for eth0: 2032297714
I1123 13:32:29.318327     602 linux.go:113] found device eth1
I1123 13:32:29.318369     602 linux.go:118] bytesRecvStr for eth1: 726433632295
I1123 13:32:29.318411     602 linux.go:119] bytesTransStr for eth1: 533974922348
I1123 13:32:29.318460     602 linux.go:113] found device eth2
I1123 13:32:29.318503     602 linux.go:118] bytesRecvStr for eth2: 1302248895
I1123 13:32:29.318545     602 linux.go:119] bytesTransStr for eth2: 1339591257
I1123 13:32:29.318596     602 linux.go:113] found device wlan0
I1123 13:32:29.318635     602 linux.go:118] bytesRecvStr for wlan0: 2480986621
I1123 13:32:29.318676     602 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:32:29.318726     602 linux.go:113] found device lo
I1123 13:32:29.318767     602 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:32:29.318810     602 linux.go:119] bytesTransStr for lo: 619794
I1123 13:32:29.318864     602 linux.go:65] updating state for eth0
I1123 13:32:29.318904     602 linux.go:65] updating state for eth1
I1123 13:32:29.318943     602 linux.go:65] updating state for eth2
I1123 13:32:29.318983     602 linux.go:65] updating state for wlan0
I1123 13:32:29.319021     602 linux.go:65] updating state for lo
I1123 13:32:29.319334     602 linux.go:113] found device eth0
I1123 13:32:29.319383     602 linux.go:118] bytesRecvStr for eth0: 1552042335
I1123 13:32:29.319426     602 linux.go:119] bytesTransStr for eth0: 2032297714
I1123 13:32:29.319476     602 linux.go:113] found device eth1
I1123 13:32:29.319518     602 linux.go:118] bytesRecvStr for eth1: 726433632295
I1123 13:32:29.319561     602 linux.go:119] bytesTransStr for eth1: 533974922348
I1123 13:32:29.319610     602 linux.go:113] found device eth2
I1123 13:32:29.319651     602 linux.go:118] bytesRecvStr for eth2: 1302248895
I1123 13:32:29.319694     602 linux.go:119] bytesTransStr for eth2: 1339591257
I1123 13:32:29.319745     602 linux.go:113] found device wlan0
I1123 13:32:29.319786     602 linux.go:118] bytesRecvStr for wlan0: 2480986621
I1123 13:32:29.319828     602 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:32:29.319941     602 linux.go:113] found device lo
I1123 13:32:29.319987     602 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:32:29.320028     602 linux.go:119] bytesTransStr for lo: 619794
I1123 13:32:29.320082     602 linux.go:65] updating state for eth0
I1123 13:32:29.320122     602 linux.go:65] updating state for eth1
I1123 13:32:29.320161     602 linux.go:65] updating state for eth2
I1123 13:32:29.320200     602 linux.go:65] updating state for wlan0
I1123 13:32:29.320238     602 linux.go:65] updating state for lo
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1123 13:32:30.320996     602 linux.go:113] found device eth0
I1123 13:32:30.321112     602 linux.go:118] bytesRecvStr for eth0: 1552188313
I1123 13:32:30.321203     602 linux.go:119] bytesTransStr for eth0: 2032301714
I1123 13:32:30.321299     602 linux.go:113] found device eth1
I1123 13:32:30.321375     602 linux.go:118] bytesRecvStr for eth1: 726433635733
I1123 13:32:30.321458     602 linux.go:119] bytesTransStr for eth1: 533975069260
I1123 13:32:30.321546     602 linux.go:113] found device eth2
I1123 13:32:30.321622     602 linux.go:118] bytesRecvStr for eth2: 1302248943
I1123 13:32:30.321701     602 linux.go:119] bytesTransStr for eth2: 1339591311
I1123 13:32:30.321791     602 linux.go:113] found device wlan0
I1123 13:32:30.321867     602 linux.go:118] bytesRecvStr for wlan0: 2480986912
I1123 13:32:30.321945     602 linux.go:119] bytesTransStr for wlan0: 100343615
I1123 13:32:30.322032     602 linux.go:113] found device lo
I1123 13:32:30.322106     602 linux.go:118] bytesRecvStr for lo: 619794
I1123 13:32:30.322184     602 linux.go:119] bytesTransStr for lo: 619794
I1123 13:32:30.322275     602 linux.go:65] updating state for eth0
I1123 13:32:30.322353     602 linux.go:65] updating state for eth1
I1123 13:32:30.322428     602 linux.go:65] updating state for eth2
I1123 13:32:30.322506     602 linux.go:65] updating state for wlan0
I1123 13:32:30.322581     602 linux.go:65] updating state for lo
I1123 13:32:30.322692     602 table.go:58] rows = 40, tableLineCount = 2
I1123 13:32:30.322776     602 table.go:69] tableLineCount = 2, rows-3 = 37
    1137.96       31.18      26.80     1145.24       0.37        0.42       2.27        0.00
```

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
