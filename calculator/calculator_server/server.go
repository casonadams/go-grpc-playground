package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/casonadams/go-grpc-playground/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	fmt.Printf("Calculate function was invoked with %v\n", req)
	a := req.GetAdd().GetA()
	b := req.GetAdd().GetB()
	result := a + b
	res := &calculatorpb.AddResponse{
		Result: result,
	}
	return res, nil

}
func main() {
	fmt.Println("Server listening on localhost:50051")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterAddServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}
