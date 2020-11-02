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
     551.56       25.73      23.13      556.30       0.25        0.28       0.00        0.00
     799.08       45.42      41.11      806.05       0.37        0.42       0.00        0.00
     884.33       34.72      31.41      890.44       0.50        0.55       0.00        0.00
     884.74       93.33      87.46      892.34       0.25        0.28       0.00        0.00
     738.71      220.99     213.41      745.71       0.37        0.42      55.90        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      54.65      951.55       0.37        0.42
      55.95      378.18       0.95        1.15
      41.96     1170.51       0.37        0.42
      25.45      701.27       0.95        1.15
      25.53     1019.50       0.37        0.42
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1101 16:40:14.089776    1179 bndstat.go:96] interval = 1.000000, count = 1
I1101 16:40:14.090405    1179 throughput.go:21] os is "linux"
I1101 16:40:14.090483    1179 throughput.go:33] running Reporter.Report() twice to prime the stats
I1101 16:40:14.090530    1179 throughput.go:35] prime 1
I1101 16:40:14.090907    1179 linux.go:113] found device eth0
I1101 16:40:14.090957    1179 linux.go:118] bytesRecvStr for eth0: 1071297025
I1101 16:40:14.091001    1179 linux.go:119] bytesTransStr for eth0: 531447962
I1101 16:40:14.091052    1179 linux.go:113] found device eth1
I1101 16:40:14.091092    1179 linux.go:118] bytesRecvStr for eth1: 509819537636
I1101 16:40:14.091137    1179 linux.go:119] bytesTransStr for eth1: 295740528348
I1101 16:40:14.091184    1179 linux.go:113] found device eth2
I1101 16:40:14.091226    1179 linux.go:118] bytesRecvStr for eth2: 289839447
I1101 16:40:14.091268    1179 linux.go:119] bytesTransStr for eth2: 396583282
I1101 16:40:14.091317    1179 linux.go:113] found device wlan0
I1101 16:40:14.091358    1179 linux.go:118] bytesRecvStr for wlan0: 3709533650
I1101 16:40:14.091399    1179 linux.go:119] bytesTransStr for wlan0: 99351616
I1101 16:40:14.091448    1179 linux.go:113] found device lo
I1101 16:40:14.091490    1179 linux.go:118] bytesRecvStr for lo: 107348
I1101 16:40:14.091531    1179 linux.go:119] bytesTransStr for lo: 107348
I1101 16:40:14.091586    1179 linux.go:65] updating state for eth0
I1101 16:40:14.091630    1179 linux.go:65] updating state for eth1
I1101 16:40:14.091670    1179 linux.go:65] updating state for eth2
I1101 16:40:14.091715    1179 linux.go:65] updating state for wlan0
I1101 16:40:14.091752    1179 linux.go:65] updating state for lo
I1101 16:40:14.091823    1179 throughput.go:38] prime 2
I1101 16:40:14.092046    1179 linux.go:113] found device eth0
I1101 16:40:14.092093    1179 linux.go:118] bytesRecvStr for eth0: 1071297025
I1101 16:40:14.092135    1179 linux.go:119] bytesTransStr for eth0: 531447962
I1101 16:40:14.092191    1179 linux.go:113] found device eth1
I1101 16:40:14.092233    1179 linux.go:118] bytesRecvStr for eth1: 509819537636
I1101 16:40:14.092274    1179 linux.go:119] bytesTransStr for eth1: 295740528348
I1101 16:40:14.092323    1179 linux.go:113] found device eth2
I1101 16:40:14.092363    1179 linux.go:118] bytesRecvStr for eth2: 289839447
I1101 16:40:14.092405    1179 linux.go:119] bytesTransStr for eth2: 396583282
I1101 16:40:14.092456    1179 linux.go:113] found device wlan0
I1101 16:40:14.092496    1179 linux.go:118] bytesRecvStr for wlan0: 3709533650
I1101 16:40:14.092537    1179 linux.go:119] bytesTransStr for wlan0: 99351616
I1101 16:40:14.092583    1179 linux.go:113] found device lo
I1101 16:40:14.092623    1179 linux.go:118] bytesRecvStr for lo: 107348
I1101 16:40:14.092667    1179 linux.go:119] bytesTransStr for lo: 107348
I1101 16:40:14.092717    1179 linux.go:65] updating state for eth0
I1101 16:40:14.092758    1179 linux.go:65] updating state for eth1
I1101 16:40:14.092796    1179 linux.go:65] updating state for eth2
I1101 16:40:14.092833    1179 linux.go:65] updating state for wlan0
I1101 16:40:14.092869    1179 linux.go:65] updating state for lo
I1101 16:40:14.093129    1179 linux.go:113] found device eth0
I1101 16:40:14.093175    1179 linux.go:118] bytesRecvStr for eth0: 1071297025
I1101 16:40:14.093216    1179 linux.go:119] bytesTransStr for eth0: 531447962
I1101 16:40:14.093264    1179 linux.go:113] found device eth1
I1101 16:40:14.093388    1179 linux.go:118] bytesRecvStr for eth1: 509819537636
I1101 16:40:14.093437    1179 linux.go:119] bytesTransStr for eth1: 295740528348
I1101 16:40:14.093489    1179 linux.go:113] found device eth2
I1101 16:40:14.093530    1179 linux.go:118] bytesRecvStr for eth2: 289839447
I1101 16:40:14.093572    1179 linux.go:119] bytesTransStr for eth2: 396583282
I1101 16:40:14.093620    1179 linux.go:113] found device wlan0
I1101 16:40:14.093661    1179 linux.go:118] bytesRecvStr for wlan0: 3709533650
I1101 16:40:14.093702    1179 linux.go:119] bytesTransStr for wlan0: 99351616
I1101 16:40:14.093752    1179 linux.go:113] found device lo
I1101 16:40:14.093792    1179 linux.go:118] bytesRecvStr for lo: 107348
I1101 16:40:14.093831    1179 linux.go:119] bytesTransStr for lo: 107348
I1101 16:40:14.093884    1179 linux.go:65] updating state for eth0
I1101 16:40:14.093924    1179 linux.go:65] updating state for eth1
I1101 16:40:14.093963    1179 linux.go:65] updating state for eth2
I1101 16:40:14.093999    1179 linux.go:65] updating state for wlan0
I1101 16:40:14.094036    1179 linux.go:65] updating state for lo
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1101 16:40:15.095108    1179 linux.go:113] found device eth0
I1101 16:40:15.095281    1179 linux.go:118] bytesRecvStr for eth0: 1071342963
I1101 16:40:15.095401    1179 linux.go:119] bytesTransStr for eth0: 531450928
I1101 16:40:15.095530    1179 linux.go:113] found device eth1
I1101 16:40:15.095638    1179 linux.go:118] bytesRecvStr for eth1: 509819540334
I1101 16:40:15.095749    1179 linux.go:119] bytesTransStr for eth1: 295740574706
I1101 16:40:15.095866    1179 linux.go:113] found device eth2
I1101 16:40:15.095971    1179 linux.go:118] bytesRecvStr for eth2: 289839447
I1101 16:40:15.096079    1179 linux.go:119] bytesTransStr for eth2: 396583336
I1101 16:40:15.096197    1179 linux.go:113] found device wlan0
I1101 16:40:15.096302    1179 linux.go:118] bytesRecvStr for wlan0: 3709533650
I1101 16:40:15.096409    1179 linux.go:119] bytesTransStr for wlan0: 99351616
I1101 16:40:15.096530    1179 linux.go:113] found device lo
I1101 16:40:15.096634    1179 linux.go:118] bytesRecvStr for lo: 107348
I1101 16:40:15.096742    1179 linux.go:119] bytesTransStr for lo: 107348
I1101 16:40:15.096862    1179 linux.go:65] updating state for eth0
I1101 16:40:15.096970    1179 linux.go:65] updating state for eth1
I1101 16:40:15.097074    1179 linux.go:65] updating state for eth2
I1101 16:40:15.097178    1179 linux.go:65] updating state for wlan0
I1101 16:40:15.097280    1179 linux.go:65] updating state for lo
I1101 16:40:15.097494    1179 table.go:58] rows = 40, tableLineCount = 2
I1101 16:40:15.097609    1179 table.go:69] tableLineCount = 2, rows-3 = 37
     357.83       23.10      21.02      361.10       0.00        0.42       0.00        0.00
```

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
