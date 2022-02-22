package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gRPC-Go/PrimeFactors/primepb"
	"google.golang.org/grpc"
)

func (*server) Prime(req *primepb.NumReq, stream primepb.PrimeDecompo_PrimeServer) error {
	fmt.Printf("Prime function was invoked with %v\n", req)
	val := req.Num
	for k := 2; val > 1; {
		if val%int64(k) == 0 {
			response := &primepb.Res{
				Facto: int32(k),
			}
			val = val / int64(k)
			stream.Send(response)
		} else {
			k++
		}
	}

	return nil
}

type server struct {
	primepb.UnimplementedPrimeDecompoServer
}

func main() {
	// Lets create a tcp connection
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	primepb.RegisterPrimeDecompoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to load server %v", err)
	}
}
