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
bndstat v0.5.7
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 82e76d4 (v0.5.7)
Compiled: Sat  1 Apr 11:02:32 PDT 2023
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
    1720.74       68.92     138.87     1733.88       1.62        1.82       0.00        0.00
    1702.46      631.33     690.51     1716.57       1.50        1.69       0.00        0.00
    1999.34      503.30     557.73     2015.37       1.62        1.82       0.00        0.00
    1839.41       76.50     130.67     1852.94       1.62        1.82       0.00        0.00
    1836.04      907.47     946.77     1853.24       1.50        1.69       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
    1754.40      602.86       1.50        1.69
    1141.27      697.44       2.25        2.47
     149.98     4042.52       1.50        1.69
     109.10     1642.84       1.50        1.69
      98.89      343.75       1.50        1.69
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0401 11:02:54.020706    2199 bndstat.go:101] interval = 1.000000, count = 1
I0401 11:02:54.021298    2199 throughput.go:21] os is "linux"
I0401 11:02:54.021400    2199 throughput.go:33] running Reporter.Report() twice to prime the stats
I0401 11:02:54.021469    2199 throughput.go:35] prime 1
I0401 11:02:54.021853    2199 linux.go:231] found device eth0
I0401 11:02:54.021924    2199 linux.go:236] bytesRecvStr for eth0: 2387516528
I0401 11:02:54.021992    2199 linux.go:237] bytesTransStr for eth0: 2229642562
I0401 11:02:54.022067    2199 linux.go:231] found device eth1
I0401 11:02:54.022128    2199 linux.go:236] bytesRecvStr for eth1: 3579729253295
I0401 11:02:54.022196    2199 linux.go:237] bytesTransStr for eth1: 3047658168024
I0401 11:02:54.022267    2199 linux.go:231] found device eth2
I0401 11:02:54.022329    2199 linux.go:236] bytesRecvStr for eth2: 2838466193
I0401 11:02:54.022392    2199 linux.go:237] bytesTransStr for eth2: 3167726779
I0401 11:02:54.022466    2199 linux.go:231] found device wlan0
I0401 11:02:54.022526    2199 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 11:02:54.022592    2199 linux.go:237] bytesTransStr for wlan0: 33550
I0401 11:02:54.022664    2199 linux.go:231] found device lo
I0401 11:02:54.022740    2199 linux.go:236] bytesRecvStr for lo: 349458
I0401 11:02:54.022805    2199 linux.go:237] bytesTransStr for lo: 349458
I0401 11:02:54.022921    2199 linux.go:124] updating state for eth0
I0401 11:02:54.023063    2199 linux.go:124] updating state for eth1
I0401 11:02:54.023125    2199 linux.go:124] updating state for eth2
I0401 11:02:54.023188    2199 linux.go:124] updating state for wlan0
I0401 11:02:54.023248    2199 linux.go:124] updating state for lo
I0401 11:02:54.023326    2199 linux.go:183] eth0: max counter seen = 2387516528, max counter guess = 4294967296
I0401 11:02:54.023421    2199 linux.go:211] eth0: in=0.0020 kbps, out=0.0019 kbps
I0401 11:02:54.023524    2199 linux.go:183] eth1: max counter seen = 3579729253295, max counter guess = 18446744069414584320
I0401 11:02:54.023738    2199 linux.go:211] eth1: in=3.0321 kbps, out=2.5815 kbps
I0401 11:02:54.023844    2199 linux.go:183] eth2: max counter seen = 3167726779, max counter guess = 4294967296
I0401 11:02:54.023930    2199 linux.go:211] eth2: in=0.0024 kbps, out=0.0027 kbps
I0401 11:02:54.024030    2199 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 11:02:54.024112    2199 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.024208    2199 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 11:02:54.024341    2199 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.024472    2199 throughput.go:38] prime 2
I0401 11:02:54.024699    2199 linux.go:231] found device eth0
I0401 11:02:54.024769    2199 linux.go:236] bytesRecvStr for eth0: 2387516528
I0401 11:02:54.024835    2199 linux.go:237] bytesTransStr for eth0: 2229642562
I0401 11:02:54.024909    2199 linux.go:231] found device eth1
I0401 11:02:54.024970    2199 linux.go:236] bytesRecvStr for eth1: 3579729253391
I0401 11:02:54.025035    2199 linux.go:237] bytesTransStr for eth1: 3047658168024
I0401 11:02:54.025106    2199 linux.go:231] found device eth2
I0401 11:02:54.025179    2199 linux.go:236] bytesRecvStr for eth2: 2838466193
I0401 11:02:54.025243    2199 linux.go:237] bytesTransStr for eth2: 3167726779
I0401 11:02:54.025315    2199 linux.go:231] found device wlan0
I0401 11:02:54.025377    2199 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 11:02:54.025474    2199 linux.go:237] bytesTransStr for wlan0: 33550
I0401 11:02:54.025597    2199 linux.go:231] found device lo
I0401 11:02:54.025658    2199 linux.go:236] bytesRecvStr for lo: 349458
I0401 11:02:54.025723    2199 linux.go:237] bytesTransStr for lo: 349458
I0401 11:02:54.025839    2199 linux.go:124] updating state for eth0
I0401 11:02:54.025906    2199 linux.go:124] updating state for eth1
I0401 11:02:54.025967    2199 linux.go:124] updating state for eth2
I0401 11:02:54.026028    2199 linux.go:124] updating state for wlan0
I0401 11:02:54.026088    2199 linux.go:124] updating state for lo
I0401 11:02:54.026164    2199 linux.go:183] eth0: max counter seen = 2387516528, max counter guess = 4294967296
I0401 11:02:54.026251    2199 linux.go:211] eth0: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.026326    2199 linux.go:183] eth1: max counter seen = 3579729253391, max counter guess = 18446744069414584320
I0401 11:02:54.026422    2199 linux.go:211] eth1: in=257.0110 kbps, out=0.0000 kbps
I0401 11:02:54.026507    2199 linux.go:183] eth2: max counter seen = 3167726779, max counter guess = 4294967296
I0401 11:02:54.026582    2199 linux.go:211] eth2: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.026652    2199 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 11:02:54.026727    2199 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.026799    2199 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 11:02:54.026873    2199 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.027168    2199 linux.go:231] found device eth0
I0401 11:02:54.027234    2199 linux.go:236] bytesRecvStr for eth0: 2387516528
I0401 11:02:54.027299    2199 linux.go:237] bytesTransStr for eth0: 2229642562
I0401 11:02:54.027370    2199 linux.go:231] found device eth1
I0401 11:02:54.027431    2199 linux.go:236] bytesRecvStr for eth1: 3579729253391
I0401 11:02:54.027494    2199 linux.go:237] bytesTransStr for eth1: 3047658168024
I0401 11:02:54.027564    2199 linux.go:231] found device eth2
I0401 11:02:54.027625    2199 linux.go:236] bytesRecvStr for eth2: 2838466193
I0401 11:02:54.027688    2199 linux.go:237] bytesTransStr for eth2: 3167726779
I0401 11:02:54.027760    2199 linux.go:231] found device wlan0
I0401 11:02:54.027824    2199 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 11:02:54.027888    2199 linux.go:237] bytesTransStr for wlan0: 33550
I0401 11:02:54.027957    2199 linux.go:231] found device lo
I0401 11:02:54.028017    2199 linux.go:236] bytesRecvStr for lo: 349458
I0401 11:02:54.028080    2199 linux.go:237] bytesTransStr for lo: 349458
I0401 11:02:54.028187    2199 linux.go:124] updating state for eth0
I0401 11:02:54.028250    2199 linux.go:124] updating state for eth1
I0401 11:02:54.028328    2199 linux.go:124] updating state for eth2
I0401 11:02:54.028391    2199 linux.go:124] updating state for wlan0
I0401 11:02:54.028449    2199 linux.go:124] updating state for lo
I0401 11:02:54.028522    2199 linux.go:183] eth0: max counter seen = 2387516528, max counter guess = 4294967296
I0401 11:02:54.028602    2199 linux.go:211] eth0: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.028675    2199 linux.go:183] eth1: max counter seen = 3579729253391, max counter guess = 18446744069414584320
I0401 11:02:54.028768    2199 linux.go:211] eth1: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.028839    2199 linux.go:183] eth2: max counter seen = 3167726779, max counter guess = 4294967296
I0401 11:02:54.028916    2199 linux.go:211] eth2: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.028989    2199 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 11:02:54.029064    2199 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:54.029161    2199 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 11:02:54.029235    2199 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I0401 11:02:55.030158    2199 linux.go:231] found device eth0
I0401 11:02:55.030294    2199 linux.go:236] bytesRecvStr for eth0: 2387995379
I0401 11:02:55.030367    2199 linux.go:237] bytesTransStr for eth0: 2229659942
I0401 11:02:55.030446    2199 linux.go:231] found device eth1
I0401 11:02:55.030511    2199 linux.go:236] bytesRecvStr for eth1: 3579729275777
I0401 11:02:55.030576    2199 linux.go:237] bytesTransStr for eth1: 3047658650062
I0401 11:02:55.030649    2199 linux.go:231] found device eth2
I0401 11:02:55.030710    2199 linux.go:236] bytesRecvStr for eth2: 2838466385
I0401 11:02:55.030774    2199 linux.go:237] bytesTransStr for eth2: 3167726995
I0401 11:02:55.030847    2199 linux.go:231] found device wlan0
I0401 11:02:55.030907    2199 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 11:02:55.030971    2199 linux.go:237] bytesTransStr for wlan0: 33550
I0401 11:02:55.031042    2199 linux.go:231] found device lo
I0401 11:02:55.031102    2199 linux.go:236] bytesRecvStr for lo: 349458
I0401 11:02:55.031165    2199 linux.go:237] bytesTransStr for lo: 349458
I0401 11:02:55.031281    2199 linux.go:124] updating state for eth0
I0401 11:02:55.031350    2199 linux.go:124] updating state for eth1
I0401 11:02:55.031412    2199 linux.go:124] updating state for eth2
I0401 11:02:55.031472    2199 linux.go:124] updating state for wlan0
I0401 11:02:55.031532    2199 linux.go:124] updating state for lo
I0401 11:02:55.031611    2199 linux.go:183] eth0: max counter seen = 2387995379, max counter guess = 4294967296
I0401 11:02:55.031732    2199 linux.go:211] eth0: in=3729.4865 kbps, out=135.3625 kbps
I0401 11:02:55.031841    2199 linux.go:183] eth1: max counter seen = 3579729275777, max counter guess = 18446744069414584320
I0401 11:02:55.031938    2199 linux.go:211] eth1: in=174.3513 kbps, out=3754.3081 kbps
I0401 11:02:55.032037    2199 linux.go:183] eth2: max counter seen = 3167726995, max counter guess = 4294967296
I0401 11:02:55.032114    2199 linux.go:211] eth2: in=1.4954 kbps, out=1.6823 kbps
I0401 11:02:55.032212    2199 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 11:02:55.032286    2199 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:55.032361    2199 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 11:02:55.032436    2199 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 11:02:55.032537    2199 table.go:58] rows = 40, tableLineCount = 2
I0401 11:02:55.032604    2199 table.go:69] tableLineCount = 2, rows-3 = 37
    3729.49      135.36     174.35     3754.31       1.50        1.68       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
