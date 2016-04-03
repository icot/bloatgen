package bloatsim

import (
    "time"
	"math/rand"
    "fmt"
    "github.com/icot/bloatgen/mydb"
)

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

func Data() string {
	n := 32 + rand.Int31n(224)
    b := make([]byte, n)
    for i := range b {
        b[i] = 'A'
    }
    return string(b)
}

func activity(c chan int32, inserts, updates float64) {
    var counter int32  = 0
    for {
        r := rand.Float64()
        data := Data()
        switch {
            case r <= inserts:
                mydb.Insert("t_char", data)
                mydb.Insert("t_text", data)
                mydb.Insert("t_varchar", data)
            case r > inserts && r <= inserts + updates:
                id := mydb.RandomId("t_char")
                mydb.Update("t_char", id, data)
                mydb.Update("t_text", id, data)
                mydb.Update("t_varchar", id, data)
            default:
                id := mydb.RandomId("t_char")
                mydb.Delete("t_char", id)
                mydb.Delete("t_text", id)
                mydb.Delete("t_varchar", id)
        }
        if counter >= 1000 {
            c <- counter
            counter = 0
        }
        counter++
    }
}

func Simulate(duration, frequency int64, inserts, updates float64) {
    stats:= make(chan int32, 1)
    ticker_chan := time.NewTicker(time.Second * time.Duration(frequency)).C
    timer_chan := time.NewTimer(time.Duration(duration) * time.Minute).C
    go activity(stats, inserts, updates)
    for {
        select {
            case <- timer_chan :
                fmt.Println("Timed out")
                return
            case <- ticker_chan :
                fmt.Println("Collecting stats")
                mydb.Stats()
            case res := <- stats:
                fmt.Printf("Completed %v iterations\n", res)
            }
    }
}

