package main

import (
	"context"
	"log"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("Client DoGreet Invoked: ")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "John",
	})

	if err != nil {
		log.Fatalf("Failed to call greet : %v\n", err)
	}

	log.Printf("Greeting : %v\n", res.Result)
}
