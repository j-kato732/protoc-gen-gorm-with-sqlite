package main

import (
	"context"
	"fmt"
	"log"
	"net"
	// "reflect"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// db "grpc_gateway_sample/db"
	"grpc_gateway_sample/db/model"
	pb "grpc_gateway_sample/proto"
)

const (
	conn = "host=db port=5432 user=admin password=password+1 dbname=testdb sslmode=disable TimeZone=Asia/Shanghai"
	port = ":8080"
)

type getPeriodService struct {
	pb.UnimplementedAimoServer
}

var (
	periods          []model.Period
	response_status  int32
	response_message string
)

func (s *getPeriodService) GetPeriod(ctx context.Context, message *pb.GetPeriodRequest) (*pb.GetPeriodResponse, error) {
	psql_db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	con, err := psql_db.DB()
	defer con.Close()

	// SELECT * from period;
	if err := psql_db.Find(&periods).Error; err != nil {
		log.Println(err)
		return nil, err
	} else {
		response_status = 1
		response_message = ""
	}

	var response_periods []*pb.Period
	for _, period := range periods {
		response_periods = append(response_periods, &pb.Period{
			Id:     int32(period.ID),
			Period: period.Period,
		})
	}

	return &pb.GetPeriodResponse{
		Response: &pb.DefaultResponse{
			Status:  response_status,
			Message: response_message,
		},
		Result: &pb.Result{
			Period: response_periods,
		},
	}, nil
}

func (s *getPeriodService) GetUserInfo(ctx context.Context, message *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	var userInfo model.UserInfo

	psql_db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	con, err := psql_db.DB()
	defer con.Close()

	// var response_status int32
	// var response_message string

	isExist := psql_db.Migrator().HasTable("userInfos")
	if isExist == false {
		psql_db.AutoMigrate(userInfo)
	}

	if err = psql_db.Where(model.UserInfo{
		UserId: 1,
		Period: "202105",
	}).Find(&userInfo).Error; err != nil {
		fmt.Println(err)
	} else {
		response_status = 1
		response_message = ""
	}

	// var response *pb.GetUserInfoResponse

	return &pb.GetUserInfoResponse{
		Response: &pb.DefaultResponse{
			Status:  response_status,
			Message: response_message,
		},
		Result: &pb.GetUserInfoResult{
			UserInfo: &pb.UserInfo{
				UserInfoId:    int32(userInfo.ID),
				UserId:        userInfo.UserId,
				LastName:      userInfo.LastName,
				FirstName:     userInfo.FirstName,
				Period:        userInfo.Period,
				DepartmentId:  userInfo.DepartmentId,
				JobId:         userInfo.JobId,
				EnrollmentFlg: userInfo.EnrollmentFlg,
				AdminFlg:      userInfo.AdminFlg,
			},
		},
	}, nil
	// SELECT * FROM userInfo where user_id = ? and period = ?;
	// if err := psql_db.Find(&)
	// if err := db.Where("user_id = ? AND period = ?", "jinzhu", "22").Find(&users)
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
