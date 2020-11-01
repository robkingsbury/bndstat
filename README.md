# bndstat
[![Actions Status](https://github.com/robkingsbury/bndstat/workflows/Build/badge.svg)](https://github.com/robkingsbury/bndstat/actions)
[![Actions Status](https://github.com/robkingsbury/bndstat/workflows/Test/badge.svg)](https://github.com/robkingsbury/bndstat/actions)

A simple Go program that displays throughput stats for network interfaces.

## Quickstart

```
$ git clone https://github.com/robkingsbury/bndstat
$ cd bndstat
$ go build
$ ./bndstat
$ ./bndstat --help
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

### Example

In this example, running on a raspberry pi that I am using as a router, eth0 is the hardline connection to wifi router, eth1 is my
primary internet provider, eth2 is my backup internet line and wlan0 is a connection to the wifi network:

```
$ bndstat 3 5
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
    1012.81       56.99      50.69     1019.71       0.25        0.53       0.00        0.00
    1108.78       48.27      44.04     1115.83       0.37        0.42       0.19        0.00
    1145.63       75.78      70.08     1153.09       0.37        0.42       0.00        0.00
     833.05      104.24      98.51      840.39       0.37        0.42       1.61        0.00
    1397.95       81.43      73.27     1407.82       0.50        0.55       0.20        0.00
```

Another example on the same machine, illustrating the device filter and using options instead of args for the interval and
count parameters:

```
$ bndstat --devices=eth1,eth2 --interval=1 --count=5
              eth1                   eth2     
         In         Out         In         Out
      48.59     1836.57       0.37        0.42
      20.58      665.11       0.37        0.42
      24.39      912.66       0.37        0.42
      23.79     1321.98       0.37        0.42
     762.78     1481.80       0.37        0.42
```

### Debug Logging
If you want to see the innerworkings of `bndstat`, you can use options from the standard Go [glog package](https://github.com/golang/glog). For example:

 ```
 $ bndstat --logtostderr --v=2 --count=1
I1025 16:04:19.351979    8592 bndstat.go:86] interval = 1.000000, count = 1
I1025 16:04:19.352790    8592 throughput.go:21] os is "linux"
I1025 16:04:19.353267    8592 linux.go:111] found device eth0
I1025 16:04:19.353331    8592 linux.go:116] bytesRecvStr for eth0: 1535674097
I1025 16:04:19.353382    8592 linux.go:117] bytesTransStr for eth0: 2934782594
I1025 16:04:19.353442    8592 linux.go:111] found device eth1
I1025 16:04:19.353491    8592 linux.go:116] bytesRecvStr for eth1: 304074782568
I1025 16:04:19.353541    8592 linux.go:117] bytesTransStr for eth1: 213873912037
I1025 16:04:19.353596    8592 linux.go:111] found device eth2
I1025 16:04:19.353640    8592 linux.go:116] bytesRecvStr for eth2: 236114548
I1025 16:04:19.353684    8592 linux.go:117] bytesTransStr for eth2: 296118878
I1025 16:04:19.353739    8592 linux.go:111] found device wlan0
I1025 16:04:19.353784    8592 linux.go:116] bytesRecvStr for wlan0: 2688856353
I1025 16:04:19.353834    8592 linux.go:117] bytesTransStr for wlan0: 99335270
I1025 16:04:19.353891    8592 linux.go:111] found device lo
I1025 16:04:19.353935    8592 linux.go:116] bytesRecvStr for lo: 69833
I1025 16:04:19.353980    8592 linux.go:117] bytesTransStr for lo: 69833
I1025 16:04:19.354280    8592 linux.go:111] found device eth0
I1025 16:04:19.354331    8592 linux.go:116] bytesRecvStr for eth0: 1535674097
I1025 16:04:19.354376    8592 linux.go:117] bytesTransStr for eth0: 2934782828
I1025 16:04:19.354442    8592 linux.go:111] found device eth1
I1025 16:04:19.354491    8592 linux.go:116] bytesRecvStr for eth1: 304074782568
I1025 16:04:19.354536    8592 linux.go:117] bytesTransStr for eth1: 213873912037
I1025 16:04:19.354589    8592 linux.go:111] found device eth2
I1025 16:04:19.354635    8592 linux.go:116] bytesRecvStr for eth2: 236114548
I1025 16:04:19.354682    8592 linux.go:117] bytesTransStr for eth2: 296118878
I1025 16:04:19.354736    8592 linux.go:111] found device wlan0
I1025 16:04:19.354780    8592 linux.go:116] bytesRecvStr for wlan0: 2688856353
I1025 16:04:19.354846    8592 linux.go:117] bytesTransStr for wlan0: 99335270
I1025 16:04:19.354900    8592 linux.go:111] found device lo
I1025 16:04:19.354943    8592 linux.go:116] bytesRecvStr for lo: 69833
I1025 16:04:19.354988    8592 linux.go:117] bytesTransStr for lo: 69833
I1025 16:04:19.355269    8592 linux.go:111] found device eth0
I1025 16:04:19.355319    8592 linux.go:116] bytesRecvStr for eth0: 1535674097
I1025 16:04:19.355363    8592 linux.go:117] bytesTransStr for eth0: 2934783430
I1025 16:04:19.355419    8592 linux.go:111] found device eth1
I1025 16:04:19.355469    8592 linux.go:116] bytesRecvStr for eth1: 304074782568
I1025 16:04:19.355514    8592 linux.go:117] bytesTransStr for eth1: 213873912037
I1025 16:04:19.355566    8592 linux.go:111] found device eth2
I1025 16:04:19.355610    8592 linux.go:116] bytesRecvStr for eth2: 236114548
I1025 16:04:19.355658    8592 linux.go:117] bytesTransStr for eth2: 296118878
I1025 16:04:19.355711    8592 linux.go:111] found device wlan0
I1025 16:04:19.355755    8592 linux.go:116] bytesRecvStr for wlan0: 2688856353
I1025 16:04:19.355799    8592 linux.go:117] bytesTransStr for wlan0: 99335270
I1025 16:04:19.355851    8592 linux.go:111] found device lo
I1025 16:04:19.355895    8592 linux.go:116] bytesRecvStr for lo: 69833
I1025 16:04:19.355938    8592 linux.go:117] bytesTransStr for lo: 69833
              eth0                   eth1                   eth2                  wlan0     
         In         Out         In         Out         In         Out         In         Out
I1025 16:04:20.356733    8592 linux.go:111] found device eth0
I1025 16:04:20.356814    8592 linux.go:116] bytesRecvStr for eth0: 1535894555
I1025 16:04:20.356868    8592 linux.go:117] bytesTransStr for eth0: 2935092996
I1025 16:04:20.356927    8592 linux.go:111] found device eth1
I1025 16:04:20.356973    8592 linux.go:116] bytesRecvStr for eth1: 304075086404
I1025 16:04:20.357019    8592 linux.go:117] bytesTransStr for eth1: 213874134335
I1025 16:04:20.357078    8592 linux.go:111] found device eth2
I1025 16:04:20.357128    8592 linux.go:116] bytesRecvStr for eth2: 236114596
I1025 16:04:20.357172    8592 linux.go:117] bytesTransStr for eth2: 296118932
I1025 16:04:20.357228    8592 linux.go:111] found device wlan0
I1025 16:04:20.357272    8592 linux.go:116] bytesRecvStr for wlan0: 2688856353
I1025 16:04:20.357317    8592 linux.go:117] bytesTransStr for wlan0: 99335270
I1025 16:04:20.357374    8592 linux.go:111] found device lo
I1025 16:04:20.357422    8592 linux.go:116] bytesRecvStr for lo: 69833
I1025 16:04:20.357472    8592 linux.go:117] bytesTransStr for lo: 69833
I1025 16:04:20.357719    8592 table.go:53] columns = 243, rows = 88, tableLineCount = 2
I1025 16:04:20.357784    8592 table.go:65] tableLineCount = 2, rows-3 = 85
    1719.67     2414.75    2370.05     1734.02       0.37        0.42       0.00        0.00
```

## Supported Platforms

As of v0.4.0, only Linux is supported. The Linux library relies on information from `/proc/net/dev` so it *should* work on most Linux systems. Very long device names would probably make the output look a little wonky since the column width is static right now.
