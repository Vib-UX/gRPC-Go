package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/gRPC-Go/greet/greetpb"
	"google.golang.org/grpc"
)

func doServerStreaming(client greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vibhav",
			LastName:  "Sharma",
		},
	}
	// Client will call our greet many times endpoint
	resStream, err := client.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling the GreetManyTimes rpc: %v\n", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// Its the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming: %v\n", err)
		}

		log.Println(msg.GetResult())
	}

}

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

	// doUnary(client)
	doServerStreaming(client)

}
