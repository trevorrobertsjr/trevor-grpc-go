package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	pb "github.com/trevorrobertsjr/trevor-grpc-go/prime/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	// Parse the command-line parameter
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -number=<number>\n", os.Args[0])
		flag.PrintDefaults()
	}
	numberFlag := flag.String("number", "", "The number to decompose into primes")
	flag.Parse()

	if *numberFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	number, err := strconv.Atoi(*numberFlag)
	if err != nil {
		log.Fatalf("Invalid number parameter: %v\n", err)
	}

	// Ensure number fits within int32
	if number > int(^uint32(0)>>1) || number < 1 {
		log.Fatalf("Number must be a positive int32 value.\n")
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer conn.Close()
	c := pb.NewThatsJustPrimeClient(conn)

	doThatsJustPrime(c, int32(number))
}
