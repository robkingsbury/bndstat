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
bndstat v0.5.2
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: 4464e1a (v0.5.2)
Compiled: Sun  4 Apr 11:53:52 PDT 2021
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
     646.86     2583.75    2587.42      659.95       2.08        2.42       0.00        0.00
     590.66       45.97      71.49      596.93       1.50        1.69       0.00        0.00
     886.32       39.89      62.60      893.71       1.50        1.69       0.00        0.00
     925.19       36.43      62.99      933.04       1.50        1.69       0.00        0.00
     720.68      658.25     674.82      728.81       1.50        1.69       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      49.41      859.14       1.50        1.69
      64.74      235.53       1.50        1.69
      70.50      264.94       1.50        1.69
     149.07     1685.54       1.50        1.69
    8392.76     1216.45       1.50        1.69
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I0404 11:54:13.600845   31849 bndstat.go:101] interval = 1.000000, count = 1
I0404 11:54:13.601278   31849 throughput.go:21] os is "linux"
I0404 11:54:13.601301   31849 throughput.go:33] running Reporter.Report() twice to prime the stats
I0404 11:54:13.601320   31849 throughput.go:35] prime 1
I0404 11:54:13.601525   31849 linux.go:219] found device eth0
I0404 11:54:13.601548   31849 linux.go:224] bytesRecvStr for eth0: 3204929122
I0404 11:54:13.601567   31849 linux.go:225] bytesTransStr for eth0: 893012390
I0404 11:54:13.601589   31849 linux.go:219] found device eth1
I0404 11:54:13.601607   31849 linux.go:224] bytesRecvStr for eth1: 1990960732599
I0404 11:54:13.601625   31849 linux.go:225] bytesTransStr for eth1: 1019348949765
I0404 11:54:13.601646   31849 linux.go:219] found device eth2
I0404 11:54:13.601663   31849 linux.go:224] bytesRecvStr for eth2: 3021473851
I0404 11:54:13.601681   31849 linux.go:225] bytesTransStr for eth2: 3855442045
I0404 11:54:13.601703   31849 linux.go:219] found device wlan0
I0404 11:54:13.601721   31849 linux.go:224] bytesRecvStr for wlan0: 529377
I0404 11:54:13.601738   31849 linux.go:225] bytesTransStr for wlan0: 1341066
I0404 11:54:13.601760   31849 linux.go:219] found device lo
I0404 11:54:13.601777   31849 linux.go:224] bytesRecvStr for lo: 2766406
I0404 11:54:13.601794   31849 linux.go:225] bytesTransStr for lo: 2766406
I0404 11:54:13.601818   31849 linux.go:122] updating state for eth0
I0404 11:54:13.601902   31849 linux.go:122] updating state for eth1
I0404 11:54:13.601923   31849 linux.go:122] updating state for eth2
I0404 11:54:13.601959   31849 linux.go:122] updating state for wlan0
I0404 11:54:13.601978   31849 linux.go:122] updating state for lo
I0404 11:54:13.602000   31849 linux.go:171] eth2: max counter seen = 3855442045, max counter guess = 4294967296
I0404 11:54:13.602034   31849 linux.go:199] eth2: in=0.0026 kbps, out=0.0033 kbps
I0404 11:54:13.602072   31849 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:54:13.602108   31849 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.602143   31849 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:54:13.602168   31849 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.602201   31849 linux.go:171] eth0: max counter seen = 3204929122, max counter guess = 4294967296
I0404 11:54:13.602227   31849 linux.go:199] eth0: in=0.0027 kbps, out=0.0008 kbps
I0404 11:54:13.602261   31849 linux.go:171] eth1: max counter seen = 1990960732599, max counter guess = 18446744069414584320
I0404 11:54:13.602296   31849 linux.go:199] eth1: in=1.6864 kbps, out=0.8634 kbps
I0404 11:54:13.602346   31849 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:54:13.602372   31849 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.602405   31849 linux.go:171] eth0: max counter seen = 3204929122, max counter guess = 4294967296
I0404 11:54:13.602429   31849 linux.go:199] eth0: in=0.0027 kbps, out=0.0008 kbps
I0404 11:54:13.602463   31849 linux.go:171] eth1: max counter seen = 1990960732599, max counter guess = 18446744069414584320
I0404 11:54:13.602498   31849 linux.go:199] eth1: in=1.6864 kbps, out=0.8634 kbps
I0404 11:54:13.602530   31849 linux.go:171] eth2: max counter seen = 3855442045, max counter guess = 4294967296
I0404 11:54:13.602555   31849 linux.go:199] eth2: in=0.0026 kbps, out=0.0033 kbps
I0404 11:54:13.602589   31849 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:54:13.602613   31849 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.602648   31849 throughput.go:38] prime 2
I0404 11:54:13.602760   31849 linux.go:219] found device eth0
I0404 11:54:13.602780   31849 linux.go:224] bytesRecvStr for eth0: 3204930596
I0404 11:54:13.602798   31849 linux.go:225] bytesTransStr for eth0: 893012390
I0404 11:54:13.602820   31849 linux.go:219] found device eth1
I0404 11:54:13.602838   31849 linux.go:224] bytesRecvStr for eth1: 1990960732599
I0404 11:54:13.602855   31849 linux.go:225] bytesTransStr for eth1: 1019348952729
I0404 11:54:13.602876   31849 linux.go:219] found device eth2
I0404 11:54:13.602893   31849 linux.go:224] bytesRecvStr for eth2: 3021473851
I0404 11:54:13.602910   31849 linux.go:225] bytesTransStr for eth2: 3855442045
I0404 11:54:13.602958   31849 linux.go:219] found device wlan0
I0404 11:54:13.602995   31849 linux.go:224] bytesRecvStr for wlan0: 529377
I0404 11:54:13.603013   31849 linux.go:225] bytesTransStr for wlan0: 1341066
I0404 11:54:13.603052   31849 linux.go:219] found device lo
I0404 11:54:13.603070   31849 linux.go:224] bytesRecvStr for lo: 2766406
I0404 11:54:13.603088   31849 linux.go:225] bytesTransStr for lo: 2766406
I0404 11:54:13.603110   31849 linux.go:122] updating state for eth0
I0404 11:54:13.603127   31849 linux.go:122] updating state for eth1
I0404 11:54:13.603161   31849 linux.go:122] updating state for eth2
I0404 11:54:13.603179   31849 linux.go:122] updating state for wlan0
I0404 11:54:13.603195   31849 linux.go:122] updating state for lo
I0404 11:54:13.603215   31849 linux.go:171] eth0: max counter seen = 3204930596, max counter guess = 4294967296
I0404 11:54:13.603238   31849 linux.go:199] eth0: in=8904.9033 kbps, out=0.0000 kbps
I0404 11:54:13.603267   31849 linux.go:171] eth1: max counter seen = 1990960732599, max counter guess = 18446744069414584320
I0404 11:54:13.603298   31849 linux.go:199] eth1: in=0.0000 kbps, out=17906.4676 kbps
I0404 11:54:13.603326   31849 linux.go:171] eth2: max counter seen = 3855442045, max counter guess = 4294967296
I0404 11:54:13.603347   31849 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.603368   31849 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:54:13.603389   31849 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.603409   31849 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:54:13.603430   31849 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.603458   31849 linux.go:171] eth1: max counter seen = 1990960732599, max counter guess = 18446744069414584320
I0404 11:54:13.603490   31849 linux.go:199] eth1: in=0.0000 kbps, out=17906.4676 kbps
I0404 11:54:13.603517   31849 linux.go:171] eth2: max counter seen = 3855442045, max counter guess = 4294967296
I0404 11:54:13.603538   31849 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.603558   31849 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:54:13.603579   31849 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.603600   31849 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:54:13.603620   31849 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.603640   31849 linux.go:171] eth0: max counter seen = 3204930596, max counter guess = 4294967296
I0404 11:54:13.603662   31849 linux.go:199] eth0: in=8904.9033 kbps, out=0.0000 kbps
I0404 11:54:13.603788   31849 linux.go:219] found device eth0
I0404 11:54:13.603808   31849 linux.go:224] bytesRecvStr for eth0: 3204930596
I0404 11:54:13.603826   31849 linux.go:225] bytesTransStr for eth0: 893012390
I0404 11:54:13.603847   31849 linux.go:219] found device eth1
I0404 11:54:13.603864   31849 linux.go:224] bytesRecvStr for eth1: 1990960732599
I0404 11:54:13.603882   31849 linux.go:225] bytesTransStr for eth1: 1019348952729
I0404 11:54:13.603902   31849 linux.go:219] found device eth2
I0404 11:54:13.603920   31849 linux.go:224] bytesRecvStr for eth2: 3021473851
I0404 11:54:13.603937   31849 linux.go:225] bytesTransStr for eth2: 3855442045
I0404 11:54:13.603958   31849 linux.go:219] found device wlan0
I0404 11:54:13.603975   31849 linux.go:224] bytesRecvStr for wlan0: 529377
I0404 11:54:13.603992   31849 linux.go:225] bytesTransStr for wlan0: 1341066
I0404 11:54:13.604020   31849 linux.go:219] found device lo
I0404 11:54:13.604038   31849 linux.go:224] bytesRecvStr for lo: 2766406
I0404 11:54:13.604055   31849 linux.go:225] bytesTransStr for lo: 2766406
I0404 11:54:13.604077   31849 linux.go:122] updating state for eth0
I0404 11:54:13.604094   31849 linux.go:122] updating state for eth1
I0404 11:54:13.604110   31849 linux.go:122] updating state for eth2
I0404 11:54:13.604126   31849 linux.go:122] updating state for wlan0
I0404 11:54:13.604142   31849 linux.go:122] updating state for lo
I0404 11:54:13.604160   31849 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:54:13.604182   31849 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.604203   31849 linux.go:171] eth0: max counter seen = 3204930596, max counter guess = 4294967296
I0404 11:54:13.604225   31849 linux.go:199] eth0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.604245   31849 linux.go:171] eth1: max counter seen = 1990960732599, max counter guess = 18446744069414584320
I0404 11:54:13.604276   31849 linux.go:199] eth1: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.604296   31849 linux.go:171] eth2: max counter seen = 3855442045, max counter guess = 4294967296
I0404 11:54:13.604317   31849 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.604337   31849 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:54:13.604358   31849 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.604386   31849 linux.go:171] eth0: max counter seen = 3204930596, max counter guess = 4294967296
I0404 11:54:13.604408   31849 linux.go:199] eth0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.604428   31849 linux.go:171] eth1: max counter seen = 1990960732599, max counter guess = 18446744069414584320
I0404 11:54:13.604458   31849 linux.go:199] eth1: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.604479   31849 linux.go:171] eth2: max counter seen = 3855442045, max counter guess = 4294967296
I0404 11:54:13.604500   31849 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.604521   31849 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:54:13.604542   31849 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:13.604562   31849 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:54:13.604582   31849 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I0404 11:54:14.605299   31849 linux.go:219] found device eth0
I0404 11:54:14.605373   31849 linux.go:224] bytesRecvStr for eth0: 3205002049
I0404 11:54:14.605427   31849 linux.go:225] bytesTransStr for eth0: 893036226
I0404 11:54:14.605486   31849 linux.go:219] found device eth1
I0404 11:54:14.605533   31849 linux.go:224] bytesRecvStr for eth1: 1990960761775
I0404 11:54:14.605577   31849 linux.go:225] bytesTransStr for eth1: 1019349025266
I0404 11:54:14.605672   31849 linux.go:219] found device eth2
I0404 11:54:14.605724   31849 linux.go:224] bytesRecvStr for eth2: 3021474043
I0404 11:54:14.605769   31849 linux.go:225] bytesTransStr for eth2: 3855442261
I0404 11:54:14.605825   31849 linux.go:219] found device wlan0
I0404 11:54:14.605868   31849 linux.go:224] bytesRecvStr for wlan0: 529377
I0404 11:54:14.605911   31849 linux.go:225] bytesTransStr for wlan0: 1341066
I0404 11:54:14.605963   31849 linux.go:219] found device lo
I0404 11:54:14.606006   31849 linux.go:224] bytesRecvStr for lo: 2766406
I0404 11:54:14.606050   31849 linux.go:225] bytesTransStr for lo: 2766406
I0404 11:54:14.606108   31849 linux.go:122] updating state for eth0
I0404 11:54:14.606151   31849 linux.go:122] updating state for eth1
I0404 11:54:14.606192   31849 linux.go:122] updating state for eth2
I0404 11:54:14.606232   31849 linux.go:122] updating state for wlan0
I0404 11:54:14.606269   31849 linux.go:122] updating state for lo
I0404 11:54:14.606321   31849 linux.go:171] eth0: max counter seen = 3205002049, max counter guess = 4294967296
I0404 11:54:14.606387   31849 linux.go:199] eth0: in=557.1011 kbps, out=185.8433 kbps
I0404 11:54:14.606479   31849 linux.go:171] eth1: max counter seen = 1990960761775, max counter guess = 18446744069414584320
I0404 11:54:14.606560   31849 linux.go:199] eth1: in=227.4779 kbps, out=565.5528 kbps
I0404 11:54:14.606646   31849 linux.go:171] eth2: max counter seen = 3855442261, max counter guess = 4294967296
I0404 11:54:14.606702   31849 linux.go:199] eth2: in=1.4970 kbps, out=1.6841 kbps
I0404 11:54:14.606787   31849 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:54:14.606842   31849 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:14.606894   31849 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:54:14.606947   31849 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:14.607021   31849 linux.go:171] lo: max counter seen = 2766406, max counter guess = 4294967296
I0404 11:54:14.607075   31849 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:14.607127   31849 linux.go:171] eth0: max counter seen = 3205002049, max counter guess = 4294967296
I0404 11:54:14.607180   31849 linux.go:199] eth0: in=557.1011 kbps, out=185.8433 kbps
I0404 11:54:14.607264   31849 linux.go:171] eth1: max counter seen = 1990960761775, max counter guess = 18446744069414584320
I0404 11:54:14.607343   31849 linux.go:199] eth1: in=227.4779 kbps, out=565.5528 kbps
I0404 11:54:14.607426   31849 linux.go:171] eth2: max counter seen = 3855442261, max counter guess = 4294967296
I0404 11:54:14.607480   31849 linux.go:199] eth2: in=1.4970 kbps, out=1.6841 kbps
I0404 11:54:14.607563   31849 linux.go:171] wlan0: max counter seen = 1341066, max counter guess = 4294967296
I0404 11:54:14.607617   31849 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I0404 11:54:14.607690   31849 table.go:58] rows = 40, tableLineCount = 2
I0404 11:54:14.607736   31849 table.go:69] tableLineCount = 2, rows-3 = 37
     557.10      185.84     227.48      565.55       1.50        1.68       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
