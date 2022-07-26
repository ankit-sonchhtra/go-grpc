package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {

	req := &pb.GreetRequest{
		FirstName: "Ankit",
	}

	stream, err := c.GreetMany(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to get stream response.")
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}

		log.Println(msg.Result)

	}

}
