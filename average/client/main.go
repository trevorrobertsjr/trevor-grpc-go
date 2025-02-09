package main

import (
	"log"
	"os"
	"strconv"

	pb "github.com/trevorrobertsjr/trevor-grpc-go/average/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	// Parse numbers from command-line arguments
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <num1> <num2> ... <numN>\n", os.Args[0])
	}

	var numbers []int32
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("Invalid number %s: %v", arg, err)
		}
		numbers = append(numbers, int32(num))
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewAverageServiceClient(conn)
	doAverage(c, numbers)
}
