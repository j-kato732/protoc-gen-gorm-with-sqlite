package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"grpc_gateway_sample/db/model"
)

const (
	conn = "host=db port=5432 user=admin password=password+1 dbname=testdb sslmode=disable TimeZone=Asia/Shanghai"
)

var period model.Period

func main() {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	v2_db, err := db.DB()
	defer v2_db.Close()

	// テーブル存在確認
	b := db.Migrator().HasTable("periods")
	if b == false {
		db.AutoMigrate(period)
	}

	// db.Create(model.Period{Period: "202105"})

	result := db.Model(&period).First(&period)
	fmt.Println(result.RowsAffected)
	fmt.Printf("(%%#v) %#v\n", result)

	fmt.Println("Success connected")
}
