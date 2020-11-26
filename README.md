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
Commit: 3c21440
Compiled: Thu 26 Nov 08:25:52 PST 2020
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
     577.80       22.33      19.28      581.76       0.57        0.67       0.00        0.00
    1089.75       38.27      34.12     1097.29       0.37        0.42       0.00        0.00
     900.10       28.84      25.39      906.68       0.55        0.67      24.19        0.00
     833.22       24.75      21.82      838.91       0.69        0.55       9.16        0.00
     352.31       34.63      32.15      356.36       0.25        0.53       1.16        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      56.12     1439.91       0.37        0.42
      99.87      925.28       0.95        1.89
      63.84      297.46       0.75        0.81
      18.62      389.83       0.37        0.42
      24.74     1806.89       0.37        0.42
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1126 08:26:13.924358    8985 bndstat.go:101] interval = 1.000000, count = 1
I1126 08:26:13.925065    8985 throughput.go:21] os is "linux"
I1126 08:26:13.925133    8985 throughput.go:33] running Reporter.Report() twice to prime the stats
I1126 08:26:13.925176    8985 throughput.go:35] prime 1
I1126 08:26:13.925571    8985 linux.go:113] found device eth0
I1126 08:26:13.925621    8985 linux.go:118] bytesRecvStr for eth0: 888115187
I1126 08:26:13.925667    8985 linux.go:119] bytesTransStr for eth0: 350690519
I1126 08:26:13.925719    8985 linux.go:113] found device eth1
I1126 08:26:13.925761    8985 linux.go:118] bytesRecvStr for eth1: 745938679616
I1126 08:26:13.925805    8985 linux.go:119] bytesTransStr for eth1: 559306659480
I1126 08:26:13.925856    8985 linux.go:113] found device eth2
I1126 08:26:13.925897    8985 linux.go:118] bytesRecvStr for eth2: 1315281289
I1126 08:26:13.925937    8985 linux.go:119] bytesTransStr for eth2: 1354594841
I1126 08:26:13.926012    8985 linux.go:113] found device wlan0
I1126 08:26:13.926055    8985 linux.go:118] bytesRecvStr for wlan0: 2931140069
I1126 08:26:13.926098    8985 linux.go:119] bytesTransStr for wlan0: 100350519
I1126 08:26:13.926147    8985 linux.go:113] found device lo
I1126 08:26:13.926189    8985 linux.go:118] bytesRecvStr for lo: 647133
I1126 08:26:13.926231    8985 linux.go:119] bytesTransStr for lo: 647133
I1126 08:26:13.926287    8985 linux.go:65] updating state for eth0
I1126 08:26:13.926363    8985 linux.go:65] updating state for eth1
I1126 08:26:13.926406    8985 linux.go:65] updating state for eth2
I1126 08:26:13.926446    8985 linux.go:65] updating state for wlan0
I1126 08:26:13.926487    8985 linux.go:65] updating state for lo
I1126 08:26:13.926558    8985 throughput.go:38] prime 2
I1126 08:26:13.926792    8985 linux.go:113] found device eth0
I1126 08:26:13.926839    8985 linux.go:118] bytesRecvStr for eth0: 888115571
I1126 08:26:13.926882    8985 linux.go:119] bytesTransStr for eth0: 350690519
I1126 08:26:13.926934    8985 linux.go:113] found device eth1
I1126 08:26:13.926976    8985 linux.go:118] bytesRecvStr for eth1: 745938679616
I1126 08:26:13.927019    8985 linux.go:119] bytesTransStr for eth1: 559306659872
I1126 08:26:13.927068    8985 linux.go:113] found device eth2
I1126 08:26:13.927110    8985 linux.go:118] bytesRecvStr for eth2: 1315281289
I1126 08:26:13.927152    8985 linux.go:119] bytesTransStr for eth2: 1354594841
I1126 08:26:13.927203    8985 linux.go:113] found device wlan0
I1126 08:26:13.927245    8985 linux.go:118] bytesRecvStr for wlan0: 2931140069
I1126 08:26:13.927287    8985 linux.go:119] bytesTransStr for wlan0: 100350519
I1126 08:26:13.927334    8985 linux.go:113] found device lo
I1126 08:26:13.927373    8985 linux.go:118] bytesRecvStr for lo: 647133
I1126 08:26:13.927416    8985 linux.go:119] bytesTransStr for lo: 647133
I1126 08:26:13.927469    8985 linux.go:65] updating state for eth0
I1126 08:26:13.927509    8985 linux.go:65] updating state for eth1
I1126 08:26:13.927548    8985 linux.go:65] updating state for eth2
I1126 08:26:13.927586    8985 linux.go:65] updating state for wlan0
I1126 08:26:13.927625    8985 linux.go:65] updating state for lo
I1126 08:26:13.927888    8985 linux.go:113] found device eth0
I1126 08:26:13.927934    8985 linux.go:118] bytesRecvStr for eth0: 888115571
I1126 08:26:13.927975    8985 linux.go:119] bytesTransStr for eth0: 350690519
I1126 08:26:13.928025    8985 linux.go:113] found device eth1
I1126 08:26:13.928067    8985 linux.go:118] bytesRecvStr for eth1: 745938679616
I1126 08:26:13.928110    8985 linux.go:119] bytesTransStr for eth1: 559306659872
I1126 08:26:13.928159    8985 linux.go:113] found device eth2
I1126 08:26:13.928198    8985 linux.go:118] bytesRecvStr for eth2: 1315281289
I1126 08:26:13.928240    8985 linux.go:119] bytesTransStr for eth2: 1354594841
I1126 08:26:13.928354    8985 linux.go:113] found device wlan0
I1126 08:26:13.928399    8985 linux.go:118] bytesRecvStr for wlan0: 2931140069
I1126 08:26:13.928441    8985 linux.go:119] bytesTransStr for wlan0: 100350519
I1126 08:26:13.928491    8985 linux.go:113] found device lo
I1126 08:26:13.928533    8985 linux.go:118] bytesRecvStr for lo: 647133
I1126 08:26:13.928574    8985 linux.go:119] bytesTransStr for lo: 647133
I1126 08:26:13.928627    8985 linux.go:65] updating state for eth0
I1126 08:26:13.928667    8985 linux.go:65] updating state for eth1
I1126 08:26:13.928705    8985 linux.go:65] updating state for eth2
I1126 08:26:13.928743    8985 linux.go:65] updating state for wlan0
I1126 08:26:13.928782    8985 linux.go:65] updating state for lo
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1126 08:26:14.929782    8985 linux.go:113] found device eth0
I1126 08:26:14.929857    8985 linux.go:118] bytesRecvStr for eth0: 888176511
I1126 08:26:14.929902    8985 linux.go:119] bytesTransStr for eth0: 350692922
I1126 08:26:14.930014    8985 linux.go:113] found device eth1
I1126 08:26:14.930058    8985 linux.go:118] bytesRecvStr for eth1: 745938681731
I1126 08:26:14.930102    8985 linux.go:119] bytesTransStr for eth1: 559306721248
I1126 08:26:14.930153    8985 linux.go:113] found device eth2
I1126 08:26:14.930195    8985 linux.go:118] bytesRecvStr for eth2: 1315281337
I1126 08:26:14.930238    8985 linux.go:119] bytesTransStr for eth2: 1354594895
I1126 08:26:14.930289    8985 linux.go:113] found device wlan0
I1126 08:26:14.930332    8985 linux.go:118] bytesRecvStr for wlan0: 2931140069
I1126 08:26:14.930375    8985 linux.go:119] bytesTransStr for wlan0: 100350519
I1126 08:26:14.930425    8985 linux.go:113] found device lo
I1126 08:26:14.930467    8985 linux.go:118] bytesRecvStr for lo: 647133
I1126 08:26:14.930509    8985 linux.go:119] bytesTransStr for lo: 647133
I1126 08:26:14.930565    8985 linux.go:65] updating state for eth0
I1126 08:26:14.930607    8985 linux.go:65] updating state for eth1
I1126 08:26:14.930647    8985 linux.go:65] updating state for eth2
I1126 08:26:14.930686    8985 linux.go:65] updating state for wlan0
I1126 08:26:14.930725    8985 linux.go:65] updating state for lo
I1126 08:26:14.930803    8985 table.go:58] rows = 40, tableLineCount = 2
I1126 08:26:14.930849    8985 table.go:69] tableLineCount = 2, rows-3 = 37
     475.17       18.74      16.49      478.57       0.37        0.42       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
