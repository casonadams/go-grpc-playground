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

func (*server) Calculator(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Calculate function was invoked with %v\n", req)
	a := req.GetCalculator().GetA()
	b := req.GetCalculator().GetB()
	result := a + b
	res := &calculatorpb.CalculatorResponse{
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
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}
