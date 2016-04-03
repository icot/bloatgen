package main

import (
    "flag"
    "fmt"
    "github.com/icot/bloatgen/bloatsim"
)

var duration, connections, frequency int64
var inserts, updates float64

func init() {
    flag.Int64Var(&connections, "conns", 1, "Number of parallel connections")
    flag.Int64Var(&duration, "duration", 1, "Test duration in minutes")
    flag.Int64Var(&frequency, "freq", 10, "Frequency of stat collection in seconds (s)")
    flag.Float64Var(&inserts, "inserts", 0.5, "Insert ratio (%)")
    flag.Float64Var(&updates, "updates", 0.4, "Update ratio (%)")
    flag.Parse()
}

func main() {
    fmt.Println("# of Connections: ", connections)
	bloatsim.Simulate(duration, frequency, inserts, updates)
}
