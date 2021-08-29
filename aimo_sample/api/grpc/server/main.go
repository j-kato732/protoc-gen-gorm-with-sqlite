package main

import (
	"context"
	"net"
	"log"

	pb "aimo_test/api/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type getPeriodService struct {
	pb.UnimplementedAimoServer
}

func (s *getPeriodService) GetPeriod(ctx context.Context, message *pb.GetPeriodRequest) (*pb.GetPeriodResponse, error){
	return &pb.GetPeriodResponse{
		Response: &pb.DefaultResponse{
			Status:	1,
			Message: "message",
		},
		Periods: []*pb.Period{
			&pb.Period{
				Id: 1,
				Period: "202105",
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