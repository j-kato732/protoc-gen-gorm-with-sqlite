package main

import (
	// "database/sql"
	"fmt"

	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	// postgres接続情報
	conn = "host=db port=5432 user=admin password=password+1 dbname=testdb sslmode=disable TimeZone=Asia/Shanghai"
)

type Periods struct {
	gorm.Model
	Period string `gorm:"size:6`
}

func main() {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Periods{})
	db.Create(&Periods{Period: "202105"})

	fmt.Println("Successfullty connected!")

	// db.HasTable(&Periods{})
	// db.HasTable("periods")

	// INSERT
}
