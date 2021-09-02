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
bndstat v0.5.3
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 136373c (v0.5.3)
Compiled: Wed  1 Sep 22:51:27 PDT 2021
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
     275.46       29.80      86.32      279.98       1.87        1.82       0.00        0.00
     388.80       17.51      69.88      393.57       1.75        1.82       0.00        0.00
     280.11       16.34      67.23      284.56       1.87        1.69       0.00        0.00
     395.03       20.09      69.06      400.00       1.87        1.82       0.00        0.00
     251.98       19.94      74.74      256.25       1.75        1.69       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      76.81      541.30       2.25        2.08
      71.02      159.18       1.87        1.69
      64.51      450.93       1.87        1.69
      59.69      168.32       1.87        1.69
     152.05      668.49       1.87        1.69
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0901 22:51:49.022472    9726 bndstat.go:101] interval = 1.000000, count = 1
I0901 22:51:49.022998    9726 throughput.go:21] os is "linux"
I0901 22:51:49.023042    9726 throughput.go:33] running Reporter.Report() twice to prime the stats
I0901 22:51:49.023078    9726 throughput.go:35] prime 1
I0901 22:51:49.023451    9726 linux.go:230] found device eth0
I0901 22:51:49.023494    9726 linux.go:235] bytesRecvStr for eth0: 1132901368
I0901 22:51:49.023532    9726 linux.go:236] bytesTransStr for eth0: 1819870690
I0901 22:51:49.023575    9726 linux.go:230] found device eth1
I0901 22:51:49.023609    9726 linux.go:235] bytesRecvStr for eth1: 28883458404
I0901 22:51:49.023643    9726 linux.go:236] bytesTransStr for eth1: 10010348898
I0901 22:51:49.023684    9726 linux.go:230] found device eth2
I0901 22:51:49.023717    9726 linux.go:235] bytesRecvStr for eth2: 28966784
I0901 22:51:49.023750    9726 linux.go:236] bytesTransStr for eth2: 31793216
I0901 22:51:49.023792    9726 linux.go:230] found device wlan0
I0901 22:51:49.023826    9726 linux.go:235] bytesRecvStr for wlan0: 0
I0901 22:51:49.023860    9726 linux.go:236] bytesTransStr for wlan0: 0
I0901 22:51:49.023900    9726 linux.go:230] found device lo
I0901 22:51:49.023934    9726 linux.go:235] bytesRecvStr for lo: 8480
I0901 22:51:49.023967    9726 linux.go:236] bytesTransStr for lo: 8480
I0901 22:51:49.024033    9726 linux.go:123] updating state for eth0
I0901 22:51:49.024070    9726 linux.go:123] updating state for eth1
I0901 22:51:49.024104    9726 linux.go:123] updating state for eth2
I0901 22:51:49.024137    9726 linux.go:123] updating state for wlan0
I0901 22:51:49.024169    9726 linux.go:123] updating state for lo
I0901 22:51:49.024219    9726 linux.go:182] eth0: max counter seen = 1819870690, max counter guess = 4294967296
I0901 22:51:49.024285    9726 linux.go:210] eth0: in=0.0010 kbps, out=0.0015 kbps
I0901 22:51:49.024360    9726 linux.go:182] eth1: max counter seen = 28883458404, max counter guess = 18446744069414584320
I0901 22:51:49.024429    9726 linux.go:210] eth1: in=0.0245 kbps, out=0.0085 kbps
I0901 22:51:49.024501    9726 linux.go:182] eth2: max counter seen = 31793216, max counter guess = 4294967296
I0901 22:51:49.024554    9726 linux.go:210] eth2: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.024623    9726 linux.go:182] lo: max counter seen = 8480, max counter guess = 4294967296
I0901 22:51:49.024674    9726 linux.go:210] lo: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.024740    9726 linux.go:182] wlan0: max counter seen = 0, max counter guess = 4294967296
I0901 22:51:49.024791    9726 linux.go:210] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.024867    9726 throughput.go:38] prime 2
I0901 22:51:49.025071    9726 linux.go:230] found device eth0
I0901 22:51:49.025110    9726 linux.go:235] bytesRecvStr for eth0: 1132901368
I0901 22:51:49.025146    9726 linux.go:236] bytesTransStr for eth0: 1819870690
I0901 22:51:49.025190    9726 linux.go:230] found device eth1
I0901 22:51:49.025224    9726 linux.go:235] bytesRecvStr for eth1: 28883458404
I0901 22:51:49.025258    9726 linux.go:236] bytesTransStr for eth1: 10010348898
I0901 22:51:49.025298    9726 linux.go:230] found device eth2
I0901 22:51:49.025331    9726 linux.go:235] bytesRecvStr for eth2: 28966784
I0901 22:51:49.025365    9726 linux.go:236] bytesTransStr for eth2: 31793216
I0901 22:51:49.025407    9726 linux.go:230] found device wlan0
I0901 22:51:49.025441    9726 linux.go:235] bytesRecvStr for wlan0: 0
I0901 22:51:49.025475    9726 linux.go:236] bytesTransStr for wlan0: 0
I0901 22:51:49.025516    9726 linux.go:230] found device lo
I0901 22:51:49.025582    9726 linux.go:235] bytesRecvStr for lo: 8480
I0901 22:51:49.025620    9726 linux.go:236] bytesTransStr for lo: 8480
I0901 22:51:49.025664    9726 linux.go:123] updating state for eth0
I0901 22:51:49.025696    9726 linux.go:123] updating state for eth1
I0901 22:51:49.025728    9726 linux.go:123] updating state for eth2
I0901 22:51:49.025759    9726 linux.go:123] updating state for wlan0
I0901 22:51:49.025791    9726 linux.go:123] updating state for lo
I0901 22:51:49.025833    9726 linux.go:182] eth0: max counter seen = 1819870690, max counter guess = 4294967296
I0901 22:51:49.025882    9726 linux.go:210] eth0: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.025925    9726 linux.go:182] eth1: max counter seen = 28883458404, max counter guess = 18446744069414584320
I0901 22:51:49.025985    9726 linux.go:210] eth1: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.026026    9726 linux.go:182] eth2: max counter seen = 31793216, max counter guess = 4294967296
I0901 22:51:49.026070    9726 linux.go:210] eth2: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.026112    9726 linux.go:182] lo: max counter seen = 8480, max counter guess = 4294967296
I0901 22:51:49.026154    9726 linux.go:210] lo: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.026194    9726 linux.go:182] wlan0: max counter seen = 0, max counter guess = 4294967296
I0901 22:51:49.026236    9726 linux.go:210] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.026503    9726 linux.go:230] found device eth0
I0901 22:51:49.026542    9726 linux.go:235] bytesRecvStr for eth0: 1132901368
I0901 22:51:49.026577    9726 linux.go:236] bytesTransStr for eth0: 1819870690
I0901 22:51:49.026622    9726 linux.go:230] found device eth1
I0901 22:51:49.026749    9726 linux.go:235] bytesRecvStr for eth1: 28883458452
I0901 22:51:49.026785    9726 linux.go:236] bytesTransStr for eth1: 10010348898
I0901 22:51:49.026825    9726 linux.go:230] found device eth2
I0901 22:51:49.026858    9726 linux.go:235] bytesRecvStr for eth2: 28966784
I0901 22:51:49.026892    9726 linux.go:236] bytesTransStr for eth2: 31793216
I0901 22:51:49.026934    9726 linux.go:230] found device wlan0
I0901 22:51:49.026968    9726 linux.go:235] bytesRecvStr for wlan0: 0
I0901 22:51:49.027002    9726 linux.go:236] bytesTransStr for wlan0: 0
I0901 22:51:49.027042    9726 linux.go:230] found device lo
I0901 22:51:49.027076    9726 linux.go:235] bytesRecvStr for lo: 8480
I0901 22:51:49.027109    9726 linux.go:236] bytesTransStr for lo: 8480
I0901 22:51:49.027153    9726 linux.go:123] updating state for eth0
I0901 22:51:49.027186    9726 linux.go:123] updating state for eth1
I0901 22:51:49.027218    9726 linux.go:123] updating state for eth2
I0901 22:51:49.027250    9726 linux.go:123] updating state for wlan0
I0901 22:51:49.027299    9726 linux.go:123] updating state for lo
I0901 22:51:49.027348    9726 linux.go:182] eth0: max counter seen = 1819870690, max counter guess = 4294967296
I0901 22:51:49.027400    9726 linux.go:210] eth0: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.027445    9726 linux.go:182] eth1: max counter seen = 28883458452, max counter guess = 18446744069414584320
I0901 22:51:49.027505    9726 linux.go:210] eth1: in=251.7377 kbps, out=0.0000 kbps
I0901 22:51:49.027560    9726 linux.go:182] eth2: max counter seen = 31793216, max counter guess = 4294967296
I0901 22:51:49.027606    9726 linux.go:210] eth2: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.027649    9726 linux.go:182] lo: max counter seen = 8480, max counter guess = 4294967296
I0901 22:51:49.027692    9726 linux.go:210] lo: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:49.027734    9726 linux.go:182] wlan0: max counter seen = 0, max counter guess = 4294967296
I0901 22:51:49.027776    9726 linux.go:210] wlan0: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I0901 22:51:50.028940    9726 linux.go:230] found device eth0
I0901 22:51:50.029096    9726 linux.go:235] bytesRecvStr for eth0: 1132924991
I0901 22:51:50.029217    9726 linux.go:236] bytesTransStr for eth0: 1820022573
I0901 22:51:50.029278    9726 linux.go:230] found device eth1
I0901 22:51:50.029326    9726 linux.go:235] bytesRecvStr for eth1: 28883616347
I0901 22:51:50.029370    9726 linux.go:236] bytesTransStr for eth1: 10010373017
I0901 22:51:50.029423    9726 linux.go:230] found device eth2
I0901 22:51:50.029468    9726 linux.go:235] bytesRecvStr for eth2: 28967072
I0901 22:51:50.029512    9726 linux.go:236] bytesTransStr for eth2: 31793482
I0901 22:51:50.029618    9726 linux.go:230] found device wlan0
I0901 22:51:50.029664    9726 linux.go:235] bytesRecvStr for wlan0: 0
I0901 22:51:50.029707    9726 linux.go:236] bytesTransStr for wlan0: 0
I0901 22:51:50.029759    9726 linux.go:230] found device lo
I0901 22:51:50.029803    9726 linux.go:235] bytesRecvStr for lo: 8480
I0901 22:51:50.029846    9726 linux.go:236] bytesTransStr for lo: 8480
I0901 22:51:50.029906    9726 linux.go:123] updating state for eth0
I0901 22:51:50.029953    9726 linux.go:123] updating state for eth1
I0901 22:51:50.029995    9726 linux.go:123] updating state for eth2
I0901 22:51:50.030035    9726 linux.go:123] updating state for wlan0
I0901 22:51:50.030074    9726 linux.go:123] updating state for lo
I0901 22:51:50.030140    9726 linux.go:182] eth0: max counter seen = 1820022573, max counter guess = 4294967296
I0901 22:51:50.030214    9726 linux.go:210] eth0: in=184.0491 kbps, out=1183.3350 kbps
I0901 22:51:50.030312    9726 linux.go:182] eth1: max counter seen = 28883616347, max counter guess = 18446744069414584320
I0901 22:51:50.030392    9726 linux.go:210] eth1: in=1230.1751 kbps, out=187.9134 kbps
I0901 22:51:50.030477    9726 linux.go:182] eth2: max counter seen = 31793482, max counter guess = 4294967296
I0901 22:51:50.030535    9726 linux.go:210] eth2: in=2.2438 kbps, out=2.0724 kbps
I0901 22:51:50.030619    9726 linux.go:182] lo: max counter seen = 8480, max counter guess = 4294967296
I0901 22:51:50.030675    9726 linux.go:210] lo: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:50.030728    9726 linux.go:182] wlan0: max counter seen = 0, max counter guess = 4294967296
I0901 22:51:50.030782    9726 linux.go:210] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0901 22:51:50.030874    9726 table.go:58] rows = 40, tableLineCount = 2
I0901 22:51:50.030922    9726 table.go:69] tableLineCount = 2, rows-3 = 37
     184.05     1183.34    1230.18      187.91       2.24        2.07       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
