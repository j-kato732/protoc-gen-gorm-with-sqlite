package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	pb "grpc_gateway_sample/proto"
)

var (
	periods   pb.PeriodORM
	userInfos pb.UserInfoORM
)

func main() {
	db, err := gorm.Open(sqlite.Open("./db/test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	con, err := db.DB()
	defer con.Close()

	if err := db.Create(&pb.PeriodORM{Period: "202111"}).Error; err != nil {
		log.Println(err)
	}
}
