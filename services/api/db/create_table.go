package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// postgres接続情報
	conn = "host=db port=5432 user=admin password=password+1 dbname=testdb sslmode=disable"
)

type PERIODS struct {
	ID     int32
	PERIOD string
}

func main() {
	db, err := sql.Open("postgres", conn)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfullty connected!")

	// INSERT
}
