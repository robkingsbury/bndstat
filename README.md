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
bndstat 0.5.10
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 135a5c4 (0.5.10)
Compiled: Sat 20 Apr 13:00:20 PDT 2024
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
    2908.46      110.27     175.06     2928.17       1.75        1.69       0.00        0.00
    3133.88      103.10     166.76     3153.83       1.44        2.04       0.00        0.00
    2827.76      988.41    1036.25     2847.38       2.07        1.91       0.00        0.00
    2843.33       74.88     128.50     2861.50       1.75        1.82       0.00        0.00
    2576.87       83.15     141.99     2594.30       1.82        2.04       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
     268.38     6826.87       1.50        1.69
      96.89      728.81       1.50        1.69
     134.22     2830.67       1.50        1.69
    2906.89     6714.61       2.59        3.09
    1322.25     1067.77       1.87        2.08
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0420 13:00:42.301273   26193 bndstat.go:102] interval = 1.000000, count = 1
I0420 13:00:42.301839   26193 throughput.go:21] os is "linux"
I0420 13:00:42.301885   26193 throughput.go:33] running Reporter.Report() twice to prime the stats
I0420 13:00:42.301923   26193 throughput.go:35] prime 1
I0420 13:00:42.302229   26193 linux.go:236] found device eth0
I0420 13:00:42.302360   26193 linux.go:241] bytesRecvStr for eth0: 3987372131
I0420 13:00:42.302408   26193 linux.go:242] bytesTransStr for eth0: 4039211551
I0420 13:00:42.302456   26193 linux.go:236] found device eth1
I0420 13:00:42.302497   26193 linux.go:241] bytesRecvStr for eth1: 5280281321644
I0420 13:00:42.302539   26193 linux.go:242] bytesTransStr for eth1: 6976147925318
I0420 13:00:42.302585   26193 linux.go:236] found device eth2
I0420 13:00:42.302626   26193 linux.go:241] bytesRecvStr for eth2: 6972102825
I0420 13:00:42.302666   26193 linux.go:242] bytesTransStr for eth2: 7276132774
I0420 13:00:42.302714   26193 linux.go:236] found device wlan0
I0420 13:00:42.302754   26193 linux.go:241] bytesRecvStr for wlan0: 0
I0420 13:00:42.302795   26193 linux.go:242] bytesTransStr for wlan0: 0
I0420 13:00:42.302838   26193 linux.go:236] found device lo
I0420 13:00:42.302878   26193 linux.go:241] bytesRecvStr for lo: 12250788
I0420 13:00:42.302918   26193 linux.go:242] bytesTransStr for lo: 12250788
I0420 13:00:42.303018   26193 linux.go:129] updating state for eth0
I0420 13:00:42.303057   26193 linux.go:129] updating state for eth1
I0420 13:00:42.303091   26193 linux.go:129] updating state for eth2
I0420 13:00:42.303124   26193 linux.go:129] updating state for wlan0
I0420 13:00:42.303157   26193 linux.go:129] updating state for lo
I0420 13:00:42.303234   26193 linux.go:188] eth0: max counter seen = 4039211551, max counter guess = 4294967296
I0420 13:00:42.303302   26193 linux.go:216] eth0: in=0.0034 kbps, out=0.0034 kbps
I0420 13:00:42.303357   26193 linux.go:188] eth1: max counter seen = 6976147925318, max counter guess = 18446744069414584320
I0420 13:00:42.303415   26193 linux.go:216] eth1: in=4.4726 kbps, out=5.9090 kbps
I0420 13:00:42.303464   26193 linux.go:188] eth2: max counter seen = 7276132774, max counter guess = 18446744069414584320
I0420 13:00:42.303520   26193 linux.go:216] eth2: in=0.0059 kbps, out=0.0062 kbps
I0420 13:00:42.303569   26193 linux.go:188] lo: max counter seen = 12250788, max counter guess = 4294967296
I0420 13:00:42.303624   26193 linux.go:216] lo: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.303673   26193 linux.go:188] wlan0: max counter seen = 0, max counter guess = 4294967296
I0420 13:00:42.303727   26193 linux.go:216] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.303805   26193 throughput.go:38] prime 2
I0420 13:00:42.304003   26193 linux.go:236] found device eth0
I0420 13:00:42.304048   26193 linux.go:241] bytesRecvStr for eth0: 3987372131
I0420 13:00:42.304089   26193 linux.go:242] bytesTransStr for eth0: 4039211551
I0420 13:00:42.304135   26193 linux.go:236] found device eth1
I0420 13:00:42.304175   26193 linux.go:241] bytesRecvStr for eth1: 5280281321692
I0420 13:00:42.304215   26193 linux.go:242] bytesTransStr for eth1: 6976147925318
I0420 13:00:42.304260   26193 linux.go:236] found device eth2
I0420 13:00:42.304300   26193 linux.go:241] bytesRecvStr for eth2: 6972102825
I0420 13:00:42.304340   26193 linux.go:242] bytesTransStr for eth2: 7276132774
I0420 13:00:42.304386   26193 linux.go:236] found device wlan0
I0420 13:00:42.304426   26193 linux.go:241] bytesRecvStr for wlan0: 0
I0420 13:00:42.304467   26193 linux.go:242] bytesTransStr for wlan0: 0
I0420 13:00:42.304512   26193 linux.go:236] found device lo
I0420 13:00:42.304551   26193 linux.go:241] bytesRecvStr for lo: 12250788
I0420 13:00:42.304593   26193 linux.go:242] bytesTransStr for lo: 12250788
I0420 13:00:42.304670   26193 linux.go:129] updating state for eth0
I0420 13:00:42.304706   26193 linux.go:129] updating state for eth1
I0420 13:00:42.304739   26193 linux.go:129] updating state for eth2
I0420 13:00:42.304771   26193 linux.go:129] updating state for wlan0
I0420 13:00:42.304804   26193 linux.go:129] updating state for lo
I0420 13:00:42.304849   26193 linux.go:188] eth0: max counter seen = 4039211551, max counter guess = 4294967296
I0420 13:00:42.304901   26193 linux.go:216] eth0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.304946   26193 linux.go:188] eth1: max counter seen = 6976147925318, max counter guess = 18446744069414584320
I0420 13:00:42.305032   26193 linux.go:216] eth1: in=226.8404 kbps, out=0.0000 kbps
I0420 13:00:42.305082   26193 linux.go:188] eth2: max counter seen = 7276132774, max counter guess = 18446744069414584320
I0420 13:00:42.305133   26193 linux.go:216] eth2: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.305176   26193 linux.go:188] lo: max counter seen = 12250788, max counter guess = 4294967296
I0420 13:00:42.305224   26193 linux.go:216] lo: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.305266   26193 linux.go:188] wlan0: max counter seen = 0, max counter guess = 4294967296
I0420 13:00:42.305314   26193 linux.go:216] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.305568   26193 linux.go:236] found device eth0
I0420 13:00:42.305612   26193 linux.go:241] bytesRecvStr for eth0: 3987372131
I0420 13:00:42.305653   26193 linux.go:242] bytesTransStr for eth0: 4039211551
I0420 13:00:42.305700   26193 linux.go:236] found device eth1
I0420 13:00:42.305740   26193 linux.go:241] bytesRecvStr for eth1: 5280281321692
I0420 13:00:42.305781   26193 linux.go:242] bytesTransStr for eth1: 6976147925318
I0420 13:00:42.305827   26193 linux.go:236] found device eth2
I0420 13:00:42.305867   26193 linux.go:241] bytesRecvStr for eth2: 6972102825
I0420 13:00:42.305907   26193 linux.go:242] bytesTransStr for eth2: 7276132774
I0420 13:00:42.305954   26193 linux.go:236] found device wlan0
I0420 13:00:42.305993   26193 linux.go:241] bytesRecvStr for wlan0: 0
I0420 13:00:42.306034   26193 linux.go:242] bytesTransStr for wlan0: 0
I0420 13:00:42.306240   26193 linux.go:236] found device lo
I0420 13:00:42.306284   26193 linux.go:241] bytesRecvStr for lo: 12250788
I0420 13:00:42.306325   26193 linux.go:242] bytesTransStr for lo: 12250788
I0420 13:00:42.306407   26193 linux.go:129] updating state for eth0
I0420 13:00:42.306443   26193 linux.go:129] updating state for eth1
I0420 13:00:42.306476   26193 linux.go:129] updating state for eth2
I0420 13:00:42.306508   26193 linux.go:129] updating state for wlan0
I0420 13:00:42.306540   26193 linux.go:129] updating state for lo
I0420 13:00:42.306588   26193 linux.go:188] eth0: max counter seen = 4039211551, max counter guess = 4294967296
I0420 13:00:42.306645   26193 linux.go:216] eth0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.306691   26193 linux.go:188] eth1: max counter seen = 6976147925318, max counter guess = 18446744069414584320
I0420 13:00:42.306741   26193 linux.go:216] eth1: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.306784   26193 linux.go:188] eth2: max counter seen = 7276132774, max counter guess = 18446744069414584320
I0420 13:00:42.306833   26193 linux.go:216] eth2: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.306875   26193 linux.go:188] lo: max counter seen = 12250788, max counter guess = 4294967296
I0420 13:00:42.306923   26193 linux.go:216] lo: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:42.306983   26193 linux.go:188] wlan0: max counter seen = 0, max counter guess = 4294967296
I0420 13:00:42.307033   26193 linux.go:216] wlan0: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I0420 13:00:43.307649   26193 linux.go:236] found device eth0
I0420 13:00:43.307712   26193 linux.go:241] bytesRecvStr for eth0: 3987853731
I0420 13:00:43.307758   26193 linux.go:242] bytesTransStr for eth0: 4039241860
I0420 13:00:43.307809   26193 linux.go:236] found device eth1
I0420 13:00:43.307851   26193 linux.go:241] bytesRecvStr for eth1: 5280281358747
I0420 13:00:43.307893   26193 linux.go:242] bytesTransStr for eth1: 6976148410102
I0420 13:00:43.307941   26193 linux.go:236] found device eth2
I0420 13:00:43.307981   26193 linux.go:241] bytesRecvStr for eth2: 6972103017
I0420 13:00:43.308023   26193 linux.go:242] bytesTransStr for eth2: 7276132990
I0420 13:00:43.308070   26193 linux.go:236] found device wlan0
I0420 13:00:43.308110   26193 linux.go:241] bytesRecvStr for wlan0: 0
I0420 13:00:43.308149   26193 linux.go:242] bytesTransStr for wlan0: 0
I0420 13:00:43.308194   26193 linux.go:236] found device lo
I0420 13:00:43.308234   26193 linux.go:241] bytesRecvStr for lo: 12250788
I0420 13:00:43.308275   26193 linux.go:242] bytesTransStr for lo: 12250788
I0420 13:00:43.308363   26193 linux.go:129] updating state for eth0
I0420 13:00:43.308401   26193 linux.go:129] updating state for eth1
I0420 13:00:43.308434   26193 linux.go:129] updating state for eth2
I0420 13:00:43.308467   26193 linux.go:129] updating state for wlan0
I0420 13:00:43.308498   26193 linux.go:129] updating state for lo
I0420 13:00:43.308595   26193 linux.go:188] eth0: max counter seen = 4039241860, max counter guess = 4294967296
I0420 13:00:43.308659   26193 linux.go:216] eth0: in=3755.1614 kbps, out=236.3272 kbps
I0420 13:00:43.308714   26193 linux.go:188] eth1: max counter seen = 6976148410102, max counter guess = 18446744069414584320
I0420 13:00:43.308766   26193 linux.go:216] eth1: in=288.9275 kbps, out=3779.9879 kbps
I0420 13:00:43.308814   26193 linux.go:188] eth2: max counter seen = 7276132990, max counter guess = 18446744069414584320
I0420 13:00:43.308866   26193 linux.go:216] eth2: in=1.4971 kbps, out=1.6842 kbps
I0420 13:00:43.308914   26193 linux.go:188] lo: max counter seen = 12250788, max counter guess = 4294967296
I0420 13:00:43.308963   26193 linux.go:216] lo: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:43.309005   26193 linux.go:188] wlan0: max counter seen = 0, max counter guess = 4294967296
I0420 13:00:43.309055   26193 linux.go:216] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0420 13:00:43.309130   26193 table.go:58] rows = 40, tableLineCount = 2
I0420 13:00:43.309173   26193 table.go:69] tableLineCount = 2, rows-3 = 37
    3755.16      236.33     288.93     3779.99       1.50        1.68       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
