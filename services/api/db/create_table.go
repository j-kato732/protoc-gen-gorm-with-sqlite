package main

import (
	"fmt"
	// "reflect"

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
	v2_db, err := db.DB()
	defer v2_db.Close()

	// テーブル存在確認
	b := db.Migrator().HasTable("periods")
	if b == false {
		db.AutoMigrate(periods)
	}

	isExist := db.Migrator().HasTable("userInfos")
	if isExist == false {
		db.AutoMigrate(userInfos)
	}

	// db.Create(&model.Period{Period: "202111"})

	// result := db.Model(&period).First(&period)
	if result := db.Find(&periods); result.Error != nil {
		fmt.Println(err)
	}

	// var response_periods []*pb.Period

	// for _, period := range periods {
	// 	fmt.Println(period.Period)
	// 	fmt.Println(period.ID)
	// 	response_periods = append(response_periods, &pb.Period{
	// 		Id:     int32(period.ID),
	// 		Period: period.Period,
	// 	})
	// // }

	// // fmt.Println(periods)
	// fmt.Println(response_periods)
	// fmt.Println(period.Period)
	// fmt.Println(result.RowsAffected)
	// fmt.Println(result.Statement.Context)
	// fmt.Println(result.Statement.TableExpr)
	// fmt.Println(result.Statement.Table)
	// fmt.Println(result.Statement.Model)
	// // fmt.Println(result.Statement.ReflectValue.Index(0))
	// fmt.Println(result.Statement.ReflectValue)
	// fmt.Println(result.Statement.ReflectValue.Field(0).FieldByIndex([]int{0}))
	// fmt.Println(result.Statement.ReflectValue.FieldByName("Period"))
	// fmt.Println(result.Statement.ReflectValue.FieldByName("Period").Kind())
	// fmt.Printf("(%%#v) %#v\n", periods)

	fmt.Println("Success connected")
}
