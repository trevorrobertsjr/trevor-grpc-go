package main

import (
	"context"
	"log"

	pb "github.com/trevorrobertsjr/trevor-grpc-go/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient) {
	var num1 uint32 = 5
	var num2 uint32 = 7
	log.Printf("doCalculator was invoked to add %d and %d\n", num1, num2)
	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		Number1: num1,
		Number2: num2,
	})
	if err != nil {
		log.Fatalf("Could not Calculate %v\n", err)

	}
	log.Printf("Calculatoring: %d\n", res.Sum)
}
