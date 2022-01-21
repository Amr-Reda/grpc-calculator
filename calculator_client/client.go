package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Amr-Reda/calculator/calculator_proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello world")

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to establish connection: %v\n", err)
	}
	defer conn.Close()

	client := calculator_proto.NewCalculatorServiceClient(conn)
	fmt.Println("<== Created client ==>")

	doSum(client)
}

func doSum(client calculator_proto.CalculatorServiceClient) {
	req := &calculator_proto.SumRequest{
		Num1: 1,
		Num2: 2,
	}
	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while call Sum RPC: %v", err)
	}
	log.Println("Response from Sum:", res.Result)
	// log.Printf("Response from Sum: %v", res.Result)
}
