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

### Examples

In this example, running on a raspberry pi that I am using as a router, eth0 is the hardline connection to wifi router, eth1 is my
primary internet provider, eth2 is my backup internet line and wlan0 is a connection to the wifi network:

```
$ bndstat 3 5
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
     783.14       26.65      23.85      788.56       0.44        0.53       0.00        0.00
     761.04      310.35     304.79      767.65       0.55        0.67       0.00        0.00
     790.48       30.69      26.92      796.46       0.37        0.42       0.00        0.00
     782.78       43.22      38.94      788.66       0.25        0.28       0.71        0.00
     920.06       32.81      29.77      926.96       0.37        0.42       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      75.99      272.63       0.37        0.42
      55.81      851.11       0.75        0.81
      53.00      924.72       0.37        0.42
      46.27     1133.64       0.00        0.42
      96.43      238.93       0.37        0.00
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1102 08:33:03.090266    3668 bndstat.go:96] interval = 1.000000, count = 1
I1102 08:33:03.091310    3668 throughput.go:21] os is "linux"
I1102 08:33:03.091451    3668 throughput.go:33] running Reporter.Report() twice to prime the stats
I1102 08:33:03.091538    3668 throughput.go:35] prime 1
I1102 08:33:03.092071    3668 linux.go:113] found device eth0
I1102 08:33:03.092162    3668 linux.go:118] bytesRecvStr for eth0: 2997753976
I1102 08:33:03.092246    3668 linux.go:119] bytesTransStr for eth0: 1969773637
I1102 08:33:03.092428    3668 linux.go:113] found device eth1
I1102 08:33:03.092505    3668 linux.go:118] bytesRecvStr for eth1: 515476766062
I1102 08:33:03.092586    3668 linux.go:119] bytesTransStr for eth1: 302016791647
I1102 08:33:03.092673    3668 linux.go:113] found device eth2
I1102 08:33:03.092750    3668 linux.go:118] bytesRecvStr for eth2: 292719455
I1102 08:33:03.092827    3668 linux.go:119] bytesTransStr for eth2: 399910994
I1102 08:33:03.092917    3668 linux.go:113] found device wlan0
I1102 08:33:03.092994    3668 linux.go:118] bytesRecvStr for wlan0: 3796373862
I1102 08:33:03.093074    3668 linux.go:119] bytesTransStr for wlan0: 99353350
I1102 08:33:03.093162    3668 linux.go:113] found device lo
I1102 08:33:03.093236    3668 linux.go:118] bytesRecvStr for lo: 107348
I1102 08:33:03.093315    3668 linux.go:119] bytesTransStr for lo: 107348
I1102 08:33:03.093406    3668 linux.go:65] updating state for eth0
I1102 08:33:03.093489    3668 linux.go:65] updating state for eth1
I1102 08:33:03.093566    3668 linux.go:65] updating state for eth2
I1102 08:33:03.093645    3668 linux.go:65] updating state for wlan0
I1102 08:33:03.093720    3668 linux.go:65] updating state for lo
I1102 08:33:03.093829    3668 throughput.go:38] prime 2
I1102 08:33:03.094089    3668 linux.go:113] found device eth0
I1102 08:33:03.094188    3668 linux.go:118] bytesRecvStr for eth0: 2997753976
I1102 08:33:03.094479    3668 linux.go:119] bytesTransStr for eth0: 1969773637
I1102 08:33:03.094704    3668 linux.go:113] found device eth1
I1102 08:33:03.094820    3668 linux.go:118] bytesRecvStr for eth1: 515476766062
I1102 08:33:03.094902    3668 linux.go:119] bytesTransStr for eth1: 302016791647
I1102 08:33:03.095010    3668 linux.go:113] found device eth2
I1102 08:33:03.095088    3668 linux.go:118] bytesRecvStr for eth2: 292719455
I1102 08:33:03.095168    3668 linux.go:119] bytesTransStr for eth2: 399910994
I1102 08:33:03.095257    3668 linux.go:113] found device wlan0
I1102 08:33:03.095333    3668 linux.go:118] bytesRecvStr for wlan0: 3796373862
I1102 08:33:03.095412    3668 linux.go:119] bytesTransStr for wlan0: 99353350
I1102 08:33:03.095497    3668 linux.go:113] found device lo
I1102 08:33:03.095630    3668 linux.go:118] bytesRecvStr for lo: 107348
I1102 08:33:03.095708    3668 linux.go:119] bytesTransStr for lo: 107348
I1102 08:33:03.095801    3668 linux.go:65] updating state for eth0
I1102 08:33:03.095877    3668 linux.go:65] updating state for eth1
I1102 08:33:03.095953    3668 linux.go:65] updating state for eth2
I1102 08:33:03.096028    3668 linux.go:65] updating state for wlan0
I1102 08:33:03.096102    3668 linux.go:65] updating state for lo
I1102 08:33:03.096416    3668 linux.go:113] found device eth0
I1102 08:33:03.096500    3668 linux.go:118] bytesRecvStr for eth0: 2997753976
I1102 08:33:03.096578    3668 linux.go:119] bytesTransStr for eth0: 1969773637
I1102 08:33:03.096666    3668 linux.go:113] found device eth1
I1102 08:33:03.096741    3668 linux.go:118] bytesRecvStr for eth1: 515476766062
I1102 08:33:03.096818    3668 linux.go:119] bytesTransStr for eth1: 302016791647
I1102 08:33:03.096903    3668 linux.go:113] found device eth2
I1102 08:33:03.096976    3668 linux.go:118] bytesRecvStr for eth2: 292719455
I1102 08:33:03.097055    3668 linux.go:119] bytesTransStr for eth2: 399910994
I1102 08:33:03.097143    3668 linux.go:113] found device wlan0
I1102 08:33:03.097220    3668 linux.go:118] bytesRecvStr for wlan0: 3796374251
I1102 08:33:03.097298    3668 linux.go:119] bytesTransStr for wlan0: 99353350
I1102 08:33:03.097383    3668 linux.go:113] found device lo
I1102 08:33:03.097486    3668 linux.go:118] bytesRecvStr for lo: 107348
I1102 08:33:03.097563    3668 linux.go:119] bytesTransStr for lo: 107348
I1102 08:33:03.097653    3668 linux.go:65] updating state for eth0
I1102 08:33:03.097728    3668 linux.go:65] updating state for eth1
I1102 08:33:03.097803    3668 linux.go:65] updating state for eth2
I1102 08:33:03.097942    3668 linux.go:65] updating state for wlan0
I1102 08:33:03.098022    3668 linux.go:65] updating state for lo
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1102 08:33:04.099402    3668 linux.go:113] found device eth0
I1102 08:33:04.099558    3668 linux.go:118] bytesRecvStr for eth0: 2998110319
I1102 08:33:04.099619    3668 linux.go:119] bytesTransStr for eth0: 1969803431
I1102 08:33:04.099675    3668 linux.go:113] found device eth1
I1102 08:33:04.099716    3668 linux.go:118] bytesRecvStr for eth1: 515476793480
I1102 08:33:04.099755    3668 linux.go:119] bytesTransStr for eth1: 302017150494
I1102 08:33:04.099803    3668 linux.go:113] found device eth2
I1102 08:33:04.099842    3668 linux.go:118] bytesRecvStr for eth2: 292719625
I1102 08:33:04.099880    3668 linux.go:119] bytesTransStr for eth2: 399911142
I1102 08:33:04.099928    3668 linux.go:113] found device wlan0
I1102 08:33:04.099967    3668 linux.go:118] bytesRecvStr for wlan0: 3796374640
I1102 08:33:04.100005    3668 linux.go:119] bytesTransStr for wlan0: 99353350
I1102 08:33:04.100051    3668 linux.go:113] found device lo
I1102 08:33:04.100089    3668 linux.go:118] bytesRecvStr for lo: 107348
I1102 08:33:04.100127    3668 linux.go:119] bytesTransStr for lo: 107348
I1102 08:33:04.100179    3668 linux.go:65] updating state for eth0
I1102 08:33:04.100218    3668 linux.go:65] updating state for eth1
I1102 08:33:04.100254    3668 linux.go:65] updating state for eth2
I1102 08:33:04.100290    3668 linux.go:65] updating state for wlan0
I1102 08:33:04.100326    3668 linux.go:65] updating state for lo
I1102 08:33:04.100400    3668 table.go:58] rows = 40, tableLineCount = 2
I1102 08:33:04.100444    3668 table.go:69] tableLineCount = 2, rows-3 = 37
    2776.91      232.18     213.66     2796.43       1.32        1.15       3.03        0.00
```

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
