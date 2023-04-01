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
bndstat v0.5.4
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 610490e (v0.5.4)
Compiled: Sat  1 Apr 10:23:58 PDT 2023
Build Host: prober
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
              eth0                  wlan0     
         In         Out         In         Out
       1.86        0.27       0.00        0.00
       0.97        0.00       0.00        0.00
       1.15        0.00       0.00        0.00
       2.51        0.00       0.00        0.00
       0.94        0.00       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
Error: device, eth1, not found
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0401 10:24:15.001323   16051 bndstat.go:101] interval = 1.000000, count = 1
I0401 10:24:15.002094   16051 throughput.go:21] os is "linux"
I0401 10:24:15.002148   16051 throughput.go:33] running Reporter.Report() twice to prime the stats
I0401 10:24:15.002194   16051 throughput.go:35] prime 1
I0401 10:24:15.002547   16051 linux.go:231] found device wlan0
I0401 10:24:15.002603   16051 linux.go:236] bytesRecvStr for wlan0: 0
I0401 10:24:15.002650   16051 linux.go:237] bytesTransStr for wlan0: 0
I0401 10:24:15.002705   16051 linux.go:231] found device eth0
I0401 10:24:15.002750   16051 linux.go:236] bytesRecvStr for eth0: 2810951704
I0401 10:24:15.002795   16051 linux.go:237] bytesTransStr for eth0: 34188419
I0401 10:24:15.002849   16051 linux.go:231] found device lo
I0401 10:24:15.002892   16051 linux.go:236] bytesRecvStr for lo: 4116
I0401 10:24:15.002936   16051 linux.go:237] bytesTransStr for lo: 4116
I0401 10:24:15.003054   16051 linux.go:124] updating state for wlan0
I0401 10:24:15.003101   16051 linux.go:124] updating state for eth0
I0401 10:24:15.003142   16051 linux.go:124] updating state for lo
I0401 10:24:15.003197   16051 linux.go:183] eth0: max counter seen = 2810951704, max counter guess = 4294967296
I0401 10:24:15.003270   16051 linux.go:211] eth0: in=0.0024 kbps, out=0.0000 kbps
I0401 10:24:15.003334   16051 linux.go:183] lo: max counter seen = 4116, max counter guess = 4294967296
I0401 10:24:15.003402   16051 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 10:24:15.003461   16051 linux.go:183] wlan0: max counter seen = 0, max counter guess = 4294967296
I0401 10:24:15.003523   16051 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:24:15.003603   16051 throughput.go:38] prime 2
I0401 10:24:15.003814   16051 linux.go:231] found device wlan0
I0401 10:24:15.003863   16051 linux.go:236] bytesRecvStr for wlan0: 0
I0401 10:24:15.003909   16051 linux.go:237] bytesTransStr for wlan0: 0
I0401 10:24:15.003963   16051 linux.go:231] found device eth0
I0401 10:24:15.004005   16051 linux.go:236] bytesRecvStr for eth0: 2810951704
I0401 10:24:15.004050   16051 linux.go:237] bytesTransStr for eth0: 34188419
I0401 10:24:15.004122   16051 linux.go:231] found device lo
I0401 10:24:15.004168   16051 linux.go:236] bytesRecvStr for lo: 4116
I0401 10:24:15.004215   16051 linux.go:237] bytesTransStr for lo: 4116
I0401 10:24:15.004309   16051 linux.go:124] updating state for wlan0
I0401 10:24:15.004352   16051 linux.go:124] updating state for eth0
I0401 10:24:15.004392   16051 linux.go:124] updating state for lo
I0401 10:24:15.004444   16051 linux.go:183] eth0: max counter seen = 2810951704, max counter guess = 4294967296
I0401 10:24:15.004501   16051 linux.go:211] eth0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:24:15.004556   16051 linux.go:183] lo: max counter seen = 4116, max counter guess = 4294967296
I0401 10:24:15.004609   16051 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 10:24:15.004660   16051 linux.go:183] wlan0: max counter seen = 0, max counter guess = 4294967296
I0401 10:24:15.004712   16051 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:24:15.004987   16051 linux.go:231] found device wlan0
I0401 10:24:15.005037   16051 linux.go:236] bytesRecvStr for wlan0: 0
I0401 10:24:15.005083   16051 linux.go:237] bytesTransStr for wlan0: 0
I0401 10:24:15.005136   16051 linux.go:231] found device eth0
I0401 10:24:15.005180   16051 linux.go:236] bytesRecvStr for eth0: 2810951704
I0401 10:24:15.005224   16051 linux.go:237] bytesTransStr for eth0: 34188419
I0401 10:24:15.005276   16051 linux.go:231] found device lo
I0401 10:24:15.005321   16051 linux.go:236] bytesRecvStr for lo: 4116
I0401 10:24:15.005368   16051 linux.go:237] bytesTransStr for lo: 4116
I0401 10:24:15.005458   16051 linux.go:124] updating state for wlan0
I0401 10:24:15.005501   16051 linux.go:124] updating state for eth0
I0401 10:24:15.005541   16051 linux.go:124] updating state for lo
I0401 10:24:15.005592   16051 linux.go:183] eth0: max counter seen = 2810951704, max counter guess = 4294967296
I0401 10:24:15.005648   16051 linux.go:211] eth0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:24:15.005703   16051 linux.go:183] lo: max counter seen = 4116, max counter guess = 4294967296
I0401 10:24:15.005756   16051 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 10:24:15.005805   16051 linux.go:183] wlan0: max counter seen = 0, max counter guess = 4294967296
I0401 10:24:15.005858   16051 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
              eth0                  wlan0     
         In         Out         In         Out
I0401 10:24:16.006422   16051 linux.go:231] found device wlan0
I0401 10:24:16.006488   16051 linux.go:236] bytesRecvStr for wlan0: 0
I0401 10:24:16.006540   16051 linux.go:237] bytesTransStr for wlan0: 0
I0401 10:24:16.006599   16051 linux.go:231] found device eth0
I0401 10:24:16.006722   16051 linux.go:236] bytesRecvStr for eth0: 2810951901
I0401 10:24:16.006773   16051 linux.go:237] bytesTransStr for eth0: 34188419
I0401 10:24:16.006829   16051 linux.go:231] found device lo
I0401 10:24:16.007007   16051 linux.go:236] bytesRecvStr for lo: 4116
I0401 10:24:16.007054   16051 linux.go:237] bytesTransStr for lo: 4116
I0401 10:24:16.007158   16051 linux.go:124] updating state for wlan0
I0401 10:24:16.007202   16051 linux.go:124] updating state for eth0
I0401 10:24:16.007242   16051 linux.go:124] updating state for lo
I0401 10:24:16.007302   16051 linux.go:183] eth0: max counter seen = 2810951901, max counter guess = 4294967296
I0401 10:24:16.007366   16051 linux.go:211] eth0: in=1.5365 kbps, out=0.0000 kbps
I0401 10:24:16.007427   16051 linux.go:183] lo: max counter seen = 4116, max counter guess = 4294967296
I0401 10:24:16.007482   16051 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 10:24:16.007532   16051 linux.go:183] wlan0: max counter seen = 0, max counter guess = 4294967296
I0401 10:24:16.007586   16051 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:24:16.007666   16051 table.go:58] rows = 40, tableLineCount = 2
I0401 10:24:16.007717   16051 table.go:69] tableLineCount = 2, rows-3 = 37
       1.54        0.00       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
