package main

import (
	"context"
	"log"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("Client doSum Invoked: ")
	res, err := c.Sum(context.Background(), &pb.Request{
		FirstNumber:  10,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Failed to get sum : %v\n", err)
	}
	log.Printf("Sum : %v\n", res.Result)
}
