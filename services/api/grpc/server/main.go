package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"reflect"

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

var period model.Period

func (s *getPeriodService) GetPeriod(ctx context.Context, message *pb.GetPeriodRequest) (*pb.GetPeriodResponse, error) {
	psql_db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	con, err := psql_db.DB()
	defer con.Close()

	result := psql_db.Find(&period)
	fmt.Println(reflect.TypeOf(result.Statement.ReflectValue.FieldByName("Period")))
	reflect_value := result.Statement.ReflectValue
	result_period := reflect_value.FieldByName("Period")
	fmt.Println(reflect.TypeOf(result_period))

	return &pb.GetPeriodResponse{
		Response: &pb.DefaultResponse{
			Status:  1,
			Message: "message",
		},
		Periods: []*pb.Period{
			&pb.Period{
				Id:     1,
				Period: period.Period,
			},
		},
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
