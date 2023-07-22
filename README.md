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
bndstat v0.5.9
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 106c4fd (v0.5.9)
Compiled: Sat 22 Jul 11:36:22 PDT 2023
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
    1898.65       64.00     140.86     1924.18       2.18        2.51       0.00        0.00
    1665.48     3199.05    3251.91     1693.75       1.50        1.69       0.00        0.00
    2135.49       83.95     157.80     2163.08       2.01        2.26       0.00        0.00
    1995.03       68.03     146.31     2021.34       2.23        1.91       0.00        0.00
    1791.22     1323.67    1383.13     1820.00       1.82        2.04       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
     588.47     1343.91       1.50        1.69
    1422.22     4156.69       2.07        2.36
    2738.94      741.38       2.45        2.75
    1002.69      785.96       1.50        1.69
     185.65     3729.78       1.50        1.69
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0722 11:36:43.524597     798 bndstat.go:102] interval = 1.000000, count = 1
I0722 11:36:43.525144     798 throughput.go:21] os is "linux"
I0722 11:36:43.525193     798 throughput.go:33] running Reporter.Report() twice to prime the stats
I0722 11:36:43.525235     798 throughput.go:35] prime 1
I0722 11:36:43.525545     798 linux.go:231] found device eth0
I0722 11:36:43.525592     798 linux.go:236] bytesRecvStr for eth0: 1658081136
I0722 11:36:43.525636     798 linux.go:237] bytesTransStr for eth0: 562883912
I0722 11:36:43.525685     798 linux.go:231] found device eth1
I0722 11:36:43.525724     798 linux.go:236] bytesRecvStr for eth1: 273429268650
I0722 11:36:43.525766     798 linux.go:237] bytesTransStr for eth1: 261653585930
I0722 11:36:43.525813     798 linux.go:231] found device eth2
I0722 11:36:43.525853     798 linux.go:236] bytesRecvStr for eth2: 517064667
I0722 11:36:43.525895     798 linux.go:237] bytesTransStr for eth2: 368001300
I0722 11:36:43.525943     798 linux.go:231] found device wlan0
I0722 11:36:43.525984     798 linux.go:236] bytesRecvStr for wlan0: 0
I0722 11:36:43.526025     798 linux.go:237] bytesTransStr for wlan0: 0
I0722 11:36:43.526070     798 linux.go:231] found device lo
I0722 11:36:43.526111     798 linux.go:236] bytesRecvStr for lo: 4160962
I0722 11:36:43.526152     798 linux.go:237] bytesTransStr for lo: 4160962
I0722 11:36:43.526248     798 linux.go:124] updating state for eth0
I0722 11:36:43.526289     798 linux.go:124] updating state for eth1
I0722 11:36:43.526324     798 linux.go:124] updating state for eth2
I0722 11:36:43.526359     798 linux.go:124] updating state for wlan0
I0722 11:36:43.526393     798 linux.go:124] updating state for lo
I0722 11:36:43.526464     798 linux.go:183] eth0: max counter seen = 1658081136, max counter guess = 4294967296
I0722 11:36:43.526533     798 linux.go:211] eth0: in=0.0014 kbps, out=0.0005 kbps
I0722 11:36:43.526589     798 linux.go:183] eth1: max counter seen = 273429268650, max counter guess = 18446744069414584320
I0722 11:36:43.526647     798 linux.go:211] eth1: in=0.2316 kbps, out=0.2216 kbps
I0722 11:36:43.526695     798 linux.go:183] eth2: max counter seen = 517064667, max counter guess = 4294967296
I0722 11:36:43.526751     798 linux.go:211] eth2: in=0.0004 kbps, out=0.0003 kbps
I0722 11:36:43.526801     798 linux.go:183] lo: max counter seen = 4160962, max counter guess = 4294967296
I0722 11:36:43.526857     798 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.526906     798 linux.go:183] wlan0: max counter seen = 0, max counter guess = 4294967296
I0722 11:36:43.526961     798 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.527041     798 throughput.go:38] prime 2
I0722 11:36:43.527242     798 linux.go:231] found device eth0
I0722 11:36:43.527287     798 linux.go:236] bytesRecvStr for eth0: 1658081136
I0722 11:36:43.527329     798 linux.go:237] bytesTransStr for eth0: 562883912
I0722 11:36:43.527376     798 linux.go:231] found device eth1
I0722 11:36:43.527417     798 linux.go:236] bytesRecvStr for eth1: 273429268650
I0722 11:36:43.527458     798 linux.go:237] bytesTransStr for eth1: 261653585930
I0722 11:36:43.527505     798 linux.go:231] found device eth2
I0722 11:36:43.527545     798 linux.go:236] bytesRecvStr for eth2: 517064667
I0722 11:36:43.527586     798 linux.go:237] bytesTransStr for eth2: 368001300
I0722 11:36:43.527632     798 linux.go:231] found device wlan0
I0722 11:36:43.527673     798 linux.go:236] bytesRecvStr for wlan0: 0
I0722 11:36:43.527715     798 linux.go:237] bytesTransStr for wlan0: 0
I0722 11:36:43.527760     798 linux.go:231] found device lo
I0722 11:36:43.527800     798 linux.go:236] bytesRecvStr for lo: 4160962
I0722 11:36:43.527842     798 linux.go:237] bytesTransStr for lo: 4160962
I0722 11:36:43.527919     798 linux.go:124] updating state for eth0
I0722 11:36:43.527956     798 linux.go:124] updating state for eth1
I0722 11:36:43.527989     798 linux.go:124] updating state for eth2
I0722 11:36:43.528022     798 linux.go:124] updating state for wlan0
I0722 11:36:43.528055     798 linux.go:124] updating state for lo
I0722 11:36:43.528102     798 linux.go:183] eth0: max counter seen = 1658081136, max counter guess = 4294967296
I0722 11:36:43.528155     798 linux.go:211] eth0: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.528200     798 linux.go:183] eth1: max counter seen = 273429268650, max counter guess = 18446744069414584320
I0722 11:36:43.528287     798 linux.go:211] eth1: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.528334     798 linux.go:183] eth2: max counter seen = 517064667, max counter guess = 4294967296
I0722 11:36:43.528385     798 linux.go:211] eth2: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.528428     798 linux.go:183] lo: max counter seen = 4160962, max counter guess = 4294967296
I0722 11:36:43.528477     798 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.528519     798 linux.go:183] wlan0: max counter seen = 0, max counter guess = 4294967296
I0722 11:36:43.528568     798 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.528817     798 linux.go:231] found device eth0
I0722 11:36:43.528862     798 linux.go:236] bytesRecvStr for eth0: 1658081136
I0722 11:36:43.528904     798 linux.go:237] bytesTransStr for eth0: 562883912
I0722 11:36:43.528952     798 linux.go:231] found device eth1
I0722 11:36:43.528992     798 linux.go:236] bytesRecvStr for eth1: 273429268650
I0722 11:36:43.529033     798 linux.go:237] bytesTransStr for eth1: 261653585930
I0722 11:36:43.529078     798 linux.go:231] found device eth2
I0722 11:36:43.529119     798 linux.go:236] bytesRecvStr for eth2: 517064667
I0722 11:36:43.529159     798 linux.go:237] bytesTransStr for eth2: 368001300
I0722 11:36:43.529206     798 linux.go:231] found device wlan0
I0722 11:36:43.529246     798 linux.go:236] bytesRecvStr for wlan0: 0
I0722 11:36:43.529287     798 linux.go:237] bytesTransStr for wlan0: 0
I0722 11:36:43.529331     798 linux.go:231] found device lo
I0722 11:36:43.529371     798 linux.go:236] bytesRecvStr for lo: 4160962
I0722 11:36:43.529411     798 linux.go:237] bytesTransStr for lo: 4160962
I0722 11:36:43.529490     798 linux.go:124] updating state for eth0
I0722 11:36:43.529526     798 linux.go:124] updating state for eth1
I0722 11:36:43.529558     798 linux.go:124] updating state for eth2
I0722 11:36:43.529591     798 linux.go:124] updating state for wlan0
I0722 11:36:43.529622     798 linux.go:124] updating state for lo
I0722 11:36:43.529670     798 linux.go:183] eth0: max counter seen = 1658081136, max counter guess = 4294967296
I0722 11:36:43.529722     798 linux.go:211] eth0: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.529767     798 linux.go:183] eth1: max counter seen = 273429268650, max counter guess = 18446744069414584320
I0722 11:36:43.529817     798 linux.go:211] eth1: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.529859     798 linux.go:183] eth2: max counter seen = 517064667, max counter guess = 4294967296
I0722 11:36:43.529908     798 linux.go:211] eth2: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.529950     798 linux.go:183] lo: max counter seen = 4160962, max counter guess = 4294967296
I0722 11:36:43.529999     798 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:43.530058     798 linux.go:183] wlan0: max counter seen = 0, max counter guess = 4294967296
I0722 11:36:43.530108     798 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I0722 11:36:44.530589     798 linux.go:231] found device eth0
I0722 11:36:44.530625     798 linux.go:236] bytesRecvStr for eth0: 1658193693
I0722 11:36:44.530648     798 linux.go:237] bytesTransStr for eth0: 562906961
I0722 11:36:44.530675     798 linux.go:231] found device eth1
I0722 11:36:44.530696     798 linux.go:236] bytesRecvStr for eth1: 273429301683
I0722 11:36:44.530718     798 linux.go:237] bytesTransStr for eth1: 261653701403
I0722 11:36:44.530742     798 linux.go:231] found device eth2
I0722 11:36:44.530762     798 linux.go:236] bytesRecvStr for eth2: 517064859
I0722 11:36:44.530783     798 linux.go:237] bytesTransStr for eth2: 368001516
I0722 11:36:44.530807     798 linux.go:231] found device wlan0
I0722 11:36:44.530827     798 linux.go:236] bytesRecvStr for wlan0: 0
I0722 11:36:44.530848     798 linux.go:237] bytesTransStr for wlan0: 0
I0722 11:36:44.530870     798 linux.go:231] found device lo
I0722 11:36:44.530890     798 linux.go:236] bytesRecvStr for lo: 4160962
I0722 11:36:44.530911     798 linux.go:237] bytesTransStr for lo: 4160962
I0722 11:36:44.530960     798 linux.go:124] updating state for eth0
I0722 11:36:44.530979     798 linux.go:124] updating state for eth1
I0722 11:36:44.530996     798 linux.go:124] updating state for eth2
I0722 11:36:44.531013     798 linux.go:124] updating state for wlan0
I0722 11:36:44.531029     798 linux.go:124] updating state for lo
I0722 11:36:44.531056     798 linux.go:183] eth0: max counter seen = 1658193693, max counter guess = 4294967296
I0722 11:36:44.531090     798 linux.go:211] eth0: in=878.0554 kbps, out=179.8049 kbps
I0722 11:36:44.531120     798 linux.go:183] eth1: max counter seen = 273429301683, max counter guess = 18446744069414584320
I0722 11:36:44.531146     798 linux.go:211] eth1: in=257.6899 kbps, out=900.8031 kbps
I0722 11:36:44.531171     798 linux.go:183] eth2: max counter seen = 517064859, max counter guess = 4294967296
I0722 11:36:44.531196     798 linux.go:211] eth2: in=1.4978 kbps, out=1.6850 kbps
I0722 11:36:44.531220     798 linux.go:183] lo: max counter seen = 4160962, max counter guess = 4294967296
I0722 11:36:44.531245     798 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:44.531266     798 linux.go:183] wlan0: max counter seen = 0, max counter guess = 4294967296
I0722 11:36:44.531291     798 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0722 11:36:44.531330     798 table.go:58] rows = 40, tableLineCount = 2
I0722 11:36:44.531352     798 table.go:69] tableLineCount = 2, rows-3 = 37
     878.06      179.80     257.69      900.80       1.50        1.69       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
