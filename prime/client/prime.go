package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/trevorrobertsjr/trevor-grpc-go/prime/proto"
)

func doThatsJustPrime(c pb.ThatsJustPrimeClient, number int32) {
	log.Println("doThatsJustPrime was invoked")
	req := &pb.PrimeRequest{
		Number: number,
	}

	stream, err := c.ThatsJustPrime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling ThatsJustPrime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		fmt.Print(msg.Prime)

		// log.Print(msg.Prime)
	}
}
