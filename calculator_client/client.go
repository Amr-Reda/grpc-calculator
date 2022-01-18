package main

import (
	"fmt"
	"log"

	"github.com/Amr-Reda/calculator/calculator_proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello world")

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to establish connection: %v", err)
	}
	defer conn.Close()

	client := calculator_proto.NewCalculatorServiceClient(conn)
	fmt.Printf("Created client: %f", client)
}
