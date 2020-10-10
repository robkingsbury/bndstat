// Bndstat displays simple network device throughput data, inspired by unix
// tools such as vmstat, iostat, mpstat, netstat, etc.
package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/kr/pretty"
	"github.com/robkingsbury/bndstat/throughput"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Info("Starting")

	s := throughput.Stat{}
	fmt.Printf("%s\n", pretty.Sprint(s))
}
