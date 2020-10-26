// Package throughput provides an interface and implementations of that
// interface to read the throughput of network devices.
//
// Synopsis
//
// A brief example of using the package to print throughput stats:
//   package main
//
//   import (
//     "fmt"
//     "github.com/robkingsbury/bndstat/throughput"
//     "time"
//   )
//
//   func main() {
//     // Throwing away errors is bad practice but is done here for brevity.
//     reporter, _ := throughput.NewReporter()
//     table := throughput.NewTable()
//
//     for {
//       stats, _ := reporter.Report()
//
//       // Directly accessing device stats
//       for _, device := range stats.Devices() {
//         in, out := stats.Avg(device, throughput.Kbps)
//         fmt.Printf("%s: in = %.2f, out = %.2f\n", device, in, out)
//       }
//
//       // Using the builtin, aligned table output
//       table.Write(throughput.TableWriteOpts{
//         Stats:    stats,
//         Devices:  stats.Devices(),
//         Unit:     throughput.Kbps,
//         ShowUnit: false,
//       })
//
//       time.Sleep(time.Second)
//     }
//   }
package throughput
