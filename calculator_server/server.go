package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Amr-Reda/calculator/calculator_proto"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculator_proto.SumRequest) (*calculator_proto.SumResponse, error) {
	result := req.GetNum1() + req.GetNum2()
	res := &calculator_proto.SumResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to list %v", err)
	}
	fmt.Println("App running on port: 50051")

	s := grpc.NewServer()
	calculator_proto.RegisterCalculatorServiceServer(s, &server{})

	errSrv := s.Serve(lis)
	if errSrv != nil {
		log.Fatalf("Failed to serve: %v", errSrv)
	}
}
