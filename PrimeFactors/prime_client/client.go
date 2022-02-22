package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/gRPC-Go/PrimeFactors/primepb"
	"google.golang.org/grpc"
)

func doServerStream(client primepb.PrimeDecompoClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &primepb.NumReq{
		Num: 120,
	}
	// Prime endpoint
	readstream, err := client.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling the Prime rpc: %v\n", err)
	}

	for {
		msg, err := readstream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v\n", err)
		}
		log.Printf("%v ", msg.GetFacto())
	}

}

func main() {
	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	defer clientConnection.Close()

	client := primepb.NewPrimeDecompoClient(clientConnection)

	doServerStream(client)

}
