package main

import (
	"context"
	"fmt"
	"log"

	"github.com/casonadams/go-grpc-playground/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm the client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v\n", err)
	}

	// Make sure to close the connection
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary RPC...")
	req := &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{
			A: 3,
			B: 10,
		},
	}

	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calcualtor RPC: %v\n", err)
	}
	log.Printf("Response from Calculate: %v\n", res.Result)

}
