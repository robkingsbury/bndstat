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
bndstat v0.5.10
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: d55749c (v0.5.10)
Compiled: Sat 20 Apr 13:20:13 PDT 2024
Build Host: bender
Go Build Version: go1.20.6
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
    2738.39      226.03     283.87     2756.59       1.82        2.04       0.00        0.00
    3002.54       79.40     141.36     3021.29       1.50        1.69       0.00        0.00
    3044.67      969.86    1020.48     3064.59       1.82        2.04       0.00        0.00
    2522.61       68.86     134.33     2539.56       1.82        2.04       0.00        0.00
    3171.64      106.54     167.67     3191.70       1.50        1.69       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      76.12      164.13       1.50        1.69
     163.76     2853.44       2.45        2.75
    2840.86     6321.62       1.50        1.68
      84.46      159.13       1.50        1.69
     192.26     2307.95       1.50        1.68
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0420 13:20:34.349868   26925 bndstat.go:102] interval = 1.000000, count = 1
I0420 13:20:34.350438   26925 throughput.go:21] os is "linux"
I0420 13:20:34.350484   26925 throughput.go:33] running Reporter.Report() twice to prime the stats
I0420 13:20:34.350523   26925 throughput.go:35] prime 1
I0420 13:20:34.350828   26925 linux.go:236] found device eth0
I0420 13:20:34.350875   26925 linux.go:241] bytesRecvStr for eth0: 148995683
I0420 13:20:34.350918   26925 linux.go:242] bytesTransStr for eth0: 4130769180
I0420 13:20:34.350966   26925 linux.go:236] found device eth1
I0420 13:20:34.351006   26925 linux.go:241] bytesRecvStr for eth1: 5280381566757
I0420 13:20:34.351048   26925 linux.go:242] bytesTransStr for eth1: 6976607489820
I0420 13:20:34.351093   26925 linux.go:236] found device eth2
I0420 13:20:34.351134   26925 linux.go:241] bytesRecvStr for eth2: 6972364709
I0420 13:20:34.351174   26925 linux.go:242] bytesTransStr for eth2: 7276427040
I0420 13:20:34.351221   26925 linux.go:236] found device wlan0
I0420 13:20:34.351261   26925 linux.go:241] bytesRecvStr for wlan0: 0
I0420 13:20:34.351302   26925 linux.go:242] bytesTransStr for wlan0: 0
I0420 13:20:34.351345   26925 linux.go:236] found device lo
I0420 13:20:34.351385   26925 linux.go:241] bytesRecvStr for lo: 12250788
I0420 13:20:34.351428   26925 linux.go:242] bytesTransStr for lo: 12250788
I0420 13:20:34.351517   26925 linux.go:129] updating state for eth0
I0420 13:20:34.351556   26925 linux.go:129] updating state for eth1
I0420 13:20:34.351590   26925 linux.go:129] updating state for eth2
I0420 13:20:34.351623   26925 linux.go:129] updating state for wlan0
I0420 13:20:34.351674   26925 linux.go:129] updating state for lo
I0420 13:20:34.351728   26925 linux.go:188] eth0: max counter seen = 4130769180, max counter guess = 4294967296
I0420 13:20:34.351796   26925 linux.go:216] eth0: in=0.0001 kbps, out=0.0035 kbps
I0420 13:20:34.351852   26925 linux.go:188] eth1: max counter seen = 6976607489820, max counter guess = 18446744069414584320
I0420 13:20:34.351909   26925 linux.go:216] eth1: in=4.4727 kbps, out=5.9094 kbps
I0420 13:20:34.351957   26925 linux.go:188] eth2: max counter seen = 7276427040, max counter guess = 18446744069414584320
I0420 13:20:34.352013   26925 linux.go:216] eth2: in=0.0059 kbps, out=0.0062 kbps
I0420 13:20:34.352062   26925 linux.go:188] lo: max counter seen = 12250788, max counter guess = 4294967296
I0420 13:20:34.352117   26925 linux.go:216] lo: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.352165   26925 linux.go:188] wlan0: max counter seen = 0, max counter guess = 4294967296
I0420 13:20:34.352220   26925 linux.go:216] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.352298   26925 throughput.go:38] prime 2
I0420 13:20:34.352493   26925 linux.go:236] found device eth0
I0420 13:20:34.352536   26925 linux.go:241] bytesRecvStr for eth0: 148995683
I0420 13:20:34.352579   26925 linux.go:242] bytesTransStr for eth0: 4130769180
I0420 13:20:34.352625   26925 linux.go:236] found device eth1
I0420 13:20:34.352666   26925 linux.go:241] bytesRecvStr for eth1: 5280381566757
I0420 13:20:34.352707   26925 linux.go:242] bytesTransStr for eth1: 6976607489820
I0420 13:20:34.352752   26925 linux.go:236] found device eth2
I0420 13:20:34.352793   26925 linux.go:241] bytesRecvStr for eth2: 6972364709
I0420 13:20:34.352833   26925 linux.go:242] bytesTransStr for eth2: 7276427040
I0420 13:20:34.352879   26925 linux.go:236] found device wlan0
I0420 13:20:34.352919   26925 linux.go:241] bytesRecvStr for wlan0: 0
I0420 13:20:34.352961   26925 linux.go:242] bytesTransStr for wlan0: 0
I0420 13:20:34.353005   26925 linux.go:236] found device lo
I0420 13:20:34.353045   26925 linux.go:241] bytesRecvStr for lo: 12250788
I0420 13:20:34.353086   26925 linux.go:242] bytesTransStr for lo: 12250788
I0420 13:20:34.353162   26925 linux.go:129] updating state for eth0
I0420 13:20:34.353199   26925 linux.go:129] updating state for eth1
I0420 13:20:34.353232   26925 linux.go:129] updating state for eth2
I0420 13:20:34.353263   26925 linux.go:129] updating state for wlan0
I0420 13:20:34.353295   26925 linux.go:129] updating state for lo
I0420 13:20:34.353340   26925 linux.go:188] eth0: max counter seen = 4130769180, max counter guess = 4294967296
I0420 13:20:34.353392   26925 linux.go:216] eth0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.353473   26925 linux.go:188] eth1: max counter seen = 6976607489820, max counter guess = 18446744069414584320
I0420 13:20:34.353528   26925 linux.go:216] eth1: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.353572   26925 linux.go:188] eth2: max counter seen = 7276427040, max counter guess = 18446744069414584320
I0420 13:20:34.353621   26925 linux.go:216] eth2: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.353663   26925 linux.go:188] lo: max counter seen = 12250788, max counter guess = 4294967296
I0420 13:20:34.353711   26925 linux.go:216] lo: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.353753   26925 linux.go:188] wlan0: max counter seen = 0, max counter guess = 4294967296
I0420 13:20:34.353801   26925 linux.go:216] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.354052   26925 linux.go:236] found device eth0
I0420 13:20:34.354095   26925 linux.go:241] bytesRecvStr for eth0: 148995683
I0420 13:20:34.354136   26925 linux.go:242] bytesTransStr for eth0: 4130769180
I0420 13:20:34.354183   26925 linux.go:236] found device eth1
I0420 13:20:34.354223   26925 linux.go:241] bytesRecvStr for eth1: 5280381566757
I0420 13:20:34.354264   26925 linux.go:242] bytesTransStr for eth1: 6976607489820
I0420 13:20:34.354310   26925 linux.go:236] found device eth2
I0420 13:20:34.354349   26925 linux.go:241] bytesRecvStr for eth2: 6972364709
I0420 13:20:34.354389   26925 linux.go:242] bytesTransStr for eth2: 7276427040
I0420 13:20:34.354436   26925 linux.go:236] found device wlan0
I0420 13:20:34.354477   26925 linux.go:241] bytesRecvStr for wlan0: 0
I0420 13:20:34.354518   26925 linux.go:242] bytesTransStr for wlan0: 0
I0420 13:20:34.354563   26925 linux.go:236] found device lo
I0420 13:20:34.354603   26925 linux.go:241] bytesRecvStr for lo: 12250788
I0420 13:20:34.354644   26925 linux.go:242] bytesTransStr for lo: 12250788
I0420 13:20:34.354729   26925 linux.go:129] updating state for eth0
I0420 13:20:34.354765   26925 linux.go:129] updating state for eth1
I0420 13:20:34.354797   26925 linux.go:129] updating state for eth2
I0420 13:20:34.354829   26925 linux.go:129] updating state for wlan0
I0420 13:20:34.354861   26925 linux.go:129] updating state for lo
I0420 13:20:34.354907   26925 linux.go:188] eth0: max counter seen = 4130769180, max counter guess = 4294967296
I0420 13:20:34.354959   26925 linux.go:216] eth0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.355003   26925 linux.go:188] eth1: max counter seen = 6976607489820, max counter guess = 18446744069414584320
I0420 13:20:34.355052   26925 linux.go:216] eth1: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.355094   26925 linux.go:188] eth2: max counter seen = 7276427040, max counter guess = 18446744069414584320
I0420 13:20:34.355143   26925 linux.go:216] eth2: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.355185   26925 linux.go:188] lo: max counter seen = 12250788, max counter guess = 4294967296
I0420 13:20:34.355250   26925 linux.go:216] lo: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:34.355294   26925 linux.go:188] wlan0: max counter seen = 0, max counter guess = 4294967296
I0420 13:20:34.355343   26925 linux.go:216] wlan0: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I0420 13:20:35.356482   26925 linux.go:236] found device eth0
I0420 13:20:35.356544   26925 linux.go:241] bytesRecvStr for eth0: 149790460
I0420 13:20:35.356590   26925 linux.go:242] bytesTransStr for eth0: 4130789475
I0420 13:20:35.356643   26925 linux.go:236] found device eth1
I0420 13:20:35.356685   26925 linux.go:241] bytesRecvStr for eth1: 5280381593142
I0420 13:20:35.356726   26925 linux.go:242] bytesTransStr for eth1: 6976608289219
I0420 13:20:35.356775   26925 linux.go:236] found device eth2
I0420 13:20:35.356815   26925 linux.go:241] bytesRecvStr for eth2: 6972364901
I0420 13:20:35.356856   26925 linux.go:242] bytesTransStr for eth2: 7276427256
I0420 13:20:35.356904   26925 linux.go:236] found device wlan0
I0420 13:20:35.356944   26925 linux.go:241] bytesRecvStr for wlan0: 0
I0420 13:20:35.356986   26925 linux.go:242] bytesTransStr for wlan0: 0
I0420 13:20:35.357029   26925 linux.go:236] found device lo
I0420 13:20:35.357068   26925 linux.go:241] bytesRecvStr for lo: 12250788
I0420 13:20:35.357109   26925 linux.go:242] bytesTransStr for lo: 12250788
I0420 13:20:35.357196   26925 linux.go:129] updating state for eth0
I0420 13:20:35.357249   26925 linux.go:129] updating state for eth1
I0420 13:20:35.357283   26925 linux.go:129] updating state for eth2
I0420 13:20:35.357316   26925 linux.go:129] updating state for wlan0
I0420 13:20:35.357348   26925 linux.go:129] updating state for lo
I0420 13:20:35.357400   26925 linux.go:188] eth0: max counter seen = 4130789475, max counter guess = 4294967296
I0420 13:20:35.357475   26925 linux.go:216] eth0: in=6193.9255 kbps, out=158.1648 kbps
I0420 13:20:35.357532   26925 linux.go:188] eth1: max counter seen = 6976608289219, max counter guess = 18446744069414584320
I0420 13:20:35.357584   26925 linux.go:216] eth1: in=205.6259 kbps, out=6229.9461 kbps
I0420 13:20:35.357632   26925 linux.go:188] eth2: max counter seen = 7276427256, max counter guess = 18446744069414584320
I0420 13:20:35.357683   26925 linux.go:216] eth2: in=1.4963 kbps, out=1.6834 kbps
I0420 13:20:35.357731   26925 linux.go:188] lo: max counter seen = 12250788, max counter guess = 4294967296
I0420 13:20:35.357780   26925 linux.go:216] lo: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:35.357822   26925 linux.go:188] wlan0: max counter seen = 0, max counter guess = 4294967296
I0420 13:20:35.357871   26925 linux.go:216] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:20:35.357947   26925 table.go:58] rows = 40, tableLineCount = 2
I0420 13:20:35.357990   26925 table.go:69] tableLineCount = 2, rows-3 = 37
    6193.93      158.16     205.63     6229.95       1.50        1.68       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
