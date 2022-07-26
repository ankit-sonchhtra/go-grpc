package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func doPrime(c pb.PrimeCalculatorServiceClient) {
	stream, _ := c.Prime(context.Background(), &pb.PrimeRequest{
		Number: 120,
	})

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receving messages %v", err)
		}

		fmt.Println(msg.Result)
	}

}
