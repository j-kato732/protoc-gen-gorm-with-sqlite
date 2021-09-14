package main

import (
	"context"
	"fmt"
	"reflect"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	pb "grpc_gateway_sample/proto"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./db/test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	con, err := db.DB()
	defer con.Close()

	var periods []pb.PeriodORM

	// select * from Periods;
	if err := db.Find(&periods).Error; err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	result, err := periods[0].ToPB(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reflect.TypeOf(result))
	fmt.Println(result.GetId())

	var userInfo pb.UserInfoORM
	// if err = db.Where("user_id = ? AND period = ?", "2", "202105").Find(&userInfo).Error; err != nil {
	// 	fmt.Println(err)
	// }
	if err = db.Where(pb.UserInfoORM{UserId: 1, Period: "202105"}).Find(&userInfo).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(userInfo)
}
