package main

import pb "github.com/trevorrobertsjr/trevor-grpc-go/average/proto"

var addr = "0.0.0.0:50051"

type Server struct {
	pb.AverageServiceServer
}

func main() {

}
