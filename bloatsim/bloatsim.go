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
                mydb.Delete("t_char")
                mydb.Delete("t_text")
                mydb.Delete("t_varchar")
            case r > inserts + deletes && r <= inserts + deletes + updates:
                mydb.Update("t_char", data)
                mydb.Update("t_text", data)
                mydb.Update("t_varchar", data)
            default:
                fmt.Println("Select")
        }
        if counter > 10 {
            c <- counter
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
    }
}

