package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Amr-Reda/calculator/calculator_proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to establish connection: %v\n", err)
	}
	defer conn.Close()

	client := calculator_proto.NewCalculatorServiceClient(conn)
	fmt.Println("<== Created client ==>")

	// Unary
	doSum(client)
	// Server Streaming
	doPrimeNumberDecomposition(client)
	// Client Streaming
	doComputeAverage(client)
}

func doSum(client calculator_proto.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
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

func doPrimeNumberDecomposition(client calculator_proto.CalculatorServiceClient) {
	fmt.Println("Starting to do a doPrimeNumberDecomposition Server Streaming RPC...")
	req := &calculator_proto.PrimeNumberDecompositionRequest{
		Num: 12,
	}
	stream, err := client.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while call PrimeNumberDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		// end of stream
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receive PrimeNumberDecomposition RPC: %v", err)
		}

		log.Println("Response from PrimeNumberDecomposition:", res.GetPrimeFactorResult())
	}
}

func doComputeAverage(client calculator_proto.CalculatorServiceClient) {
	fmt.Println("Starting to do a ComputeAverage Server Streaming RPC...")
	stream, err := client.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while call ComputeAverage RPC: %v", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		fmt.Println("Sending number:", number)
		stream.Send(&calculator_proto.ComputeAverageRequest{
			Num: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response fromComputeAverage RPC: %v", err)
	}
	log.Println("Response from ComputeAverage:", res.GetAverageResult())
}
