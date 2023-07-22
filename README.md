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
bndstat v0.5.8
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 6d0aadd (v0.5.8)
Compiled: Sat Jul 22 10:16:49 PDT 2023
Build Host: janet
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
              ens4     
         In         Out
       0.34        0.17
       0.00        0.17
       0.00        0.00
       0.00        0.17
      19.70       19.09
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
I0722 10:17:05.754971   22299 bndstat.go:102] interval = 1.000000, count = 1
I0722 10:17:05.755096   22299 throughput.go:21] os is "linux"
I0722 10:17:05.755103   22299 throughput.go:33] running Reporter.Report() twice to prime the stats
I0722 10:17:05.755109   22299 throughput.go:35] prime 1
I0722 10:17:05.755238   22299 linux.go:231] found device lo
I0722 10:17:05.755251   22299 linux.go:236] bytesRecvStr for lo: 96236
I0722 10:17:05.755257   22299 linux.go:237] bytesTransStr for lo: 96236
I0722 10:17:05.755265   22299 linux.go:231] found device ens4
I0722 10:17:05.755271   22299 linux.go:236] bytesRecvStr for ens4: 201542786
I0722 10:17:05.755276   22299 linux.go:237] bytesTransStr for ens4: 14414118
I0722 10:17:05.755316   22299 linux.go:124] updating state for lo
I0722 10:17:05.755336   22299 linux.go:124] updating state for ens4
I0722 10:17:05.755346   22299 linux.go:183] ens4: max counter seen = 201542786, max counter guess = 4294967296
I0722 10:17:05.755361   22299 linux.go:211] ens4: in=0.0002 kbps, out=0.0000 kbps
I0722 10:17:05.755391   22299 linux.go:183] lo: max counter seen = 96236, max counter guess = 4294967296
I0722 10:17:05.755401   22299 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0722 10:17:05.755415   22299 throughput.go:38] prime 2
I0722 10:17:05.755487   22299 linux.go:231] found device lo
I0722 10:17:05.755502   22299 linux.go:236] bytesRecvStr for lo: 96236
I0722 10:17:05.755509   22299 linux.go:237] bytesTransStr for lo: 96236
I0722 10:17:05.755516   22299 linux.go:231] found device ens4
I0722 10:17:05.755521   22299 linux.go:236] bytesRecvStr for ens4: 201542786
I0722 10:17:05.755527   22299 linux.go:237] bytesTransStr for ens4: 14414118
I0722 10:17:05.755555   22299 linux.go:124] updating state for lo
I0722 10:17:05.755564   22299 linux.go:124] updating state for ens4
I0722 10:17:05.755571   22299 linux.go:183] ens4: max counter seen = 201542786, max counter guess = 4294967296
I0722 10:17:05.755583   22299 linux.go:211] ens4: in=0.0000 kbps, out=0.0000 kbps
I0722 10:17:05.755594   22299 linux.go:183] lo: max counter seen = 96236, max counter guess = 4294967296
I0722 10:17:05.755605   22299 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0722 10:17:05.755665   22299 linux.go:231] found device lo
I0722 10:17:05.755676   22299 linux.go:236] bytesRecvStr for lo: 96236
I0722 10:17:05.755683   22299 linux.go:237] bytesTransStr for lo: 96236
I0722 10:17:05.755693   22299 linux.go:231] found device ens4
I0722 10:17:05.755702   22299 linux.go:236] bytesRecvStr for ens4: 201542786
I0722 10:17:05.755709   22299 linux.go:237] bytesTransStr for ens4: 14414118
I0722 10:17:05.755729   22299 linux.go:124] updating state for lo
I0722 10:17:05.755754   22299 linux.go:124] updating state for ens4
I0722 10:17:05.755766   22299 linux.go:183] ens4: max counter seen = 201542786, max counter guess = 4294967296
I0722 10:17:05.755778   22299 linux.go:211] ens4: in=0.0000 kbps, out=0.0000 kbps
I0722 10:17:05.755787   22299 linux.go:183] lo: max counter seen = 96236, max counter guess = 4294967296
I0722 10:17:05.755795   22299 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
              ens4     
         In         Out
I0722 10:17:06.755995   22299 linux.go:231] found device lo
I0722 10:17:06.756017   22299 linux.go:236] bytesRecvStr for lo: 96236
I0722 10:17:06.756025   22299 linux.go:237] bytesTransStr for lo: 96236
I0722 10:17:06.756035   22299 linux.go:231] found device ens4
I0722 10:17:06.756041   22299 linux.go:236] bytesRecvStr for ens4: 201550896
I0722 10:17:06.756047   22299 linux.go:237] bytesTransStr for ens4: 14417261
I0722 10:17:06.756069   22299 linux.go:124] updating state for lo
I0722 10:17:06.756075   22299 linux.go:124] updating state for ens4
I0722 10:17:06.756085   22299 linux.go:183] ens4: max counter seen = 201550896, max counter guess = 4294967296
I0722 10:17:06.756096   22299 linux.go:211] ens4: in=63.3379 kbps, out=24.5464 kbps
I0722 10:17:06.756106   22299 linux.go:183] lo: max counter seen = 96236, max counter guess = 4294967296
I0722 10:17:06.756114   22299 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0722 10:17:06.756126   22299 table.go:58] rows = 40, tableLineCount = 2
I0722 10:17:06.756137   22299 table.go:69] tableLineCount = 2, rows-3 = 37
      63.34       24.55
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
