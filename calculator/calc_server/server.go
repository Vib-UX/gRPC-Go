package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gRPC-Go/calculator/calcpb"
	"google.golang.org/grpc"
)

func (*server) Add(ctx context.Context, req *calcpb.SumReq) (*calcpb.SumRes, error) {
	fmt.Printf("Add fxn was invoked %v\n", req)
	// sum of a and b stored in the result
	result := req.GetA() + req.GetB()
	response := &calcpb.SumRes{
		Res: result,
	}
	return response, nil
}

type server struct {
	calcpb.UnimplementedCalcServer
}

func main() {
	// Lets first listen to the tcp connection
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	newServer := grpc.NewServer()
	calcpb.RegisterCalcServer(newServer, &server{})

	if err := newServer.Serve(lis); err != nil {
		log.Fatalf("Failed to load server %v", err)
	}
}
