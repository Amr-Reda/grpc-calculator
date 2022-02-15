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
	fmt.Printf("Recieved Sum RPC\n")
	result := req.GetNum1() + req.GetNum2()
	res := &calculator_proto.SumResponse{
		Result: result,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculator_proto.PrimeNumberDecompositionRequest, stream calculator_proto.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Println("Recieved PrimeNumberDecomposition RPC")
	num := req.GetNum()
	divisor := int64(2)
	for num > 1 {
		if num%divisor == 0 {
			stream.Send(&calculator_proto.PrimeNumberDecompositionResponse{
				PrimeFactorResult: divisor,
			})
			num = num / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
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
