package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	// "reflect"

	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// db "grpc_gateway_sample/db"
	db "grpc_gateway_sample/action"
	pb "grpc_gateway_sample/proto"
)

const (
	db_path = "./db/test.db"
	port    = ":8080"
)

type getPeriodService struct {
	pb.UnimplementedAimoServer
}

var (
	periods_orm      []pb.PeriodORM
	response_status  int32
	response_message string
)

func (s *getPeriodService) GetPeriod(ctx context.Context, message *pb.GetPeriodRequest) (*pb.GetPeriodResponse, error) {
	result, err := db.GetPeriod(ctx)
	if err != nil {
		response_status = 255
		response_message = "faild GetPeriod"
	}

	return &pb.GetPeriodResponse{
		Response: &pb.DefaultResponse{
			Status:  response_status,
			Message: response_message,
		},
		Result: &pb.Result{
			Period: result,
		},
	}, nil
}

func (s *getPeriodService) GetUserInfo(ctx context.Context, message *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	if message.UserId == 0 || len(message.Period) == 0 {
		return nil, errors.New("invalid params")
	}

	var userInfo pb.UserInfoORM

	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	con, err := db.DB()
	defer con.Close()

	// var response_status int32
	// var response_message string

	isExist := db.Migrator().HasTable(userInfo.TableName())
	if isExist == false {
		db.AutoMigrate(userInfo)
	}

	result := db.Where(pb.UserInfoORM{
		UserId: message.UserId,
		Period: message.Period,
	}).Find(&userInfo)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	if result.RowsAffected == 0 {
		response_status = 10
		response_message = "userId 1 was not found"
	} else {
		response_status = 1
		response_message = ""
	}

	response_userInfo, _ := userInfo.ToPB(ctx)
	// var response *pb.GetUserInfoResponse

	return &pb.GetUserInfoResponse{
		Response: &pb.DefaultResponse{
			Status:  response_status,
			Message: response_message,
		},
		Result: &pb.GetUserInfoResult{
			UserInfo: &response_userInfo,
		},
	}, nil
	// SELECT * FROM userInfo where user_id = ? and period = ?;
	// if err := psql_db.Find(&)
	// if err := db.Where("user_id = ? AND period = ?", "jinzhu", "22").Find(&users)
}

func (s *getPeriodService) PostUserInfo(ctx context.Context, request *pb.UserInfo) (*pb.PostUserInfoResponse, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	con, err := db.DB()
	defer con.Close()

	request_orm, err := request.ToORM(ctx)
	if err != nil {
		fmt.Println(err)
	}

	result := db.Create(&request_orm)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	if result.RowsAffected == 0 {
		response_status = 10
		response_message = "failed create"
	} else {
		response_status = 1
		response_message = ""
	}

	return &pb.PostUserInfoResponse{
		Response: &pb.DefaultResponse{
			Status:  response_status,
			Message: response_message,
		},
		Result: &pb.PostUserInfoResponsePostUserInfoResult{
			UserId: request.Id,
		},
	}, nil
}

func (s *getPeriodService) UpdateUserInfo(ctx context.Context, request *pb.UserInfo) (*pb.DefaultResponse, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	con, err := db.DB()
	defer con.Close()

	request_orm, err := request.ToORM(ctx)
	if err != nil {
		fmt.Println(err)
	}

	result := db.Model(&request_orm).Updates(request_orm)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		response_status = 1
		response_message = ""
	}

	return &pb.DefaultResponse{
		Status:  response_status,
		Message: response_message,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v¥n", err)
	}
	server := grpc.NewServer()
	// 実行したい実処理をserverに登録する
	// periodService := &getPeriodService{}
	// pb.RegisterAimoServer(server, periodService)
	pb.RegisterAimoServer(server, &getPeriodService{})
	log.Printf("server listening at %v\n", lis.Addr())
	if err != nil {
		log.Fatalf("faild to serve: %v\n", err)
	}
	server.Serve(lis)
}
