package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ankit-sonchhtra/go-grpc/proto"
)

func doAvg(c pb.AvgCalculatorServiceClient) {
	reqs := []*pb.AvgRequest{
		{Number: 2},
		{Number: 4},
		{Number: 6},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Failed to get steream %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while close the stream : %v", err)
	}

	fmt.Println("Average of requested numbers : ", res.Result)
}
