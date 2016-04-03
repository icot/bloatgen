package main

import (
    "flag"
    "fmt"
    "github.com/icot/bloatgen/bloatsim"
)

var connections, inserts, updates int

func init() {
    flag.IntVar(&connections, "connections", 1, "Number of parallel connections")
    flag.Parse()
}

func main() {
    fmt.Println("# of Connections: ", connections)
	bloatsim.Simulate(60, 0.3, 0.1, 0.6)
}
