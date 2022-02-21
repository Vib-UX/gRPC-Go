package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gRPC-Go/greet/greetpb"
	"google.golang.org/grpc"
)

func (s *server) mustEmbedUnimplementedGreetServiceServer() {}

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func main() {
	fmt.Println("Hello world!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to load server %v", err)
	}

}
