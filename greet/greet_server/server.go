package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/gRPC-Go/greet/greetpb"
	"google.golang.org/grpc"
)

// Greeting function endpoint handled by server
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

// Greeting many times endpoint handled by the server
func (*server) GreetManyTimes(req *greetpb.GreetManyRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	for i := 0; i < 10; i++ {
		result := "Hello " + req.GetGreeting().GetFirstName() + " " + req.Greeting.GetLastName() + " time " + strconv.Itoa(i)
		response := &greetpb.GreetManyResponse{
			Result: result,
		}

		stream.Send(response)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

// GreetEveryone handled by the server
func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function was invoked with streaming req \n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error in receiving the stream %v\n", err)
		}
		result := "Hello " + req.GetGreeting().GetFirstName() + " !"
		if recErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		}); recErr != nil {
			log.Fatalf("Error while sending the data to client %v\n", recErr)
			return recErr
		}
	}
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
