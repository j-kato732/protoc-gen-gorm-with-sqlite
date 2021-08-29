package main

import (
	"context"
	"fmt"
	"log"
	
	pb "aimo_test/api/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewAimoClient(conn)
	res, err := client.GetPeriod(context.TODO(), &pb.GetPeriodRequest{})
	fmt.Printf("Result:%#v \n", res.Response.GetStatus())
	fmt.Printf("Result:%#v \n", res.Response.GetMessage())
	fmt.Printf("Result:%#v \n", res.Periods[0].GetId())
	fmt.Printf("Result:%#v \n", res.Periods[0].GetPeriod())
	if err != nil {
		fmt.Printf("error:%#v \n", err)
	}
}