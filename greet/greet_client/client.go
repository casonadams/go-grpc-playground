package main

import (
	"fmt"
	"log"

	"github.com/casonadams/go-grpc-playground/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm the client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	// This gets closed when things are done
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created Client: %f", c)
}
