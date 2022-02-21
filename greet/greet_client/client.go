package main

import (
	"fmt"
	"log"

	"github.com/gRPC-Go/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	//WithInsecure() not to be opened in production
	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}
	defer clientConnection.Close()
	client := greetpb.NewGreetServiceClient(clientConnection)
	fmt.Printf("Created client! %f", client)
}
