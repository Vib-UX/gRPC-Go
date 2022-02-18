package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gRPC-Go/greet/greetpb"
	"google.golang.org/grpc"
)

// func (s *server) mustEmbedUnimplementedGreetServiceServer() {}
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet fxn was invoked %v\n", req)
	// Extracting data from the request
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName // declares to be of tyope string as required in the Respone struct

	// Creating our response message
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

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
