package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gRPC-Go/calculator/calcpb"
	"google.golang.org/grpc"
)

func doUnary(client calcpb.CalcClient) {
	fmt.Printf("Unary rpc established %v\n", client)

	//Lets create the request to the server
	req := &calcpb.SumReq{
		A: 2,
		B: 3,
	}

	// Lets receive the response
	res, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling add rpc %v", err)
	}
	log.Printf("Response from the add rpc : = %v", res)

}

func main() {
	// Lets connect to the server
	clientConnection, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Coudn't connect to the server %v", err)
	}

	defer clientConnection.Close()

	client := calcpb.NewCalcClient(clientConnection)

	doUnary(client)
}
