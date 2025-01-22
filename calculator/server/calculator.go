package main

import (
	"context"
	"log"

	pb "github.com/trevorrobertsjr/trevor-grpc-go/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)
	return &pb.CalculatorResponse{
		Sum: in.Number1 + in.Number2,
	}, nil
}
