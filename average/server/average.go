package main

import (
	"io"
	"log"

	pb "github.com/trevorrobertsjr/trevor-grpc-go/average/proto"
)

func (s *Server) Average(stream pb.AverageService_AverageServer) error {

	log.Println("Average function was invoked")
	var sum int32 = 0
	var counter int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if counter <= 0 {
				log.Fatalf("No numbers were entered")

			} else {
				log.Println("Returning the average")
				return stream.SendAndClose(&pb.AverageResponse{
					Result: float64(sum) / float64(counter),
				})
			}
		}
		sum += req.Number
		counter++
		log.Println(sum)
		log.Println(counter)

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
	}

}
