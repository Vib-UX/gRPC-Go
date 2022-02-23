package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/gRPC-Go/MaxAPI--Bidirectional_Streaming/maxpb"
	"google.golang.org/grpc"
)

type server struct {
	maxpb.UnimplementedAPIServer
}

func (*server) FindMax(stream maxpb.API_FindMaxServer) error {
	fmt.Printf("Max API function was invoked with streaming req \n")
	max := math.MinInt
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving the stream from client %v", err)
		}
		if req.GetNum() > int32(max) {
			max = int(req.GetNum()) // Updates the current max
			if resErr := stream.Send(&maxpb.Res{
				Max: int32(max),
			}); resErr != nil {
				log.Fatalf("%v", resErr)
				return resErr
			}
		}

	}
}

func main() {
	//lets create a tcp server connection
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error listening to the client %v\n", err)
	}

	s := grpc.NewServer()
	maxpb.RegisterAPIServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to load server %v\n", err)
	}
}
