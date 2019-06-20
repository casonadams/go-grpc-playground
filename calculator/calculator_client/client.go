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
	add(c)
	subtract(c)
}

func add(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary Add RPC...")
	add_req := &calculatorpb.AddRequest{
		Add: &calculatorpb.Add{
			A: 3,
			B: 10,
		},
	}

	add, err := c.Add(context.Background(), add_req)
	if err != nil {
		log.Fatalf("Error while calling Calcualtor Add RPC: %v\n", err)
	}
	log.Printf("Response from Calculate Add: %v\n", add.Result)
}

func subtract(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary Subtract RPC...")
	subtract_req := &calculatorpb.SubtractRequest{
		Subtract: &calculatorpb.Subtract{
			A: 10,
			B: 3,
		},
	}

	subtract, err := c.Subtract(context.Background(), subtract_req)
	if err != nil {
		log.Fatalf("Error while calling Calculator Subtract RPC: %v\n", err)
	}
	log.Printf("Response from Calculate Subtract: %v\n", subtract.Result)
}
