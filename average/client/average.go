package main

import (
	"context"
	"log"

	pb "github.com/trevorrobertsjr/trevor-grpc-go/average/proto"
)

func doAverage(c pb.AverageServiceClient, numbers []int32) {
	log.Println("doAverage was invoked")

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	// Convert numbers to gRPC request messages and send them
	for _, num := range numbers {
		req := &pb.AverageRequest{Number: num}
		log.Printf("Sending request: %v\n", req)
		stream.Send(req)
	}

	// Receive and print the response
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}
	log.Printf("Average: %f\n", res.Result)
}
