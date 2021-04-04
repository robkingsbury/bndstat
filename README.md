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
bndstat v0.5.1
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 3a1d777 (v0.5.1)
Compiled: Sun  4 Apr 11:13:38 PDT 2021
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
    1286.16     8019.28    7979.20     1318.68       1.62        1.82       0.00        0.00
     855.14     4061.25    4053.13      873.05       1.62        1.82       0.00        0.00
     536.52       41.75      68.65      542.22       1.50        1.69       0.00        0.00
     882.93     4066.67    4128.34      901.41       1.50        1.69       0.00        0.00
    1012.84     4494.17    4484.54     1032.76       1.50        1.69       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      52.56      213.19       1.50        1.69
      86.76      977.28       1.50        1.69
      94.45     1169.99       1.50        1.69
    7823.64     1033.32       1.50        1.69
      46.16      238.79       1.50        1.69
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0404 11:13:59.631996   30678 bndstat.go:101] interval = 1.000000, count = 1
I0404 11:13:59.632693   30678 throughput.go:21] os is "linux"
I0404 11:13:59.632808   30678 throughput.go:33] running Reporter.Report() twice to prime the stats
I0404 11:13:59.632877   30678 throughput.go:35] prime 1
I0404 11:13:59.633229   30678 linux.go:219] found device eth0
I0404 11:13:59.633302   30678 linux.go:224] bytesRecvStr for eth0: 2953134396
I0404 11:13:59.633409   30678 linux.go:225] bytesTransStr for eth0: 431635735
I0404 11:13:59.633486   30678 linux.go:219] found device eth1
I0404 11:13:59.633548   30678 linux.go:224] bytesRecvStr for eth1: 1990494721586
I0404 11:13:59.633613   30678 linux.go:225] bytesTransStr for eth1: 1019094263704
I0404 11:13:59.633684   30678 linux.go:219] found device eth2
I0404 11:13:59.633747   30678 linux.go:224] bytesRecvStr for eth2: 3020979899
I0404 11:13:59.633810   30678 linux.go:225] bytesTransStr for eth2: 3854881559
I0404 11:13:59.633883   30678 linux.go:219] found device wlan0
I0404 11:13:59.633944   30678 linux.go:224] bytesRecvStr for wlan0: 529377
I0404 11:13:59.634009   30678 linux.go:225] bytesTransStr for wlan0: 1341066
I0404 11:13:59.634082   30678 linux.go:219] found device lo
I0404 11:13:59.634142   30678 linux.go:224] bytesRecvStr for lo: 2766406
I0404 11:13:59.634206   30678 linux.go:225] bytesTransStr for lo: 2766406
I0404 11:13:59.634283   30678 linux.go:122] updating state for eth0
I0404 11:13:59.634348   30678 linux.go:122] updating state for eth1
I0404 11:13:59.634413   30678 linux.go:122] updating state for eth2
I0404 11:13:59.634478   30678 linux.go:122] updating state for wlan0
I0404 11:13:59.634540   30678 linux.go:122] updating state for lo
I0404 11:13:59.634609   30678 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:13:59.634702   30678 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.634805   30678 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:13:59.634889   30678 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.634990   30678 linux.go:171] eth0: max counter seen = 2953134396, max counter guess = 4294967296
I0404 11:13:59.635073   30678 linux.go:199] eth0: in=0.0025 kbps, out=0.0004 kbps
I0404 11:13:59.635173   30678 linux.go:171] eth1: max counter seen = 1990494721586, max counter guess = 18446744069414584320
I0404 11:13:59.635277   30678 linux.go:199] eth1: in=1.6860 kbps, out=0.8632 kbps
I0404 11:13:59.635377   30678 linux.go:171] eth2: max counter seen = 3854881559, max counter guess = 4294967296
I0404 11:13:59.635528   30678 linux.go:199] eth2: in=0.0026 kbps, out=0.0033 kbps
I0404 11:13:59.635714   30678 linux.go:171] eth0: max counter seen = 2953134396, max counter guess = 4294967296
I0404 11:13:59.635808   30678 linux.go:199] eth0: in=0.0025 kbps, out=0.0004 kbps
I0404 11:13:59.635911   30678 linux.go:171] eth1: max counter seen = 1990494721586, max counter guess = 18446744069414584320
I0404 11:13:59.636016   30678 linux.go:199] eth1: in=1.6860 kbps, out=0.8632 kbps
I0404 11:13:59.636114   30678 linux.go:171] eth2: max counter seen = 3854881559, max counter guess = 4294967296
I0404 11:13:59.636224   30678 linux.go:199] eth2: in=0.0026 kbps, out=0.0033 kbps
I0404 11:13:59.636323   30678 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:13:59.636408   30678 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.636506   30678 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:13:59.636617   30678 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.636716   30678 throughput.go:38] prime 2
I0404 11:13:59.636986   30678 linux.go:219] found device eth0
I0404 11:13:59.637053   30678 linux.go:224] bytesRecvStr for eth0: 2953135798
I0404 11:13:59.637120   30678 linux.go:225] bytesTransStr for eth0: 431635735
I0404 11:13:59.637192   30678 linux.go:219] found device eth1
I0404 11:13:59.637254   30678 linux.go:224] bytesRecvStr for eth1: 1990494721586
I0404 11:13:59.637317   30678 linux.go:225] bytesTransStr for eth1: 1019094265114
I0404 11:13:59.637391   30678 linux.go:219] found device eth2
I0404 11:13:59.637452   30678 linux.go:224] bytesRecvStr for eth2: 3020979899
I0404 11:13:59.637548   30678 linux.go:225] bytesTransStr for eth2: 3854881559
I0404 11:13:59.637624   30678 linux.go:219] found device wlan0
I0404 11:13:59.637684   30678 linux.go:224] bytesRecvStr for wlan0: 529377
I0404 11:13:59.637747   30678 linux.go:225] bytesTransStr for wlan0: 1341066
I0404 11:13:59.637818   30678 linux.go:219] found device lo
I0404 11:13:59.637878   30678 linux.go:224] bytesRecvStr for lo: 2766406
I0404 11:13:59.637940   30678 linux.go:225] bytesTransStr for lo: 2766406
I0404 11:13:59.638014   30678 linux.go:122] updating state for eth0
I0404 11:13:59.638075   30678 linux.go:122] updating state for eth1
I0404 11:13:59.638136   30678 linux.go:122] updating state for eth2
I0404 11:13:59.638198   30678 linux.go:122] updating state for wlan0
I0404 11:13:59.638257   30678 linux.go:122] updating state for lo
I0404 11:13:59.638323   30678 linux.go:171] eth0: max counter seen = 2953135798, max counter guess = 4294967296
I0404 11:13:59.638401   30678 linux.go:199] eth0: in=2934.6308 kbps, out=0.0000 kbps
I0404 11:13:59.638488   30678 linux.go:171] eth1: max counter seen = 1990494721586, max counter guess = 18446744069414584320
I0404 11:13:59.638583   30678 linux.go:199] eth1: in=0.0000 kbps, out=2951.3762 kbps
I0404 11:13:59.638668   30678 linux.go:171] eth2: max counter seen = 3854881559, max counter guess = 4294967296
I0404 11:13:59.638743   30678 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.638813   30678 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:13:59.638887   30678 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.638958   30678 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:13:59.639032   30678 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.639120   30678 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:13:59.639196   30678 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.639266   30678 linux.go:171] eth0: max counter seen = 2953135798, max counter guess = 4294967296
I0404 11:13:59.639341   30678 linux.go:199] eth0: in=2934.6308 kbps, out=0.0000 kbps
I0404 11:13:59.639424   30678 linux.go:171] eth1: max counter seen = 1990494721586, max counter guess = 18446744069414584320
I0404 11:13:59.639519   30678 linux.go:199] eth1: in=0.0000 kbps, out=2951.3762 kbps
I0404 11:13:59.639603   30678 linux.go:171] eth2: max counter seen = 3854881559, max counter guess = 4294967296
I0404 11:13:59.639678   30678 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.639748   30678 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:13:59.639821   30678 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.640077   30678 linux.go:219] found device eth0
I0404 11:13:59.640144   30678 linux.go:224] bytesRecvStr for eth0: 2953135798
I0404 11:13:59.640208   30678 linux.go:225] bytesTransStr for eth0: 431635735
I0404 11:13:59.640297   30678 linux.go:219] found device eth1
I0404 11:13:59.640359   30678 linux.go:224] bytesRecvStr for eth1: 1990494721586
I0404 11:13:59.640422   30678 linux.go:225] bytesTransStr for eth1: 1019094265114
I0404 11:13:59.640493   30678 linux.go:219] found device eth2
I0404 11:13:59.640556   30678 linux.go:224] bytesRecvStr for eth2: 3020979899
I0404 11:13:59.640620   30678 linux.go:225] bytesTransStr for eth2: 3854881559
I0404 11:13:59.640692   30678 linux.go:219] found device wlan0
I0404 11:13:59.640777   30678 linux.go:224] bytesRecvStr for wlan0: 529377
I0404 11:13:59.640840   30678 linux.go:225] bytesTransStr for wlan0: 1341066
I0404 11:13:59.640910   30678 linux.go:219] found device lo
I0404 11:13:59.640970   30678 linux.go:224] bytesRecvStr for lo: 2766406
I0404 11:13:59.641034   30678 linux.go:225] bytesTransStr for lo: 2766406
I0404 11:13:59.641107   30678 linux.go:122] updating state for eth0
I0404 11:13:59.641168   30678 linux.go:122] updating state for eth1
I0404 11:13:59.641229   30678 linux.go:122] updating state for eth2
I0404 11:13:59.641289   30678 linux.go:122] updating state for wlan0
I0404 11:13:59.641350   30678 linux.go:122] updating state for lo
I0404 11:13:59.641415   30678 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:13:59.641491   30678 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.641562   30678 linux.go:171] eth0: max counter seen = 2953135798, max counter guess = 4294967296
I0404 11:13:59.641638   30678 linux.go:199] eth0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.641709   30678 linux.go:171] eth1: max counter seen = 1990494721586, max counter guess = 18446744069414584320
I0404 11:13:59.641802   30678 linux.go:199] eth1: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.641872   30678 linux.go:171] eth2: max counter seen = 3854881559, max counter guess = 4294967296
I0404 11:13:59.641946   30678 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.642016   30678 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:13:59.642091   30678 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.642179   30678 linux.go:171] eth0: max counter seen = 2953135798, max counter guess = 4294967296
I0404 11:13:59.642254   30678 linux.go:199] eth0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.642324   30678 linux.go:171] eth1: max counter seen = 1990494721586, max counter guess = 18446744069414584320
I0404 11:13:59.642416   30678 linux.go:199] eth1: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.642486   30678 linux.go:171] eth2: max counter seen = 3854881559, max counter guess = 4294967296
I0404 11:13:59.642560   30678 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.642631   30678 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:13:59.642730   30678 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:13:59.642803   30678 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:13:59.642878   30678 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I0404 11:14:00.643698   30678 linux.go:219] found device eth0
I0404 11:14:00.643752   30678 linux.go:224] bytesRecvStr for eth0: 2953301344
I0404 11:14:00.643791   30678 linux.go:225] bytesTransStr for eth0: 433350349
I0404 11:14:00.643830   30678 linux.go:219] found device eth1
I0404 11:14:00.643862   30678 linux.go:224] bytesRecvStr for eth1: 1990496425602
I0404 11:14:00.643895   30678 linux.go:225] bytesTransStr for eth1: 1019094436459
I0404 11:14:00.643938   30678 linux.go:219] found device eth2
I0404 11:14:00.643969   30678 linux.go:224] bytesRecvStr for eth2: 3020980091
I0404 11:14:00.644004   30678 linux.go:225] bytesTransStr for eth2: 3854881775
I0404 11:14:00.644041   30678 linux.go:219] found device wlan0
I0404 11:14:00.644073   30678 linux.go:224] bytesRecvStr for wlan0: 529377
I0404 11:14:00.644105   30678 linux.go:225] bytesTransStr for wlan0: 1341066
I0404 11:14:00.644140   30678 linux.go:219] found device lo
I0404 11:14:00.644176   30678 linux.go:224] bytesRecvStr for lo: 2766406
I0404 11:14:00.644209   30678 linux.go:225] bytesTransStr for lo: 2766406
I0404 11:14:00.644248   30678 linux.go:122] updating state for eth0
I0404 11:14:00.644280   30678 linux.go:122] updating state for eth1
I0404 11:14:00.644311   30678 linux.go:122] updating state for eth2
I0404 11:14:00.644359   30678 linux.go:122] updating state for wlan0
I0404 11:14:00.644389   30678 linux.go:122] updating state for lo
I0404 11:14:00.644425   30678 linux.go:171] eth0: max counter seen = 2953301344, max counter guess = 4294967296
I0404 11:14:00.644469   30678 linux.go:199] eth0: in=1289.2738 kbps, out=13353.4304 kbps
I0404 11:14:00.644523   30678 linux.go:171] eth1: max counter seen = 1990496425602, max counter guess = 18446744069414584320
I0404 11:14:00.644572   30678 linux.go:199] eth1: in=13270.8931 kbps, out=1334.4365 kbps
I0404 11:14:00.644621   30678 linux.go:171] eth2: max counter seen = 3854881775, max counter guess = 4294967296
I0404 11:14:00.644659   30678 linux.go:199] eth2: in=1.4953 kbps, out=1.6822 kbps
I0404 11:14:00.644708   30678 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:14:00.644746   30678 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:14:00.644782   30678 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:14:00.644819   30678 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:14:00.644865   30678 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:14:00.644903   30678 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:14:00.644948   30678 linux.go:171] eth0: max counter seen = 2953301344, max counter guess = 4294967296
I0404 11:14:00.644987   30678 linux.go:199] eth0: in=1289.2738 kbps, out=13353.4304 kbps
I0404 11:14:00.645041   30678 linux.go:171] eth1: max counter seen = 1990496425602, max counter guess = 18446744069414584320
I0404 11:14:00.645089   30678 linux.go:199] eth1: in=13270.8931 kbps, out=1334.4365 kbps
I0404 11:14:00.645138   30678 linux.go:171] eth2: max counter seen = 3854881775, max counter guess = 4294967296
I0404 11:14:00.645175   30678 linux.go:199] eth2: in=1.4953 kbps, out=1.6822 kbps
I0404 11:14:00.645224   30678 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:14:00.645261   30678 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:14:00.645309   30678 table.go:58] rows = 40, tableLineCount = 2
I0404 11:14:00.645342   30678 table.go:69] tableLineCount = 2, rows-3 = 37
    1289.27    13353.43   13270.89     1334.44       1.50        1.68       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
