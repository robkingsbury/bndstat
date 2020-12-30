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
bndstat v0.5.0
Rob Kingsbury
https://github.com/robkingsbury/bndstat
Commit: e32bc6e (v0.5.0)
Compiled: Wed 30 Dec 08:51:06 PST 2020
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
     701.29       42.92      79.81      712.80       0.37        0.28       0.00        0.00
     772.66      342.74     374.51      781.88       0.37        0.42       0.00        0.00
     695.81       26.66      58.63      704.15       0.37        0.42       0.00        0.00
     922.46       34.48      66.02      935.52       0.62        0.68       0.00        0.00
     724.73       43.31      76.06      730.33       0.37        0.42       0.00        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      53.10      930.95       0.95        1.15
      54.88      863.91       0.95        1.16
     105.97     1047.70       0.37        0.42
     438.00      255.29       0.37        0.42
      61.69      885.49       0.37        0.42
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

```
$ bndstat --logtostderr --v=2 --count=1
I1230 08:51:27.553837    1794 bndstat.go:101] interval = 1.000000, count = 1
I1230 08:51:27.554634    1794 throughput.go:21] os is "linux"
I1230 08:51:27.554689    1794 throughput.go:33] running Reporter.Report() twice to prime the stats
I1230 08:51:27.554733    1794 throughput.go:35] prime 1
I1230 08:51:27.555147    1794 linux.go:219] found device eth0
I1230 08:51:27.555199    1794 linux.go:224] bytesRecvStr for eth0: 3465594967
I1230 08:51:27.555246    1794 linux.go:225] bytesTransStr for eth0: 4001197974
I1230 08:51:27.555300    1794 linux.go:219] found device eth1
I1230 08:51:27.555342    1794 linux.go:224] bytesRecvStr for eth1: 12543945220
I1230 08:51:27.555387    1794 linux.go:225] bytesTransStr for eth1: 16457091079
I1230 08:51:27.555440    1794 linux.go:219] found device eth2
I1230 08:51:27.555484    1794 linux.go:224] bytesRecvStr for eth2: 161476369
I1230 08:51:27.555527    1794 linux.go:225] bytesTransStr for eth2: 193676395
I1230 08:51:27.555581    1794 linux.go:219] found device wlan0
I1230 08:51:27.555625    1794 linux.go:224] bytesRecvStr for wlan0: 0
I1230 08:51:27.555668    1794 linux.go:225] bytesTransStr for wlan0: 0
I1230 08:51:27.555719    1794 linux.go:219] found device lo
I1230 08:51:27.555763    1794 linux.go:224] bytesRecvStr for lo: 8218
I1230 08:51:27.555806    1794 linux.go:225] bytesTransStr for lo: 8218
I1230 08:51:27.555863    1794 linux.go:122] updating state for eth0
I1230 08:51:27.555908    1794 linux.go:122] updating state for eth1
I1230 08:51:27.555951    1794 linux.go:122] updating state for eth2
I1230 08:51:27.555993    1794 linux.go:122] updating state for wlan0
I1230 08:51:27.556034    1794 linux.go:122] updating state for lo
I1230 08:51:27.556081    1794 linux.go:171] eth0: max counter seen = 4001197974, max counter guess = 4294967296
I1230 08:51:27.556156    1794 linux.go:199] eth0: in=0.0029 kbps, out=0.0034 kbps
I1230 08:51:27.556248    1794 linux.go:171] eth1: max counter seen = 16457091079, max counter guess = 18446744069414584320
I1230 08:51:27.556549    1794 linux.go:199] eth1: in=0.0106 kbps, out=0.0139 kbps
I1230 08:51:27.556652    1794 linux.go:171] eth2: max counter seen = 193676395, max counter guess = 4294967296
I1230 08:51:27.556719    1794 linux.go:199] eth2: in=0.0001 kbps, out=0.0002 kbps
I1230 08:51:27.556806    1794 linux.go:171] wlan0: max counter seen = 0, max counter guess = 4294967296
I1230 08:51:27.556868    1794 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.556920    1794 linux.go:171] lo: max counter seen = 8218, max counter guess = 4294967296
I1230 08:51:27.556982    1794 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.557107    1794 linux.go:171] eth0: max counter seen = 4001197974, max counter guess = 4294967296
I1230 08:51:27.557171    1794 linux.go:199] eth0: in=0.0029 kbps, out=0.0034 kbps
I1230 08:51:27.557256    1794 linux.go:171] eth1: max counter seen = 16457091079, max counter guess = 18446744069414584320
I1230 08:51:27.557339    1794 linux.go:199] eth1: in=0.0106 kbps, out=0.0139 kbps
I1230 08:51:27.557424    1794 linux.go:171] eth2: max counter seen = 193676395, max counter guess = 4294967296
I1230 08:51:27.557486    1794 linux.go:199] eth2: in=0.0001 kbps, out=0.0002 kbps
I1230 08:51:27.557570    1794 linux.go:171] wlan0: max counter seen = 0, max counter guess = 4294967296
I1230 08:51:27.557628    1794 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.557680    1794 linux.go:171] lo: max counter seen = 8218, max counter guess = 4294967296
I1230 08:51:27.557744    1794 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.557827    1794 throughput.go:38] prime 2
I1230 08:51:27.558083    1794 linux.go:219] found device eth0
I1230 08:51:27.558133    1794 linux.go:224] bytesRecvStr for eth0: 3465594967
I1230 08:51:27.558175    1794 linux.go:225] bytesTransStr for eth0: 4001197974
I1230 08:51:27.558230    1794 linux.go:219] found device eth1
I1230 08:51:27.558274    1794 linux.go:224] bytesRecvStr for eth1: 12543945220
I1230 08:51:27.558318    1794 linux.go:225] bytesTransStr for eth1: 16457091079
I1230 08:51:27.558369    1794 linux.go:219] found device eth2
I1230 08:51:27.558412    1794 linux.go:224] bytesRecvStr for eth2: 161476435
I1230 08:51:27.558496    1794 linux.go:225] bytesTransStr for eth2: 193676489
I1230 08:51:27.558553    1794 linux.go:219] found device wlan0
I1230 08:51:27.558597    1794 linux.go:224] bytesRecvStr for wlan0: 0
I1230 08:51:27.558640    1794 linux.go:225] bytesTransStr for wlan0: 0
I1230 08:51:27.558690    1794 linux.go:219] found device lo
I1230 08:51:27.558733    1794 linux.go:224] bytesRecvStr for lo: 8218
I1230 08:51:27.558776    1794 linux.go:225] bytesTransStr for lo: 8218
I1230 08:51:27.558832    1794 linux.go:122] updating state for eth0
I1230 08:51:27.558873    1794 linux.go:122] updating state for eth1
I1230 08:51:27.558913    1794 linux.go:122] updating state for eth2
I1230 08:51:27.558953    1794 linux.go:122] updating state for wlan0
I1230 08:51:27.558990    1794 linux.go:122] updating state for lo
I1230 08:51:27.559036    1794 linux.go:171] eth0: max counter seen = 4001197974, max counter guess = 4294967296
I1230 08:51:27.559093    1794 linux.go:199] eth0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.559147    1794 linux.go:171] eth1: max counter seen = 16457091079, max counter guess = 18446744069414584320
I1230 08:51:27.559221    1794 linux.go:199] eth1: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.559273    1794 linux.go:171] eth2: max counter seen = 193676489, max counter guess = 4294967296
I1230 08:51:27.559327    1794 linux.go:199] eth2: in=173.6412 kbps, out=247.3071 kbps
I1230 08:51:27.559410    1794 linux.go:171] wlan0: max counter seen = 0, max counter guess = 4294967296
I1230 08:51:27.559463    1794 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.559514    1794 linux.go:171] lo: max counter seen = 8218, max counter guess = 4294967296
I1230 08:51:27.559565    1794 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.559635    1794 linux.go:171] eth1: max counter seen = 16457091079, max counter guess = 18446744069414584320
I1230 08:51:27.559710    1794 linux.go:199] eth1: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.559761    1794 linux.go:171] eth2: max counter seen = 193676489, max counter guess = 4294967296
I1230 08:51:27.559814    1794 linux.go:199] eth2: in=173.6412 kbps, out=247.3071 kbps
I1230 08:51:27.559896    1794 linux.go:171] wlan0: max counter seen = 0, max counter guess = 4294967296
I1230 08:51:27.559947    1794 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.559998    1794 linux.go:171] lo: max counter seen = 8218, max counter guess = 4294967296
I1230 08:51:27.560049    1794 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.560099    1794 linux.go:171] eth0: max counter seen = 4001197974, max counter guess = 4294967296
I1230 08:51:27.560152    1794 linux.go:199] eth0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.560434    1794 linux.go:219] found device eth0
I1230 08:51:27.560484    1794 linux.go:224] bytesRecvStr for eth0: 3465594967
I1230 08:51:27.560529    1794 linux.go:225] bytesTransStr for eth0: 4001197974
I1230 08:51:27.560604    1794 linux.go:219] found device eth1
I1230 08:51:27.560650    1794 linux.go:224] bytesRecvStr for eth1: 12543945220
I1230 08:51:27.560694    1794 linux.go:225] bytesTransStr for eth1: 16457091079
I1230 08:51:27.560746    1794 linux.go:219] found device eth2
I1230 08:51:27.560790    1794 linux.go:224] bytesRecvStr for eth2: 161476435
I1230 08:51:27.560835    1794 linux.go:225] bytesTransStr for eth2: 193676489
I1230 08:51:27.560888    1794 linux.go:219] found device wlan0
I1230 08:51:27.560932    1794 linux.go:224] bytesRecvStr for wlan0: 0
I1230 08:51:27.560990    1794 linux.go:225] bytesTransStr for wlan0: 0
I1230 08:51:27.561051    1794 linux.go:219] found device lo
I1230 08:51:27.561100    1794 linux.go:224] bytesRecvStr for lo: 8218
I1230 08:51:27.561148    1794 linux.go:225] bytesTransStr for lo: 8218
I1230 08:51:27.561204    1794 linux.go:122] updating state for eth0
I1230 08:51:27.561245    1794 linux.go:122] updating state for eth1
I1230 08:51:27.561285    1794 linux.go:122] updating state for eth2
I1230 08:51:27.561325    1794 linux.go:122] updating state for wlan0
I1230 08:51:27.561364    1794 linux.go:122] updating state for lo
I1230 08:51:27.561411    1794 linux.go:171] eth0: max counter seen = 4001197974, max counter guess = 4294967296
I1230 08:51:27.561470    1794 linux.go:199] eth0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.561524    1794 linux.go:171] eth1: max counter seen = 16457091079, max counter guess = 18446744069414584320
I1230 08:51:27.561596    1794 linux.go:199] eth1: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.561647    1794 linux.go:171] eth2: max counter seen = 193676489, max counter guess = 4294967296
I1230 08:51:27.561701    1794 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.561752    1794 linux.go:171] wlan0: max counter seen = 0, max counter guess = 4294967296
I1230 08:51:27.561803    1794 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.561855    1794 linux.go:171] lo: max counter seen = 8218, max counter guess = 4294967296
I1230 08:51:27.561907    1794 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.561977    1794 linux.go:171] eth0: max counter seen = 4001197974, max counter guess = 4294967296
I1230 08:51:27.562031    1794 linux.go:199] eth0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.562082    1794 linux.go:171] eth1: max counter seen = 16457091079, max counter guess = 18446744069414584320
I1230 08:51:27.562152    1794 linux.go:199] eth1: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.562203    1794 linux.go:171] eth2: max counter seen = 193676489, max counter guess = 4294967296
I1230 08:51:27.562255    1794 linux.go:199] eth2: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.562306    1794 linux.go:171] wlan0: max counter seen = 0, max counter guess = 4294967296
I1230 08:51:27.562393    1794 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:27.562447    1794 linux.go:171] lo: max counter seen = 8218, max counter guess = 4294967296
I1230 08:51:27.562501    1794 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1230 08:51:28.563567    1794 linux.go:219] found device eth0
I1230 08:51:28.563744    1794 linux.go:224] bytesRecvStr for eth0: 3465664716
I1230 08:51:28.563863    1794 linux.go:225] bytesTransStr for eth0: 4001199997
I1230 08:51:28.563989    1794 linux.go:219] found device eth1
I1230 08:51:28.564099    1794 linux.go:224] bytesRecvStr for eth1: 12543953047
I1230 08:51:28.564211    1794 linux.go:225] bytesTransStr for eth1: 16457162919
I1230 08:51:28.564330    1794 linux.go:219] found device eth2
I1230 08:51:28.564441    1794 linux.go:224] bytesRecvStr for eth2: 161476483
I1230 08:51:28.564552    1794 linux.go:225] bytesTransStr for eth2: 193676543
I1230 08:51:28.564696    1794 linux.go:219] found device wlan0
I1230 08:51:28.564842    1794 linux.go:224] bytesRecvStr for wlan0: 0
I1230 08:51:28.564951    1794 linux.go:225] bytesTransStr for wlan0: 0
I1230 08:51:28.565068    1794 linux.go:219] found device lo
I1230 08:51:28.565173    1794 linux.go:224] bytesRecvStr for lo: 8218
I1230 08:51:28.565282    1794 linux.go:225] bytesTransStr for lo: 8218
I1230 08:51:28.565407    1794 linux.go:122] updating state for eth0
I1230 08:51:28.565517    1794 linux.go:122] updating state for eth1
I1230 08:51:28.565622    1794 linux.go:122] updating state for eth2
I1230 08:51:28.565728    1794 linux.go:122] updating state for wlan0
I1230 08:51:28.565832    1794 linux.go:122] updating state for lo
I1230 08:51:28.565977    1794 linux.go:171] lo: max counter seen = 8218, max counter guess = 4294967296
I1230 08:51:28.566111    1794 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:28.566280    1794 linux.go:171] eth0: max counter seen = 4001199997, max counter guess = 4294967296
I1230 08:51:28.566436    1794 linux.go:199] eth0: in=542.6347 kbps, out=15.7386 kbps
I1230 08:51:28.566630    1794 linux.go:171] eth1: max counter seen = 16457162919, max counter guess = 18446744069414584320
I1230 08:51:28.566817    1794 linux.go:199] eth1: in=60.8927 kbps, out=558.9023 kbps
I1230 08:51:28.567005    1794 linux.go:171] eth2: max counter seen = 193676543, max counter guess = 4294967296
I1230 08:51:28.567131    1794 linux.go:199] eth2: in=0.3734 kbps, out=0.4201 kbps
I1230 08:51:28.567318    1794 linux.go:171] wlan0: max counter seen = 0, max counter guess = 4294967296
I1230 08:51:28.567441    1794 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:28.567585    1794 linux.go:171] eth0: max counter seen = 4001199997, max counter guess = 4294967296
I1230 08:51:28.567752    1794 linux.go:199] eth0: in=542.6347 kbps, out=15.7386 kbps
I1230 08:51:28.567959    1794 linux.go:171] eth1: max counter seen = 16457162919, max counter guess = 18446744069414584320
I1230 08:51:28.568141    1794 linux.go:199] eth1: in=60.8927 kbps, out=558.9023 kbps
I1230 08:51:28.568326    1794 linux.go:171] eth2: max counter seen = 193676543, max counter guess = 4294967296
I1230 08:51:28.568450    1794 linux.go:199] eth2: in=0.3734 kbps, out=0.4201 kbps
I1230 08:51:28.568639    1794 linux.go:171] wlan0: max counter seen = 0, max counter guess = 4294967296
I1230 08:51:28.568765    1794 linux.go:199] wlan0: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:28.568853    1794 linux.go:171] lo: max counter seen = 8218, max counter guess = 4294967296
I1230 08:51:28.568980    1794 linux.go:199] lo: in=0.0000 kbps, out=0.0000 kbps
I1230 08:51:28.569089    1794 table.go:58] rows = 40, tableLineCount = 2
I1230 08:51:28.569205    1794 table.go:69] tableLineCount = 2, rows-3 = 37
     542.63       15.74      60.89      558.90       0.37        0.42       0.00        0.00
```

## Throughput Package

Device stats are available programmatically via the *throughput* package. See http://godoc.org/github.com/robkingsbury/bndstat/throughput for the GoDoc package documentation.

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
