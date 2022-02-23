package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gRPC-Go/MaxAPI--Bidirectional_Streaming/maxpb"
	"google.golang.org/grpc"
)

// Lets implement biDi streaming api call from client side
func doBiDiStreaming(client maxpb.APIClient) {

	fmt.Println("Starting to do a BiDi Streaming RPC...")

	stream, err := client.FindMax(context.Background())
	if err != nil {
		log.Fatalf("Error while creating the stream %v\n", err)
	}
	requests := []*maxpb.Req{
		{
			Num: 1,
		},
		{
			Num: 5,
		},
		{
			Num: 3,
		},
		{
			Num: 6,
		},
		{
			Num: 2,
		},
		{
			Num: 20,
		},
	}

	waitc := make(chan struct{}) //wait channel go routine
	go func() {
		//Function to send multiple req to the server
		for _, req := range requests {
			fmt.Printf("Sending the Req with val: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// Receiving stream messages from the server
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error encountered while receiving the stream %v\n", err)
				break
			}
			log.Printf("%v\n", msg.GetMax())

		}
		close(waitc)
	}()

	<-waitc

}

func main() {
	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect  %v\n", err)
	}
	// defer here will close the client connection once main block is done
	defer clientConnection.Close()
	client := maxpb.NewAPIClient(clientConnection)
	doBiDiStreaming(client)
}
