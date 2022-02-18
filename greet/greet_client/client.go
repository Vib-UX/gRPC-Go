package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gRPC-Go/greet/greetpb"
	"google.golang.org/grpc"
)

func doUnary(client greetpb.GreetServiceClient) {
	fmt.Printf("Unary rpc %v", client)

	// Lest create the request given as per the type
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vibhav",
			LastName:  "Sharma",
		},
	}
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet rpc: %v", err)
	}
	log.Printf("Response from the greet RPC %v", res)

}

func main() {
	//WithInsecure() not to be opened in production
	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}
	// defer here will close the client connection once main block is done
	defer clientConnection.Close()
	client := greetpb.NewGreetServiceClient(clientConnection)

	doUnary(client)

}
