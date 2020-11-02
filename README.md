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
    1446.57       79.22      71.29     1457.11       0.42        0.53      24.32        0.00
     605.04       37.15      33.64      610.34       0.37        0.42       0.71        0.00
     705.35      152.65     148.08      711.32       0.37        0.42       0.00        0.00
     504.74       22.83      20.27      509.53       0.25        0.28       0.00        0.00
    1029.10       47.39      44.24     1035.44       0.37        0.42       0.24        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      30.20     1719.60       0.37        0.42
      49.34     2223.48       0.75        0.81
      33.13      931.73       0.00        0.42
      26.73     1092.68       0.37        0.00
      27.33     1078.78       0.37        0.42
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1101 20:22:25.569184    1940 bndstat.go:96] interval = 1.000000, count = 1
I1101 20:22:25.569703    1940 throughput.go:21] os is "linux"
I1101 20:22:25.569766    1940 throughput.go:33] running Reporter.Report() twice to prime the stats
I1101 20:22:25.569873    1940 throughput.go:35] prime 1
I1101 20:22:25.570210    1940 linux.go:113] found device eth0
I1101 20:22:25.570280    1940 linux.go:118] bytesRecvStr for eth0: 2529840594
I1101 20:22:25.570348    1940 linux.go:119] bytesTransStr for eth0: 2415767617
I1101 20:22:25.570422    1940 linux.go:113] found device eth1
I1101 20:22:25.570482    1940 linux.go:118] bytesRecvStr for eth1: 511681534127
I1101 20:22:25.570546    1940 linux.go:119] bytesTransStr for eth1: 297212822734
I1101 20:22:25.570616    1940 linux.go:113] found device eth2
I1101 20:22:25.570680    1940 linux.go:118] bytesRecvStr for eth2: 290513709
I1101 20:22:25.570742    1940 linux.go:119] bytesTransStr for eth2: 397362250
I1101 20:22:25.570814    1940 linux.go:113] found device wlan0
I1101 20:22:25.570874    1940 linux.go:118] bytesRecvStr for wlan0: 3726428217
I1101 20:22:25.570937    1940 linux.go:119] bytesTransStr for wlan0: 99352316
I1101 20:22:25.571006    1940 linux.go:113] found device lo
I1101 20:22:25.571067    1940 linux.go:118] bytesRecvStr for lo: 107348
I1101 20:22:25.571129    1940 linux.go:119] bytesTransStr for lo: 107348
I1101 20:22:25.571202    1940 linux.go:65] updating state for eth0
I1101 20:22:25.571266    1940 linux.go:65] updating state for eth1
I1101 20:22:25.571328    1940 linux.go:65] updating state for eth2
I1101 20:22:25.571391    1940 linux.go:65] updating state for wlan0
I1101 20:22:25.571451    1940 linux.go:65] updating state for lo
I1101 20:22:25.571538    1940 throughput.go:38] prime 2
I1101 20:22:25.571745    1940 linux.go:113] found device eth0
I1101 20:22:25.571814    1940 linux.go:118] bytesRecvStr for eth0: 2529840594
I1101 20:22:25.571877    1940 linux.go:119] bytesTransStr for eth0: 2415767617
I1101 20:22:25.572033    1940 linux.go:113] found device eth1
I1101 20:22:25.572096    1940 linux.go:118] bytesRecvStr for eth1: 511681534127
I1101 20:22:25.572159    1940 linux.go:119] bytesTransStr for eth1: 297212822734
I1101 20:22:25.572232    1940 linux.go:113] found device eth2
I1101 20:22:25.572292    1940 linux.go:118] bytesRecvStr for eth2: 290513709
I1101 20:22:25.572355    1940 linux.go:119] bytesTransStr for eth2: 397362250
I1101 20:22:25.572425    1940 linux.go:113] found device wlan0
I1101 20:22:25.572512    1940 linux.go:118] bytesRecvStr for wlan0: 3726428217
I1101 20:22:25.572576    1940 linux.go:119] bytesTransStr for wlan0: 99352316
I1101 20:22:25.572644    1940 linux.go:113] found device lo
I1101 20:22:25.572704    1940 linux.go:118] bytesRecvStr for lo: 107348
I1101 20:22:25.572767    1940 linux.go:119] bytesTransStr for lo: 107348
I1101 20:22:25.572840    1940 linux.go:65] updating state for eth0
I1101 20:22:25.572901    1940 linux.go:65] updating state for eth1
I1101 20:22:25.572962    1940 linux.go:65] updating state for eth2
I1101 20:22:25.573022    1940 linux.go:65] updating state for wlan0
I1101 20:22:25.573082    1940 linux.go:65] updating state for lo
I1101 20:22:25.573339    1940 linux.go:113] found device eth0
I1101 20:22:25.573406    1940 linux.go:118] bytesRecvStr for eth0: 2529840594
I1101 20:22:25.573468    1940 linux.go:119] bytesTransStr for eth0: 2415767617
I1101 20:22:25.573562    1940 linux.go:113] found device eth1
I1101 20:22:25.573675    1940 linux.go:118] bytesRecvStr for eth1: 511681534127
I1101 20:22:25.573741    1940 linux.go:119] bytesTransStr for eth1: 297212822734
I1101 20:22:25.573812    1940 linux.go:113] found device eth2
I1101 20:22:25.573873    1940 linux.go:118] bytesRecvStr for eth2: 290513709
I1101 20:22:25.573936    1940 linux.go:119] bytesTransStr for eth2: 397362250
I1101 20:22:25.574007    1940 linux.go:113] found device wlan0
I1101 20:22:25.574066    1940 linux.go:118] bytesRecvStr for wlan0: 3726428217
I1101 20:22:25.574129    1940 linux.go:119] bytesTransStr for wlan0: 99352316
I1101 20:22:25.574197    1940 linux.go:113] found device lo
I1101 20:22:25.574257    1940 linux.go:118] bytesRecvStr for lo: 107348
I1101 20:22:25.574318    1940 linux.go:119] bytesTransStr for lo: 107348
I1101 20:22:25.574390    1940 linux.go:65] updating state for eth0
I1101 20:22:25.574450    1940 linux.go:65] updating state for eth1
I1101 20:22:25.574509    1940 linux.go:65] updating state for eth2
I1101 20:22:25.574568    1940 linux.go:65] updating state for wlan0
I1101 20:22:25.574628    1940 linux.go:65] updating state for lo
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1101 20:22:26.575771    1940 linux.go:113] found device eth0
I1101 20:22:26.575854    1940 linux.go:118] bytesRecvStr for eth0: 2529922030
I1101 20:22:26.575900    1940 linux.go:119] bytesTransStr for eth0: 2415770345
I1101 20:22:26.575958    1940 linux.go:113] found device eth1
I1101 20:22:26.576003    1940 linux.go:118] bytesRecvStr for eth1: 511681536611
I1101 20:22:26.576046    1940 linux.go:119] bytesTransStr for eth1: 297212904795
I1101 20:22:26.576095    1940 linux.go:113] found device eth2
I1101 20:22:26.576134    1940 linux.go:118] bytesRecvStr for eth2: 290513757
I1101 20:22:26.576172    1940 linux.go:119] bytesTransStr for eth2: 397362304
I1101 20:22:26.576221    1940 linux.go:113] found device wlan0
I1101 20:22:26.576260    1940 linux.go:118] bytesRecvStr for wlan0: 3726428217
I1101 20:22:26.576298    1940 linux.go:119] bytesTransStr for wlan0: 99352316
I1101 20:22:26.576346    1940 linux.go:113] found device lo
I1101 20:22:26.576383    1940 linux.go:118] bytesRecvStr for lo: 107348
I1101 20:22:26.576421    1940 linux.go:119] bytesTransStr for lo: 107348
I1101 20:22:26.576472    1940 linux.go:65] updating state for eth0
I1101 20:22:26.576511    1940 linux.go:65] updating state for eth1
I1101 20:22:26.576547    1940 linux.go:65] updating state for eth2
I1101 20:22:26.576583    1940 linux.go:65] updating state for wlan0
I1101 20:22:26.576619    1940 linux.go:65] updating state for lo
I1101 20:22:26.576728    1940 table.go:58] rows = 40, tableLineCount = 2
I1101 20:22:26.576772    1940 table.go:69] tableLineCount = 2, rows-3 = 37
     634.90       21.27      19.37      639.77       0.37        0.42       0.00        0.00
```

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.

## Current Version

```
$ bndstat --version
bndstat v0.4.2
Rob Kingsbury
https://github.com/robkingsbury/bndstat
```
