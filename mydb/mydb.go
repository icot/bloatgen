package mydb

import (
    "fmt"
    "time"
    "math/rand"
    "database/sql"
    _ "github.com/lib/pq"
)

var db *sql.DB
var err error

func init() {
    db, err = sql.Open("postgres", "host=/var/run/postgresql dbname=bloat password=bloat")
	if err != nil {
        fmt.Println("Error: ", err)
	}
}

func RandomId(table string) int32 {
    query := fmt.Sprintf("select max(id) from %v", table)
	row, err := db.Query(query)
	defer row.Close()
    if err != nil {
        fmt.Println(err)
    }
    var max_id sql.NullInt64
    row.Next()
    err = row.Scan(&max_id)
    if err != nil {
        panic(err)
    }
    if max_id.Valid {
        return 1 + rand.Int31n(int32(max_id.Int64))
    } else {
        fmt.Println("Empty table")
        return 0
    }
}

func Stats() {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("insert into stats.table_bloat select * from tables_bloat()");
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
    _, err = stmt.Exec()
    if err != nil {
        panic(err)
    }
	err = tx.Commit()
	if err != nil {
	    panic(err)
	}
}

func Insert(table, data string) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
    query := fmt.Sprintf("insert into %v(tstamp, data) values($1, $2)", table)
	stmt, err := tx.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
    _, err = stmt.Exec(time.Now(), data)
    if err != nil {
        panic(err)
    }
	err = tx.Commit()
	if err != nil {
	    panic(err)
	}
}

func Update(table string, id int32, data string) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
    query := fmt.Sprintf("update %v set data = $1 where id = $2", table)
	stmt, err := tx.Prepare(query)
	if err != nil {
        fmt.Println("Error preparing statement")
		panic(err)
	}
	defer stmt.Close()
    _, err = stmt.Exec(data, id)
    if err != nil {
        fmt.Println("Error executing statement")
        panic(err)
    }
	err = tx.Commit()
	if err != nil {
        fmt.Println("Error commiting transaction")
        panic(err)
	}

}

func Delete(table string, id int32) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
    query := fmt.Sprintf("delete from %v where id = $1", table)
	stmt, err := tx.Prepare(query)
	if err != nil {
        fmt.Println("Error preparing statement")
		panic(err)
	}
	defer stmt.Close()
    _, err = stmt.Exec(id)
    if err != nil {
        fmt.Println("Error executing statement")
        panic(err)
    }
	err = tx.Commit()
	if err != nil {
        fmt.Println("Error commiting transaction")
        panic(err)
	}
}

