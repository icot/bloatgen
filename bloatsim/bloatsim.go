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

func activity(c chan int32, inserts, deletes, updates float32) int32 {
    var counter int32  = 0
    for {
        r := rand.Float32()
        data := Data()
        switch {
            case r <= inserts:
                mydb.Insert("t_char", data)
                mydb.Insert("t_text", data)
                mydb.Insert("t_varchar", data)
            case r > inserts && r <= inserts + deletes:
                id := mydb.RandomId("t_char")
                mydb.Delete("t_char", id)
                mydb.Delete("t_text", id)
                mydb.Delete("t_varchar", id)
            default:
                id := mydb.RandomId("t_char")
                mydb.Update("t_char", id, data)
                mydb.Update("t_text", id, data)
                mydb.Update("t_varchar", id, data)
        }
        if counter >= 1000 {
            c <- counter
            counter = 0
        }
        counter++
    }
}

func Simulate(duration int64, inserts, deletes, updates float32) {
    c:= make(chan int32, 1)
    go activity(c, inserts, deletes, updates)
    select {
        case <- time.After(time.Duration(duration) * time.Second):
            fmt.Println("Timed out")
        case res := <- c:
            fmt.Printf("Completed %v iterations", res)
            mydb.Stats()
    }
}

