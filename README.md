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
bndstat v0.5.5
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 80b92ab (v0.5.5)
Compiled: Sat  1 Apr 10:51:41 PDT 2023
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
    1640.97      942.13     990.94     1654.97       1.62        1.82       0.00        0.00
    1505.31       88.05     140.05     1517.61       1.62        1.82       0.00        0.00
    1631.44       91.68     140.40     1644.87       1.62        1.82       0.00        0.00
    1673.42      967.90    1006.87     1687.91       1.50        1.69       0.00        0.00
    1691.99       62.81     111.38     1704.87       1.62        1.82       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
     143.70     2814.02       1.50        1.69
      76.73      640.15       1.50        1.69
      80.97      506.16       1.50        1.69
     140.37     2605.08       1.87        2.08
     108.15     1888.54       1.50        1.69
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0401 10:52:02.208542    1430 bndstat.go:101] interval = 1.000000, count = 1
I0401 10:52:02.209069    1430 throughput.go:21] os is "linux"
I0401 10:52:02.209111    1430 throughput.go:33] running Reporter.Report() twice to prime the stats
I0401 10:52:02.209181    1430 throughput.go:35] prime 1
I0401 10:52:02.209540    1430 linux.go:231] found device eth0
I0401 10:52:02.209582    1430 linux.go:236] bytesRecvStr for eth0: 2230105326
I0401 10:52:02.209619    1430 linux.go:237] bytesTransStr for eth0: 2181020571
I0401 10:52:02.209664    1430 linux.go:231] found device eth1
I0401 10:52:02.209699    1430 linux.go:236] bytesRecvStr for eth1: 3579676644183
I0401 10:52:02.209734    1430 linux.go:237] bytesTransStr for eth1: 3047499564215
I0401 10:52:02.209777    1430 linux.go:231] found device eth2
I0401 10:52:02.209811    1430 linux.go:236] bytesRecvStr for eth2: 2838333599
I0401 10:52:02.209845    1430 linux.go:237] bytesTransStr for eth2: 3167578447
I0401 10:52:02.209888    1430 linux.go:231] found device wlan0
I0401 10:52:02.209923    1430 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 10:52:02.209957    1430 linux.go:237] bytesTransStr for wlan0: 33550
I0401 10:52:02.209998    1430 linux.go:231] found device lo
I0401 10:52:02.210032    1430 linux.go:236] bytesRecvStr for lo: 349458
I0401 10:52:02.210066    1430 linux.go:237] bytesTransStr for lo: 349458
I0401 10:52:02.210161    1430 linux.go:124] updating state for eth0
I0401 10:52:02.210200    1430 linux.go:124] updating state for eth1
I0401 10:52:02.210258    1430 linux.go:124] updating state for eth2
I0401 10:52:02.210292    1430 linux.go:124] updating state for wlan0
I0401 10:52:02.210324    1430 linux.go:124] updating state for lo
I0401 10:52:02.210374    1430 linux.go:183] eth0: max counter seen = 2230105326, max counter guess = 4294967296
I0401 10:52:02.210436    1430 linux.go:211] eth0: in=0.0019 kbps, out=0.0018 kbps
I0401 10:52:02.210511    1430 linux.go:183] eth1: max counter seen = 3579676644183, max counter guess = 18446744069414584320
I0401 10:52:02.210582    1430 linux.go:211] eth1: in=3.0321 kbps, out=2.5813 kbps
I0401 10:52:02.210652    1430 linux.go:183] eth2: max counter seen = 3167578447, max counter guess = 4294967296
I0401 10:52:02.210703    1430 linux.go:211] eth2: in=0.0024 kbps, out=0.0027 kbps
I0401 10:52:02.210773    1430 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 10:52:02.210824    1430 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.210891    1430 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 10:52:02.210942    1430 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.211042    1430 throughput.go:38] prime 2
I0401 10:52:02.211230    1430 linux.go:231] found device eth0
I0401 10:52:02.211269    1430 linux.go:236] bytesRecvStr for eth0: 2230105326
I0401 10:52:02.211304    1430 linux.go:237] bytesTransStr for eth0: 2181020571
I0401 10:52:02.211348    1430 linux.go:231] found device eth1
I0401 10:52:02.211382    1430 linux.go:236] bytesRecvStr for eth1: 3579676644183
I0401 10:52:02.211417    1430 linux.go:237] bytesTransStr for eth1: 3047499564215
I0401 10:52:02.211458    1430 linux.go:231] found device eth2
I0401 10:52:02.211493    1430 linux.go:236] bytesRecvStr for eth2: 2838333599
I0401 10:52:02.211528    1430 linux.go:237] bytesTransStr for eth2: 3167578447
I0401 10:52:02.211571    1430 linux.go:231] found device wlan0
I0401 10:52:02.211606    1430 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 10:52:02.211640    1430 linux.go:237] bytesTransStr for wlan0: 33550
I0401 10:52:02.211681    1430 linux.go:231] found device lo
I0401 10:52:02.211714    1430 linux.go:236] bytesRecvStr for lo: 349458
I0401 10:52:02.211747    1430 linux.go:237] bytesTransStr for lo: 349458
I0401 10:52:02.211823    1430 linux.go:124] updating state for eth0
I0401 10:52:02.211858    1430 linux.go:124] updating state for eth1
I0401 10:52:02.211890    1430 linux.go:124] updating state for eth2
I0401 10:52:02.211954    1430 linux.go:124] updating state for wlan0
I0401 10:52:02.211990    1430 linux.go:124] updating state for lo
I0401 10:52:02.212035    1430 linux.go:183] eth0: max counter seen = 2230105326, max counter guess = 4294967296
I0401 10:52:02.212082    1430 linux.go:211] eth0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.212125    1430 linux.go:183] eth1: max counter seen = 3579676644183, max counter guess = 18446744069414584320
I0401 10:52:02.212189    1430 linux.go:211] eth1: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.212232    1430 linux.go:183] eth2: max counter seen = 3167578447, max counter guess = 4294967296
I0401 10:52:02.212277    1430 linux.go:211] eth2: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.212319    1430 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 10:52:02.212362    1430 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.212403    1430 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 10:52:02.212446    1430 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.212691    1430 linux.go:231] found device eth0
I0401 10:52:02.212729    1430 linux.go:236] bytesRecvStr for eth0: 2230105326
I0401 10:52:02.212764    1430 linux.go:237] bytesTransStr for eth0: 2181020571
I0401 10:52:02.212807    1430 linux.go:231] found device eth1
I0401 10:52:02.212841    1430 linux.go:236] bytesRecvStr for eth1: 3579676644231
I0401 10:52:02.212876    1430 linux.go:237] bytesTransStr for eth1: 3047499564215
I0401 10:52:02.212917    1430 linux.go:231] found device eth2
I0401 10:52:02.212952    1430 linux.go:236] bytesRecvStr for eth2: 2838333599
I0401 10:52:02.212986    1430 linux.go:237] bytesTransStr for eth2: 3167578447
I0401 10:52:02.213028    1430 linux.go:231] found device wlan0
I0401 10:52:02.213062    1430 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 10:52:02.213096    1430 linux.go:237] bytesTransStr for wlan0: 33550
I0401 10:52:02.213136    1430 linux.go:231] found device lo
I0401 10:52:02.213170    1430 linux.go:236] bytesRecvStr for lo: 349458
I0401 10:52:02.213204    1430 linux.go:237] bytesTransStr for lo: 349458
I0401 10:52:02.213280    1430 linux.go:124] updating state for eth0
I0401 10:52:02.213315    1430 linux.go:124] updating state for eth1
I0401 10:52:02.213347    1430 linux.go:124] updating state for eth2
I0401 10:52:02.213379    1430 linux.go:124] updating state for wlan0
I0401 10:52:02.213410    1430 linux.go:124] updating state for lo
I0401 10:52:02.213452    1430 linux.go:183] eth0: max counter seen = 2230105326, max counter guess = 4294967296
I0401 10:52:02.213497    1430 linux.go:211] eth0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.213541    1430 linux.go:183] eth1: max counter seen = 3579676644231, max counter guess = 18446744069414584320
I0401 10:52:02.213621    1430 linux.go:211] eth1: in=257.1955 kbps, out=0.0000 kbps
I0401 10:52:02.213746    1430 linux.go:183] eth2: max counter seen = 3167578447, max counter guess = 4294967296
I0401 10:52:02.213792    1430 linux.go:211] eth2: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.213833    1430 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 10:52:02.213876    1430 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:02.213917    1430 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 10:52:02.213959    1430 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I0401 10:52:03.214708    1430 linux.go:231] found device eth0
I0401 10:52:03.214786    1430 linux.go:236] bytesRecvStr for eth0: 2230188637
I0401 10:52:03.214835    1430 linux.go:237] bytesTransStr for eth0: 2181367884
I0401 10:52:03.214893    1430 linux.go:231] found device eth1
I0401 10:52:03.214951    1430 linux.go:236] bytesRecvStr for eth1: 3579676994814
I0401 10:52:03.214997    1430 linux.go:237] bytesTransStr for eth1: 3047499649064
I0401 10:52:03.215051    1430 linux.go:231] found device eth2
I0401 10:52:03.215095    1430 linux.go:236] bytesRecvStr for eth2: 2838333791
I0401 10:52:03.215139    1430 linux.go:237] bytesTransStr for eth2: 3167578663
I0401 10:52:03.215190    1430 linux.go:231] found device wlan0
I0401 10:52:03.215233    1430 linux.go:236] bytesRecvStr for wlan0: 59937
I0401 10:52:03.215276    1430 linux.go:237] bytesTransStr for wlan0: 33550
I0401 10:52:03.215325    1430 linux.go:231] found device lo
I0401 10:52:03.215368    1430 linux.go:236] bytesRecvStr for lo: 349458
I0401 10:52:03.215414    1430 linux.go:237] bytesTransStr for lo: 349458
I0401 10:52:03.215526    1430 linux.go:124] updating state for eth0
I0401 10:52:03.215572    1430 linux.go:124] updating state for eth1
I0401 10:52:03.215610    1430 linux.go:124] updating state for eth2
I0401 10:52:03.215649    1430 linux.go:124] updating state for wlan0
I0401 10:52:03.215687    1430 linux.go:124] updating state for lo
I0401 10:52:03.215748    1430 linux.go:183] eth0: max counter seen = 2230188637, max counter guess = 4294967296
I0401 10:52:03.215819    1430 linux.go:211] eth0: in=649.4125 kbps, out=2707.3186 kbps
I0401 10:52:03.215913    1430 linux.go:183] eth1: max counter seen = 3579676994814, max counter guess = 18446744069414584320
I0401 10:52:03.215996    1430 linux.go:211] eth1: in=2732.8084 kbps, out=661.4013 kbps
I0401 10:52:03.216083    1430 linux.go:183] eth2: max counter seen = 3167578663, max counter guess = 4294967296
I0401 10:52:03.216141    1430 linux.go:211] eth2: in=1.4966 kbps, out=1.6837 kbps
I0401 10:52:03.216225    1430 linux.go:183] lo: max counter seen = 349458, max counter guess = 4294967296
I0401 10:52:03.216282    1430 linux.go:211] lo: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:03.216377    1430 linux.go:183] wlan0: max counter seen = 59937, max counter guess = 4294967296
I0401 10:52:03.216434    1430 linux.go:211] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0401 10:52:03.216525    1430 table.go:58] rows = 40, tableLineCount = 2
I0401 10:52:03.216572    1430 table.go:69] tableLineCount = 2, rows-3 = 37
     649.41     2707.32    2732.81      661.40       1.50        1.68       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
