package main

import (
	"fmt"
	"log"

	pb "github.com/trevorrobertsjr/trevor-grpc-go/prime/proto"
	"google.golang.org/grpc"
)

func (s *Server) ThatsJustPrime(in *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	log.Printf("ThatsJustPrime function was invoked with: %v\n", in)

	res := fmt.Sprintf("Prime decomposition of %d: ", in.Number)
	stream.Send(&pb.PrimeResponse{
		Prime: res,
	})
	N := in.Number

	var k int32 = 2 // Start with the smallest prime number

	// Loop until N becomes 1
	for N > 1 {
		if N%k == 0 {
			res := fmt.Sprintf("%d", k)
			stream.Send(&pb.PrimeResponse{
				Prime: res,
			})
			N = N / k // Reduce N by dividing it by k
			if N > 1 {
				res := ", "
				stream.Send(&pb.PrimeResponse{
					Prime: res,
				})
			}
		} else {
			k++ // Increment k to check the next potential factor
		}
	}
	res = "\n"
	stream.Send(&pb.PrimeResponse{
		Prime: res,
	}) // Finish the line

	return nil
}
