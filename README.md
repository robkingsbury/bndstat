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
bndstat v0.5.6
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 8d180d9 (v0.5.6)
Compiled: Sat  1 Apr 11:01:27 PDT 2023
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
    1822.18     1222.93    1265.50     1837.08       1.75        1.95       0.00        0.00
    1869.38      115.12     194.92     1884.01       1.50        1.69       0.00        0.00
    1860.76      257.15     326.84     1876.01       1.62        1.82       0.00        0.00
    2144.12     2364.78    2415.03     2166.17       1.50        1.69       0.00        0.00
    1944.84       84.80     152.25     1959.31       1.62        1.82       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
     140.51     1197.13       1.50        1.69
    1243.41     1800.73       1.50        1.69
     978.56     3757.86       1.87        2.08
     472.70      786.82       1.50        1.69
     210.91     1066.74       1.50        1.69
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0401 11:01:48.371932    1966 bndstat.go:101] interval = 1.000000, count = 1
I0401 11:01:48.372558    1966 throughput.go:21] os is "linux"
I0401 11:01:48.372659    1966 throughput.go:33] running Reporter.Report() twice to prime the stats
I0401 11:01:48.372727    1966 throughput.go:35] prime 1
I0401 11:01:48.373088    1966 linux.go:231] found device eth0
I0401 11:01:48.373160    1966 linux.go:236] bytesRecvStr for eth0: 2371406906
I0401 11:01:48.373231    1966 linux.go:237] bytesTransStr for eth0: 2225423577
I0401 11:01:48.373307    1966 linux.go:231] found device eth1
I0401 11:01:48.373369    1966 linux.go:236] bytesRecvStr for eth1: 3579724653843
I0401 11:01:48.373435    1966 linux.go:237] bytesTransStr for eth1: 3047641951951
I0401 11:01:48.373508    1966 linux.go:231] found device eth2
I0401 11:01:48.373570    1966 linux.go:236] bytesRecvStr for eth2: 2838452801
I0401 11:01:48.373711    1966 linux.go:237] bytesTransStr for eth2: 3167711773
I0401 11:01:48.373793    1966 linux.go:231] found device wlan0
I0401 11:01:48.373856    1966 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 11:01:48.373920    1966 linux.go:237] bytesTransStr for wlan0: 33550
I0401 11:01:48.373992    1966 linux.go:231] found device lo
I0401 11:01:48.374070    1966 linux.go:236] bytesRecvStr for lo: 349458
I0401 11:01:48.374136    1966 linux.go:237] bytesTransStr for lo: 349458
I0401 11:01:48.374262    1966 linux.go:124] updating state for eth0
I0401 11:01:48.374406    1966 linux.go:124] updating state for eth1
I0401 11:01:48.374470    1966 linux.go:124] updating state for eth2
I0401 11:01:48.374533    1966 linux.go:124] updating state for wlan0
I0401 11:01:48.374593    1966 linux.go:124] updating state for lo
I0401 11:01:48.374673    1966 linux.go:183] eth0: max counter seen = 2371406906, max counter guess = 4294967296
I0401 11:01:48.374769    1966 linux.go:211] eth0: in=0.0020 kbps, out=0.0019 kbps
I0401 11:01:48.374874    1966 linux.go:183] eth1: max counter seen = 3579724653843, max counter guess = 18446744069414584320
I0401 11:01:48.374979    1966 linux.go:211] eth1: in=3.0321 kbps, out=2.5815 kbps
I0401 11:01:48.375079    1966 linux.go:183] eth2: max counter seen = 3167711773, max counter guess = 4294967296
I0401 11:01:48.375164    1966 linux.go:211] eth2: in=0.0024 kbps, out=0.0027 kbps
I0401 11:01:48.375266    1966 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 11:01:48.375349    1966 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.375447    1966 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 11:01:48.375576    1966 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.375708    1966 throughput.go:38] prime 2
I0401 11:01:48.375934    1966 linux.go:231] found device eth0
I0401 11:01:48.376000    1966 linux.go:236] bytesRecvStr for eth0: 2371406906
I0401 11:01:48.376066    1966 linux.go:237] bytesTransStr for eth0: 2225423577
I0401 11:01:48.376139    1966 linux.go:231] found device eth1
I0401 11:01:48.376200    1966 linux.go:236] bytesRecvStr for eth1: 3579724653843
I0401 11:01:48.376264    1966 linux.go:237] bytesTransStr for eth1: 3047641951951
I0401 11:01:48.376337    1966 linux.go:231] found device eth2
I0401 11:01:48.376410    1966 linux.go:236] bytesRecvStr for eth2: 2838452801
I0401 11:01:48.376474    1966 linux.go:237] bytesTransStr for eth2: 3167711773
I0401 11:01:48.376547    1966 linux.go:231] found device wlan0
I0401 11:01:48.376608    1966 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 11:01:48.376722    1966 linux.go:237] bytesTransStr for wlan0: 33550
I0401 11:01:48.376843    1966 linux.go:231] found device lo
I0401 11:01:48.376905    1966 linux.go:236] bytesRecvStr for lo: 349458
I0401 11:01:48.376969    1966 linux.go:237] bytesTransStr for lo: 349458
I0401 11:01:48.377076    1966 linux.go:124] updating state for eth0
I0401 11:01:48.377143    1966 linux.go:124] updating state for eth1
I0401 11:01:48.377203    1966 linux.go:124] updating state for eth2
I0401 11:01:48.377264    1966 linux.go:124] updating state for wlan0
I0401 11:01:48.377324    1966 linux.go:124] updating state for lo
I0401 11:01:48.377399    1966 linux.go:183] eth0: max counter seen = 2371406906, max counter guess = 4294967296
I0401 11:01:48.377480    1966 linux.go:211] eth0: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.377555    1966 linux.go:183] eth1: max counter seen = 3579724653843, max counter guess = 18446744069414584320
I0401 11:01:48.377650    1966 linux.go:211] eth1: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.377722    1966 linux.go:183] eth2: max counter seen = 3167711773, max counter guess = 4294967296
I0401 11:01:48.377798    1966 linux.go:211] eth2: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.377869    1966 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 11:01:48.377943    1966 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.378015    1966 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 11:01:48.378089    1966 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.378390    1966 linux.go:231] found device eth0
I0401 11:01:48.378461    1966 linux.go:236] bytesRecvStr for eth0: 2371408380
I0401 11:01:48.378526    1966 linux.go:237] bytesTransStr for eth0: 2225423577
I0401 11:01:48.378598    1966 linux.go:231] found device eth1
I0401 11:01:48.378659    1966 linux.go:236] bytesRecvStr for eth1: 3579724653843
I0401 11:01:48.378722    1966 linux.go:237] bytesTransStr for eth1: 3047641951951
I0401 11:01:48.378793    1966 linux.go:231] found device eth2
I0401 11:01:48.378855    1966 linux.go:236] bytesRecvStr for eth2: 2838452801
I0401 11:01:48.378918    1966 linux.go:237] bytesTransStr for eth2: 3167711773
I0401 11:01:48.378990    1966 linux.go:231] found device wlan0
I0401 11:01:48.379051    1966 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 11:01:48.379115    1966 linux.go:237] bytesTransStr for wlan0: 33550
I0401 11:01:48.379184    1966 linux.go:231] found device lo
I0401 11:01:48.379245    1966 linux.go:236] bytesRecvStr for lo: 349458
I0401 11:01:48.379307    1966 linux.go:237] bytesTransStr for lo: 349458
I0401 11:01:48.379415    1966 linux.go:124] updating state for eth0
I0401 11:01:48.379482    1966 linux.go:124] updating state for eth1
I0401 11:01:48.379559    1966 linux.go:124] updating state for eth2
I0401 11:01:48.379622    1966 linux.go:124] updating state for wlan0
I0401 11:01:48.379681    1966 linux.go:124] updating state for lo
I0401 11:01:48.379755    1966 linux.go:183] eth0: max counter seen = 2371408380, max counter guess = 4294967296
I0401 11:01:48.379835    1966 linux.go:211] eth0: in=4920.4166 kbps, out=0.0000 kbps
I0401 11:01:48.379922    1966 linux.go:183] eth1: max counter seen = 3579724653843, max counter guess = 18446744069414584320
I0401 11:01:48.380017    1966 linux.go:211] eth1: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.380089    1966 linux.go:183] eth2: max counter seen = 3167711773, max counter guess = 4294967296
I0401 11:01:48.380166    1966 linux.go:211] eth2: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.380236    1966 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 11:01:48.380313    1966 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:48.380411    1966 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 11:01:48.380489    1966 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I0401 11:01:49.381765    1966 linux.go:231] found device eth0
I0401 11:01:49.381894    1966 linux.go:236] bytesRecvStr for eth0: 2371965339
I0401 11:01:49.381988    1966 linux.go:237] bytesTransStr for eth0: 2225440757
I0401 11:01:49.382084    1966 linux.go:231] found device eth1
I0401 11:01:49.382163    1966 linux.go:236] bytesRecvStr for eth1: 3579724676225
I0401 11:01:49.382247    1966 linux.go:237] bytesTransStr for eth1: 3047642513971
I0401 11:01:49.382342    1966 linux.go:231] found device eth2
I0401 11:01:49.382421    1966 linux.go:236] bytesRecvStr for eth2: 2838452993
I0401 11:01:49.382501    1966 linux.go:237] bytesTransStr for eth2: 3167711989
I0401 11:01:49.382605    1966 linux.go:231] found device wlan0
I0401 11:01:49.382691    1966 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 11:01:49.382780    1966 linux.go:237] bytesTransStr for wlan0: 33550
I0401 11:01:49.382872    1966 linux.go:231] found device lo
I0401 11:01:49.382951    1966 linux.go:236] bytesRecvStr for lo: 349458
I0401 11:01:49.383032    1966 linux.go:237] bytesTransStr for lo: 349458
I0401 11:01:49.383194    1966 linux.go:124] updating state for eth0
I0401 11:01:49.383281    1966 linux.go:124] updating state for eth1
I0401 11:01:49.383359    1966 linux.go:124] updating state for eth2
I0401 11:01:49.383436    1966 linux.go:124] updating state for wlan0
I0401 11:01:49.383513    1966 linux.go:124] updating state for lo
I0401 11:01:49.383612    1966 linux.go:183] eth0: max counter seen = 2371965339, max counter guess = 4294967296
I0401 11:01:49.383785    1966 linux.go:211] eth0: in=4334.8834 kbps, out=133.7141 kbps
I0401 11:01:49.383920    1966 linux.go:183] eth1: max counter seen = 3579724676225, max counter guess = 18446744069414584320
I0401 11:01:49.384047    1966 linux.go:211] eth1: in=174.2020 kbps, out=4374.2738 kbps
I0401 11:01:49.384193    1966 linux.go:183] eth2: max counter seen = 3167711989, max counter guess = 4294967296
I0401 11:01:49.384292    1966 linux.go:211] eth2: in=1.4944 kbps, out=1.6812 kbps
I0401 11:01:49.384417    1966 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 11:01:49.384512    1966 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:49.384603    1966 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 11:01:49.384699    1966 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 11:01:49.384828    1966 table.go:58] rows = 40, tableLineCount = 2
I0401 11:01:49.384912    1966 table.go:69] tableLineCount = 2, rows-3 = 37
    4334.88      133.71     174.20     4374.27       1.49        1.68       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
