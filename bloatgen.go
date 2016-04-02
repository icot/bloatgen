package main

import (
    "flag"
    "fmt"
    "github.com/icot/bloatgen/mydb"
    "github.com/icot/bloatgen/bloatsim"
)

var connections, inserts, updates int

func init() {
    flag.IntVar(&connections, "connections", 1, "Number of parallel connections")
    flag.Parse()
}

func main() {
    fmt.Println("# of Connections: ", connections)
	//bloatsim.Simulate(2, 0.1, 0.05, 0.2)
    mydb.Insert("t_char", bloatsim.Data())
    mydb.Insert("t_char", bloatsim.Data())
    mydb.Insert("t_char", bloatsim.Data())
    mydb.Test()
    mydb.Delete("t_char")
    mydb.Update("t_char", bloatsim.Data())
    mydb.Update("t_char", bloatsim.Data())
    mydb.Delete("t_char")
    mydb.Test()
}
