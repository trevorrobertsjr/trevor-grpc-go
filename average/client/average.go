package main

import (
	"context"
	"log"

	pb "github.com/trevorrobertsjr/trevor-grpc-go/average/proto"
)

func doAverage(c pb.AverageServiceClient) {
	log.Println("doAverage was invoked")

	reqs := []*pb.AverageRequest{
		// {Number: 3},
		// {Number: 4},
		// {Number: 3},
	}
	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending a req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}
	log.Printf("Average: %f\n", res.Result)
}
